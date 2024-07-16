# UserSessionPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Username** | **string** | The username used to initiate the session. | 
**Password** | **string** | The associated password for the username used to initiate the session. | 
**ApiToken** | Pointer to **string** | The API token of the user session. | [optional] [readonly] 
**ApiTokenExpirationDateTime** | Pointer to **time.Time** | The timestamp of the API token expiration. | [optional] [readonly] 
**RemoteAddress** | Pointer to **string** | The IP address of the device that initiated the user session. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | Indicates whether the initiated session is a read-only session. When database replication is configured between Address Manager servers, you can initiate a read-only session to call API methods with read-only support and retrieve information from the Secondary or Tertiary Address Manager server. | [optional] 
**LoginDateTime** | Pointer to **time.Time** | A timestamp of when the session was initiated. | [optional] [readonly] 
**LogoutDateTime** | Pointer to **time.Time** | A timestamp of when the session was terminated. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the user session. | [optional] 
**Response** | Pointer to **string** | The response from the authenticated user session. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Authenticator** | Pointer to [**Authenticator**](Authenticator.md) |  | [optional] 

## Methods

### NewUserSessionPostRequestBody

`func NewUserSessionPostRequestBody(username string, password string, ) *UserSessionPostRequestBody`

NewUserSessionPostRequestBody instantiates a new UserSessionPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionPostRequestBodyWithDefaults

`func NewUserSessionPostRequestBodyWithDefaults() *UserSessionPostRequestBody`

NewUserSessionPostRequestBodyWithDefaults instantiates a new UserSessionPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSessionPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSessionPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSessionPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserSessionPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserSessionPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSessionPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *UserSessionPostRequestBody) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSessionPostRequestBody) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSessionPostRequestBody) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserSessionPostRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserSessionPostRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserSessionPostRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetApiToken

`func (o *UserSessionPostRequestBody) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *UserSessionPostRequestBody) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *UserSessionPostRequestBody) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *UserSessionPostRequestBody) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetApiTokenExpirationDateTime

`func (o *UserSessionPostRequestBody) GetApiTokenExpirationDateTime() time.Time`

GetApiTokenExpirationDateTime returns the ApiTokenExpirationDateTime field if non-nil, zero value otherwise.

### GetApiTokenExpirationDateTimeOk

`func (o *UserSessionPostRequestBody) GetApiTokenExpirationDateTimeOk() (*time.Time, bool)`

GetApiTokenExpirationDateTimeOk returns a tuple with the ApiTokenExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenExpirationDateTime

`func (o *UserSessionPostRequestBody) SetApiTokenExpirationDateTime(v time.Time)`

SetApiTokenExpirationDateTime sets ApiTokenExpirationDateTime field to given value.

### HasApiTokenExpirationDateTime

`func (o *UserSessionPostRequestBody) HasApiTokenExpirationDateTime() bool`

HasApiTokenExpirationDateTime returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *UserSessionPostRequestBody) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *UserSessionPostRequestBody) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *UserSessionPostRequestBody) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *UserSessionPostRequestBody) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetReadOnly

`func (o *UserSessionPostRequestBody) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UserSessionPostRequestBody) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UserSessionPostRequestBody) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UserSessionPostRequestBody) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetLoginDateTime

`func (o *UserSessionPostRequestBody) GetLoginDateTime() time.Time`

GetLoginDateTime returns the LoginDateTime field if non-nil, zero value otherwise.

### GetLoginDateTimeOk

`func (o *UserSessionPostRequestBody) GetLoginDateTimeOk() (*time.Time, bool)`

GetLoginDateTimeOk returns a tuple with the LoginDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDateTime

`func (o *UserSessionPostRequestBody) SetLoginDateTime(v time.Time)`

SetLoginDateTime sets LoginDateTime field to given value.

### HasLoginDateTime

`func (o *UserSessionPostRequestBody) HasLoginDateTime() bool`

HasLoginDateTime returns a boolean if a field has been set.

### GetLogoutDateTime

`func (o *UserSessionPostRequestBody) GetLogoutDateTime() time.Time`

GetLogoutDateTime returns the LogoutDateTime field if non-nil, zero value otherwise.

### GetLogoutDateTimeOk

`func (o *UserSessionPostRequestBody) GetLogoutDateTimeOk() (*time.Time, bool)`

GetLogoutDateTimeOk returns a tuple with the LogoutDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutDateTime

`func (o *UserSessionPostRequestBody) SetLogoutDateTime(v time.Time)`

SetLogoutDateTime sets LogoutDateTime field to given value.

### HasLogoutDateTime

`func (o *UserSessionPostRequestBody) HasLogoutDateTime() bool`

HasLogoutDateTime returns a boolean if a field has been set.

### GetState

`func (o *UserSessionPostRequestBody) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserSessionPostRequestBody) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserSessionPostRequestBody) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserSessionPostRequestBody) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResponse

`func (o *UserSessionPostRequestBody) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *UserSessionPostRequestBody) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *UserSessionPostRequestBody) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *UserSessionPostRequestBody) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetUser

`func (o *UserSessionPostRequestBody) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSessionPostRequestBody) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSessionPostRequestBody) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserSessionPostRequestBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserSessionPostRequestBody) GetAuthenticator() Authenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserSessionPostRequestBody) GetAuthenticatorOk() (*Authenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserSessionPostRequestBody) SetAuthenticator(v Authenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserSessionPostRequestBody) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


