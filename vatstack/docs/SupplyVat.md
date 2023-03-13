# SupplyVat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | VAT amount in cents. | [optional] 
**Inclusive** | Pointer to **string** | Specifies if the &#x60;amount_total&#x60; is inclusive (common for EU consumers) or exclusive of VAT. This affects how the &#x60;vat.amount&#x60; is calculated. | [optional] [default to "false"]
**Rate** | Pointer to **float64** | VAT rate applicable for the place of supply established in &#x60;country_code&#x60;. | [optional] 
**RateType** | Pointer to **string** | Automatically determined type of VAT rate based on inputs. Can be &#x60;null&#x60;, &#x60;exempt&#x60;, &#x60;reduced&#x60;, &#x60;reverse_charge&#x60;, &#x60;standard&#x60; or &#x60;zero&#x60;. | [optional] 

## Methods

### NewSupplyVat

`func NewSupplyVat() *SupplyVat`

NewSupplyVat instantiates a new SupplyVat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyVatWithDefaults

`func NewSupplyVatWithDefaults() *SupplyVat`

NewSupplyVatWithDefaults instantiates a new SupplyVat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SupplyVat) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SupplyVat) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SupplyVat) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SupplyVat) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetInclusive

`func (o *SupplyVat) GetInclusive() string`

GetInclusive returns the Inclusive field if non-nil, zero value otherwise.

### GetInclusiveOk

`func (o *SupplyVat) GetInclusiveOk() (*string, bool)`

GetInclusiveOk returns a tuple with the Inclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusive

`func (o *SupplyVat) SetInclusive(v string)`

SetInclusive sets Inclusive field to given value.

### HasInclusive

`func (o *SupplyVat) HasInclusive() bool`

HasInclusive returns a boolean if a field has been set.

### GetRate

`func (o *SupplyVat) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *SupplyVat) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *SupplyVat) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *SupplyVat) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateType

`func (o *SupplyVat) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *SupplyVat) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *SupplyVat) SetRateType(v string)`

SetRateType sets RateType field to given value.

### HasRateType

`func (o *SupplyVat) HasRateType() bool`

HasRateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


