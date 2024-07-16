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

// checks if the TemplateApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateApplication{}

// TemplateApplication struct for TemplateApplication
type TemplateApplication struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The state of the template application operation.
	State *string `json:"state,omitempty"`
	// The date and time that the template application operation started.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The date and time that the template application operation completed.
	CompletionDateTime *time.Time `json:"completionDateTime,omitempty"`
	// The method used to resolve conflicts between template items and the resource the template is applied to.
	ConflictResolutionStrategy *string `json:"conflictResolutionStrategy,omitempty"`
}

// NewTemplateApplication instantiates a new TemplateApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateApplication() *TemplateApplication {
	this := TemplateApplication{}
	return &this
}

// NewTemplateApplicationWithDefaults instantiates a new TemplateApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateApplicationWithDefaults() *TemplateApplication {
	this := TemplateApplication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TemplateApplication) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TemplateApplication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TemplateApplication) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TemplateApplication) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TemplateApplication) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TemplateApplication) SetType(v string) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TemplateApplication) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TemplateApplication) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TemplateApplication) SetState(v string) {
	o.State = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *TemplateApplication) GetStartDateTime() time.Time {
	if o == nil || IsNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *TemplateApplication) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *TemplateApplication) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetCompletionDateTime returns the CompletionDateTime field value if set, zero value otherwise.
func (o *TemplateApplication) GetCompletionDateTime() time.Time {
	if o == nil || IsNil(o.CompletionDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CompletionDateTime
}

// GetCompletionDateTimeOk returns a tuple with the CompletionDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetCompletionDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletionDateTime) {
		return nil, false
	}
	return o.CompletionDateTime, true
}

// HasCompletionDateTime returns a boolean if a field has been set.
func (o *TemplateApplication) HasCompletionDateTime() bool {
	if o != nil && !IsNil(o.CompletionDateTime) {
		return true
	}

	return false
}

// SetCompletionDateTime gets a reference to the given time.Time and assigns it to the CompletionDateTime field.
func (o *TemplateApplication) SetCompletionDateTime(v time.Time) {
	o.CompletionDateTime = &v
}

// GetConflictResolutionStrategy returns the ConflictResolutionStrategy field value if set, zero value otherwise.
func (o *TemplateApplication) GetConflictResolutionStrategy() string {
	if o == nil || IsNil(o.ConflictResolutionStrategy) {
		var ret string
		return ret
	}
	return *o.ConflictResolutionStrategy
}

// GetConflictResolutionStrategyOk returns a tuple with the ConflictResolutionStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateApplication) GetConflictResolutionStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.ConflictResolutionStrategy) {
		return nil, false
	}
	return o.ConflictResolutionStrategy, true
}

// HasConflictResolutionStrategy returns a boolean if a field has been set.
func (o *TemplateApplication) HasConflictResolutionStrategy() bool {
	if o != nil && !IsNil(o.ConflictResolutionStrategy) {
		return true
	}

	return false
}

// SetConflictResolutionStrategy gets a reference to the given string and assigns it to the ConflictResolutionStrategy field.
func (o *TemplateApplication) SetConflictResolutionStrategy(v string) {
	o.ConflictResolutionStrategy = &v
}

func (o TemplateApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.CompletionDateTime) {
		toSerialize["completionDateTime"] = o.CompletionDateTime
	}
	if !IsNil(o.ConflictResolutionStrategy) {
		toSerialize["conflictResolutionStrategy"] = o.ConflictResolutionStrategy
	}
	return toSerialize, nil
}

type NullableTemplateApplication struct {
	value *TemplateApplication
	isSet bool
}

func (v NullableTemplateApplication) Get() *TemplateApplication {
	return v.value
}

func (v *NullableTemplateApplication) Set(val *TemplateApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateApplication(val *TemplateApplication) *NullableTemplateApplication {
	return &NullableTemplateApplication{value: val, isSet: true}
}

func (v NullableTemplateApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

