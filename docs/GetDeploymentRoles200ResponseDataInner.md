# GetDeploymentRoles200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**RoleType** | Pointer to **string** |  | [optional] 
**Interfaces** | Pointer to [**[]InlinedDnsRoleServerInterface**](InlinedDnsRoleServerInterface.md) | The server interfaces to which the DHCP deployment role is assigned. | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**NsRecordTtl** | Pointer to **int64** | The time-to-live (TTL) value for the name server and glue records deployed via the deployment role. | [optional] 
**Server** | Pointer to [**InlinedServer**](InlinedServer.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetDeploymentRoles200ResponseDataInner

`func NewGetDeploymentRoles200ResponseDataInner() *GetDeploymentRoles200ResponseDataInner`

NewGetDeploymentRoles200ResponseDataInner instantiates a new GetDeploymentRoles200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentRoles200ResponseDataInnerWithDefaults

`func NewGetDeploymentRoles200ResponseDataInnerWithDefaults() *GetDeploymentRoles200ResponseDataInner`

NewGetDeploymentRoles200ResponseDataInnerWithDefaults instantiates a new GetDeploymentRoles200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeploymentRoles200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeploymentRoles200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeploymentRoles200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetDeploymentRoles200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDeploymentRoles200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDeploymentRoles200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetDeploymentRoles200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeploymentRoles200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeploymentRoles200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetDeploymentRoles200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetDeploymentRoles200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetDeploymentRoles200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetDeploymentRoles200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetDeploymentRoles200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetDeploymentRoles200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetRoleType

`func (o *GetDeploymentRoles200ResponseDataInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *GetDeploymentRoles200ResponseDataInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *GetDeploymentRoles200ResponseDataInner) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetDeploymentRoles200ResponseDataInner) GetInterfaces() []InlinedDnsRoleServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetInterfacesOk() (*[]InlinedDnsRoleServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetDeploymentRoles200ResponseDataInner) SetInterfaces(v []InlinedDnsRoleServerInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetDeploymentRoles200ResponseDataInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetView

`func (o *GetDeploymentRoles200ResponseDataInner) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetDeploymentRoles200ResponseDataInner) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *GetDeploymentRoles200ResponseDataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetNsRecordTtl

`func (o *GetDeploymentRoles200ResponseDataInner) GetNsRecordTtl() int64`

GetNsRecordTtl returns the NsRecordTtl field if non-nil, zero value otherwise.

### GetNsRecordTtlOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetNsRecordTtlOk() (*int64, bool)`

GetNsRecordTtlOk returns a tuple with the NsRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsRecordTtl

`func (o *GetDeploymentRoles200ResponseDataInner) SetNsRecordTtl(v int64)`

SetNsRecordTtl sets NsRecordTtl field to given value.

### HasNsRecordTtl

`func (o *GetDeploymentRoles200ResponseDataInner) HasNsRecordTtl() bool`

HasNsRecordTtl returns a boolean if a field has been set.

### GetServer

`func (o *GetDeploymentRoles200ResponseDataInner) GetServer() InlinedServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetServerOk() (*InlinedServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetDeploymentRoles200ResponseDataInner) SetServer(v InlinedServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetDeploymentRoles200ResponseDataInner) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetLinks

`func (o *GetDeploymentRoles200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeploymentRoles200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeploymentRoles200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeploymentRoles200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


