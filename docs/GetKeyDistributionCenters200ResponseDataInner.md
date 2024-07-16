# GetKeyDistributionCenters200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The resource identifier. | [optional] 
**Type** | Pointer to **string** | The resource type. | [optional] 
**Name** | Pointer to **string** | The name of the Kerberos key distribution center. | [optional] 
**UserDefinedFields** | Pointer to **map[string]string** | User-defined fields set for the resource. | [optional] 
**Host** | Pointer to **string** | The IP address or hostname of the Kerberos key distribution center | [optional] 
**Port** | Pointer to **int32** | The port of the Kerberos key distribution center | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGetKeyDistributionCenters200ResponseDataInner

`func NewGetKeyDistributionCenters200ResponseDataInner() *GetKeyDistributionCenters200ResponseDataInner`

NewGetKeyDistributionCenters200ResponseDataInner instantiates a new GetKeyDistributionCenters200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKeyDistributionCenters200ResponseDataInnerWithDefaults

`func NewGetKeyDistributionCenters200ResponseDataInnerWithDefaults() *GetKeyDistributionCenters200ResponseDataInner`

NewGetKeyDistributionCenters200ResponseDataInnerWithDefaults instantiates a new GetKeyDistributionCenters200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetUserDefinedFields() map[string]string`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetUserDefinedFieldsOk() (*map[string]string, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetUserDefinedFields(v map[string]string)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetHost

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLinks

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetKeyDistributionCenters200ResponseDataInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetKeyDistributionCenters200ResponseDataInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetKeyDistributionCenters200ResponseDataInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


