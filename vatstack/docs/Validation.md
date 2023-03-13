# Validation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Boolean indicating whether the company exists and is active. Use &#x60;valid&#x60; to check whether the business is also VAT-registered. | [optional] 
**Code** | Pointer to **string** | In the event of an error, this field will contain the error code. | [optional] 
**CompanyAddress** | Pointer to **string** | Address of the company the VAT number is associated with. Servers of Germany and Spain won’t return a value for privacy reasons and will default to &#x60;null&#x60;. | [optional] 
**CompanyName** | Pointer to **string** | Name of the company the VAT number is associated with. Servers of Germany and Spain won’t return a value for privacy reasons and will default to &#x60;null&#x60;. | [optional] 
**CompanyType** | Pointer to **string** | Type of the company entity returned by the respective government service (where available). | [optional] 
**ConsultationNumber** | Pointer to **string** | If you save your own VAT number in your dashboard, the reply will contain a unique consultation number. The consultation number enables you to prove to a tax administration of a Member State that you have checked a VAT number at the &#x60;requested&#x60; date, and obtained a validation result. | [optional] 
**CountryCode** | Pointer to **string** | 2-letter ISO country code. Note that while Greek VAT numbers contain the &#x60;EL&#x60; country code, our response will return the ISO country code &#x60;GR&#x60;. | [optional] 
**Created** | Pointer to **time.Time** | ISO date at which the object was created. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**Query** | **string** | Your original query. | 
**Requested** | Pointer to **time.Time** | ISO date at which the validation request was originally performed. Types &#x60;eu_vat&#x60; and &#x60;gb_vat&#x60; do not specify a time. | [optional] 
**Type** | Pointer to **string** | Type of VAT number. | [optional] 
**Updated** | Pointer to **time.Time** | ISO date at which the object was updated. | [optional] 
**Valid** | Pointer to **bool** | Boolean indicating whether the vat_number is registered for VAT. If government services are down, the value will be &#x60;null&#x60; and re-checked automatically for you. | [optional] 
**ValidFormat** | Pointer to **bool** | Boolean indicating whether the VAT number contained in &#x60;query&#x60; is in a valid format. | [optional] 
**VatNumber** | Pointer to **string** | VAT number extracted from your query without the country code. | [optional] 

## Methods

### NewValidation

`func NewValidation(query string, ) *Validation`

NewValidation instantiates a new Validation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationWithDefaults

`func NewValidationWithDefaults() *Validation`

NewValidationWithDefaults instantiates a new Validation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *Validation) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Validation) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Validation) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Validation) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCode

`func (o *Validation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Validation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Validation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Validation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCompanyAddress

`func (o *Validation) GetCompanyAddress() string`

GetCompanyAddress returns the CompanyAddress field if non-nil, zero value otherwise.

### GetCompanyAddressOk

`func (o *Validation) GetCompanyAddressOk() (*string, bool)`

GetCompanyAddressOk returns a tuple with the CompanyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress

`func (o *Validation) SetCompanyAddress(v string)`

SetCompanyAddress sets CompanyAddress field to given value.

### HasCompanyAddress

`func (o *Validation) HasCompanyAddress() bool`

HasCompanyAddress returns a boolean if a field has been set.

### GetCompanyName

`func (o *Validation) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Validation) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Validation) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Validation) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyType

`func (o *Validation) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *Validation) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *Validation) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *Validation) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetConsultationNumber

`func (o *Validation) GetConsultationNumber() string`

GetConsultationNumber returns the ConsultationNumber field if non-nil, zero value otherwise.

### GetConsultationNumberOk

`func (o *Validation) GetConsultationNumberOk() (*string, bool)`

GetConsultationNumberOk returns a tuple with the ConsultationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsultationNumber

`func (o *Validation) SetConsultationNumber(v string)`

SetConsultationNumber sets ConsultationNumber field to given value.

### HasConsultationNumber

`func (o *Validation) HasConsultationNumber() bool`

HasConsultationNumber returns a boolean if a field has been set.

### GetCountryCode

`func (o *Validation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Validation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Validation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Validation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreated

`func (o *Validation) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Validation) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Validation) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Validation) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Validation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Validation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Validation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Validation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuery

`func (o *Validation) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Validation) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Validation) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRequested

`func (o *Validation) GetRequested() time.Time`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *Validation) GetRequestedOk() (*time.Time, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *Validation) SetRequested(v time.Time)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *Validation) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetType

`func (o *Validation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Validation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Validation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Validation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *Validation) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Validation) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Validation) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Validation) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValid

`func (o *Validation) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *Validation) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *Validation) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *Validation) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetValidFormat

`func (o *Validation) GetValidFormat() bool`

GetValidFormat returns the ValidFormat field if non-nil, zero value otherwise.

### GetValidFormatOk

`func (o *Validation) GetValidFormatOk() (*bool, bool)`

GetValidFormatOk returns a tuple with the ValidFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFormat

`func (o *Validation) SetValidFormat(v bool)`

SetValidFormat sets ValidFormat field to given value.

### HasValidFormat

`func (o *Validation) HasValidFormat() bool`

HasValidFormat returns a boolean if a field has been set.

### GetVatNumber

`func (o *Validation) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *Validation) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *Validation) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *Validation) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


