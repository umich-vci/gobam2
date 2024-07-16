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

// checks if the GetUserTasks200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserTasks200ResponseDataInner{}

// GetUserTasks200ResponseDataInner struct for GetUserTasks200ResponseDataInner
type GetUserTasks200ResponseDataInner struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name or description for the task.
	Name *string `json:"name,omitempty" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	// The priority level of the task.
	Priority *string `json:"priority,omitempty"`
	// The current status of the task.
	State *string `json:"state,omitempty"`
	// The current progress percentage of the task.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// The start date for the task.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The due date for the task.
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// Custom comments for the task.
	Comment *string `json:"comment,omitempty"`
	Links *ResourceLinks `json:"_links,omitempty"`
}

// NewGetUserTasks200ResponseDataInner instantiates a new GetUserTasks200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserTasks200ResponseDataInner() *GetUserTasks200ResponseDataInner {
	this := GetUserTasks200ResponseDataInner{}
	return &this
}

// NewGetUserTasks200ResponseDataInnerWithDefaults instantiates a new GetUserTasks200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserTasks200ResponseDataInnerWithDefaults() *GetUserTasks200ResponseDataInner {
	this := GetUserTasks200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetUserTasks200ResponseDataInner) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetUserTasks200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetUserTasks200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *GetUserTasks200ResponseDataInner) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *GetUserTasks200ResponseDataInner) SetPriority(v string) {
	o.Priority = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetUserTasks200ResponseDataInner) SetState(v string) {
	o.State = &v
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetPercentComplete() int32 {
	if o == nil || IsNil(o.PercentComplete) {
		var ret int32
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetPercentCompleteOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentComplete) {
		return nil, false
	}
	return o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasPercentComplete() bool {
	if o != nil && !IsNil(o.PercentComplete) {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given int32 and assigns it to the PercentComplete field.
func (o *GetUserTasks200ResponseDataInner) SetPercentComplete(v int32) {
	o.PercentComplete = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetStartDateTime() time.Time {
	if o == nil || IsNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *GetUserTasks200ResponseDataInner) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetDueDateTime returns the DueDateTime field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetDueDateTime() time.Time {
	if o == nil || IsNil(o.DueDateTime) {
		var ret time.Time
		return ret
	}
	return *o.DueDateTime
}

// GetDueDateTimeOk returns a tuple with the DueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetDueDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDateTime) {
		return nil, false
	}
	return o.DueDateTime, true
}

// HasDueDateTime returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasDueDateTime() bool {
	if o != nil && !IsNil(o.DueDateTime) {
		return true
	}

	return false
}

// SetDueDateTime gets a reference to the given time.Time and assigns it to the DueDateTime field.
func (o *GetUserTasks200ResponseDataInner) SetDueDateTime(v time.Time) {
	o.DueDateTime = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *GetUserTasks200ResponseDataInner) SetComment(v string) {
	o.Comment = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetUserTasks200ResponseDataInner) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTasks200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetUserTasks200ResponseDataInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GetUserTasks200ResponseDataInner) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GetUserTasks200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserTasks200ResponseDataInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PercentComplete) {
		toSerialize["percentComplete"] = o.PercentComplete
	}
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.DueDateTime) {
		toSerialize["dueDateTime"] = o.DueDateTime
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetUserTasks200ResponseDataInner struct {
	value *GetUserTasks200ResponseDataInner
	isSet bool
}

func (v NullableGetUserTasks200ResponseDataInner) Get() *GetUserTasks200ResponseDataInner {
	return v.value
}

func (v *NullableGetUserTasks200ResponseDataInner) Set(val *GetUserTasks200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserTasks200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserTasks200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserTasks200ResponseDataInner(val *GetUserTasks200ResponseDataInner) *NullableGetUserTasks200ResponseDataInner {
	return &NullableGetUserTasks200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetUserTasks200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserTasks200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

