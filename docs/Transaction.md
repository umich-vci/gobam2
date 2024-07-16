# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**TransactionType** | Pointer to **string** | The type of transaction performed. | [optional] 
**Operation** | Pointer to **string** | The operation performed. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the operation performed. | [optional] 
**Comment** | Pointer to **string** | The change control comment of the transaction. | [optional] 
**CreationDateTime** | Pointer to **string** | The date and time of the transaction. | [optional] 
**TransactionId** | Pointer to **int64** | The unique ID of the transaction performed. | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**RevertingTransaction** | Pointer to [**Transaction**](Transaction.md) |  | [optional] [readonly] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTransactionType

`func (o *Transaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Transaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Transaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *Transaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetOperation

`func (o *Transaction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Transaction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Transaction) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Transaction) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComment

`func (o *Transaction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Transaction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Transaction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Transaction) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *Transaction) GetCreationDateTime() string`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *Transaction) GetCreationDateTimeOk() (*string, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *Transaction) SetCreationDateTime(v string)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *Transaction) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetTransactionId

`func (o *Transaction) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Transaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUser

`func (o *Transaction) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Transaction) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Transaction) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Transaction) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRevertingTransaction

`func (o *Transaction) GetRevertingTransaction() Transaction`

GetRevertingTransaction returns the RevertingTransaction field if non-nil, zero value otherwise.

### GetRevertingTransactionOk

`func (o *Transaction) GetRevertingTransactionOk() (*Transaction, bool)`

GetRevertingTransactionOk returns a tuple with the RevertingTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertingTransaction

`func (o *Transaction) SetRevertingTransaction(v Transaction)`

SetRevertingTransaction sets RevertingTransaction field to given value.

### HasRevertingTransaction

`func (o *Transaction) HasRevertingTransaction() bool`

HasRevertingTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


