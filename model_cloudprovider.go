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

// checks if the Cloudprovider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cloudprovider{}

// Cloudprovider Cloud providers (platforms) database
type Cloudprovider struct {
	ProviderId *string `json:"provider_id,omitempty"`
	Name string `json:"name"`
	ParentCompany *string `json:"parent_company,omitempty"`
	MarketingHomepageUrl *string `json:"marketing_homepage_url,omitempty"`
	DateCreated *string `json:"date_created,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	Active *string `json:"active,omitempty"`
}

type _Cloudprovider Cloudprovider

// NewCloudprovider instantiates a new Cloudprovider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudprovider(name string) *Cloudprovider {
	this := Cloudprovider{}
	this.Name = name
	return &this
}

// NewCloudproviderWithDefaults instantiates a new Cloudprovider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudproviderWithDefaults() *Cloudprovider {
	this := Cloudprovider{}
	return &this
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *Cloudprovider) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// HasProviderId returns a boolean if a field has been set.
func (o *Cloudprovider) HasProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *Cloudprovider) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetName returns the Name field value
func (o *Cloudprovider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Cloudprovider) SetName(v string) {
	o.Name = v
}

// GetParentCompany returns the ParentCompany field value if set, zero value otherwise.
func (o *Cloudprovider) GetParentCompany() string {
	if o == nil || IsNil(o.ParentCompany) {
		var ret string
		return ret
	}
	return *o.ParentCompany
}

// GetParentCompanyOk returns a tuple with the ParentCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetParentCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.ParentCompany) {
		return nil, false
	}
	return o.ParentCompany, true
}

// HasParentCompany returns a boolean if a field has been set.
func (o *Cloudprovider) HasParentCompany() bool {
	if o != nil && !IsNil(o.ParentCompany) {
		return true
	}

	return false
}

// SetParentCompany gets a reference to the given string and assigns it to the ParentCompany field.
func (o *Cloudprovider) SetParentCompany(v string) {
	o.ParentCompany = &v
}

// GetMarketingHomepageUrl returns the MarketingHomepageUrl field value if set, zero value otherwise.
func (o *Cloudprovider) GetMarketingHomepageUrl() string {
	if o == nil || IsNil(o.MarketingHomepageUrl) {
		var ret string
		return ret
	}
	return *o.MarketingHomepageUrl
}

// GetMarketingHomepageUrlOk returns a tuple with the MarketingHomepageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetMarketingHomepageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.MarketingHomepageUrl) {
		return nil, false
	}
	return o.MarketingHomepageUrl, true
}

// HasMarketingHomepageUrl returns a boolean if a field has been set.
func (o *Cloudprovider) HasMarketingHomepageUrl() bool {
	if o != nil && !IsNil(o.MarketingHomepageUrl) {
		return true
	}

	return false
}

// SetMarketingHomepageUrl gets a reference to the given string and assigns it to the MarketingHomepageUrl field.
func (o *Cloudprovider) SetMarketingHomepageUrl(v string) {
	o.MarketingHomepageUrl = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Cloudprovider) GetDateCreated() string {
	if o == nil || IsNil(o.DateCreated) {
		var ret string
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetDateCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Cloudprovider) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given string and assigns it to the DateCreated field.
func (o *Cloudprovider) SetDateCreated(v string) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Cloudprovider) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Cloudprovider) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *Cloudprovider) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Cloudprovider) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cloudprovider) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *Cloudprovider) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *Cloudprovider) SetActive(v string) {
	o.Active = &v
}

func (o Cloudprovider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cloudprovider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderId) {
		toSerialize["provider_id"] = o.ProviderId
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentCompany) {
		toSerialize["parent_company"] = o.ParentCompany
	}
	if !IsNil(o.MarketingHomepageUrl) {
		toSerialize["marketing_homepage_url"] = o.MarketingHomepageUrl
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["last_updated"] = o.LastUpdated
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

func (o *Cloudprovider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCloudprovider := _Cloudprovider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCloudprovider)

	if err != nil {
		return err
	}

	*o = Cloudprovider(varCloudprovider)

	return err
}

type NullableCloudprovider struct {
	value *Cloudprovider
	isSet bool
}

func (v NullableCloudprovider) Get() *Cloudprovider {
	return v.value
}

func (v *NullableCloudprovider) Set(val *Cloudprovider) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudprovider) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudprovider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudprovider(val *Cloudprovider) *NullableCloudprovider {
	return &NullableCloudprovider{value: val, isSet: true}
}

func (v NullableCloudprovider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudprovider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


