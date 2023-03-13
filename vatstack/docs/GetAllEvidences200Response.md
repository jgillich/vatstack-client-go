# GetAllEvidences200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | True if this list has another page of items after this one that can be fetched. | 
**EvidencesCount** | **int32** |  | 
**Evidences** | [**[]Evidence**](Evidence.md) |  | 

## Methods

### NewGetAllEvidences200Response

`func NewGetAllEvidences200Response(hasMore bool, evidencesCount int32, evidences []Evidence, ) *GetAllEvidences200Response`

NewGetAllEvidences200Response instantiates a new GetAllEvidences200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllEvidences200ResponseWithDefaults

`func NewGetAllEvidences200ResponseWithDefaults() *GetAllEvidences200Response`

NewGetAllEvidences200ResponseWithDefaults instantiates a new GetAllEvidences200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllEvidences200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllEvidences200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllEvidences200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetEvidencesCount

`func (o *GetAllEvidences200Response) GetEvidencesCount() int32`

GetEvidencesCount returns the EvidencesCount field if non-nil, zero value otherwise.

### GetEvidencesCountOk

`func (o *GetAllEvidences200Response) GetEvidencesCountOk() (*int32, bool)`

GetEvidencesCountOk returns a tuple with the EvidencesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidencesCount

`func (o *GetAllEvidences200Response) SetEvidencesCount(v int32)`

SetEvidencesCount sets EvidencesCount field to given value.


### GetEvidences

`func (o *GetAllEvidences200Response) GetEvidences() []Evidence`

GetEvidences returns the Evidences field if non-nil, zero value otherwise.

### GetEvidencesOk

`func (o *GetAllEvidences200Response) GetEvidencesOk() (*[]Evidence, bool)`

GetEvidencesOk returns a tuple with the Evidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidences

`func (o *GetAllEvidences200Response) SetEvidences(v []Evidence)`

SetEvidences sets Evidences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


