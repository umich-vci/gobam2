# ViewPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the view. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**DeviceRegistrationEnabled** | Pointer to **bool** | Indicates whether DNS redirection has been enabled for use with the BlueCat Device Registration Portal. | [optional] 
**DeviceRegistrationPortalAddress** | Pointer to **string** | The IPv4 address of the BlueCat Device Registration Portal. | [optional] 

## Methods

### NewViewPostRequestBody

`func NewViewPostRequestBody(name string, ) *ViewPostRequestBody`

NewViewPostRequestBody instantiates a new ViewPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewPostRequestBodyWithDefaults

`func NewViewPostRequestBodyWithDefaults() *ViewPostRequestBody`

NewViewPostRequestBodyWithDefaults instantiates a new ViewPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ViewPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ViewPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ViewPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ViewPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ViewPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ViewPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ViewPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *ViewPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *ViewPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *ViewPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *ViewPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *ViewPostRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ViewPostRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ViewPostRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ViewPostRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeviceRegistrationEnabled

`func (o *ViewPostRequestBody) GetDeviceRegistrationEnabled() bool`

GetDeviceRegistrationEnabled returns the DeviceRegistrationEnabled field if non-nil, zero value otherwise.

### GetDeviceRegistrationEnabledOk

`func (o *ViewPostRequestBody) GetDeviceRegistrationEnabledOk() (*bool, bool)`

GetDeviceRegistrationEnabledOk returns a tuple with the DeviceRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationEnabled

`func (o *ViewPostRequestBody) SetDeviceRegistrationEnabled(v bool)`

SetDeviceRegistrationEnabled sets DeviceRegistrationEnabled field to given value.

### HasDeviceRegistrationEnabled

`func (o *ViewPostRequestBody) HasDeviceRegistrationEnabled() bool`

HasDeviceRegistrationEnabled returns a boolean if a field has been set.

### GetDeviceRegistrationPortalAddress

`func (o *ViewPostRequestBody) GetDeviceRegistrationPortalAddress() string`

GetDeviceRegistrationPortalAddress returns the DeviceRegistrationPortalAddress field if non-nil, zero value otherwise.

### GetDeviceRegistrationPortalAddressOk

`func (o *ViewPostRequestBody) GetDeviceRegistrationPortalAddressOk() (*string, bool)`

GetDeviceRegistrationPortalAddressOk returns a tuple with the DeviceRegistrationPortalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationPortalAddress

`func (o *ViewPostRequestBody) SetDeviceRegistrationPortalAddress(v string)`

SetDeviceRegistrationPortalAddress sets DeviceRegistrationPortalAddress field to given value.

### HasDeviceRegistrationPortalAddress

`func (o *ViewPostRequestBody) HasDeviceRegistrationPortalAddress() bool`

HasDeviceRegistrationPortalAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


