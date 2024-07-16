# SyslogRedirectionServerBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerType** | Pointer to **string** | The syslog redirection server type. | [optional] 
**Address** | Pointer to **string** | The IP address of the syslog server. | [optional] 

## Methods

### NewSyslogRedirectionServerBean

`func NewSyslogRedirectionServerBean() *SyslogRedirectionServerBean`

NewSyslogRedirectionServerBean instantiates a new SyslogRedirectionServerBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogRedirectionServerBeanWithDefaults

`func NewSyslogRedirectionServerBeanWithDefaults() *SyslogRedirectionServerBean`

NewSyslogRedirectionServerBeanWithDefaults instantiates a new SyslogRedirectionServerBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerType

`func (o *SyslogRedirectionServerBean) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *SyslogRedirectionServerBean) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *SyslogRedirectionServerBean) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *SyslogRedirectionServerBean) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetAddress

`func (o *SyslogRedirectionServerBean) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SyslogRedirectionServerBean) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SyslogRedirectionServerBean) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SyslogRedirectionServerBean) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


