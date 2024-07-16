# GetFiles200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Version** | Pointer to **string** | The version of the TFTP file. | [optional] 
**Description** | Pointer to **string** | The description of the TFTP file. | [optional] 
**Size** | Pointer to **int64** | The size of the TFTP file, in bytes. | [optional] [readonly] 
**Md5Checksum** | Pointer to **string** | The MD5 checksum of the TFTP file. | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetFiles200ResponseDataInner

`func NewGetFiles200ResponseDataInner() *GetFiles200ResponseDataInner`

NewGetFiles200ResponseDataInner instantiates a new GetFiles200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiles200ResponseDataInnerWithDefaults

`func NewGetFiles200ResponseDataInnerWithDefaults() *GetFiles200ResponseDataInner`

NewGetFiles200ResponseDataInnerWithDefaults instantiates a new GetFiles200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetFiles200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFiles200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFiles200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetFiles200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetFiles200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetFiles200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetFiles200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetFiles200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetFiles200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFiles200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFiles200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFiles200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetFiles200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetFiles200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetFiles200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetFiles200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetVersion

`func (o *GetFiles200ResponseDataInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetFiles200ResponseDataInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetFiles200ResponseDataInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetFiles200ResponseDataInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *GetFiles200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetFiles200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetFiles200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetFiles200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *GetFiles200ResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetFiles200ResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetFiles200ResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetFiles200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5Checksum

`func (o *GetFiles200ResponseDataInner) GetMd5Checksum() string`

GetMd5Checksum returns the Md5Checksum field if non-nil, zero value otherwise.

### GetMd5ChecksumOk

`func (o *GetFiles200ResponseDataInner) GetMd5ChecksumOk() (*string, bool)`

GetMd5ChecksumOk returns a tuple with the Md5Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Checksum

`func (o *GetFiles200ResponseDataInner) SetMd5Checksum(v string)`

SetMd5Checksum sets Md5Checksum field to given value.

### HasMd5Checksum

`func (o *GetFiles200ResponseDataInner) HasMd5Checksum() bool`

HasMd5Checksum returns a boolean if a field has been set.

### GetLinks

`func (o *GetFiles200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetFiles200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetFiles200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetFiles200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


