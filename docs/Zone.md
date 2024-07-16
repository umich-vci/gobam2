# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] 
**DeploymentEnabled** | Pointer to **bool** | Indicates whether the zone is deployable. | [optional] 
**DynamicUpdateEnabled** | Pointer to **bool** | Indicates whether the dynamic update feature is enabled on the zone. When enabled, any resource records that are added, updated, or deleted within the zone will be automatically deployed to the primary DNS/DHCP Server of that zone using selective deployment. | [optional] 
**Template** | Pointer to [**InlinedZoneTemplate**](InlinedZoneTemplate.md) |  | [optional] 
**Signed** | Pointer to **bool** | Indicates if the zone is currently signed using a DNSSEC Signing Policy. | [optional] 
**SigningPolicy** | Pointer to [**InlinedDNSSECSigningPolicy**](InlinedDNSSECSigningPolicy.md) |  | [optional] 
**AbsoluteName** | Pointer to **string** | The fully qualified name of the zone. | [optional] 

## Methods

### NewZone

`func NewZone() *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Zone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Zone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Zone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Zone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Zone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeploymentEnabled

`func (o *Zone) GetDeploymentEnabled() bool`

GetDeploymentEnabled returns the DeploymentEnabled field if non-nil, zero value otherwise.

### GetDeploymentEnabledOk

`func (o *Zone) GetDeploymentEnabledOk() (*bool, bool)`

GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentEnabled

`func (o *Zone) SetDeploymentEnabled(v bool)`

SetDeploymentEnabled sets DeploymentEnabled field to given value.

### HasDeploymentEnabled

`func (o *Zone) HasDeploymentEnabled() bool`

HasDeploymentEnabled returns a boolean if a field has been set.

### GetDynamicUpdateEnabled

`func (o *Zone) GetDynamicUpdateEnabled() bool`

GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field if non-nil, zero value otherwise.

### GetDynamicUpdateEnabledOk

`func (o *Zone) GetDynamicUpdateEnabledOk() (*bool, bool)`

GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicUpdateEnabled

`func (o *Zone) SetDynamicUpdateEnabled(v bool)`

SetDynamicUpdateEnabled sets DynamicUpdateEnabled field to given value.

### HasDynamicUpdateEnabled

`func (o *Zone) HasDynamicUpdateEnabled() bool`

HasDynamicUpdateEnabled returns a boolean if a field has been set.

### GetTemplate

`func (o *Zone) GetTemplate() InlinedZoneTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Zone) GetTemplateOk() (*InlinedZoneTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Zone) SetTemplate(v InlinedZoneTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Zone) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSigned

`func (o *Zone) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *Zone) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *Zone) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *Zone) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSigningPolicy

`func (o *Zone) GetSigningPolicy() InlinedDNSSECSigningPolicy`

GetSigningPolicy returns the SigningPolicy field if non-nil, zero value otherwise.

### GetSigningPolicyOk

`func (o *Zone) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool)`

GetSigningPolicyOk returns a tuple with the SigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPolicy

`func (o *Zone) SetSigningPolicy(v InlinedDNSSECSigningPolicy)`

SetSigningPolicy sets SigningPolicy field to given value.

### HasSigningPolicy

`func (o *Zone) HasSigningPolicy() bool`

HasSigningPolicy returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *Zone) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *Zone) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *Zone) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *Zone) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


