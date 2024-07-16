# TagPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the tag. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 

## Methods

### NewTagPutRequestBody

`func NewTagPutRequestBody(name string, ) *TagPutRequestBody`

NewTagPutRequestBody instantiates a new TagPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPutRequestBodyWithDefaults

`func NewTagPutRequestBodyWithDefaults() *TagPutRequestBody`

NewTagPutRequestBodyWithDefaults instantiates a new TagPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TagPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TagPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TagPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *TagPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *TagPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *TagPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *TagPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *TagPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


