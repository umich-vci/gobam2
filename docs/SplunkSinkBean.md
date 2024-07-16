# SplunkSinkBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The URI of the Splunk HEC host. | [optional] 
**Token** | Pointer to **string** | The Splunk HEC token. | [optional] 

## Methods

### NewSplunkSinkBean

`func NewSplunkSinkBean() *SplunkSinkBean`

NewSplunkSinkBean instantiates a new SplunkSinkBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplunkSinkBeanWithDefaults

`func NewSplunkSinkBeanWithDefaults() *SplunkSinkBean`

NewSplunkSinkBeanWithDefaults instantiates a new SplunkSinkBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SplunkSinkBean) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SplunkSinkBean) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SplunkSinkBean) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SplunkSinkBean) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *SplunkSinkBean) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SplunkSinkBean) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SplunkSinkBean) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SplunkSinkBean) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


