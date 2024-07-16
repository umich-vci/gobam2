# SnmpTrapServerBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether the SNMP trap server is enabled. | [optional] 
**Address** | Pointer to **string** | The IP address of the SNMP trap server. | [optional] 
**Port** | Pointer to **int32** | The SNMP trap server port. | [optional] 
**Snmp** | Pointer to [**SnmpBean**](SnmpBean.md) |  | [optional] 

## Methods

### NewSnmpTrapServerBean

`func NewSnmpTrapServerBean() *SnmpTrapServerBean`

NewSnmpTrapServerBean instantiates a new SnmpTrapServerBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpTrapServerBeanWithDefaults

`func NewSnmpTrapServerBeanWithDefaults() *SnmpTrapServerBean`

NewSnmpTrapServerBeanWithDefaults instantiates a new SnmpTrapServerBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SnmpTrapServerBean) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpTrapServerBean) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpTrapServerBean) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SnmpTrapServerBean) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAddress

`func (o *SnmpTrapServerBean) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SnmpTrapServerBean) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SnmpTrapServerBean) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SnmpTrapServerBean) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *SnmpTrapServerBean) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SnmpTrapServerBean) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SnmpTrapServerBean) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SnmpTrapServerBean) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSnmp

`func (o *SnmpTrapServerBean) GetSnmp() SnmpBean`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *SnmpTrapServerBean) GetSnmpOk() (*SnmpBean, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *SnmpTrapServerBean) SetSnmp(v SnmpBean)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *SnmpTrapServerBean) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


