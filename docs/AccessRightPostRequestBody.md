# AccessRightPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**DefaultAccessLevel** | **string** | The default access right assigned to the user or group. | 
**UserScope** | [**InlinedUserScope**](InlinedUserScope.md) |  | 
**Resource** | Pointer to [**InlinedResource**](InlinedResource.md) |  | [optional] 
**DeploymentsAllowed** | Pointer to **bool** | Indicates whether or not the user or group can perform a full deployment of data from the configuration to a managed server. | [optional] 
**QuickDeploymentsAllowed** | Pointer to **bool** | Indicates whether or not the user or group can instantly deploy changed DNS resource records. | [optional] 
**WorkflowLevel** | Pointer to **string** | The workflow level assigned to the user. | [optional] 
**AccessOverrides** | Pointer to [**[]AccessRightOverrideBean**](AccessRightOverrideBean.md) | A list of overridden access rights. | [optional] 
**SelectiveDeploymentsAllowed** | Pointer to **bool** | Indicates whether a user or group can perform a selection deployment of data to a managed server, and dynamically deploy resource records that are added, updated, and deleted within a DNS zone that has the Dynamic Update feature enabled. | [optional] 

## Methods

### NewAccessRightPostRequestBody

`func NewAccessRightPostRequestBody(defaultAccessLevel string, userScope InlinedUserScope, ) *AccessRightPostRequestBody`

NewAccessRightPostRequestBody instantiates a new AccessRightPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRightPostRequestBodyWithDefaults

`func NewAccessRightPostRequestBodyWithDefaults() *AccessRightPostRequestBody`

NewAccessRightPostRequestBodyWithDefaults instantiates a new AccessRightPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessRightPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessRightPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessRightPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccessRightPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AccessRightPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessRightPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessRightPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessRightPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultAccessLevel

`func (o *AccessRightPostRequestBody) GetDefaultAccessLevel() string`

GetDefaultAccessLevel returns the DefaultAccessLevel field if non-nil, zero value otherwise.

### GetDefaultAccessLevelOk

`func (o *AccessRightPostRequestBody) GetDefaultAccessLevelOk() (*string, bool)`

GetDefaultAccessLevelOk returns a tuple with the DefaultAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessLevel

`func (o *AccessRightPostRequestBody) SetDefaultAccessLevel(v string)`

SetDefaultAccessLevel sets DefaultAccessLevel field to given value.


### GetUserScope

`func (o *AccessRightPostRequestBody) GetUserScope() InlinedUserScope`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *AccessRightPostRequestBody) GetUserScopeOk() (*InlinedUserScope, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *AccessRightPostRequestBody) SetUserScope(v InlinedUserScope)`

SetUserScope sets UserScope field to given value.


### GetResource

`func (o *AccessRightPostRequestBody) GetResource() InlinedResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AccessRightPostRequestBody) GetResourceOk() (*InlinedResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AccessRightPostRequestBody) SetResource(v InlinedResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AccessRightPostRequestBody) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDeploymentsAllowed

`func (o *AccessRightPostRequestBody) GetDeploymentsAllowed() bool`

GetDeploymentsAllowed returns the DeploymentsAllowed field if non-nil, zero value otherwise.

### GetDeploymentsAllowedOk

`func (o *AccessRightPostRequestBody) GetDeploymentsAllowedOk() (*bool, bool)`

GetDeploymentsAllowedOk returns a tuple with the DeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsAllowed

`func (o *AccessRightPostRequestBody) SetDeploymentsAllowed(v bool)`

SetDeploymentsAllowed sets DeploymentsAllowed field to given value.

### HasDeploymentsAllowed

`func (o *AccessRightPostRequestBody) HasDeploymentsAllowed() bool`

HasDeploymentsAllowed returns a boolean if a field has been set.

### GetQuickDeploymentsAllowed

`func (o *AccessRightPostRequestBody) GetQuickDeploymentsAllowed() bool`

GetQuickDeploymentsAllowed returns the QuickDeploymentsAllowed field if non-nil, zero value otherwise.

### GetQuickDeploymentsAllowedOk

`func (o *AccessRightPostRequestBody) GetQuickDeploymentsAllowedOk() (*bool, bool)`

GetQuickDeploymentsAllowedOk returns a tuple with the QuickDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickDeploymentsAllowed

`func (o *AccessRightPostRequestBody) SetQuickDeploymentsAllowed(v bool)`

SetQuickDeploymentsAllowed sets QuickDeploymentsAllowed field to given value.

### HasQuickDeploymentsAllowed

`func (o *AccessRightPostRequestBody) HasQuickDeploymentsAllowed() bool`

HasQuickDeploymentsAllowed returns a boolean if a field has been set.

### GetWorkflowLevel

`func (o *AccessRightPostRequestBody) GetWorkflowLevel() string`

GetWorkflowLevel returns the WorkflowLevel field if non-nil, zero value otherwise.

### GetWorkflowLevelOk

`func (o *AccessRightPostRequestBody) GetWorkflowLevelOk() (*string, bool)`

GetWorkflowLevelOk returns a tuple with the WorkflowLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowLevel

`func (o *AccessRightPostRequestBody) SetWorkflowLevel(v string)`

SetWorkflowLevel sets WorkflowLevel field to given value.

### HasWorkflowLevel

`func (o *AccessRightPostRequestBody) HasWorkflowLevel() bool`

HasWorkflowLevel returns a boolean if a field has been set.

### GetAccessOverrides

`func (o *AccessRightPostRequestBody) GetAccessOverrides() []AccessRightOverrideBean`

GetAccessOverrides returns the AccessOverrides field if non-nil, zero value otherwise.

### GetAccessOverridesOk

`func (o *AccessRightPostRequestBody) GetAccessOverridesOk() (*[]AccessRightOverrideBean, bool)`

GetAccessOverridesOk returns a tuple with the AccessOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOverrides

`func (o *AccessRightPostRequestBody) SetAccessOverrides(v []AccessRightOverrideBean)`

SetAccessOverrides sets AccessOverrides field to given value.

### HasAccessOverrides

`func (o *AccessRightPostRequestBody) HasAccessOverrides() bool`

HasAccessOverrides returns a boolean if a field has been set.

### GetSelectiveDeploymentsAllowed

`func (o *AccessRightPostRequestBody) GetSelectiveDeploymentsAllowed() bool`

GetSelectiveDeploymentsAllowed returns the SelectiveDeploymentsAllowed field if non-nil, zero value otherwise.

### GetSelectiveDeploymentsAllowedOk

`func (o *AccessRightPostRequestBody) GetSelectiveDeploymentsAllowedOk() (*bool, bool)`

GetSelectiveDeploymentsAllowedOk returns a tuple with the SelectiveDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveDeploymentsAllowed

`func (o *AccessRightPostRequestBody) SetSelectiveDeploymentsAllowed(v bool)`

SetSelectiveDeploymentsAllowed sets SelectiveDeploymentsAllowed field to given value.

### HasSelectiveDeploymentsAllowed

`func (o *AccessRightPostRequestBody) HasSelectiveDeploymentsAllowed() bool`

HasSelectiveDeploymentsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


