# DHCPClientIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Uid** | Pointer to [**UID**](UID.md) |  | [optional] 

## Methods

### NewDHCPClientIdentifier

`func NewDHCPClientIdentifier() *DHCPClientIdentifier`

NewDHCPClientIdentifier instantiates a new DHCPClientIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPClientIdentifierWithDefaults

`func NewDHCPClientIdentifierWithDefaults() *DHCPClientIdentifier`

NewDHCPClientIdentifierWithDefaults instantiates a new DHCPClientIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPClientIdentifier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPClientIdentifier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPClientIdentifier) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPClientIdentifier) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *DHCPClientIdentifier) GetUid() UID`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DHCPClientIdentifier) GetUidOk() (*UID, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DHCPClientIdentifier) SetUid(v UID)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DHCPClientIdentifier) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


