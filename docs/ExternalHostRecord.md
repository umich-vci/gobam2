# ExternalHostRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 

## Methods

### NewExternalHostRecord

`func NewExternalHostRecord() *ExternalHostRecord`

NewExternalHostRecord instantiates a new ExternalHostRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalHostRecordWithDefaults

`func NewExternalHostRecordWithDefaults() *ExternalHostRecord`

NewExternalHostRecordWithDefaults instantiates a new ExternalHostRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalHostRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalHostRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalHostRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalHostRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ExternalHostRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalHostRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalHostRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalHostRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetView

`func (o *ExternalHostRecord) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ExternalHostRecord) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ExternalHostRecord) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *ExternalHostRecord) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


