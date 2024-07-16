# LinkedHostRecordAllOfViews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The resource identifier. | 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the view. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**DeviceRegistrationEnabled** | Pointer to **bool** | Indicates whether DNS redirection has been enabled for use with the BlueCat Device Registration Portal. | [optional] 
**DeviceRegistrationPortalAddress** | Pointer to **string** | The IPv4 address of the BlueCat Device Registration Portal. | [optional] 

## Methods

### NewLinkedHostRecordAllOfViews

`func NewLinkedHostRecordAllOfViews(id int64, ) *LinkedHostRecordAllOfViews`

NewLinkedHostRecordAllOfViews instantiates a new LinkedHostRecordAllOfViews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedHostRecordAllOfViewsWithDefaults

`func NewLinkedHostRecordAllOfViewsWithDefaults() *LinkedHostRecordAllOfViews`

NewLinkedHostRecordAllOfViewsWithDefaults instantiates a new LinkedHostRecordAllOfViews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedHostRecordAllOfViews) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedHostRecordAllOfViews) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedHostRecordAllOfViews) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *LinkedHostRecordAllOfViews) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedHostRecordAllOfViews) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedHostRecordAllOfViews) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedHostRecordAllOfViews) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *LinkedHostRecordAllOfViews) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkedHostRecordAllOfViews) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkedHostRecordAllOfViews) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinkedHostRecordAllOfViews) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *LinkedHostRecordAllOfViews) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *LinkedHostRecordAllOfViews) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *LinkedHostRecordAllOfViews) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *LinkedHostRecordAllOfViews) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *LinkedHostRecordAllOfViews) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *LinkedHostRecordAllOfViews) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *LinkedHostRecordAllOfViews) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *LinkedHostRecordAllOfViews) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeviceRegistrationEnabled

`func (o *LinkedHostRecordAllOfViews) GetDeviceRegistrationEnabled() bool`

GetDeviceRegistrationEnabled returns the DeviceRegistrationEnabled field if non-nil, zero value otherwise.

### GetDeviceRegistrationEnabledOk

`func (o *LinkedHostRecordAllOfViews) GetDeviceRegistrationEnabledOk() (*bool, bool)`

GetDeviceRegistrationEnabledOk returns a tuple with the DeviceRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationEnabled

`func (o *LinkedHostRecordAllOfViews) SetDeviceRegistrationEnabled(v bool)`

SetDeviceRegistrationEnabled sets DeviceRegistrationEnabled field to given value.

### HasDeviceRegistrationEnabled

`func (o *LinkedHostRecordAllOfViews) HasDeviceRegistrationEnabled() bool`

HasDeviceRegistrationEnabled returns a boolean if a field has been set.

### GetDeviceRegistrationPortalAddress

`func (o *LinkedHostRecordAllOfViews) GetDeviceRegistrationPortalAddress() string`

GetDeviceRegistrationPortalAddress returns the DeviceRegistrationPortalAddress field if non-nil, zero value otherwise.

### GetDeviceRegistrationPortalAddressOk

`func (o *LinkedHostRecordAllOfViews) GetDeviceRegistrationPortalAddressOk() (*string, bool)`

GetDeviceRegistrationPortalAddressOk returns a tuple with the DeviceRegistrationPortalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationPortalAddress

`func (o *LinkedHostRecordAllOfViews) SetDeviceRegistrationPortalAddress(v string)`

SetDeviceRegistrationPortalAddress sets DeviceRegistrationPortalAddress field to given value.

### HasDeviceRegistrationPortalAddress

`func (o *LinkedHostRecordAllOfViews) HasDeviceRegistrationPortalAddress() bool`

HasDeviceRegistrationPortalAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


