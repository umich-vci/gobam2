# SyslogTransportBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | The syslog transport protocol. | [optional] 

## Methods

### NewSyslogTransportBean

`func NewSyslogTransportBean() *SyslogTransportBean`

NewSyslogTransportBean instantiates a new SyslogTransportBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogTransportBeanWithDefaults

`func NewSyslogTransportBeanWithDefaults() *SyslogTransportBean`

NewSyslogTransportBeanWithDefaults instantiates a new SyslogTransportBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *SyslogTransportBean) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyslogTransportBean) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyslogTransportBean) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyslogTransportBean) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


