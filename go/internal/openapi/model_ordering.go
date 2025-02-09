/*
Svix API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Ordering Defines the ordering in a listing of results.
type Ordering string

// List of Ordering
const (
	ORDERING_ASCENDING  Ordering = "ascending"
	ORDERING_DESCENDING Ordering = "descending"
)

// All allowed values of Ordering enum
var AllowedOrderingEnumValues = []Ordering{
	"ascending",
	"descending",
}

func (v *Ordering) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Ordering(value)
	for _, existing := range AllowedOrderingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Ordering", value)
}

// NewOrderingFromValue returns a pointer to a valid Ordering
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderingFromValue(v string) (*Ordering, error) {
	ev := Ordering(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Ordering: valid values are %v", v, AllowedOrderingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Ordering) IsValid() bool {
	for _, existing := range AllowedOrderingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Ordering value
func (v Ordering) Ptr() *Ordering {
	return &v
}

type NullableOrdering struct {
	value *Ordering
	isSet bool
}

func (v NullableOrdering) Get() *Ordering {
	return v.value
}

func (v *NullableOrdering) Set(val *Ordering) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdering) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdering(val *Ordering) *NullableOrdering {
	return &NullableOrdering{value: val, isSet: true}
}

func (v NullableOrdering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
