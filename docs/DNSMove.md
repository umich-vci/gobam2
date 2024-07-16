# DNSMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**AbsoluteName** | Pointer to **string** | The absolute name of the destination zone resource where the DNS resource is to be moved. | [optional] 

## Methods

### NewDNSMove

`func NewDNSMove() *DNSMove`

NewDNSMove instantiates a new DNSMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSMoveWithDefaults

`func NewDNSMoveWithDefaults() *DNSMove`

NewDNSMoveWithDefaults instantiates a new DNSMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSMove) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSMove) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSMove) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSMove) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAbsoluteName

`func (o *DNSMove) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *DNSMove) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *DNSMove) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *DNSMove) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


