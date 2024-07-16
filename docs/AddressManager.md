# AddressManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Version** | Pointer to **string** | The Address Manager version. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the Address Manager appliance | [optional] 

## Methods

### NewAddressManager

`func NewAddressManager() *AddressManager`

NewAddressManager instantiates a new AddressManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressManagerWithDefaults

`func NewAddressManagerWithDefaults() *AddressManager`

NewAddressManagerWithDefaults instantiates a new AddressManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddressManager) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressManager) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressManager) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddressManager) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddressManager) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressManager) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressManager) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressManager) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AddressManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressManager) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddressManager) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *AddressManager) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *AddressManager) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *AddressManager) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *AddressManager) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetVersion

`func (o *AddressManager) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddressManager) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddressManager) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddressManager) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHostname

`func (o *AddressManager) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddressManager) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddressManager) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AddressManager) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


