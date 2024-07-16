# \DeploymentOptionResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeploymentOption**](DeploymentOptionResourcesAPI.md#DeleteDeploymentOption) | **Delete** /api/v2/deploymentOptions/{id} | Delete a deployment option
[**DeleteDeploymentOptionDefinition**](DeploymentOptionResourcesAPI.md#DeleteDeploymentOptionDefinition) | **Delete** /api/v2/deploymentOptionDefinitions/{id} | Delete a deployment option definition
[**GetCollectionDeploymentOptionDefinitions**](DeploymentOptionResourcesAPI.md#GetCollectionDeploymentOptionDefinitions) | **Get** /api/v2/{collection}/{collectionId}/deploymentOptionDefinitions | Retrieve a collection of deployment option definitions
[**GetCollectionDeploymentOptions**](DeploymentOptionResourcesAPI.md#GetCollectionDeploymentOptions) | **Get** /api/v2/{collection}/{collectionId}/deploymentOptions | Retrieve a collection of deployment options
[**GetDeploymentOption**](DeploymentOptionResourcesAPI.md#GetDeploymentOption) | **Get** /api/v2/deploymentOptions/{id} | Retrieve a single deployment option
[**GetDeploymentOptionDefinition**](DeploymentOptionResourcesAPI.md#GetDeploymentOptionDefinition) | **Get** /api/v2/deploymentOptionDefinitions/{id} | Retrieve a single deployment option definition
[**GetDeploymentOptionDefinitions**](DeploymentOptionResourcesAPI.md#GetDeploymentOptionDefinitions) | **Get** /api/v2/deploymentOptionDefinitions | Retrieve a collection of deployment option definitions
[**GetDeploymentOptions**](DeploymentOptionResourcesAPI.md#GetDeploymentOptions) | **Get** /api/v2/deploymentOptions | Retrieve a collection of deployment options
[**PostCollectionDeploymentOption**](DeploymentOptionResourcesAPI.md#PostCollectionDeploymentOption) | **Post** /api/v2/{collection}/{collectionId}/deploymentOptions | Create a new deployment option
[**PostCollectionDeploymentOptionDefinition**](DeploymentOptionResourcesAPI.md#PostCollectionDeploymentOptionDefinition) | **Post** /api/v2/{collection}/{collectionId}/deploymentOptionDefinitions | Create a new deployment option definition
[**PutDeploymentOption**](DeploymentOptionResourcesAPI.md#PutDeploymentOption) | **Put** /api/v2/deploymentOptions/{id} | Update a deployment option
[**PutDeploymentOptionDefinition**](DeploymentOptionResourcesAPI.md#PutDeploymentOptionDefinition) | **Put** /api/v2/deploymentOptionDefinitions/{id} | Update a deployment option definition



## DeleteDeploymentOption

> DeleteDeploymentOption(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a deployment option



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
	r, err := apiClient.DeploymentOptionResourcesAPI.DeleteDeploymentOption(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.DeleteDeploymentOption``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeploymentOptionRequest struct via the builder pattern


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


## DeleteDeploymentOptionDefinition

> DeleteDeploymentOptionDefinition(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a deployment option definition



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
	r, err := apiClient.DeploymentOptionResourcesAPI.DeleteDeploymentOptionDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.DeleteDeploymentOptionDefinition``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeploymentOptionDefinitionRequest struct via the builder pattern


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


## GetCollectionDeploymentOptionDefinitions

> GetCollectionDeploymentOptionDefinitions200Response GetCollectionDeploymentOptionDefinitions(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of deployment option definitions



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
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.GetCollectionDeploymentOptionDefinitions(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.GetCollectionDeploymentOptionDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionDeploymentOptionDefinitions`: GetCollectionDeploymentOptionDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.GetCollectionDeploymentOptionDefinitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionDeploymentOptionDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetCollectionDeploymentOptionDefinitions200Response**](GetCollectionDeploymentOptionDefinitions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionDeploymentOptions

> GetCollectionDeploymentOptions200Response GetCollectionDeploymentOptions(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of deployment options



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
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.GetCollectionDeploymentOptions(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.GetCollectionDeploymentOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionDeploymentOptions`: GetCollectionDeploymentOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.GetCollectionDeploymentOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionDeploymentOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetCollectionDeploymentOptions200Response**](GetCollectionDeploymentOptions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentOption

> GetDeploymentOptions200ResponseDataInner GetDeploymentOption(ctx, id).Fields(fields).Execute()

Retrieve a single deployment option



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
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.GetDeploymentOption(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.GetDeploymentOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentOption`: GetDeploymentOptions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.GetDeploymentOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetDeploymentOptions200ResponseDataInner**](GetDeploymentOptions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentOptionDefinition

> GetDeploymentOptionDefinitions200ResponseDataInner GetDeploymentOptionDefinition(ctx, id).Fields(fields).Execute()

Retrieve a single deployment option definition



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
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.GetDeploymentOptionDefinition(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.GetDeploymentOptionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentOptionDefinition`: GetDeploymentOptionDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.GetDeploymentOptionDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentOptionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetDeploymentOptionDefinitions200ResponseDataInner**](GetDeploymentOptionDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentOptionDefinitions

> GetDeploymentOptionDefinitions200Response GetDeploymentOptionDefinitions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of deployment option definitions



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
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.GetDeploymentOptionDefinitions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.GetDeploymentOptionDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentOptionDefinitions`: GetDeploymentOptionDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.GetDeploymentOptionDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentOptionDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetDeploymentOptionDefinitions200Response**](GetDeploymentOptionDefinitions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploymentOptions

> GetDeploymentOptions200Response GetDeploymentOptions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of deployment options



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
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.GetDeploymentOptions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.GetDeploymentOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentOptions`: GetDeploymentOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.GetDeploymentOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploymentOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetDeploymentOptions200Response**](GetDeploymentOptions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionDeploymentOption

> GetDeploymentOptions200ResponseDataInner PostCollectionDeploymentOption(ctx, collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostCollectionDeploymentOptionRequest(postCollectionDeploymentOptionRequest).Execute()

Create a new deployment option



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
	postCollectionDeploymentOptionRequest := openapiclient.postCollectionDeploymentOption_request{DHCPVendorOptionPostRequestBody: openapiclient.NewDHCPVendorOptionPostRequestBody("Type_example", *openapiclient.NewInlinedDHCPVendorOptionDefinition())} // PostCollectionDeploymentOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.PostCollectionDeploymentOption(context.Background(), collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostCollectionDeploymentOptionRequest(postCollectionDeploymentOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.PostCollectionDeploymentOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollectionDeploymentOption`: GetDeploymentOptions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.PostCollectionDeploymentOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionDeploymentOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **postCollectionDeploymentOptionRequest** | [**PostCollectionDeploymentOptionRequest**](PostCollectionDeploymentOptionRequest.md) |  | 

### Return type

[**GetDeploymentOptions200ResponseDataInner**](GetDeploymentOptions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionDeploymentOptionDefinition

> PutDeploymentOptionDefinition200Response PostCollectionDeploymentOptionDefinition(ctx, collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostCollectionDeploymentOptionDefinitionRequest(postCollectionDeploymentOptionDefinitionRequest).Execute()

Create a new deployment option definition



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
	postCollectionDeploymentOptionDefinitionRequest := openapiclient.postCollectionDeploymentOptionDefinition_request{DHCPVendorOptionDefinitionPostRequestBody: openapiclient.NewDHCPVendorOptionDefinitionPostRequestBody("Type_example", "allow-recursion", int32(123), "DisplayName_example", "Description_example", map[string]interface{}(123))} // PostCollectionDeploymentOptionDefinitionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.PostCollectionDeploymentOptionDefinition(context.Background(), collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostCollectionDeploymentOptionDefinitionRequest(postCollectionDeploymentOptionDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.PostCollectionDeploymentOptionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollectionDeploymentOptionDefinition`: PutDeploymentOptionDefinition200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.PostCollectionDeploymentOptionDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionDeploymentOptionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **postCollectionDeploymentOptionDefinitionRequest** | [**PostCollectionDeploymentOptionDefinitionRequest**](PostCollectionDeploymentOptionDefinitionRequest.md) |  | 

### Return type

[**PutDeploymentOptionDefinition200Response**](PutDeploymentOptionDefinition200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeploymentOption

> GetDeploymentOptions200ResponseDataInner PutDeploymentOption(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PutDeploymentOptionRequest(putDeploymentOptionRequest).Execute()

Update a deployment option



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
	putDeploymentOptionRequest := openapiclient.putDeploymentOption_request{DHCPVendorOptionPutRequestBody: openapiclient.NewDHCPVendorOptionPutRequestBody("Type_example")} // PutDeploymentOptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.PutDeploymentOption(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutDeploymentOptionRequest(putDeploymentOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.PutDeploymentOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDeploymentOption`: GetDeploymentOptions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.PutDeploymentOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **putDeploymentOptionRequest** | [**PutDeploymentOptionRequest**](PutDeploymentOptionRequest.md) |  | 

### Return type

[**GetDeploymentOptions200ResponseDataInner**](GetDeploymentOptions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDeploymentOptionDefinition

> PutDeploymentOptionDefinition200Response PutDeploymentOptionDefinition(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PutDeploymentOptionDefinitionRequest(putDeploymentOptionDefinitionRequest).Execute()

Update a deployment option definition



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
	putDeploymentOptionDefinitionRequest := openapiclient.putDeploymentOptionDefinition_request{DHCPVendorOptionDefinitionPutRequestBody: openapiclient.NewDHCPVendorOptionDefinitionPutRequestBody("Type_example", "allow-recursion", "DisplayName_example", "Description_example")} // PutDeploymentOptionDefinitionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentOptionResourcesAPI.PutDeploymentOptionDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutDeploymentOptionDefinitionRequest(putDeploymentOptionDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentOptionResourcesAPI.PutDeploymentOptionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDeploymentOptionDefinition`: PutDeploymentOptionDefinition200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentOptionResourcesAPI.PutDeploymentOptionDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeploymentOptionDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **putDeploymentOptionDefinitionRequest** | [**PutDeploymentOptionDefinitionRequest**](PutDeploymentOptionDefinitionRequest.md) |  | 

### Return type

[**PutDeploymentOptionDefinition200Response**](PutDeploymentOptionDefinition200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

