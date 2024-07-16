# GetImportLogs200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the log file. | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the log file, in bytes. | [optional] [readonly] 
**Contents** | Pointer to **string** | The contents of the log file. | [optional] [readonly] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetImportLogs200ResponseDataInner

`func NewGetImportLogs200ResponseDataInner() *GetImportLogs200ResponseDataInner`

NewGetImportLogs200ResponseDataInner instantiates a new GetImportLogs200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImportLogs200ResponseDataInnerWithDefaults

`func NewGetImportLogs200ResponseDataInnerWithDefaults() *GetImportLogs200ResponseDataInner`

NewGetImportLogs200ResponseDataInnerWithDefaults instantiates a new GetImportLogs200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImportLogs200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImportLogs200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImportLogs200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImportLogs200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetImportLogs200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetImportLogs200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetImportLogs200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetImportLogs200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetImportLogs200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetImportLogs200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetImportLogs200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetImportLogs200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *GetImportLogs200ResponseDataInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetImportLogs200ResponseDataInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetImportLogs200ResponseDataInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetImportLogs200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContents

`func (o *GetImportLogs200ResponseDataInner) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *GetImportLogs200ResponseDataInner) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *GetImportLogs200ResponseDataInner) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *GetImportLogs200ResponseDataInner) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetLinks

`func (o *GetImportLogs200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetImportLogs200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetImportLogs200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetImportLogs200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


