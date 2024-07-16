# ResponsePolicyItemImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The type of import | [optional] [readonly] 
**FileInfo** | Pointer to **interface{}** |  | [optional] 
**FileData** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** | The name of the import | [optional] [readonly] 
**ContentType** | Pointer to **string** | The content type of the import data | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the import, in bytes | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the import | [optional] [readonly] 
**CreationDateTime** | Pointer to **time.Time** | The date and time the import resource was created | [optional] [readonly] 
**StartDateTime** | Pointer to **time.Time** | The date and time the import was started | [optional] [readonly] 
**CompletionDateTime** | Pointer to **time.Time** | The date and time the import completed | [optional] [readonly] 

## Methods

### NewResponsePolicyItemImport

`func NewResponsePolicyItemImport() *ResponsePolicyItemImport`

NewResponsePolicyItemImport instantiates a new ResponsePolicyItemImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePolicyItemImportWithDefaults

`func NewResponsePolicyItemImportWithDefaults() *ResponsePolicyItemImport`

NewResponsePolicyItemImportWithDefaults instantiates a new ResponsePolicyItemImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsePolicyItemImport) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsePolicyItemImport) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsePolicyItemImport) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsePolicyItemImport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ResponsePolicyItemImport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePolicyItemImport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePolicyItemImport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePolicyItemImport) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFileInfo

`func (o *ResponsePolicyItemImport) GetFileInfo() interface{}`

GetFileInfo returns the FileInfo field if non-nil, zero value otherwise.

### GetFileInfoOk

`func (o *ResponsePolicyItemImport) GetFileInfoOk() (*interface{}, bool)`

GetFileInfoOk returns a tuple with the FileInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfo

`func (o *ResponsePolicyItemImport) SetFileInfo(v interface{})`

SetFileInfo sets FileInfo field to given value.

### HasFileInfo

`func (o *ResponsePolicyItemImport) HasFileInfo() bool`

HasFileInfo returns a boolean if a field has been set.

### SetFileInfoNil

`func (o *ResponsePolicyItemImport) SetFileInfoNil(b bool)`

 SetFileInfoNil sets the value for FileInfo to be an explicit nil

### UnsetFileInfo
`func (o *ResponsePolicyItemImport) UnsetFileInfo()`

UnsetFileInfo ensures that no value is present for FileInfo, not even an explicit nil
### GetFileData

`func (o *ResponsePolicyItemImport) GetFileData() interface{}`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ResponsePolicyItemImport) GetFileDataOk() (*interface{}, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ResponsePolicyItemImport) SetFileData(v interface{})`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ResponsePolicyItemImport) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *ResponsePolicyItemImport) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *ResponsePolicyItemImport) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetName

`func (o *ResponsePolicyItemImport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePolicyItemImport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePolicyItemImport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsePolicyItemImport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *ResponsePolicyItemImport) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ResponsePolicyItemImport) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ResponsePolicyItemImport) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ResponsePolicyItemImport) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *ResponsePolicyItemImport) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResponsePolicyItemImport) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResponsePolicyItemImport) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ResponsePolicyItemImport) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetState

`func (o *ResponsePolicyItemImport) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResponsePolicyItemImport) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResponsePolicyItemImport) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResponsePolicyItemImport) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *ResponsePolicyItemImport) GetCreationDateTime() time.Time`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *ResponsePolicyItemImport) GetCreationDateTimeOk() (*time.Time, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *ResponsePolicyItemImport) SetCreationDateTime(v time.Time)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *ResponsePolicyItemImport) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetStartDateTime

`func (o *ResponsePolicyItemImport) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ResponsePolicyItemImport) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ResponsePolicyItemImport) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ResponsePolicyItemImport) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### GetCompletionDateTime

`func (o *ResponsePolicyItemImport) GetCompletionDateTime() time.Time`

GetCompletionDateTime returns the CompletionDateTime field if non-nil, zero value otherwise.

### GetCompletionDateTimeOk

`func (o *ResponsePolicyItemImport) GetCompletionDateTimeOk() (*time.Time, bool)`

GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDateTime

`func (o *ResponsePolicyItemImport) SetCompletionDateTime(v time.Time)`

SetCompletionDateTime sets CompletionDateTime field to given value.

### HasCompletionDateTime

`func (o *ResponsePolicyItemImport) HasCompletionDateTime() bool`

HasCompletionDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


