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

// checks if the SAMLIdentityProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLIdentityProvider{}

// SAMLIdentityProvider struct for SAMLIdentityProvider
type SAMLIdentityProvider struct {
	Authenticator
	// The type of authenticator.
	Type *string `json:"type,omitempty"`
	// The description of the SAML IdP configuration.
	Description *string `json:"description,omitempty"`
	// Indicates whether SAML IdP authentication is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The attribute name for group in the SAML response.
	GroupAttribute *string `json:"groupAttribute,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The attribute name for email in the SAML response.
	EmailAttribute *string `json:"emailAttribute,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The sign-in URL of the SAMP IdP.
	SignOnUrl *string `json:"signOnUrl,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The logout URL of the SAMP IdP.
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// The entity ID of the SAMP IdP.
	EntityId *string `json:"entityId,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The SAML IdP server signing certificate.
	SigningCertificate *string `json:"signingCertificate,omitempty"`
}

// NewSAMLIdentityProvider instantiates a new SAMLIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLIdentityProvider() *SAMLIdentityProvider {
	this := SAMLIdentityProvider{}
	return &this
}

// NewSAMLIdentityProviderWithDefaults instantiates a new SAMLIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLIdentityProviderWithDefaults() *SAMLIdentityProvider {
	this := SAMLIdentityProvider{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SAMLIdentityProvider) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SAMLIdentityProvider) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SAMLIdentityProvider) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetGroupAttribute returns the GroupAttribute field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetGroupAttribute() string {
	if o == nil || IsNil(o.GroupAttribute) {
		var ret string
		return ret
	}
	return *o.GroupAttribute
}

// GetGroupAttributeOk returns a tuple with the GroupAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetGroupAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupAttribute) {
		return nil, false
	}
	return o.GroupAttribute, true
}

// HasGroupAttribute returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasGroupAttribute() bool {
	if o != nil && !IsNil(o.GroupAttribute) {
		return true
	}

	return false
}

// SetGroupAttribute gets a reference to the given string and assigns it to the GroupAttribute field.
func (o *SAMLIdentityProvider) SetGroupAttribute(v string) {
	o.GroupAttribute = &v
}

// GetEmailAttribute returns the EmailAttribute field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetEmailAttribute() string {
	if o == nil || IsNil(o.EmailAttribute) {
		var ret string
		return ret
	}
	return *o.EmailAttribute
}

// GetEmailAttributeOk returns a tuple with the EmailAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetEmailAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttribute) {
		return nil, false
	}
	return o.EmailAttribute, true
}

// HasEmailAttribute returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasEmailAttribute() bool {
	if o != nil && !IsNil(o.EmailAttribute) {
		return true
	}

	return false
}

// SetEmailAttribute gets a reference to the given string and assigns it to the EmailAttribute field.
func (o *SAMLIdentityProvider) SetEmailAttribute(v string) {
	o.EmailAttribute = &v
}

// GetSignOnUrl returns the SignOnUrl field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetSignOnUrl() string {
	if o == nil || IsNil(o.SignOnUrl) {
		var ret string
		return ret
	}
	return *o.SignOnUrl
}

// GetSignOnUrlOk returns a tuple with the SignOnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetSignOnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SignOnUrl) {
		return nil, false
	}
	return o.SignOnUrl, true
}

// HasSignOnUrl returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasSignOnUrl() bool {
	if o != nil && !IsNil(o.SignOnUrl) {
		return true
	}

	return false
}

// SetSignOnUrl gets a reference to the given string and assigns it to the SignOnUrl field.
func (o *SAMLIdentityProvider) SetSignOnUrl(v string) {
	o.SignOnUrl = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetLogoutUrl() string {
	if o == nil || IsNil(o.LogoutUrl) {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutUrl) {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasLogoutUrl() bool {
	if o != nil && !IsNil(o.LogoutUrl) {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SAMLIdentityProvider) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SAMLIdentityProvider) SetEntityId(v string) {
	o.EntityId = &v
}

// GetSigningCertificate returns the SigningCertificate field value if set, zero value otherwise.
func (o *SAMLIdentityProvider) GetSigningCertificate() string {
	if o == nil || IsNil(o.SigningCertificate) {
		var ret string
		return ret
	}
	return *o.SigningCertificate
}

// GetSigningCertificateOk returns a tuple with the SigningCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLIdentityProvider) GetSigningCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.SigningCertificate) {
		return nil, false
	}
	return o.SigningCertificate, true
}

// HasSigningCertificate returns a boolean if a field has been set.
func (o *SAMLIdentityProvider) HasSigningCertificate() bool {
	if o != nil && !IsNil(o.SigningCertificate) {
		return true
	}

	return false
}

// SetSigningCertificate gets a reference to the given string and assigns it to the SigningCertificate field.
func (o *SAMLIdentityProvider) SetSigningCertificate(v string) {
	o.SigningCertificate = &v
}

func (o SAMLIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLIdentityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.GroupAttribute) {
		toSerialize["groupAttribute"] = o.GroupAttribute
	}
	if !IsNil(o.EmailAttribute) {
		toSerialize["emailAttribute"] = o.EmailAttribute
	}
	if !IsNil(o.SignOnUrl) {
		toSerialize["signOnUrl"] = o.SignOnUrl
	}
	if !IsNil(o.LogoutUrl) {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.SigningCertificate) {
		toSerialize["signingCertificate"] = o.SigningCertificate
	}
	return toSerialize, nil
}

type NullableSAMLIdentityProvider struct {
	value *SAMLIdentityProvider
	isSet bool
}

func (v NullableSAMLIdentityProvider) Get() *SAMLIdentityProvider {
	return v.value
}

func (v *NullableSAMLIdentityProvider) Set(val *SAMLIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLIdentityProvider(val *SAMLIdentityProvider) *NullableSAMLIdentityProvider {
	return &NullableSAMLIdentityProvider{value: val, isSet: true}
}

func (v NullableSAMLIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


