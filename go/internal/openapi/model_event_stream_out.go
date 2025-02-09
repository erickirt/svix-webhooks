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

// checks if the EventStreamOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventStreamOut{}

// EventStreamOut struct for EventStreamOut
type EventStreamOut struct {
	Data     []EventOut `json:"data"`
	Done     bool       `json:"done"`
	Iterator string     `json:"iterator"`
}

type _EventStreamOut EventStreamOut

// NewEventStreamOut instantiates a new EventStreamOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventStreamOut(data []EventOut, done bool, iterator string) *EventStreamOut {
	this := EventStreamOut{}
	this.Data = data
	this.Done = done
	this.Iterator = iterator
	return &this
}

// NewEventStreamOutWithDefaults instantiates a new EventStreamOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventStreamOutWithDefaults() *EventStreamOut {
	this := EventStreamOut{}
	return &this
}

// GetData returns the Data field value
func (o *EventStreamOut) GetData() []EventOut {
	if o == nil {
		var ret []EventOut
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EventStreamOut) GetDataOk() ([]EventOut, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EventStreamOut) SetData(v []EventOut) {
	o.Data = v
}

// GetDone returns the Done field value
func (o *EventStreamOut) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *EventStreamOut) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *EventStreamOut) SetDone(v bool) {
	o.Done = v
}

// GetIterator returns the Iterator field value
func (o *EventStreamOut) GetIterator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iterator
}

// GetIteratorOk returns a tuple with the Iterator field value
// and a boolean to check if the value has been set.
func (o *EventStreamOut) GetIteratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iterator, true
}

// SetIterator sets field value
func (o *EventStreamOut) SetIterator(v string) {
	o.Iterator = v
}

func (o EventStreamOut) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventStreamOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["done"] = o.Done
	toSerialize["iterator"] = o.Iterator
	return toSerialize, nil
}

func (o *EventStreamOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"done",
		"iterator",
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

	varEventStreamOut := _EventStreamOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventStreamOut)

	if err != nil {
		return err
	}

	*o = EventStreamOut(varEventStreamOut)

	return err
}

type NullableEventStreamOut struct {
	value *EventStreamOut
	isSet bool
}

func (v NullableEventStreamOut) Get() *EventStreamOut {
	return v.value
}

func (v *NullableEventStreamOut) Set(val *EventStreamOut) {
	v.value = val
	v.isSet = true
}

func (v NullableEventStreamOut) IsSet() bool {
	return v.isSet
}

func (v *NullableEventStreamOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventStreamOut(val *EventStreamOut) *NullableEventStreamOut {
	return &NullableEventStreamOut{value: val, isSet: true}
}

func (v NullableEventStreamOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventStreamOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
