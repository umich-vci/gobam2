# GetNamingPolicies200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the naming policy. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**PolicyValues** | Pointer to [**[]NamingPolicyValue**](NamingPolicyValue.md) |  | [optional] 
**PolicyRestrictions** | Pointer to [**[]NamingPolicyRestriction**](NamingPolicyRestriction.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetNamingPolicies200ResponseDataInner

`func NewGetNamingPolicies200ResponseDataInner() *GetNamingPolicies200ResponseDataInner`

NewGetNamingPolicies200ResponseDataInner instantiates a new GetNamingPolicies200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNamingPolicies200ResponseDataInnerWithDefaults

`func NewGetNamingPolicies200ResponseDataInnerWithDefaults() *GetNamingPolicies200ResponseDataInner`

NewGetNamingPolicies200ResponseDataInnerWithDefaults instantiates a new GetNamingPolicies200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNamingPolicies200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNamingPolicies200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNamingPolicies200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNamingPolicies200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetNamingPolicies200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNamingPolicies200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNamingPolicies200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNamingPolicies200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetNamingPolicies200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNamingPolicies200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNamingPolicies200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNamingPolicies200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetNamingPolicies200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetNamingPolicies200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetNamingPolicies200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetNamingPolicies200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetPolicyValues

`func (o *GetNamingPolicies200ResponseDataInner) GetPolicyValues() []NamingPolicyValue`

GetPolicyValues returns the PolicyValues field if non-nil, zero value otherwise.

### GetPolicyValuesOk

`func (o *GetNamingPolicies200ResponseDataInner) GetPolicyValuesOk() (*[]NamingPolicyValue, bool)`

GetPolicyValuesOk returns a tuple with the PolicyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyValues

`func (o *GetNamingPolicies200ResponseDataInner) SetPolicyValues(v []NamingPolicyValue)`

SetPolicyValues sets PolicyValues field to given value.

### HasPolicyValues

`func (o *GetNamingPolicies200ResponseDataInner) HasPolicyValues() bool`

HasPolicyValues returns a boolean if a field has been set.

### GetPolicyRestrictions

`func (o *GetNamingPolicies200ResponseDataInner) GetPolicyRestrictions() []NamingPolicyRestriction`

GetPolicyRestrictions returns the PolicyRestrictions field if non-nil, zero value otherwise.

### GetPolicyRestrictionsOk

`func (o *GetNamingPolicies200ResponseDataInner) GetPolicyRestrictionsOk() (*[]NamingPolicyRestriction, bool)`

GetPolicyRestrictionsOk returns a tuple with the PolicyRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyRestrictions

`func (o *GetNamingPolicies200ResponseDataInner) SetPolicyRestrictions(v []NamingPolicyRestriction)`

SetPolicyRestrictions sets PolicyRestrictions field to given value.

### HasPolicyRestrictions

`func (o *GetNamingPolicies200ResponseDataInner) HasPolicyRestrictions() bool`

HasPolicyRestrictions returns a boolean if a field has been set.

### GetLinks

`func (o *GetNamingPolicies200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetNamingPolicies200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetNamingPolicies200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetNamingPolicies200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


