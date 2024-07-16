# DHCPv4ServiceOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **float32** | The code for the DHCP option (per RFC 2132). | [optional] 
**Definition** | Pointer to [**InlinedDHCPv4ServiceOptionDefinition**](InlinedDHCPv4ServiceOptionDefinition.md) |  | [optional] 

## Methods

### NewDHCPv4ServiceOption

`func NewDHCPv4ServiceOption() *DHCPv4ServiceOption`

NewDHCPv4ServiceOption instantiates a new DHCPv4ServiceOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPv4ServiceOptionWithDefaults

`func NewDHCPv4ServiceOptionWithDefaults() *DHCPv4ServiceOption`

NewDHCPv4ServiceOptionWithDefaults instantiates a new DHCPv4ServiceOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPv4ServiceOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPv4ServiceOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPv4ServiceOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPv4ServiceOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DHCPv4ServiceOption) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DHCPv4ServiceOption) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DHCPv4ServiceOption) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DHCPv4ServiceOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *DHCPv4ServiceOption) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DHCPv4ServiceOption) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DHCPv4ServiceOption) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DHCPv4ServiceOption) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *DHCPv4ServiceOption) GetDefinition() InlinedDHCPv4ServiceOptionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DHCPv4ServiceOption) GetDefinitionOk() (*InlinedDHCPv4ServiceOptionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DHCPv4ServiceOption) SetDefinition(v InlinedDHCPv4ServiceOptionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DHCPv4ServiceOption) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

