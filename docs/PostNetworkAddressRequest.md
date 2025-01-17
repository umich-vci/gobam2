# PostNetworkAddressRequest

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

### NewPostNetworkAddressRequest

`func NewPostNetworkAddressRequest(type_ string, state string, ) *PostNetworkAddressRequest`

NewPostNetworkAddressRequest instantiates a new PostNetworkAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostNetworkAddressRequestWithDefaults

`func NewPostNetworkAddressRequestWithDefaults() *PostNetworkAddressRequest`

NewPostNetworkAddressRequestWithDefaults instantiates a new PostNetworkAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostNetworkAddressRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostNetworkAddressRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostNetworkAddressRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostNetworkAddressRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostNetworkAddressRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostNetworkAddressRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostNetworkAddressRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostNetworkAddressRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostNetworkAddressRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostNetworkAddressRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostNetworkAddressRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostNetworkAddressRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostNetworkAddressRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostNetworkAddressRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostNetworkAddressRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostNetworkAddressRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostNetworkAddressRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostNetworkAddressRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostNetworkAddressRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetState

`func (o *PostNetworkAddressRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostNetworkAddressRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostNetworkAddressRequest) SetState(v string)`

SetState sets State field to given value.


### GetAddress

`func (o *PostNetworkAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostNetworkAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostNetworkAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PostNetworkAddressRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *PostNetworkAddressRequest) GetMacAddress() AddressMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *PostNetworkAddressRequest) GetMacAddressOk() (*AddressMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *PostNetworkAddressRequest) SetMacAddress(v AddressMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *PostNetworkAddressRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDevice

`func (o *PostNetworkAddressRequest) GetDevice() InlinedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PostNetworkAddressRequest) GetDeviceOk() (*InlinedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PostNetworkAddressRequest) SetDevice(v InlinedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PostNetworkAddressRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *PostNetworkAddressRequest) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PostNetworkAddressRequest) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PostNetworkAddressRequest) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PostNetworkAddressRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResourceRecords

`func (o *PostNetworkAddressRequest) GetResourceRecords() []ResourceRecord`

GetResourceRecords returns the ResourceRecords field if non-nil, zero value otherwise.

### GetResourceRecordsOk

`func (o *PostNetworkAddressRequest) GetResourceRecordsOk() (*[]ResourceRecord, bool)`

GetResourceRecordsOk returns a tuple with the ResourceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecords

`func (o *PostNetworkAddressRequest) SetResourceRecords(v []ResourceRecord)`

SetResourceRecords sets ResourceRecords field to given value.

### HasResourceRecords

`func (o *PostNetworkAddressRequest) HasResourceRecords() bool`

HasResourceRecords returns a boolean if a field has been set.

### GetLeaseDateTime

`func (o *PostNetworkAddressRequest) GetLeaseDateTime() time.Time`

GetLeaseDateTime returns the LeaseDateTime field if non-nil, zero value otherwise.

### GetLeaseDateTimeOk

`func (o *PostNetworkAddressRequest) GetLeaseDateTimeOk() (*time.Time, bool)`

GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDateTime

`func (o *PostNetworkAddressRequest) SetLeaseDateTime(v time.Time)`

SetLeaseDateTime sets LeaseDateTime field to given value.

### HasLeaseDateTime

`func (o *PostNetworkAddressRequest) HasLeaseDateTime() bool`

HasLeaseDateTime returns a boolean if a field has been set.

### GetLeaseExpirationDateTime

`func (o *PostNetworkAddressRequest) GetLeaseExpirationDateTime() time.Time`

GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field if non-nil, zero value otherwise.

### GetLeaseExpirationDateTimeOk

`func (o *PostNetworkAddressRequest) GetLeaseExpirationDateTimeOk() (*time.Time, bool)`

GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpirationDateTime

`func (o *PostNetworkAddressRequest) SetLeaseExpirationDateTime(v time.Time)`

SetLeaseExpirationDateTime sets LeaseExpirationDateTime field to given value.

### HasLeaseExpirationDateTime

`func (o *PostNetworkAddressRequest) HasLeaseExpirationDateTime() bool`

HasLeaseExpirationDateTime returns a boolean if a field has been set.

### GetRemoteId

`func (o *PostNetworkAddressRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PostNetworkAddressRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PostNetworkAddressRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PostNetworkAddressRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetReservedUsing

`func (o *PostNetworkAddressRequest) GetReservedUsing() string`

GetReservedUsing returns the ReservedUsing field if non-nil, zero value otherwise.

### GetReservedUsingOk

`func (o *PostNetworkAddressRequest) GetReservedUsingOk() (*string, bool)`

GetReservedUsingOk returns a tuple with the ReservedUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUsing

`func (o *PostNetworkAddressRequest) SetReservedUsing(v string)`

SetReservedUsing sets ReservedUsing field to given value.

### HasReservedUsing

`func (o *PostNetworkAddressRequest) HasReservedUsing() bool`

HasReservedUsing returns a boolean if a field has been set.

### GetIdentityAssociationIdentifier

`func (o *PostNetworkAddressRequest) GetIdentityAssociationIdentifier() string`

GetIdentityAssociationIdentifier returns the IdentityAssociationIdentifier field if non-nil, zero value otherwise.

### GetIdentityAssociationIdentifierOk

`func (o *PostNetworkAddressRequest) GetIdentityAssociationIdentifierOk() (*string, bool)`

GetIdentityAssociationIdentifierOk returns a tuple with the IdentityAssociationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAssociationIdentifier

`func (o *PostNetworkAddressRequest) SetIdentityAssociationIdentifier(v string)`

SetIdentityAssociationIdentifier sets IdentityAssociationIdentifier field to given value.

### HasIdentityAssociationIdentifier

`func (o *PostNetworkAddressRequest) HasIdentityAssociationIdentifier() bool`

HasIdentityAssociationIdentifier returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *PostNetworkAddressRequest) GetClientIdentifier() InlinedDHCPUniqueIdentifier`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *PostNetworkAddressRequest) GetClientIdentifierOk() (*InlinedDHCPUniqueIdentifier, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *PostNetworkAddressRequest) SetClientIdentifier(v InlinedDHCPUniqueIdentifier)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *PostNetworkAddressRequest) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetInterfaceId

`func (o *PostNetworkAddressRequest) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PostNetworkAddressRequest) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PostNetworkAddressRequest) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PostNetworkAddressRequest) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


