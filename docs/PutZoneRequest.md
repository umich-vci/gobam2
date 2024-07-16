# PutZoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | **string** | The resource type. | 
**Name** | **string** | The name of the resource. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**View** | Pointer to [**View**](View.md) |  | [optional] [readonly] 
**DeploymentEnabled** | **bool** | Indicates whether the zone is deployable. | 
**AbsoluteName** | Pointer to **string** | The fully qualified name of the zone. | [optional] 
**DynamicUpdateEnabled** | **bool** | Indicates whether the dynamic update feature is enabled on the zone. When enabled, any resource records that are added, updated, or deleted within the zone will be automatically deployed to the primary DNS/DHCP Server of that zone using selective deployment. | 
**Template** | Pointer to [**InlinedZoneTemplate**](InlinedZoneTemplate.md) |  | [optional] 
**Signed** | **bool** | Indicates if the zone is currently signed using a DNSSEC Signing Policy. | 
**SigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**ResponsePolicyZoneType** | **string** | The type of response policy zone. | 
**ResponsePolicy** | Pointer to [**InlinedResponsePolicy**](InlinedResponsePolicy.md) |  | [optional] 
**OverridePolicyType** | Pointer to **string** | The response policy type set for BlueCat Security Feed. | [optional] 
**OverrideRefreshTime** | Pointer to **string** | The custom refresh time interval set for BlueCat Security Feed, in seconds. | [optional] 
**RedirectTarget** | Pointer to **string** | The target FQDN to which all undesirable connection attemps are redirected. | [optional] 
**FeedCategories** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPutZoneRequest

`func NewPutZoneRequest(type_ string, name string, deploymentEnabled bool, dynamicUpdateEnabled bool, signed bool, responsePolicyZoneType string, ) *PutZoneRequest`

NewPutZoneRequest instantiates a new PutZoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutZoneRequestWithDefaults

`func NewPutZoneRequestWithDefaults() *PutZoneRequest`

NewPutZoneRequestWithDefaults instantiates a new PutZoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutZoneRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutZoneRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutZoneRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PutZoneRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PutZoneRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutZoneRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutZoneRequest) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PutZoneRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutZoneRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutZoneRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *PutZoneRequest) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PutZoneRequest) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PutZoneRequest) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PutZoneRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *PutZoneRequest) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PutZoneRequest) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PutZoneRequest) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PutZoneRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetView

`func (o *PutZoneRequest) GetView() View`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PutZoneRequest) GetViewOk() (*View, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PutZoneRequest) SetView(v View)`

SetView sets View field to given value.

### HasView

`func (o *PutZoneRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *PutZoneRequest) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *PutZoneRequest) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *PutZoneRequest) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.


### GetAbsoluteName

`func (o *PutZoneRequest) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *PutZoneRequest) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *PutZoneRequest) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *PutZoneRequest) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetDynamicUpdateEnabled

`func (o *PutZoneRequest) GetDynamicUpdateEnabled() bool`

GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field if non-nil, zero value otherwise.

### GetDynamicUpdateEnabledOk

`func (o *PutZoneRequest) GetDynamicUpdateEnabledOk() (*bool, bool)`

GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateEnabled

`func (o *PutZoneRequest) SetDynamicUpdateEnabled(v bool)`

SetDynamicUpdateEnabled sets DynamicUpdateEnabled field to given value.


### GetTemplate

`func (o *PutZoneRequest) GetTemplate() InlinedZoneTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PutZoneRequest) GetTemplateOk() (*InlinedZoneTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PutZoneRequest) SetTemplate(v InlinedZoneTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PutZoneRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSigned

`func (o *PutZoneRequest) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *PutZoneRequest) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *PutZoneRequest) SetSigned(v bool)`

SetSigned sets Signed field to given value.


### GetSigningPolicy

`func (o *PutZoneRequest) GetSigningPolicy() InlinedDNSSECSigningPolicy`

GetSigningPolicy returns the SigningPolicy field if non-nil, zero value otherwise.

### GetSigningPolicyOk

`func (o *PutZoneRequest) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetSigningPolicyOk returns a tuple with the SigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPolicy

`func (o *PutZoneRequest) SetSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetSigningPolicy sets SigningPolicy field to given value.

### HasSigningPolicy

`func (o *PutZoneRequest) HasSigningPolicy() bool`

HasSigningPolicy returns a boolean if a field has been set.

### GetResponsePolicyZoneType

`func (o *PutZoneRequest) GetResponsePolicyZoneType() string`

GetResponsePolicyZoneType returns the ResponsePolicyZoneType field if non-nil, zero value otherwise.

### GetResponsePolicyZoneTypeOk

`func (o *PutZoneRequest) GetResponsePolicyZoneTypeOk() (*string, bool)`

GetResponsePolicyZoneTypeOk returns a tuple with the ResponsePolicyZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicyZoneType

`func (o *PutZoneRequest) SetResponsePolicyZoneType(v string)`

SetResponsePolicyZoneType sets ResponsePolicyZoneType field to given value.


### GetResponsePolicy

`func (o *PutZoneRequest) GetResponsePolicy() InlinedResponsePolicy`

GetResponsePolicy returns the ResponsePolicy field if non-nil, zero value otherwise.

### GetResponsePolicyOk

`func (o *PutZoneRequest) GetResponsePolicyOk() (*InlinedResponsePolicy, bool)`

GetResponsePolicyOk returns a tuple with the ResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicy

`func (o *PutZoneRequest) SetResponsePolicy(v InlinedResponsePolicy)`

SetResponsePolicy sets ResponsePolicy field to given value.

### HasResponsePolicy

`func (o *PutZoneRequest) HasResponsePolicy() bool`

HasResponsePolicy returns a boolean if a field has been set.

### GetOverridePolicyType

`func (o *PutZoneRequest) GetOverridePolicyType() string`

GetOverridePolicyType returns the OverridePolicyType field if non-nil, zero value otherwise.

### GetOverridePolicyTypeOk

`func (o *PutZoneRequest) GetOverridePolicyTypeOk() (*string, bool)`

GetOverridePolicyTypeOk returns a tuple with the OverridePolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePolicyType

`func (o *PutZoneRequest) SetOverridePolicyType(v string)`

SetOverridePolicyType sets OverridePolicyType field to given value.

### HasOverridePolicyType

`func (o *PutZoneRequest) HasOverridePolicyType() bool`

HasOverridePolicyType returns a boolean if a field has been set.

### GetOverrideRefreshTime

`func (o *PutZoneRequest) GetOverrideRefreshTime() string`

GetOverrideRefreshTime returns the OverrideRefreshTime field if non-nil, zero value otherwise.

### GetOverrideRefreshTimeOk

`func (o *PutZoneRequest) GetOverrideRefreshTimeOk() (*string, bool)`

GetOverrideRefreshTimeOk returns a tuple with the OverrideRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRefreshTime

`func (o *PutZoneRequest) SetOverrideRefreshTime(v string)`

SetOverrideRefreshTime sets OverrideRefreshTime field to given value.

### HasOverrideRefreshTime

`func (o *PutZoneRequest) HasOverrideRefreshTime() bool`

HasOverrideRefreshTime returns a boolean if a field has been set.

### GetRedirectTarget

`func (o *PutZoneRequest) GetRedirectTarget() string`

GetRedirectTarget returns the RedirectTarget field if non-nil, zero value otherwise.

### GetRedirectTargetOk

`func (o *PutZoneRequest) GetRedirectTargetOk() (*string, bool)`

GetRedirectTargetOk returns a tuple with the RedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTarget

`func (o *PutZoneRequest) SetRedirectTarget(v string)`

SetRedirectTarget sets RedirectTarget field to given value.

### HasRedirectTarget

`func (o *PutZoneRequest) HasRedirectTarget() bool`

HasRedirectTarget returns a boolean if a field has been set.

### GetFeedCategories

`func (o *PutZoneRequest) GetFeedCategories() []string`

GetFeedCategories returns the FeedCategories field if non-nil, zero value otherwise.

### GetFeedCategoriesOk

`func (o *PutZoneRequest) GetFeedCategoriesOk() (*[]string, bool)`

GetFeedCategoriesOk returns a tuple with the FeedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategories

`func (o *PutZoneRequest) SetFeedCategories(v []string)`

SetFeedCategories sets FeedCategories field to given value.

### HasFeedCategories

`func (o *PutZoneRequest) HasFeedCategories() bool`

HasFeedCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


