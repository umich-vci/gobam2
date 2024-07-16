# DHCPv6ClientOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **float32** | The code for the DHCP option (per RFC 2132). | [optional] 
**Definition** | Pointer to [**InlinedDHCPv6ClientOptionDefinition**](InlinedDHCPv6ClientOptionDefinition.md) |  | [optional] 

## Methods

### NewDHCPv6ClientOption

`func NewDHCPv6ClientOption() *DHCPv6ClientOption`

NewDHCPv6ClientOption instantiates a new DHCPv6ClientOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPv6ClientOptionWithDefaults

`func NewDHCPv6ClientOptionWithDefaults() *DHCPv6ClientOption`

NewDHCPv6ClientOptionWithDefaults instantiates a new DHCPv6ClientOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPv6ClientOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPv6ClientOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPv6ClientOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPv6ClientOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DHCPv6ClientOption) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DHCPv6ClientOption) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DHCPv6ClientOption) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DHCPv6ClientOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *DHCPv6ClientOption) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DHCPv6ClientOption) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DHCPv6ClientOption) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DHCPv6ClientOption) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *DHCPv6ClientOption) GetDefinition() InlinedDHCPv6ClientOptionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DHCPv6ClientOption) GetDefinitionOk() (*InlinedDHCPv6ClientOptionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DHCPv6ClientOption) SetDefinition(v InlinedDHCPv6ClientOptionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DHCPv6ClientOption) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


