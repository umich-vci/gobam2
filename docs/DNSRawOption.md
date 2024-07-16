# DNSRawOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Value** | Pointer to **string** | The deployment option value. | [optional] 

## Methods

### NewDNSRawOption

`func NewDNSRawOption() *DNSRawOption`

NewDNSRawOption instantiates a new DNSRawOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRawOptionWithDefaults

`func NewDNSRawOptionWithDefaults() *DNSRawOption`

NewDNSRawOptionWithDefaults instantiates a new DNSRawOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSRawOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSRawOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSRawOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSRawOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DNSRawOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DNSRawOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DNSRawOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DNSRawOption) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


