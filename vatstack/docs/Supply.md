# Supply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**AmountRefunded** | Pointer to **int32** |  | [optional] 
**AmountTotal** | Pointer to **int32** |  | [optional] 
**CountryCode** | **string** | 2-letter ISO country code of the place of supply that is relevant for the &#x60;vat.rate&#x60;. | 
**Created** | Pointer to **time.Time** | ISO date at which the object was created. | [optional] 
**Currency** | **string** | 3-letter ISO 4217 currency code used to charge the &#x60;amount&#x60;. | 
**Description** | Pointer to **string** | An arbitrary string to describe the supplied item. Often useful for displaying to users. | [optional] 
**Evidence** | Pointer to [**Evidence**](Evidence.md) |  | [optional] 
**EvidenceStatus** | Pointer to **string** | Status of whether the attached evidence object sufficiently proves the place of supply established in &#x60;country_code&#x60;. Will be either &#x60;sufficient&#x60; or &#x60;insufficient&#x60;. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**InvoiceNumber** | **string** | A custom string for the invoice number issued to the customer. It’s advisable to follow sequential numbering. | 
**Issued** | **time.Time** | ISO date at which &#x60;invoice_number&#x60; was issued. | 
**Name** | Pointer to **string** | A custom string for the name of the customer. | [optional] 
**Notes** | Pointer to **string** | A custom string for additional notes. | [optional] 
**Updated** | Pointer to **time.Time** | ISO date at which the object was updated. | [optional] 
**Validation** | Pointer to [**Validation**](Validation.md) |  | [optional] 
**Vat** | Pointer to [**SupplyVat**](SupplyVat.md) |  | [optional] 

## Methods

### NewSupply

`func NewSupply(amount int32, countryCode string, currency string, invoiceNumber string, issued time.Time, ) *Supply`

NewSupply instantiates a new Supply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyWithDefaults

`func NewSupplyWithDefaults() *Supply`

NewSupplyWithDefaults instantiates a new Supply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *Supply) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Supply) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Supply) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAmountRefunded

`func (o *Supply) GetAmountRefunded() int32`

GetAmountRefunded returns the AmountRefunded field if non-nil, zero value otherwise.

### GetAmountRefundedOk

`func (o *Supply) GetAmountRefundedOk() (*int32, bool)`

GetAmountRefundedOk returns a tuple with the AmountRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefunded

`func (o *Supply) SetAmountRefunded(v int32)`

SetAmountRefunded sets AmountRefunded field to given value.

### HasAmountRefunded

`func (o *Supply) HasAmountRefunded() bool`

HasAmountRefunded returns a boolean if a field has been set.

### GetAmountTotal

`func (o *Supply) GetAmountTotal() int32`

GetAmountTotal returns the AmountTotal field if non-nil, zero value otherwise.

### GetAmountTotalOk

`func (o *Supply) GetAmountTotalOk() (*int32, bool)`

GetAmountTotalOk returns a tuple with the AmountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountTotal

`func (o *Supply) SetAmountTotal(v int32)`

SetAmountTotal sets AmountTotal field to given value.

### HasAmountTotal

`func (o *Supply) HasAmountTotal() bool`

HasAmountTotal returns a boolean if a field has been set.

### GetCountryCode

`func (o *Supply) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Supply) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Supply) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCreated

`func (o *Supply) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Supply) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Supply) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Supply) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCurrency

`func (o *Supply) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Supply) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Supply) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *Supply) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Supply) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Supply) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Supply) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvidence

`func (o *Supply) GetEvidence() Evidence`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *Supply) GetEvidenceOk() (*Evidence, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *Supply) SetEvidence(v Evidence)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *Supply) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetEvidenceStatus

`func (o *Supply) GetEvidenceStatus() string`

GetEvidenceStatus returns the EvidenceStatus field if non-nil, zero value otherwise.

### GetEvidenceStatusOk

`func (o *Supply) GetEvidenceStatusOk() (*string, bool)`

GetEvidenceStatusOk returns a tuple with the EvidenceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceStatus

`func (o *Supply) SetEvidenceStatus(v string)`

SetEvidenceStatus sets EvidenceStatus field to given value.

### HasEvidenceStatus

`func (o *Supply) HasEvidenceStatus() bool`

HasEvidenceStatus returns a boolean if a field has been set.

### GetId

`func (o *Supply) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Supply) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Supply) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Supply) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *Supply) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *Supply) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *Supply) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetIssued

`func (o *Supply) GetIssued() time.Time`

GetIssued returns the Issued field if non-nil, zero value otherwise.

### GetIssuedOk

`func (o *Supply) GetIssuedOk() (*time.Time, bool)`

GetIssuedOk returns a tuple with the Issued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssued

`func (o *Supply) SetIssued(v time.Time)`

SetIssued sets Issued field to given value.


### GetName

`func (o *Supply) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Supply) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Supply) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Supply) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *Supply) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Supply) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Supply) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Supply) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUpdated

`func (o *Supply) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Supply) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Supply) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Supply) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValidation

`func (o *Supply) GetValidation() Validation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *Supply) GetValidationOk() (*Validation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *Supply) SetValidation(v Validation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *Supply) HasValidation() bool`

HasValidation returns a boolean if a field has been set.

### GetVat

`func (o *Supply) GetVat() SupplyVat`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *Supply) GetVatOk() (*SupplyVat, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *Supply) SetVat(v SupplyVat)`

SetVat sets Vat field to given value.

### HasVat

`func (o *Supply) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


