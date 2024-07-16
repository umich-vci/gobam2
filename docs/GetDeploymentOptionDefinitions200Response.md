# GetDeploymentOptionDefinitions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionHalLinks**](CollectionHalLinks.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]GetDeploymentOptionDefinitions200ResponseDataInner**](GetDeploymentOptionDefinitions200ResponseDataInner.md) |  | [optional] [readonly] 

## Methods

### NewGetDeploymentOptionDefinitions200Response

`func NewGetDeploymentOptionDefinitions200Response() *GetDeploymentOptionDefinitions200Response`

NewGetDeploymentOptionDefinitions200Response instantiates a new GetDeploymentOptionDefinitions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentOptionDefinitions200ResponseWithDefaults

`func NewGetDeploymentOptionDefinitions200ResponseWithDefaults() *GetDeploymentOptionDefinitions200Response`

NewGetDeploymentOptionDefinitions200ResponseWithDefaults instantiates a new GetDeploymentOptionDefinitions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetDeploymentOptionDefinitions200Response) GetLinks() CollectionHalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeploymentOptionDefinitions200Response) GetLinksOk() (*CollectionHalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeploymentOptionDefinitions200Response) SetLinks(v CollectionHalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeploymentOptionDefinitions200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCount

`func (o *GetDeploymentOptionDefinitions200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetDeploymentOptionDefinitions200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetDeploymentOptionDefinitions200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetDeploymentOptionDefinitions200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetDeploymentOptionDefinitions200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetDeploymentOptionDefinitions200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetDeploymentOptionDefinitions200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetDeploymentOptionDefinitions200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetDeploymentOptionDefinitions200Response) GetData() []GetDeploymentOptionDefinitions200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDeploymentOptionDefinitions200Response) GetDataOk() (*[]GetDeploymentOptionDefinitions200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDeploymentOptionDefinitions200Response) SetData(v []GetDeploymentOptionDefinitions200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetDeploymentOptionDefinitions200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


