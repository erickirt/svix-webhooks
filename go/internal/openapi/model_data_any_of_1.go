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

// checks if the DataAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAnyOf1{}

// DataAnyOf1 struct for DataAnyOf1
type DataAnyOf1 struct {
	AppStats []ApplicationStatsOut `json:"app_stats"`
}

type _DataAnyOf1 DataAnyOf1

// NewDataAnyOf1 instantiates a new DataAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAnyOf1(appStats []ApplicationStatsOut) *DataAnyOf1 {
	this := DataAnyOf1{}
	this.AppStats = appStats
	return &this
}

// NewDataAnyOf1WithDefaults instantiates a new DataAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAnyOf1WithDefaults() *DataAnyOf1 {
	this := DataAnyOf1{}
	return &this
}

// GetAppStats returns the AppStats field value
func (o *DataAnyOf1) GetAppStats() []ApplicationStatsOut {
	if o == nil {
		var ret []ApplicationStatsOut
		return ret
	}

	return o.AppStats
}

// GetAppStatsOk returns a tuple with the AppStats field value
// and a boolean to check if the value has been set.
func (o *DataAnyOf1) GetAppStatsOk() ([]ApplicationStatsOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppStats, true
}

// SetAppStats sets field value
func (o *DataAnyOf1) SetAppStats(v []ApplicationStatsOut) {
	o.AppStats = v
}

func (o DataAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_stats"] = o.AppStats
	return toSerialize, nil
}

func (o *DataAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app_stats",
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

	varDataAnyOf1 := _DataAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDataAnyOf1)

	if err != nil {
		return err
	}

	*o = DataAnyOf1(varDataAnyOf1)

	return err
}

type NullableDataAnyOf1 struct {
	value *DataAnyOf1
	isSet bool
}

func (v NullableDataAnyOf1) Get() *DataAnyOf1 {
	return v.value
}

func (v *NullableDataAnyOf1) Set(val *DataAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAnyOf1(val *DataAnyOf1) *NullableDataAnyOf1 {
	return &NullableDataAnyOf1{value: val, isSet: true}
}

func (v NullableDataAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
