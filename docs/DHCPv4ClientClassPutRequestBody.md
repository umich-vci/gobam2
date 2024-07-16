# DHCPv4ClientClassPutRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the DHCP client class. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Description** | Pointer to **string** | The description for the DHCP client class. | [optional] 
**Option** | Pointer to **string** | The DHCP option to match. | [optional] 
**MatchOffset** | Pointer to **int32** | The offset that determines where to start the check for matching values. The match check is applied offset bytes from the beginning of the identifier. Offset is not configurable for custom match statements. | [optional] 
**MatchLength** | Pointer to **int32** | The length determines the portion of the identifier where the check for matching values will be applied. The match check is applied offset bytes from the beginning of the identifier, continuing for length bytes. Length is not configurable for custom match statements. | [optional] 
**MatchExpression** | Pointer to **string** | The data expression for custom match statements, or boolean expression for custom match if statements. The expression must match the supported syntax for the ISC&#39;s DHCP daemon. | [optional] 

## Methods

### NewDHCPv4ClientClassPutRequestBody

`func NewDHCPv4ClientClassPutRequestBody(name string, ) *DHCPv4ClientClassPutRequestBody`

NewDHCPv4ClientClassPutRequestBody instantiates a new DHCPv4ClientClassPutRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPv4ClientClassPutRequestBodyWithDefaults

`func NewDHCPv4ClientClassPutRequestBodyWithDefaults() *DHCPv4ClientClassPutRequestBody`

NewDHCPv4ClientClassPutRequestBodyWithDefaults instantiates a new DHCPv4ClientClassPutRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DHCPv4ClientClassPutRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DHCPv4ClientClassPutRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DHCPv4ClientClassPutRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DHCPv4ClientClassPutRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *DHCPv4ClientClassPutRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DHCPv4ClientClassPutRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DHCPv4ClientClassPutRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DHCPv4ClientClassPutRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *DHCPv4ClientClassPutRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DHCPv4ClientClassPutRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DHCPv4ClientClassPutRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *DHCPv4ClientClassPutRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *DHCPv4ClientClassPutRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *DHCPv4ClientClassPutRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *DHCPv4ClientClassPutRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *DHCPv4ClientClassPutRequestBody) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DHCPv4ClientClassPutRequestBody) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DHCPv4ClientClassPutRequestBody) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DHCPv4ClientClassPutRequestBody) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *DHCPv4ClientClassPutRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DHCPv4ClientClassPutRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DHCPv4ClientClassPutRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DHCPv4ClientClassPutRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOption

`func (o *DHCPv4ClientClassPutRequestBody) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *DHCPv4ClientClassPutRequestBody) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *DHCPv4ClientClassPutRequestBody) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *DHCPv4ClientClassPutRequestBody) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetMatchOffset

`func (o *DHCPv4ClientClassPutRequestBody) GetMatchOffset() int32`

GetMatchOffset returns the MatchOffset field if non-nil, zero value otherwise.

### GetMatchOffsetOk

`func (o *DHCPv4ClientClassPutRequestBody) GetMatchOffsetOk() (*int32, bool)`

GetMatchOffsetOk returns a tuple with the MatchOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOffset

`func (o *DHCPv4ClientClassPutRequestBody) SetMatchOffset(v int32)`

SetMatchOffset sets MatchOffset field to given value.

### HasMatchOffset

`func (o *DHCPv4ClientClassPutRequestBody) HasMatchOffset() bool`

HasMatchOffset returns a boolean if a field has been set.

### GetMatchLength

`func (o *DHCPv4ClientClassPutRequestBody) GetMatchLength() int32`

GetMatchLength returns the MatchLength field if non-nil, zero value otherwise.

### GetMatchLengthOk

`func (o *DHCPv4ClientClassPutRequestBody) GetMatchLengthOk() (*int32, bool)`

GetMatchLengthOk returns a tuple with the MatchLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLength

`func (o *DHCPv4ClientClassPutRequestBody) SetMatchLength(v int32)`

SetMatchLength sets MatchLength field to given value.

### HasMatchLength

`func (o *DHCPv4ClientClassPutRequestBody) HasMatchLength() bool`

HasMatchLength returns a boolean if a field has been set.

### GetMatchExpression

`func (o *DHCPv4ClientClassPutRequestBody) GetMatchExpression() string`

GetMatchExpression returns the MatchExpression field if non-nil, zero value otherwise.

### GetMatchExpressionOk

`func (o *DHCPv4ClientClassPutRequestBody) GetMatchExpressionOk() (*string, bool)`

GetMatchExpressionOk returns a tuple with the MatchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpression

`func (o *DHCPv4ClientClassPutRequestBody) SetMatchExpression(v string)`

SetMatchExpression sets MatchExpression field to given value.

### HasMatchExpression

`func (o *DHCPv4ClientClassPutRequestBody) HasMatchExpression() bool`

HasMatchExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


