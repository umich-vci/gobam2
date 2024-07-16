# TransactionPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**TransactionType** | **string** | The type of transaction performed. | 
**Operation** | Pointer to **string** | The operation performed. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the operation performed. | [optional] 
**Comment** | Pointer to **string** | The change control comment of the transaction. | [optional] 
**CreationDateTime** | Pointer to **string** | The date and time of the transaction. | [optional] 
**TransactionId** | **int64** | The unique ID of the transaction performed. | 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**RevertingTransaction** | Pointer to [**Transaction**](Transaction.md) |  | [optional] [readonly] 

## Methods

### NewTransactionPostRequestBody

`func NewTransactionPostRequestBody(transactionType string, transactionId int64, ) *TransactionPostRequestBody`

NewTransactionPostRequestBody instantiates a new TransactionPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPostRequestBodyWithDefaults

`func NewTransactionPostRequestBodyWithDefaults() *TransactionPostRequestBody`

NewTransactionPostRequestBodyWithDefaults instantiates a new TransactionPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *TransactionPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTransactionType

`func (o *TransactionPostRequestBody) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *TransactionPostRequestBody) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *TransactionPostRequestBody) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.


### GetOperation

`func (o *TransactionPostRequestBody) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *TransactionPostRequestBody) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *TransactionPostRequestBody) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *TransactionPostRequestBody) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionPostRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionPostRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionPostRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionPostRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComment

`func (o *TransactionPostRequestBody) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TransactionPostRequestBody) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TransactionPostRequestBody) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TransactionPostRequestBody) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *TransactionPostRequestBody) GetCreationDateTime() string`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *TransactionPostRequestBody) GetCreationDateTimeOk() (*string, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *TransactionPostRequestBody) SetCreationDateTime(v string)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *TransactionPostRequestBody) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetTransactionId

`func (o *TransactionPostRequestBody) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionPostRequestBody) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionPostRequestBody) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.


### GetUser

`func (o *TransactionPostRequestBody) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransactionPostRequestBody) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransactionPostRequestBody) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *TransactionPostRequestBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRevertingTransaction

`func (o *TransactionPostRequestBody) GetRevertingTransaction() Transaction`

GetRevertingTransaction returns the RevertingTransaction field if non-nil, zero value otherwise.

### GetRevertingTransactionOk

`func (o *TransactionPostRequestBody) GetRevertingTransactionOk() (*Transaction, bool)`

GetRevertingTransactionOk returns a tuple with the RevertingTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertingTransaction

`func (o *TransactionPostRequestBody) SetRevertingTransaction(v Transaction)`

SetRevertingTransaction sets RevertingTransaction field to given value.

### HasRevertingTransaction

`func (o *TransactionPostRequestBody) HasRevertingTransaction() bool`

HasRevertingTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


