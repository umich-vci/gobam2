# GetNamingPolicyRestrictions200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Restrictions** | Pointer to [**[]ValueMatchTypePairBean**](ValueMatchTypePairBean.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetNamingPolicyRestrictions200ResponseDataInner

`func NewGetNamingPolicyRestrictions200ResponseDataInner() *GetNamingPolicyRestrictions200ResponseDataInner`

NewGetNamingPolicyRestrictions200ResponseDataInner instantiates a new GetNamingPolicyRestrictions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNamingPolicyRestrictions200ResponseDataInnerWithDefaults

`func NewGetNamingPolicyRestrictions200ResponseDataInnerWithDefaults() *GetNamingPolicyRestrictions200ResponseDataInner`

NewGetNamingPolicyRestrictions200ResponseDataInnerWithDefaults instantiates a new GetNamingPolicyRestrictions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetRestrictions

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetRestrictions() []ValueMatchTypePairBean`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetRestrictionsOk() (*[]ValueMatchTypePairBean, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) SetRestrictions(v []ValueMatchTypePairBean)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetLinks

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetNamingPolicyRestrictions200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


