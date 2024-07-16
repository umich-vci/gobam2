# GetClientClasses200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the DHCP client class. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Configuration** | Pointer to [**InlinedConfiguration**](InlinedConfiguration.md) |  | [optional] [readonly] 
**Description** | Pointer to **string** | The description for the DHCP client class. | [optional] 
**Option** | Pointer to **string** | The DHCP option to match. | [optional] 
**MatchOffset** | Pointer to **int32** | The offset that determines where to start the check for matching values. The match check is applied offset bytes from the beginning of the identifier. Offset is not configurable for custom match statements. | [optional] 
**MatchLength** | Pointer to **int32** | The length determines the portion of the identifier where the check for matching values will be applied. The match check is applied offset bytes from the beginning of the identifier, continuing for length bytes. Length is not configurable for custom match statements. | [optional] 
**MatchExpression** | Pointer to **string** | The data expression for custom match statements, or boolean expression for custom match if statements. The expression must match the supported syntax for the ISC&#39;s DHCP daemon. | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetClientClasses200ResponseDataInner

`func NewGetClientClasses200ResponseDataInner() *GetClientClasses200ResponseDataInner`

NewGetClientClasses200ResponseDataInner instantiates a new GetClientClasses200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClientClasses200ResponseDataInnerWithDefaults

`func NewGetClientClasses200ResponseDataInnerWithDefaults() *GetClientClasses200ResponseDataInner`

NewGetClientClasses200ResponseDataInnerWithDefaults instantiates a new GetClientClasses200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClientClasses200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClientClasses200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClientClasses200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClientClasses200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetClientClasses200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetClientClasses200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetClientClasses200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetClientClasses200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetClientClasses200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClientClasses200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClientClasses200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClientClasses200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetClientClasses200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetClientClasses200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetClientClasses200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetClientClasses200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetConfiguration

`func (o *GetClientClasses200ResponseDataInner) GetConfiguration() InlinedConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetClientClasses200ResponseDataInner) GetConfigurationOk() (*InlinedConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetClientClasses200ResponseDataInner) SetConfiguration(v InlinedConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetClientClasses200ResponseDataInner) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *GetClientClasses200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClientClasses200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClientClasses200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClientClasses200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOption

`func (o *GetClientClasses200ResponseDataInner) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *GetClientClasses200ResponseDataInner) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *GetClientClasses200ResponseDataInner) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *GetClientClasses200ResponseDataInner) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetMatchOffset

`func (o *GetClientClasses200ResponseDataInner) GetMatchOffset() int32`

GetMatchOffset returns the MatchOffset field if non-nil, zero value otherwise.

### GetMatchOffsetOk

`func (o *GetClientClasses200ResponseDataInner) GetMatchOffsetOk() (*int32, bool)`

GetMatchOffsetOk returns a tuple with the MatchOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOffset

`func (o *GetClientClasses200ResponseDataInner) SetMatchOffset(v int32)`

SetMatchOffset sets MatchOffset field to given value.

### HasMatchOffset

`func (o *GetClientClasses200ResponseDataInner) HasMatchOffset() bool`

HasMatchOffset returns a boolean if a field has been set.

### GetMatchLength

`func (o *GetClientClasses200ResponseDataInner) GetMatchLength() int32`

GetMatchLength returns the MatchLength field if non-nil, zero value otherwise.

### GetMatchLengthOk

`func (o *GetClientClasses200ResponseDataInner) GetMatchLengthOk() (*int32, bool)`

GetMatchLengthOk returns a tuple with the MatchLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLength

`func (o *GetClientClasses200ResponseDataInner) SetMatchLength(v int32)`

SetMatchLength sets MatchLength field to given value.

### HasMatchLength

`func (o *GetClientClasses200ResponseDataInner) HasMatchLength() bool`

HasMatchLength returns a boolean if a field has been set.

### GetMatchExpression

`func (o *GetClientClasses200ResponseDataInner) GetMatchExpression() string`

GetMatchExpression returns the MatchExpression field if non-nil, zero value otherwise.

### GetMatchExpressionOk

`func (o *GetClientClasses200ResponseDataInner) GetMatchExpressionOk() (*string, bool)`

GetMatchExpressionOk returns a tuple with the MatchExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpression

`func (o *GetClientClasses200ResponseDataInner) SetMatchExpression(v string)`

SetMatchExpression sets MatchExpression field to given value.

### HasMatchExpression

`func (o *GetClientClasses200ResponseDataInner) HasMatchExpression() bool`

HasMatchExpression returns a boolean if a field has been set.

### GetLinks

`func (o *GetClientClasses200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetClientClasses200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetClientClasses200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetClientClasses200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


