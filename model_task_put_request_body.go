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
	"bytes"
	"fmt"
)

// checks if the TaskPutRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskPutRequestBody{}

// TaskPutRequestBody struct for TaskPutRequestBody
type TaskPutRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The name or description for the task.
	Name string `json:"name" validate:"regexp=^.*\\\\S+.*$"`
	// User-defined fields set for the resource.
	UserDefinedFields *map[string]string `json:"userDefinedFields,omitempty"`
	// The priority level of the task.
	Priority string `json:"priority"`
	// The current status of the task.
	State string `json:"state"`
	// The current progress percentage of the task.
	PercentComplete int32 `json:"percentComplete"`
	// The start date for the task.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The due date for the task.
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// Custom comments for the task.
	Comment *string `json:"comment,omitempty"`
}

type _TaskPutRequestBody TaskPutRequestBody

// NewTaskPutRequestBody instantiates a new TaskPutRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskPutRequestBody(name string, priority string, state string, percentComplete int32) *TaskPutRequestBody {
	this := TaskPutRequestBody{}
	this.Name = name
	this.Priority = priority
	this.State = state
	this.PercentComplete = percentComplete
	return &this
}

// NewTaskPutRequestBodyWithDefaults instantiates a new TaskPutRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskPutRequestBodyWithDefaults() *TaskPutRequestBody {
	this := TaskPutRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskPutRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskPutRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TaskPutRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskPutRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskPutRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaskPutRequestBody) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *TaskPutRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaskPutRequestBody) SetName(v string) {
	o.Name = v
}

// GetUserDefinedFields returns the UserDefinedFields field value if set, zero value otherwise.
func (o *TaskPutRequestBody) GetUserDefinedFields() map[string]string {
	if o == nil || IsNil(o.UserDefinedFields) {
		var ret map[string]string
		return ret
	}
	return *o.UserDefinedFields
}

// GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.UserDefinedFields) {
		return nil, false
	}
	return o.UserDefinedFields, true
}

// HasUserDefinedFields returns a boolean if a field has been set.
func (o *TaskPutRequestBody) HasUserDefinedFields() bool {
	if o != nil && !IsNil(o.UserDefinedFields) {
		return true
	}

	return false
}

// SetUserDefinedFields gets a reference to the given map[string]string and assigns it to the UserDefinedFields field.
func (o *TaskPutRequestBody) SetUserDefinedFields(v map[string]string) {
	o.UserDefinedFields = &v
}

// GetPriority returns the Priority field value
func (o *TaskPutRequestBody) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *TaskPutRequestBody) SetPriority(v string) {
	o.Priority = v
}

// GetState returns the State field value
func (o *TaskPutRequestBody) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TaskPutRequestBody) SetState(v string) {
	o.State = v
}

// GetPercentComplete returns the PercentComplete field value
func (o *TaskPutRequestBody) GetPercentComplete() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetPercentCompleteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PercentComplete, true
}

// SetPercentComplete sets field value
func (o *TaskPutRequestBody) SetPercentComplete(v int32) {
	o.PercentComplete = v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *TaskPutRequestBody) GetStartDateTime() time.Time {
	if o == nil || IsNil(o.StartDateTime) {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDateTime) {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *TaskPutRequestBody) HasStartDateTime() bool {
	if o != nil && !IsNil(o.StartDateTime) {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *TaskPutRequestBody) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetDueDateTime returns the DueDateTime field value if set, zero value otherwise.
func (o *TaskPutRequestBody) GetDueDateTime() time.Time {
	if o == nil || IsNil(o.DueDateTime) {
		var ret time.Time
		return ret
	}
	return *o.DueDateTime
}

// GetDueDateTimeOk returns a tuple with the DueDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetDueDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDateTime) {
		return nil, false
	}
	return o.DueDateTime, true
}

// HasDueDateTime returns a boolean if a field has been set.
func (o *TaskPutRequestBody) HasDueDateTime() bool {
	if o != nil && !IsNil(o.DueDateTime) {
		return true
	}

	return false
}

// SetDueDateTime gets a reference to the given time.Time and assigns it to the DueDateTime field.
func (o *TaskPutRequestBody) SetDueDateTime(v time.Time) {
	o.DueDateTime = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TaskPutRequestBody) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskPutRequestBody) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TaskPutRequestBody) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TaskPutRequestBody) SetComment(v string) {
	o.Comment = &v
}

func (o TaskPutRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskPutRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.UserDefinedFields) {
		toSerialize["userDefinedFields"] = o.UserDefinedFields
	}
	toSerialize["priority"] = o.Priority
	toSerialize["state"] = o.State
	toSerialize["percentComplete"] = o.PercentComplete
	if !IsNil(o.StartDateTime) {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if !IsNil(o.DueDateTime) {
		toSerialize["dueDateTime"] = o.DueDateTime
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

func (o *TaskPutRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"priority",
		"state",
		"percentComplete",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTaskPutRequestBody := _TaskPutRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaskPutRequestBody)

	if err != nil {
		return err
	}

	*o = TaskPutRequestBody(varTaskPutRequestBody)

	return err
}

type NullableTaskPutRequestBody struct {
	value *TaskPutRequestBody
	isSet bool
}

func (v NullableTaskPutRequestBody) Get() *TaskPutRequestBody {
	return v.value
}

func (v *NullableTaskPutRequestBody) Set(val *TaskPutRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskPutRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskPutRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskPutRequestBody(val *TaskPutRequestBody) *NullableTaskPutRequestBody {
	return &NullableTaskPutRequestBody{value: val, isSet: true}
}

func (v NullableTaskPutRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskPutRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

