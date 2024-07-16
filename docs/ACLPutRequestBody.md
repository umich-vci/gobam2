# ACLPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Predefined** | Pointer to **bool** | Indicates whether the access control list is predefined by the system (true) or user-created (false). | [optional] [readonly] 
**MatchElements** | **[]string** |  | 

## Methods

### NewACLPutRequestBody

`func NewACLPutRequestBody(name string, matchElements []string, ) *ACLPutRequestBody`

NewACLPutRequestBody instantiates a new ACLPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLPutRequestBodyWithDefaults

`func NewACLPutRequestBodyWithDefaults() *ACLPutRequestBody`

NewACLPutRequestBodyWithDefaults instantiates a new ACLPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ACLPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ACLPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ACLPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ACLPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ACLPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ACLPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ACLPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ACLPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ACLPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ACLPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ACLPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ACLPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ACLPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ACLPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ACLPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ACLPutRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ACLPutRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ACLPutRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ACLPutRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPredefined

`func (o *ACLPutRequestBody) GetPredefined() bool`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *ACLPutRequestBody) GetPredefinedOk() (*bool, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *ACLPutRequestBody) SetPredefined(v bool)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *ACLPutRequestBody) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.

### GetMatchElements

`func (o *ACLPutRequestBody) GetMatchElements() []string`

GetMatchElements returns the MatchElements field if non-nil, zero value otherwise.

### GetMatchElementsOk

`func (o *ACLPutRequestBody) GetMatchElementsOk() (*[]string, bool)`

GetMatchElementsOk returns a tuple with the MatchElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchElements

`func (o *ACLPutRequestBody) SetMatchElements(v []string)`

SetMatchElements sets MatchElements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


