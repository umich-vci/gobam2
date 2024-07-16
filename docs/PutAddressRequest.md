# PutAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The type of IP address. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**State** | **string** | The state of the IP address. | 
**Address** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to [**AddressMacAddress**](AddressMacAddress.md) |  | [optional] 
**Device** | Pointer to [**InlinedDevice**](InlinedDevice.md) |  | [optional] 
**Location** | Pointer to [**InlinedLocation**](InlinedLocation.md) |  | [optional] 
**ResourceRecords** | Pointer to [**[]ResourceRecord**](ResourceRecord.md) |  | [optional] 
**LeaseDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address was offered. | [optional] [readonly] 
**LeaseExpirationDateTime** | Pointer to **time.Time** | The timestamp detailing when the lease for the IP address expires. | [optional] [readonly] 
**RemoteId** | Pointer to **string** | The ID of the relay agent. | [optional] [readonly] 
**ReservedUsing** | Pointer to **string** | The type of DHCP reservation. | [optional] 
**IdentityAssociationIdentifier** | Pointer to **string** | The Identity Association Identifier for the network interface of the client system. | [optional] [readonly] 
**ClientIdentifier** | Pointer to [**InlinedDHCPUniqueIdentifier**](InlinedDHCPUniqueIdentifier.md) |  | [optional] 
**InterfaceId** | Pointer to **string** | The ID of the interface from which the request came. | [optional] [readonly] 

## Methods

### NewPutAddressRequest

`func NewPutAddressRequest(type_ string, state string, ) *PutAddressRequest`

NewPutAddressRequest instantiates a new PutAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutAddressRequestWithDefaults

`func NewPutAddressRequestWithDefaults() *PutAddressRequest`

NewPutAddressRequestWithDefaults instantiates a new PutAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutAddressRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutAddressRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutAddressRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutAddressRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutAddressRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutAddressRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutAddressRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PutAddressRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutAddressRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutAddressRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutAddressRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PutAddressRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutAddressRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutAddressRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutAddressRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PutAddressRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PutAddressRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PutAddressRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PutAddressRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetState

`func (o *PutAddressRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PutAddressRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PutAddressRequest) SetState(v string)`

SetState sets State field to given value.


### GetAddress

`func (o *PutAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PutAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PutAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PutAddressRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *PutAddressRequest) GetMacAddress() AddressMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PutAddressRequest) GetMacAddressOk() (*AddressMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PutAddressRequest) SetMacAddress(v AddressMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PutAddressRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDevice

`func (o *PutAddressRequest) GetDevice() InlinedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PutAddressRequest) GetDeviceOk() (*InlinedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PutAddressRequest) SetDevice(v InlinedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PutAddressRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *PutAddressRequest) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PutAddressRequest) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PutAddressRequest) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PutAddressRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResourceRecords

`func (o *PutAddressRequest) GetResourceRecords() []ResourceRecord`

GetResourceRecords returns the ResourceRecords field if non-nil, zero value otherwise.

### GetResourceRecordsOk

`func (o *PutAddressRequest) GetResourceRecordsOk() (*[]ResourceRecord, bool)`

GetResourceRecordsOk returns a tuple with the ResourceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecords

`func (o *PutAddressRequest) SetResourceRecords(v []ResourceRecord)`

SetResourceRecords sets ResourceRecords field to given value.

### HasResourceRecords

`func (o *PutAddressRequest) HasResourceRecords() bool`

HasResourceRecords returns a boolean if a field has been set.

### GetLeaseDateTime

`func (o *PutAddressRequest) GetLeaseDateTime() time.Time`

GetLeaseDateTime returns the LeaseDateTime field if non-nil, zero value otherwise.

### GetLeaseDateTimeOk

`func (o *PutAddressRequest) GetLeaseDateTimeOk() (*time.Time, bool)`

GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDateTime

`func (o *PutAddressRequest) SetLeaseDateTime(v time.Time)`

SetLeaseDateTime sets LeaseDateTime field to given value.

### HasLeaseDateTime

`func (o *PutAddressRequest) HasLeaseDateTime() bool`

HasLeaseDateTime returns a boolean if a field has been set.

### GetLeaseExpirationDateTime

`func (o *PutAddressRequest) GetLeaseExpirationDateTime() time.Time`

GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field if non-nil, zero value otherwise.

### GetLeaseExpirationDateTimeOk

`func (o *PutAddressRequest) GetLeaseExpirationDateTimeOk() (*time.Time, bool)`

GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpirationDateTime

`func (o *PutAddressRequest) SetLeaseExpirationDateTime(v time.Time)`

SetLeaseExpirationDateTime sets LeaseExpirationDateTime field to given value.

### HasLeaseExpirationDateTime

`func (o *PutAddressRequest) HasLeaseExpirationDateTime() bool`

HasLeaseExpirationDateTime returns a boolean if a field has been set.

### GetRemoteId

`func (o *PutAddressRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PutAddressRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PutAddressRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PutAddressRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetReservedUsing

`func (o *PutAddressRequest) GetReservedUsing() string`

GetReservedUsing returns the ReservedUsing field if non-nil, zero value otherwise.

### GetReservedUsingOk

`func (o *PutAddressRequest) GetReservedUsingOk() (*string, bool)`

GetReservedUsingOk returns a tuple with the ReservedUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUsing

`func (o *PutAddressRequest) SetReservedUsing(v string)`

SetReservedUsing sets ReservedUsing field to given value.

### HasReservedUsing

`func (o *PutAddressRequest) HasReservedUsing() bool`

HasReservedUsing returns a boolean if a field has been set.

### GetIdentityAssociationIdentifier

`func (o *PutAddressRequest) GetIdentityAssociationIdentifier() string`

GetIdentityAssociationIdentifier returns the IdentityAssociationIdentifier field if non-nil, zero value otherwise.

### GetIdentityAssociationIdentifierOk

`func (o *PutAddressRequest) GetIdentityAssociationIdentifierOk() (*string, bool)`

GetIdentityAssociationIdentifierOk returns a tuple with the IdentityAssociationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAssociationIdentifier

`func (o *PutAddressRequest) SetIdentityAssociationIdentifier(v string)`

SetIdentityAssociationIdentifier sets IdentityAssociationIdentifier field to given value.

### HasIdentityAssociationIdentifier

`func (o *PutAddressRequest) HasIdentityAssociationIdentifier() bool`

HasIdentityAssociationIdentifier returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *PutAddressRequest) GetClientIdentifier() InlinedDHCPUniqueIdentifier`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *PutAddressRequest) GetClientIdentifierOk() (*InlinedDHCPUniqueIdentifier, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *PutAddressRequest) SetClientIdentifier(v InlinedDHCPUniqueIdentifier)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *PutAddressRequest) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetInterfaceId

`func (o *PutAddressRequest) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PutAddressRequest) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PutAddressRequest) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PutAddressRequest) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


