# MACPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**InstantDeploymentEnabled** | Pointer to **bool** | Indicates whether MAC pool changes are instantly deployed to associated DNS/DHCP servers. | [optional] 

## Methods

### NewMACPool

`func NewMACPool() *MACPool`

NewMACPool instantiates a new MACPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMACPoolWithDefaults

`func NewMACPoolWithDefaults() *MACPool`

NewMACPoolWithDefaults instantiates a new MACPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MACPool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MACPool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MACPool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MACPool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetInstantDeploymentEnabled

`func (o *MACPool) GetInstantDeploymentEnabled() bool`

GetInstantDeploymentEnabled returns the InstantDeploymentEnabled field if non-nil, zero value otherwise.

### GetInstantDeploymentEnabledOk

`func (o *MACPool) GetInstantDeploymentEnabledOk() (*bool, bool)`

GetInstantDeploymentEnabledOk returns a tuple with the InstantDeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantDeploymentEnabled

`func (o *MACPool) SetInstantDeploymentEnabled(v bool)`

SetInstantDeploymentEnabled sets InstantDeploymentEnabled field to given value.

### HasInstantDeploymentEnabled

`func (o *MACPool) HasInstantDeploymentEnabled() bool`

HasInstantDeploymentEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


