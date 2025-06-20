/*
Spartera API Documentation

Auto-generated API documentation for REST services of the Spartera platform

API version: 0.0.0
Contact: support@spartera.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sparteraapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CompaniesCompanyIdApiKeysGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompaniesCompanyIdApiKeysGet200Response{}

// CompaniesCompanyIdApiKeysGet200Response struct for CompaniesCompanyIdApiKeysGet200Response
type CompaniesCompanyIdApiKeysGet200Response struct {
	// Response status message
	Message string `json:"message"`
	// Response data
	Data map[string]interface{} `json:"data,omitempty"`
}

type _CompaniesCompanyIdApiKeysGet200Response CompaniesCompanyIdApiKeysGet200Response

// NewCompaniesCompanyIdApiKeysGet200Response instantiates a new CompaniesCompanyIdApiKeysGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompaniesCompanyIdApiKeysGet200Response(message string) *CompaniesCompanyIdApiKeysGet200Response {
	this := CompaniesCompanyIdApiKeysGet200Response{}
	this.Message = message
	return &this
}

// NewCompaniesCompanyIdApiKeysGet200ResponseWithDefaults instantiates a new CompaniesCompanyIdApiKeysGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompaniesCompanyIdApiKeysGet200ResponseWithDefaults() *CompaniesCompanyIdApiKeysGet200Response {
	this := CompaniesCompanyIdApiKeysGet200Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *CompaniesCompanyIdApiKeysGet200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdApiKeysGet200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CompaniesCompanyIdApiKeysGet200Response) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CompaniesCompanyIdApiKeysGet200Response) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdApiKeysGet200Response) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CompaniesCompanyIdApiKeysGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CompaniesCompanyIdApiKeysGet200Response) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o CompaniesCompanyIdApiKeysGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompaniesCompanyIdApiKeysGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

func (o *CompaniesCompanyIdApiKeysGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCompaniesCompanyIdApiKeysGet200Response := _CompaniesCompanyIdApiKeysGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompaniesCompanyIdApiKeysGet200Response)

	if err != nil {
		return err
	}

	*o = CompaniesCompanyIdApiKeysGet200Response(varCompaniesCompanyIdApiKeysGet200Response)

	return err
}

type NullableCompaniesCompanyIdApiKeysGet200Response struct {
	value *CompaniesCompanyIdApiKeysGet200Response
	isSet bool
}

func (v NullableCompaniesCompanyIdApiKeysGet200Response) Get() *CompaniesCompanyIdApiKeysGet200Response {
	return v.value
}

func (v *NullableCompaniesCompanyIdApiKeysGet200Response) Set(val *CompaniesCompanyIdApiKeysGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCompaniesCompanyIdApiKeysGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCompaniesCompanyIdApiKeysGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompaniesCompanyIdApiKeysGet200Response(val *CompaniesCompanyIdApiKeysGet200Response) *NullableCompaniesCompanyIdApiKeysGet200Response {
	return &NullableCompaniesCompanyIdApiKeysGet200Response{value: val, isSet: true}
}

func (v NullableCompaniesCompanyIdApiKeysGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompaniesCompanyIdApiKeysGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


