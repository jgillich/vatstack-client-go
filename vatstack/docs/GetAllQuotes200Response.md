# GetAllQuotes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | True if this list has another page of items after this one that can be fetched. | 
**QuotesCount** | **int32** |  | 
**Quotes** | [**[]Quote**](Quote.md) |  | 

## Methods

### NewGetAllQuotes200Response

`func NewGetAllQuotes200Response(hasMore bool, quotesCount int32, quotes []Quote, ) *GetAllQuotes200Response`

NewGetAllQuotes200Response instantiates a new GetAllQuotes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllQuotes200ResponseWithDefaults

`func NewGetAllQuotes200ResponseWithDefaults() *GetAllQuotes200Response`

NewGetAllQuotes200ResponseWithDefaults instantiates a new GetAllQuotes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAllQuotes200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAllQuotes200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAllQuotes200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetQuotesCount

`func (o *GetAllQuotes200Response) GetQuotesCount() int32`

GetQuotesCount returns the QuotesCount field if non-nil, zero value otherwise.

### GetQuotesCountOk

`func (o *GetAllQuotes200Response) GetQuotesCountOk() (*int32, bool)`

GetQuotesCountOk returns a tuple with the QuotesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotesCount

`func (o *GetAllQuotes200Response) SetQuotesCount(v int32)`

SetQuotesCount sets QuotesCount field to given value.


### GetQuotes

`func (o *GetAllQuotes200Response) GetQuotes() []Quote`

GetQuotes returns the Quotes field if non-nil, zero value otherwise.

### GetQuotesOk

`func (o *GetAllQuotes200Response) GetQuotesOk() (*[]Quote, bool)`

GetQuotesOk returns a tuple with the Quotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotes

`func (o *GetAllQuotes200Response) SetQuotes(v []Quote)`

SetQuotes sets Quotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


