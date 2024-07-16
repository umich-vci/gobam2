/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
)

// checks if the GetCollectionZones200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCollectionZones200ResponseDataInner{}

// GetCollectionZones200ResponseDataInner struct for GetCollectionZones200ResponseDataInner
type GetCollectionZones200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	View *View `json:"view,omitempty"`
	DeploymentEnabled *bool `json:"deploymentEnabled,omitempty"`
	// Indicates whether the dynamic update feature is enabled on the zone. When enabled, any resource records that are added, updated, or deleted within the zone will be automatically deployed to the primary DNS/DHCP Server of that zone using selective deployment.
	DynamicUpdateEnabled *bool `json:"dynamicUpdateEnabled,omitempty"`
	Template *InlinedZoneTemplate `json:"template,omitempty"`
	// Indicates if the zone is currently signed using a DNSSEC Signing Policy.
	Signed *bool `json:"signed,omitempty"`
	SigningPolicy *InlinedDNSSECSigningPolicy `json:"signingPolicy,omitempty"`
	AbsoluteName *string `json:"absoluteName,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetCollectionZones200ResponseDataInner instantiates a new GetCollectionZones200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCollectionZones200ResponseDataInner() *GetCollectionZones200ResponseDataInner {
	this := GetCollectionZones200ResponseDataInner{}
	return &this
}

// NewGetCollectionZones200ResponseDataInnerWithDefaults instantiates a new GetCollectionZones200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCollectionZones200ResponseDataInnerWithDefaults() *GetCollectionZones200ResponseDataInner {
	this := GetCollectionZones200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetCollectionZones200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetCollectionZones200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCollectionZones200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *GetCollectionZones200ResponseDataInner) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *GetCollectionZones200ResponseDataInner) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetView() View {
	if o == nil || IsNil(o.View) {
		var ret View
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetViewOk() (*View, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given View and assigns it to the View field.
func (o *GetCollectionZones200ResponseDataInner) SetView(v View) {
	o.View = &v
}

// GetDeploymentEnabled returns the DeploymentEnabled field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetDeploymentEnabled() bool {
	if o == nil || IsNil(o.DeploymentEnabled) {
		var ret bool
		return ret
	}
	return *o.DeploymentEnabled
}

// GetDeploymentEnabledOk returns a tuple with the DeploymentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetDeploymentEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeploymentEnabled) {
		return nil, false
	}
	return o.DeploymentEnabled, true
}

// HasDeploymentEnabled returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasDeploymentEnabled() bool {
	if o != nil && !IsNil(o.DeploymentEnabled) {
		return true
	}

	return false
}

// SetDeploymentEnabled gets a reference to the given bool and assigns it to the DeploymentEnabled field.
func (o *GetCollectionZones200ResponseDataInner) SetDeploymentEnabled(v bool) {
	o.DeploymentEnabled = &v
}

// GetDynamicUpdateEnabled returns the DynamicUpdateEnabled field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetDynamicUpdateEnabled() bool {
	if o == nil || IsNil(o.DynamicUpdateEnabled) {
		var ret bool
		return ret
	}
	return *o.DynamicUpdateEnabled
}

// GetDynamicUpdateEnabledOk returns a tuple with the DynamicUpdateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetDynamicUpdateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DynamicUpdateEnabled) {
		return nil, false
	}
	return o.DynamicUpdateEnabled, true
}

// HasDynamicUpdateEnabled returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasDynamicUpdateEnabled() bool {
	if o != nil && !IsNil(o.DynamicUpdateEnabled) {
		return true
	}

	return false
}

// SetDynamicUpdateEnabled gets a reference to the given bool and assigns it to the DynamicUpdateEnabled field.
func (o *GetCollectionZones200ResponseDataInner) SetDynamicUpdateEnabled(v bool) {
	o.DynamicUpdateEnabled = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetTemplate() InlinedZoneTemplate {
	if o == nil || IsNil(o.Template) {
		var ret InlinedZoneTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetTemplateOk() (*InlinedZoneTemplate, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given InlinedZoneTemplate and assigns it to the Template field.
func (o *GetCollectionZones200ResponseDataInner) SetTemplate(v InlinedZoneTemplate) {
	o.Template = &v
}

// GetSigned returns the Signed field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetSigned() bool {
	if o == nil || IsNil(o.Signed) {
		var ret bool
		return ret
	}
	return *o.Signed
}

// GetSignedOk returns a tuple with the Signed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.Signed) {
		return nil, false
	}
	return o.Signed, true
}

// HasSigned returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasSigned() bool {
	if o != nil && !IsNil(o.Signed) {
		return true
	}

	return false
}

// SetSigned gets a reference to the given bool and assigns it to the Signed field.
func (o *GetCollectionZones200ResponseDataInner) SetSigned(v bool) {
	o.Signed = &v
}

// GetSigningPolicy returns the SigningPolicy field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetSigningPolicy() InlinedDNSSECSigningPolicy {
	if o == nil || IsNil(o.SigningPolicy) {
		var ret InlinedDNSSECSigningPolicy
		return ret
	}
	return *o.SigningPolicy
}

// GetSigningPolicyOk returns a tuple with the SigningPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetSigningPolicyOk() (*InlinedDNSSECSigningPolicy, bool) {
	if o == nil || IsNil(o.SigningPolicy) {
		return nil, false
	}
	return o.SigningPolicy, true
}

// HasSigningPolicy returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasSigningPolicy() bool {
	if o != nil && !IsNil(o.SigningPolicy) {
		return true
	}

	return false
}

// SetSigningPolicy gets a reference to the given InlinedDNSSECSigningPolicy and assigns it to the SigningPolicy field.
func (o *GetCollectionZones200ResponseDataInner) SetSigningPolicy(v InlinedDNSSECSigningPolicy) {
	o.SigningPolicy = &v
}

// GetAbsoluteName returns the AbsoluteName field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetAbsoluteName() string {
	if o == nil || IsNil(o.AbsoluteName) {
		var ret string
		return ret
	}
	return *o.AbsoluteName
}

// GetAbsoluteNameOk returns a tuple with the AbsoluteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetAbsoluteNameOk() (*string, bool) {
	if o == nil || IsNil(o.AbsoluteName) {
		return nil, false
	}
	return o.AbsoluteName, true
}

// HasAbsoluteName returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasAbsoluteName() bool {
	if o != nil && !IsNil(o.AbsoluteName) {
		return true
	}

	return false
}

// SetAbsoluteName gets a reference to the given string and assigns it to the AbsoluteName field.
func (o *GetCollectionZones200ResponseDataInner) SetAbsoluteName(v string) {
	o.AbsoluteName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetCollectionZones200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCollectionZones200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetCollectionZones200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetCollectionZones200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetCollectionZones200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCollectionZones200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.DeploymentEnabled) {
		toSerialize["deploymentEnabled"] = o.DeploymentEnabled
	}
	if !IsNil(o.DynamicUpdateEnabled) {
		toSerialize["dynamicUpdateEnabled"] = o.DynamicUpdateEnabled
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Signed) {
		toSerialize["signed"] = o.Signed
	}
	if !IsNil(o.SigningPolicy) {
		toSerialize["signingPolicy"] = o.SigningPolicy
	}
	if !IsNil(o.AbsoluteName) {
		toSerialize["absoluteName"] = o.AbsoluteName
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetCollectionZones200ResponseDataInner struct {
	value *GetCollectionZones200ResponseDataInner
	isSet bool
}

func (v NullableGetCollectionZones200ResponseDataInner) Get() *GetCollectionZones200ResponseDataInner {
	return v.value
}

func (v *NullableGetCollectionZones200ResponseDataInner) Set(val *GetCollectionZones200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCollectionZones200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCollectionZones200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCollectionZones200ResponseDataInner(val *GetCollectionZones200ResponseDataInner) *NullableGetCollectionZones200ResponseDataInner {
	return &NullableGetCollectionZones200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetCollectionZones200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCollectionZones200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


