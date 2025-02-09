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

// checks if the OperationalWebhookEndpointUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationalWebhookEndpointUpdate{}

// OperationalWebhookEndpointUpdate struct for OperationalWebhookEndpointUpdate
type OperationalWebhookEndpointUpdate struct {
	Description *string            `json:"description,omitempty"`
	Disabled    *bool              `json:"disabled,omitempty"`
	FilterTypes []string           `json:"filterTypes,omitempty"`
	Metadata    *map[string]string `json:"metadata,omitempty"`
	RateLimit   *int32             `json:"rateLimit,omitempty"`
	// Optional unique identifier for the endpoint.
	Uid *string `json:"uid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	Url string  `json:"url"`
}

type _OperationalWebhookEndpointUpdate OperationalWebhookEndpointUpdate

// NewOperationalWebhookEndpointUpdate instantiates a new OperationalWebhookEndpointUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationalWebhookEndpointUpdate(url string) *OperationalWebhookEndpointUpdate {
	this := OperationalWebhookEndpointUpdate{}
	var description string = ""
	this.Description = &description
	var disabled bool = false
	this.Disabled = &disabled
	this.Url = url
	return &this
}

// NewOperationalWebhookEndpointUpdateWithDefaults instantiates a new OperationalWebhookEndpointUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationalWebhookEndpointUpdateWithDefaults() *OperationalWebhookEndpointUpdate {
	this := OperationalWebhookEndpointUpdate{}
	var description string = ""
	this.Description = &description
	var disabled bool = false
	this.Disabled = &disabled
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OperationalWebhookEndpointUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointUpdate) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointUpdate) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *OperationalWebhookEndpointUpdate) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointUpdate) GetFilterTypes() []string {
	if o == nil || IsNil(o.FilterTypes) {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointUpdate) HasFilterTypes() bool {
	if o != nil && !IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *OperationalWebhookEndpointUpdate) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointUpdate) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointUpdate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OperationalWebhookEndpointUpdate) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointUpdate) GetRateLimit() int32 {
	if o == nil || IsNil(o.RateLimit) {
		var ret int32
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetRateLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.RateLimit) {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointUpdate) HasRateLimit() bool {
	if o != nil && !IsNil(o.RateLimit) {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given int32 and assigns it to the RateLimit field.
func (o *OperationalWebhookEndpointUpdate) SetRateLimit(v int32) {
	o.RateLimit = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *OperationalWebhookEndpointUpdate) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *OperationalWebhookEndpointUpdate) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *OperationalWebhookEndpointUpdate) SetUid(v string) {
	o.Uid = &v
}

// GetUrl returns the Url field value
func (o *OperationalWebhookEndpointUpdate) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *OperationalWebhookEndpointUpdate) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *OperationalWebhookEndpointUpdate) SetUrl(v string) {
	o.Url = v
}

func (o OperationalWebhookEndpointUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationalWebhookEndpointUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.FilterTypes) {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RateLimit) {
		toSerialize["rateLimit"] = o.RateLimit
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *OperationalWebhookEndpointUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varOperationalWebhookEndpointUpdate := _OperationalWebhookEndpointUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOperationalWebhookEndpointUpdate)

	if err != nil {
		return err
	}

	*o = OperationalWebhookEndpointUpdate(varOperationalWebhookEndpointUpdate)

	return err
}

type NullableOperationalWebhookEndpointUpdate struct {
	value *OperationalWebhookEndpointUpdate
	isSet bool
}

func (v NullableOperationalWebhookEndpointUpdate) Get() *OperationalWebhookEndpointUpdate {
	return v.value
}

func (v *NullableOperationalWebhookEndpointUpdate) Set(val *OperationalWebhookEndpointUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationalWebhookEndpointUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationalWebhookEndpointUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationalWebhookEndpointUpdate(val *OperationalWebhookEndpointUpdate) *NullableOperationalWebhookEndpointUpdate {
	return &NullableOperationalWebhookEndpointUpdate{value: val, isSet: true}
}

func (v NullableOperationalWebhookEndpointUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationalWebhookEndpointUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
