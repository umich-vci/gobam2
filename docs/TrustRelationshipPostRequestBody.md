# TrustRelationshipPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Address** | **string** | The IP address of the remote Address Manager server that is part of the trust relationship. | 
**Username** | **string** | The API username of the remote Address Manager server that is used to establish the trust relationship. | 
**Password** | **string** | The password of the API user that is used to establish the trust relationship with the remote Address Manager server. | 
**KeyGenerationDateTime** | Pointer to **time.Time** | The timestamp of when the SSH key was created for exchange with the remote Address Manager server to establish the trust relationship. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the trust relationship. | [optional] 

## Methods

### NewTrustRelationshipPostRequestBody

`func NewTrustRelationshipPostRequestBody(address string, username string, password string, ) *TrustRelationshipPostRequestBody`

NewTrustRelationshipPostRequestBody instantiates a new TrustRelationshipPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustRelationshipPostRequestBodyWithDefaults

`func NewTrustRelationshipPostRequestBodyWithDefaults() *TrustRelationshipPostRequestBody`

NewTrustRelationshipPostRequestBodyWithDefaults instantiates a new TrustRelationshipPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrustRelationshipPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrustRelationshipPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrustRelationshipPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TrustRelationshipPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TrustRelationshipPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrustRelationshipPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrustRelationshipPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TrustRelationshipPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *TrustRelationshipPostRequestBody) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TrustRelationshipPostRequestBody) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TrustRelationshipPostRequestBody) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetUsername

`func (o *TrustRelationshipPostRequestBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TrustRelationshipPostRequestBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TrustRelationshipPostRequestBody) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *TrustRelationshipPostRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TrustRelationshipPostRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TrustRelationshipPostRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetKeyGenerationDateTime

`func (o *TrustRelationshipPostRequestBody) GetKeyGenerationDateTime() time.Time`

GetKeyGenerationDateTime returns the KeyGenerationDateTime field if non-nil, zero value otherwise.

### GetKeyGenerationDateTimeOk

`func (o *TrustRelationshipPostRequestBody) GetKeyGenerationDateTimeOk() (*time.Time, bool)`

GetKeyGenerationDateTimeOk returns a tuple with the KeyGenerationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyGenerationDateTime

`func (o *TrustRelationshipPostRequestBody) SetKeyGenerationDateTime(v time.Time)`

SetKeyGenerationDateTime sets KeyGenerationDateTime field to given value.

### HasKeyGenerationDateTime

`func (o *TrustRelationshipPostRequestBody) HasKeyGenerationDateTime() bool`

HasKeyGenerationDateTime returns a boolean if a field has been set.

### GetState

`func (o *TrustRelationshipPostRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TrustRelationshipPostRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TrustRelationshipPostRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TrustRelationshipPostRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


