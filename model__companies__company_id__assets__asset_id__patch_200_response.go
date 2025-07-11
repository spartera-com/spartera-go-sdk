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

// checks if the CompaniesCompanyIdAssetsAssetIdPatch200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompaniesCompanyIdAssetsAssetIdPatch200Response{}

// CompaniesCompanyIdAssetsAssetIdPatch200Response struct for CompaniesCompanyIdAssetsAssetIdPatch200Response
type CompaniesCompanyIdAssetsAssetIdPatch200Response struct {
	// Response status message
	Message string `json:"message"`
	Data CompaniesCompanyIdAssetsAssetIdPatch200ResponseData `json:"data"`
}

type _CompaniesCompanyIdAssetsAssetIdPatch200Response CompaniesCompanyIdAssetsAssetIdPatch200Response

// NewCompaniesCompanyIdAssetsAssetIdPatch200Response instantiates a new CompaniesCompanyIdAssetsAssetIdPatch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompaniesCompanyIdAssetsAssetIdPatch200Response(message string, data CompaniesCompanyIdAssetsAssetIdPatch200ResponseData) *CompaniesCompanyIdAssetsAssetIdPatch200Response {
	this := CompaniesCompanyIdAssetsAssetIdPatch200Response{}
	this.Message = message
	this.Data = data
	return &this
}

// NewCompaniesCompanyIdAssetsAssetIdPatch200ResponseWithDefaults instantiates a new CompaniesCompanyIdAssetsAssetIdPatch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompaniesCompanyIdAssetsAssetIdPatch200ResponseWithDefaults() *CompaniesCompanyIdAssetsAssetIdPatch200Response {
	this := CompaniesCompanyIdAssetsAssetIdPatch200Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) GetData() CompaniesCompanyIdAssetsAssetIdPatch200ResponseData {
	if o == nil {
		var ret CompaniesCompanyIdAssetsAssetIdPatch200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) GetDataOk() (*CompaniesCompanyIdAssetsAssetIdPatch200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) SetData(v CompaniesCompanyIdAssetsAssetIdPatch200ResponseData) {
	o.Data = v
}

func (o CompaniesCompanyIdAssetsAssetIdPatch200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompaniesCompanyIdAssetsAssetIdPatch200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CompaniesCompanyIdAssetsAssetIdPatch200Response) UnmarshalJSON(data []byte) (err error) {
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

	varCompaniesCompanyIdAssetsAssetIdPatch200Response := _CompaniesCompanyIdAssetsAssetIdPatch200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompaniesCompanyIdAssetsAssetIdPatch200Response)

	if err != nil {
		return err
	}

	*o = CompaniesCompanyIdAssetsAssetIdPatch200Response(varCompaniesCompanyIdAssetsAssetIdPatch200Response)

	return err
}

type NullableCompaniesCompanyIdAssetsAssetIdPatch200Response struct {
	value *CompaniesCompanyIdAssetsAssetIdPatch200Response
	isSet bool
}

func (v NullableCompaniesCompanyIdAssetsAssetIdPatch200Response) Get() *CompaniesCompanyIdAssetsAssetIdPatch200Response {
	return v.value
}

func (v *NullableCompaniesCompanyIdAssetsAssetIdPatch200Response) Set(val *CompaniesCompanyIdAssetsAssetIdPatch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCompaniesCompanyIdAssetsAssetIdPatch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCompaniesCompanyIdAssetsAssetIdPatch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompaniesCompanyIdAssetsAssetIdPatch200Response(val *CompaniesCompanyIdAssetsAssetIdPatch200Response) *NullableCompaniesCompanyIdAssetsAssetIdPatch200Response {
	return &NullableCompaniesCompanyIdAssetsAssetIdPatch200Response{value: val, isSet: true}
}

func (v NullableCompaniesCompanyIdAssetsAssetIdPatch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompaniesCompanyIdAssetsAssetIdPatch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


