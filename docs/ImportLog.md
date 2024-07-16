# ImportLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the log file. | [optional] [readonly] 
**Size** | Pointer to **int64** | The size of the log file, in bytes. | [optional] [readonly] 
**Contents** | Pointer to **string** | The contents of the log file. | [optional] [readonly] 

## Methods

### NewImportLog

`func NewImportLog() *ImportLog`

NewImportLog instantiates a new ImportLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportLogWithDefaults

`func NewImportLogWithDefaults() *ImportLog`

NewImportLogWithDefaults instantiates a new ImportLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImportLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ImportLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ImportLog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportLog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportLog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ImportLog) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ImportLog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportLog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportLog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportLog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ImportLog) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImportLog) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImportLog) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImportLog) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetContents

`func (o *ImportLog) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ImportLog) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ImportLog) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ImportLog) HasContents() bool`

HasContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


