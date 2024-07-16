# SyslogServerBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | The port used to connect to the syslog server. | [optional] 
**SyslogProtocol** | Pointer to **string** | The syslog protocol used to format syslog messages. | [optional] 
**SeverityLevel** | Pointer to **string** | The severity level of messages that are sent to the syslog server. | [optional] 
**Transport** | Pointer to [**SyslogTransportBean**](SyslogTransportBean.md) |  | [optional] 
**Iso8601TimestampsEnabled** | Pointer to **bool** | Indicates whether the ISO 8601 timestamp format is used for syslog messages. | [optional] 
**ServiceTypes** | Pointer to **[]string** | The services for which syslog messages are generated. | [optional] 

## Methods

### NewSyslogServerBean

`func NewSyslogServerBean() *SyslogServerBean`

NewSyslogServerBean instantiates a new SyslogServerBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogServerBeanWithDefaults

`func NewSyslogServerBeanWithDefaults() *SyslogServerBean`

NewSyslogServerBeanWithDefaults instantiates a new SyslogServerBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *SyslogServerBean) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogServerBean) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogServerBean) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogServerBean) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSyslogProtocol

`func (o *SyslogServerBean) GetSyslogProtocol() string`

GetSyslogProtocol returns the SyslogProtocol field if non-nil, zero value otherwise.

### GetSyslogProtocolOk

`func (o *SyslogServerBean) GetSyslogProtocolOk() (*string, bool)`

GetSyslogProtocolOk returns a tuple with the SyslogProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogProtocol

`func (o *SyslogServerBean) SetSyslogProtocol(v string)`

SetSyslogProtocol sets SyslogProtocol field to given value.

### HasSyslogProtocol

`func (o *SyslogServerBean) HasSyslogProtocol() bool`

HasSyslogProtocol returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *SyslogServerBean) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *SyslogServerBean) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *SyslogServerBean) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *SyslogServerBean) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetTransport

`func (o *SyslogServerBean) GetTransport() SyslogTransportBean`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *SyslogServerBean) GetTransportOk() (*SyslogTransportBean, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *SyslogServerBean) SetTransport(v SyslogTransportBean)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *SyslogServerBean) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetIso8601TimestampsEnabled

`func (o *SyslogServerBean) GetIso8601TimestampsEnabled() bool`

GetIso8601TimestampsEnabled returns the Iso8601TimestampsEnabled field if non-nil, zero value otherwise.

### GetIso8601TimestampsEnabledOk

`func (o *SyslogServerBean) GetIso8601TimestampsEnabledOk() (*bool, bool)`

GetIso8601TimestampsEnabledOk returns a tuple with the Iso8601TimestampsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIso8601TimestampsEnabled

`func (o *SyslogServerBean) SetIso8601TimestampsEnabled(v bool)`

SetIso8601TimestampsEnabled sets Iso8601TimestampsEnabled field to given value.

### HasIso8601TimestampsEnabled

`func (o *SyslogServerBean) HasIso8601TimestampsEnabled() bool`

HasIso8601TimestampsEnabled returns a boolean if a field has been set.

### GetServiceTypes

`func (o *SyslogServerBean) GetServiceTypes() []string`

GetServiceTypes returns the ServiceTypes field if non-nil, zero value otherwise.

### GetServiceTypesOk

`func (o *SyslogServerBean) GetServiceTypesOk() (*[]string, bool)`

GetServiceTypesOk returns a tuple with the ServiceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypes

`func (o *SyslogServerBean) SetServiceTypes(v []string)`

SetServiceTypes sets ServiceTypes field to given value.

### HasServiceTypes

`func (o *SyslogServerBean) HasServiceTypes() bool`

HasServiceTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


