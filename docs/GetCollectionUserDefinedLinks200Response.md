# GetCollectionUserDefinedLinks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionHalLinks**](CollectionHalLinks.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]GetCollectionUserDefinedLinks200ResponseDataInner**](GetCollectionUserDefinedLinks200ResponseDataInner.md) |  | [optional] [readonly] 

## Methods

### NewGetCollectionUserDefinedLinks200Response

`func NewGetCollectionUserDefinedLinks200Response() *GetCollectionUserDefinedLinks200Response`

NewGetCollectionUserDefinedLinks200Response instantiates a new GetCollectionUserDefinedLinks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionUserDefinedLinks200ResponseWithDefaults

`func NewGetCollectionUserDefinedLinks200ResponseWithDefaults() *GetCollectionUserDefinedLinks200Response`

NewGetCollectionUserDefinedLinks200ResponseWithDefaults instantiates a new GetCollectionUserDefinedLinks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetCollectionUserDefinedLinks200Response) GetLinks() CollectionHalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetCollectionUserDefinedLinks200Response) GetLinksOk() (*CollectionHalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetCollectionUserDefinedLinks200Response) SetLinks(v CollectionHalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetCollectionUserDefinedLinks200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCount

`func (o *GetCollectionUserDefinedLinks200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetCollectionUserDefinedLinks200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetCollectionUserDefinedLinks200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetCollectionUserDefinedLinks200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetCollectionUserDefinedLinks200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetCollectionUserDefinedLinks200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetCollectionUserDefinedLinks200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetCollectionUserDefinedLinks200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetCollectionUserDefinedLinks200Response) GetData() []GetCollectionUserDefinedLinks200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCollectionUserDefinedLinks200Response) GetDataOk() (*[]GetCollectionUserDefinedLinks200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCollectionUserDefinedLinks200Response) SetData(v []GetCollectionUserDefinedLinks200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetCollectionUserDefinedLinks200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


