# GetDeploymentRoles200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**RoleType** | Pointer to **string** |  | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**NsRecordTtl** | Pointer to **int64** | The time-to-live (TTL) value for the name server and glue records deployed via the deployment role. | [optional] 
**Interfaces** | Pointer to [**[]InlinedDnsRoleServerInterface**](InlinedDnsRoleServerInterface.md) | The server interfaces to which the DHCP deployment role is assigned. | [optional] 
**Server** | Pointer to [**InlinedServer**](InlinedServer.md) |  | [optional] 

## Methods

### NewGetDeploymentRoles200Response1DataInner

`func NewGetDeploymentRoles200Response1DataInner() *GetDeploymentRoles200Response1DataInner`

NewGetDeploymentRoles200Response1DataInner instantiates a new GetDeploymentRoles200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentRoles200Response1DataInnerWithDefaults

`func NewGetDeploymentRoles200Response1DataInnerWithDefaults() *GetDeploymentRoles200Response1DataInner`

NewGetDeploymentRoles200Response1DataInnerWithDefaults instantiates a new GetDeploymentRoles200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeploymentRoles200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeploymentRoles200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeploymentRoles200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeploymentRoles200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetDeploymentRoles200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeploymentRoles200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeploymentRoles200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeploymentRoles200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetDeploymentRoles200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeploymentRoles200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeploymentRoles200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeploymentRoles200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetDeploymentRoles200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetDeploymentRoles200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetDeploymentRoles200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetDeploymentRoles200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetDeploymentRoles200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetDeploymentRoles200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetDeploymentRoles200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetDeploymentRoles200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRoleType

`func (o *GetDeploymentRoles200Response1DataInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *GetDeploymentRoles200Response1DataInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *GetDeploymentRoles200Response1DataInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *GetDeploymentRoles200Response1DataInner) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetView

`func (o *GetDeploymentRoles200Response1DataInner) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetDeploymentRoles200Response1DataInner) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetDeploymentRoles200Response1DataInner) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *GetDeploymentRoles200Response1DataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetNsRecordTtl

`func (o *GetDeploymentRoles200Response1DataInner) GetNsRecordTtl() int64`

GetNsRecordTtl returns the NsRecordTtl field if non-nil, zero value otherwise.

### GetNsRecordTtlOk

`func (o *GetDeploymentRoles200Response1DataInner) GetNsRecordTtlOk() (*int64, bool)`

GetNsRecordTtlOk returns a tuple with the NsRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsRecordTtl

`func (o *GetDeploymentRoles200Response1DataInner) SetNsRecordTtl(v int64)`

SetNsRecordTtl sets NsRecordTtl field to given value.

### HasNsRecordTtl

`func (o *GetDeploymentRoles200Response1DataInner) HasNsRecordTtl() bool`

HasNsRecordTtl returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetDeploymentRoles200Response1DataInner) GetInterfaces() []InlinedDnsRoleServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetDeploymentRoles200Response1DataInner) GetInterfacesOk() (*[]InlinedDnsRoleServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetDeploymentRoles200Response1DataInner) SetInterfaces(v []InlinedDnsRoleServerInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetDeploymentRoles200Response1DataInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetServer

`func (o *GetDeploymentRoles200Response1DataInner) GetServer() InlinedServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetDeploymentRoles200Response1DataInner) GetServerOk() (*InlinedServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetDeploymentRoles200Response1DataInner) SetServer(v InlinedServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetDeploymentRoles200Response1DataInner) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


