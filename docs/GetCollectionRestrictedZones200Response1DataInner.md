# GetCollectionRestrictedZones200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the view. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**DeviceRegistrationEnabled** | Pointer to **bool** | Indicates whether DNS redirection has been enabled for use with the BlueCat Device Registration Portal. | [optional] 
**DeviceRegistrationPortalAddress** | Pointer to **string** | The IPv4 address of the BlueCat Device Registration Portal. | [optional] 

## Methods

### NewGetCollectionRestrictedZones200Response1DataInner

`func NewGetCollectionRestrictedZones200Response1DataInner() *GetCollectionRestrictedZones200Response1DataInner`

NewGetCollectionRestrictedZones200Response1DataInner instantiates a new GetCollectionRestrictedZones200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionRestrictedZones200Response1DataInnerWithDefaults

`func NewGetCollectionRestrictedZones200Response1DataInnerWithDefaults() *GetCollectionRestrictedZones200Response1DataInner`

NewGetCollectionRestrictedZones200Response1DataInnerWithDefaults instantiates a new GetCollectionRestrictedZones200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeviceRegistrationEnabled

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetDeviceRegistrationEnabled() bool`

GetDeviceRegistrationEnabled returns the DeviceRegistrationEnabled field if non-nil, zero value otherwise.

### GetDeviceRegistrationEnabledOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetDeviceRegistrationEnabledOk() (*bool, bool)`

GetDeviceRegistrationEnabledOk returns a tuple with the DeviceRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationEnabled

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetDeviceRegistrationEnabled(v bool)`

SetDeviceRegistrationEnabled sets DeviceRegistrationEnabled field to given value.

### HasDeviceRegistrationEnabled

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasDeviceRegistrationEnabled() bool`

HasDeviceRegistrationEnabled returns a boolean if a field has been set.

### GetDeviceRegistrationPortalAddress

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetDeviceRegistrationPortalAddress() string`

GetDeviceRegistrationPortalAddress returns the DeviceRegistrationPortalAddress field if non-nil, zero value otherwise.

### GetDeviceRegistrationPortalAddressOk

`func (o *GetCollectionRestrictedZones200Response1DataInner) GetDeviceRegistrationPortalAddressOk() (*string, bool)`

GetDeviceRegistrationPortalAddressOk returns a tuple with the DeviceRegistrationPortalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationPortalAddress

`func (o *GetCollectionRestrictedZones200Response1DataInner) SetDeviceRegistrationPortalAddress(v string)`

SetDeviceRegistrationPortalAddress sets DeviceRegistrationPortalAddress field to given value.

### HasDeviceRegistrationPortalAddress

`func (o *GetCollectionRestrictedZones200Response1DataInner) HasDeviceRegistrationPortalAddress() bool`

HasDeviceRegistrationPortalAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


