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

// checks if the StreamPortalAccessIn type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamPortalAccessIn{}

// StreamPortalAccessIn struct for StreamPortalAccessIn
type StreamPortalAccessIn struct {
	// How long the token will be valid for, in seconds.  Valid values are between 1 hour and 7 days. The default is 7 days.
	Expiry *int32 `json:"expiry,omitempty"`
	// The set of feature flags the created token will have access to.
	FeatureFlags []string `json:"featureFlags,omitempty"`
	// Whether the app portal should be in read-only mode.
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// NewStreamPortalAccessIn instantiates a new StreamPortalAccessIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamPortalAccessIn() *StreamPortalAccessIn {
	this := StreamPortalAccessIn{}
	var expiry int32 = 604800
	this.Expiry = &expiry
	return &this
}

// NewStreamPortalAccessInWithDefaults instantiates a new StreamPortalAccessIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamPortalAccessInWithDefaults() *StreamPortalAccessIn {
	this := StreamPortalAccessIn{}
	var expiry int32 = 604800
	this.Expiry = &expiry
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *StreamPortalAccessIn) GetExpiry() int32 {
	if o == nil || IsNil(o.Expiry) {
		var ret int32
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamPortalAccessIn) GetExpiryOk() (*int32, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *StreamPortalAccessIn) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given int32 and assigns it to the Expiry field.
func (o *StreamPortalAccessIn) SetExpiry(v int32) {
	o.Expiry = &v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *StreamPortalAccessIn) GetFeatureFlags() []string {
	if o == nil || IsNil(o.FeatureFlags) {
		var ret []string
		return ret
	}
	return o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamPortalAccessIn) GetFeatureFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureFlags) {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *StreamPortalAccessIn) HasFeatureFlags() bool {
	if o != nil && !IsNil(o.FeatureFlags) {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given []string and assigns it to the FeatureFlags field.
func (o *StreamPortalAccessIn) SetFeatureFlags(v []string) {
	o.FeatureFlags = v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *StreamPortalAccessIn) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamPortalAccessIn) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *StreamPortalAccessIn) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *StreamPortalAccessIn) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o StreamPortalAccessIn) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamPortalAccessIn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.FeatureFlags) {
		toSerialize["featureFlags"] = o.FeatureFlags
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return toSerialize, nil
}

type NullableStreamPortalAccessIn struct {
	value *StreamPortalAccessIn
	isSet bool
}

func (v NullableStreamPortalAccessIn) Get() *StreamPortalAccessIn {
	return v.value
}

func (v *NullableStreamPortalAccessIn) Set(val *StreamPortalAccessIn) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamPortalAccessIn) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamPortalAccessIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamPortalAccessIn(val *StreamPortalAccessIn) *NullableStreamPortalAccessIn {
	return &NullableStreamPortalAccessIn{value: val, isSet: true}
}

func (v NullableStreamPortalAccessIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamPortalAccessIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
