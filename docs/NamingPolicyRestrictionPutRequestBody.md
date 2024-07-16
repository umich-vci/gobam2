# NamingPolicyRestrictionPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Restrictions** | [**[]ValueMatchTypePairBean**](ValueMatchTypePairBean.md) |  | 

## Methods

### NewNamingPolicyRestrictionPutRequestBody

`func NewNamingPolicyRestrictionPutRequestBody(name string, restrictions []ValueMatchTypePairBean, ) *NamingPolicyRestrictionPutRequestBody`

NewNamingPolicyRestrictionPutRequestBody instantiates a new NamingPolicyRestrictionPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamingPolicyRestrictionPutRequestBodyWithDefaults

`func NewNamingPolicyRestrictionPutRequestBodyWithDefaults() *NamingPolicyRestrictionPutRequestBody`

NewNamingPolicyRestrictionPutRequestBodyWithDefaults instantiates a new NamingPolicyRestrictionPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NamingPolicyRestrictionPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamingPolicyRestrictionPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamingPolicyRestrictionPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NamingPolicyRestrictionPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NamingPolicyRestrictionPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NamingPolicyRestrictionPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NamingPolicyRestrictionPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NamingPolicyRestrictionPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NamingPolicyRestrictionPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamingPolicyRestrictionPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamingPolicyRestrictionPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *NamingPolicyRestrictionPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *NamingPolicyRestrictionPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *NamingPolicyRestrictionPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *NamingPolicyRestrictionPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetRestrictions

`func (o *NamingPolicyRestrictionPutRequestBody) GetRestrictions() []ValueMatchTypePairBean`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *NamingPolicyRestrictionPutRequestBody) GetRestrictionsOk() (*[]ValueMatchTypePairBean, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *NamingPolicyRestrictionPutRequestBody) SetRestrictions(v []ValueMatchTypePairBean)`

SetRestrictions sets Restrictions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


