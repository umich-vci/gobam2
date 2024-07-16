# VendorProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the DHCP vendor profile. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Identifier** | Pointer to **string** | The Vendor Class Identifier. | [optional] 
**Description** | Pointer to **string** | The description of the DHCP vendor profile. | [optional] 

## Methods

### NewVendorProfile

`func NewVendorProfile() *VendorProfile`

NewVendorProfile instantiates a new VendorProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorProfileWithDefaults

`func NewVendorProfileWithDefaults() *VendorProfile`

NewVendorProfileWithDefaults instantiates a new VendorProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorProfile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorProfile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorProfile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VendorProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VendorProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VendorProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *VendorProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VendorProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VendorProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VendorProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *VendorProfile) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *VendorProfile) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *VendorProfile) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *VendorProfile) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetIdentifier

`func (o *VendorProfile) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VendorProfile) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VendorProfile) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *VendorProfile) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *VendorProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VendorProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VendorProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VendorProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


