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

// checks if the SNMPService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SNMPService{}

// SNMPService struct for SNMPService
type SNMPService struct {
	Service
	// The resource type.
	Type *string `json:"type,omitempty"`
	// Indicates whether the SNMP service is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates whether the SNMP service configuration has been manually overridden.
	ManualOverrideEnabled *bool `json:"manualOverrideEnabled,omitempty"`
	// The system name to be reported through SNMP.
	ContactName *string `json:"contactName,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The email address for the system contact to be reported through SNMP.
	ContactEmail *string `json:"contactEmail,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The description of the system's location to be reported through SNMP.
	Location *string `json:"location,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The description of the system to be reported through SNMP.
	Description *string `json:"description,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// The SNMP polling interval, in ISO 8601 duration format, that determines the frequency that SNMP values are refreshed on the system.
	PollingInterval *string `json:"pollingInterval,omitempty"`
	// The logging level for SNMP events.
	LogLevel *string `json:"logLevel,omitempty"`
	Snmpv1 *Snmpv1Bean `json:"snmpv1,omitempty"`
	Snmpv2c *Snmpv2cBean `json:"snmpv2c,omitempty"`
	Snmpv3 *Snmpv3Bean `json:"snmpv3,omitempty"`
	TrapServers []SnmpTrapServerBean `json:"trapServers,omitempty"`
}

// NewSNMPService instantiates a new SNMPService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSNMPService() *SNMPService {
	this := SNMPService{}
	return &this
}

// NewSNMPServiceWithDefaults instantiates a new SNMPService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSNMPServiceWithDefaults() *SNMPService {
	this := SNMPService{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SNMPService) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SNMPService) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SNMPService) SetType(v string) {
	o.Type = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SNMPService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SNMPService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SNMPService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetManualOverrideEnabled returns the ManualOverrideEnabled field value if set, zero value otherwise.
func (o *SNMPService) GetManualOverrideEnabled() bool {
	if o == nil || IsNil(o.ManualOverrideEnabled) {
		var ret bool
		return ret
	}
	return *o.ManualOverrideEnabled
}

// GetManualOverrideEnabledOk returns a tuple with the ManualOverrideEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetManualOverrideEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualOverrideEnabled) {
		return nil, false
	}
	return o.ManualOverrideEnabled, true
}

// HasManualOverrideEnabled returns a boolean if a field has been set.
func (o *SNMPService) HasManualOverrideEnabled() bool {
	if o != nil && !IsNil(o.ManualOverrideEnabled) {
		return true
	}

	return false
}

// SetManualOverrideEnabled gets a reference to the given bool and assigns it to the ManualOverrideEnabled field.
func (o *SNMPService) SetManualOverrideEnabled(v bool) {
	o.ManualOverrideEnabled = &v
}

// GetContactName returns the ContactName field value if set, zero value otherwise.
func (o *SNMPService) GetContactName() string {
	if o == nil || IsNil(o.ContactName) {
		var ret string
		return ret
	}
	return *o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.ContactName) {
		return nil, false
	}
	return o.ContactName, true
}

// HasContactName returns a boolean if a field has been set.
func (o *SNMPService) HasContactName() bool {
	if o != nil && !IsNil(o.ContactName) {
		return true
	}

	return false
}

// SetContactName gets a reference to the given string and assigns it to the ContactName field.
func (o *SNMPService) SetContactName(v string) {
	o.ContactName = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *SNMPService) GetContactEmail() string {
	if o == nil || IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetContactEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *SNMPService) HasContactEmail() bool {
	if o != nil && !IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *SNMPService) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SNMPService) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SNMPService) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *SNMPService) SetLocation(v string) {
	o.Location = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SNMPService) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SNMPService) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SNMPService) SetDescription(v string) {
	o.Description = &v
}

// GetPollingInterval returns the PollingInterval field value if set, zero value otherwise.
func (o *SNMPService) GetPollingInterval() string {
	if o == nil || IsNil(o.PollingInterval) {
		var ret string
		return ret
	}
	return *o.PollingInterval
}

// GetPollingIntervalOk returns a tuple with the PollingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetPollingIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.PollingInterval) {
		return nil, false
	}
	return o.PollingInterval, true
}

// HasPollingInterval returns a boolean if a field has been set.
func (o *SNMPService) HasPollingInterval() bool {
	if o != nil && !IsNil(o.PollingInterval) {
		return true
	}

	return false
}

// SetPollingInterval gets a reference to the given string and assigns it to the PollingInterval field.
func (o *SNMPService) SetPollingInterval(v string) {
	o.PollingInterval = &v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *SNMPService) GetLogLevel() string {
	if o == nil || IsNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetLogLevelOk() (*string, bool) {
	if o == nil || IsNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *SNMPService) HasLogLevel() bool {
	if o != nil && !IsNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *SNMPService) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetSnmpv1 returns the Snmpv1 field value if set, zero value otherwise.
func (o *SNMPService) GetSnmpv1() Snmpv1Bean {
	if o == nil || IsNil(o.Snmpv1) {
		var ret Snmpv1Bean
		return ret
	}
	return *o.Snmpv1
}

// GetSnmpv1Ok returns a tuple with the Snmpv1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetSnmpv1Ok() (*Snmpv1Bean, bool) {
	if o == nil || IsNil(o.Snmpv1) {
		return nil, false
	}
	return o.Snmpv1, true
}

// HasSnmpv1 returns a boolean if a field has been set.
func (o *SNMPService) HasSnmpv1() bool {
	if o != nil && !IsNil(o.Snmpv1) {
		return true
	}

	return false
}

// SetSnmpv1 gets a reference to the given Snmpv1Bean and assigns it to the Snmpv1 field.
func (o *SNMPService) SetSnmpv1(v Snmpv1Bean) {
	o.Snmpv1 = &v
}

// GetSnmpv2c returns the Snmpv2c field value if set, zero value otherwise.
func (o *SNMPService) GetSnmpv2c() Snmpv2cBean {
	if o == nil || IsNil(o.Snmpv2c) {
		var ret Snmpv2cBean
		return ret
	}
	return *o.Snmpv2c
}

// GetSnmpv2cOk returns a tuple with the Snmpv2c field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetSnmpv2cOk() (*Snmpv2cBean, bool) {
	if o == nil || IsNil(o.Snmpv2c) {
		return nil, false
	}
	return o.Snmpv2c, true
}

// HasSnmpv2c returns a boolean if a field has been set.
func (o *SNMPService) HasSnmpv2c() bool {
	if o != nil && !IsNil(o.Snmpv2c) {
		return true
	}

	return false
}

// SetSnmpv2c gets a reference to the given Snmpv2cBean and assigns it to the Snmpv2c field.
func (o *SNMPService) SetSnmpv2c(v Snmpv2cBean) {
	o.Snmpv2c = &v
}

// GetSnmpv3 returns the Snmpv3 field value if set, zero value otherwise.
func (o *SNMPService) GetSnmpv3() Snmpv3Bean {
	if o == nil || IsNil(o.Snmpv3) {
		var ret Snmpv3Bean
		return ret
	}
	return *o.Snmpv3
}

// GetSnmpv3Ok returns a tuple with the Snmpv3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetSnmpv3Ok() (*Snmpv3Bean, bool) {
	if o == nil || IsNil(o.Snmpv3) {
		return nil, false
	}
	return o.Snmpv3, true
}

// HasSnmpv3 returns a boolean if a field has been set.
func (o *SNMPService) HasSnmpv3() bool {
	if o != nil && !IsNil(o.Snmpv3) {
		return true
	}

	return false
}

// SetSnmpv3 gets a reference to the given Snmpv3Bean and assigns it to the Snmpv3 field.
func (o *SNMPService) SetSnmpv3(v Snmpv3Bean) {
	o.Snmpv3 = &v
}

// GetTrapServers returns the TrapServers field value if set, zero value otherwise.
func (o *SNMPService) GetTrapServers() []SnmpTrapServerBean {
	if o == nil || IsNil(o.TrapServers) {
		var ret []SnmpTrapServerBean
		return ret
	}
	return o.TrapServers
}

// GetTrapServersOk returns a tuple with the TrapServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SNMPService) GetTrapServersOk() ([]SnmpTrapServerBean, bool) {
	if o == nil || IsNil(o.TrapServers) {
		return nil, false
	}
	return o.TrapServers, true
}

// HasTrapServers returns a boolean if a field has been set.
func (o *SNMPService) HasTrapServers() bool {
	if o != nil && !IsNil(o.TrapServers) {
		return true
	}

	return false
}

// SetTrapServers gets a reference to the given []SnmpTrapServerBean and assigns it to the TrapServers field.
func (o *SNMPService) SetTrapServers(v []SnmpTrapServerBean) {
	o.TrapServers = v
}

func (o SNMPService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SNMPService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ManualOverrideEnabled) {
		toSerialize["manualOverrideEnabled"] = o.ManualOverrideEnabled
	}
	if !IsNil(o.ContactName) {
		toSerialize["contactName"] = o.ContactName
	}
	if !IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.PollingInterval) {
		toSerialize["pollingInterval"] = o.PollingInterval
	}
	if !IsNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !IsNil(o.Snmpv1) {
		toSerialize["snmpv1"] = o.Snmpv1
	}
	if !IsNil(o.Snmpv2c) {
		toSerialize["snmpv2c"] = o.Snmpv2c
	}
	if !IsNil(o.Snmpv3) {
		toSerialize["snmpv3"] = o.Snmpv3
	}
	if !IsNil(o.TrapServers) {
		toSerialize["trapServers"] = o.TrapServers
	}
	return toSerialize, nil
}

type NullableSNMPService struct {
	value *SNMPService
	isSet bool
}

func (v NullableSNMPService) Get() *SNMPService {
	return v.value
}

func (v *NullableSNMPService) Set(val *SNMPService) {
	v.value = val
	v.isSet = true
}

func (v NullableSNMPService) IsSet() bool {
	return v.isSet
}

func (v *NullableSNMPService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSNMPService(val *SNMPService) *NullableSNMPService {
	return &NullableSNMPService{value: val, isSet: true}
}

func (v NullableSNMPService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSNMPService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


