# AccessRightPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**DefaultAccessLevel** | **string** | The default access right assigned to the user or group. | 
**UserScope** | Pointer to [**InlinedUserScope**](InlinedUserScope.md) |  | [optional] 
**Resource** | Pointer to [**InlinedResource**](InlinedResource.md) |  | [optional] 
**DeploymentsAllowed** | **bool** | Indicates whether or not the user or group can perform a full deployment of data from the configuration to a managed server. | 
**QuickDeploymentsAllowed** | **bool** | Indicates whether or not the user or group can instantly deploy changed DNS resource records. | 
**WorkflowLevel** | **string** | The workflow level assigned to the user. | 
**AccessOverrides** | [**[]AccessRightOverrideBean**](AccessRightOverrideBean.md) | A list of overridden access rights. | 
**SelectiveDeploymentsAllowed** | **bool** | Indicates whether a user or group can perform a selection deployment of data to a managed server, and dynamically deploy resource records that are added, updated, and deleted within a DNS zone that has the Dynamic Update feature enabled. | 

## Methods

### NewAccessRightPutRequestBody

`func NewAccessRightPutRequestBody(defaultAccessLevel string, deploymentsAllowed bool, quickDeploymentsAllowed bool, workflowLevel string, accessOverrides []AccessRightOverrideBean, selectiveDeploymentsAllowed bool, ) *AccessRightPutRequestBody`

NewAccessRightPutRequestBody instantiates a new AccessRightPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRightPutRequestBodyWithDefaults

`func NewAccessRightPutRequestBodyWithDefaults() *AccessRightPutRequestBody`

NewAccessRightPutRequestBodyWithDefaults instantiates a new AccessRightPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessRightPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessRightPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessRightPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccessRightPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AccessRightPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessRightPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessRightPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessRightPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultAccessLevel

`func (o *AccessRightPutRequestBody) GetDefaultAccessLevel() string`

GetDefaultAccessLevel returns the DefaultAccessLevel field if non-nil, zero value otherwise.

### GetDefaultAccessLevelOk

`func (o *AccessRightPutRequestBody) GetDefaultAccessLevelOk() (*string, bool)`

GetDefaultAccessLevelOk returns a tuple with the DefaultAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessLevel

`func (o *AccessRightPutRequestBody) SetDefaultAccessLevel(v string)`

SetDefaultAccessLevel sets DefaultAccessLevel field to given value.


### GetUserScope

`func (o *AccessRightPutRequestBody) GetUserScope() InlinedUserScope`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *AccessRightPutRequestBody) GetUserScopeOk() (*InlinedUserScope, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *AccessRightPutRequestBody) SetUserScope(v InlinedUserScope)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *AccessRightPutRequestBody) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.

### GetResource

`func (o *AccessRightPutRequestBody) GetResource() InlinedResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AccessRightPutRequestBody) GetResourceOk() (*InlinedResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AccessRightPutRequestBody) SetResource(v InlinedResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AccessRightPutRequestBody) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDeploymentsAllowed

`func (o *AccessRightPutRequestBody) GetDeploymentsAllowed() bool`

GetDeploymentsAllowed returns the DeploymentsAllowed field if non-nil, zero value otherwise.

### GetDeploymentsAllowedOk

`func (o *AccessRightPutRequestBody) GetDeploymentsAllowedOk() (*bool, bool)`

GetDeploymentsAllowedOk returns a tuple with the DeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsAllowed

`func (o *AccessRightPutRequestBody) SetDeploymentsAllowed(v bool)`

SetDeploymentsAllowed sets DeploymentsAllowed field to given value.


### GetQuickDeploymentsAllowed

`func (o *AccessRightPutRequestBody) GetQuickDeploymentsAllowed() bool`

GetQuickDeploymentsAllowed returns the QuickDeploymentsAllowed field if non-nil, zero value otherwise.

### GetQuickDeploymentsAllowedOk

`func (o *AccessRightPutRequestBody) GetQuickDeploymentsAllowedOk() (*bool, bool)`

GetQuickDeploymentsAllowedOk returns a tuple with the QuickDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickDeploymentsAllowed

`func (o *AccessRightPutRequestBody) SetQuickDeploymentsAllowed(v bool)`

SetQuickDeploymentsAllowed sets QuickDeploymentsAllowed field to given value.


### GetWorkflowLevel

`func (o *AccessRightPutRequestBody) GetWorkflowLevel() string`

GetWorkflowLevel returns the WorkflowLevel field if non-nil, zero value otherwise.

### GetWorkflowLevelOk

`func (o *AccessRightPutRequestBody) GetWorkflowLevelOk() (*string, bool)`

GetWorkflowLevelOk returns a tuple with the WorkflowLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowLevel

`func (o *AccessRightPutRequestBody) SetWorkflowLevel(v string)`

SetWorkflowLevel sets WorkflowLevel field to given value.


### GetAccessOverrides

`func (o *AccessRightPutRequestBody) GetAccessOverrides() []AccessRightOverrideBean`

GetAccessOverrides returns the AccessOverrides field if non-nil, zero value otherwise.

### GetAccessOverridesOk

`func (o *AccessRightPutRequestBody) GetAccessOverridesOk() (*[]AccessRightOverrideBean, bool)`

GetAccessOverridesOk returns a tuple with the AccessOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOverrides

`func (o *AccessRightPutRequestBody) SetAccessOverrides(v []AccessRightOverrideBean)`

SetAccessOverrides sets AccessOverrides field to given value.


### GetSelectiveDeploymentsAllowed

`func (o *AccessRightPutRequestBody) GetSelectiveDeploymentsAllowed() bool`

GetSelectiveDeploymentsAllowed returns the SelectiveDeploymentsAllowed field if non-nil, zero value otherwise.

### GetSelectiveDeploymentsAllowedOk

`func (o *AccessRightPutRequestBody) GetSelectiveDeploymentsAllowedOk() (*bool, bool)`

GetSelectiveDeploymentsAllowedOk returns a tuple with the SelectiveDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveDeploymentsAllowed

`func (o *AccessRightPutRequestBody) SetSelectiveDeploymentsAllowed(v bool)`

SetSelectiveDeploymentsAllowed sets SelectiveDeploymentsAllowed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


