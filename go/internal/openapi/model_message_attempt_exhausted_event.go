/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the MessageAttemptExhaustedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageAttemptExhaustedEvent{}

// MessageAttemptExhaustedEvent Sent when a message delivery has failed (all of the retry attempts have been exhausted).
type MessageAttemptExhaustedEvent struct {
	Data MessageAttemptExhaustedEventData `json:"data"`
	Type string                           `json:"type"`
}

type _MessageAttemptExhaustedEvent MessageAttemptExhaustedEvent

// NewMessageAttemptExhaustedEvent instantiates a new MessageAttemptExhaustedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageAttemptExhaustedEvent(data MessageAttemptExhaustedEventData, type_ string) *MessageAttemptExhaustedEvent {
	this := MessageAttemptExhaustedEvent{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewMessageAttemptExhaustedEventWithDefaults instantiates a new MessageAttemptExhaustedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageAttemptExhaustedEventWithDefaults() *MessageAttemptExhaustedEvent {
	this := MessageAttemptExhaustedEvent{}
	var type_ string = "message.attempt.exhausted"
	this.Type = type_
	return &this
}

// GetData returns the Data field value
func (o *MessageAttemptExhaustedEvent) GetData() MessageAttemptExhaustedEventData {
	if o == nil {
		var ret MessageAttemptExhaustedEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEvent) GetDataOk() (*MessageAttemptExhaustedEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MessageAttemptExhaustedEvent) SetData(v MessageAttemptExhaustedEventData) {
	o.Data = v
}

// GetType returns the Type field value
func (o *MessageAttemptExhaustedEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageAttemptExhaustedEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageAttemptExhaustedEvent) SetType(v string) {
	o.Type = v
}

func (o MessageAttemptExhaustedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageAttemptExhaustedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *MessageAttemptExhaustedEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMessageAttemptExhaustedEvent := _MessageAttemptExhaustedEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageAttemptExhaustedEvent)

	if err != nil {
		return err
	}

	*o = MessageAttemptExhaustedEvent(varMessageAttemptExhaustedEvent)

	return err
}

type NullableMessageAttemptExhaustedEvent struct {
	value *MessageAttemptExhaustedEvent
	isSet bool
}

func (v NullableMessageAttemptExhaustedEvent) Get() *MessageAttemptExhaustedEvent {
	return v.value
}

func (v *NullableMessageAttemptExhaustedEvent) Set(val *MessageAttemptExhaustedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageAttemptExhaustedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageAttemptExhaustedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageAttemptExhaustedEvent(val *MessageAttemptExhaustedEvent) *NullableMessageAttemptExhaustedEvent {
	return &NullableMessageAttemptExhaustedEvent{value: val, isSet: true}
}

func (v NullableMessageAttemptExhaustedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageAttemptExhaustedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
