# GetResponsePolicyImports200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**CollectionHalLinks**](CollectionHalLinks.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [readonly] 
**TotalCount** | Pointer to **int32** |  | [optional] [readonly] 
**Data** | Pointer to [**[]GetResponsePolicyImports200ResponseDataInner**](GetResponsePolicyImports200ResponseDataInner.md) |  | [optional] [readonly] 

## Methods

### NewGetResponsePolicyImports200Response

`func NewGetResponsePolicyImports200Response() *GetResponsePolicyImports200Response`

NewGetResponsePolicyImports200Response instantiates a new GetResponsePolicyImports200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResponsePolicyImports200ResponseWithDefaults

`func NewGetResponsePolicyImports200ResponseWithDefaults() *GetResponsePolicyImports200Response`

NewGetResponsePolicyImports200ResponseWithDefaults instantiates a new GetResponsePolicyImports200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GetResponsePolicyImports200Response) GetLinks() CollectionHalLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetResponsePolicyImports200Response) GetLinksOk() (*CollectionHalLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetResponsePolicyImports200Response) SetLinks(v CollectionHalLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetResponsePolicyImports200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCount

`func (o *GetResponsePolicyImports200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetResponsePolicyImports200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetResponsePolicyImports200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetResponsePolicyImports200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetResponsePolicyImports200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetResponsePolicyImports200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetResponsePolicyImports200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetResponsePolicyImports200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetData

`func (o *GetResponsePolicyImports200Response) GetData() []GetResponsePolicyImports200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetResponsePolicyImports200Response) GetDataOk() (*[]GetResponsePolicyImports200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetResponsePolicyImports200Response) SetData(v []GetResponsePolicyImports200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetResponsePolicyImports200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


