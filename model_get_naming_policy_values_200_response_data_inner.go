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

// checks if the GetNamingPolicyValues200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNamingPolicyValues200ResponseDataInner{}

// GetNamingPolicyValues200ResponseDataInner struct for GetNamingPolicyValues200ResponseDataInner
type GetNamingPolicyValues200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name of the naming policy value.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// A character used to separate elements in a filename; typically, - (hyphen) or _ (underscore) are used.
	Connector *string `json:"connector,omitempty"`
	// The incremental counter type. The value can be Counter to make the value a sequential counter or Unique Name to use the value to ensure that the names are unique. When you select Counter, the value starts at a specified value and increments each time the naming policy creates a name. When you select Unique Name, the value increments only to ensure that generated names are unique.
	IncrementalRole *string `json:"incrementalRole,omitempty"`
	// The number system for the incremental value.
	IncrementType *string `json:"incrementType,omitempty"`
	// The padding that is added to the incremental value. Simple padding pads the incremental value with a fixed number of leading zeros. Global padding pads the incremental value with leading zeros to make the entire name generated by the policy a specific length.
	PaddingType *string `json:"paddingType,omitempty"`
	// The starting value for the incremental value.
	SequenceStart *int32 `json:"sequenceStart,omitempty"`
	// The amount by which to increment the value each time it's used.
	SequenceIncrement *int32 `json:"sequenceIncrement,omitempty"`
	// The length of the padding. For simple padding, the value determines how many digits are used for the incremental value. For global padding, the value determines the overall length of the name generated by the naming policy.
	PaddingLength *int32 `json:"paddingLength,omitempty"`
	// Determines whether the naming policy reuses numeric values if they're available.
	MissingValueReuseEnabled *bool `json:"missingValueReuseEnabled,omitempty"`
	// The start value for the numeric range.
	Min *int32 `json:"min,omitempty"`
	// The end value for the numeric range. A value of 0 indicates that the range is unbounded.
	Max *int32 `json:"max,omitempty"`
	Elements []NameValuePairBean `json:"elements,omitempty"`
	TextType *string `json:"textType,omitempty"`
	// The shortest string allowed in the name. When set of 0, the text string is optional.
	MinLength *int32 `json:"minLength,omitempty"`
	// The longest string allowed in the name. When set of 0, the range is unbound.
	MaxLength *int32 `json:"maxLength,omitempty"`
	// The regular expression to restrict the text string.
	RegularExpression *string `json:"regularExpression,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetNamingPolicyValues200ResponseDataInner instantiates a new GetNamingPolicyValues200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNamingPolicyValues200ResponseDataInner() *GetNamingPolicyValues200ResponseDataInner {
	this := GetNamingPolicyValues200ResponseDataInner{}
	return &this
}

// NewGetNamingPolicyValues200ResponseDataInnerWithDefaults instantiates a new GetNamingPolicyValues200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNamingPolicyValues200ResponseDataInnerWithDefaults() *GetNamingPolicyValues200ResponseDataInner {
	this := GetNamingPolicyValues200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetConnector() string {
	if o == nil || IsNil(o.Connector) {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetConnectorOk() (*string, bool) {
	if o == nil || IsNil(o.Connector) {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasConnector() bool {
	if o != nil && !IsNil(o.Connector) {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetConnector(v string) {
	o.Connector = &v
}

// GetIncrementalRole returns the IncrementalRole field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementalRole() string {
	if o == nil || IsNil(o.IncrementalRole) {
		var ret string
		return ret
	}
	return *o.IncrementalRole
}

// GetIncrementalRoleOk returns a tuple with the IncrementalRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementalRoleOk() (*string, bool) {
	if o == nil || IsNil(o.IncrementalRole) {
		return nil, false
	}
	return o.IncrementalRole, true
}

// HasIncrementalRole returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasIncrementalRole() bool {
	if o != nil && !IsNil(o.IncrementalRole) {
		return true
	}

	return false
}

// SetIncrementalRole gets a reference to the given string and assigns it to the IncrementalRole field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetIncrementalRole(v string) {
	o.IncrementalRole = &v
}

// GetIncrementType returns the IncrementType field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementType() string {
	if o == nil || IsNil(o.IncrementType) {
		var ret string
		return ret
	}
	return *o.IncrementType
}

// GetIncrementTypeOk returns a tuple with the IncrementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetIncrementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IncrementType) {
		return nil, false
	}
	return o.IncrementType, true
}

// HasIncrementType returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasIncrementType() bool {
	if o != nil && !IsNil(o.IncrementType) {
		return true
	}

	return false
}

// SetIncrementType gets a reference to the given string and assigns it to the IncrementType field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetIncrementType(v string) {
	o.IncrementType = &v
}

// GetPaddingType returns the PaddingType field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingType() string {
	if o == nil || IsNil(o.PaddingType) {
		var ret string
		return ret
	}
	return *o.PaddingType
}

// GetPaddingTypeOk returns a tuple with the PaddingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaddingType) {
		return nil, false
	}
	return o.PaddingType, true
}

// HasPaddingType returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasPaddingType() bool {
	if o != nil && !IsNil(o.PaddingType) {
		return true
	}

	return false
}

// SetPaddingType gets a reference to the given string and assigns it to the PaddingType field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetPaddingType(v string) {
	o.PaddingType = &v
}

// GetSequenceStart returns the SequenceStart field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceStart() int32 {
	if o == nil || IsNil(o.SequenceStart) {
		var ret int32
		return ret
	}
	return *o.SequenceStart
}

// GetSequenceStartOk returns a tuple with the SequenceStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceStartOk() (*int32, bool) {
	if o == nil || IsNil(o.SequenceStart) {
		return nil, false
	}
	return o.SequenceStart, true
}

// HasSequenceStart returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasSequenceStart() bool {
	if o != nil && !IsNil(o.SequenceStart) {
		return true
	}

	return false
}

// SetSequenceStart gets a reference to the given int32 and assigns it to the SequenceStart field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetSequenceStart(v int32) {
	o.SequenceStart = &v
}

// GetSequenceIncrement returns the SequenceIncrement field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceIncrement() int32 {
	if o == nil || IsNil(o.SequenceIncrement) {
		var ret int32
		return ret
	}
	return *o.SequenceIncrement
}

// GetSequenceIncrementOk returns a tuple with the SequenceIncrement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetSequenceIncrementOk() (*int32, bool) {
	if o == nil || IsNil(o.SequenceIncrement) {
		return nil, false
	}
	return o.SequenceIncrement, true
}

// HasSequenceIncrement returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasSequenceIncrement() bool {
	if o != nil && !IsNil(o.SequenceIncrement) {
		return true
	}

	return false
}

// SetSequenceIncrement gets a reference to the given int32 and assigns it to the SequenceIncrement field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetSequenceIncrement(v int32) {
	o.SequenceIncrement = &v
}

// GetPaddingLength returns the PaddingLength field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingLength() int32 {
	if o == nil || IsNil(o.PaddingLength) {
		var ret int32
		return ret
	}
	return *o.PaddingLength
}

// GetPaddingLengthOk returns a tuple with the PaddingLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetPaddingLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PaddingLength) {
		return nil, false
	}
	return o.PaddingLength, true
}

// HasPaddingLength returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasPaddingLength() bool {
	if o != nil && !IsNil(o.PaddingLength) {
		return true
	}

	return false
}

// SetPaddingLength gets a reference to the given int32 and assigns it to the PaddingLength field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetPaddingLength(v int32) {
	o.PaddingLength = &v
}

// GetMissingValueReuseEnabled returns the MissingValueReuseEnabled field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMissingValueReuseEnabled() bool {
	if o == nil || IsNil(o.MissingValueReuseEnabled) {
		var ret bool
		return ret
	}
	return *o.MissingValueReuseEnabled
}

// GetMissingValueReuseEnabledOk returns a tuple with the MissingValueReuseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMissingValueReuseEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MissingValueReuseEnabled) {
		return nil, false
	}
	return o.MissingValueReuseEnabled, true
}

// HasMissingValueReuseEnabled returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasMissingValueReuseEnabled() bool {
	if o != nil && !IsNil(o.MissingValueReuseEnabled) {
		return true
	}

	return false
}

// SetMissingValueReuseEnabled gets a reference to the given bool and assigns it to the MissingValueReuseEnabled field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetMissingValueReuseEnabled(v bool) {
	o.MissingValueReuseEnabled = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetMax(v int32) {
	o.Max = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetElements() []NameValuePairBean {
	if o == nil || IsNil(o.Elements) {
		var ret []NameValuePairBean
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetElementsOk() ([]NameValuePairBean, bool) {
	if o == nil || IsNil(o.Elements) {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasElements() bool {
	if o != nil && !IsNil(o.Elements) {
		return true
	}

	return false
}

// SetElements gets a reference to the given []NameValuePairBean and assigns it to the Elements field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetElements(v []NameValuePairBean) {
	o.Elements = v
}

// GetTextType returns the TextType field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetTextType() string {
	if o == nil || IsNil(o.TextType) {
		var ret string
		return ret
	}
	return *o.TextType
}

// GetTextTypeOk returns a tuple with the TextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetTextTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TextType) {
		return nil, false
	}
	return o.TextType, true
}

// HasTextType returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasTextType() bool {
	if o != nil && !IsNil(o.TextType) {
		return true
	}

	return false
}

// SetTextType gets a reference to the given string and assigns it to the TextType field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetTextType(v string) {
	o.TextType = &v
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMinLength() int32 {
	if o == nil || IsNil(o.MinLength) {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMinLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MinLength) {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasMinLength() bool {
	if o != nil && !IsNil(o.MinLength) {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMaxLength() int32 {
	if o == nil || IsNil(o.MaxLength) {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetMaxLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLength) {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasMaxLength() bool {
	if o != nil && !IsNil(o.MaxLength) {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetRegularExpression returns the RegularExpression field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetRegularExpression() string {
	if o == nil || IsNil(o.RegularExpression) {
		var ret string
		return ret
	}
	return *o.RegularExpression
}

// GetRegularExpressionOk returns a tuple with the RegularExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetRegularExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.RegularExpression) {
		return nil, false
	}
	return o.RegularExpression, true
}

// HasRegularExpression returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasRegularExpression() bool {
	if o != nil && !IsNil(o.RegularExpression) {
		return true
	}

	return false
}

// SetRegularExpression gets a reference to the given string and assigns it to the RegularExpression field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetRegularExpression(v string) {
	o.RegularExpression = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetNamingPolicyValues200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetNamingPolicyValues200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetNamingPolicyValues200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetNamingPolicyValues200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNamingPolicyValues200ResponseDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Connector) {
		toSerialize["connector"] = o.Connector
	}
	if !IsNil(o.IncrementalRole) {
		toSerialize["incrementalRole"] = o.IncrementalRole
	}
	if !IsNil(o.IncrementType) {
		toSerialize["incrementType"] = o.IncrementType
	}
	if !IsNil(o.PaddingType) {
		toSerialize["paddingType"] = o.PaddingType
	}
	if !IsNil(o.SequenceStart) {
		toSerialize["sequenceStart"] = o.SequenceStart
	}
	if !IsNil(o.SequenceIncrement) {
		toSerialize["sequenceIncrement"] = o.SequenceIncrement
	}
	if !IsNil(o.PaddingLength) {
		toSerialize["paddingLength"] = o.PaddingLength
	}
	if !IsNil(o.MissingValueReuseEnabled) {
		toSerialize["missingValueReuseEnabled"] = o.MissingValueReuseEnabled
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Elements) {
		toSerialize["elements"] = o.Elements
	}
	if !IsNil(o.TextType) {
		toSerialize["textType"] = o.TextType
	}
	if !IsNil(o.MinLength) {
		toSerialize["minLength"] = o.MinLength
	}
	if !IsNil(o.MaxLength) {
		toSerialize["maxLength"] = o.MaxLength
	}
	if !IsNil(o.RegularExpression) {
		toSerialize["regularExpression"] = o.RegularExpression
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetNamingPolicyValues200ResponseDataInner struct {
	value *GetNamingPolicyValues200ResponseDataInner
	isSet bool
}

func (v NullableGetNamingPolicyValues200ResponseDataInner) Get() *GetNamingPolicyValues200ResponseDataInner {
	return v.value
}

func (v *NullableGetNamingPolicyValues200ResponseDataInner) Set(val *GetNamingPolicyValues200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNamingPolicyValues200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNamingPolicyValues200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNamingPolicyValues200ResponseDataInner(val *GetNamingPolicyValues200ResponseDataInner) *NullableGetNamingPolicyValues200ResponseDataInner {
	return &NullableGetNamingPolicyValues200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetNamingPolicyValues200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNamingPolicyValues200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


