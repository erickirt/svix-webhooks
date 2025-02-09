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

// checks if the EventTypeFromOpenApi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventTypeFromOpenApi{}

// EventTypeFromOpenApi struct for EventTypeFromOpenApi
type EventTypeFromOpenApi struct {
	Deprecated  bool    `json:"deprecated"`
	Description string  `json:"description"`
	FeatureFlag *string `json:"featureFlag,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The event type group's name
	GroupName *string `json:"groupName,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The event type's name
	Name    string                 `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	Schemas map[string]interface{} `json:"schemas,omitempty"`
}

type _EventTypeFromOpenApi EventTypeFromOpenApi

// NewEventTypeFromOpenApi instantiates a new EventTypeFromOpenApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeFromOpenApi(deprecated bool, description string, name string) *EventTypeFromOpenApi {
	this := EventTypeFromOpenApi{}
	this.Deprecated = deprecated
	this.Description = description
	this.Name = name
	return &this
}

// NewEventTypeFromOpenApiWithDefaults instantiates a new EventTypeFromOpenApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeFromOpenApiWithDefaults() *EventTypeFromOpenApi {
	this := EventTypeFromOpenApi{}
	return &this
}

// GetDeprecated returns the Deprecated field value
func (o *EventTypeFromOpenApi) GetDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value
// and a boolean to check if the value has been set.
func (o *EventTypeFromOpenApi) GetDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deprecated, true
}

// SetDeprecated sets field value
func (o *EventTypeFromOpenApi) SetDeprecated(v bool) {
	o.Deprecated = v
}

// GetDescription returns the Description field value
func (o *EventTypeFromOpenApi) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *EventTypeFromOpenApi) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *EventTypeFromOpenApi) SetDescription(v string) {
	o.Description = v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise.
func (o *EventTypeFromOpenApi) GetFeatureFlag() string {
	if o == nil || IsNil(o.FeatureFlag) {
		var ret string
		return ret
	}
	return *o.FeatureFlag
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeFromOpenApi) GetFeatureFlagOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureFlag) {
		return nil, false
	}
	return o.FeatureFlag, true
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *EventTypeFromOpenApi) HasFeatureFlag() bool {
	if o != nil && !IsNil(o.FeatureFlag) {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given string and assigns it to the FeatureFlag field.
func (o *EventTypeFromOpenApi) SetFeatureFlag(v string) {
	o.FeatureFlag = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *EventTypeFromOpenApi) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeFromOpenApi) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *EventTypeFromOpenApi) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *EventTypeFromOpenApi) SetGroupName(v string) {
	o.GroupName = &v
}

// GetName returns the Name field value
func (o *EventTypeFromOpenApi) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventTypeFromOpenApi) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventTypeFromOpenApi) SetName(v string) {
	o.Name = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *EventTypeFromOpenApi) GetSchemas() map[string]interface{} {
	if o == nil || IsNil(o.Schemas) {
		var ret map[string]interface{}
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeFromOpenApi) GetSchemasOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Schemas) {
		return map[string]interface{}{}, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *EventTypeFromOpenApi) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given map[string]interface{} and assigns it to the Schemas field.
func (o *EventTypeFromOpenApi) SetSchemas(v map[string]interface{}) {
	o.Schemas = v
}

func (o EventTypeFromOpenApi) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventTypeFromOpenApi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deprecated"] = o.Deprecated
	toSerialize["description"] = o.Description
	if !IsNil(o.FeatureFlag) {
		toSerialize["featureFlag"] = o.FeatureFlag
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	return toSerialize, nil
}

func (o *EventTypeFromOpenApi) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deprecated",
		"description",
		"name",
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

	varEventTypeFromOpenApi := _EventTypeFromOpenApi{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventTypeFromOpenApi)

	if err != nil {
		return err
	}

	*o = EventTypeFromOpenApi(varEventTypeFromOpenApi)

	return err
}

type NullableEventTypeFromOpenApi struct {
	value *EventTypeFromOpenApi
	isSet bool
}

func (v NullableEventTypeFromOpenApi) Get() *EventTypeFromOpenApi {
	return v.value
}

func (v *NullableEventTypeFromOpenApi) Set(val *EventTypeFromOpenApi) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeFromOpenApi) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeFromOpenApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeFromOpenApi(val *EventTypeFromOpenApi) *NullableEventTypeFromOpenApi {
	return &NullableEventTypeFromOpenApi{value: val, isSet: true}
}

func (v NullableEventTypeFromOpenApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeFromOpenApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
