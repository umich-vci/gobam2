# GetMacAddresses200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Address** | Pointer to **string** | The MAC address. | [optional] 
**MacPool** | Pointer to [**InlinedMacPool**](InlinedMacPool.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetMacAddresses200ResponseDataInner

`func NewGetMacAddresses200ResponseDataInner() *GetMacAddresses200ResponseDataInner`

NewGetMacAddresses200ResponseDataInner instantiates a new GetMacAddresses200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMacAddresses200ResponseDataInnerWithDefaults

`func NewGetMacAddresses200ResponseDataInnerWithDefaults() *GetMacAddresses200ResponseDataInner`

NewGetMacAddresses200ResponseDataInnerWithDefaults instantiates a new GetMacAddresses200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMacAddresses200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMacAddresses200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMacAddresses200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetMacAddresses200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMacAddresses200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMacAddresses200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMacAddresses200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMacAddresses200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetMacAddresses200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMacAddresses200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMacAddresses200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMacAddresses200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetMacAddresses200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetMacAddresses200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetMacAddresses200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetMacAddresses200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetMacAddresses200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetMacAddresses200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetMacAddresses200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetMacAddresses200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetAddress

`func (o *GetMacAddresses200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetMacAddresses200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetMacAddresses200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetMacAddresses200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMacPool

`func (o *GetMacAddresses200ResponseDataInner) GetMacPool() InlinedMacPool`

GetMacPool returns the MacPool field if non-nil, zero value otherwise.

### GetMacPoolOk

`func (o *GetMacAddresses200ResponseDataInner) GetMacPoolOk() (*InlinedMacPool, bool)`

GetMacPoolOk returns a tuple with the MacPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPool

`func (o *GetMacAddresses200ResponseDataInner) SetMacPool(v InlinedMacPool)`

SetMacPool sets MacPool field to given value.

### HasMacPool

`func (o *GetMacAddresses200ResponseDataInner) HasMacPool() bool`

HasMacPool returns a boolean if a field has been set.

### GetLinks

`func (o *GetMacAddresses200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetMacAddresses200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetMacAddresses200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetMacAddresses200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


