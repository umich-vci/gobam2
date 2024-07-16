# DHCPVendorOptionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of deployment option definition. | [optional] 
**Code** | Pointer to **int32** | The ID of the vendor profile option definition. | [optional] 
**DisplayName** | Pointer to **string** | The display name for the vendor profile option definition. | [optional] 
**Description** | Pointer to **string** | The description of the information passed by the option definition. | [optional] 
**ValueFormat** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDHCPVendorOptionDefinition

`func NewDHCPVendorOptionDefinition() *DHCPVendorOptionDefinition`

NewDHCPVendorOptionDefinition instantiates a new DHCPVendorOptionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPVendorOptionDefinitionWithDefaults

`func NewDHCPVendorOptionDefinitionWithDefaults() *DHCPVendorOptionDefinition`

NewDHCPVendorOptionDefinitionWithDefaults instantiates a new DHCPVendorOptionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DHCPVendorOptionDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPVendorOptionDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPVendorOptionDefinition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPVendorOptionDefinition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *DHCPVendorOptionDefinition) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DHCPVendorOptionDefinition) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DHCPVendorOptionDefinition) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *DHCPVendorOptionDefinition) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDisplayName

`func (o *DHCPVendorOptionDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DHCPVendorOptionDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DHCPVendorOptionDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DHCPVendorOptionDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *DHCPVendorOptionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DHCPVendorOptionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DHCPVendorOptionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DHCPVendorOptionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValueFormat

`func (o *DHCPVendorOptionDefinition) GetValueFormat() map[string]interface{}`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *DHCPVendorOptionDefinition) GetValueFormatOk() (*map[string]interface{}, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *DHCPVendorOptionDefinition) SetValueFormat(v map[string]interface{})`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *DHCPVendorOptionDefinition) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


