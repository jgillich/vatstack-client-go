# Batch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | ISO date at which the object was created. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**Name** | **string** | Descriptive name of the object. | 
**Queries** | Pointer to **[]string** | Array of all queries in the order of your input. | [optional] 
**QueriesIgnored** | Pointer to **[]string** | Array of all queries which are of invalid format. Ignored queries are automatically identified upon your input. Its items will not be validated. | [optional] 
**Scheduled** | Pointer to **time.Time** | ISO date at which the object status was changed to &#x60;scheduled&#x60;. Defaults to &#x60;null&#x60; while the status is &#x60;pending&#x60;. | [optional] 
**SucceededCount** | Pointer to **int32** | The number of &#x60;validations&#x60; where &#x60;valid&#x60; is &#x60;true&#x60; or &#x60;false&#x60;. It’s an indicator for how far the batch has progressed. | [optional] 
**Updated** | Pointer to **time.Time** | ISO date at which the object was updated. | [optional] 
**Validations** | Pointer to [**[]Validation**](Validation.md) | Array of the first 20 validation objects in the order of &#x60;queries&#x60;. If an &#x60;id&#x60; is shown in the object, it was created for the batch and you can retrieve it individually. See [validation object](https://vatstack.com/docs/validations) for reference. | [optional] 

## Methods

### NewBatch

`func NewBatch(name string, ) *Batch`

NewBatch instantiates a new Batch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchWithDefaults

`func NewBatchWithDefaults() *Batch`

NewBatchWithDefaults instantiates a new Batch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Batch) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Batch) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Batch) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Batch) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Batch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Batch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Batch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Batch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Batch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Batch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Batch) SetName(v string)`

SetName sets Name field to given value.


### GetQueries

`func (o *Batch) GetQueries() []string`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *Batch) GetQueriesOk() (*[]string, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *Batch) SetQueries(v []string)`

SetQueries sets Queries field to given value.

### HasQueries

`func (o *Batch) HasQueries() bool`

HasQueries returns a boolean if a field has been set.

### GetQueriesIgnored

`func (o *Batch) GetQueriesIgnored() []string`

GetQueriesIgnored returns the QueriesIgnored field if non-nil, zero value otherwise.

### GetQueriesIgnoredOk

`func (o *Batch) GetQueriesIgnoredOk() (*[]string, bool)`

GetQueriesIgnoredOk returns a tuple with the QueriesIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriesIgnored

`func (o *Batch) SetQueriesIgnored(v []string)`

SetQueriesIgnored sets QueriesIgnored field to given value.

### HasQueriesIgnored

`func (o *Batch) HasQueriesIgnored() bool`

HasQueriesIgnored returns a boolean if a field has been set.

### GetScheduled

`func (o *Batch) GetScheduled() time.Time`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *Batch) GetScheduledOk() (*time.Time, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *Batch) SetScheduled(v time.Time)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *Batch) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetSucceededCount

`func (o *Batch) GetSucceededCount() int32`

GetSucceededCount returns the SucceededCount field if non-nil, zero value otherwise.

### GetSucceededCountOk

`func (o *Batch) GetSucceededCountOk() (*int32, bool)`

GetSucceededCountOk returns a tuple with the SucceededCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceededCount

`func (o *Batch) SetSucceededCount(v int32)`

SetSucceededCount sets SucceededCount field to given value.

### HasSucceededCount

`func (o *Batch) HasSucceededCount() bool`

HasSucceededCount returns a boolean if a field has been set.

### GetUpdated

`func (o *Batch) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Batch) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Batch) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Batch) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValidations

`func (o *Batch) GetValidations() []Validation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *Batch) GetValidationsOk() (*[]Validation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *Batch) SetValidations(v []Validation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *Batch) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


