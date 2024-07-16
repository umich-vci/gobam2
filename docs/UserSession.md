# UserSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Username** | Pointer to **string** | The username used to initiate the session. | [optional] 
**Password** | Pointer to **string** | The associated password for the username used to initiate the session. | [optional] 
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

### NewUserSession

`func NewUserSession() *UserSession`

NewUserSession instantiates a new UserSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionWithDefaults

`func NewUserSessionWithDefaults() *UserSession`

NewUserSessionWithDefaults instantiates a new UserSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSession) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSession) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSession) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UserSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserSession) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSession) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSession) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSession) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *UserSession) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSession) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSession) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSession) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *UserSession) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserSession) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserSession) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserSession) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiToken

`func (o *UserSession) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *UserSession) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *UserSession) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *UserSession) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetApiTokenExpirationDateTime

`func (o *UserSession) GetApiTokenExpirationDateTime() time.Time`

GetApiTokenExpirationDateTime returns the ApiTokenExpirationDateTime field if non-nil, zero value otherwise.

### GetApiTokenExpirationDateTimeOk

`func (o *UserSession) GetApiTokenExpirationDateTimeOk() (*time.Time, bool)`

GetApiTokenExpirationDateTimeOk returns a tuple with the ApiTokenExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenExpirationDateTime

`func (o *UserSession) SetApiTokenExpirationDateTime(v time.Time)`

SetApiTokenExpirationDateTime sets ApiTokenExpirationDateTime field to given value.

### HasApiTokenExpirationDateTime

`func (o *UserSession) HasApiTokenExpirationDateTime() bool`

HasApiTokenExpirationDateTime returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *UserSession) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *UserSession) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *UserSession) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *UserSession) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetReadOnly

`func (o *UserSession) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UserSession) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UserSession) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UserSession) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetLoginDateTime

`func (o *UserSession) GetLoginDateTime() time.Time`

GetLoginDateTime returns the LoginDateTime field if non-nil, zero value otherwise.

### GetLoginDateTimeOk

`func (o *UserSession) GetLoginDateTimeOk() (*time.Time, bool)`

GetLoginDateTimeOk returns a tuple with the LoginDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDateTime

`func (o *UserSession) SetLoginDateTime(v time.Time)`

SetLoginDateTime sets LoginDateTime field to given value.

### HasLoginDateTime

`func (o *UserSession) HasLoginDateTime() bool`

HasLoginDateTime returns a boolean if a field has been set.

### GetLogoutDateTime

`func (o *UserSession) GetLogoutDateTime() time.Time`

GetLogoutDateTime returns the LogoutDateTime field if non-nil, zero value otherwise.

### GetLogoutDateTimeOk

`func (o *UserSession) GetLogoutDateTimeOk() (*time.Time, bool)`

GetLogoutDateTimeOk returns a tuple with the LogoutDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutDateTime

`func (o *UserSession) SetLogoutDateTime(v time.Time)`

SetLogoutDateTime sets LogoutDateTime field to given value.

### HasLogoutDateTime

`func (o *UserSession) HasLogoutDateTime() bool`

HasLogoutDateTime returns a boolean if a field has been set.

### GetState

`func (o *UserSession) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserSession) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserSession) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserSession) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResponse

`func (o *UserSession) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *UserSession) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *UserSession) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *UserSession) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetUser

`func (o *UserSession) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSession) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSession) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserSession) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserSession) GetAuthenticator() Authenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserSession) GetAuthenticatorOk() (*Authenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserSession) SetAuthenticator(v Authenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserSession) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


