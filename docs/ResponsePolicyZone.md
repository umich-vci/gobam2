# ResponsePolicyZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**ResponsePolicyZoneType** | Pointer to **string** | The type of response policy zone. | [optional] 
**ResponsePolicy** | Pointer to [**InlinedResponsePolicy**](InlinedResponsePolicy.md) |  | [optional] 
**OverridePolicyType** | Pointer to **string** | The response policy type set for BlueCat Security Feed. | [optional] 
**OverrideRefreshTime** | Pointer to **string** | The custom refresh time interval set for BlueCat Security Feed, in seconds. | [optional] 
**RedirectTarget** | Pointer to **string** | The target FQDN to which all undesirable connection attemps are redirected. | [optional] 
**FeedCategories** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResponsePolicyZone

`func NewResponsePolicyZone() *ResponsePolicyZone`

NewResponsePolicyZone instantiates a new ResponsePolicyZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePolicyZoneWithDefaults

`func NewResponsePolicyZoneWithDefaults() *ResponsePolicyZone`

NewResponsePolicyZoneWithDefaults instantiates a new ResponsePolicyZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponsePolicyZone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsePolicyZone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsePolicyZone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsePolicyZone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ResponsePolicyZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsePolicyZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsePolicyZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsePolicyZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResponsePolicyZoneType

`func (o *ResponsePolicyZone) GetResponsePolicyZoneType() string`

GetResponsePolicyZoneType returns the ResponsePolicyZoneType field if non-nil, zero value otherwise.

### GetResponsePolicyZoneTypeOk

`func (o *ResponsePolicyZone) GetResponsePolicyZoneTypeOk() (*string, bool)`

GetResponsePolicyZoneTypeOk returns a tuple with the ResponsePolicyZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicyZoneType

`func (o *ResponsePolicyZone) SetResponsePolicyZoneType(v string)`

SetResponsePolicyZoneType sets ResponsePolicyZoneType field to given value.

### HasResponsePolicyZoneType

`func (o *ResponsePolicyZone) HasResponsePolicyZoneType() bool`

HasResponsePolicyZoneType returns a boolean if a field has been set.

### GetResponsePolicy

`func (o *ResponsePolicyZone) GetResponsePolicy() InlinedResponsePolicy`

GetResponsePolicy returns the ResponsePolicy field if non-nil, zero value otherwise.

### GetResponsePolicyOk

`func (o *ResponsePolicyZone) GetResponsePolicyOk() (*InlinedResponsePolicy, bool)`

GetResponsePolicyOk returns a tuple with the ResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePolicy

`func (o *ResponsePolicyZone) SetResponsePolicy(v InlinedResponsePolicy)`

SetResponsePolicy sets ResponsePolicy field to given value.

### HasResponsePolicy

`func (o *ResponsePolicyZone) HasResponsePolicy() bool`

HasResponsePolicy returns a boolean if a field has been set.

### GetOverridePolicyType

`func (o *ResponsePolicyZone) GetOverridePolicyType() string`

GetOverridePolicyType returns the OverridePolicyType field if non-nil, zero value otherwise.

### GetOverridePolicyTypeOk

`func (o *ResponsePolicyZone) GetOverridePolicyTypeOk() (*string, bool)`

GetOverridePolicyTypeOk returns a tuple with the OverridePolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePolicyType

`func (o *ResponsePolicyZone) SetOverridePolicyType(v string)`

SetOverridePolicyType sets OverridePolicyType field to given value.

### HasOverridePolicyType

`func (o *ResponsePolicyZone) HasOverridePolicyType() bool`

HasOverridePolicyType returns a boolean if a field has been set.

### GetOverrideRefreshTime

`func (o *ResponsePolicyZone) GetOverrideRefreshTime() string`

GetOverrideRefreshTime returns the OverrideRefreshTime field if non-nil, zero value otherwise.

### GetOverrideRefreshTimeOk

`func (o *ResponsePolicyZone) GetOverrideRefreshTimeOk() (*string, bool)`

GetOverrideRefreshTimeOk returns a tuple with the OverrideRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideRefreshTime

`func (o *ResponsePolicyZone) SetOverrideRefreshTime(v string)`

SetOverrideRefreshTime sets OverrideRefreshTime field to given value.

### HasOverrideRefreshTime

`func (o *ResponsePolicyZone) HasOverrideRefreshTime() bool`

HasOverrideRefreshTime returns a boolean if a field has been set.

### GetRedirectTarget

`func (o *ResponsePolicyZone) GetRedirectTarget() string`

GetRedirectTarget returns the RedirectTarget field if non-nil, zero value otherwise.

### GetRedirectTargetOk

`func (o *ResponsePolicyZone) GetRedirectTargetOk() (*string, bool)`

GetRedirectTargetOk returns a tuple with the RedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTarget

`func (o *ResponsePolicyZone) SetRedirectTarget(v string)`

SetRedirectTarget sets RedirectTarget field to given value.

### HasRedirectTarget

`func (o *ResponsePolicyZone) HasRedirectTarget() bool`

HasRedirectTarget returns a boolean if a field has been set.

### GetFeedCategories

`func (o *ResponsePolicyZone) GetFeedCategories() []string`

GetFeedCategories returns the FeedCategories field if non-nil, zero value otherwise.

### GetFeedCategoriesOk

`func (o *ResponsePolicyZone) GetFeedCategoriesOk() (*[]string, bool)`

GetFeedCategoriesOk returns a tuple with the FeedCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedCategories

`func (o *ResponsePolicyZone) SetFeedCategories(v []string)`

SetFeedCategories sets FeedCategories field to given value.

### HasFeedCategories

`func (o *ResponsePolicyZone) HasFeedCategories() bool`

HasFeedCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


