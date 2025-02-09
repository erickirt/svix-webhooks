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

// checks if the GenerateOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateOut{}

// GenerateOut struct for GenerateOut
type GenerateOut struct {
	Choices []CompletionChoice `json:"choices"`
	Created int64              `json:"created"`
	Id      string             `json:"id"`
	Model   string             `json:"model"`
	Object  string             `json:"object"`
}

type _GenerateOut GenerateOut

// NewGenerateOut instantiates a new GenerateOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateOut(choices []CompletionChoice, created int64, id string, model string, object string) *GenerateOut {
	this := GenerateOut{}
	this.Choices = choices
	this.Created = created
	this.Id = id
	this.Model = model
	this.Object = object
	return &this
}

// NewGenerateOutWithDefaults instantiates a new GenerateOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateOutWithDefaults() *GenerateOut {
	this := GenerateOut{}
	return &this
}

// GetChoices returns the Choices field value
func (o *GenerateOut) GetChoices() []CompletionChoice {
	if o == nil {
		var ret []CompletionChoice
		return ret
	}

	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value
// and a boolean to check if the value has been set.
func (o *GenerateOut) GetChoicesOk() ([]CompletionChoice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Choices, true
}

// SetChoices sets field value
func (o *GenerateOut) SetChoices(v []CompletionChoice) {
	o.Choices = v
}

// GetCreated returns the Created field value
func (o *GenerateOut) GetCreated() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GenerateOut) GetCreatedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GenerateOut) SetCreated(v int64) {
	o.Created = v
}

// GetId returns the Id field value
func (o *GenerateOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenerateOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenerateOut) SetId(v string) {
	o.Id = v
}

// GetModel returns the Model field value
func (o *GenerateOut) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *GenerateOut) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *GenerateOut) SetModel(v string) {
	o.Model = v
}

// GetObject returns the Object field value
func (o *GenerateOut) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *GenerateOut) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *GenerateOut) SetObject(v string) {
	o.Object = v
}

func (o GenerateOut) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["choices"] = o.Choices
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["model"] = o.Model
	toSerialize["object"] = o.Object
	return toSerialize, nil
}

func (o *GenerateOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"choices",
		"created",
		"id",
		"model",
		"object",
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

	varGenerateOut := _GenerateOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGenerateOut)

	if err != nil {
		return err
	}

	*o = GenerateOut(varGenerateOut)

	return err
}

type NullableGenerateOut struct {
	value *GenerateOut
	isSet bool
}

func (v NullableGenerateOut) Get() *GenerateOut {
	return v.value
}

func (v *NullableGenerateOut) Set(val *GenerateOut) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateOut) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateOut(val *GenerateOut) *NullableGenerateOut {
	return &NullableGenerateOut{value: val, isSet: true}
}

func (v NullableGenerateOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
