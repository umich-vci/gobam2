# AccessRight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**DefaultAccessLevel** | Pointer to **string** | The default access right assigned to the user or group. | [optional] 
**UserScope** | Pointer to [**InlinedUserScope**](InlinedUserScope.md) |  | [optional] 
**Resource** | Pointer to [**InlinedResource**](InlinedResource.md) |  | [optional] 
**DeploymentsAllowed** | Pointer to **bool** | Indicates whether or not the user or group can perform a full deployment of data from the configuration to a managed server. | [optional] 
**QuickDeploymentsAllowed** | Pointer to **bool** | Indicates whether or not the user or group can instantly deploy changed DNS resource records. | [optional] 
**WorkflowLevel** | Pointer to **string** | The workflow level assigned to the user. | [optional] 
**AccessOverrides** | Pointer to [**[]AccessRightOverrideBean**](AccessRightOverrideBean.md) | A list of overridden access rights. | [optional] 
**SelectiveDeploymentsAllowed** | Pointer to **bool** | Indicates whether a user or group can perform a selection deployment of data to a managed server, and dynamically deploy resource records that are added, updated, and deleted within a DNS zone that has the Dynamic Update feature enabled. | [optional] 

## Methods

### NewAccessRight

`func NewAccessRight() *AccessRight`

NewAccessRight instantiates a new AccessRight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRightWithDefaults

`func NewAccessRightWithDefaults() *AccessRight`

NewAccessRightWithDefaults instantiates a new AccessRight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessRight) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessRight) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessRight) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccessRight) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AccessRight) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessRight) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessRight) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccessRight) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultAccessLevel

`func (o *AccessRight) GetDefaultAccessLevel() string`

GetDefaultAccessLevel returns the DefaultAccessLevel field if non-nil, zero value otherwise.

### GetDefaultAccessLevelOk

`func (o *AccessRight) GetDefaultAccessLevelOk() (*string, bool)`

GetDefaultAccessLevelOk returns a tuple with the DefaultAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessLevel

`func (o *AccessRight) SetDefaultAccessLevel(v string)`

SetDefaultAccessLevel sets DefaultAccessLevel field to given value.

### HasDefaultAccessLevel

`func (o *AccessRight) HasDefaultAccessLevel() bool`

HasDefaultAccessLevel returns a boolean if a field has been set.

### GetUserScope

`func (o *AccessRight) GetUserScope() InlinedUserScope`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *AccessRight) GetUserScopeOk() (*InlinedUserScope, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *AccessRight) SetUserScope(v InlinedUserScope)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *AccessRight) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.

### GetResource

`func (o *AccessRight) GetResource() InlinedResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *AccessRight) GetResourceOk() (*InlinedResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *AccessRight) SetResource(v InlinedResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *AccessRight) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDeploymentsAllowed

`func (o *AccessRight) GetDeploymentsAllowed() bool`

GetDeploymentsAllowed returns the DeploymentsAllowed field if non-nil, zero value otherwise.

### GetDeploymentsAllowedOk

`func (o *AccessRight) GetDeploymentsAllowedOk() (*bool, bool)`

GetDeploymentsAllowedOk returns a tuple with the DeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsAllowed

`func (o *AccessRight) SetDeploymentsAllowed(v bool)`

SetDeploymentsAllowed sets DeploymentsAllowed field to given value.

### HasDeploymentsAllowed

`func (o *AccessRight) HasDeploymentsAllowed() bool`

HasDeploymentsAllowed returns a boolean if a field has been set.

### GetQuickDeploymentsAllowed

`func (o *AccessRight) GetQuickDeploymentsAllowed() bool`

GetQuickDeploymentsAllowed returns the QuickDeploymentsAllowed field if non-nil, zero value otherwise.

### GetQuickDeploymentsAllowedOk

`func (o *AccessRight) GetQuickDeploymentsAllowedOk() (*bool, bool)`

GetQuickDeploymentsAllowedOk returns a tuple with the QuickDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickDeploymentsAllowed

`func (o *AccessRight) SetQuickDeploymentsAllowed(v bool)`

SetQuickDeploymentsAllowed sets QuickDeploymentsAllowed field to given value.

### HasQuickDeploymentsAllowed

`func (o *AccessRight) HasQuickDeploymentsAllowed() bool`

HasQuickDeploymentsAllowed returns a boolean if a field has been set.

### GetWorkflowLevel

`func (o *AccessRight) GetWorkflowLevel() string`

GetWorkflowLevel returns the WorkflowLevel field if non-nil, zero value otherwise.

### GetWorkflowLevelOk

`func (o *AccessRight) GetWorkflowLevelOk() (*string, bool)`

GetWorkflowLevelOk returns a tuple with the WorkflowLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowLevel

`func (o *AccessRight) SetWorkflowLevel(v string)`

SetWorkflowLevel sets WorkflowLevel field to given value.

### HasWorkflowLevel

`func (o *AccessRight) HasWorkflowLevel() bool`

HasWorkflowLevel returns a boolean if a field has been set.

### GetAccessOverrides

`func (o *AccessRight) GetAccessOverrides() []AccessRightOverrideBean`

GetAccessOverrides returns the AccessOverrides field if non-nil, zero value otherwise.

### GetAccessOverridesOk

`func (o *AccessRight) GetAccessOverridesOk() (*[]AccessRightOverrideBean, bool)`

GetAccessOverridesOk returns a tuple with the AccessOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOverrides

`func (o *AccessRight) SetAccessOverrides(v []AccessRightOverrideBean)`

SetAccessOverrides sets AccessOverrides field to given value.

### HasAccessOverrides

`func (o *AccessRight) HasAccessOverrides() bool`

HasAccessOverrides returns a boolean if a field has been set.

### GetSelectiveDeploymentsAllowed

`func (o *AccessRight) GetSelectiveDeploymentsAllowed() bool`

GetSelectiveDeploymentsAllowed returns the SelectiveDeploymentsAllowed field if non-nil, zero value otherwise.

### GetSelectiveDeploymentsAllowedOk

`func (o *AccessRight) GetSelectiveDeploymentsAllowedOk() (*bool, bool)`

GetSelectiveDeploymentsAllowedOk returns a tuple with the SelectiveDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveDeploymentsAllowed

`func (o *AccessRight) SetSelectiveDeploymentsAllowed(v bool)`

SetSelectiveDeploymentsAllowed sets SelectiveDeploymentsAllowed field to given value.

### HasSelectiveDeploymentsAllowed

`func (o *AccessRight) HasSelectiveDeploymentsAllowed() bool`

HasSelectiveDeploymentsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


