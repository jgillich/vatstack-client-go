# Evidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankAddress** | Pointer to [**EvidenceBankAddress**](EvidenceBankAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**EvidenceBillingAddress**](EvidenceBillingAddress.md) |  | [optional] 
**Created** | Pointer to **time.Time** | ISO date at which the object was created. | [optional] 
**IpAddress** | Pointer to [**EvidenceIpAddress**](EvidenceIpAddress.md) |  | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**RequiredCount** | Pointer to **int32** | Required pieces of evidence according to your account’s regional settings. | [optional] 
**Updated** | Pointer to **time.Time** | ISO date at which the object was updated. | [optional] 

## Methods

### NewEvidence

`func NewEvidence() *Evidence`

NewEvidence instantiates a new Evidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceWithDefaults

`func NewEvidenceWithDefaults() *Evidence`

NewEvidenceWithDefaults instantiates a new Evidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankAddress

`func (o *Evidence) GetBankAddress() EvidenceBankAddress`

GetBankAddress returns the BankAddress field if non-nil, zero value otherwise.

### GetBankAddressOk

`func (o *Evidence) GetBankAddressOk() (*EvidenceBankAddress, bool)`

GetBankAddressOk returns a tuple with the BankAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAddress

`func (o *Evidence) SetBankAddress(v EvidenceBankAddress)`

SetBankAddress sets BankAddress field to given value.

### HasBankAddress

`func (o *Evidence) HasBankAddress() bool`

HasBankAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *Evidence) GetBillingAddress() EvidenceBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *Evidence) GetBillingAddressOk() (*EvidenceBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *Evidence) SetBillingAddress(v EvidenceBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *Evidence) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetCreated

`func (o *Evidence) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Evidence) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Evidence) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Evidence) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetIpAddress

`func (o *Evidence) GetIpAddress() EvidenceIpAddress`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Evidence) GetIpAddressOk() (*EvidenceIpAddress, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Evidence) SetIpAddress(v EvidenceIpAddress)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Evidence) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetId

`func (o *Evidence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Evidence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Evidence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Evidence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequiredCount

`func (o *Evidence) GetRequiredCount() int32`

GetRequiredCount returns the RequiredCount field if non-nil, zero value otherwise.

### GetRequiredCountOk

`func (o *Evidence) GetRequiredCountOk() (*int32, bool)`

GetRequiredCountOk returns a tuple with the RequiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCount

`func (o *Evidence) SetRequiredCount(v int32)`

SetRequiredCount sets RequiredCount field to given value.

### HasRequiredCount

`func (o *Evidence) HasRequiredCount() bool`

HasRequiredCount returns a boolean if a field has been set.

### GetUpdated

`func (o *Evidence) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Evidence) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Evidence) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Evidence) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


