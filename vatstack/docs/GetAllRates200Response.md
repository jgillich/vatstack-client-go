# GetAllRates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | True if this list has another page of items after this one that can be fetched. | 
**RatesCount** | **int32** |  | 
**Rates** | [**[]Rate**](Rate.md) |  | 

## Methods

### NewGetAllRates200Response

`func NewGetAllRates200Response(hasMore bool, ratesCount int32, rates []Rate, ) *GetAllRates200Response`

NewGetAllRates200Response instantiates a new GetAllRates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllRates200ResponseWithDefaults

`func NewGetAllRates200ResponseWithDefaults() *GetAllRates200Response`

NewGetAllRates200ResponseWithDefaults instantiates a new GetAllRates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllRates200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllRates200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllRates200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetRatesCount

`func (o *GetAllRates200Response) GetRatesCount() int32`

GetRatesCount returns the RatesCount field if non-nil, zero value otherwise.

### GetRatesCountOk

`func (o *GetAllRates200Response) GetRatesCountOk() (*int32, bool)`

GetRatesCountOk returns a tuple with the RatesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatesCount

`func (o *GetAllRates200Response) SetRatesCount(v int32)`

SetRatesCount sets RatesCount field to given value.


### GetRates

`func (o *GetAllRates200Response) GetRates() []Rate`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *GetAllRates200Response) GetRatesOk() (*[]Rate, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *GetAllRates200Response) SetRates(v []Rate)`

SetRates sets Rates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


