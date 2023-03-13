# EvidenceBankAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | 2-letter ISO country code of the bank or payment source. | [optional] 
**Name** | Pointer to **string** | Name of the bank or payment source. | [optional] 

## Methods

### NewEvidenceBankAddress

`func NewEvidenceBankAddress() *EvidenceBankAddress`

NewEvidenceBankAddress instantiates a new EvidenceBankAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceBankAddressWithDefaults

`func NewEvidenceBankAddressWithDefaults() *EvidenceBankAddress`

NewEvidenceBankAddressWithDefaults instantiates a new EvidenceBankAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *EvidenceBankAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *EvidenceBankAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *EvidenceBankAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *EvidenceBankAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetName

`func (o *EvidenceBankAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvidenceBankAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvidenceBankAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvidenceBankAddress) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


