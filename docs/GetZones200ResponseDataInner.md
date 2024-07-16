# GetZones200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**View** | Pointer to [**View**](View.md) |  | [optional] [readonly] 
**NumberName** | Pointer to **string** |  | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified name of the zone. | [optional] 
**DeploymentEnabled** | Pointer to **bool** | Indicates whether the zone is deployable. | [optional] 
**ResponsePolicyZoneType** | Pointer to **string** | The type of response policy zone. | [optional] 
**ResponsePolicy** | Pointer to [**InlinedResponsePolicy**](InlinedResponsePolicy.md) |  | [optional] 
**OverridePolicyType** | Pointer to **string** | The response policy type set for BlueCat Security Feed. | [optional] 
**OverrideRefreshTime** | Pointer to **string** | The custom refresh time interval set for BlueCat Security Feed, in seconds. | [optional] 
**RedirectTarget** | Pointer to **string** | The target FQDN to which all undesirable connection attemps are redirected. | [optional] 
**FeedCategories** | Pointer to **[]string** |  | [optional] 
**DynamicUpdateEnabled** | Pointer to **bool** | Indicates whether the dynamic update feature is enabled on the zone. When enabled, any resource records that are added, updated, or deleted within the zone will be automatically deployed to the primary DNS/DHCP Server of that zone using selective deployment. | [optional] 
**Template** | Pointer to [**InlinedZoneTemplate**](InlinedZoneTemplate.md) |  | [optional] 
**Signed** | Pointer to **bool** | Indicates if the zone is currently signed using a DNSSEC Signing Policy. | [optional] 
**SigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetZones200ResponseDataInner

`func NewGetZones200ResponseDataInner() *GetZones200ResponseDataInner`

NewGetZones200ResponseDataInner instantiates a new GetZones200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZones200ResponseDataInnerWithDefaults

`func NewGetZones200ResponseDataInnerWithDefaults() *GetZones200ResponseDataInner`

NewGetZones200ResponseDataInnerWithDefaults instantiates a new GetZones200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetZones200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetZones200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetZones200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetZones200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetZones200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetZones200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetZones200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetZones200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetZones200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetZones200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetZones200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetZones200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetZones200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetZones200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetZones200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetZones200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetZones200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetZones200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetZones200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetZones200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetView

`func (o *GetZones200ResponseDataInner) GetView() View`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetZones200ResponseDataInner) GetViewOk() (*View, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetZones200ResponseDataInner) SetView(v View)`

SetView sets View field to given value.

### HasView

`func (o *GetZones200ResponseDataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetNumberName

`func (o *GetZones200ResponseDataInner) GetNumberName() string`

GetNumberName returns the NumberName field if non-nil, zero value otherwise.

### GetNumberNameOk

`func (o *GetZones200ResponseDataInner) GetNumberNameOk() (*string, bool)`

GetNumberNameOk returns a tuple with the NumberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberName

`func (o *GetZones200ResponseDataInner) SetNumberName(v string)`

SetNumberName sets NumberName field to given value.

### HasNumberName

`func (o *GetZones200ResponseDataInner) HasNumberName() bool`

HasNumberName returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetZones200ResponseDataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetZones200ResponseDataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetZones200ResponseDataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetZones200ResponseDataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *GetZones200ResponseDataInner) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *GetZones200ResponseDataInner) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *GetZones200ResponseDataInner) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.

### HasDeploymentEnabled

`func (o *GetZones200ResponseDataInner) HasDeploymentEnabled() bool`

HasDeploymentEnabled returns a boolean if a field has been set.

### GetResponsePolicyZoneType

`func (o *GetZones200ResponseDataInner) GetResponsePolicyZoneType() string`

GetResponsePolicyZoneType returns the ResponsePolicyZoneType field if non-nil, zero value otherwise.

### GetResponsePolicyZoneTypeOk

`func (o *GetZones200ResponseDataInner) GetResponsePolicyZoneTypeOk() (*string, bool)`

GetResponsePolicyZoneTypeOk returns a tuple with the ResponsePolicyZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicyZoneType

`func (o *GetZones200ResponseDataInner) SetResponsePolicyZoneType(v string)`

SetResponsePolicyZoneType sets ResponsePolicyZoneType field to given value.

### HasResponsePolicyZoneType

`func (o *GetZones200ResponseDataInner) HasResponsePolicyZoneType() bool`

HasResponsePolicyZoneType returns a boolean if a field has been set.

### GetResponsePolicy

`func (o *GetZones200ResponseDataInner) GetResponsePolicy() InlinedResponsePolicy`

GetResponsePolicy returns the ResponsePolicy field if non-nil, zero value otherwise.

### GetResponsePolicyOk

`func (o *GetZones200ResponseDataInner) GetResponsePolicyOk() (*InlinedResponsePolicy, bool)`

GetResponsePolicyOk returns a tuple with the ResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicy

`func (o *GetZones200ResponseDataInner) SetResponsePolicy(v InlinedResponsePolicy)`

SetResponsePolicy sets ResponsePolicy field to given value.

### HasResponsePolicy

`func (o *GetZones200ResponseDataInner) HasResponsePolicy() bool`

HasResponsePolicy returns a boolean if a field has been set.

### GetOverridePolicyType

`func (o *GetZones200ResponseDataInner) GetOverridePolicyType() string`

GetOverridePolicyType returns the OverridePolicyType field if non-nil, zero value otherwise.

### GetOverridePolicyTypeOk

`func (o *GetZones200ResponseDataInner) GetOverridePolicyTypeOk() (*string, bool)`

GetOverridePolicyTypeOk returns a tuple with the OverridePolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePolicyType

`func (o *GetZones200ResponseDataInner) SetOverridePolicyType(v string)`

SetOverridePolicyType sets OverridePolicyType field to given value.

### HasOverridePolicyType

`func (o *GetZones200ResponseDataInner) HasOverridePolicyType() bool`

HasOverridePolicyType returns a boolean if a field has been set.

### GetOverrideRefreshTime

`func (o *GetZones200ResponseDataInner) GetOverrideRefreshTime() string`

GetOverrideRefreshTime returns the OverrideRefreshTime field if non-nil, zero value otherwise.

### GetOverrideRefreshTimeOk

`func (o *GetZones200ResponseDataInner) GetOverrideRefreshTimeOk() (*string, bool)`

GetOverrideRefreshTimeOk returns a tuple with the OverrideRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRefreshTime

`func (o *GetZones200ResponseDataInner) SetOverrideRefreshTime(v string)`

SetOverrideRefreshTime sets OverrideRefreshTime field to given value.

### HasOverrideRefreshTime

`func (o *GetZones200ResponseDataInner) HasOverrideRefreshTime() bool`

HasOverrideRefreshTime returns a boolean if a field has been set.

### GetRedirectTarget

`func (o *GetZones200ResponseDataInner) GetRedirectTarget() string`

GetRedirectTarget returns the RedirectTarget field if non-nil, zero value otherwise.

### GetRedirectTargetOk

`func (o *GetZones200ResponseDataInner) GetRedirectTargetOk() (*string, bool)`

GetRedirectTargetOk returns a tuple with the RedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTarget

`func (o *GetZones200ResponseDataInner) SetRedirectTarget(v string)`

SetRedirectTarget sets RedirectTarget field to given value.

### HasRedirectTarget

`func (o *GetZones200ResponseDataInner) HasRedirectTarget() bool`

HasRedirectTarget returns a boolean if a field has been set.

### GetFeedCategories

`func (o *GetZones200ResponseDataInner) GetFeedCategories() []string`

GetFeedCategories returns the FeedCategories field if non-nil, zero value otherwise.

### GetFeedCategoriesOk

`func (o *GetZones200ResponseDataInner) GetFeedCategoriesOk() (*[]string, bool)`

GetFeedCategoriesOk returns a tuple with the FeedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategories

`func (o *GetZones200ResponseDataInner) SetFeedCategories(v []string)`

SetFeedCategories sets FeedCategories field to given value.

### HasFeedCategories

`func (o *GetZones200ResponseDataInner) HasFeedCategories() bool`

HasFeedCategories returns a boolean if a field has been set.

### GetDynamicUpdateEnabled

`func (o *GetZones200ResponseDataInner) GetDynamicUpdateEnabled() bool`

GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field if non-nil, zero value otherwise.

### GetDynamicUpdateEnabledOk

`func (o *GetZones200ResponseDataInner) GetDynamicUpdateEnabledOk() (*bool, bool)`

GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateEnabled

`func (o *GetZones200ResponseDataInner) SetDynamicUpdateEnabled(v bool)`

SetDynamicUpdateEnabled sets DynamicUpdateEnabled field to given value.

### HasDynamicUpdateEnabled

`func (o *GetZones200ResponseDataInner) HasDynamicUpdateEnabled() bool`

HasDynamicUpdateEnabled returns a boolean if a field has been set.

### GetTemplate

`func (o *GetZones200ResponseDataInner) GetTemplate() InlinedZoneTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetZones200ResponseDataInner) GetTemplateOk() (*InlinedZoneTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetZones200ResponseDataInner) SetTemplate(v InlinedZoneTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetZones200ResponseDataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSigned

`func (o *GetZones200ResponseDataInner) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *GetZones200ResponseDataInner) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *GetZones200ResponseDataInner) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *GetZones200ResponseDataInner) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSigningPolicy

`func (o *GetZones200ResponseDataInner) GetSigningPolicy() InlinedDNSSECSigningPolicy`

GetSigningPolicy returns the SigningPolicy field if non-nil, zero value otherwise.

### GetSigningPolicyOk

`func (o *GetZones200ResponseDataInner) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetSigningPolicyOk returns a tuple with the SigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPolicy

`func (o *GetZones200ResponseDataInner) SetSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetSigningPolicy sets SigningPolicy field to given value.

### HasSigningPolicy

`func (o *GetZones200ResponseDataInner) HasSigningPolicy() bool`

HasSigningPolicy returns a boolean if a field has been set.

### GetLinks

`func (o *GetZones200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetZones200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetZones200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetZones200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


