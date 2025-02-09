/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OperationalWebhookEndpointSecretIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationalWebhookEndpointSecretIn{}

// OperationalWebhookEndpointSecretIn struct for OperationalWebhookEndpointSecretIn
type OperationalWebhookEndpointSecretIn struct {
	// The endpoint's verification secret.  Format: `base64` encoded random bytes optionally prefixed with `whsec_`. It is recommended to not set this and let the server generate the secret.
	Key *string `json:"key,omitempty" validate:"regexp=^(whsec_)?[a-zA-Z0-9+\\/=]{32,100}$"`
}

// NewOperationalWebhookEndpointSecretIn instantiates a new OperationalWebhookEndpointSecretIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationalWebhookEndpointSecretIn() *OperationalWebhookEndpointSecretIn {
	this := OperationalWebhookEndpointSecretIn{}
	return &this
}

// NewOperationalWebhookEndpointSecretInWithDefaults instantiates a new OperationalWebhookEndpointSecretIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationalWebhookEndpointSecretInWithDefaults() *OperationalWebhookEndpointSecretIn {
	this := OperationalWebhookEndpointSecretIn{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointSecretIn) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointSecretIn) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointSecretIn) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *OperationalWebhookEndpointSecretIn) SetKey(v string) {
	o.Key = &v
}

func (o OperationalWebhookEndpointSecretIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationalWebhookEndpointSecretIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableOperationalWebhookEndpointSecretIn struct {
	value *OperationalWebhookEndpointSecretIn
	isSet bool
}

func (v NullableOperationalWebhookEndpointSecretIn) Get() *OperationalWebhookEndpointSecretIn {
	return v.value
}

func (v *NullableOperationalWebhookEndpointSecretIn) Set(val *OperationalWebhookEndpointSecretIn) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationalWebhookEndpointSecretIn) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationalWebhookEndpointSecretIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationalWebhookEndpointSecretIn(val *OperationalWebhookEndpointSecretIn) *NullableOperationalWebhookEndpointSecretIn {
	return &NullableOperationalWebhookEndpointSecretIn{value: val, isSet: true}
}

func (v NullableOperationalWebhookEndpointSecretIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationalWebhookEndpointSecretIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
