# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of IP address. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**State** | Pointer to **string** | The state of the IP address. | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to [**AddressMacAddress**](AddressMacAddress.md) |  | [optional] 
**Device** | Pointer to [**InlinedDevice**](InlinedDevice.md) |  | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**ResourceRecords** | Pointer to [**[]ResourceRecord**](ResourceRecord.md) |  | [optional] 
**LeaseDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address was offered. | [optional] [readonly] 
**LeaseExpirationDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address expires. | [optional] [readonly] 
**RemoteId** | Pointer to **string** | The ID of the relay agent. | [optional] [readonly] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Address) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Address) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Address) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Address) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Address) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Address) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Address) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Address) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Address) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *Address) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *Address) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *Address) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *Address) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *Address) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Address) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Address) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Address) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAddress

`func (o *Address) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Address) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Address) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Address) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *Address) GetMacAddress() AddressMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Address) GetMacAddressOk() (*AddressMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Address) SetMacAddress(v AddressMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Address) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDevice

`func (o *Address) GetDevice() InlinedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Address) GetDeviceOk() (*InlinedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Address) SetDevice(v InlinedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Address) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *Address) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Address) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Address) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Address) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResourceRecords

`func (o *Address) GetResourceRecords() []ResourceRecord`

GetResourceRecords returns the ResourceRecords field if non-nil, zero value otherwise.

### GetResourceRecordsOk

`func (o *Address) GetResourceRecordsOk() (*[]ResourceRecord, bool)`

GetResourceRecordsOk returns a tuple with the ResourceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecords

`func (o *Address) SetResourceRecords(v []ResourceRecord)`

SetResourceRecords sets ResourceRecords field to given value.

### HasResourceRecords

`func (o *Address) HasResourceRecords() bool`

HasResourceRecords returns a boolean if a field has been set.

### GetLeaseDateTime

`func (o *Address) GetLeaseDateTime() time.Time`

GetLeaseDateTime returns the LeaseDateTime field if non-nil, zero value otherwise.

### GetLeaseDateTimeOk

`func (o *Address) GetLeaseDateTimeOk() (*time.Time, bool)`

GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDateTime

`func (o *Address) SetLeaseDateTime(v time.Time)`

SetLeaseDateTime sets LeaseDateTime field to given value.

### HasLeaseDateTime

`func (o *Address) HasLeaseDateTime() bool`

HasLeaseDateTime returns a boolean if a field has been set.

### GetLeaseExpirationDateTime

`func (o *Address) GetLeaseExpirationDateTime() time.Time`

GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field if non-nil, zero value otherwise.

### GetLeaseExpirationDateTimeOk

`func (o *Address) GetLeaseExpirationDateTimeOk() (*time.Time, bool)`

GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpirationDateTime

`func (o *Address) SetLeaseExpirationDateTime(v time.Time)`

SetLeaseExpirationDateTime sets LeaseExpirationDateTime field to given value.

### HasLeaseExpirationDateTime

`func (o *Address) HasLeaseExpirationDateTime() bool`

HasLeaseExpirationDateTime returns a boolean if a field has been set.

### GetRemoteId

`func (o *Address) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Address) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Address) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Address) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


