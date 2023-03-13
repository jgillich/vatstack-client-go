# EvidenceIpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The same IP address coming from the &#x60;ip_address.label&#x60; body parameter which will be geolocated automatically. | [optional] 
**Provider** | Pointer to **string** | Provider used to geolocate &#x60;ip_address.label&#x60;. We use MaxMind® GeoIP2 geolocation technology by default but have a number of fallback providers. | [optional] 
**City** | Pointer to **string** | City of the geolocated IP address. | [optional] 
**CountryCode** | Pointer to **string** | 2-letter ISO country code of the geolocated IP address. | [optional] 
**State** | Pointer to **string** | State of the geolocated IP address. | [optional] 

## Methods

### NewEvidenceIpAddress

`func NewEvidenceIpAddress() *EvidenceIpAddress`

NewEvidenceIpAddress instantiates a new EvidenceIpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceIpAddressWithDefaults

`func NewEvidenceIpAddressWithDefaults() *EvidenceIpAddress`

NewEvidenceIpAddressWithDefaults instantiates a new EvidenceIpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *EvidenceIpAddress) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EvidenceIpAddress) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EvidenceIpAddress) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EvidenceIpAddress) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProvider

`func (o *EvidenceIpAddress) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EvidenceIpAddress) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EvidenceIpAddress) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EvidenceIpAddress) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCity

`func (o *EvidenceIpAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EvidenceIpAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EvidenceIpAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *EvidenceIpAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *EvidenceIpAddress) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *EvidenceIpAddress) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *EvidenceIpAddress) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *EvidenceIpAddress) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetState

`func (o *EvidenceIpAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EvidenceIpAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EvidenceIpAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EvidenceIpAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


