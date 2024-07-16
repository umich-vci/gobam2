# UID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewUID

`func NewUID() *UID`

NewUID instantiates a new UID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIDWithDefaults

`func NewUIDWithDefaults() *UID`

NewUIDWithDefaults instantiates a new UID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIdentifier

`func (o *UID) GetClientIdentifier() string`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *UID) GetClientIdentifierOk() (*string, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *UID) SetClientIdentifier(v string)`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *UID) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


