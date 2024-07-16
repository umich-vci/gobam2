# ResponsePolicyPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the response policy | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**PolicyType** | **string** | The type of response policy configured. | 
**Ttl** | Pointer to **int64** | The TTL value for each type of response policy, in seconds. | [optional] 
**RedirectTarget** | Pointer to **string** | For response policies configured as REDIRECT policies, the fully qualified domain name to which users will be redirected. | [optional] 

## Methods

### NewResponsePolicyPostRequestBody

`func NewResponsePolicyPostRequestBody(name string, policyType string, ) *ResponsePolicyPostRequestBody`

NewResponsePolicyPostRequestBody instantiates a new ResponsePolicyPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePolicyPostRequestBodyWithDefaults

`func NewResponsePolicyPostRequestBodyWithDefaults() *ResponsePolicyPostRequestBody`

NewResponsePolicyPostRequestBodyWithDefaults instantiates a new ResponsePolicyPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsePolicyPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePolicyPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePolicyPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsePolicyPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResponsePolicyPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePolicyPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePolicyPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePolicyPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponsePolicyPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePolicyPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePolicyPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ResponsePolicyPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ResponsePolicyPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ResponsePolicyPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ResponsePolicyPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ResponsePolicyPostRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ResponsePolicyPostRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ResponsePolicyPostRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ResponsePolicyPostRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPolicyType

`func (o *ResponsePolicyPostRequestBody) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *ResponsePolicyPostRequestBody) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *ResponsePolicyPostRequestBody) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.


### GetTtl

`func (o *ResponsePolicyPostRequestBody) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ResponsePolicyPostRequestBody) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ResponsePolicyPostRequestBody) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ResponsePolicyPostRequestBody) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetRedirectTarget

`func (o *ResponsePolicyPostRequestBody) GetRedirectTarget() string`

GetRedirectTarget returns the RedirectTarget field if non-nil, zero value otherwise.

### GetRedirectTargetOk

`func (o *ResponsePolicyPostRequestBody) GetRedirectTargetOk() (*string, bool)`

GetRedirectTargetOk returns a tuple with the RedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTarget

`func (o *ResponsePolicyPostRequestBody) SetRedirectTarget(v string)`

SetRedirectTarget sets RedirectTarget field to given value.

### HasRedirectTarget

`func (o *ResponsePolicyPostRequestBody) HasRedirectTarget() bool`

HasRedirectTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


