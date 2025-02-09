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

// checks if the SnowflakeConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnowflakeConfig{}

// SnowflakeConfig Configuration parameters for defining a Snowflake sink.
type SnowflakeConfig struct {
	// Snowflake account identifier, which includes both the organization and account IDs separated by a hyphen.
	AccountIdentifier string `json:"accountIdentifier"`
	// Database name.  Only required if not using transformations.
	DbName *string `json:"dbName,omitempty"`
	// PEM-encoded private key used for signing token-based requests to the Snowflake API.  Beginning/end delimiters are not required.
	PrivateKey string `json:"privateKey"`
	// Schema name.  Only required if not using transformations.
	SchemaName *string `json:"schemaName,omitempty"`
	// Table name.  Only required if not using transformations.
	TableName *string `json:"tableName,omitempty"`
	// The Snowflake user id.
	UserId string `json:"userId"`
}

type _SnowflakeConfig SnowflakeConfig

// NewSnowflakeConfig instantiates a new SnowflakeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnowflakeConfig(accountIdentifier string, privateKey string, userId string) *SnowflakeConfig {
	this := SnowflakeConfig{}
	this.AccountIdentifier = accountIdentifier
	this.PrivateKey = privateKey
	this.UserId = userId
	return &this
}

// NewSnowflakeConfigWithDefaults instantiates a new SnowflakeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnowflakeConfigWithDefaults() *SnowflakeConfig {
	this := SnowflakeConfig{}
	return &this
}

// GetAccountIdentifier returns the AccountIdentifier field value
func (o *SnowflakeConfig) GetAccountIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountIdentifier
}

// GetAccountIdentifierOk returns a tuple with the AccountIdentifier field value
// and a boolean to check if the value has been set.
func (o *SnowflakeConfig) GetAccountIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountIdentifier, true
}

// SetAccountIdentifier sets field value
func (o *SnowflakeConfig) SetAccountIdentifier(v string) {
	o.AccountIdentifier = v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *SnowflakeConfig) GetDbName() string {
	if o == nil || IsNil(o.DbName) {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeConfig) GetDbNameOk() (*string, bool) {
	if o == nil || IsNil(o.DbName) {
		return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *SnowflakeConfig) HasDbName() bool {
	if o != nil && !IsNil(o.DbName) {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *SnowflakeConfig) SetDbName(v string) {
	o.DbName = &v
}

// GetPrivateKey returns the PrivateKey field value
func (o *SnowflakeConfig) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *SnowflakeConfig) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *SnowflakeConfig) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *SnowflakeConfig) GetSchemaName() string {
	if o == nil || IsNil(o.SchemaName) {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeConfig) GetSchemaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaName) {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *SnowflakeConfig) HasSchemaName() bool {
	if o != nil && !IsNil(o.SchemaName) {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *SnowflakeConfig) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *SnowflakeConfig) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeConfig) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *SnowflakeConfig) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *SnowflakeConfig) SetTableName(v string) {
	o.TableName = &v
}

// GetUserId returns the UserId field value
func (o *SnowflakeConfig) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SnowflakeConfig) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SnowflakeConfig) SetUserId(v string) {
	o.UserId = v
}

func (o SnowflakeConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnowflakeConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountIdentifier"] = o.AccountIdentifier
	if !IsNil(o.DbName) {
		toSerialize["dbName"] = o.DbName
	}
	toSerialize["privateKey"] = o.PrivateKey
	if !IsNil(o.SchemaName) {
		toSerialize["schemaName"] = o.SchemaName
	}
	if !IsNil(o.TableName) {
		toSerialize["tableName"] = o.TableName
	}
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *SnowflakeConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountIdentifier",
		"privateKey",
		"userId",
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

	varSnowflakeConfig := _SnowflakeConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSnowflakeConfig)

	if err != nil {
		return err
	}

	*o = SnowflakeConfig(varSnowflakeConfig)

	return err
}

type NullableSnowflakeConfig struct {
	value *SnowflakeConfig
	isSet bool
}

func (v NullableSnowflakeConfig) Get() *SnowflakeConfig {
	return v.value
}

func (v *NullableSnowflakeConfig) Set(val *SnowflakeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSnowflakeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSnowflakeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnowflakeConfig(val *SnowflakeConfig) *NullableSnowflakeConfig {
	return &NullableSnowflakeConfig{value: val, isSet: true}
}

func (v NullableSnowflakeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnowflakeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
