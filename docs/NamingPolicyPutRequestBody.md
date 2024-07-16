# NamingPolicyPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the naming policy. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**PolicyValues** | Pointer to [**[]NamingPolicyValue**](NamingPolicyValue.md) |  | [optional] 
**PolicyRestrictions** | Pointer to [**[]NamingPolicyRestriction**](NamingPolicyRestriction.md) |  | [optional] 

## Methods

### NewNamingPolicyPutRequestBody

`func NewNamingPolicyPutRequestBody(name string, ) *NamingPolicyPutRequestBody`

NewNamingPolicyPutRequestBody instantiates a new NamingPolicyPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyPutRequestBodyWithDefaults

`func NewNamingPolicyPutRequestBodyWithDefaults() *NamingPolicyPutRequestBody`

NewNamingPolicyPutRequestBodyWithDefaults instantiates a new NamingPolicyPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingPolicyPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingPolicyPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingPolicyPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NamingPolicyPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NamingPolicyPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicyPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicyPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicyPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NamingPolicyPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamingPolicyPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamingPolicyPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *NamingPolicyPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *NamingPolicyPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *NamingPolicyPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *NamingPolicyPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPolicyValues

`func (o *NamingPolicyPutRequestBody) GetPolicyValues() []NamingPolicyValue`

GetPolicyValues returns the PolicyValues field if non-nil, zero value otherwise.

### GetPolicyValuesOk

`func (o *NamingPolicyPutRequestBody) GetPolicyValuesOk() (*[]NamingPolicyValue, bool)`

GetPolicyValuesOk returns a tuple with the PolicyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyValues

`func (o *NamingPolicyPutRequestBody) SetPolicyValues(v []NamingPolicyValue)`

SetPolicyValues sets PolicyValues field to given value.

### HasPolicyValues

`func (o *NamingPolicyPutRequestBody) HasPolicyValues() bool`

HasPolicyValues returns a boolean if a field has been set.

### GetPolicyRestrictions

`func (o *NamingPolicyPutRequestBody) GetPolicyRestrictions() []NamingPolicyRestriction`

GetPolicyRestrictions returns the PolicyRestrictions field if non-nil, zero value otherwise.

### GetPolicyRestrictionsOk

`func (o *NamingPolicyPutRequestBody) GetPolicyRestrictionsOk() (*[]NamingPolicyRestriction, bool)`

GetPolicyRestrictionsOk returns a tuple with the PolicyRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRestrictions

`func (o *NamingPolicyPutRequestBody) SetPolicyRestrictions(v []NamingPolicyRestriction)`

SetPolicyRestrictions sets PolicyRestrictions field to given value.

### HasPolicyRestrictions

`func (o *NamingPolicyPutRequestBody) HasPolicyRestrictions() bool`

HasPolicyRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


