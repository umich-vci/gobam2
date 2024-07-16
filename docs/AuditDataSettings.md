# AuditDataSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**AdministrativeHistoryRetentionPeriod** | Pointer to **string** | The number of days to retain the administrative history records. A null retention period indicates that the data will be retained indefinitely. | [optional] 
**SessionAndEventsRetentionPeriod** | Pointer to **string** | The number of days to retain the session and event history records. A null retention period indicates that the data will be retained indefinitely. | [optional] 
**DhcpTransactionRetentionPeriod** | Pointer to **string** | The number of days to retain the DHCP transaction history records. A null retention period indicates that the data will be retained indefinitely. A retention period of 0 seconds indicates that the data wil not be retained. | [optional] 
**DdnsTransactionRetentionPeriod** | Pointer to **string** | The number of days to retain the DDNS transaction history records. A null retention period indicates that the data will be retained indefinitely. A retention period of 0 seconds indicates that the data wil not be retained. | [optional] 
**ExportEnabled** | Pointer to **bool** | Indicates if the audit data export service is enabled. | [optional] 
**Destination** | Pointer to [**SinkBean**](SinkBean.md) |  | [optional] 

## Methods

### NewAuditDataSettings

`func NewAuditDataSettings() *AuditDataSettings`

NewAuditDataSettings instantiates a new AuditDataSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditDataSettingsWithDefaults

`func NewAuditDataSettingsWithDefaults() *AuditDataSettings`

NewAuditDataSettingsWithDefaults instantiates a new AuditDataSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AuditDataSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditDataSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditDataSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditDataSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdministrativeHistoryRetentionPeriod

`func (o *AuditDataSettings) GetAdministrativeHistoryRetentionPeriod() string`

GetAdministrativeHistoryRetentionPeriod returns the AdministrativeHistoryRetentionPeriod field if non-nil, zero value otherwise.

### GetAdministrativeHistoryRetentionPeriodOk

`func (o *AuditDataSettings) GetAdministrativeHistoryRetentionPeriodOk() (*string, bool)`

GetAdministrativeHistoryRetentionPeriodOk returns a tuple with the AdministrativeHistoryRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeHistoryRetentionPeriod

`func (o *AuditDataSettings) SetAdministrativeHistoryRetentionPeriod(v string)`

SetAdministrativeHistoryRetentionPeriod sets AdministrativeHistoryRetentionPeriod field to given value.

### HasAdministrativeHistoryRetentionPeriod

`func (o *AuditDataSettings) HasAdministrativeHistoryRetentionPeriod() bool`

HasAdministrativeHistoryRetentionPeriod returns a boolean if a field has been set.

### GetSessionAndEventsRetentionPeriod

`func (o *AuditDataSettings) GetSessionAndEventsRetentionPeriod() string`

GetSessionAndEventsRetentionPeriod returns the SessionAndEventsRetentionPeriod field if non-nil, zero value otherwise.

### GetSessionAndEventsRetentionPeriodOk

`func (o *AuditDataSettings) GetSessionAndEventsRetentionPeriodOk() (*string, bool)`

GetSessionAndEventsRetentionPeriodOk returns a tuple with the SessionAndEventsRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAndEventsRetentionPeriod

`func (o *AuditDataSettings) SetSessionAndEventsRetentionPeriod(v string)`

SetSessionAndEventsRetentionPeriod sets SessionAndEventsRetentionPeriod field to given value.

### HasSessionAndEventsRetentionPeriod

`func (o *AuditDataSettings) HasSessionAndEventsRetentionPeriod() bool`

HasSessionAndEventsRetentionPeriod returns a boolean if a field has been set.

### GetDhcpTransactionRetentionPeriod

`func (o *AuditDataSettings) GetDhcpTransactionRetentionPeriod() string`

GetDhcpTransactionRetentionPeriod returns the DhcpTransactionRetentionPeriod field if non-nil, zero value otherwise.

### GetDhcpTransactionRetentionPeriodOk

`func (o *AuditDataSettings) GetDhcpTransactionRetentionPeriodOk() (*string, bool)`

GetDhcpTransactionRetentionPeriodOk returns a tuple with the DhcpTransactionRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTransactionRetentionPeriod

`func (o *AuditDataSettings) SetDhcpTransactionRetentionPeriod(v string)`

SetDhcpTransactionRetentionPeriod sets DhcpTransactionRetentionPeriod field to given value.

### HasDhcpTransactionRetentionPeriod

`func (o *AuditDataSettings) HasDhcpTransactionRetentionPeriod() bool`

HasDhcpTransactionRetentionPeriod returns a boolean if a field has been set.

### GetDdnsTransactionRetentionPeriod

`func (o *AuditDataSettings) GetDdnsTransactionRetentionPeriod() string`

GetDdnsTransactionRetentionPeriod returns the DdnsTransactionRetentionPeriod field if non-nil, zero value otherwise.

### GetDdnsTransactionRetentionPeriodOk

`func (o *AuditDataSettings) GetDdnsTransactionRetentionPeriodOk() (*string, bool)`

GetDdnsTransactionRetentionPeriodOk returns a tuple with the DdnsTransactionRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsTransactionRetentionPeriod

`func (o *AuditDataSettings) SetDdnsTransactionRetentionPeriod(v string)`

SetDdnsTransactionRetentionPeriod sets DdnsTransactionRetentionPeriod field to given value.

### HasDdnsTransactionRetentionPeriod

`func (o *AuditDataSettings) HasDdnsTransactionRetentionPeriod() bool`

HasDdnsTransactionRetentionPeriod returns a boolean if a field has been set.

### GetExportEnabled

`func (o *AuditDataSettings) GetExportEnabled() bool`

GetExportEnabled returns the ExportEnabled field if non-nil, zero value otherwise.

### GetExportEnabledOk

`func (o *AuditDataSettings) GetExportEnabledOk() (*bool, bool)`

GetExportEnabledOk returns a tuple with the ExportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportEnabled

`func (o *AuditDataSettings) SetExportEnabled(v bool)`

SetExportEnabled sets ExportEnabled field to given value.

### HasExportEnabled

`func (o *AuditDataSettings) HasExportEnabled() bool`

HasExportEnabled returns a boolean if a field has been set.

### GetDestination

`func (o *AuditDataSettings) GetDestination() SinkBean`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *AuditDataSettings) GetDestinationOk() (*SinkBean, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *AuditDataSettings) SetDestination(v SinkBean)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *AuditDataSettings) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


