# ResourceImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of import | [optional] 
**FileInfo** | Pointer to **interface{}** |  | [optional] 
**FileData** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** | The name of the import | [optional] [readonly] 
**ContentType** | Pointer to **string** | The content type of the import data | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the import, in bytes | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the import | [optional] 
**CreationDateTime** | Pointer to **time.Time** | The date and time the import resource was created | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The date and time the import was started | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The date and time the import completed | [optional] [readonly] 

## Methods

### NewResourceImport

`func NewResourceImport() *ResourceImport`

NewResourceImport instantiates a new ResourceImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceImportWithDefaults

`func NewResourceImportWithDefaults() *ResourceImport`

NewResourceImportWithDefaults instantiates a new ResourceImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceImport) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceImport) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceImport) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceImport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResourceImport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceImport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceImport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceImport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFileInfo

`func (o *ResourceImport) GetFileInfo() interface{}`

GetFileInfo returns the FileInfo field if non-nil, zero value otherwise.

### GetFileInfoOk

`func (o *ResourceImport) GetFileInfoOk() (*interface{}, bool)`

GetFileInfoOk returns a tuple with the FileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfo

`func (o *ResourceImport) SetFileInfo(v interface{})`

SetFileInfo sets FileInfo field to given value.

### HasFileInfo

`func (o *ResourceImport) HasFileInfo() bool`

HasFileInfo returns a boolean if a field has been set.

### SetFileInfoNil

`func (o *ResourceImport) SetFileInfoNil(b bool)`

 SetFileInfoNil sets the value for FileInfo to be an explicit nil

### UnsetFileInfo
`func (o *ResourceImport) UnsetFileInfo()`

UnsetFileInfo ensures that no value is present for FileInfo, not even an explicit nil
### GetFileData

`func (o *ResourceImport) GetFileData() interface{}`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ResourceImport) GetFileDataOk() (*interface{}, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ResourceImport) SetFileData(v interface{})`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ResourceImport) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *ResourceImport) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *ResourceImport) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetName

`func (o *ResourceImport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceImport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceImport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceImport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *ResourceImport) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ResourceImport) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ResourceImport) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ResourceImport) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *ResourceImport) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResourceImport) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResourceImport) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResourceImport) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *ResourceImport) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceImport) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceImport) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResourceImport) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *ResourceImport) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *ResourceImport) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *ResourceImport) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *ResourceImport) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *ResourceImport) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ResourceImport) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ResourceImport) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ResourceImport) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *ResourceImport) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *ResourceImport) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *ResourceImport) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *ResourceImport) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


