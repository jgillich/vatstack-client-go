# GetAllSupplies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | True if this list has another page of items after this one that can be fetched. | 
**SuppliesCount** | **int32** |  | 
**Supplies** | [**[]Supply**](Supply.md) |  | 

## Methods

### NewGetAllSupplies200Response

`func NewGetAllSupplies200Response(hasMore bool, suppliesCount int32, supplies []Supply, ) *GetAllSupplies200Response`

NewGetAllSupplies200Response instantiates a new GetAllSupplies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllSupplies200ResponseWithDefaults

`func NewGetAllSupplies200ResponseWithDefaults() *GetAllSupplies200Response`

NewGetAllSupplies200ResponseWithDefaults instantiates a new GetAllSupplies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllSupplies200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllSupplies200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllSupplies200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetSuppliesCount

`func (o *GetAllSupplies200Response) GetSuppliesCount() int32`

GetSuppliesCount returns the SuppliesCount field if non-nil, zero value otherwise.

### GetSuppliesCountOk

`func (o *GetAllSupplies200Response) GetSuppliesCountOk() (*int32, bool)`

GetSuppliesCountOk returns a tuple with the SuppliesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppliesCount

`func (o *GetAllSupplies200Response) SetSuppliesCount(v int32)`

SetSuppliesCount sets SuppliesCount field to given value.


### GetSupplies

`func (o *GetAllSupplies200Response) GetSupplies() []Supply`

GetSupplies returns the Supplies field if non-nil, zero value otherwise.

### GetSuppliesOk

`func (o *GetAllSupplies200Response) GetSuppliesOk() (*[]Supply, bool)`

GetSuppliesOk returns a tuple with the Supplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplies

`func (o *GetAllSupplies200Response) SetSupplies(v []Supply)`

SetSupplies sets Supplies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


