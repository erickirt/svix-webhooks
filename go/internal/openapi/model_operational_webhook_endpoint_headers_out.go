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

// checks if the OperationalWebhookEndpointHeadersOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationalWebhookEndpointHeadersOut{}

// OperationalWebhookEndpointHeadersOut struct for OperationalWebhookEndpointHeadersOut
type OperationalWebhookEndpointHeadersOut struct {
	Headers   map[string]string `json:"headers"`
	Sensitive []string          `json:"sensitive"`
}

type _OperationalWebhookEndpointHeadersOut OperationalWebhookEndpointHeadersOut

// NewOperationalWebhookEndpointHeadersOut instantiates a new OperationalWebhookEndpointHeadersOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationalWebhookEndpointHeadersOut(headers map[string]string, sensitive []string) *OperationalWebhookEndpointHeadersOut {
	this := OperationalWebhookEndpointHeadersOut{}
	this.Headers = headers
	this.Sensitive = sensitive
	return &this
}

// NewOperationalWebhookEndpointHeadersOutWithDefaults instantiates a new OperationalWebhookEndpointHeadersOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationalWebhookEndpointHeadersOutWithDefaults() *OperationalWebhookEndpointHeadersOut {
	this := OperationalWebhookEndpointHeadersOut{}
	return &this
}

// GetHeaders returns the Headers field value
func (o *OperationalWebhookEndpointHeadersOut) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointHeadersOut) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *OperationalWebhookEndpointHeadersOut) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetSensitive returns the Sensitive field value
func (o *OperationalWebhookEndpointHeadersOut) GetSensitive() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointHeadersOut) GetSensitiveOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// SetSensitive sets field value
func (o *OperationalWebhookEndpointHeadersOut) SetSensitive(v []string) {
	o.Sensitive = v
}

func (o OperationalWebhookEndpointHeadersOut) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationalWebhookEndpointHeadersOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["headers"] = o.Headers
	toSerialize["sensitive"] = o.Sensitive
	return toSerialize, nil
}

func (o *OperationalWebhookEndpointHeadersOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"headers",
		"sensitive",
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

	varOperationalWebhookEndpointHeadersOut := _OperationalWebhookEndpointHeadersOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperationalWebhookEndpointHeadersOut)

	if err != nil {
		return err
	}

	*o = OperationalWebhookEndpointHeadersOut(varOperationalWebhookEndpointHeadersOut)

	return err
}

type NullableOperationalWebhookEndpointHeadersOut struct {
	value *OperationalWebhookEndpointHeadersOut
	isSet bool
}

func (v NullableOperationalWebhookEndpointHeadersOut) Get() *OperationalWebhookEndpointHeadersOut {
	return v.value
}

func (v *NullableOperationalWebhookEndpointHeadersOut) Set(val *OperationalWebhookEndpointHeadersOut) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationalWebhookEndpointHeadersOut) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationalWebhookEndpointHeadersOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationalWebhookEndpointHeadersOut(val *OperationalWebhookEndpointHeadersOut) *NullableOperationalWebhookEndpointHeadersOut {
	return &NullableOperationalWebhookEndpointHeadersOut{value: val, isSet: true}
}

func (v NullableOperationalWebhookEndpointHeadersOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationalWebhookEndpointHeadersOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
