# QuoteVat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abbreviation** | Pointer to **string** | Abbreviation of &#x60;vat.local_name&#x60;. | [optional] 
**Amount** | Pointer to **int32** | VAT amount in cents. | [optional] 
**Inclusive** | Pointer to **string** | Specifies if the &#x60;amount_total&#x60; is inclusive (common for EU consumers) or exclusive of VAT. This affects how the &#x60;vat.amount&#x60; is calculated. If &#x60;false&#x60;, you should present &#x60;amount&#x60; plus &#x60;vat.amount&#x60; to your customer as the final price to pay. | [optional] [default to "false"]
**LocalName** | Pointer to **string** | Localized name of the VAT. | [optional] 
**Rate** | Pointer to **float64** | VAT rate applied for the calculation. If member_state is &#x60;false&#x60;, the value will be &#x60;0&#x60;. | [optional] 
**RateType** | Pointer to **string** | Automatically determined type of VAT rate based on inputs. Can be &#x60;null&#x60;, &#x60;exempt&#x60;, &#x60;reduced&#x60;, &#x60;reverse_charge&#x60;, &#x60;standard&#x60; or &#x60;zero&#x60;. | [optional] 

## Methods

### NewQuoteVat

`func NewQuoteVat() *QuoteVat`

NewQuoteVat instantiates a new QuoteVat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteVatWithDefaults

`func NewQuoteVatWithDefaults() *QuoteVat`

NewQuoteVatWithDefaults instantiates a new QuoteVat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbbreviation

`func (o *QuoteVat) GetAbbreviation() string`

GetAbbreviation returns the Abbreviation field if non-nil, zero value otherwise.

### GetAbbreviationOk

`func (o *QuoteVat) GetAbbreviationOk() (*string, bool)`

GetAbbreviationOk returns a tuple with the Abbreviation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbbreviation

`func (o *QuoteVat) SetAbbreviation(v string)`

SetAbbreviation sets Abbreviation field to given value.

### HasAbbreviation

`func (o *QuoteVat) HasAbbreviation() bool`

HasAbbreviation returns a boolean if a field has been set.

### GetAmount

`func (o *QuoteVat) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuoteVat) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuoteVat) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuoteVat) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetInclusive

`func (o *QuoteVat) GetInclusive() string`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *QuoteVat) GetInclusiveOk() (*string, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *QuoteVat) SetInclusive(v string)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *QuoteVat) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.

### GetLocalName

`func (o *QuoteVat) GetLocalName() string`

GetLocalName returns the LocalName field if non-nil, zero value otherwise.

### GetLocalNameOk

`func (o *QuoteVat) GetLocalNameOk() (*string, bool)`

GetLocalNameOk returns a tuple with the LocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalName

`func (o *QuoteVat) SetLocalName(v string)`

SetLocalName sets LocalName field to given value.

### HasLocalName

`func (o *QuoteVat) HasLocalName() bool`

HasLocalName returns a boolean if a field has been set.

### GetRate

`func (o *QuoteVat) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *QuoteVat) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *QuoteVat) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *QuoteVat) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateType

`func (o *QuoteVat) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *QuoteVat) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *QuoteVat) SetRateType(v string)`

SetRateType sets RateType field to given value.

### HasRateType

`func (o *QuoteVat) HasRateType() bool`

HasRateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


