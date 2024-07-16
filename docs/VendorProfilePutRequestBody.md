# VendorProfilePutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the DHCP vendor profile. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Identifier** | **string** | The Vendor Class Identifier. | 
**Description** | **string** | The description of the DHCP vendor profile. | 

## Methods

### NewVendorProfilePutRequestBody

`func NewVendorProfilePutRequestBody(name string, identifier string, description string, ) *VendorProfilePutRequestBody`

NewVendorProfilePutRequestBody instantiates a new VendorProfilePutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorProfilePutRequestBodyWithDefaults

`func NewVendorProfilePutRequestBodyWithDefaults() *VendorProfilePutRequestBody`

NewVendorProfilePutRequestBodyWithDefaults instantiates a new VendorProfilePutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorProfilePutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorProfilePutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorProfilePutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VendorProfilePutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VendorProfilePutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorProfilePutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorProfilePutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VendorProfilePutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *VendorProfilePutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VendorProfilePutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VendorProfilePutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *VendorProfilePutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *VendorProfilePutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *VendorProfilePutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *VendorProfilePutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetIdentifier

`func (o *VendorProfilePutRequestBody) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VendorProfilePutRequestBody) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VendorProfilePutRequestBody) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDescription

`func (o *VendorProfilePutRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VendorProfilePutRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VendorProfilePutRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


