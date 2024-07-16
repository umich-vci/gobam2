# DHCPv4CustomClientOptionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of deployment option definition. | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ValueFormat** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDHCPv4CustomClientOptionDefinition

`func NewDHCPv4CustomClientOptionDefinition() *DHCPv4CustomClientOptionDefinition`

NewDHCPv4CustomClientOptionDefinition instantiates a new DHCPv4CustomClientOptionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPv4CustomClientOptionDefinitionWithDefaults

`func NewDHCPv4CustomClientOptionDefinitionWithDefaults() *DHCPv4CustomClientOptionDefinition`

NewDHCPv4CustomClientOptionDefinitionWithDefaults instantiates a new DHCPv4CustomClientOptionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPv4CustomClientOptionDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPv4CustomClientOptionDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPv4CustomClientOptionDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPv4CustomClientOptionDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *DHCPv4CustomClientOptionDefinition) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DHCPv4CustomClientOptionDefinition) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DHCPv4CustomClientOptionDefinition) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DHCPv4CustomClientOptionDefinition) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *DHCPv4CustomClientOptionDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DHCPv4CustomClientOptionDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DHCPv4CustomClientOptionDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DHCPv4CustomClientOptionDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *DHCPv4CustomClientOptionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DHCPv4CustomClientOptionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DHCPv4CustomClientOptionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DHCPv4CustomClientOptionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueFormat

`func (o *DHCPv4CustomClientOptionDefinition) GetValueFormat() map[string]interface{}`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *DHCPv4CustomClientOptionDefinition) GetValueFormatOk() (*map[string]interface{}, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *DHCPv4CustomClientOptionDefinition) SetValueFormat(v map[string]interface{})`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *DHCPv4CustomClientOptionDefinition) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


