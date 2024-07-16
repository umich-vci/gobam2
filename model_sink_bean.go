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

// checks if the SinkBean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SinkBean{}

// SinkBean struct for SinkBean
type SinkBean struct {
	// The output sink type.
	SinkType *string `json:"sinkType,omitempty"`
	// Indicates whether health check service is enabled on the downstream sink server.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty"`
	// The CA certificate used to authenticate with the remote host.
	CaCertificate *string `json:"caCertificate,omitempty"`
	// Indicates whether the server attempts a TLS handshake using the uploaded CA certificate with the remote host's TLS server certificate.
	CertificateVerificationEnabled *bool `json:"certificateVerificationEnabled,omitempty"`
	// Indicates whether the server validates the hostname part of the URI against the CN (Common Name) or SAN (Subject Alternative Name) of the server certificate during the TLS handshake.
	HostnameVerificationEnabled *bool `json:"hostnameVerificationEnabled,omitempty"`
}

// NewSinkBean instantiates a new SinkBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSinkBean() *SinkBean {
	this := SinkBean{}
	return &this
}

// NewSinkBeanWithDefaults instantiates a new SinkBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSinkBeanWithDefaults() *SinkBean {
	this := SinkBean{}
	return &this
}

// GetSinkType returns the SinkType field value if set, zero value otherwise.
func (o *SinkBean) GetSinkType() string {
	if o == nil || IsNil(o.SinkType) {
		var ret string
		return ret
	}
	return *o.SinkType
}

// GetSinkTypeOk returns a tuple with the SinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SinkBean) GetSinkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SinkType) {
		return nil, false
	}
	return o.SinkType, true
}

// HasSinkType returns a boolean if a field has been set.
func (o *SinkBean) HasSinkType() bool {
	if o != nil && !IsNil(o.SinkType) {
		return true
	}

	return false
}

// SetSinkType gets a reference to the given string and assigns it to the SinkType field.
func (o *SinkBean) SetSinkType(v string) {
	o.SinkType = &v
}

// GetHealthCheckEnabled returns the HealthCheckEnabled field value if set, zero value otherwise.
func (o *SinkBean) GetHealthCheckEnabled() bool {
	if o == nil || IsNil(o.HealthCheckEnabled) {
		var ret bool
		return ret
	}
	return *o.HealthCheckEnabled
}

// GetHealthCheckEnabledOk returns a tuple with the HealthCheckEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SinkBean) GetHealthCheckEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HealthCheckEnabled) {
		return nil, false
	}
	return o.HealthCheckEnabled, true
}

// HasHealthCheckEnabled returns a boolean if a field has been set.
func (o *SinkBean) HasHealthCheckEnabled() bool {
	if o != nil && !IsNil(o.HealthCheckEnabled) {
		return true
	}

	return false
}

// SetHealthCheckEnabled gets a reference to the given bool and assigns it to the HealthCheckEnabled field.
func (o *SinkBean) SetHealthCheckEnabled(v bool) {
	o.HealthCheckEnabled = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *SinkBean) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SinkBean) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *SinkBean) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *SinkBean) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCertificateVerificationEnabled returns the CertificateVerificationEnabled field value if set, zero value otherwise.
func (o *SinkBean) GetCertificateVerificationEnabled() bool {
	if o == nil || IsNil(o.CertificateVerificationEnabled) {
		var ret bool
		return ret
	}
	return *o.CertificateVerificationEnabled
}

// GetCertificateVerificationEnabledOk returns a tuple with the CertificateVerificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SinkBean) GetCertificateVerificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CertificateVerificationEnabled) {
		return nil, false
	}
	return o.CertificateVerificationEnabled, true
}

// HasCertificateVerificationEnabled returns a boolean if a field has been set.
func (o *SinkBean) HasCertificateVerificationEnabled() bool {
	if o != nil && !IsNil(o.CertificateVerificationEnabled) {
		return true
	}

	return false
}

// SetCertificateVerificationEnabled gets a reference to the given bool and assigns it to the CertificateVerificationEnabled field.
func (o *SinkBean) SetCertificateVerificationEnabled(v bool) {
	o.CertificateVerificationEnabled = &v
}

// GetHostnameVerificationEnabled returns the HostnameVerificationEnabled field value if set, zero value otherwise.
func (o *SinkBean) GetHostnameVerificationEnabled() bool {
	if o == nil || IsNil(o.HostnameVerificationEnabled) {
		var ret bool
		return ret
	}
	return *o.HostnameVerificationEnabled
}

// GetHostnameVerificationEnabledOk returns a tuple with the HostnameVerificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SinkBean) GetHostnameVerificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.HostnameVerificationEnabled) {
		return nil, false
	}
	return o.HostnameVerificationEnabled, true
}

// HasHostnameVerificationEnabled returns a boolean if a field has been set.
func (o *SinkBean) HasHostnameVerificationEnabled() bool {
	if o != nil && !IsNil(o.HostnameVerificationEnabled) {
		return true
	}

	return false
}

// SetHostnameVerificationEnabled gets a reference to the given bool and assigns it to the HostnameVerificationEnabled field.
func (o *SinkBean) SetHostnameVerificationEnabled(v bool) {
	o.HostnameVerificationEnabled = &v
}

func (o SinkBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SinkBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SinkType) {
		toSerialize["sinkType"] = o.SinkType
	}
	if !IsNil(o.HealthCheckEnabled) {
		toSerialize["healthCheckEnabled"] = o.HealthCheckEnabled
	}
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	if !IsNil(o.CertificateVerificationEnabled) {
		toSerialize["certificateVerificationEnabled"] = o.CertificateVerificationEnabled
	}
	if !IsNil(o.HostnameVerificationEnabled) {
		toSerialize["hostnameVerificationEnabled"] = o.HostnameVerificationEnabled
	}
	return toSerialize, nil
}

type NullableSinkBean struct {
	value *SinkBean
	isSet bool
}

func (v NullableSinkBean) Get() *SinkBean {
	return v.value
}

func (v *NullableSinkBean) Set(val *SinkBean) {
	v.value = val
	v.isSet = true
}

func (v NullableSinkBean) IsSet() bool {
	return v.isSet
}

func (v *NullableSinkBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSinkBean(val *SinkBean) *NullableSinkBean {
	return &NullableSinkBean{value: val, isSet: true}
}

func (v NullableSinkBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSinkBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

