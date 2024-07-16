# GetAddresses200Response1DataInner

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
**ReservedUsing** | Pointer to **string** | The type of DHCP reservation. | [optional] 
**IdentityAssociationIdentifier** | Pointer to **string** | The Identity Association Identifier for the network interface of the client system. | [optional] [readonly] 
**ClientIdentifier** | Pointer to [**InlinedDHCPUniqueIdentifier**](InlinedDHCPUniqueIdentifier.md) |  | [optional] 
**InterfaceId** | Pointer to **string** | The ID of the interface from which the request came. | [optional] [readonly] 

## Methods

### NewGetAddresses200Response1DataInner

`func NewGetAddresses200Response1DataInner() *GetAddresses200Response1DataInner`

NewGetAddresses200Response1DataInner instantiates a new GetAddresses200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddresses200Response1DataInnerWithDefaults

`func NewGetAddresses200Response1DataInnerWithDefaults() *GetAddresses200Response1DataInner`

NewGetAddresses200Response1DataInnerWithDefaults instantiates a new GetAddresses200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAddresses200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAddresses200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAddresses200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAddresses200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetAddresses200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAddresses200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAddresses200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAddresses200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetAddresses200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAddresses200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAddresses200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAddresses200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetAddresses200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetAddresses200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetAddresses200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetAddresses200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetAddresses200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetAddresses200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetAddresses200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetAddresses200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetState

`func (o *GetAddresses200Response1DataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAddresses200Response1DataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAddresses200Response1DataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAddresses200Response1DataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAddress

`func (o *GetAddresses200Response1DataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetAddresses200Response1DataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetAddresses200Response1DataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetAddresses200Response1DataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetAddresses200Response1DataInner) GetMacAddress() AddressMacAddress`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetAddresses200Response1DataInner) GetMacAddressOk() (*AddressMacAddress, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetAddresses200Response1DataInner) SetMacAddress(v AddressMacAddress)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetAddresses200Response1DataInner) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDevice

`func (o *GetAddresses200Response1DataInner) GetDevice() InlinedDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetAddresses200Response1DataInner) GetDeviceOk() (*InlinedDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetAddresses200Response1DataInner) SetDevice(v InlinedDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetAddresses200Response1DataInner) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetLocation

`func (o *GetAddresses200Response1DataInner) GetLocation() InlinedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetAddresses200Response1DataInner) GetLocationOk() (*InlinedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetAddresses200Response1DataInner) SetLocation(v InlinedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetAddresses200Response1DataInner) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetResourceRecords

`func (o *GetAddresses200Response1DataInner) GetResourceRecords() []ResourceRecord`

GetResourceRecords returns the ResourceRecords field if non-nil, zero value otherwise.

### GetResourceRecordsOk

`func (o *GetAddresses200Response1DataInner) GetResourceRecordsOk() (*[]ResourceRecord, bool)`

GetResourceRecordsOk returns a tuple with the ResourceRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecords

`func (o *GetAddresses200Response1DataInner) SetResourceRecords(v []ResourceRecord)`

SetResourceRecords sets ResourceRecords field to given value.

### HasResourceRecords

`func (o *GetAddresses200Response1DataInner) HasResourceRecords() bool`

HasResourceRecords returns a boolean if a field has been set.

### GetLeaseDateTime

`func (o *GetAddresses200Response1DataInner) GetLeaseDateTime() time.Time`

GetLeaseDateTime returns the LeaseDateTime field if non-nil, zero value otherwise.

### GetLeaseDateTimeOk

`func (o *GetAddresses200Response1DataInner) GetLeaseDateTimeOk() (*time.Time, bool)`

GetLeaseDateTimeOk returns a tuple with the LeaseDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDateTime

`func (o *GetAddresses200Response1DataInner) SetLeaseDateTime(v time.Time)`

SetLeaseDateTime sets LeaseDateTime field to given value.

### HasLeaseDateTime

`func (o *GetAddresses200Response1DataInner) HasLeaseDateTime() bool`

HasLeaseDateTime returns a boolean if a field has been set.

### GetLeaseExpirationDateTime

`func (o *GetAddresses200Response1DataInner) GetLeaseExpirationDateTime() time.Time`

GetLeaseExpirationDateTime returns the LeaseExpirationDateTime field if non-nil, zero value otherwise.

### GetLeaseExpirationDateTimeOk

`func (o *GetAddresses200Response1DataInner) GetLeaseExpirationDateTimeOk() (*time.Time, bool)`

GetLeaseExpirationDateTimeOk returns a tuple with the LeaseExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseExpirationDateTime

`func (o *GetAddresses200Response1DataInner) SetLeaseExpirationDateTime(v time.Time)`

SetLeaseExpirationDateTime sets LeaseExpirationDateTime field to given value.

### HasLeaseExpirationDateTime

`func (o *GetAddresses200Response1DataInner) HasLeaseExpirationDateTime() bool`

HasLeaseExpirationDateTime returns a boolean if a field has been set.

### GetRemoteId

`func (o *GetAddresses200Response1DataInner) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *GetAddresses200Response1DataInner) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *GetAddresses200Response1DataInner) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *GetAddresses200Response1DataInner) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetReservedUsing

`func (o *GetAddresses200Response1DataInner) GetReservedUsing() string`

GetReservedUsing returns the ReservedUsing field if non-nil, zero value otherwise.

### GetReservedUsingOk

`func (o *GetAddresses200Response1DataInner) GetReservedUsingOk() (*string, bool)`

GetReservedUsingOk returns a tuple with the ReservedUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedUsing

`func (o *GetAddresses200Response1DataInner) SetReservedUsing(v string)`

SetReservedUsing sets ReservedUsing field to given value.

### HasReservedUsing

`func (o *GetAddresses200Response1DataInner) HasReservedUsing() bool`

HasReservedUsing returns a boolean if a field has been set.

### GetIdentityAssociationIdentifier

`func (o *GetAddresses200Response1DataInner) GetIdentityAssociationIdentifier() string`

GetIdentityAssociationIdentifier returns the IdentityAssociationIdentifier field if non-nil, zero value otherwise.

### GetIdentityAssociationIdentifierOk

`func (o *GetAddresses200Response1DataInner) GetIdentityAssociationIdentifierOk() (*string, bool)`

GetIdentityAssociationIdentifierOk returns a tuple with the IdentityAssociationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityAssociationIdentifier

`func (o *GetAddresses200Response1DataInner) SetIdentityAssociationIdentifier(v string)`

SetIdentityAssociationIdentifier sets IdentityAssociationIdentifier field to given value.

### HasIdentityAssociationIdentifier

`func (o *GetAddresses200Response1DataInner) HasIdentityAssociationIdentifier() bool`

HasIdentityAssociationIdentifier returns a boolean if a field has been set.

### GetClientIdentifier

`func (o *GetAddresses200Response1DataInner) GetClientIdentifier() InlinedDHCPUniqueIdentifier`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *GetAddresses200Response1DataInner) GetClientIdentifierOk() (*InlinedDHCPUniqueIdentifier, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *GetAddresses200Response1DataInner) SetClientIdentifier(v InlinedDHCPUniqueIdentifier)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *GetAddresses200Response1DataInner) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### GetInterfaceId

`func (o *GetAddresses200Response1DataInner) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *GetAddresses200Response1DataInner) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *GetAddresses200Response1DataInner) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *GetAddresses200Response1DataInner) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


