# LogLevelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**AddressManager** | Pointer to **string** |  | [optional] 
**RestV2** | Pointer to **string** |  | [optional] 
**ZoneImport** | Pointer to **string** |  | [optional] 
**ReconciliationService** | Pointer to **string** |  | [optional] 
**DiscoveryEngine** | Pointer to **string** |  | [optional] 
**SnmpClient** | Pointer to **string** |  | [optional] 
**MonitoringService** | Pointer to **string** |  | [optional] 
**RrdTool** | Pointer to **string** |  | [optional] 

## Methods

### NewLogLevelSettings

`func NewLogLevelSettings() *LogLevelSettings`

NewLogLevelSettings instantiates a new LogLevelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogLevelSettingsWithDefaults

`func NewLogLevelSettingsWithDefaults() *LogLevelSettings`

NewLogLevelSettingsWithDefaults instantiates a new LogLevelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LogLevelSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogLevelSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogLevelSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogLevelSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddressManager

`func (o *LogLevelSettings) GetAddressManager() string`

GetAddressManager returns the AddressManager field if non-nil, zero value otherwise.

### GetAddressManagerOk

`func (o *LogLevelSettings) GetAddressManagerOk() (*string, bool)`

GetAddressManagerOk returns a tuple with the AddressManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressManager

`func (o *LogLevelSettings) SetAddressManager(v string)`

SetAddressManager sets AddressManager field to given value.

### HasAddressManager

`func (o *LogLevelSettings) HasAddressManager() bool`

HasAddressManager returns a boolean if a field has been set.

### GetRestV2

`func (o *LogLevelSettings) GetRestV2() string`

GetRestV2 returns the RestV2 field if non-nil, zero value otherwise.

### GetRestV2Ok

`func (o *LogLevelSettings) GetRestV2Ok() (*string, bool)`

GetRestV2Ok returns a tuple with the RestV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestV2

`func (o *LogLevelSettings) SetRestV2(v string)`

SetRestV2 sets RestV2 field to given value.

### HasRestV2

`func (o *LogLevelSettings) HasRestV2() bool`

HasRestV2 returns a boolean if a field has been set.

### GetZoneImport

`func (o *LogLevelSettings) GetZoneImport() string`

GetZoneImport returns the ZoneImport field if non-nil, zero value otherwise.

### GetZoneImportOk

`func (o *LogLevelSettings) GetZoneImportOk() (*string, bool)`

GetZoneImportOk returns a tuple with the ZoneImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneImport

`func (o *LogLevelSettings) SetZoneImport(v string)`

SetZoneImport sets ZoneImport field to given value.

### HasZoneImport

`func (o *LogLevelSettings) HasZoneImport() bool`

HasZoneImport returns a boolean if a field has been set.

### GetReconciliationService

`func (o *LogLevelSettings) GetReconciliationService() string`

GetReconciliationService returns the ReconciliationService field if non-nil, zero value otherwise.

### GetReconciliationServiceOk

`func (o *LogLevelSettings) GetReconciliationServiceOk() (*string, bool)`

GetReconciliationServiceOk returns a tuple with the ReconciliationService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconciliationService

`func (o *LogLevelSettings) SetReconciliationService(v string)`

SetReconciliationService sets ReconciliationService field to given value.

### HasReconciliationService

`func (o *LogLevelSettings) HasReconciliationService() bool`

HasReconciliationService returns a boolean if a field has been set.

### GetDiscoveryEngine

`func (o *LogLevelSettings) GetDiscoveryEngine() string`

GetDiscoveryEngine returns the DiscoveryEngine field if non-nil, zero value otherwise.

### GetDiscoveryEngineOk

`func (o *LogLevelSettings) GetDiscoveryEngineOk() (*string, bool)`

GetDiscoveryEngineOk returns a tuple with the DiscoveryEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEngine

`func (o *LogLevelSettings) SetDiscoveryEngine(v string)`

SetDiscoveryEngine sets DiscoveryEngine field to given value.

### HasDiscoveryEngine

`func (o *LogLevelSettings) HasDiscoveryEngine() bool`

HasDiscoveryEngine returns a boolean if a field has been set.

### GetSnmpClient

`func (o *LogLevelSettings) GetSnmpClient() string`

GetSnmpClient returns the SnmpClient field if non-nil, zero value otherwise.

### GetSnmpClientOk

`func (o *LogLevelSettings) GetSnmpClientOk() (*string, bool)`

GetSnmpClientOk returns a tuple with the SnmpClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpClient

`func (o *LogLevelSettings) SetSnmpClient(v string)`

SetSnmpClient sets SnmpClient field to given value.

### HasSnmpClient

`func (o *LogLevelSettings) HasSnmpClient() bool`

HasSnmpClient returns a boolean if a field has been set.

### GetMonitoringService

`func (o *LogLevelSettings) GetMonitoringService() string`

GetMonitoringService returns the MonitoringService field if non-nil, zero value otherwise.

### GetMonitoringServiceOk

`func (o *LogLevelSettings) GetMonitoringServiceOk() (*string, bool)`

GetMonitoringServiceOk returns a tuple with the MonitoringService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringService

`func (o *LogLevelSettings) SetMonitoringService(v string)`

SetMonitoringService sets MonitoringService field to given value.

### HasMonitoringService

`func (o *LogLevelSettings) HasMonitoringService() bool`

HasMonitoringService returns a boolean if a field has been set.

### GetRrdTool

`func (o *LogLevelSettings) GetRrdTool() string`

GetRrdTool returns the RrdTool field if non-nil, zero value otherwise.

### GetRrdToolOk

`func (o *LogLevelSettings) GetRrdToolOk() (*string, bool)`

GetRrdToolOk returns a tuple with the RrdTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrdTool

`func (o *LogLevelSettings) SetRrdTool(v string)`

SetRrdTool sets RrdTool field to given value.

### HasRrdTool

`func (o *LogLevelSettings) HasRrdTool() bool`

HasRrdTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


