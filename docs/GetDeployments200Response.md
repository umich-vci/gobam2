# GetDeployments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionHalLinks**](CollectionHalLinks.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]GetDeployments200ResponseDataInner**](GetDeployments200ResponseDataInner.md) |  | [optional] [readonly] 

## Methods

### NewGetDeployments200Response

`func NewGetDeployments200Response() *GetDeployments200Response`

NewGetDeployments200Response instantiates a new GetDeployments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeployments200ResponseWithDefaults

`func NewGetDeployments200ResponseWithDefaults() *GetDeployments200Response`

NewGetDeployments200ResponseWithDefaults instantiates a new GetDeployments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetDeployments200Response) GetLinks() CollectionHalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetDeployments200Response) GetLinksOk() (*CollectionHalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetDeployments200Response) SetLinks(v CollectionHalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetDeployments200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCount

`func (o *GetDeployments200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetDeployments200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetDeployments200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetDeployments200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetDeployments200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetDeployments200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetDeployments200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetDeployments200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetDeployments200Response) GetData() []GetDeployments200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDeployments200Response) GetDataOk() (*[]GetDeployments200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDeployments200Response) SetData(v []GetDeployments200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetDeployments200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


