# GetInterfaces200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Server** | Pointer to [**AbstractServer**](AbstractServer.md) |  | [optional] [readonly] 
**SecondaryAddress** | Pointer to **string** | The secondary IP address of the interface. | [optional] 
**PrimaryAddress** | Pointer to **string** | The primary IP address of the published interface. | [optional] 

## Methods

### NewGetInterfaces200Response1DataInner

`func NewGetInterfaces200Response1DataInner() *GetInterfaces200Response1DataInner`

NewGetInterfaces200Response1DataInner instantiates a new GetInterfaces200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInterfaces200Response1DataInnerWithDefaults

`func NewGetInterfaces200Response1DataInnerWithDefaults() *GetInterfaces200Response1DataInner`

NewGetInterfaces200Response1DataInnerWithDefaults instantiates a new GetInterfaces200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInterfaces200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInterfaces200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInterfaces200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInterfaces200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetInterfaces200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetInterfaces200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetInterfaces200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetInterfaces200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetInterfaces200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInterfaces200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInterfaces200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInterfaces200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetInterfaces200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetInterfaces200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetInterfaces200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetInterfaces200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetInterfaces200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetInterfaces200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetInterfaces200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetInterfaces200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetServer

`func (o *GetInterfaces200Response1DataInner) GetServer() AbstractServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetInterfaces200Response1DataInner) GetServerOk() (*AbstractServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetInterfaces200Response1DataInner) SetServer(v AbstractServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetInterfaces200Response1DataInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSecondaryAddress

`func (o *GetInterfaces200Response1DataInner) GetSecondaryAddress() string`

GetSecondaryAddress returns the SecondaryAddress field if non-nil, zero value otherwise.

### GetSecondaryAddressOk

`func (o *GetInterfaces200Response1DataInner) GetSecondaryAddressOk() (*string, bool)`

GetSecondaryAddressOk returns a tuple with the SecondaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAddress

`func (o *GetInterfaces200Response1DataInner) SetSecondaryAddress(v string)`

SetSecondaryAddress sets SecondaryAddress field to given value.

### HasSecondaryAddress

`func (o *GetInterfaces200Response1DataInner) HasSecondaryAddress() bool`

HasSecondaryAddress returns a boolean if a field has been set.

### GetPrimaryAddress

`func (o *GetInterfaces200Response1DataInner) GetPrimaryAddress() string`

GetPrimaryAddress returns the PrimaryAddress field if non-nil, zero value otherwise.

### GetPrimaryAddressOk

`func (o *GetInterfaces200Response1DataInner) GetPrimaryAddressOk() (*string, bool)`

GetPrimaryAddressOk returns a tuple with the PrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddress

`func (o *GetInterfaces200Response1DataInner) SetPrimaryAddress(v string)`

SetPrimaryAddress sets PrimaryAddress field to given value.

### HasPrimaryAddress

`func (o *GetInterfaces200Response1DataInner) HasPrimaryAddress() bool`

HasPrimaryAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


