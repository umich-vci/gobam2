# TrustRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Address** | Pointer to **string** | The IP address of the remote Address Manager server that is part of the trust relationship. | [optional] 
**Username** | Pointer to **string** | The API username of the remote Address Manager server that is used to establish the trust relationship. | [optional] 
**Password** | Pointer to **string** | The password of the API user that is used to establish the trust relationship with the remote Address Manager server. | [optional] 
**KeyGenerationDateTime** | Pointer to **time.Time** | The timestamp of when the SSH key was created for exchange with the remote Address Manager server to establish the trust relationship. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the trust relationship. | [optional] 

## Methods

### NewTrustRelationship

`func NewTrustRelationship() *TrustRelationship`

NewTrustRelationship instantiates a new TrustRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustRelationshipWithDefaults

`func NewTrustRelationshipWithDefaults() *TrustRelationship`

NewTrustRelationshipWithDefaults instantiates a new TrustRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustRelationship) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustRelationship) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustRelationship) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TrustRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TrustRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrustRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrustRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrustRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *TrustRelationship) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TrustRelationship) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TrustRelationship) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TrustRelationship) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetUsername

`func (o *TrustRelationship) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TrustRelationship) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TrustRelationship) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TrustRelationship) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *TrustRelationship) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TrustRelationship) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TrustRelationship) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TrustRelationship) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetKeyGenerationDateTime

`func (o *TrustRelationship) GetKeyGenerationDateTime() time.Time`

GetKeyGenerationDateTime returns the KeyGenerationDateTime field if non-nil, zero value otherwise.

### GetKeyGenerationDateTimeOk

`func (o *TrustRelationship) GetKeyGenerationDateTimeOk() (*time.Time, bool)`

GetKeyGenerationDateTimeOk returns a tuple with the KeyGenerationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationDateTime

`func (o *TrustRelationship) SetKeyGenerationDateTime(v time.Time)`

SetKeyGenerationDateTime sets KeyGenerationDateTime field to given value.

### HasKeyGenerationDateTime

`func (o *TrustRelationship) HasKeyGenerationDateTime() bool`

HasKeyGenerationDateTime returns a boolean if a field has been set.

### GetState

`func (o *TrustRelationship) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrustRelationship) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrustRelationship) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrustRelationship) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


