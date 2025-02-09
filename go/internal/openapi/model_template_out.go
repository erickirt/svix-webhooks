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
	"time"
)

// checks if the TemplateOut type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateOut{}

// TemplateOut struct for TemplateOut
type TemplateOut struct {
	CreatedAt        time.Time     `json:"createdAt"`
	Description      string        `json:"description"`
	FeatureFlag      *string       `json:"featureFlag,omitempty" validate:"regexp=^[a-zA-Z0-9\\\\-_.]+$"`
	FilterTypes      []string      `json:"filterTypes,omitempty"`
	Id               string        `json:"id"`
	Instructions     string        `json:"instructions"`
	InstructionsLink *string       `json:"instructionsLink,omitempty"`
	Kind             ConnectorKind `json:"kind"`
	Logo             string        `json:"logo"`
	Name             string        `json:"name"`
	OrgId            string        `json:"orgId"`
	Transformation   string        `json:"transformation"`
	UpdatedAt        time.Time     `json:"updatedAt"`
}

type _TemplateOut TemplateOut

// NewTemplateOut instantiates a new TemplateOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateOut(createdAt time.Time, description string, id string, instructions string, kind ConnectorKind, logo string, name string, orgId string, transformation string, updatedAt time.Time) *TemplateOut {
	this := TemplateOut{}
	this.CreatedAt = createdAt
	this.Description = description
	this.Id = id
	this.Instructions = instructions
	this.Kind = kind
	this.Logo = logo
	this.Name = name
	this.OrgId = orgId
	this.Transformation = transformation
	this.UpdatedAt = updatedAt
	return &this
}

// NewTemplateOutWithDefaults instantiates a new TemplateOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateOutWithDefaults() *TemplateOut {
	this := TemplateOut{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *TemplateOut) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TemplateOut) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *TemplateOut) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *TemplateOut) SetDescription(v string) {
	o.Description = v
}

// GetFeatureFlag returns the FeatureFlag field value if set, zero value otherwise.
func (o *TemplateOut) GetFeatureFlag() string {
	if o == nil || IsNil(o.FeatureFlag) {
		var ret string
		return ret
	}
	return *o.FeatureFlag
}

// GetFeatureFlagOk returns a tuple with the FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetFeatureFlagOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureFlag) {
		return nil, false
	}
	return o.FeatureFlag, true
}

// HasFeatureFlag returns a boolean if a field has been set.
func (o *TemplateOut) HasFeatureFlag() bool {
	if o != nil && !IsNil(o.FeatureFlag) {
		return true
	}

	return false
}

// SetFeatureFlag gets a reference to the given string and assigns it to the FeatureFlag field.
func (o *TemplateOut) SetFeatureFlag(v string) {
	o.FeatureFlag = &v
}

// GetFilterTypes returns the FilterTypes field value if set, zero value otherwise.
func (o *TemplateOut) GetFilterTypes() []string {
	if o == nil || IsNil(o.FilterTypes) {
		var ret []string
		return ret
	}
	return o.FilterTypes
}

// GetFilterTypesOk returns a tuple with the FilterTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetFilterTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.FilterTypes) {
		return nil, false
	}
	return o.FilterTypes, true
}

// HasFilterTypes returns a boolean if a field has been set.
func (o *TemplateOut) HasFilterTypes() bool {
	if o != nil && !IsNil(o.FilterTypes) {
		return true
	}

	return false
}

// SetFilterTypes gets a reference to the given []string and assigns it to the FilterTypes field.
func (o *TemplateOut) SetFilterTypes(v []string) {
	o.FilterTypes = v
}

// GetId returns the Id field value
func (o *TemplateOut) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemplateOut) SetId(v string) {
	o.Id = v
}

// GetInstructions returns the Instructions field value
func (o *TemplateOut) GetInstructions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetInstructionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instructions, true
}

// SetInstructions sets field value
func (o *TemplateOut) SetInstructions(v string) {
	o.Instructions = v
}

// GetInstructionsLink returns the InstructionsLink field value if set, zero value otherwise.
func (o *TemplateOut) GetInstructionsLink() string {
	if o == nil || IsNil(o.InstructionsLink) {
		var ret string
		return ret
	}
	return *o.InstructionsLink
}

// GetInstructionsLinkOk returns a tuple with the InstructionsLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetInstructionsLinkOk() (*string, bool) {
	if o == nil || IsNil(o.InstructionsLink) {
		return nil, false
	}
	return o.InstructionsLink, true
}

// HasInstructionsLink returns a boolean if a field has been set.
func (o *TemplateOut) HasInstructionsLink() bool {
	if o != nil && !IsNil(o.InstructionsLink) {
		return true
	}

	return false
}

// SetInstructionsLink gets a reference to the given string and assigns it to the InstructionsLink field.
func (o *TemplateOut) SetInstructionsLink(v string) {
	o.InstructionsLink = &v
}

// GetKind returns the Kind field value
func (o *TemplateOut) GetKind() ConnectorKind {
	if o == nil {
		var ret ConnectorKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetKindOk() (*ConnectorKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TemplateOut) SetKind(v ConnectorKind) {
	o.Kind = v
}

// GetLogo returns the Logo field value
func (o *TemplateOut) GetLogo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logo, true
}

// SetLogo sets field value
func (o *TemplateOut) SetLogo(v string) {
	o.Logo = v
}

// GetName returns the Name field value
func (o *TemplateOut) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TemplateOut) SetName(v string) {
	o.Name = v
}

// GetOrgId returns the OrgId field value
func (o *TemplateOut) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *TemplateOut) SetOrgId(v string) {
	o.OrgId = v
}

// GetTransformation returns the Transformation field value
func (o *TemplateOut) GetTransformation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transformation
}

// GetTransformationOk returns a tuple with the Transformation field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetTransformationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transformation, true
}

// SetTransformation sets field value
func (o *TemplateOut) SetTransformation(v string) {
	o.Transformation = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TemplateOut) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TemplateOut) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TemplateOut) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o TemplateOut) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateOut) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["description"] = o.Description
	if !IsNil(o.FeatureFlag) {
		toSerialize["featureFlag"] = o.FeatureFlag
	}
	if !IsNil(o.FilterTypes) {
		toSerialize["filterTypes"] = o.FilterTypes
	}
	toSerialize["id"] = o.Id
	toSerialize["instructions"] = o.Instructions
	if !IsNil(o.InstructionsLink) {
		toSerialize["instructionsLink"] = o.InstructionsLink
	}
	toSerialize["kind"] = o.Kind
	toSerialize["logo"] = o.Logo
	toSerialize["name"] = o.Name
	toSerialize["orgId"] = o.OrgId
	toSerialize["transformation"] = o.Transformation
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *TemplateOut) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"description",
		"id",
		"instructions",
		"kind",
		"logo",
		"name",
		"orgId",
		"transformation",
		"updatedAt",
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

	varTemplateOut := _TemplateOut{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTemplateOut)

	if err != nil {
		return err
	}

	*o = TemplateOut(varTemplateOut)

	return err
}

type NullableTemplateOut struct {
	value *TemplateOut
	isSet bool
}

func (v NullableTemplateOut) Get() *TemplateOut {
	return v.value
}

func (v *NullableTemplateOut) Set(val *TemplateOut) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateOut) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateOut(val *TemplateOut) *NullableTemplateOut {
	return &NullableTemplateOut{value: val, isSet: true}
}

func (v NullableTemplateOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
