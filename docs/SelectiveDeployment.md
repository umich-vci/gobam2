# SelectiveDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Resources** | Pointer to [**[]InlinedDeployableResource**](InlinedDeployableResource.md) | A list of resource record and IP address resources to deploy. | [optional] 
**BatchMode** | Pointer to **string** | The method used for batching selective deployment tasks. | [optional] 
**RecordScope** | Pointer to **string** | Defines whether selective deployment tasks include resources that are related to the defined DNS records. | [optional] 
**ContinueOnFailureEnabled** | Pointer to **bool** | Indicates the mode of operation on a failed resource record. If false, the deployment stops when a record fails. If true, the deployment continues when a record fails and the deployment moves to the next record. | [optional] 
**DynamicRecordDeploymentStrategy** | Pointer to **string** | Defines how dynamic records are handled with selective deployment tasks. | [optional] 

## Methods

### NewSelectiveDeployment

`func NewSelectiveDeployment() *SelectiveDeployment`

NewSelectiveDeployment instantiates a new SelectiveDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectiveDeploymentWithDefaults

`func NewSelectiveDeploymentWithDefaults() *SelectiveDeployment`

NewSelectiveDeploymentWithDefaults instantiates a new SelectiveDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SelectiveDeployment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectiveDeployment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectiveDeployment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SelectiveDeployment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *SelectiveDeployment) GetResources() []InlinedDeployableResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *SelectiveDeployment) GetResourcesOk() (*[]InlinedDeployableResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *SelectiveDeployment) SetResources(v []InlinedDeployableResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *SelectiveDeployment) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetBatchMode

`func (o *SelectiveDeployment) GetBatchMode() string`

GetBatchMode returns the BatchMode field if non-nil, zero value otherwise.

### GetBatchModeOk

`func (o *SelectiveDeployment) GetBatchModeOk() (*string, bool)`

GetBatchModeOk returns a tuple with the BatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchMode

`func (o *SelectiveDeployment) SetBatchMode(v string)`

SetBatchMode sets BatchMode field to given value.

### HasBatchMode

`func (o *SelectiveDeployment) HasBatchMode() bool`

HasBatchMode returns a boolean if a field has been set.

### GetRecordScope

`func (o *SelectiveDeployment) GetRecordScope() string`

GetRecordScope returns the RecordScope field if non-nil, zero value otherwise.

### GetRecordScopeOk

`func (o *SelectiveDeployment) GetRecordScopeOk() (*string, bool)`

GetRecordScopeOk returns a tuple with the RecordScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordScope

`func (o *SelectiveDeployment) SetRecordScope(v string)`

SetRecordScope sets RecordScope field to given value.

### HasRecordScope

`func (o *SelectiveDeployment) HasRecordScope() bool`

HasRecordScope returns a boolean if a field has been set.

### GetContinueOnFailureEnabled

`func (o *SelectiveDeployment) GetContinueOnFailureEnabled() bool`

GetContinueOnFailureEnabled returns the ContinueOnFailureEnabled field if non-nil, zero value otherwise.

### GetContinueOnFailureEnabledOk

`func (o *SelectiveDeployment) GetContinueOnFailureEnabledOk() (*bool, bool)`

GetContinueOnFailureEnabledOk returns a tuple with the ContinueOnFailureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnFailureEnabled

`func (o *SelectiveDeployment) SetContinueOnFailureEnabled(v bool)`

SetContinueOnFailureEnabled sets ContinueOnFailureEnabled field to given value.

### HasContinueOnFailureEnabled

`func (o *SelectiveDeployment) HasContinueOnFailureEnabled() bool`

HasContinueOnFailureEnabled returns a boolean if a field has been set.

### GetDynamicRecordDeploymentStrategy

`func (o *SelectiveDeployment) GetDynamicRecordDeploymentStrategy() string`

GetDynamicRecordDeploymentStrategy returns the DynamicRecordDeploymentStrategy field if non-nil, zero value otherwise.

### GetDynamicRecordDeploymentStrategyOk

`func (o *SelectiveDeployment) GetDynamicRecordDeploymentStrategyOk() (*string, bool)`

GetDynamicRecordDeploymentStrategyOk returns a tuple with the DynamicRecordDeploymentStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicRecordDeploymentStrategy

`func (o *SelectiveDeployment) SetDynamicRecordDeploymentStrategy(v string)`

SetDynamicRecordDeploymentStrategy sets DynamicRecordDeploymentStrategy field to given value.

### HasDynamicRecordDeploymentStrategy

`func (o *SelectiveDeployment) HasDynamicRecordDeploymentStrategy() bool`

HasDynamicRecordDeploymentStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


