# PublishedInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**PrimaryAddress** | Pointer to **string** | The primary IP address of the published interface. | [optional] 

## Methods

### NewPublishedInterface

`func NewPublishedInterface() *PublishedInterface`

NewPublishedInterface instantiates a new PublishedInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedInterfaceWithDefaults

`func NewPublishedInterfaceWithDefaults() *PublishedInterface`

NewPublishedInterfaceWithDefaults instantiates a new PublishedInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublishedInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublishedInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublishedInterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PublishedInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrimaryAddress

`func (o *PublishedInterface) GetPrimaryAddress() string`

GetPrimaryAddress returns the PrimaryAddress field if non-nil, zero value otherwise.

### GetPrimaryAddressOk

`func (o *PublishedInterface) GetPrimaryAddressOk() (*string, bool)`

GetPrimaryAddressOk returns a tuple with the PrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddress

`func (o *PublishedInterface) SetPrimaryAddress(v string)`

SetPrimaryAddress sets PrimaryAddress field to given value.

### HasPrimaryAddress

`func (o *PublishedInterface) HasPrimaryAddress() bool`

HasPrimaryAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


