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

// checks if the EventTypeImportOpenApiIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventTypeImportOpenApiIn{}

// EventTypeImportOpenApiIn Import a list of event types from webhooks defined in an OpenAPI spec.  The OpenAPI spec can be specified as either `spec` given the spec as a JSON object, or as `specRaw` (a `string`) which will be parsed as YAML or JSON by the server. Sending neither or both is invalid, resulting in a `400` **Bad Request**.
type EventTypeImportOpenApiIn struct {
	// If `true`, return the event types that would be modified without actually modifying them.
	DryRun *bool `json:"dryRun,omitempty"`
	// If `true`, all existing event types that are not in the spec will be archived.
	ReplaceAll *bool `json:"replaceAll,omitempty"`
	// A pre-parsed JSON spec.
	Spec map[string]interface{} `json:"spec,omitempty"`
	// A string, parsed by the server as YAML or JSON.
	SpecRaw *string `json:"specRaw,omitempty"`
}

// NewEventTypeImportOpenApiIn instantiates a new EventTypeImportOpenApiIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventTypeImportOpenApiIn() *EventTypeImportOpenApiIn {
	this := EventTypeImportOpenApiIn{}
	var dryRun bool = false
	this.DryRun = &dryRun
	var replaceAll bool = false
	this.ReplaceAll = &replaceAll
	return &this
}

// NewEventTypeImportOpenApiInWithDefaults instantiates a new EventTypeImportOpenApiIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventTypeImportOpenApiInWithDefaults() *EventTypeImportOpenApiIn {
	this := EventTypeImportOpenApiIn{}
	var dryRun bool = false
	this.DryRun = &dryRun
	var replaceAll bool = false
	this.ReplaceAll = &replaceAll
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *EventTypeImportOpenApiIn) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeImportOpenApiIn) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *EventTypeImportOpenApiIn) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *EventTypeImportOpenApiIn) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetReplaceAll returns the ReplaceAll field value if set, zero value otherwise.
func (o *EventTypeImportOpenApiIn) GetReplaceAll() bool {
	if o == nil || IsNil(o.ReplaceAll) {
		var ret bool
		return ret
	}
	return *o.ReplaceAll
}

// GetReplaceAllOk returns a tuple with the ReplaceAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeImportOpenApiIn) GetReplaceAllOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceAll) {
		return nil, false
	}
	return o.ReplaceAll, true
}

// HasReplaceAll returns a boolean if a field has been set.
func (o *EventTypeImportOpenApiIn) HasReplaceAll() bool {
	if o != nil && !IsNil(o.ReplaceAll) {
		return true
	}

	return false
}

// SetReplaceAll gets a reference to the given bool and assigns it to the ReplaceAll field.
func (o *EventTypeImportOpenApiIn) SetReplaceAll(v bool) {
	o.ReplaceAll = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *EventTypeImportOpenApiIn) GetSpec() map[string]interface{} {
	if o == nil || IsNil(o.Spec) {
		var ret map[string]interface{}
		return ret
	}
	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeImportOpenApiIn) GetSpecOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Spec) {
		return map[string]interface{}{}, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *EventTypeImportOpenApiIn) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given map[string]interface{} and assigns it to the Spec field.
func (o *EventTypeImportOpenApiIn) SetSpec(v map[string]interface{}) {
	o.Spec = v
}

// GetSpecRaw returns the SpecRaw field value if set, zero value otherwise.
func (o *EventTypeImportOpenApiIn) GetSpecRaw() string {
	if o == nil || IsNil(o.SpecRaw) {
		var ret string
		return ret
	}
	return *o.SpecRaw
}

// GetSpecRawOk returns a tuple with the SpecRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventTypeImportOpenApiIn) GetSpecRawOk() (*string, bool) {
	if o == nil || IsNil(o.SpecRaw) {
		return nil, false
	}
	return o.SpecRaw, true
}

// HasSpecRaw returns a boolean if a field has been set.
func (o *EventTypeImportOpenApiIn) HasSpecRaw() bool {
	if o != nil && !IsNil(o.SpecRaw) {
		return true
	}

	return false
}

// SetSpecRaw gets a reference to the given string and assigns it to the SpecRaw field.
func (o *EventTypeImportOpenApiIn) SetSpecRaw(v string) {
	o.SpecRaw = &v
}

func (o EventTypeImportOpenApiIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventTypeImportOpenApiIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !IsNil(o.ReplaceAll) {
		toSerialize["replaceAll"] = o.ReplaceAll
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.SpecRaw) {
		toSerialize["specRaw"] = o.SpecRaw
	}
	return toSerialize, nil
}

type NullableEventTypeImportOpenApiIn struct {
	value *EventTypeImportOpenApiIn
	isSet bool
}

func (v NullableEventTypeImportOpenApiIn) Get() *EventTypeImportOpenApiIn {
	return v.value
}

func (v *NullableEventTypeImportOpenApiIn) Set(val *EventTypeImportOpenApiIn) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeImportOpenApiIn) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeImportOpenApiIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeImportOpenApiIn(val *EventTypeImportOpenApiIn) *NullableEventTypeImportOpenApiIn {
	return &NullableEventTypeImportOpenApiIn{value: val, isSet: true}
}

func (v NullableEventTypeImportOpenApiIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeImportOpenApiIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
