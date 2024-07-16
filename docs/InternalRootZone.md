# InternalRootZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**DeploymentEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewInternalRootZone

`func NewInternalRootZone() *InternalRootZone`

NewInternalRootZone instantiates a new InternalRootZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalRootZoneWithDefaults

`func NewInternalRootZoneWithDefaults() *InternalRootZone`

NewInternalRootZoneWithDefaults instantiates a new InternalRootZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InternalRootZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternalRootZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternalRootZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InternalRootZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *InternalRootZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalRootZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalRootZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalRootZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *InternalRootZone) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *InternalRootZone) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *InternalRootZone) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.

### HasDeploymentEnabled

`func (o *InternalRootZone) HasDeploymentEnabled() bool`

HasDeploymentEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


