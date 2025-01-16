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
pub struct EventStreamOut {
    #[serde(rename = "data")]
    pub data: Vec<models::EventOut>,
    #[serde(rename = "done")]
    pub done: bool,
    #[serde(rename = "iterator")]
    pub iterator: String,
}

impl EventStreamOut {
    pub fn new(data: Vec<models::EventOut>, done: bool, iterator: String) -> EventStreamOut {
        EventStreamOut {
            data,
            done,
            iterator,
        }
    }
}
