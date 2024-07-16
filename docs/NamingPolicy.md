# NamingPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the naming policy. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**PolicyValues** | Pointer to [**[]NamingPolicyValue**](NamingPolicyValue.md) |  | [optional] 
**PolicyRestrictions** | Pointer to [**[]NamingPolicyRestriction**](NamingPolicyRestriction.md) |  | [optional] 

## Methods

### NewNamingPolicy

`func NewNamingPolicy() *NamingPolicy`

NewNamingPolicy instantiates a new NamingPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyWithDefaults

`func NewNamingPolicyWithDefaults() *NamingPolicy`

NewNamingPolicyWithDefaults instantiates a new NamingPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NamingPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NamingPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicy) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NamingPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamingPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamingPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NamingPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *NamingPolicy) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *NamingPolicy) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *NamingPolicy) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *NamingPolicy) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPolicyValues

`func (o *NamingPolicy) GetPolicyValues() []NamingPolicyValue`

GetPolicyValues returns the PolicyValues field if non-nil, zero value otherwise.

### GetPolicyValuesOk

`func (o *NamingPolicy) GetPolicyValuesOk() (*[]NamingPolicyValue, bool)`

GetPolicyValuesOk returns a tuple with the PolicyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyValues

`func (o *NamingPolicy) SetPolicyValues(v []NamingPolicyValue)`

SetPolicyValues sets PolicyValues field to given value.

### HasPolicyValues

`func (o *NamingPolicy) HasPolicyValues() bool`

HasPolicyValues returns a boolean if a field has been set.

### GetPolicyRestrictions

`func (o *NamingPolicy) GetPolicyRestrictions() []NamingPolicyRestriction`

GetPolicyRestrictions returns the PolicyRestrictions field if non-nil, zero value otherwise.

### GetPolicyRestrictionsOk

`func (o *NamingPolicy) GetPolicyRestrictionsOk() (*[]NamingPolicyRestriction, bool)`

GetPolicyRestrictionsOk returns a tuple with the PolicyRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRestrictions

`func (o *NamingPolicy) SetPolicyRestrictions(v []NamingPolicyRestriction)`

SetPolicyRestrictions sets PolicyRestrictions field to given value.

### HasPolicyRestrictions

`func (o *NamingPolicy) HasPolicyRestrictions() bool`

HasPolicyRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


