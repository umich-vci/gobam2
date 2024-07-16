# SnmpBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The SNMP version for the router or switch. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether SNMP service is enabled. | [optional] 
**Port** | Pointer to **int32** | The SNMP port that is used to communicate with the router or switch. | [optional] 

## Methods

### NewSnmpBean

`func NewSnmpBean() *SnmpBean`

NewSnmpBean instantiates a new SnmpBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpBeanWithDefaults

`func NewSnmpBeanWithDefaults() *SnmpBean`

NewSnmpBeanWithDefaults instantiates a new SnmpBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SnmpBean) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnmpBean) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnmpBean) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnmpBean) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPort

`func (o *SnmpBean) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpBean) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpBean) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpBean) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


