# GetZones200Response1DataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**View** | Pointer to [**View**](View.md) |  | [optional] [readonly] 
**DeploymentEnabled** | Pointer to **bool** | Indicates whether the zone is deployable. | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified name of the zone. | [optional] 
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

## Methods

### NewGetZones200Response1DataInner

`func NewGetZones200Response1DataInner() *GetZones200Response1DataInner`

NewGetZones200Response1DataInner instantiates a new GetZones200Response1DataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZones200Response1DataInnerWithDefaults

`func NewGetZones200Response1DataInnerWithDefaults() *GetZones200Response1DataInner`

NewGetZones200Response1DataInnerWithDefaults instantiates a new GetZones200Response1DataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetZones200Response1DataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetZones200Response1DataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetZones200Response1DataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetZones200Response1DataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetZones200Response1DataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetZones200Response1DataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetZones200Response1DataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetZones200Response1DataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetZones200Response1DataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetZones200Response1DataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetZones200Response1DataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetZones200Response1DataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetZones200Response1DataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetZones200Response1DataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetZones200Response1DataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetZones200Response1DataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetZones200Response1DataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetZones200Response1DataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetZones200Response1DataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetZones200Response1DataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetView

`func (o *GetZones200Response1DataInner) GetView() View`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetZones200Response1DataInner) GetViewOk() (*View, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetZones200Response1DataInner) SetView(v View)`

SetView sets View field to given value.

### HasView

`func (o *GetZones200Response1DataInner) HasView() bool`

HasView returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *GetZones200Response1DataInner) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *GetZones200Response1DataInner) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *GetZones200Response1DataInner) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.

### HasDeploymentEnabled

`func (o *GetZones200Response1DataInner) HasDeploymentEnabled() bool`

HasDeploymentEnabled returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *GetZones200Response1DataInner) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *GetZones200Response1DataInner) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *GetZones200Response1DataInner) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *GetZones200Response1DataInner) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetResponsePolicyZoneType

`func (o *GetZones200Response1DataInner) GetResponsePolicyZoneType() string`

GetResponsePolicyZoneType returns the ResponsePolicyZoneType field if non-nil, zero value otherwise.

### GetResponsePolicyZoneTypeOk

`func (o *GetZones200Response1DataInner) GetResponsePolicyZoneTypeOk() (*string, bool)`

GetResponsePolicyZoneTypeOk returns a tuple with the ResponsePolicyZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicyZoneType

`func (o *GetZones200Response1DataInner) SetResponsePolicyZoneType(v string)`

SetResponsePolicyZoneType sets ResponsePolicyZoneType field to given value.

### HasResponsePolicyZoneType

`func (o *GetZones200Response1DataInner) HasResponsePolicyZoneType() bool`

HasResponsePolicyZoneType returns a boolean if a field has been set.

### GetResponsePolicy

`func (o *GetZones200Response1DataInner) GetResponsePolicy() InlinedResponsePolicy`

GetResponsePolicy returns the ResponsePolicy field if non-nil, zero value otherwise.

### GetResponsePolicyOk

`func (o *GetZones200Response1DataInner) GetResponsePolicyOk() (*InlinedResponsePolicy, bool)`

GetResponsePolicyOk returns a tuple with the ResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicy

`func (o *GetZones200Response1DataInner) SetResponsePolicy(v InlinedResponsePolicy)`

SetResponsePolicy sets ResponsePolicy field to given value.

### HasResponsePolicy

`func (o *GetZones200Response1DataInner) HasResponsePolicy() bool`

HasResponsePolicy returns a boolean if a field has been set.

### GetOverridePolicyType

`func (o *GetZones200Response1DataInner) GetOverridePolicyType() string`

GetOverridePolicyType returns the OverridePolicyType field if non-nil, zero value otherwise.

### GetOverridePolicyTypeOk

`func (o *GetZones200Response1DataInner) GetOverridePolicyTypeOk() (*string, bool)`

GetOverridePolicyTypeOk returns a tuple with the OverridePolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePolicyType

`func (o *GetZones200Response1DataInner) SetOverridePolicyType(v string)`

SetOverridePolicyType sets OverridePolicyType field to given value.

### HasOverridePolicyType

`func (o *GetZones200Response1DataInner) HasOverridePolicyType() bool`

HasOverridePolicyType returns a boolean if a field has been set.

### GetOverrideRefreshTime

`func (o *GetZones200Response1DataInner) GetOverrideRefreshTime() string`

GetOverrideRefreshTime returns the OverrideRefreshTime field if non-nil, zero value otherwise.

### GetOverrideRefreshTimeOk

`func (o *GetZones200Response1DataInner) GetOverrideRefreshTimeOk() (*string, bool)`

GetOverrideRefreshTimeOk returns a tuple with the OverrideRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRefreshTime

`func (o *GetZones200Response1DataInner) SetOverrideRefreshTime(v string)`

SetOverrideRefreshTime sets OverrideRefreshTime field to given value.

### HasOverrideRefreshTime

`func (o *GetZones200Response1DataInner) HasOverrideRefreshTime() bool`

HasOverrideRefreshTime returns a boolean if a field has been set.

### GetRedirectTarget

`func (o *GetZones200Response1DataInner) GetRedirectTarget() string`

GetRedirectTarget returns the RedirectTarget field if non-nil, zero value otherwise.

### GetRedirectTargetOk

`func (o *GetZones200Response1DataInner) GetRedirectTargetOk() (*string, bool)`

GetRedirectTargetOk returns a tuple with the RedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTarget

`func (o *GetZones200Response1DataInner) SetRedirectTarget(v string)`

SetRedirectTarget sets RedirectTarget field to given value.

### HasRedirectTarget

`func (o *GetZones200Response1DataInner) HasRedirectTarget() bool`

HasRedirectTarget returns a boolean if a field has been set.

### GetFeedCategories

`func (o *GetZones200Response1DataInner) GetFeedCategories() []string`

GetFeedCategories returns the FeedCategories field if non-nil, zero value otherwise.

### GetFeedCategoriesOk

`func (o *GetZones200Response1DataInner) GetFeedCategoriesOk() (*[]string, bool)`

GetFeedCategoriesOk returns a tuple with the FeedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategories

`func (o *GetZones200Response1DataInner) SetFeedCategories(v []string)`

SetFeedCategories sets FeedCategories field to given value.

### HasFeedCategories

`func (o *GetZones200Response1DataInner) HasFeedCategories() bool`

HasFeedCategories returns a boolean if a field has been set.

### GetDynamicUpdateEnabled

`func (o *GetZones200Response1DataInner) GetDynamicUpdateEnabled() bool`

GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field if non-nil, zero value otherwise.

### GetDynamicUpdateEnabledOk

`func (o *GetZones200Response1DataInner) GetDynamicUpdateEnabledOk() (*bool, bool)`

GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateEnabled

`func (o *GetZones200Response1DataInner) SetDynamicUpdateEnabled(v bool)`

SetDynamicUpdateEnabled sets DynamicUpdateEnabled field to given value.

### HasDynamicUpdateEnabled

`func (o *GetZones200Response1DataInner) HasDynamicUpdateEnabled() bool`

HasDynamicUpdateEnabled returns a boolean if a field has been set.

### GetTemplate

`func (o *GetZones200Response1DataInner) GetTemplate() InlinedZoneTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetZones200Response1DataInner) GetTemplateOk() (*InlinedZoneTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetZones200Response1DataInner) SetTemplate(v InlinedZoneTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetZones200Response1DataInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSigned

`func (o *GetZones200Response1DataInner) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *GetZones200Response1DataInner) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *GetZones200Response1DataInner) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *GetZones200Response1DataInner) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSigningPolicy

`func (o *GetZones200Response1DataInner) GetSigningPolicy() InlinedDNSSECSigningPolicy`

GetSigningPolicy returns the SigningPolicy field if non-nil, zero value otherwise.

### GetSigningPolicyOk

`func (o *GetZones200Response1DataInner) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetSigningPolicyOk returns a tuple with the SigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPolicy

`func (o *GetZones200Response1DataInner) SetSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetSigningPolicy sets SigningPolicy field to given value.

### HasSigningPolicy

`func (o *GetZones200Response1DataInner) HasSigningPolicy() bool`

HasSigningPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


