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

// checks if the TFTPFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TFTPFile{}

// TFTPFile struct for TFTPFile
type TFTPFile struct {
	FileInfo
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The version of the TFTP file.
	Version *string `json:"version,omitempty"`
	// The description of the TFTP file.
	Description *string `json:"description,omitempty"`
	// The size of the TFTP file, in bytes.
	Size *int64 `json:"size,omitempty"`
	// The MD5 checksum of the TFTP file.
	Md5Checksum *string `json:"md5Checksum,omitempty"`
}

// NewTFTPFile instantiates a new TFTPFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTFTPFile() *TFTPFile {
	this := TFTPFile{}
	return &this
}

// NewTFTPFileWithDefaults instantiates a new TFTPFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTFTPFileWithDefaults() *TFTPFile {
	this := TFTPFile{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TFTPFile) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPFile) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TFTPFile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TFTPFile) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TFTPFile) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPFile) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TFTPFile) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *TFTPFile) SetVersion(v string) {
	o.Version = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TFTPFile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPFile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TFTPFile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TFTPFile) SetDescription(v string) {
	o.Description = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TFTPFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TFTPFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *TFTPFile) SetSize(v int64) {
	o.Size = &v
}

// GetMd5Checksum returns the Md5Checksum field value if set, zero value otherwise.
func (o *TFTPFile) GetMd5Checksum() string {
	if o == nil || IsNil(o.Md5Checksum) {
		var ret string
		return ret
	}
	return *o.Md5Checksum
}

// GetMd5ChecksumOk returns a tuple with the Md5Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFTPFile) GetMd5ChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Md5Checksum) {
		return nil, false
	}
	return o.Md5Checksum, true
}

// HasMd5Checksum returns a boolean if a field has been set.
func (o *TFTPFile) HasMd5Checksum() bool {
	if o != nil && !IsNil(o.Md5Checksum) {
		return true
	}

	return false
}

// SetMd5Checksum gets a reference to the given string and assigns it to the Md5Checksum field.
func (o *TFTPFile) SetMd5Checksum(v string) {
	o.Md5Checksum = &v
}

func (o TFTPFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TFTPFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Md5Checksum) {
		toSerialize["md5Checksum"] = o.Md5Checksum
	}
	return toSerialize, nil
}

type NullableTFTPFile struct {
	value *TFTPFile
	isSet bool
}

func (v NullableTFTPFile) Get() *TFTPFile {
	return v.value
}

func (v *NullableTFTPFile) Set(val *TFTPFile) {
	v.value = val
	v.isSet = true
}

func (v NullableTFTPFile) IsSet() bool {
	return v.isSet
}

func (v *NullableTFTPFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTFTPFile(val *TFTPFile) *NullableTFTPFile {
	return &NullableTFTPFile{value: val, isSet: true}
}

func (v NullableTFTPFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTFTPFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


