# DHCPVendorOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The resource type. | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**Code** | Pointer to **float32** | The vendor option code. | [optional] [readonly] 
**Definition** | Pointer to [**InlinedDHCPVendorOptionDefinition**](InlinedDHCPVendorOptionDefinition.md) |  | [optional] 

## Methods

### NewDHCPVendorOption

`func NewDHCPVendorOption() *DHCPVendorOption`

NewDHCPVendorOption instantiates a new DHCPVendorOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPVendorOptionWithDefaults

`func NewDHCPVendorOptionWithDefaults() *DHCPVendorOption`

NewDHCPVendorOptionWithDefaults instantiates a new DHCPVendorOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPVendorOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPVendorOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPVendorOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPVendorOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DHCPVendorOption) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DHCPVendorOption) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DHCPVendorOption) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DHCPVendorOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCode

`func (o *DHCPVendorOption) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DHCPVendorOption) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DHCPVendorOption) SetCode(v float32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DHCPVendorOption) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefinition

`func (o *DHCPVendorOption) GetDefinition() InlinedDHCPVendorOptionDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DHCPVendorOption) GetDefinitionOk() (*InlinedDHCPVendorOptionDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DHCPVendorOption) SetDefinition(v InlinedDHCPVendorOptionDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DHCPVendorOption) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


