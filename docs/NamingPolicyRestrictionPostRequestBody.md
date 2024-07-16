# NamingPolicyRestrictionPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Restrictions** | [**[]ValueMatchTypePairBean**](ValueMatchTypePairBean.md) |  | 

## Methods

### NewNamingPolicyRestrictionPostRequestBody

`func NewNamingPolicyRestrictionPostRequestBody(name string, restrictions []ValueMatchTypePairBean, ) *NamingPolicyRestrictionPostRequestBody`

NewNamingPolicyRestrictionPostRequestBody instantiates a new NamingPolicyRestrictionPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyRestrictionPostRequestBodyWithDefaults

`func NewNamingPolicyRestrictionPostRequestBodyWithDefaults() *NamingPolicyRestrictionPostRequestBody`

NewNamingPolicyRestrictionPostRequestBodyWithDefaults instantiates a new NamingPolicyRestrictionPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingPolicyRestrictionPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingPolicyRestrictionPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingPolicyRestrictionPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NamingPolicyRestrictionPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NamingPolicyRestrictionPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicyRestrictionPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicyRestrictionPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicyRestrictionPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NamingPolicyRestrictionPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamingPolicyRestrictionPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamingPolicyRestrictionPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *NamingPolicyRestrictionPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *NamingPolicyRestrictionPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *NamingPolicyRestrictionPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *NamingPolicyRestrictionPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetRestrictions

`func (o *NamingPolicyRestrictionPostRequestBody) GetRestrictions() []ValueMatchTypePairBean`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *NamingPolicyRestrictionPostRequestBody) GetRestrictionsOk() (*[]ValueMatchTypePairBean, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *NamingPolicyRestrictionPostRequestBody) SetRestrictions(v []ValueMatchTypePairBean)`

SetRestrictions sets Restrictions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


