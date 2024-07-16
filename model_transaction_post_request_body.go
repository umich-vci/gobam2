/*
BlueCat Address Manager 9.5 RESTful v2 API

The **BlueCat Address Manager 9.5 RESTful v2 API** is a new RESTful API for Address Manager that presents Address Manager objects as resources. Each resource has a unique endpoint, and related resources are grouped in collections. To fetch an individual resource, a `GET` request is sent to the resource's endpoint. To fetch all resources in a collection, a `GET` request is sent to the collection's endpoint.  The RESTful v2 API is [hypermedia driven](https://en.wikipedia.org/wiki/HATEOAS) and uses the [HAL](https://en.wikipedia.org/wiki/Hypertext_Application_Language) format for representing links. When navigating through the API, you can use those links to navigate to related resources or child resources of the requested endpoint. The API supports the following media types for most endpoints:  `application/hal+json`, `application/json`, and `text/csv`.  For authentication, the API supports both `Basic` and `Bearer` HTTP authentication schemes.  To get started, create a session and receive credentials for `Basic` authentication by sending a `POST` request to the [/sessions](#/Admin%20Resources/postSession) endpoint, with a message body containing the user's credentials:  ```{\"username\":\"apiuser\", \"password\":\"apipass\"}```  The response will contain an `apiToken` field that can be entered with the username in the Swagger UI **Authorize** dialog. The response will also contain a `basicAuthenticationCredentials` field containing a base64 encoding of the requester's username and API token delimited by a colon. This string can be injected directly into request `Authorization` headers in the following format:  ```Authorization: Basic YXBpOlQ0OExOT3VIRGhDcnVBNEo1bGFES3JuS3hTZC9QK3pjczZXTzBJMDY=```  For full details on API format and authentication methods, refer to the Address Manager RESTful v2 API Guide on the BlueCat Documentation Portal.

API version: 9.5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gobam2

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TransactionPostRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionPostRequestBody{}

// TransactionPostRequestBody struct for TransactionPostRequestBody
type TransactionPostRequestBody struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The type of transaction performed.
	TransactionType string `json:"transactionType"`
	// The operation performed.
	Operation *string `json:"operation,omitempty"`
	// The description of the operation performed.
	Description *string `json:"description,omitempty"`
	// The change control comment of the transaction.
	Comment *string `json:"comment,omitempty"`
	// The date and time of the transaction.
	CreationDateTime *string `json:"creationDateTime,omitempty"`
	// The unique ID of the transaction performed.
	TransactionId int64 `json:"transactionId"`
	User *User `json:"user,omitempty"`
	RevertingTransaction *Transaction `json:"revertingTransaction,omitempty"`
}

type _TransactionPostRequestBody TransactionPostRequestBody

// NewTransactionPostRequestBody instantiates a new TransactionPostRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionPostRequestBody(transactionType string, transactionId int64) *TransactionPostRequestBody {
	this := TransactionPostRequestBody{}
	this.TransactionType = transactionType
	this.TransactionId = transactionId
	return &this
}

// NewTransactionPostRequestBodyWithDefaults instantiates a new TransactionPostRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionPostRequestBodyWithDefaults() *TransactionPostRequestBody {
	this := TransactionPostRequestBody{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TransactionPostRequestBody) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransactionPostRequestBody) SetType(v string) {
	o.Type = &v
}

// GetTransactionType returns the TransactionType field value
func (o *TransactionPostRequestBody) GetTransactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetTransactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionType, true
}

// SetTransactionType sets field value
func (o *TransactionPostRequestBody) SetTransactionType(v string) {
	o.TransactionType = v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *TransactionPostRequestBody) SetOperation(v string) {
	o.Operation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransactionPostRequestBody) SetDescription(v string) {
	o.Description = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TransactionPostRequestBody) SetComment(v string) {
	o.Comment = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetCreationDateTime() string {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret string
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetCreationDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given string and assigns it to the CreationDateTime field.
func (o *TransactionPostRequestBody) SetCreationDateTime(v string) {
	o.CreationDateTime = &v
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionPostRequestBody) GetTransactionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetTransactionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionPostRequestBody) SetTransactionId(v int64) {
	o.TransactionId = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *TransactionPostRequestBody) SetUser(v User) {
	o.User = &v
}

// GetRevertingTransaction returns the RevertingTransaction field value if set, zero value otherwise.
func (o *TransactionPostRequestBody) GetRevertingTransaction() Transaction {
	if o == nil || IsNil(o.RevertingTransaction) {
		var ret Transaction
		return ret
	}
	return *o.RevertingTransaction
}

// GetRevertingTransactionOk returns a tuple with the RevertingTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionPostRequestBody) GetRevertingTransactionOk() (*Transaction, bool) {
	if o == nil || IsNil(o.RevertingTransaction) {
		return nil, false
	}
	return o.RevertingTransaction, true
}

// HasRevertingTransaction returns a boolean if a field has been set.
func (o *TransactionPostRequestBody) HasRevertingTransaction() bool {
	if o != nil && !IsNil(o.RevertingTransaction) {
		return true
	}

	return false
}

// SetRevertingTransaction gets a reference to the given Transaction and assigns it to the RevertingTransaction field.
func (o *TransactionPostRequestBody) SetRevertingTransaction(v Transaction) {
	o.RevertingTransaction = &v
}

func (o TransactionPostRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionPostRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["transactionType"] = o.TransactionType
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreationDateTime) {
		toSerialize["creationDateTime"] = o.CreationDateTime
	}
	toSerialize["transactionId"] = o.TransactionId
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.RevertingTransaction) {
		toSerialize["revertingTransaction"] = o.RevertingTransaction
	}
	return toSerialize, nil
}

func (o *TransactionPostRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactionType",
		"transactionId",
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

	varTransactionPostRequestBody := _TransactionPostRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactionPostRequestBody)

	if err != nil {
		return err
	}

	*o = TransactionPostRequestBody(varTransactionPostRequestBody)

	return err
}

type NullableTransactionPostRequestBody struct {
	value *TransactionPostRequestBody
	isSet bool
}

func (v NullableTransactionPostRequestBody) Get() *TransactionPostRequestBody {
	return v.value
}

func (v *NullableTransactionPostRequestBody) Set(val *TransactionPostRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionPostRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionPostRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionPostRequestBody(val *TransactionPostRequestBody) *NullableTransactionPostRequestBody {
	return &NullableTransactionPostRequestBody{value: val, isSet: true}
}

func (v NullableTransactionPostRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionPostRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


