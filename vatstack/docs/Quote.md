# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] 
**AmountTotal** | Pointer to **int32** |  | [optional] 
**Category** | Pointer to **string** | Category of the digital product. Defaults to &#x60;null&#x60; if no category was specified in the request or if the category cannot be applied for the &#x60;country_code&#x60;. | [optional] 
**CountryCode** | Pointer to **string** | 2-letter ISO country code. | [optional] 
**CountryName** | Pointer to **string** | Corresponding English name of &#x60;country_code&#x60;. | [optional] 
**Created** | Pointer to **time.Time** | ISO date at which the object was created. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**IpAddress** | Pointer to **string** | The same IP address coming from the &#x60;ip_address&#x60; body parameter, or the geolocated IP address if none was provided. Value is &#x60;null&#x60; if &#x60;country_code&#x60; was provided. | [optional] 
**LocalName** | Pointer to **string** | Localized name of the VAT identification number. | [optional] 
**MemberState** | Pointer to **bool** | Boolean indicating whether &#x60;country_code&#x60; is an EU member state. | [optional] 
**Updated** | Pointer to **time.Time** | ISO date at which the object was updated. | [optional] 
**Validation** | Pointer to [**Validation**](Validation.md) |  | [optional] 
**Vat** | Pointer to [**QuoteVat**](QuoteVat.md) |  | [optional] 

## Methods

### NewQuote

`func NewQuote() *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Quote) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Quote) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Quote) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Quote) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountTotal

`func (o *Quote) GetAmountTotal() int32`

GetAmountTotal returns the AmountTotal field if non-nil, zero value otherwise.

### GetAmountTotalOk

`func (o *Quote) GetAmountTotalOk() (*int32, bool)`

GetAmountTotalOk returns a tuple with the AmountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotal

`func (o *Quote) SetAmountTotal(v int32)`

SetAmountTotal sets AmountTotal field to given value.

### HasAmountTotal

`func (o *Quote) HasAmountTotal() bool`

HasAmountTotal returns a boolean if a field has been set.

### GetCategory

`func (o *Quote) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Quote) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Quote) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Quote) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCountryCode

`func (o *Quote) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Quote) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Quote) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Quote) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *Quote) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *Quote) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *Quote) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *Quote) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCreated

`func (o *Quote) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Quote) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Quote) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Quote) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Quote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Quote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Quote) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Quote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *Quote) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Quote) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Quote) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Quote) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLocalName

`func (o *Quote) GetLocalName() string`

GetLocalName returns the LocalName field if non-nil, zero value otherwise.

### GetLocalNameOk

`func (o *Quote) GetLocalNameOk() (*string, bool)`

GetLocalNameOk returns a tuple with the LocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalName

`func (o *Quote) SetLocalName(v string)`

SetLocalName sets LocalName field to given value.

### HasLocalName

`func (o *Quote) HasLocalName() bool`

HasLocalName returns a boolean if a field has been set.

### GetMemberState

`func (o *Quote) GetMemberState() bool`

GetMemberState returns the MemberState field if non-nil, zero value otherwise.

### GetMemberStateOk

`func (o *Quote) GetMemberStateOk() (*bool, bool)`

GetMemberStateOk returns a tuple with the MemberState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberState

`func (o *Quote) SetMemberState(v bool)`

SetMemberState sets MemberState field to given value.

### HasMemberState

`func (o *Quote) HasMemberState() bool`

HasMemberState returns a boolean if a field has been set.

### GetUpdated

`func (o *Quote) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Quote) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Quote) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Quote) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValidation

`func (o *Quote) GetValidation() Validation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *Quote) GetValidationOk() (*Validation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *Quote) SetValidation(v Validation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *Quote) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### GetVat

`func (o *Quote) GetVat() QuoteVat`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *Quote) GetVatOk() (*QuoteVat, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *Quote) SetVat(v QuoteVat)`

SetVat sets Vat field to given value.

### HasVat

`func (o *Quote) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


