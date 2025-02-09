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
	"time"
)

// checks if the RecoverIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoverIn{}

// RecoverIn struct for RecoverIn
type RecoverIn struct {
	Since time.Time  `json:"since"`
	Until *time.Time `json:"until,omitempty"`
}

type _RecoverIn RecoverIn

// NewRecoverIn instantiates a new RecoverIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoverIn(since time.Time) *RecoverIn {
	this := RecoverIn{}
	this.Since = since
	return &this
}

// NewRecoverInWithDefaults instantiates a new RecoverIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoverInWithDefaults() *RecoverIn {
	this := RecoverIn{}
	return &this
}

// GetSince returns the Since field value
func (o *RecoverIn) GetSince() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Since
}

// GetSinceOk returns a tuple with the Since field value
// and a boolean to check if the value has been set.
func (o *RecoverIn) GetSinceOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Since, true
}

// SetSince sets field value
func (o *RecoverIn) SetSince(v time.Time) {
	o.Since = v
}

// GetUntil returns the Until field value if set, zero value otherwise.
func (o *RecoverIn) GetUntil() time.Time {
	if o == nil || IsNil(o.Until) {
		var ret time.Time
		return ret
	}
	return *o.Until
}

// GetUntilOk returns a tuple with the Until field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoverIn) GetUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Until) {
		return nil, false
	}
	return o.Until, true
}

// HasUntil returns a boolean if a field has been set.
func (o *RecoverIn) HasUntil() bool {
	if o != nil && !IsNil(o.Until) {
		return true
	}

	return false
}

// SetUntil gets a reference to the given time.Time and assigns it to the Until field.
func (o *RecoverIn) SetUntil(v time.Time) {
	o.Until = &v
}

func (o RecoverIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoverIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["since"] = o.Since
	if !IsNil(o.Until) {
		toSerialize["until"] = o.Until
	}
	return toSerialize, nil
}

func (o *RecoverIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"since",
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

	varRecoverIn := _RecoverIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecoverIn)

	if err != nil {
		return err
	}

	*o = RecoverIn(varRecoverIn)

	return err
}

type NullableRecoverIn struct {
	value *RecoverIn
	isSet bool
}

func (v NullableRecoverIn) Get() *RecoverIn {
	return v.value
}

func (v *NullableRecoverIn) Set(val *RecoverIn) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoverIn) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoverIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoverIn(val *RecoverIn) *NullableRecoverIn {
	return &NullableRecoverIn{value: val, isSet: true}
}

func (v NullableRecoverIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoverIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
