use rdkafka::{consumer::StreamConsumer, error::KafkaResult, ClientConfig};
use serde::Deserialize;
use svix_bridge_types::{SenderInput, SenderOutputOpts, TransformationConfig};

use crate::{input::KafkaConsumer, Result};

#[derive(Clone, Deserialize)]
pub struct KafkaInputOpts {
    /// Comma-separated list of addresses.
    ///
    /// Example: `localhost:9094`
    #[serde(rename = "kafka_bootstrap_brokers")]
    pub bootstrap_brokers: String,

    /// The consumer group ID, used to track the stream offset between restarts
    /// (due to host maintenance, upgrades, crashes, etc.).
    #[serde(rename = "kafka_group_id")]
    pub group_id: String,

    /// The topic to listen to.
    #[serde(rename = "kafka_topic")]
    pub topic: String,

    /// The value for 'security.protocol' in the kafka config.
    #[serde(flatten)]
    pub security_protocol: KafkaSecurityProtocol,

    /// The 'debug' config value for rdkafka - enables more verbose logging
    /// for the selected 'contexts'
    #[serde(rename = "kafka_debug_contexts")]
    pub debug_contexts: Option<String>,
}

impl KafkaInputOpts {
    pub(crate) fn create_consumer(self) -> KafkaResult<StreamConsumer> {
        let mut config = ClientConfig::new();
        config
            .set("group.id", self.group_id)
            .set("bootstrap.servers", self.bootstrap_brokers)
            // messages are committed manually after webhook delivery was successful.
            .set("enable.auto.commit", "false");

        match self.security_protocol {
            KafkaSecurityProtocol::Plaintext => {
                config.set("security.protocol", "plaintext");
            }
            KafkaSecurityProtocol::Ssl => {
                config.set("security.protocol", "ssl");
            }
            KafkaSecurityProtocol::SaslSsl {
                sasl_username,
                sasl_password,
            } => {
                config
                    .set("security.protocol", "sasl_ssl")
                    .set("sasl.mechanisms", "SCRAM-SHA-512")
                    .set("sasl.username", sasl_username)
                    .set("sasl.password", sasl_password);
            }
        }

        if let Some(debug_contexts) = self.debug_contexts {
            if !debug_contexts.is_empty() {
                config.set("debug", debug_contexts);
            }
        }

        config.create()
    }
}

#[derive(Clone, Debug, Deserialize)]
#[serde(tag = "kafka_security_protocol", rename_all = "snake_case")]
pub enum KafkaSecurityProtocol {
    Plaintext,
    Ssl,
    SaslSsl {
        #[serde(rename = "kafka_sasl_username")]
        sasl_username: String,
        #[serde(rename = "kafka_sasl_password")]
        sasl_password: String,
    },
}

pub fn into_sender_input(
    name: String,
    opts: KafkaInputOpts,
    transformation: Option<TransformationConfig>,
    output: SenderOutputOpts,
) -> Result<Box<dyn SenderInput>> {
    Ok(Box::new(KafkaConsumer::new(
        name,
        opts,
        transformation,
        output,
    )?))
}
