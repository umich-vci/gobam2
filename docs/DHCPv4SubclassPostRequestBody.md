# DHCPv4SubclassPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the resource. | [optional] [readonly] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Description** | Pointer to **string** | The description for the DHCP subclass value. | [optional] 
**Value** | **string** | The DHCP subclass value. | 

## Methods

### NewDHCPv4SubclassPostRequestBody

`func NewDHCPv4SubclassPostRequestBody(value string, ) *DHCPv4SubclassPostRequestBody`

NewDHCPv4SubclassPostRequestBody instantiates a new DHCPv4SubclassPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPv4SubclassPostRequestBodyWithDefaults

`func NewDHCPv4SubclassPostRequestBodyWithDefaults() *DHCPv4SubclassPostRequestBody`

NewDHCPv4SubclassPostRequestBodyWithDefaults instantiates a new DHCPv4SubclassPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DHCPv4SubclassPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DHCPv4SubclassPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DHCPv4SubclassPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DHCPv4SubclassPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DHCPv4SubclassPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPv4SubclassPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPv4SubclassPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPv4SubclassPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DHCPv4SubclassPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DHCPv4SubclassPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DHCPv4SubclassPostRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DHCPv4SubclassPostRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *DHCPv4SubclassPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DHCPv4SubclassPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DHCPv4SubclassPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DHCPv4SubclassPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *DHCPv4SubclassPostRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DHCPv4SubclassPostRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DHCPv4SubclassPostRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DHCPv4SubclassPostRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *DHCPv4SubclassPostRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DHCPv4SubclassPostRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DHCPv4SubclassPostRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DHCPv4SubclassPostRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *DHCPv4SubclassPostRequestBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DHCPv4SubclassPostRequestBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DHCPv4SubclassPostRequestBody) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


