# ViewPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the view. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**DeviceRegistrationEnabled** | **bool** | Indicates whether DNS redirection has been enabled for use with the BlueCat Device Registration Portal. | 
**DeviceRegistrationPortalAddress** | Pointer to **string** | The IPv4 address of the BlueCat Device Registration Portal. | [optional] 

## Methods

### NewViewPutRequestBody

`func NewViewPutRequestBody(deviceRegistrationEnabled bool, ) *ViewPutRequestBody`

NewViewPutRequestBody instantiates a new ViewPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPutRequestBodyWithDefaults

`func NewViewPutRequestBodyWithDefaults() *ViewPutRequestBody`

NewViewPutRequestBodyWithDefaults instantiates a new ViewPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ViewPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ViewPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ViewPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewPutRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ViewPutRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *ViewPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ViewPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ViewPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ViewPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ViewPutRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ViewPutRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ViewPutRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ViewPutRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeviceRegistrationEnabled

`func (o *ViewPutRequestBody) GetDeviceRegistrationEnabled() bool`

GetDeviceRegistrationEnabled returns the DeviceRegistrationEnabled field if non-nil, zero value otherwise.

### GetDeviceRegistrationEnabledOk

`func (o *ViewPutRequestBody) GetDeviceRegistrationEnabledOk() (*bool, bool)`

GetDeviceRegistrationEnabledOk returns a tuple with the DeviceRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationEnabled

`func (o *ViewPutRequestBody) SetDeviceRegistrationEnabled(v bool)`

SetDeviceRegistrationEnabled sets DeviceRegistrationEnabled field to given value.


### GetDeviceRegistrationPortalAddress

`func (o *ViewPutRequestBody) GetDeviceRegistrationPortalAddress() string`

GetDeviceRegistrationPortalAddress returns the DeviceRegistrationPortalAddress field if non-nil, zero value otherwise.

### GetDeviceRegistrationPortalAddressOk

`func (o *ViewPutRequestBody) GetDeviceRegistrationPortalAddressOk() (*string, bool)`

GetDeviceRegistrationPortalAddressOk returns a tuple with the DeviceRegistrationPortalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationPortalAddress

`func (o *ViewPutRequestBody) SetDeviceRegistrationPortalAddress(v string)`

SetDeviceRegistrationPortalAddress sets DeviceRegistrationPortalAddress field to given value.

### HasDeviceRegistrationPortalAddress

`func (o *ViewPutRequestBody) HasDeviceRegistrationPortalAddress() bool`

HasDeviceRegistrationPortalAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


