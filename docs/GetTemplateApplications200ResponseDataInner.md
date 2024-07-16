# GetTemplateApplications200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**State** | Pointer to **string** | The state of the template application operation. | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The date and time that the template application operation started. | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The date and time that the template application operation completed. | [optional] [readonly] 
**ConflictResolutionStrategy** | Pointer to **string** | The method used to resolve conflicts between template items and the resource the template is applied to. | [optional] 
**DeploymentOptionsApplied** | Pointer to **bool** | Indicates whether deployment options set in the template are applied. | [optional] 
**GatewayItemApplied** | Pointer to **bool** | Indicates whether the template settings for the gateway address are applied. | [optional] 
**ReservedRangeItemsApplied** | Pointer to **bool** | Indicates whether the reserved address ranges set in the template are applied. | [optional] 
**DhcpRangeItemsApplied** | Pointer to **bool** | Indicates whether the DHCP reserved address ranges set in the template are applied. | [optional] 
**IpGroupItemsApplied** | Pointer to **bool** | Indicates whether the IP groups set in the template are applied. | [optional] 
**OrphanedAddressConversionStrategy** | Pointer to **string** | Determines whether orphaned DHCP allocated IPv4 addresses are converted into another assignment type and the type of host record assigned to the orphaned record. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetTemplateApplications200ResponseDataInner

`func NewGetTemplateApplications200ResponseDataInner() *GetTemplateApplications200ResponseDataInner`

NewGetTemplateApplications200ResponseDataInner instantiates a new GetTemplateApplications200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTemplateApplications200ResponseDataInnerWithDefaults

`func NewGetTemplateApplications200ResponseDataInnerWithDefaults() *GetTemplateApplications200ResponseDataInner`

NewGetTemplateApplications200ResponseDataInnerWithDefaults instantiates a new GetTemplateApplications200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTemplateApplications200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTemplateApplications200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTemplateApplications200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTemplateApplications200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetTemplateApplications200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTemplateApplications200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTemplateApplications200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTemplateApplications200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *GetTemplateApplications200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetTemplateApplications200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetTemplateApplications200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetTemplateApplications200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStartDateTime

`func (o *GetTemplateApplications200ResponseDataInner) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *GetTemplateApplications200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *GetTemplateApplications200ResponseDataInner) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *GetTemplateApplications200ResponseDataInner) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *GetTemplateApplications200ResponseDataInner) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *GetTemplateApplications200ResponseDataInner) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *GetTemplateApplications200ResponseDataInner) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *GetTemplateApplications200ResponseDataInner) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.

### GetConflictResolutionStrategy

`func (o *GetTemplateApplications200ResponseDataInner) GetConflictResolutionStrategy() string`

GetConflictResolutionStrategy returns the ConflictResolutionStrategy field if non-nil, zero value otherwise.

### GetConflictResolutionStrategyOk

`func (o *GetTemplateApplications200ResponseDataInner) GetConflictResolutionStrategyOk() (*string, bool)`

GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolutionStrategy

`func (o *GetTemplateApplications200ResponseDataInner) SetConflictResolutionStrategy(v string)`

SetConflictResolutionStrategy sets ConflictResolutionStrategy field to given value.

### HasConflictResolutionStrategy

`func (o *GetTemplateApplications200ResponseDataInner) HasConflictResolutionStrategy() bool`

HasConflictResolutionStrategy returns a boolean if a field has been set.

### GetDeploymentOptionsApplied

`func (o *GetTemplateApplications200ResponseDataInner) GetDeploymentOptionsApplied() bool`

GetDeploymentOptionsApplied returns the DeploymentOptionsApplied field if non-nil, zero value otherwise.

### GetDeploymentOptionsAppliedOk

`func (o *GetTemplateApplications200ResponseDataInner) GetDeploymentOptionsAppliedOk() (*bool, bool)`

GetDeploymentOptionsAppliedOk returns a tuple with the DeploymentOptionsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentOptionsApplied

`func (o *GetTemplateApplications200ResponseDataInner) SetDeploymentOptionsApplied(v bool)`

SetDeploymentOptionsApplied sets DeploymentOptionsApplied field to given value.

### HasDeploymentOptionsApplied

`func (o *GetTemplateApplications200ResponseDataInner) HasDeploymentOptionsApplied() bool`

HasDeploymentOptionsApplied returns a boolean if a field has been set.

### GetGatewayItemApplied

`func (o *GetTemplateApplications200ResponseDataInner) GetGatewayItemApplied() bool`

GetGatewayItemApplied returns the GatewayItemApplied field if non-nil, zero value otherwise.

### GetGatewayItemAppliedOk

`func (o *GetTemplateApplications200ResponseDataInner) GetGatewayItemAppliedOk() (*bool, bool)`

GetGatewayItemAppliedOk returns a tuple with the GatewayItemApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayItemApplied

`func (o *GetTemplateApplications200ResponseDataInner) SetGatewayItemApplied(v bool)`

SetGatewayItemApplied sets GatewayItemApplied field to given value.

### HasGatewayItemApplied

`func (o *GetTemplateApplications200ResponseDataInner) HasGatewayItemApplied() bool`

HasGatewayItemApplied returns a boolean if a field has been set.

### GetReservedRangeItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) GetReservedRangeItemsApplied() bool`

GetReservedRangeItemsApplied returns the ReservedRangeItemsApplied field if non-nil, zero value otherwise.

### GetReservedRangeItemsAppliedOk

`func (o *GetTemplateApplications200ResponseDataInner) GetReservedRangeItemsAppliedOk() (*bool, bool)`

GetReservedRangeItemsAppliedOk returns a tuple with the ReservedRangeItemsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedRangeItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) SetReservedRangeItemsApplied(v bool)`

SetReservedRangeItemsApplied sets ReservedRangeItemsApplied field to given value.

### HasReservedRangeItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) HasReservedRangeItemsApplied() bool`

HasReservedRangeItemsApplied returns a boolean if a field has been set.

### GetDhcpRangeItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) GetDhcpRangeItemsApplied() bool`

GetDhcpRangeItemsApplied returns the DhcpRangeItemsApplied field if non-nil, zero value otherwise.

### GetDhcpRangeItemsAppliedOk

`func (o *GetTemplateApplications200ResponseDataInner) GetDhcpRangeItemsAppliedOk() (*bool, bool)`

GetDhcpRangeItemsAppliedOk returns a tuple with the DhcpRangeItemsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRangeItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) SetDhcpRangeItemsApplied(v bool)`

SetDhcpRangeItemsApplied sets DhcpRangeItemsApplied field to given value.

### HasDhcpRangeItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) HasDhcpRangeItemsApplied() bool`

HasDhcpRangeItemsApplied returns a boolean if a field has been set.

### GetIpGroupItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) GetIpGroupItemsApplied() bool`

GetIpGroupItemsApplied returns the IpGroupItemsApplied field if non-nil, zero value otherwise.

### GetIpGroupItemsAppliedOk

`func (o *GetTemplateApplications200ResponseDataInner) GetIpGroupItemsAppliedOk() (*bool, bool)`

GetIpGroupItemsAppliedOk returns a tuple with the IpGroupItemsApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpGroupItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) SetIpGroupItemsApplied(v bool)`

SetIpGroupItemsApplied sets IpGroupItemsApplied field to given value.

### HasIpGroupItemsApplied

`func (o *GetTemplateApplications200ResponseDataInner) HasIpGroupItemsApplied() bool`

HasIpGroupItemsApplied returns a boolean if a field has been set.

### GetOrphanedAddressConversionStrategy

`func (o *GetTemplateApplications200ResponseDataInner) GetOrphanedAddressConversionStrategy() string`

GetOrphanedAddressConversionStrategy returns the OrphanedAddressConversionStrategy field if non-nil, zero value otherwise.

### GetOrphanedAddressConversionStrategyOk

`func (o *GetTemplateApplications200ResponseDataInner) GetOrphanedAddressConversionStrategyOk() (*string, bool)`

GetOrphanedAddressConversionStrategyOk returns a tuple with the OrphanedAddressConversionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanedAddressConversionStrategy

`func (o *GetTemplateApplications200ResponseDataInner) SetOrphanedAddressConversionStrategy(v string)`

SetOrphanedAddressConversionStrategy sets OrphanedAddressConversionStrategy field to given value.

### HasOrphanedAddressConversionStrategy

`func (o *GetTemplateApplications200ResponseDataInner) HasOrphanedAddressConversionStrategy() bool`

HasOrphanedAddressConversionStrategy returns a boolean if a field has been set.

### GetLinks

`func (o *GetTemplateApplications200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTemplateApplications200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTemplateApplications200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetTemplateApplications200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


