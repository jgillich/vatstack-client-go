/*
Vatstack API Specification

A suite of VAT APIs which let you validate VAT numbers with government services, look up VAT rates by ISO country code, calculate price quotes with centralized VAT rules and store transactions with each sale for VAT reporting.

API version: 1.3.0
Contact: team@vatstack.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vatstack

import (
	"encoding/json"
)

// checks if the GetAllSupplies200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllSupplies200Response{}

// GetAllSupplies200Response struct for GetAllSupplies200Response
type GetAllSupplies200Response struct {
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	SuppliesCount int32 `json:"supplies_count"`
	Supplies []Supply `json:"supplies"`
}

// NewGetAllSupplies200Response instantiates a new GetAllSupplies200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllSupplies200Response(hasMore bool, suppliesCount int32, supplies []Supply) *GetAllSupplies200Response {
	this := GetAllSupplies200Response{}
	this.HasMore = hasMore
	this.SuppliesCount = suppliesCount
	this.Supplies = supplies
	return &this
}

// NewGetAllSupplies200ResponseWithDefaults instantiates a new GetAllSupplies200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllSupplies200ResponseWithDefaults() *GetAllSupplies200Response {
	this := GetAllSupplies200Response{}
	return &this
}

// GetHasMore returns the HasMore field value
func (o *GetAllSupplies200Response) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *GetAllSupplies200Response) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *GetAllSupplies200Response) SetHasMore(v bool) {
	o.HasMore = v
}

// GetSuppliesCount returns the SuppliesCount field value
func (o *GetAllSupplies200Response) GetSuppliesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SuppliesCount
}

// GetSuppliesCountOk returns a tuple with the SuppliesCount field value
// and a boolean to check if the value has been set.
func (o *GetAllSupplies200Response) GetSuppliesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppliesCount, true
}

// SetSuppliesCount sets field value
func (o *GetAllSupplies200Response) SetSuppliesCount(v int32) {
	o.SuppliesCount = v
}

// GetSupplies returns the Supplies field value
func (o *GetAllSupplies200Response) GetSupplies() []Supply {
	if o == nil {
		var ret []Supply
		return ret
	}

	return o.Supplies
}

// GetSuppliesOk returns a tuple with the Supplies field value
// and a boolean to check if the value has been set.
func (o *GetAllSupplies200Response) GetSuppliesOk() ([]Supply, bool) {
	if o == nil {
		return nil, false
	}
	return o.Supplies, true
}

// SetSupplies sets field value
func (o *GetAllSupplies200Response) SetSupplies(v []Supply) {
	o.Supplies = v
}

func (o GetAllSupplies200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllSupplies200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["has_more"] = o.HasMore
	toSerialize["supplies_count"] = o.SuppliesCount
	toSerialize["supplies"] = o.Supplies
	return toSerialize, nil
}

type NullableGetAllSupplies200Response struct {
	value *GetAllSupplies200Response
	isSet bool
}

func (v NullableGetAllSupplies200Response) Get() *GetAllSupplies200Response {
	return v.value
}

func (v *NullableGetAllSupplies200Response) Set(val *GetAllSupplies200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllSupplies200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllSupplies200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllSupplies200Response(val *GetAllSupplies200Response) *NullableGetAllSupplies200Response {
	return &NullableGetAllSupplies200Response{value: val, isSet: true}
}

func (v NullableGetAllSupplies200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllSupplies200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


