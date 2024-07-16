# \NamingPolicyResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCollectionNamingPolicy**](NamingPolicyResourcesAPI.md#DeleteCollectionNamingPolicy) | **Delete** /api/v2/{collection}/{collectionId}/namingPolicies/{id} | Unapply naming policy
[**DeleteNamingPolicy**](NamingPolicyResourcesAPI.md#DeleteNamingPolicy) | **Delete** /api/v2/namingPolicies/{id} | Delete a naming policy
[**DeleteNamingPolicyRestriction**](NamingPolicyResourcesAPI.md#DeleteNamingPolicyRestriction) | **Delete** /api/v2/namingPolicyRestrictions/{id} | Delete a naming restriction
[**DeleteNamingPolicyValue**](NamingPolicyResourcesAPI.md#DeleteNamingPolicyValue) | **Delete** /api/v2/namingPolicyValues/{id} | Delete a naming policy value
[**GetCollectionNamingPolicies**](NamingPolicyResourcesAPI.md#GetCollectionNamingPolicies) | **Get** /api/v2/{collection}/{collectionId}/namingPolicies | Retrieve a collection of applied naming policies
[**GetCollectionNamingPolicy**](NamingPolicyResourcesAPI.md#GetCollectionNamingPolicy) | **Get** /api/v2/{collection}/{collectionId}/namingPolicies/{id} | Retrieve an applied naming policy
[**GetNamingPolicies**](NamingPolicyResourcesAPI.md#GetNamingPolicies) | **Get** /api/v2/namingPolicies | Retrieve a collection of naming policies
[**GetNamingPolicy**](NamingPolicyResourcesAPI.md#GetNamingPolicy) | **Get** /api/v2/namingPolicies/{id} | Retrieve a single naming policy
[**GetNamingPolicyPolicyRestriction**](NamingPolicyResourcesAPI.md#GetNamingPolicyPolicyRestriction) | **Get** /api/v2/namingPolicies/{collectionId}/policyRestrictions/{id} | Retrieve a single naming policy restriction
[**GetNamingPolicyPolicyRestrictions**](NamingPolicyResourcesAPI.md#GetNamingPolicyPolicyRestrictions) | **Get** /api/v2/namingPolicies/{collectionId}/policyRestrictions | Retrieve a collection of naming policy restrictions
[**GetNamingPolicyPolicyValue**](NamingPolicyResourcesAPI.md#GetNamingPolicyPolicyValue) | **Get** /api/v2/namingPolicies/{collectionId}/policyValues/{id} | Retrieve a single naming policy value
[**GetNamingPolicyPolicyValues**](NamingPolicyResourcesAPI.md#GetNamingPolicyPolicyValues) | **Get** /api/v2/namingPolicies/{collectionId}/policyValues | Retrieve a collection of naming policy values
[**GetNamingPolicyRestriction**](NamingPolicyResourcesAPI.md#GetNamingPolicyRestriction) | **Get** /api/v2/namingPolicyRestrictions/{id} | Retrieve a single naming restriction
[**GetNamingPolicyRestrictions**](NamingPolicyResourcesAPI.md#GetNamingPolicyRestrictions) | **Get** /api/v2/namingPolicyRestrictions | Retrieve a collection of naming restrictions
[**GetNamingPolicyValue**](NamingPolicyResourcesAPI.md#GetNamingPolicyValue) | **Get** /api/v2/namingPolicyValues/{id} | Retrieve a single naming policy value
[**GetNamingPolicyValues**](NamingPolicyResourcesAPI.md#GetNamingPolicyValues) | **Get** /api/v2/namingPolicyValues | Retrieve a collection of naming policy values
[**PostCollectionNamingPolicy**](NamingPolicyResourcesAPI.md#PostCollectionNamingPolicy) | **Post** /api/v2/{collection}/{collectionId}/namingPolicies | Apply a naming policy
[**PostNamingPolicy**](NamingPolicyResourcesAPI.md#PostNamingPolicy) | **Post** /api/v2/namingPolicies | Create a new naming policy
[**PostNamingPolicyRestriction**](NamingPolicyResourcesAPI.md#PostNamingPolicyRestriction) | **Post** /api/v2/namingPolicyRestrictions | Create a new naming restriction
[**PostNamingPolicyValue**](NamingPolicyResourcesAPI.md#PostNamingPolicyValue) | **Post** /api/v2/namingPolicyValues | Create a new naming policy value
[**PutNamingPolicy**](NamingPolicyResourcesAPI.md#PutNamingPolicy) | **Put** /api/v2/namingPolicies/{id} | Update a naming policy
[**PutNamingPolicyRestriction**](NamingPolicyResourcesAPI.md#PutNamingPolicyRestriction) | **Put** /api/v2/namingPolicyRestrictions/{id} | Update a naming restriction
[**PutNamingPolicyValue**](NamingPolicyResourcesAPI.md#PutNamingPolicyValue) | **Put** /api/v2/namingPolicyValues/{id} | Update a naming policy value



## DeleteCollectionNamingPolicy

> DeleteCollectionNamingPolicy(ctx, collection, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Unapply naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collection := "collection_example" // string | The name of the collection.
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamingPolicyResourcesAPI.DeleteCollectionNamingPolicy(context.Background(), collection, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.DeleteCollectionNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamingPolicy

> DeleteNamingPolicy(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamingPolicyResourcesAPI.DeleteNamingPolicy(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.DeleteNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamingPolicyRestriction

> DeleteNamingPolicyRestriction(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a naming restriction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamingPolicyResourcesAPI.DeleteNamingPolicyRestriction(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.DeleteNamingPolicyRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamingPolicyRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamingPolicyValue

> DeleteNamingPolicyValue(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a naming policy value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NamingPolicyResourcesAPI.DeleteNamingPolicyValue(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.DeleteNamingPolicyValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamingPolicyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionNamingPolicies

> GetNamingPolicies200Response GetCollectionNamingPolicies(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of applied naming policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collection := "collection_example" // string | The name of the collection.
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetCollectionNamingPolicies(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetCollectionNamingPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionNamingPolicies`: GetNamingPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetCollectionNamingPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionNamingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetNamingPolicies200Response**](GetNamingPolicies200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionNamingPolicy

> GetNamingPolicies200ResponseDataInner GetCollectionNamingPolicy(ctx, collection, collectionId, id).Fields(fields).Execute()

Retrieve an applied naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collection := "collection_example" // string | The name of the collection.
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetCollectionNamingPolicy(context.Background(), collection, collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetCollectionNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionNamingPolicy`: GetNamingPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetCollectionNamingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetNamingPolicies200ResponseDataInner**](GetNamingPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicies

> GetNamingPolicies200Response GetNamingPolicies(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of naming policies



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicies(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicies`: GetNamingPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetNamingPolicies200Response**](GetNamingPolicies200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicy

> GetNamingPolicies200ResponseDataInner GetNamingPolicy(ctx, id).Fields(fields).Execute()

Retrieve a single naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicy(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicy`: GetNamingPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetNamingPolicies200ResponseDataInner**](GetNamingPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyPolicyRestriction

> GetLocationAnnotatedResources200ResponseDataInner GetNamingPolicyPolicyRestriction(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single naming policy restriction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestriction(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyPolicyRestriction`: GetLocationAnnotatedResources200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyPolicyRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetLocationAnnotatedResources200ResponseDataInner**](GetLocationAnnotatedResources200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyPolicyRestrictions

> GetLocationAnnotatedResources200Response GetNamingPolicyPolicyRestrictions(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of naming policy restrictions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestrictions(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyPolicyRestrictions`: GetLocationAnnotatedResources200Response
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestrictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyPolicyRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetLocationAnnotatedResources200Response**](GetLocationAnnotatedResources200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyPolicyValue

> GetLocationAnnotatedResources200ResponseDataInner GetNamingPolicyPolicyValue(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single naming policy value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyValue(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyPolicyValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyPolicyValue`: GetLocationAnnotatedResources200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyPolicyValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyPolicyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetLocationAnnotatedResources200ResponseDataInner**](GetLocationAnnotatedResources200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyPolicyValues

> GetLocationAnnotatedResources200Response GetNamingPolicyPolicyValues(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of naming policy values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyValues(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyPolicyValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyPolicyValues`: GetLocationAnnotatedResources200Response
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyPolicyValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyPolicyValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetLocationAnnotatedResources200Response**](GetLocationAnnotatedResources200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyRestriction

> GetNamingPolicyRestrictions200ResponseDataInner GetNamingPolicyRestriction(ctx, id).Fields(fields).Execute()

Retrieve a single naming restriction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyRestriction(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyRestriction`: GetNamingPolicyRestrictions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetNamingPolicyRestrictions200ResponseDataInner**](GetNamingPolicyRestrictions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyRestrictions

> GetNamingPolicyRestrictions200Response GetNamingPolicyRestrictions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of naming restrictions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyRestrictions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyRestrictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyRestrictions`: GetNamingPolicyRestrictions200Response
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyRestrictions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyRestrictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetNamingPolicyRestrictions200Response**](GetNamingPolicyRestrictions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyValue

> GetNamingPolicyValues200ResponseDataInner GetNamingPolicyValue(ctx, id).Fields(fields).Execute()

Retrieve a single naming policy value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	fields := "fields_example" // string | The subset of response resource field paths. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyValue(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyValue`: GetNamingPolicyValues200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetNamingPolicyValues200ResponseDataInner**](GetNamingPolicyValues200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamingPolicyValues

> GetNamingPolicyValues200Response GetNamingPolicyValues(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of naming policy values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	fields := "fields_example" // string | The subset of response resource field paths. (optional)
	limit := int32(56) // int32 | The maximum number of resources returned in the response. (optional)
	offset := int32(56) // int32 | The offset of the first resource returned in the response. (optional)
	filter := "filter_example" // string | The query filter string. (optional)
	orderBy := "orderBy_example" // string | The field path to order the response resources by. (optional)
	total := true // bool | Return the total number of resources matching the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyValues(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.GetNamingPolicyValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamingPolicyValues`: GetNamingPolicyValues200Response
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.GetNamingPolicyValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamingPolicyValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetNamingPolicyValues200Response**](GetNamingPolicyValues200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionNamingPolicy

> GetNamingPolicies200ResponseDataInner PostCollectionNamingPolicy(ctx, collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedNamingPolicyPostRequestBody(linkedNamingPolicyPostRequestBody).Execute()

Apply a naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	collection := "collection_example" // string | The name of the collection.
	collectionId := int64(12345) // int64 | The ID of the collection resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	linkedNamingPolicyPostRequestBody := *openapiclient.NewLinkedNamingPolicyPostRequestBody(int64(103307)) // LinkedNamingPolicyPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PostCollectionNamingPolicy(context.Background(), collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedNamingPolicyPostRequestBody(linkedNamingPolicyPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PostCollectionNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollectionNamingPolicy`: GetNamingPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PostCollectionNamingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **linkedNamingPolicyPostRequestBody** | [**LinkedNamingPolicyPostRequestBody**](LinkedNamingPolicyPostRequestBody.md) |  | 

### Return type

[**GetNamingPolicies200ResponseDataInner**](GetNamingPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNamingPolicy

> GetNamingPolicies200ResponseDataInner PostNamingPolicy(ctx).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyPostRequestBody(namingPolicyPostRequestBody).Execute()

Create a new naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	namingPolicyPostRequestBody := *openapiclient.NewNamingPolicyPostRequestBody("naming policy name") // NamingPolicyPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PostNamingPolicy(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyPostRequestBody(namingPolicyPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PostNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNamingPolicy`: GetNamingPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PostNamingPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **namingPolicyPostRequestBody** | [**NamingPolicyPostRequestBody**](NamingPolicyPostRequestBody.md) |  | 

### Return type

[**GetNamingPolicies200ResponseDataInner**](GetNamingPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNamingPolicyRestriction

> GetNamingPolicyRestrictions200ResponseDataInner PostNamingPolicyRestriction(ctx).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyRestrictionPostRequestBody(namingPolicyRestrictionPostRequestBody).Execute()

Create a new naming restriction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	namingPolicyRestrictionPostRequestBody := *openapiclient.NewNamingPolicyRestrictionPostRequestBody("name", []openapiclient.ValueMatchTypePairBean{*openapiclient.NewValueMatchTypePairBean()}) // NamingPolicyRestrictionPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PostNamingPolicyRestriction(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyRestrictionPostRequestBody(namingPolicyRestrictionPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PostNamingPolicyRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNamingPolicyRestriction`: GetNamingPolicyRestrictions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PostNamingPolicyRestriction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNamingPolicyRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **namingPolicyRestrictionPostRequestBody** | [**NamingPolicyRestrictionPostRequestBody**](NamingPolicyRestrictionPostRequestBody.md) |  | 

### Return type

[**GetNamingPolicyRestrictions200ResponseDataInner**](GetNamingPolicyRestrictions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNamingPolicyValue

> GetNamingPolicyValues200ResponseDataInner PostNamingPolicyValue(ctx).XBcnChangeControlComment(xBcnChangeControlComment).PostNamingPolicyValueRequest(postNamingPolicyValueRequest).Execute()

Create a new naming policy value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	postNamingPolicyValueRequest := openapiclient.postNamingPolicyValue_request{NamingPolicyConnectorValuePostRequestBody: openapiclient.NewNamingPolicyConnectorValuePostRequestBody("Type_example", "naming policy value name", "-")} // PostNamingPolicyValueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PostNamingPolicyValue(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).PostNamingPolicyValueRequest(postNamingPolicyValueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PostNamingPolicyValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNamingPolicyValue`: GetNamingPolicyValues200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PostNamingPolicyValue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNamingPolicyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **postNamingPolicyValueRequest** | [**PostNamingPolicyValueRequest**](PostNamingPolicyValueRequest.md) |  | 

### Return type

[**GetNamingPolicyValues200ResponseDataInner**](GetNamingPolicyValues200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNamingPolicy

> GetNamingPolicies200ResponseDataInner PutNamingPolicy(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyPutRequestBody(namingPolicyPutRequestBody).Execute()

Update a naming policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	namingPolicyPutRequestBody := *openapiclient.NewNamingPolicyPutRequestBody("naming policy name") // NamingPolicyPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PutNamingPolicy(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyPutRequestBody(namingPolicyPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PutNamingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNamingPolicy`: GetNamingPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PutNamingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNamingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **namingPolicyPutRequestBody** | [**NamingPolicyPutRequestBody**](NamingPolicyPutRequestBody.md) |  | 

### Return type

[**GetNamingPolicies200ResponseDataInner**](GetNamingPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNamingPolicyRestriction

> GetNamingPolicyRestrictions200ResponseDataInner PutNamingPolicyRestriction(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyRestrictionPutRequestBody(namingPolicyRestrictionPutRequestBody).Execute()

Update a naming restriction



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	namingPolicyRestrictionPutRequestBody := *openapiclient.NewNamingPolicyRestrictionPutRequestBody("name", []openapiclient.ValueMatchTypePairBean{*openapiclient.NewValueMatchTypePairBean()}) // NamingPolicyRestrictionPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PutNamingPolicyRestriction(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).NamingPolicyRestrictionPutRequestBody(namingPolicyRestrictionPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PutNamingPolicyRestriction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNamingPolicyRestriction`: GetNamingPolicyRestrictions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PutNamingPolicyRestriction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNamingPolicyRestrictionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **namingPolicyRestrictionPutRequestBody** | [**NamingPolicyRestrictionPutRequestBody**](NamingPolicyRestrictionPutRequestBody.md) |  | 

### Return type

[**GetNamingPolicyRestrictions200ResponseDataInner**](GetNamingPolicyRestrictions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNamingPolicyValue

> GetNamingPolicyValues200ResponseDataInner PutNamingPolicyValue(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PutNamingPolicyValueRequest(putNamingPolicyValueRequest).Execute()

Update a naming policy value



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/umich-vci/gobam2"
)

func main() {
	id := int64(12345) // int64 | The ID of the resource.
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	putNamingPolicyValueRequest := openapiclient.putNamingPolicyValue_request{NamingPolicyConnectorValuePutRequestBody: openapiclient.NewNamingPolicyConnectorValuePutRequestBody("Type_example", "naming policy value name", "-")} // PutNamingPolicyValueRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NamingPolicyResourcesAPI.PutNamingPolicyValue(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutNamingPolicyValueRequest(putNamingPolicyValueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NamingPolicyResourcesAPI.PutNamingPolicyValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNamingPolicyValue`: GetNamingPolicyValues200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NamingPolicyResourcesAPI.PutNamingPolicyValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNamingPolicyValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **putNamingPolicyValueRequest** | [**PutNamingPolicyValueRequest**](PutNamingPolicyValueRequest.md) |  | 

### Return type

[**GetNamingPolicyValues200ResponseDataInner**](GetNamingPolicyValues200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

