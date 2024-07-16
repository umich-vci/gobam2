# AbstractZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**View** | Pointer to [**View**](View.md) |  | [optional] [readonly] 

## Methods

### NewAbstractZone

`func NewAbstractZone() *AbstractZone`

NewAbstractZone instantiates a new AbstractZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractZoneWithDefaults

`func NewAbstractZoneWithDefaults() *AbstractZone`

NewAbstractZoneWithDefaults instantiates a new AbstractZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AbstractZone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractZone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractZone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AbstractZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbstractZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbstractZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AbstractZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *AbstractZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *AbstractZone) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *AbstractZone) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *AbstractZone) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *AbstractZone) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *AbstractZone) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AbstractZone) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AbstractZone) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AbstractZone) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetView

`func (o *AbstractZone) GetView() View`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *AbstractZone) GetViewOk() (*View, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *AbstractZone) SetView(v View)`

SetView sets View field to given value.

### HasView

`func (o *AbstractZone) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


