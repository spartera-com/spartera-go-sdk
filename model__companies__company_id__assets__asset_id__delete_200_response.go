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

// checks if the CompaniesCompanyIdAssetsAssetIdDelete200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompaniesCompanyIdAssetsAssetIdDelete200Response{}

// CompaniesCompanyIdAssetsAssetIdDelete200Response struct for CompaniesCompanyIdAssetsAssetIdDelete200Response
type CompaniesCompanyIdAssetsAssetIdDelete200Response struct {
	// Response status message
	Message string `json:"message"`
	Data CompaniesCompanyIdAssetsAssetIdDelete200ResponseData `json:"data"`
}

type _CompaniesCompanyIdAssetsAssetIdDelete200Response CompaniesCompanyIdAssetsAssetIdDelete200Response

// NewCompaniesCompanyIdAssetsAssetIdDelete200Response instantiates a new CompaniesCompanyIdAssetsAssetIdDelete200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompaniesCompanyIdAssetsAssetIdDelete200Response(message string, data CompaniesCompanyIdAssetsAssetIdDelete200ResponseData) *CompaniesCompanyIdAssetsAssetIdDelete200Response {
	this := CompaniesCompanyIdAssetsAssetIdDelete200Response{}
	this.Message = message
	this.Data = data
	return &this
}

// NewCompaniesCompanyIdAssetsAssetIdDelete200ResponseWithDefaults instantiates a new CompaniesCompanyIdAssetsAssetIdDelete200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompaniesCompanyIdAssetsAssetIdDelete200ResponseWithDefaults() *CompaniesCompanyIdAssetsAssetIdDelete200Response {
	this := CompaniesCompanyIdAssetsAssetIdDelete200Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) SetMessage(v string) {
	o.Message = v
}

// GetData returns the Data field value
func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) GetData() CompaniesCompanyIdAssetsAssetIdDelete200ResponseData {
	if o == nil {
		var ret CompaniesCompanyIdAssetsAssetIdDelete200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) GetDataOk() (*CompaniesCompanyIdAssetsAssetIdDelete200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) SetData(v CompaniesCompanyIdAssetsAssetIdDelete200ResponseData) {
	o.Data = v
}

func (o CompaniesCompanyIdAssetsAssetIdDelete200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompaniesCompanyIdAssetsAssetIdDelete200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CompaniesCompanyIdAssetsAssetIdDelete200Response) UnmarshalJSON(data []byte) (err error) {
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

	varCompaniesCompanyIdAssetsAssetIdDelete200Response := _CompaniesCompanyIdAssetsAssetIdDelete200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompaniesCompanyIdAssetsAssetIdDelete200Response)

	if err != nil {
		return err
	}

	*o = CompaniesCompanyIdAssetsAssetIdDelete200Response(varCompaniesCompanyIdAssetsAssetIdDelete200Response)

	return err
}

type NullableCompaniesCompanyIdAssetsAssetIdDelete200Response struct {
	value *CompaniesCompanyIdAssetsAssetIdDelete200Response
	isSet bool
}

func (v NullableCompaniesCompanyIdAssetsAssetIdDelete200Response) Get() *CompaniesCompanyIdAssetsAssetIdDelete200Response {
	return v.value
}

func (v *NullableCompaniesCompanyIdAssetsAssetIdDelete200Response) Set(val *CompaniesCompanyIdAssetsAssetIdDelete200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCompaniesCompanyIdAssetsAssetIdDelete200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCompaniesCompanyIdAssetsAssetIdDelete200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompaniesCompanyIdAssetsAssetIdDelete200Response(val *CompaniesCompanyIdAssetsAssetIdDelete200Response) *NullableCompaniesCompanyIdAssetsAssetIdDelete200Response {
	return &NullableCompaniesCompanyIdAssetsAssetIdDelete200Response{value: val, isSet: true}
}

func (v NullableCompaniesCompanyIdAssetsAssetIdDelete200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompaniesCompanyIdAssetsAssetIdDelete200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


