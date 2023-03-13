# HitValidations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int32** | Included validations in the plan you’re subscribed to. | [optional] 
**Used** | Pointer to **int32** | Validations created during the running month. Additional requests beyond &#x60;capacity&#x60; will be charged according to the plan you’re subscribed to. | [optional] 

## Methods

### NewHitValidations

`func NewHitValidations() *HitValidations`

NewHitValidations instantiates a new HitValidations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHitValidationsWithDefaults

`func NewHitValidationsWithDefaults() *HitValidations`

NewHitValidationsWithDefaults instantiates a new HitValidations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *HitValidations) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *HitValidations) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *HitValidations) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *HitValidations) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetUsed

`func (o *HitValidations) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *HitValidations) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *HitValidations) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *HitValidations) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


