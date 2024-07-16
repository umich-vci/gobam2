# PostCollectionDeploymentRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**RoleType** | **string** |  | 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**NsRecordTtl** | Pointer to **int64** | The time-to-live (TTL) value for the name server and glue records deployed via the deployment role. | [optional] 
**Interfaces** | [**[]InlinedDnsRoleServerInterface**](InlinedDnsRoleServerInterface.md) | The server interfaces to which the DHCP deployment role is assigned. | 
**Server** | [**InlinedServer**](InlinedServer.md) |  | 

## Methods

### NewPostCollectionDeploymentRoleRequest

`func NewPostCollectionDeploymentRoleRequest(type_ string, roleType string, interfaces []InlinedDnsRoleServerInterface, server InlinedServer, ) *PostCollectionDeploymentRoleRequest`

NewPostCollectionDeploymentRoleRequest instantiates a new PostCollectionDeploymentRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCollectionDeploymentRoleRequestWithDefaults

`func NewPostCollectionDeploymentRoleRequestWithDefaults() *PostCollectionDeploymentRoleRequest`

NewPostCollectionDeploymentRoleRequestWithDefaults instantiates a new PostCollectionDeploymentRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostCollectionDeploymentRoleRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostCollectionDeploymentRoleRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostCollectionDeploymentRoleRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PostCollectionDeploymentRoleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PostCollectionDeploymentRoleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCollectionDeploymentRoleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCollectionDeploymentRoleRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PostCollectionDeploymentRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostCollectionDeploymentRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostCollectionDeploymentRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostCollectionDeploymentRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostCollectionDeploymentRoleRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostCollectionDeploymentRoleRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostCollectionDeploymentRoleRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostCollectionDeploymentRoleRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PostCollectionDeploymentRoleRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostCollectionDeploymentRoleRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostCollectionDeploymentRoleRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostCollectionDeploymentRoleRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRoleType

`func (o *PostCollectionDeploymentRoleRequest) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *PostCollectionDeploymentRoleRequest) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *PostCollectionDeploymentRoleRequest) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetView

`func (o *PostCollectionDeploymentRoleRequest) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PostCollectionDeploymentRoleRequest) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PostCollectionDeploymentRoleRequest) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *PostCollectionDeploymentRoleRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetNsRecordTtl

`func (o *PostCollectionDeploymentRoleRequest) GetNsRecordTtl() int64`

GetNsRecordTtl returns the NsRecordTtl field if non-nil, zero value otherwise.

### GetNsRecordTtlOk

`func (o *PostCollectionDeploymentRoleRequest) GetNsRecordTtlOk() (*int64, bool)`

GetNsRecordTtlOk returns a tuple with the NsRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsRecordTtl

`func (o *PostCollectionDeploymentRoleRequest) SetNsRecordTtl(v int64)`

SetNsRecordTtl sets NsRecordTtl field to given value.

### HasNsRecordTtl

`func (o *PostCollectionDeploymentRoleRequest) HasNsRecordTtl() bool`

HasNsRecordTtl returns a boolean if a field has been set.

### GetInterfaces

`func (o *PostCollectionDeploymentRoleRequest) GetInterfaces() []InlinedDnsRoleServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *PostCollectionDeploymentRoleRequest) GetInterfacesOk() (*[]InlinedDnsRoleServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *PostCollectionDeploymentRoleRequest) SetInterfaces(v []InlinedDnsRoleServerInterface)`

SetInterfaces sets Interfaces field to given value.


### GetServer

`func (o *PostCollectionDeploymentRoleRequest) GetServer() InlinedServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PostCollectionDeploymentRoleRequest) GetServerOk() (*InlinedServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PostCollectionDeploymentRoleRequest) SetServer(v InlinedServer)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


