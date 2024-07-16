# GetAccessRights200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetAccessRights200ResponseDataInner

`func NewGetAccessRights200ResponseDataInner() *GetAccessRights200ResponseDataInner`

NewGetAccessRights200ResponseDataInner instantiates a new GetAccessRights200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessRights200ResponseDataInnerWithDefaults

`func NewGetAccessRights200ResponseDataInnerWithDefaults() *GetAccessRights200ResponseDataInner`

NewGetAccessRights200ResponseDataInnerWithDefaults instantiates a new GetAccessRights200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccessRights200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccessRights200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccessRights200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAccessRights200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetAccessRights200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAccessRights200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAccessRights200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAccessRights200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDefaultAccessLevel

`func (o *GetAccessRights200ResponseDataInner) GetDefaultAccessLevel() string`

GetDefaultAccessLevel returns the DefaultAccessLevel field if non-nil, zero value otherwise.

### GetDefaultAccessLevelOk

`func (o *GetAccessRights200ResponseDataInner) GetDefaultAccessLevelOk() (*string, bool)`

GetDefaultAccessLevelOk returns a tuple with the DefaultAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAccessLevel

`func (o *GetAccessRights200ResponseDataInner) SetDefaultAccessLevel(v string)`

SetDefaultAccessLevel sets DefaultAccessLevel field to given value.

### HasDefaultAccessLevel

`func (o *GetAccessRights200ResponseDataInner) HasDefaultAccessLevel() bool`

HasDefaultAccessLevel returns a boolean if a field has been set.

### GetUserScope

`func (o *GetAccessRights200ResponseDataInner) GetUserScope() InlinedUserScope`

GetUserScope returns the UserScope field if non-nil, zero value otherwise.

### GetUserScopeOk

`func (o *GetAccessRights200ResponseDataInner) GetUserScopeOk() (*InlinedUserScope, bool)`

GetUserScopeOk returns a tuple with the UserScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScope

`func (o *GetAccessRights200ResponseDataInner) SetUserScope(v InlinedUserScope)`

SetUserScope sets UserScope field to given value.

### HasUserScope

`func (o *GetAccessRights200ResponseDataInner) HasUserScope() bool`

HasUserScope returns a boolean if a field has been set.

### GetResource

`func (o *GetAccessRights200ResponseDataInner) GetResource() InlinedResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetAccessRights200ResponseDataInner) GetResourceOk() (*InlinedResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetAccessRights200ResponseDataInner) SetResource(v InlinedResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetAccessRights200ResponseDataInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) GetDeploymentsAllowed() bool`

GetDeploymentsAllowed returns the DeploymentsAllowed field if non-nil, zero value otherwise.

### GetDeploymentsAllowedOk

`func (o *GetAccessRights200ResponseDataInner) GetDeploymentsAllowedOk() (*bool, bool)`

GetDeploymentsAllowedOk returns a tuple with the DeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) SetDeploymentsAllowed(v bool)`

SetDeploymentsAllowed sets DeploymentsAllowed field to given value.

### HasDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) HasDeploymentsAllowed() bool`

HasDeploymentsAllowed returns a boolean if a field has been set.

### GetQuickDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) GetQuickDeploymentsAllowed() bool`

GetQuickDeploymentsAllowed returns the QuickDeploymentsAllowed field if non-nil, zero value otherwise.

### GetQuickDeploymentsAllowedOk

`func (o *GetAccessRights200ResponseDataInner) GetQuickDeploymentsAllowedOk() (*bool, bool)`

GetQuickDeploymentsAllowedOk returns a tuple with the QuickDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) SetQuickDeploymentsAllowed(v bool)`

SetQuickDeploymentsAllowed sets QuickDeploymentsAllowed field to given value.

### HasQuickDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) HasQuickDeploymentsAllowed() bool`

HasQuickDeploymentsAllowed returns a boolean if a field has been set.

### GetWorkflowLevel

`func (o *GetAccessRights200ResponseDataInner) GetWorkflowLevel() string`

GetWorkflowLevel returns the WorkflowLevel field if non-nil, zero value otherwise.

### GetWorkflowLevelOk

`func (o *GetAccessRights200ResponseDataInner) GetWorkflowLevelOk() (*string, bool)`

GetWorkflowLevelOk returns a tuple with the WorkflowLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowLevel

`func (o *GetAccessRights200ResponseDataInner) SetWorkflowLevel(v string)`

SetWorkflowLevel sets WorkflowLevel field to given value.

### HasWorkflowLevel

`func (o *GetAccessRights200ResponseDataInner) HasWorkflowLevel() bool`

HasWorkflowLevel returns a boolean if a field has been set.

### GetAccessOverrides

`func (o *GetAccessRights200ResponseDataInner) GetAccessOverrides() []AccessRightOverrideBean`

GetAccessOverrides returns the AccessOverrides field if non-nil, zero value otherwise.

### GetAccessOverridesOk

`func (o *GetAccessRights200ResponseDataInner) GetAccessOverridesOk() (*[]AccessRightOverrideBean, bool)`

GetAccessOverridesOk returns a tuple with the AccessOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessOverrides

`func (o *GetAccessRights200ResponseDataInner) SetAccessOverrides(v []AccessRightOverrideBean)`

SetAccessOverrides sets AccessOverrides field to given value.

### HasAccessOverrides

`func (o *GetAccessRights200ResponseDataInner) HasAccessOverrides() bool`

HasAccessOverrides returns a boolean if a field has been set.

### GetSelectiveDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) GetSelectiveDeploymentsAllowed() bool`

GetSelectiveDeploymentsAllowed returns the SelectiveDeploymentsAllowed field if non-nil, zero value otherwise.

### GetSelectiveDeploymentsAllowedOk

`func (o *GetAccessRights200ResponseDataInner) GetSelectiveDeploymentsAllowedOk() (*bool, bool)`

GetSelectiveDeploymentsAllowedOk returns a tuple with the SelectiveDeploymentsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) SetSelectiveDeploymentsAllowed(v bool)`

SetSelectiveDeploymentsAllowed sets SelectiveDeploymentsAllowed field to given value.

### HasSelectiveDeploymentsAllowed

`func (o *GetAccessRights200ResponseDataInner) HasSelectiveDeploymentsAllowed() bool`

HasSelectiveDeploymentsAllowed returns a boolean if a field has been set.

### GetLinks

`func (o *GetAccessRights200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAccessRights200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAccessRights200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAccessRights200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


