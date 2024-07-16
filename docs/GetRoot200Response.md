# GetRoot200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Version** | Pointer to **string** | The Address Manager version. | [optional] 
**Hostname** | Pointer to **string** | The hostname of the Address Manager appliance | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetRoot200Response

`func NewGetRoot200Response() *GetRoot200Response`

NewGetRoot200Response instantiates a new GetRoot200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoot200ResponseWithDefaults

`func NewGetRoot200ResponseWithDefaults() *GetRoot200Response`

NewGetRoot200ResponseWithDefaults instantiates a new GetRoot200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRoot200Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRoot200Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRoot200Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetRoot200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetRoot200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRoot200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRoot200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRoot200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetRoot200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRoot200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRoot200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRoot200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetRoot200Response) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetRoot200Response) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetRoot200Response) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetRoot200Response) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetVersion

`func (o *GetRoot200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRoot200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRoot200Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetRoot200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHostname

`func (o *GetRoot200Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetRoot200Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetRoot200Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetRoot200Response) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLinks

`func (o *GetRoot200Response) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetRoot200Response) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetRoot200Response) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetRoot200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


