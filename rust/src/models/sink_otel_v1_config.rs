/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 *
 * Generated by: https://openapi-generator.tech
 */

#[allow(unused_imports)]
use crate::models;
#[allow(unused_imports)]
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SinkOtelV1Config {
    #[serde(rename = "url")]
    pub url: String,
}

impl SinkOtelV1Config {
    pub fn new(url: String) -> SinkOtelV1Config {
        SinkOtelV1Config { url }
    }
}
