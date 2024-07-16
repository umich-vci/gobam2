# GetCollectionRestrictedZones200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**DeviceRegistrationEnabled** | Pointer to **bool** | Indicates whether DNS redirection has been enabled for use with the BlueCat Device Registration Portal. | [optional] 
**DeviceRegistrationPortalAddress** | Pointer to **string** | The IPv4 address of the BlueCat Device Registration Portal. | [optional] 
**View** | Pointer to [**View**](View.md) |  | [optional] [readonly] 
**DeploymentEnabled** | Pointer to **bool** | Indicates whether the zone is deployable. | [optional] 
**DynamicUpdateEnabled** | Pointer to **bool** | Indicates whether the dynamic update feature is enabled on the zone. When enabled, any resource records that are added, updated, or deleted within the zone will be automatically deployed to the primary DNS/DHCP Server of that zone using selective deployment. | [optional] 
**Template** | Pointer to [**InlinedZoneTemplate**](InlinedZoneTemplate.md) |  | [optional] 
**Signed** | Pointer to **bool** | Indicates if the zone is currently signed using a DNSSEC Signing Policy. | [optional] 
**SigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified name of the zone. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetCollectionRestrictedZones200ResponseDataInner

`func NewGetCollectionRestrictedZones200ResponseDataInner() *GetCollectionRestrictedZones200ResponseDataInner`

NewGetCollectionRestrictedZones200ResponseDataInner instantiates a new GetCollectionRestrictedZones200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionRestrictedZones200ResponseDataInnerWithDefaults

`func NewGetCollectionRestrictedZones200ResponseDataInnerWithDefaults() *GetCollectionRestrictedZones200ResponseDataInner`

NewGetCollectionRestrictedZones200ResponseDataInnerWithDefaults instantiates a new GetCollectionRestrictedZones200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDeviceRegistrationEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDeviceRegistrationEnabled() bool`

GetDeviceRegistrationEnabled returns the DeviceRegistrationEnabled field if non-nil, zero value otherwise.

### GetDeviceRegistrationEnabledOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDeviceRegistrationEnabledOk() (*bool, bool)`

GetDeviceRegistrationEnabledOk returns a tuple with the DeviceRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetDeviceRegistrationEnabled(v bool)`

SetDeviceRegistrationEnabled sets DeviceRegistrationEnabled field to given value.

### HasDeviceRegistrationEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasDeviceRegistrationEnabled() bool`

HasDeviceRegistrationEnabled returns a boolean if a field has been set.

### GetDeviceRegistrationPortalAddress

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDeviceRegistrationPortalAddress() string`

GetDeviceRegistrationPortalAddress returns the DeviceRegistrationPortalAddress field if non-nil, zero value otherwise.

### GetDeviceRegistrationPortalAddressOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDeviceRegistrationPortalAddressOk() (*string, bool)`

GetDeviceRegistrationPortalAddressOk returns a tuple with the DeviceRegistrationPortalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistrationPortalAddress

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetDeviceRegistrationPortalAddress(v string)`

SetDeviceRegistrationPortalAddress sets DeviceRegistrationPortalAddress field to given value.

### HasDeviceRegistrationPortalAddress

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasDeviceRegistrationPortalAddress() bool`

HasDeviceRegistrationPortalAddress returns a boolean if a field has been set.

### GetView

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetView() View`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetViewOk() (*View, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetView(v View)`

SetView sets View field to given value.

### HasView

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.

### HasDeploymentEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasDeploymentEnabled() bool`

HasDeploymentEnabled returns a boolean if a field has been set.

### GetDynamicUpdateEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDynamicUpdateEnabled() bool`

GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field if non-nil, zero value otherwise.

### GetDynamicUpdateEnabledOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetDynamicUpdateEnabledOk() (*bool, bool)`

GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetDynamicUpdateEnabled(v bool)`

SetDynamicUpdateEnabled sets DynamicUpdateEnabled field to given value.

### HasDynamicUpdateEnabled

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasDynamicUpdateEnabled() bool`

HasDynamicUpdateEnabled returns a boolean if a field has been set.

### GetTemplate

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetTemplate() InlinedZoneTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetTemplateOk() (*InlinedZoneTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetTemplate(v InlinedZoneTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSigned

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSigningPolicy

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetSigningPolicy() InlinedDNSSECSigningPolicy`

GetSigningPolicy returns the SigningPolicy field if non-nil, zero value otherwise.

### GetSigningPolicyOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetSigningPolicyOk returns a tuple with the SigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPolicy

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetSigningPolicy sets SigningPolicy field to given value.

### HasSigningPolicy

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasSigningPolicy() bool`

HasSigningPolicy returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetLinks

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCollectionRestrictedZones200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCollectionRestrictedZones200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetCollectionRestrictedZones200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


