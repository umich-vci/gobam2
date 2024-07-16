# GetSessions200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetSessions200ResponseDataInner

`func NewGetSessions200ResponseDataInner() *GetSessions200ResponseDataInner`

NewGetSessions200ResponseDataInner instantiates a new GetSessions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessions200ResponseDataInnerWithDefaults

`func NewGetSessions200ResponseDataInnerWithDefaults() *GetSessions200ResponseDataInner`

NewGetSessions200ResponseDataInnerWithDefaults instantiates a new GetSessions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSessions200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSessions200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSessions200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSessions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetSessions200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSessions200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSessions200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSessions200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *GetSessions200ResponseDataInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetSessions200ResponseDataInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetSessions200ResponseDataInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetSessions200ResponseDataInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetSessions200ResponseDataInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetSessions200ResponseDataInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetSessions200ResponseDataInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetSessions200ResponseDataInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetApiToken

`func (o *GetSessions200ResponseDataInner) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *GetSessions200ResponseDataInner) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *GetSessions200ResponseDataInner) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *GetSessions200ResponseDataInner) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetApiTokenExpirationDateTime

`func (o *GetSessions200ResponseDataInner) GetApiTokenExpirationDateTime() time.Time`

GetApiTokenExpirationDateTime returns the ApiTokenExpirationDateTime field if non-nil, zero value otherwise.

### GetApiTokenExpirationDateTimeOk

`func (o *GetSessions200ResponseDataInner) GetApiTokenExpirationDateTimeOk() (*time.Time, bool)`

GetApiTokenExpirationDateTimeOk returns a tuple with the ApiTokenExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTokenExpirationDateTime

`func (o *GetSessions200ResponseDataInner) SetApiTokenExpirationDateTime(v time.Time)`

SetApiTokenExpirationDateTime sets ApiTokenExpirationDateTime field to given value.

### HasApiTokenExpirationDateTime

`func (o *GetSessions200ResponseDataInner) HasApiTokenExpirationDateTime() bool`

HasApiTokenExpirationDateTime returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *GetSessions200ResponseDataInner) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *GetSessions200ResponseDataInner) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *GetSessions200ResponseDataInner) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *GetSessions200ResponseDataInner) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetReadOnly

`func (o *GetSessions200ResponseDataInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GetSessions200ResponseDataInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GetSessions200ResponseDataInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GetSessions200ResponseDataInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetLoginDateTime

`func (o *GetSessions200ResponseDataInner) GetLoginDateTime() time.Time`

GetLoginDateTime returns the LoginDateTime field if non-nil, zero value otherwise.

### GetLoginDateTimeOk

`func (o *GetSessions200ResponseDataInner) GetLoginDateTimeOk() (*time.Time, bool)`

GetLoginDateTimeOk returns a tuple with the LoginDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDateTime

`func (o *GetSessions200ResponseDataInner) SetLoginDateTime(v time.Time)`

SetLoginDateTime sets LoginDateTime field to given value.

### HasLoginDateTime

`func (o *GetSessions200ResponseDataInner) HasLoginDateTime() bool`

HasLoginDateTime returns a boolean if a field has been set.

### GetLogoutDateTime

`func (o *GetSessions200ResponseDataInner) GetLogoutDateTime() time.Time`

GetLogoutDateTime returns the LogoutDateTime field if non-nil, zero value otherwise.

### GetLogoutDateTimeOk

`func (o *GetSessions200ResponseDataInner) GetLogoutDateTimeOk() (*time.Time, bool)`

GetLogoutDateTimeOk returns a tuple with the LogoutDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutDateTime

`func (o *GetSessions200ResponseDataInner) SetLogoutDateTime(v time.Time)`

SetLogoutDateTime sets LogoutDateTime field to given value.

### HasLogoutDateTime

`func (o *GetSessions200ResponseDataInner) HasLogoutDateTime() bool`

HasLogoutDateTime returns a boolean if a field has been set.

### GetState

`func (o *GetSessions200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetSessions200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetSessions200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetSessions200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetResponse

`func (o *GetSessions200ResponseDataInner) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *GetSessions200ResponseDataInner) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *GetSessions200ResponseDataInner) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *GetSessions200ResponseDataInner) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetUser

`func (o *GetSessions200ResponseDataInner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSessions200ResponseDataInner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSessions200ResponseDataInner) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *GetSessions200ResponseDataInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAuthenticator

`func (o *GetSessions200ResponseDataInner) GetAuthenticator() Authenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *GetSessions200ResponseDataInner) GetAuthenticatorOk() (*Authenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *GetSessions200ResponseDataInner) SetAuthenticator(v Authenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *GetSessions200ResponseDataInner) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetLinks

`func (o *GetSessions200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetSessions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetSessions200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetSessions200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


