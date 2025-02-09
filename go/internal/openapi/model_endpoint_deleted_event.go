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

// checks if the EndpointDeletedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointDeletedEvent{}

// EndpointDeletedEvent Sent when an endpoint is deleted.
type EndpointDeletedEvent struct {
	Data EndpointDeletedEventData `json:"data"`
	Type string                   `json:"type"`
}

type _EndpointDeletedEvent EndpointDeletedEvent

// NewEndpointDeletedEvent instantiates a new EndpointDeletedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointDeletedEvent(data EndpointDeletedEventData, type_ string) *EndpointDeletedEvent {
	this := EndpointDeletedEvent{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewEndpointDeletedEventWithDefaults instantiates a new EndpointDeletedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointDeletedEventWithDefaults() *EndpointDeletedEvent {
	this := EndpointDeletedEvent{}
	var type_ string = "endpoint.deleted"
	this.Type = type_
	return &this
}

// GetData returns the Data field value
func (o *EndpointDeletedEvent) GetData() EndpointDeletedEventData {
	if o == nil {
		var ret EndpointDeletedEventData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EndpointDeletedEvent) GetDataOk() (*EndpointDeletedEventData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EndpointDeletedEvent) SetData(v EndpointDeletedEventData) {
	o.Data = v
}

// GetType returns the Type field value
func (o *EndpointDeletedEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndpointDeletedEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EndpointDeletedEvent) SetType(v string) {
	o.Type = v
}

func (o EndpointDeletedEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointDeletedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *EndpointDeletedEvent) UnmarshalJSON(data []byte) (err error) {
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

	varEndpointDeletedEvent := _EndpointDeletedEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointDeletedEvent)

	if err != nil {
		return err
	}

	*o = EndpointDeletedEvent(varEndpointDeletedEvent)

	return err
}

type NullableEndpointDeletedEvent struct {
	value *EndpointDeletedEvent
	isSet bool
}

func (v NullableEndpointDeletedEvent) Get() *EndpointDeletedEvent {
	return v.value
}

func (v *NullableEndpointDeletedEvent) Set(val *EndpointDeletedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointDeletedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointDeletedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointDeletedEvent(val *EndpointDeletedEvent) *NullableEndpointDeletedEvent {
	return &NullableEndpointDeletedEvent{value: val, isSet: true}
}

func (v NullableEndpointDeletedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointDeletedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
