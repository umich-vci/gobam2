# GetDevices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Addresses** | Pointer to [**[]DeviceAddressesInner**](DeviceAddressesInner.md) |  | [optional] 
**DeviceType** | Pointer to [**InlinedDeviceType**](InlinedDeviceType.md) |  | [optional] 
**DeviceSubtype** | Pointer to [**InlinedDeviceSubtype**](InlinedDeviceSubtype.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetDevices200ResponseDataInner

`func NewGetDevices200ResponseDataInner() *GetDevices200ResponseDataInner`

NewGetDevices200ResponseDataInner instantiates a new GetDevices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDevices200ResponseDataInnerWithDefaults

`func NewGetDevices200ResponseDataInnerWithDefaults() *GetDevices200ResponseDataInner`

NewGetDevices200ResponseDataInnerWithDefaults instantiates a new GetDevices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDevices200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDevices200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDevices200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDevices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetDevices200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDevices200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDevices200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDevices200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetDevices200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDevices200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDevices200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDevices200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetDevices200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetDevices200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetDevices200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetDevices200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetDevices200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetDevices200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetDevices200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetDevices200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAddresses

`func (o *GetDevices200ResponseDataInner) GetAddresses() []DeviceAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetDevices200ResponseDataInner) GetAddressesOk() (*[]DeviceAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetDevices200ResponseDataInner) SetAddresses(v []DeviceAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetDevices200ResponseDataInner) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetDeviceType

`func (o *GetDevices200ResponseDataInner) GetDeviceType() InlinedDeviceType`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GetDevices200ResponseDataInner) GetDeviceTypeOk() (*InlinedDeviceType, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GetDevices200ResponseDataInner) SetDeviceType(v InlinedDeviceType)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *GetDevices200ResponseDataInner) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceSubtype

`func (o *GetDevices200ResponseDataInner) GetDeviceSubtype() InlinedDeviceSubtype`

GetDeviceSubtype returns the DeviceSubtype field if non-nil, zero value otherwise.

### GetDeviceSubtypeOk

`func (o *GetDevices200ResponseDataInner) GetDeviceSubtypeOk() (*InlinedDeviceSubtype, bool)`

GetDeviceSubtypeOk returns a tuple with the DeviceSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSubtype

`func (o *GetDevices200ResponseDataInner) SetDeviceSubtype(v InlinedDeviceSubtype)`

SetDeviceSubtype sets DeviceSubtype field to given value.

### HasDeviceSubtype

`func (o *GetDevices200ResponseDataInner) HasDeviceSubtype() bool`

HasDeviceSubtype returns a boolean if a field has been set.

### GetLinks

`func (o *GetDevices200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDevices200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDevices200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDevices200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


