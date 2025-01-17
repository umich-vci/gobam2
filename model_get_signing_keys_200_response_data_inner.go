/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"time"
)

// checks if the GetSigningKeys200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSigningKeys200ResponseDataInner{}

// GetSigningKeys200ResponseDataInner struct for GetSigningKeys200ResponseDataInner
type GetSigningKeys200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	Configuration *InlinedConfiguration `json:"configuration,omitempty"`
	// The key length, in bits.
	Length *int32 `json:"length,omitempty"`
	// The state of the key.
	State *string `json:"state,omitempty"`
	// The date and time that they key was generated.
	KeyGenerationDateTime *time.Time `json:"keyGenerationDateTime,omitempty"`
	// The private key.
	PrivateKey *string `json:"privateKey,omitempty"`
	// The algorithm used to sign the key.
	Algorithm *string `json:"algorithm,omitempty"`
	// The date and time for the beginning of the key's validity period.
	KeyStartDateTime *time.Time `json:"keyStartDateTime,omitempty"`
	// The date and time at which the key expires.
	KeyExpirationDateTime *time.Time `json:"keyExpirationDateTime,omitempty"`
	// The TTL (time to live) for the key if an override TTL is specified when the key is created.
	Ttl *string `json:"ttl,omitempty"`
	// The key tag value for the key. The key tag is used during DNSSEC validation and when signing and re-signing zones.
	KeyTag *int32 `json:"keyTag,omitempty"`
	// The public key text.
	PublicKey *string `json:"publicKey,omitempty"`
	// The flag used to verify the DNS record signature for resource records.
	Flags *int32 `json:"flags,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetSigningKeys200ResponseDataInner instantiates a new GetSigningKeys200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSigningKeys200ResponseDataInner() *GetSigningKeys200ResponseDataInner {
	this := GetSigningKeys200ResponseDataInner{}
	return &this
}

// NewGetSigningKeys200ResponseDataInnerWithDefaults instantiates a new GetSigningKeys200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSigningKeys200ResponseDataInnerWithDefaults() *GetSigningKeys200ResponseDataInner {
	this := GetSigningKeys200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetSigningKeys200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetSigningKeys200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetSigningKeys200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *GetSigningKeys200ResponseDataInner) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetConfiguration() InlinedConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret InlinedConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given InlinedConfiguration and assigns it to the Configuration field.
func (o *GetSigningKeys200ResponseDataInner) SetConfiguration(v InlinedConfiguration) {
	o.Configuration = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *GetSigningKeys200ResponseDataInner) SetLength(v int32) {
	o.Length = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetSigningKeys200ResponseDataInner) SetState(v string) {
	o.State = &v
}

// GetKeyGenerationDateTime returns the KeyGenerationDateTime field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetKeyGenerationDateTime() time.Time {
	if o == nil || IsNil(o.KeyGenerationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.KeyGenerationDateTime
}

// GetKeyGenerationDateTimeOk returns a tuple with the KeyGenerationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetKeyGenerationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.KeyGenerationDateTime) {
		return nil, false
	}
	return o.KeyGenerationDateTime, true
}

// HasKeyGenerationDateTime returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasKeyGenerationDateTime() bool {
	if o != nil && !IsNil(o.KeyGenerationDateTime) {
		return true
	}

	return false
}

// SetKeyGenerationDateTime gets a reference to the given time.Time and assigns it to the KeyGenerationDateTime field.
func (o *GetSigningKeys200ResponseDataInner) SetKeyGenerationDateTime(v time.Time) {
	o.KeyGenerationDateTime = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *GetSigningKeys200ResponseDataInner) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *GetSigningKeys200ResponseDataInner) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetKeyStartDateTime returns the KeyStartDateTime field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetKeyStartDateTime() time.Time {
	if o == nil || IsNil(o.KeyStartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.KeyStartDateTime
}

// GetKeyStartDateTimeOk returns a tuple with the KeyStartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetKeyStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.KeyStartDateTime) {
		return nil, false
	}
	return o.KeyStartDateTime, true
}

// HasKeyStartDateTime returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasKeyStartDateTime() bool {
	if o != nil && !IsNil(o.KeyStartDateTime) {
		return true
	}

	return false
}

// SetKeyStartDateTime gets a reference to the given time.Time and assigns it to the KeyStartDateTime field.
func (o *GetSigningKeys200ResponseDataInner) SetKeyStartDateTime(v time.Time) {
	o.KeyStartDateTime = &v
}

// GetKeyExpirationDateTime returns the KeyExpirationDateTime field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetKeyExpirationDateTime() time.Time {
	if o == nil || IsNil(o.KeyExpirationDateTime) {
		var ret time.Time
		return ret
	}
	return *o.KeyExpirationDateTime
}

// GetKeyExpirationDateTimeOk returns a tuple with the KeyExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetKeyExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.KeyExpirationDateTime) {
		return nil, false
	}
	return o.KeyExpirationDateTime, true
}

// HasKeyExpirationDateTime returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasKeyExpirationDateTime() bool {
	if o != nil && !IsNil(o.KeyExpirationDateTime) {
		return true
	}

	return false
}

// SetKeyExpirationDateTime gets a reference to the given time.Time and assigns it to the KeyExpirationDateTime field.
func (o *GetSigningKeys200ResponseDataInner) SetKeyExpirationDateTime(v time.Time) {
	o.KeyExpirationDateTime = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *GetSigningKeys200ResponseDataInner) SetTtl(v string) {
	o.Ttl = &v
}

// GetKeyTag returns the KeyTag field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetKeyTag() int32 {
	if o == nil || IsNil(o.KeyTag) {
		var ret int32
		return ret
	}
	return *o.KeyTag
}

// GetKeyTagOk returns a tuple with the KeyTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetKeyTagOk() (*int32, bool) {
	if o == nil || IsNil(o.KeyTag) {
		return nil, false
	}
	return o.KeyTag, true
}

// HasKeyTag returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasKeyTag() bool {
	if o != nil && !IsNil(o.KeyTag) {
		return true
	}

	return false
}

// SetKeyTag gets a reference to the given int32 and assigns it to the KeyTag field.
func (o *GetSigningKeys200ResponseDataInner) SetKeyTag(v int32) {
	o.KeyTag = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *GetSigningKeys200ResponseDataInner) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *GetSigningKeys200ResponseDataInner) SetFlags(v int32) {
	o.Flags = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetSigningKeys200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSigningKeys200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetSigningKeys200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetSigningKeys200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetSigningKeys200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSigningKeys200ResponseDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.KeyGenerationDateTime) {
		toSerialize["keyGenerationDateTime"] = o.KeyGenerationDateTime
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.KeyStartDateTime) {
		toSerialize["keyStartDateTime"] = o.KeyStartDateTime
	}
	if !IsNil(o.KeyExpirationDateTime) {
		toSerialize["keyExpirationDateTime"] = o.KeyExpirationDateTime
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.KeyTag) {
		toSerialize["keyTag"] = o.KeyTag
	}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetSigningKeys200ResponseDataInner struct {
	value *GetSigningKeys200ResponseDataInner
	isSet bool
}

func (v NullableGetSigningKeys200ResponseDataInner) Get() *GetSigningKeys200ResponseDataInner {
	return v.value
}

func (v *NullableGetSigningKeys200ResponseDataInner) Set(val *GetSigningKeys200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSigningKeys200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSigningKeys200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSigningKeys200ResponseDataInner(val *GetSigningKeys200ResponseDataInner) *NullableGetSigningKeys200ResponseDataInner {
	return &NullableGetSigningKeys200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetSigningKeys200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSigningKeys200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


