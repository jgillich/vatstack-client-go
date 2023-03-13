# GetAllValidations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | True if this list has another page of items after this one that can be fetched. | 
**ValidationsCount** | **int32** |  | 
**Validations** | [**[]Validation**](Validation.md) |  | 

## Methods

### NewGetAllValidations200Response

`func NewGetAllValidations200Response(hasMore bool, validationsCount int32, validations []Validation, ) *GetAllValidations200Response`

NewGetAllValidations200Response instantiates a new GetAllValidations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllValidations200ResponseWithDefaults

`func NewGetAllValidations200ResponseWithDefaults() *GetAllValidations200Response`

NewGetAllValidations200ResponseWithDefaults instantiates a new GetAllValidations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllValidations200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllValidations200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllValidations200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetValidationsCount

`func (o *GetAllValidations200Response) GetValidationsCount() int32`

GetValidationsCount returns the ValidationsCount field if non-nil, zero value otherwise.

### GetValidationsCountOk

`func (o *GetAllValidations200Response) GetValidationsCountOk() (*int32, bool)`

GetValidationsCountOk returns a tuple with the ValidationsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationsCount

`func (o *GetAllValidations200Response) SetValidationsCount(v int32)`

SetValidationsCount sets ValidationsCount field to given value.


### GetValidations

`func (o *GetAllValidations200Response) GetValidations() []Validation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *GetAllValidations200Response) GetValidationsOk() (*[]Validation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *GetAllValidations200Response) SetValidations(v []Validation)`

SetValidations sets Validations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


