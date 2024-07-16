# VendorProfilePostRequestBody

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

### NewVendorProfilePostRequestBody

`func NewVendorProfilePostRequestBody(name string, identifier string, description string, ) *VendorProfilePostRequestBody`

NewVendorProfilePostRequestBody instantiates a new VendorProfilePostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorProfilePostRequestBodyWithDefaults

`func NewVendorProfilePostRequestBodyWithDefaults() *VendorProfilePostRequestBody`

NewVendorProfilePostRequestBodyWithDefaults instantiates a new VendorProfilePostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VendorProfilePostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VendorProfilePostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VendorProfilePostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VendorProfilePostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VendorProfilePostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VendorProfilePostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VendorProfilePostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VendorProfilePostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *VendorProfilePostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VendorProfilePostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VendorProfilePostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *VendorProfilePostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *VendorProfilePostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *VendorProfilePostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *VendorProfilePostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetIdentifier

`func (o *VendorProfilePostRequestBody) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *VendorProfilePostRequestBody) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *VendorProfilePostRequestBody) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetDescription

`func (o *VendorProfilePostRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VendorProfilePostRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VendorProfilePostRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


