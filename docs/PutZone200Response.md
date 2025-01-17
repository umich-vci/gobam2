# PutZone200Response

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
**DynamicUpdateEnabled** | Pointer to **bool** | Indicates whether the dynamic update feature is enabled on the zone. When enabled, any resource records that are added, updated, or deleted within the zone will be automatically deployed to the primary DNS/DHCP Server of that zone using selective deployment. | [optional] 
**Template** | Pointer to [**InlinedZoneTemplate**](InlinedZoneTemplate.md) |  | [optional] 
**Signed** | Pointer to **bool** | Indicates if the zone is currently signed using a DNSSEC Signing Policy. | [optional] 
**SigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**ResponsePolicyZoneType** | Pointer to **string** | The type of response policy zone. | [optional] 
**ResponsePolicy** | Pointer to [**InlinedResponsePolicy**](InlinedResponsePolicy.md) |  | [optional] 
**OverridePolicyType** | Pointer to **string** | The response policy type set for BlueCat Security Feed. | [optional] 
**OverrideRefreshTime** | Pointer to **string** | The custom refresh time interval set for BlueCat Security Feed, in seconds. | [optional] 
**RedirectTarget** | Pointer to **string** | The target FQDN to which all undesirable connection attemps are redirected. | [optional] 
**FeedCategories** | Pointer to **[]string** |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewPutZone200Response

`func NewPutZone200Response() *PutZone200Response`

NewPutZone200Response instantiates a new PutZone200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutZone200ResponseWithDefaults

`func NewPutZone200ResponseWithDefaults() *PutZone200Response`

NewPutZone200ResponseWithDefaults instantiates a new PutZone200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutZone200Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutZone200Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutZone200Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutZone200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutZone200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutZone200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutZone200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutZone200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PutZone200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutZone200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutZone200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutZone200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PutZone200Response) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutZone200Response) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutZone200Response) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutZone200Response) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PutZone200Response) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PutZone200Response) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PutZone200Response) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PutZone200Response) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetView

`func (o *PutZone200Response) GetView() View`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PutZone200Response) GetViewOk() (*View, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PutZone200Response) SetView(v View)`

SetView sets View field to given value.

### HasView

`func (o *PutZone200Response) HasView() bool`

HasView returns a boolean if a field has been set.

### GetNumberName

`func (o *PutZone200Response) GetNumberName() string`

GetNumberName returns the NumberName field if non-nil, zero value otherwise.

### GetNumberNameOk

`func (o *PutZone200Response) GetNumberNameOk() (*string, bool)`

GetNumberNameOk returns a tuple with the NumberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberName

`func (o *PutZone200Response) SetNumberName(v string)`

SetNumberName sets NumberName field to given value.

### HasNumberName

`func (o *PutZone200Response) HasNumberName() bool`

HasNumberName returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *PutZone200Response) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PutZone200Response) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PutZone200Response) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PutZone200Response) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *PutZone200Response) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *PutZone200Response) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *PutZone200Response) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.

### HasDeploymentEnabled

`func (o *PutZone200Response) HasDeploymentEnabled() bool`

HasDeploymentEnabled returns a boolean if a field has been set.

### GetDynamicUpdateEnabled

`func (o *PutZone200Response) GetDynamicUpdateEnabled() bool`

GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field if non-nil, zero value otherwise.

### GetDynamicUpdateEnabledOk

`func (o *PutZone200Response) GetDynamicUpdateEnabledOk() (*bool, bool)`

GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateEnabled

`func (o *PutZone200Response) SetDynamicUpdateEnabled(v bool)`

SetDynamicUpdateEnabled sets DynamicUpdateEnabled field to given value.

### HasDynamicUpdateEnabled

`func (o *PutZone200Response) HasDynamicUpdateEnabled() bool`

HasDynamicUpdateEnabled returns a boolean if a field has been set.

### GetTemplate

`func (o *PutZone200Response) GetTemplate() InlinedZoneTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PutZone200Response) GetTemplateOk() (*InlinedZoneTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PutZone200Response) SetTemplate(v InlinedZoneTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PutZone200Response) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSigned

`func (o *PutZone200Response) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *PutZone200Response) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *PutZone200Response) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *PutZone200Response) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSigningPolicy

`func (o *PutZone200Response) GetSigningPolicy() InlinedDNSSECSigningPolicy`

GetSigningPolicy returns the SigningPolicy field if non-nil, zero value otherwise.

### GetSigningPolicyOk

`func (o *PutZone200Response) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetSigningPolicyOk returns a tuple with the SigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPolicy

`func (o *PutZone200Response) SetSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetSigningPolicy sets SigningPolicy field to given value.

### HasSigningPolicy

`func (o *PutZone200Response) HasSigningPolicy() bool`

HasSigningPolicy returns a boolean if a field has been set.

### GetResponsePolicyZoneType

`func (o *PutZone200Response) GetResponsePolicyZoneType() string`

GetResponsePolicyZoneType returns the ResponsePolicyZoneType field if non-nil, zero value otherwise.

### GetResponsePolicyZoneTypeOk

`func (o *PutZone200Response) GetResponsePolicyZoneTypeOk() (*string, bool)`

GetResponsePolicyZoneTypeOk returns a tuple with the ResponsePolicyZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicyZoneType

`func (o *PutZone200Response) SetResponsePolicyZoneType(v string)`

SetResponsePolicyZoneType sets ResponsePolicyZoneType field to given value.

### HasResponsePolicyZoneType

`func (o *PutZone200Response) HasResponsePolicyZoneType() bool`

HasResponsePolicyZoneType returns a boolean if a field has been set.

### GetResponsePolicy

`func (o *PutZone200Response) GetResponsePolicy() InlinedResponsePolicy`

GetResponsePolicy returns the ResponsePolicy field if non-nil, zero value otherwise.

### GetResponsePolicyOk

`func (o *PutZone200Response) GetResponsePolicyOk() (*InlinedResponsePolicy, bool)`

GetResponsePolicyOk returns a tuple with the ResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicy

`func (o *PutZone200Response) SetResponsePolicy(v InlinedResponsePolicy)`

SetResponsePolicy sets ResponsePolicy field to given value.

### HasResponsePolicy

`func (o *PutZone200Response) HasResponsePolicy() bool`

HasResponsePolicy returns a boolean if a field has been set.

### GetOverridePolicyType

`func (o *PutZone200Response) GetOverridePolicyType() string`

GetOverridePolicyType returns the OverridePolicyType field if non-nil, zero value otherwise.

### GetOverridePolicyTypeOk

`func (o *PutZone200Response) GetOverridePolicyTypeOk() (*string, bool)`

GetOverridePolicyTypeOk returns a tuple with the OverridePolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePolicyType

`func (o *PutZone200Response) SetOverridePolicyType(v string)`

SetOverridePolicyType sets OverridePolicyType field to given value.

### HasOverridePolicyType

`func (o *PutZone200Response) HasOverridePolicyType() bool`

HasOverridePolicyType returns a boolean if a field has been set.

### GetOverrideRefreshTime

`func (o *PutZone200Response) GetOverrideRefreshTime() string`

GetOverrideRefreshTime returns the OverrideRefreshTime field if non-nil, zero value otherwise.

### GetOverrideRefreshTimeOk

`func (o *PutZone200Response) GetOverrideRefreshTimeOk() (*string, bool)`

GetOverrideRefreshTimeOk returns a tuple with the OverrideRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRefreshTime

`func (o *PutZone200Response) SetOverrideRefreshTime(v string)`

SetOverrideRefreshTime sets OverrideRefreshTime field to given value.

### HasOverrideRefreshTime

`func (o *PutZone200Response) HasOverrideRefreshTime() bool`

HasOverrideRefreshTime returns a boolean if a field has been set.

### GetRedirectTarget

`func (o *PutZone200Response) GetRedirectTarget() string`

GetRedirectTarget returns the RedirectTarget field if non-nil, zero value otherwise.

### GetRedirectTargetOk

`func (o *PutZone200Response) GetRedirectTargetOk() (*string, bool)`

GetRedirectTargetOk returns a tuple with the RedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTarget

`func (o *PutZone200Response) SetRedirectTarget(v string)`

SetRedirectTarget sets RedirectTarget field to given value.

### HasRedirectTarget

`func (o *PutZone200Response) HasRedirectTarget() bool`

HasRedirectTarget returns a boolean if a field has been set.

### GetFeedCategories

`func (o *PutZone200Response) GetFeedCategories() []string`

GetFeedCategories returns the FeedCategories field if non-nil, zero value otherwise.

### GetFeedCategoriesOk

`func (o *PutZone200Response) GetFeedCategoriesOk() (*[]string, bool)`

GetFeedCategoriesOk returns a tuple with the FeedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategories

`func (o *PutZone200Response) SetFeedCategories(v []string)`

SetFeedCategories sets FeedCategories field to given value.

### HasFeedCategories

`func (o *PutZone200Response) HasFeedCategories() bool`

HasFeedCategories returns a boolean if a field has been set.

### GetLinks

`func (o *PutZone200Response) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PutZone200Response) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PutZone200Response) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PutZone200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


