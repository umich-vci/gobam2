# HAServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**Profile** | Pointer to **string** | The profile of the server. | [optional] [readonly] 
**Interfaces** | Pointer to [**[]HAServerAllOfInterfaces**](HAServerAllOfInterfaces.md) | The list of network interfaces of the server. | [optional] 
**ActiveServer** | Pointer to [**InlinedServer**](InlinedServer.md) |  | [optional] 
**PassiveServer** | Pointer to [**InlinedServer**](InlinedServer.md) |  | [optional] 

## Methods

### NewHAServer

`func NewHAServer() *HAServer`

NewHAServer instantiates a new HAServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAServerWithDefaults

`func NewHAServerWithDefaults() *HAServer`

NewHAServerWithDefaults instantiates a new HAServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HAServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HAServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HAServer) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HAServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *HAServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HAServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HAServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HAServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfile

`func (o *HAServer) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *HAServer) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *HAServer) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *HAServer) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetInterfaces

`func (o *HAServer) GetInterfaces() []HAServerAllOfInterfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *HAServer) GetInterfacesOk() (*[]HAServerAllOfInterfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *HAServer) SetInterfaces(v []HAServerAllOfInterfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *HAServer) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetActiveServer

`func (o *HAServer) GetActiveServer() InlinedServer`

GetActiveServer returns the ActiveServer field if non-nil, zero value otherwise.

### GetActiveServerOk

`func (o *HAServer) GetActiveServerOk() (*InlinedServer, bool)`

GetActiveServerOk returns a tuple with the ActiveServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveServer

`func (o *HAServer) SetActiveServer(v InlinedServer)`

SetActiveServer sets ActiveServer field to given value.

### HasActiveServer

`func (o *HAServer) HasActiveServer() bool`

HasActiveServer returns a boolean if a field has been set.

### GetPassiveServer

`func (o *HAServer) GetPassiveServer() InlinedServer`

GetPassiveServer returns the PassiveServer field if non-nil, zero value otherwise.

### GetPassiveServerOk

`func (o *HAServer) GetPassiveServerOk() (*InlinedServer, bool)`

GetPassiveServerOk returns a tuple with the PassiveServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveServer

`func (o *HAServer) SetPassiveServer(v InlinedServer)`

SetPassiveServer sets PassiveServer field to given value.

### HasPassiveServer

`func (o *HAServer) HasPassiveServer() bool`

HasPassiveServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


