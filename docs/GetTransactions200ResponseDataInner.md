# GetTransactions200ResponseDataInner

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
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetTransactions200ResponseDataInner

`func NewGetTransactions200ResponseDataInner() *GetTransactions200ResponseDataInner`

NewGetTransactions200ResponseDataInner instantiates a new GetTransactions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactions200ResponseDataInnerWithDefaults

`func NewGetTransactions200ResponseDataInnerWithDefaults() *GetTransactions200ResponseDataInner`

NewGetTransactions200ResponseDataInnerWithDefaults instantiates a new GetTransactions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTransactions200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTransactions200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTransactions200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetTransactions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetTransactions200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTransactions200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTransactions200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTransactions200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTransactionType

`func (o *GetTransactions200ResponseDataInner) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *GetTransactions200ResponseDataInner) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *GetTransactions200ResponseDataInner) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *GetTransactions200ResponseDataInner) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetOperation

`func (o *GetTransactions200ResponseDataInner) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *GetTransactions200ResponseDataInner) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *GetTransactions200ResponseDataInner) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *GetTransactions200ResponseDataInner) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetDescription

`func (o *GetTransactions200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTransactions200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTransactions200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTransactions200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetComment

`func (o *GetTransactions200ResponseDataInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetTransactions200ResponseDataInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetTransactions200ResponseDataInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetTransactions200ResponseDataInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationDateTime

`func (o *GetTransactions200ResponseDataInner) GetCreationDateTime() string`

GetCreationDateTime returns the CreationDateTime field if non-nil, zero value otherwise.

### GetCreationDateTimeOk

`func (o *GetTransactions200ResponseDataInner) GetCreationDateTimeOk() (*string, bool)`

GetCreationDateTimeOk returns a tuple with the CreationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDateTime

`func (o *GetTransactions200ResponseDataInner) SetCreationDateTime(v string)`

SetCreationDateTime sets CreationDateTime field to given value.

### HasCreationDateTime

`func (o *GetTransactions200ResponseDataInner) HasCreationDateTime() bool`

HasCreationDateTime returns a boolean if a field has been set.

### GetTransactionId

`func (o *GetTransactions200ResponseDataInner) GetTransactionId() int64`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetTransactions200ResponseDataInner) GetTransactionIdOk() (*int64, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetTransactions200ResponseDataInner) SetTransactionId(v int64)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetTransactions200ResponseDataInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUser

`func (o *GetTransactions200ResponseDataInner) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetTransactions200ResponseDataInner) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetTransactions200ResponseDataInner) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *GetTransactions200ResponseDataInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRevertingTransaction

`func (o *GetTransactions200ResponseDataInner) GetRevertingTransaction() Transaction`

GetRevertingTransaction returns the RevertingTransaction field if non-nil, zero value otherwise.

### GetRevertingTransactionOk

`func (o *GetTransactions200ResponseDataInner) GetRevertingTransactionOk() (*Transaction, bool)`

GetRevertingTransactionOk returns a tuple with the RevertingTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertingTransaction

`func (o *GetTransactions200ResponseDataInner) SetRevertingTransaction(v Transaction)`

SetRevertingTransaction sets RevertingTransaction field to given value.

### HasRevertingTransaction

`func (o *GetTransactions200ResponseDataInner) HasRevertingTransaction() bool`

HasRevertingTransaction returns a boolean if a field has been set.

### GetLinks

`func (o *GetTransactions200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetTransactions200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetTransactions200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetTransactions200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


