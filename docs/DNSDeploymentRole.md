# DNSDeploymentRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**RoleType** | Pointer to **string** | The type of DNS deployment role. Roles set to NONE are not deployed, instead clearing data from the server they are applied to. | [optional] 
**View** | Pointer to [**InlinedView**](InlinedView.md) |  | [optional] 
**NsRecordTtl** | Pointer to **int64** | The time-to-live (TTL) value for the name server and glue records deployed via the deployment role. | [optional] 
**Interfaces** | Pointer to [**[]InlinedDnsRoleServerInterface**](InlinedDnsRoleServerInterface.md) | The server interfaces to which the DHCP deployment role is assigned. | [optional] 

## Methods

### NewDNSDeploymentRole

`func NewDNSDeploymentRole() *DNSDeploymentRole`

NewDNSDeploymentRole instantiates a new DNSDeploymentRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSDeploymentRoleWithDefaults

`func NewDNSDeploymentRoleWithDefaults() *DNSDeploymentRole`

NewDNSDeploymentRoleWithDefaults instantiates a new DNSDeploymentRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSDeploymentRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSDeploymentRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSDeploymentRole) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSDeploymentRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRoleType

`func (o *DNSDeploymentRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DNSDeploymentRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DNSDeploymentRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *DNSDeploymentRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetView

`func (o *DNSDeploymentRole) GetView() InlinedView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *DNSDeploymentRole) GetViewOk() (*InlinedView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *DNSDeploymentRole) SetView(v InlinedView)`

SetView sets View field to given value.

### HasView

`func (o *DNSDeploymentRole) HasView() bool`

HasView returns a boolean if a field has been set.

### GetNsRecordTtl

`func (o *DNSDeploymentRole) GetNsRecordTtl() int64`

GetNsRecordTtl returns the NsRecordTtl field if non-nil, zero value otherwise.

### GetNsRecordTtlOk

`func (o *DNSDeploymentRole) GetNsRecordTtlOk() (*int64, bool)`

GetNsRecordTtlOk returns a tuple with the NsRecordTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsRecordTtl

`func (o *DNSDeploymentRole) SetNsRecordTtl(v int64)`

SetNsRecordTtl sets NsRecordTtl field to given value.

### HasNsRecordTtl

`func (o *DNSDeploymentRole) HasNsRecordTtl() bool`

HasNsRecordTtl returns a boolean if a field has been set.

### GetInterfaces

`func (o *DNSDeploymentRole) GetInterfaces() []InlinedDnsRoleServerInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DNSDeploymentRole) GetInterfacesOk() (*[]InlinedDnsRoleServerInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DNSDeploymentRole) SetInterfaces(v []InlinedDnsRoleServerInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *DNSDeploymentRole) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


