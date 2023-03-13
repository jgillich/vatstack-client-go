# Rate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbreviation** | Pointer to **string** | Abbreviation of &#x60;local_name&#x60;. | [optional] 
**Categories** | Pointer to [**RateCategories**](RateCategories.md) |  | [optional] 
**CountryCode** | Pointer to **string** | 2-letter ISO country code. | [optional] 
**CountryName** | Pointer to **string** | Corresponding English name of &#x60;country_code&#x60;. | [optional] 
**Currency** | Pointer to **string** | 3-letter ISO 4217 local currency code. | [optional] 
**LocalName** | Pointer to **string** | Localized name of the VAT identification number. | [optional] 
**MemberState** | Pointer to **bool** | Boolean indicating whether the country is an EU Member State. | [optional] 
**ReducedRates** | Pointer to **[]float64** | 3-letter ISO 4217 local currency code. | [optional] 
**StandardRate** | Pointer to **float64** | Standard VAT rate in percent. | [optional] 
**VatAbbreviation** | Pointer to **string** | Abbreviation of &#x60;vat_local_name&#x60;. | [optional] 
**VatLocalName** | Pointer to **string** | Localized name of the VAT. | [optional] 

## Methods

### NewRate

`func NewRate() *Rate`

NewRate instantiates a new Rate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateWithDefaults

`func NewRateWithDefaults() *Rate`

NewRateWithDefaults instantiates a new Rate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *Rate) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *Rate) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *Rate) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *Rate) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### GetCategories

`func (o *Rate) GetCategories() RateCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Rate) GetCategoriesOk() (*RateCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Rate) SetCategories(v RateCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Rate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCountryCode

`func (o *Rate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Rate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Rate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Rate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *Rate) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *Rate) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *Rate) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *Rate) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCurrency

`func (o *Rate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Rate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Rate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Rate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetLocalName

`func (o *Rate) GetLocalName() string`

GetLocalName returns the LocalName field if non-nil, zero value otherwise.

### GetLocalNameOk

`func (o *Rate) GetLocalNameOk() (*string, bool)`

GetLocalNameOk returns a tuple with the LocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalName

`func (o *Rate) SetLocalName(v string)`

SetLocalName sets LocalName field to given value.

### HasLocalName

`func (o *Rate) HasLocalName() bool`

HasLocalName returns a boolean if a field has been set.

### GetMemberState

`func (o *Rate) GetMemberState() bool`

GetMemberState returns the MemberState field if non-nil, zero value otherwise.

### GetMemberStateOk

`func (o *Rate) GetMemberStateOk() (*bool, bool)`

GetMemberStateOk returns a tuple with the MemberState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberState

`func (o *Rate) SetMemberState(v bool)`

SetMemberState sets MemberState field to given value.

### HasMemberState

`func (o *Rate) HasMemberState() bool`

HasMemberState returns a boolean if a field has been set.

### GetReducedRates

`func (o *Rate) GetReducedRates() []float64`

GetReducedRates returns the ReducedRates field if non-nil, zero value otherwise.

### GetReducedRatesOk

`func (o *Rate) GetReducedRatesOk() (*[]float64, bool)`

GetReducedRatesOk returns a tuple with the ReducedRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedRates

`func (o *Rate) SetReducedRates(v []float64)`

SetReducedRates sets ReducedRates field to given value.

### HasReducedRates

`func (o *Rate) HasReducedRates() bool`

HasReducedRates returns a boolean if a field has been set.

### GetStandardRate

`func (o *Rate) GetStandardRate() float64`

GetStandardRate returns the StandardRate field if non-nil, zero value otherwise.

### GetStandardRateOk

`func (o *Rate) GetStandardRateOk() (*float64, bool)`

GetStandardRateOk returns a tuple with the StandardRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardRate

`func (o *Rate) SetStandardRate(v float64)`

SetStandardRate sets StandardRate field to given value.

### HasStandardRate

`func (o *Rate) HasStandardRate() bool`

HasStandardRate returns a boolean if a field has been set.

### GetVatAbbreviation

`func (o *Rate) GetVatAbbreviation() string`

GetVatAbbreviation returns the VatAbbreviation field if non-nil, zero value otherwise.

### GetVatAbbreviationOk

`func (o *Rate) GetVatAbbreviationOk() (*string, bool)`

GetVatAbbreviationOk returns a tuple with the VatAbbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatAbbreviation

`func (o *Rate) SetVatAbbreviation(v string)`

SetVatAbbreviation sets VatAbbreviation field to given value.

### HasVatAbbreviation

`func (o *Rate) HasVatAbbreviation() bool`

HasVatAbbreviation returns a boolean if a field has been set.

### GetVatLocalName

`func (o *Rate) GetVatLocalName() string`

GetVatLocalName returns the VatLocalName field if non-nil, zero value otherwise.

### GetVatLocalNameOk

`func (o *Rate) GetVatLocalNameOk() (*string, bool)`

GetVatLocalNameOk returns a tuple with the VatLocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatLocalName

`func (o *Rate) SetVatLocalName(v string)`

SetVatLocalName sets VatLocalName field to given value.

### HasVatLocalName

`func (o *Rate) HasVatLocalName() bool`

HasVatLocalName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


