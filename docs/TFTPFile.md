# TFTPFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Version** | Pointer to **string** | The version of the TFTP file. | [optional] 
**Description** | Pointer to **string** | The description of the TFTP file. | [optional] 
**Size** | Pointer to **int64** | The size of the TFTP file, in bytes. | [optional] [readonly] 
**Md5Checksum** | Pointer to **string** | The MD5 checksum of the TFTP file. | [optional] [readonly] 

## Methods

### NewTFTPFile

`func NewTFTPFile() *TFTPFile`

NewTFTPFile instantiates a new TFTPFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTFTPFileWithDefaults

`func NewTFTPFileWithDefaults() *TFTPFile`

NewTFTPFileWithDefaults instantiates a new TFTPFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TFTPFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TFTPFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TFTPFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TFTPFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *TFTPFile) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TFTPFile) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TFTPFile) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TFTPFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *TFTPFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TFTPFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TFTPFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TFTPFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *TFTPFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TFTPFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TFTPFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TFTPFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMd5Checksum

`func (o *TFTPFile) GetMd5Checksum() string`

GetMd5Checksum returns the Md5Checksum field if non-nil, zero value otherwise.

### GetMd5ChecksumOk

`func (o *TFTPFile) GetMd5ChecksumOk() (*string, bool)`

GetMd5ChecksumOk returns a tuple with the Md5Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Checksum

`func (o *TFTPFile) SetMd5Checksum(v string)`

SetMd5Checksum sets Md5Checksum field to given value.

### HasMd5Checksum

`func (o *TFTPFile) HasMd5Checksum() bool`

HasMd5Checksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


