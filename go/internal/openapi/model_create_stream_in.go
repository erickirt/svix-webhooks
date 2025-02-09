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

// checks if the CreateStreamIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStreamIn{}

// CreateStreamIn struct for CreateStreamIn
type CreateStreamIn struct {
	Events []EventIn `json:"events"`
	Stream *StreamIn `json:"stream,omitempty"`
}

type _CreateStreamIn CreateStreamIn

// NewCreateStreamIn instantiates a new CreateStreamIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStreamIn(events []EventIn) *CreateStreamIn {
	this := CreateStreamIn{}
	this.Events = events
	return &this
}

// NewCreateStreamInWithDefaults instantiates a new CreateStreamIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStreamInWithDefaults() *CreateStreamIn {
	this := CreateStreamIn{}
	return &this
}

// GetEvents returns the Events field value
func (o *CreateStreamIn) GetEvents() []EventIn {
	if o == nil {
		var ret []EventIn
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *CreateStreamIn) GetEventsOk() ([]EventIn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *CreateStreamIn) SetEvents(v []EventIn) {
	o.Events = v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *CreateStreamIn) GetStream() StreamIn {
	if o == nil || IsNil(o.Stream) {
		var ret StreamIn
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStreamIn) GetStreamOk() (*StreamIn, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *CreateStreamIn) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given StreamIn and assigns it to the Stream field.
func (o *CreateStreamIn) SetStream(v StreamIn) {
	o.Stream = &v
}

func (o CreateStreamIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateStreamIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	return toSerialize, nil
}

func (o *CreateStreamIn) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varCreateStreamIn := _CreateStreamIn{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateStreamIn)

	if err != nil {
		return err
	}

	*o = CreateStreamIn(varCreateStreamIn)

	return err
}

type NullableCreateStreamIn struct {
	value *CreateStreamIn
	isSet bool
}

func (v NullableCreateStreamIn) Get() *CreateStreamIn {
	return v.value
}

func (v *NullableCreateStreamIn) Set(val *CreateStreamIn) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStreamIn) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStreamIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStreamIn(val *CreateStreamIn) *NullableCreateStreamIn {
	return &NullableCreateStreamIn{value: val, isSet: true}
}

func (v NullableCreateStreamIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateStreamIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
