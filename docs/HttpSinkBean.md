# HttpSinkBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**HealthCheckUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpSinkBean

`func NewHttpSinkBean() *HttpSinkBean`

NewHttpSinkBean instantiates a new HttpSinkBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpSinkBeanWithDefaults

`func NewHttpSinkBeanWithDefaults() *HttpSinkBean`

NewHttpSinkBeanWithDefaults instantiates a new HttpSinkBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HttpSinkBean) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpSinkBean) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpSinkBean) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpSinkBean) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *HttpSinkBean) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *HttpSinkBean) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *HttpSinkBean) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *HttpSinkBean) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetHealthCheckUrl

`func (o *HttpSinkBean) GetHealthCheckUrl() string`

GetHealthCheckUrl returns the HealthCheckUrl field if non-nil, zero value otherwise.

### GetHealthCheckUrlOk

`func (o *HttpSinkBean) GetHealthCheckUrlOk() (*string, bool)`

GetHealthCheckUrlOk returns a tuple with the HealthCheckUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUrl

`func (o *HttpSinkBean) SetHealthCheckUrl(v string)`

SetHealthCheckUrl sets HealthCheckUrl field to given value.

### HasHealthCheckUrl

`func (o *HttpSinkBean) HasHealthCheckUrl() bool`

HasHealthCheckUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


