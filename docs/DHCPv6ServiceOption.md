# DHCPv6ServiceOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **float32** | The code for the DHCP option (per RFC 2132). | [optional] 
**Definition** | Pointer to [**InlinedDHCPv6ServiceOptionDefinition**](InlinedDHCPv6ServiceOptionDefinition.md) |  | [optional] 

## Methods

### NewDHCPv6ServiceOption

`func NewDHCPv6ServiceOption() *DHCPv6ServiceOption`

NewDHCPv6ServiceOption instantiates a new DHCPv6ServiceOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPv6ServiceOptionWithDefaults

`func NewDHCPv6ServiceOptionWithDefaults() *DHCPv6ServiceOption`

NewDHCPv6ServiceOptionWithDefaults instantiates a new DHCPv6ServiceOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPv6ServiceOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPv6ServiceOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPv6ServiceOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPv6ServiceOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DHCPv6ServiceOption) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DHCPv6ServiceOption) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DHCPv6ServiceOption) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DHCPv6ServiceOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *DHCPv6ServiceOption) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DHCPv6ServiceOption) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DHCPv6ServiceOption) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DHCPv6ServiceOption) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *DHCPv6ServiceOption) GetDefinition() InlinedDHCPv6ServiceOptionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DHCPv6ServiceOption) GetDefinitionOk() (*InlinedDHCPv6ServiceOptionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DHCPv6ServiceOption) SetDefinition(v InlinedDHCPv6ServiceOptionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DHCPv6ServiceOption) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


