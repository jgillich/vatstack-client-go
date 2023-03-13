# GetAllBatches200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | True if this list has another page of items after this one that can be fetched. | 
**BatchesCount** | **int32** |  | 
**Batches** | [**[]Batch**](Batch.md) |  | 

## Methods

### NewGetAllBatches200Response

`func NewGetAllBatches200Response(hasMore bool, batchesCount int32, batches []Batch, ) *GetAllBatches200Response`

NewGetAllBatches200Response instantiates a new GetAllBatches200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllBatches200ResponseWithDefaults

`func NewGetAllBatches200ResponseWithDefaults() *GetAllBatches200Response`

NewGetAllBatches200ResponseWithDefaults instantiates a new GetAllBatches200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllBatches200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllBatches200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllBatches200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetBatchesCount

`func (o *GetAllBatches200Response) GetBatchesCount() int32`

GetBatchesCount returns the BatchesCount field if non-nil, zero value otherwise.

### GetBatchesCountOk

`func (o *GetAllBatches200Response) GetBatchesCountOk() (*int32, bool)`

GetBatchesCountOk returns a tuple with the BatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchesCount

`func (o *GetAllBatches200Response) SetBatchesCount(v int32)`

SetBatchesCount sets BatchesCount field to given value.


### GetBatches

`func (o *GetAllBatches200Response) GetBatches() []Batch`

GetBatches returns the Batches field if non-nil, zero value otherwise.

### GetBatchesOk

`func (o *GetAllBatches200Response) GetBatchesOk() (*[]Batch, bool)`

GetBatchesOk returns a tuple with the Batches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatches

`func (o *GetAllBatches200Response) SetBatches(v []Batch)`

SetBatches sets Batches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


