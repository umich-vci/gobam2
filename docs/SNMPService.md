# SNMPService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the SNMP service is enabled. | [optional] 
**ManualOverrideEnabled** | Pointer to **bool** | Indicates whether the SNMP service configuration has been manually overridden. | [optional] [readonly] 
**ContactName** | Pointer to **string** | The system name to be reported through SNMP. | [optional] 
**ContactEmail** | Pointer to **string** | The email address for the system contact to be reported through SNMP. | [optional] 
**Location** | Pointer to **string** | The description of the system&#39;s location to be reported through SNMP. | [optional] 
**Description** | Pointer to **string** | The description of the system to be reported through SNMP. | [optional] 
**PollingInterval** | Pointer to **string** | The SNMP polling interval, in ISO 8601 duration format, that determines the frequency that SNMP values are refreshed on the system. | [optional] 
**LogLevel** | Pointer to **string** | The logging level for SNMP events. | [optional] 
**Snmpv1** | Pointer to [**Snmpv1Bean**](Snmpv1Bean.md) |  | [optional] 
**Snmpv2c** | Pointer to [**Snmpv2cBean**](Snmpv2cBean.md) |  | [optional] 
**Snmpv3** | Pointer to [**Snmpv3Bean**](Snmpv3Bean.md) |  | [optional] 
**TrapServers** | Pointer to [**[]SnmpTrapServerBean**](SnmpTrapServerBean.md) |  | [optional] 

## Methods

### NewSNMPService

`func NewSNMPService() *SNMPService`

NewSNMPService instantiates a new SNMPService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPServiceWithDefaults

`func NewSNMPServiceWithDefaults() *SNMPService`

NewSNMPServiceWithDefaults instantiates a new SNMPService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SNMPService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SNMPService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SNMPService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SNMPService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *SNMPService) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SNMPService) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SNMPService) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SNMPService) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetManualOverrideEnabled

`func (o *SNMPService) GetManualOverrideEnabled() bool`

GetManualOverrideEnabled returns the ManualOverrideEnabled field if non-nil, zero value otherwise.

### GetManualOverrideEnabledOk

`func (o *SNMPService) GetManualOverrideEnabledOk() (*bool, bool)`

GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualOverrideEnabled

`func (o *SNMPService) SetManualOverrideEnabled(v bool)`

SetManualOverrideEnabled sets ManualOverrideEnabled field to given value.

### HasManualOverrideEnabled

`func (o *SNMPService) HasManualOverrideEnabled() bool`

HasManualOverrideEnabled returns a boolean if a field has been set.

### GetContactName

`func (o *SNMPService) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *SNMPService) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *SNMPService) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *SNMPService) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetContactEmail

`func (o *SNMPService) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *SNMPService) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *SNMPService) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *SNMPService) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetLocation

`func (o *SNMPService) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SNMPService) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SNMPService) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SNMPService) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDescription

`func (o *SNMPService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SNMPService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SNMPService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SNMPService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPollingInterval

`func (o *SNMPService) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *SNMPService) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *SNMPService) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *SNMPService) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetLogLevel

`func (o *SNMPService) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *SNMPService) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *SNMPService) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *SNMPService) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetSnmpv1

`func (o *SNMPService) GetSnmpv1() Snmpv1Bean`

GetSnmpv1 returns the Snmpv1 field if non-nil, zero value otherwise.

### GetSnmpv1Ok

`func (o *SNMPService) GetSnmpv1Ok() (*Snmpv1Bean, bool)`

GetSnmpv1Ok returns a tuple with the Snmpv1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1

`func (o *SNMPService) SetSnmpv1(v Snmpv1Bean)`

SetSnmpv1 sets Snmpv1 field to given value.

### HasSnmpv1

`func (o *SNMPService) HasSnmpv1() bool`

HasSnmpv1 returns a boolean if a field has been set.

### GetSnmpv2c

`func (o *SNMPService) GetSnmpv2c() Snmpv2cBean`

GetSnmpv2c returns the Snmpv2c field if non-nil, zero value otherwise.

### GetSnmpv2cOk

`func (o *SNMPService) GetSnmpv2cOk() (*Snmpv2cBean, bool)`

GetSnmpv2cOk returns a tuple with the Snmpv2c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv2c

`func (o *SNMPService) SetSnmpv2c(v Snmpv2cBean)`

SetSnmpv2c sets Snmpv2c field to given value.

### HasSnmpv2c

`func (o *SNMPService) HasSnmpv2c() bool`

HasSnmpv2c returns a boolean if a field has been set.

### GetSnmpv3

`func (o *SNMPService) GetSnmpv3() Snmpv3Bean`

GetSnmpv3 returns the Snmpv3 field if non-nil, zero value otherwise.

### GetSnmpv3Ok

`func (o *SNMPService) GetSnmpv3Ok() (*Snmpv3Bean, bool)`

GetSnmpv3Ok returns a tuple with the Snmpv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3

`func (o *SNMPService) SetSnmpv3(v Snmpv3Bean)`

SetSnmpv3 sets Snmpv3 field to given value.

### HasSnmpv3

`func (o *SNMPService) HasSnmpv3() bool`

HasSnmpv3 returns a boolean if a field has been set.

### GetTrapServers

`func (o *SNMPService) GetTrapServers() []SnmpTrapServerBean`

GetTrapServers returns the TrapServers field if non-nil, zero value otherwise.

### GetTrapServersOk

`func (o *SNMPService) GetTrapServersOk() (*[]SnmpTrapServerBean, bool)`

GetTrapServersOk returns a tuple with the TrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapServers

`func (o *SNMPService) SetTrapServers(v []SnmpTrapServerBean)`

SetTrapServers sets TrapServers field to given value.

### HasTrapServers

`func (o *SNMPService) HasTrapServers() bool`

HasTrapServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


