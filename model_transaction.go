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

// checks if the Transaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transaction{}

// Transaction struct for Transaction
type Transaction struct {
	// The resource identifier.
	Id *int64 `json:"id,omitempty"`
	// The resource type.
	Type *string `json:"type,omitempty"`
	// The type of transaction performed.
	TransactionType *string `json:"transactionType,omitempty"`
	// The operation performed.
	Operation *string `json:"operation,omitempty"`
	// The description of the operation performed.
	Description *string `json:"description,omitempty"`
	// The change control comment of the transaction.
	Comment *string `json:"comment,omitempty"`
	// The date and time of the transaction.
	CreationDateTime *string `json:"creationDateTime,omitempty"`
	// The unique ID of the transaction performed.
	TransactionId *int64 `json:"transactionId,omitempty"`
	User *User `json:"user,omitempty"`
	RevertingTransaction *Transaction `json:"revertingTransaction,omitempty"`
}

// NewTransaction instantiates a new Transaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransaction() *Transaction {
	this := Transaction{}
	return &this
}

// NewTransactionWithDefaults instantiates a new Transaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionWithDefaults() *Transaction {
	this := Transaction{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Transaction) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Transaction) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Transaction) SetId(v int64) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Transaction) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Transaction) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Transaction) SetType(v string) {
	o.Type = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *Transaction) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *Transaction) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *Transaction) SetTransactionType(v string) {
	o.TransactionType = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *Transaction) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *Transaction) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *Transaction) SetOperation(v string) {
	o.Operation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Transaction) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Transaction) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Transaction) SetDescription(v string) {
	o.Description = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Transaction) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Transaction) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Transaction) SetComment(v string) {
	o.Comment = &v
}

// GetCreationDateTime returns the CreationDateTime field value if set, zero value otherwise.
func (o *Transaction) GetCreationDateTime() string {
	if o == nil || IsNil(o.CreationDateTime) {
		var ret string
		return ret
	}
	return *o.CreationDateTime
}

// GetCreationDateTimeOk returns a tuple with the CreationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetCreationDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationDateTime) {
		return nil, false
	}
	return o.CreationDateTime, true
}

// HasCreationDateTime returns a boolean if a field has been set.
func (o *Transaction) HasCreationDateTime() bool {
	if o != nil && !IsNil(o.CreationDateTime) {
		return true
	}

	return false
}

// SetCreationDateTime gets a reference to the given string and assigns it to the CreationDateTime field.
func (o *Transaction) SetCreationDateTime(v string) {
	o.CreationDateTime = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *Transaction) GetTransactionId() int64 {
	if o == nil || IsNil(o.TransactionId) {
		var ret int64
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetTransactionIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *Transaction) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given int64 and assigns it to the TransactionId field.
func (o *Transaction) SetTransactionId(v int64) {
	o.TransactionId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Transaction) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Transaction) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Transaction) SetUser(v User) {
	o.User = &v
}

// GetRevertingTransaction returns the RevertingTransaction field value if set, zero value otherwise.
func (o *Transaction) GetRevertingTransaction() Transaction {
	if o == nil || IsNil(o.RevertingTransaction) {
		var ret Transaction
		return ret
	}
	return *o.RevertingTransaction
}

// GetRevertingTransactionOk returns a tuple with the RevertingTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transaction) GetRevertingTransactionOk() (*Transaction, bool) {
	if o == nil || IsNil(o.RevertingTransaction) {
		return nil, false
	}
	return o.RevertingTransaction, true
}

// HasRevertingTransaction returns a boolean if a field has been set.
func (o *Transaction) HasRevertingTransaction() bool {
	if o != nil && !IsNil(o.RevertingTransaction) {
		return true
	}

	return false
}

// SetRevertingTransaction gets a reference to the given Transaction and assigns it to the RevertingTransaction field.
func (o *Transaction) SetRevertingTransaction(v Transaction) {
	o.RevertingTransaction = &v
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transactionType"] = o.TransactionType
	}
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
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.RevertingTransaction) {
		toSerialize["revertingTransaction"] = o.RevertingTransaction
	}
	return toSerialize, nil
}

type NullableTransaction struct {
	value *Transaction
	isSet bool
}

func (v NullableTransaction) Get() *Transaction {
	return v.value
}

func (v *NullableTransaction) Set(val *Transaction) {
	v.value = val
	v.isSet = true
}

func (v NullableTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransaction(val *Transaction) *NullableTransaction {
	return &NullableTransaction{value: val, isSet: true}
}

func (v NullableTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


