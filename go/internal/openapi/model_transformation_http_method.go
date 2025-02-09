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

// TransformationHttpMethod the model 'TransformationHttpMethod'
type TransformationHttpMethod string

// List of TransformationHttpMethod
const (
	TRANSFORMATIONHTTPMETHOD_POST  TransformationHttpMethod = "POST"
	TRANSFORMATIONHTTPMETHOD_PUT   TransformationHttpMethod = "PUT"
	TRANSFORMATIONHTTPMETHOD_PATCH TransformationHttpMethod = "PATCH"
)

// All allowed values of TransformationHttpMethod enum
var AllowedTransformationHttpMethodEnumValues = []TransformationHttpMethod{
	"POST",
	"PUT",
	"PATCH",
}

func (v *TransformationHttpMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransformationHttpMethod(value)
	for _, existing := range AllowedTransformationHttpMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransformationHttpMethod", value)
}

// NewTransformationHttpMethodFromValue returns a pointer to a valid TransformationHttpMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransformationHttpMethodFromValue(v string) (*TransformationHttpMethod, error) {
	ev := TransformationHttpMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransformationHttpMethod: valid values are %v", v, AllowedTransformationHttpMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransformationHttpMethod) IsValid() bool {
	for _, existing := range AllowedTransformationHttpMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransformationHttpMethod value
func (v TransformationHttpMethod) Ptr() *TransformationHttpMethod {
	return &v
}

type NullableTransformationHttpMethod struct {
	value *TransformationHttpMethod
	isSet bool
}

func (v NullableTransformationHttpMethod) Get() *TransformationHttpMethod {
	return v.value
}

func (v *NullableTransformationHttpMethod) Set(val *TransformationHttpMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationHttpMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationHttpMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationHttpMethod(val *TransformationHttpMethod) *NullableTransformationHttpMethod {
	return &NullableTransformationHttpMethod{value: val, isSet: true}
}

func (v NullableTransformationHttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationHttpMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
