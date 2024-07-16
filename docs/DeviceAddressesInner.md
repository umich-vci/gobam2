# DeviceAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 
**Address** | **string** |  | 
**MacAddress** | Pointer to [**AddressMacAddress**](AddressMacAddress.md) |  | [optional] 
**Device** | Pointer to [**InlinedDevice**](InlinedDevice.md) |  | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**ResourceRecords** | Pointer to [**[]ResourceRecord**](ResourceRecord.md) |  | [optional] 
**LeaseDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address was offered. | [optional] [readonly] 
**LeaseExpirationDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address expires. | [optional] [readonly] 
**RemoteId** | Pointer to **string** | The ID of the relay agent. | [optional] [readonly] 

## Methods

### NewDeviceAddressesInner

`func NewDeviceAddressesInner(id int64, address string, ) *DeviceAddressesInner`

NewDeviceAddressesInner instantiates a new DeviceAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAddressesInnerWithDefaults

`func NewDeviceAddressesInnerWithDefaults() *DeviceAddressesInner`

NewDeviceAddressesInnerWithDefaults instantiates a new DeviceAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAddressesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAddressesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAddressesInner) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *DeviceAddressesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceAddressesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceAddressesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceAddressesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DeviceAddressesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAddressesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAddressesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceAddressesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *DeviceAddressesInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DeviceAddressesInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DeviceAddressesInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DeviceAddressesInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *DeviceAddressesInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DeviceAddressesInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DeviceAddressesInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DeviceAddressesInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetState

`func (o *DeviceAddressesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DeviceAddressesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DeviceAddressesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DeviceAddressesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAddress

`func (o *DeviceAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DeviceAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DeviceAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMacAddress

`func (o *DeviceAddressesInner) GetMacAddress() AddressMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *DeviceAddressesInner) GetMacAddressOk() (*AddressMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *DeviceAddressesInner) SetMacAddress(v AddressMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *DeviceAddressesInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDevice

`func (o *DeviceAddressesInner) GetDevice() InlinedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DeviceAddressesInner) GetDeviceOk() (*InlinedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DeviceAddressesInner) SetDevice(v InlinedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DeviceAddressesInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *DeviceAddressesInner) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeviceAddressesInner) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeviceAddressesInner) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeviceAddressesInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResourceRecords

`func (o *DeviceAddressesInner) GetResourceRecords() []ResourceRecord`

GetResourceRecords returns the ResourceRecords field if non-nil, zero value otherwise.

### GetResourceRecordsOk

`func (o *DeviceAddressesInner) GetResourceRecordsOk() (*[]ResourceRecord, bool)`

GetResourceRecordsOk returns a tuple with the ResourceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecords

`func (o *DeviceAddressesInner) SetResourceRecords(v []ResourceRecord)`

SetResourceRecords sets ResourceRecords field to given value.

### HasResourceRecords

`func (o *DeviceAddressesInner) HasResourceRecords() bool`

HasResourceRecords returns a boolean if a field has been set.

### GetLeaseDateTime

`func (o *DeviceAddressesInner) GetLeaseDateTime() time.Time`

GetLeaseDateTime returns the LeaseDateTime field if non-nil, zero value otherwise.

### GetLeaseDateTimeOk

`func (o *DeviceAddressesInner) GetLeaseDateTimeOk() (*time.Time, bool)`

GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDateTime

`func (o *DeviceAddressesInner) SetLeaseDateTime(v time.Time)`

SetLeaseDateTime sets LeaseDateTime field to given value.

### HasLeaseDateTime

`func (o *DeviceAddressesInner) HasLeaseDateTime() bool`

HasLeaseDateTime returns a boolean if a field has been set.

### GetLeaseExpirationDateTime

`func (o *DeviceAddressesInner) GetLeaseExpirationDateTime() time.Time`

GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field if non-nil, zero value otherwise.

### GetLeaseExpirationDateTimeOk

`func (o *DeviceAddressesInner) GetLeaseExpirationDateTimeOk() (*time.Time, bool)`

GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpirationDateTime

`func (o *DeviceAddressesInner) SetLeaseExpirationDateTime(v time.Time)`

SetLeaseExpirationDateTime sets LeaseExpirationDateTime field to given value.

### HasLeaseExpirationDateTime

`func (o *DeviceAddressesInner) HasLeaseExpirationDateTime() bool`

HasLeaseExpirationDateTime returns a boolean if a field has been set.

### GetRemoteId

`func (o *DeviceAddressesInner) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *DeviceAddressesInner) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *DeviceAddressesInner) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *DeviceAddressesInner) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


