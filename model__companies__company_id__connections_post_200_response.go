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

// checks if the CompaniesCompanyIdConnectionsPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompaniesCompanyIdConnectionsPost200Response{}

// CompaniesCompanyIdConnectionsPost200Response struct for CompaniesCompanyIdConnectionsPost200Response
type CompaniesCompanyIdConnectionsPost200Response struct {
	// Response status message
	Message string `json:"message"`
	Data CompaniesCompanyIdConnectionsPost200ResponseData `json:"data"`
}

type _CompaniesCompanyIdConnectionsPost200Response CompaniesCompanyIdConnectionsPost200Response

// NewCompaniesCompanyIdConnectionsPost200Response instantiates a new CompaniesCompanyIdConnectionsPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompaniesCompanyIdConnectionsPost200Response(message string, data CompaniesCompanyIdConnectionsPost200ResponseData) *CompaniesCompanyIdConnectionsPost200Response {
	this := CompaniesCompanyIdConnectionsPost200Response{}
	this.Message = message
	this.Data = data
	return &this
}

// NewCompaniesCompanyIdConnectionsPost200ResponseWithDefaults instantiates a new CompaniesCompanyIdConnectionsPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompaniesCompanyIdConnectionsPost200ResponseWithDefaults() *CompaniesCompanyIdConnectionsPost200Response {
	this := CompaniesCompanyIdConnectionsPost200Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *CompaniesCompanyIdConnectionsPost200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdConnectionsPost200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CompaniesCompanyIdConnectionsPost200Response) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *CompaniesCompanyIdConnectionsPost200Response) GetData() CompaniesCompanyIdConnectionsPost200ResponseData {
	if o == nil {
		var ret CompaniesCompanyIdConnectionsPost200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdConnectionsPost200Response) GetDataOk() (*CompaniesCompanyIdConnectionsPost200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CompaniesCompanyIdConnectionsPost200Response) SetData(v CompaniesCompanyIdConnectionsPost200ResponseData) {
	o.Data = v
}

func (o CompaniesCompanyIdConnectionsPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompaniesCompanyIdConnectionsPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CompaniesCompanyIdConnectionsPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"data",
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

	varCompaniesCompanyIdConnectionsPost200Response := _CompaniesCompanyIdConnectionsPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompaniesCompanyIdConnectionsPost200Response)

	if err != nil {
		return err
	}

	*o = CompaniesCompanyIdConnectionsPost200Response(varCompaniesCompanyIdConnectionsPost200Response)

	return err
}

type NullableCompaniesCompanyIdConnectionsPost200Response struct {
	value *CompaniesCompanyIdConnectionsPost200Response
	isSet bool
}

func (v NullableCompaniesCompanyIdConnectionsPost200Response) Get() *CompaniesCompanyIdConnectionsPost200Response {
	return v.value
}

func (v *NullableCompaniesCompanyIdConnectionsPost200Response) Set(val *CompaniesCompanyIdConnectionsPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCompaniesCompanyIdConnectionsPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCompaniesCompanyIdConnectionsPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompaniesCompanyIdConnectionsPost200Response(val *CompaniesCompanyIdConnectionsPost200Response) *NullableCompaniesCompanyIdConnectionsPost200Response {
	return &NullableCompaniesCompanyIdConnectionsPost200Response{value: val, isSet: true}
}

func (v NullableCompaniesCompanyIdConnectionsPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompaniesCompanyIdConnectionsPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


