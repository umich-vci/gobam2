# KeyDistributionCenterPostRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | **string** | The name of the Kerberos key distribution center. | 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Host** | **string** | The IP address or hostname of the Kerberos key distribution center | 
**Port** | Pointer to **int32** | The port of the Kerberos key distribution center | [optional] 

## Methods

### NewKeyDistributionCenterPostRequestBody

`func NewKeyDistributionCenterPostRequestBody(name string, host string, ) *KeyDistributionCenterPostRequestBody`

NewKeyDistributionCenterPostRequestBody instantiates a new KeyDistributionCenterPostRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyDistributionCenterPostRequestBodyWithDefaults

`func NewKeyDistributionCenterPostRequestBodyWithDefaults() *KeyDistributionCenterPostRequestBody`

NewKeyDistributionCenterPostRequestBodyWithDefaults instantiates a new KeyDistributionCenterPostRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyDistributionCenterPostRequestBody) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyDistributionCenterPostRequestBody) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyDistributionCenterPostRequestBody) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyDistributionCenterPostRequestBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KeyDistributionCenterPostRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyDistributionCenterPostRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyDistributionCenterPostRequestBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyDistributionCenterPostRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *KeyDistributionCenterPostRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyDistributionCenterPostRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyDistributionCenterPostRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetUserDefinedFields

`func (o *KeyDistributionCenterPostRequestBody) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *KeyDistributionCenterPostRequestBody) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *KeyDistributionCenterPostRequestBody) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *KeyDistributionCenterPostRequestBody) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetHost

`func (o *KeyDistributionCenterPostRequestBody) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KeyDistributionCenterPostRequestBody) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KeyDistributionCenterPostRequestBody) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *KeyDistributionCenterPostRequestBody) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KeyDistributionCenterPostRequestBody) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KeyDistributionCenterPostRequestBody) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KeyDistributionCenterPostRequestBody) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


