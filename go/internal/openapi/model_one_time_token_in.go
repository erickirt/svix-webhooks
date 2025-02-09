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

// checks if the OneTimeTokenIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneTimeTokenIn{}

// OneTimeTokenIn struct for OneTimeTokenIn
type OneTimeTokenIn struct {
	OneTimeToken string `json:"oneTimeToken"`
}

type _OneTimeTokenIn OneTimeTokenIn

// NewOneTimeTokenIn instantiates a new OneTimeTokenIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneTimeTokenIn(oneTimeToken string) *OneTimeTokenIn {
	this := OneTimeTokenIn{}
	this.OneTimeToken = oneTimeToken
	return &this
}

// NewOneTimeTokenInWithDefaults instantiates a new OneTimeTokenIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneTimeTokenInWithDefaults() *OneTimeTokenIn {
	this := OneTimeTokenIn{}
	return &this
}

// GetOneTimeToken returns the OneTimeToken field value
func (o *OneTimeTokenIn) GetOneTimeToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OneTimeToken
}

// GetOneTimeTokenOk returns a tuple with the OneTimeToken field value
// and a boolean to check if the value has been set.
func (o *OneTimeTokenIn) GetOneTimeTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OneTimeToken, true
}

// SetOneTimeToken sets field value
func (o *OneTimeTokenIn) SetOneTimeToken(v string) {
	o.OneTimeToken = v
}

func (o OneTimeTokenIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OneTimeTokenIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oneTimeToken"] = o.OneTimeToken
	return toSerialize, nil
}

func (o *OneTimeTokenIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"oneTimeToken",
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

	varOneTimeTokenIn := _OneTimeTokenIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOneTimeTokenIn)

	if err != nil {
		return err
	}

	*o = OneTimeTokenIn(varOneTimeTokenIn)

	return err
}

type NullableOneTimeTokenIn struct {
	value *OneTimeTokenIn
	isSet bool
}

func (v NullableOneTimeTokenIn) Get() *OneTimeTokenIn {
	return v.value
}

func (v *NullableOneTimeTokenIn) Set(val *OneTimeTokenIn) {
	v.value = val
	v.isSet = true
}

func (v NullableOneTimeTokenIn) IsSet() bool {
	return v.isSet
}

func (v *NullableOneTimeTokenIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneTimeTokenIn(val *OneTimeTokenIn) *NullableOneTimeTokenIn {
	return &NullableOneTimeTokenIn{value: val, isSet: true}
}

func (v NullableOneTimeTokenIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOneTimeTokenIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
