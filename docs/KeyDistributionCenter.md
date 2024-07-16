# KeyDistributionCenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the Kerberos key distribution center. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Host** | Pointer to **string** | The IP address or hostname of the Kerberos key distribution center | [optional] 
**Port** | Pointer to **int32** | The port of the Kerberos key distribution center | [optional] 

## Methods

### NewKeyDistributionCenter

`func NewKeyDistributionCenter() *KeyDistributionCenter`

NewKeyDistributionCenter instantiates a new KeyDistributionCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyDistributionCenterWithDefaults

`func NewKeyDistributionCenterWithDefaults() *KeyDistributionCenter`

NewKeyDistributionCenterWithDefaults instantiates a new KeyDistributionCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyDistributionCenter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyDistributionCenter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyDistributionCenter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyDistributionCenter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *KeyDistributionCenter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyDistributionCenter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyDistributionCenter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyDistributionCenter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *KeyDistributionCenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyDistributionCenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyDistributionCenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyDistributionCenter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *KeyDistributionCenter) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *KeyDistributionCenter) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *KeyDistributionCenter) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *KeyDistributionCenter) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetHost

`func (o *KeyDistributionCenter) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KeyDistributionCenter) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KeyDistributionCenter) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KeyDistributionCenter) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *KeyDistributionCenter) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KeyDistributionCenter) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KeyDistributionCenter) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KeyDistributionCenter) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

