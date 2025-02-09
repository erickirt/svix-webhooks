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

// checks if the EndpointEnabledEventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointEnabledEventData{}

// EndpointEnabledEventData Sent when an endpoint has been enabled.
type EndpointEnabledEventData struct {
	// The app's ID
	AppId string `json:"appId"`
	// The app's UID
	AppUid *string `json:"appUid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	// The ep's ID
	EndpointId string `json:"endpointId"`
	// The ep's UID
	EndpointUid *string `json:"endpointUid,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
}

type _EndpointEnabledEventData EndpointEnabledEventData

// NewEndpointEnabledEventData instantiates a new EndpointEnabledEventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointEnabledEventData(appId string, endpointId string) *EndpointEnabledEventData {
	this := EndpointEnabledEventData{}
	this.AppId = appId
	this.EndpointId = endpointId
	return &this
}

// NewEndpointEnabledEventDataWithDefaults instantiates a new EndpointEnabledEventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointEnabledEventDataWithDefaults() *EndpointEnabledEventData {
	this := EndpointEnabledEventData{}
	return &this
}

// GetAppId returns the AppId field value
func (o *EndpointEnabledEventData) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *EndpointEnabledEventData) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *EndpointEnabledEventData) SetAppId(v string) {
	o.AppId = v
}

// GetAppUid returns the AppUid field value if set, zero value otherwise.
func (o *EndpointEnabledEventData) GetAppUid() string {
	if o == nil || IsNil(o.AppUid) {
		var ret string
		return ret
	}
	return *o.AppUid
}

// GetAppUidOk returns a tuple with the AppUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointEnabledEventData) GetAppUidOk() (*string, bool) {
	if o == nil || IsNil(o.AppUid) {
		return nil, false
	}
	return o.AppUid, true
}

// HasAppUid returns a boolean if a field has been set.
func (o *EndpointEnabledEventData) HasAppUid() bool {
	if o != nil && !IsNil(o.AppUid) {
		return true
	}

	return false
}

// SetAppUid gets a reference to the given string and assigns it to the AppUid field.
func (o *EndpointEnabledEventData) SetAppUid(v string) {
	o.AppUid = &v
}

// GetEndpointId returns the EndpointId field value
func (o *EndpointEnabledEventData) GetEndpointId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *EndpointEnabledEventData) GetEndpointIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *EndpointEnabledEventData) SetEndpointId(v string) {
	o.EndpointId = v
}

// GetEndpointUid returns the EndpointUid field value if set, zero value otherwise.
func (o *EndpointEnabledEventData) GetEndpointUid() string {
	if o == nil || IsNil(o.EndpointUid) {
		var ret string
		return ret
	}
	return *o.EndpointUid
}

// GetEndpointUidOk returns a tuple with the EndpointUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointEnabledEventData) GetEndpointUidOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointUid) {
		return nil, false
	}
	return o.EndpointUid, true
}

// HasEndpointUid returns a boolean if a field has been set.
func (o *EndpointEnabledEventData) HasEndpointUid() bool {
	if o != nil && !IsNil(o.EndpointUid) {
		return true
	}

	return false
}

// SetEndpointUid gets a reference to the given string and assigns it to the EndpointUid field.
func (o *EndpointEnabledEventData) SetEndpointUid(v string) {
	o.EndpointUid = &v
}

func (o EndpointEnabledEventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointEnabledEventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !IsNil(o.AppUid) {
		toSerialize["appUid"] = o.AppUid
	}
	toSerialize["endpointId"] = o.EndpointId
	if !IsNil(o.EndpointUid) {
		toSerialize["endpointUid"] = o.EndpointUid
	}
	return toSerialize, nil
}

func (o *EndpointEnabledEventData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appId",
		"endpointId",
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

	varEndpointEnabledEventData := _EndpointEnabledEventData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointEnabledEventData)

	if err != nil {
		return err
	}

	*o = EndpointEnabledEventData(varEndpointEnabledEventData)

	return err
}

type NullableEndpointEnabledEventData struct {
	value *EndpointEnabledEventData
	isSet bool
}

func (v NullableEndpointEnabledEventData) Get() *EndpointEnabledEventData {
	return v.value
}

func (v *NullableEndpointEnabledEventData) Set(val *EndpointEnabledEventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointEnabledEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointEnabledEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointEnabledEventData(val *EndpointEnabledEventData) *NullableEndpointEnabledEventData {
	return &NullableEndpointEnabledEventData{value: val, isSet: true}
}

func (v NullableEndpointEnabledEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointEnabledEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
