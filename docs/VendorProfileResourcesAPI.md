# \VendorProfileResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDeploymentOptionDefinition**](VendorProfileResourcesAPI.md#DeleteDeploymentOptionDefinition) | **Delete** /api/v2/deploymentOptionDefinitions/{id} | Delete a deployment option definition
[**DeleteVendorProfile**](VendorProfileResourcesAPI.md#DeleteVendorProfile) | **Delete** /api/v2/vendorProfiles/{id} | Delete a vendor profile
[**GetCollectionDeploymentOptionDefinitions**](VendorProfileResourcesAPI.md#GetCollectionDeploymentOptionDefinitions) | **Get** /api/v2/{collection}/{collectionId}/deploymentOptionDefinitions | Retrieve a collection of deployment option definitions
[**GetDeploymentOptionDefinition**](VendorProfileResourcesAPI.md#GetDeploymentOptionDefinition) | **Get** /api/v2/deploymentOptionDefinitions/{id} | Retrieve a single deployment option definition
[**GetDeploymentOptionDefinitions**](VendorProfileResourcesAPI.md#GetDeploymentOptionDefinitions) | **Get** /api/v2/deploymentOptionDefinitions | Retrieve a collection of deployment option definitions
[**GetVendorProfile**](VendorProfileResourcesAPI.md#GetVendorProfile) | **Get** /api/v2/vendorProfiles/{id} | Retrieve a single vendor profile
[**GetVendorProfiles**](VendorProfileResourcesAPI.md#GetVendorProfiles) | **Get** /api/v2/vendorProfiles | Retrieve a collection of vendor profiles
[**PostCollectionDeploymentOptionDefinition**](VendorProfileResourcesAPI.md#PostCollectionDeploymentOptionDefinition) | **Post** /api/v2/{collection}/{collectionId}/deploymentOptionDefinitions | Create a new deployment option definition
[**PostVendorProfile**](VendorProfileResourcesAPI.md#PostVendorProfile) | **Post** /api/v2/vendorProfiles | Create a new vendor profile
[**PutDeploymentOptionDefinition**](VendorProfileResourcesAPI.md#PutDeploymentOptionDefinition) | **Put** /api/v2/deploymentOptionDefinitions/{id} | Update a deployment option definition
[**PutVendorProfile**](VendorProfileResourcesAPI.md#PutVendorProfile) | **Put** /api/v2/vendorProfiles/{id} | Update a vendor profile



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
	r, err := apiClient.VendorProfileResourcesAPI.DeleteDeploymentOptionDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.DeleteDeploymentOptionDefinition``: %v\n", err)
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


## DeleteVendorProfile

> DeleteVendorProfile(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a vendor profile



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
	r, err := apiClient.VendorProfileResourcesAPI.DeleteVendorProfile(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.DeleteVendorProfile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVendorProfileRequest struct via the builder pattern


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
	resp, r, err := apiClient.VendorProfileResourcesAPI.GetCollectionDeploymentOptionDefinitions(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.GetCollectionDeploymentOptionDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionDeploymentOptionDefinitions`: GetCollectionDeploymentOptionDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.GetCollectionDeploymentOptionDefinitions`: %v\n", resp)
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
	resp, r, err := apiClient.VendorProfileResourcesAPI.GetDeploymentOptionDefinition(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.GetDeploymentOptionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentOptionDefinition`: GetDeploymentOptionDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.GetDeploymentOptionDefinition`: %v\n", resp)
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
	resp, r, err := apiClient.VendorProfileResourcesAPI.GetDeploymentOptionDefinitions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.GetDeploymentOptionDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploymentOptionDefinitions`: GetDeploymentOptionDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.GetDeploymentOptionDefinitions`: %v\n", resp)
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


## GetVendorProfile

> GetVendorProfiles200ResponseDataInner GetVendorProfile(ctx, id).Fields(fields).Execute()

Retrieve a single vendor profile



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
	resp, r, err := apiClient.VendorProfileResourcesAPI.GetVendorProfile(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.GetVendorProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorProfile`: GetVendorProfiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.GetVendorProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetVendorProfiles200ResponseDataInner**](GetVendorProfiles200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVendorProfiles

> GetVendorProfiles200Response GetVendorProfiles(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of vendor profiles



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
	resp, r, err := apiClient.VendorProfileResourcesAPI.GetVendorProfiles(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.GetVendorProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVendorProfiles`: GetVendorProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.GetVendorProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVendorProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetVendorProfiles200Response**](GetVendorProfiles200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

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
	resp, r, err := apiClient.VendorProfileResourcesAPI.PostCollectionDeploymentOptionDefinition(context.Background(), collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostCollectionDeploymentOptionDefinitionRequest(postCollectionDeploymentOptionDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.PostCollectionDeploymentOptionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollectionDeploymentOptionDefinition`: PutDeploymentOptionDefinition200Response
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.PostCollectionDeploymentOptionDefinition`: %v\n", resp)
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


## PostVendorProfile

> GetVendorProfiles200ResponseDataInner PostVendorProfile(ctx).XBcnChangeControlComment(xBcnChangeControlComment).VendorProfilePostRequestBody(vendorProfilePostRequestBody).Execute()

Create a new vendor profile



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
	vendorProfilePostRequestBody := *openapiclient.NewVendorProfilePostRequestBody("Cisco IP phones", "Cisco Systems, Inc. IP Phone CP-7940G", "Vendor profile for Cisco IP phones") // VendorProfilePostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorProfileResourcesAPI.PostVendorProfile(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).VendorProfilePostRequestBody(vendorProfilePostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.PostVendorProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVendorProfile`: GetVendorProfiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.PostVendorProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVendorProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **vendorProfilePostRequestBody** | [**VendorProfilePostRequestBody**](VendorProfilePostRequestBody.md) |  | 

### Return type

[**GetVendorProfiles200ResponseDataInner**](GetVendorProfiles200ResponseDataInner.md)

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
	resp, r, err := apiClient.VendorProfileResourcesAPI.PutDeploymentOptionDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutDeploymentOptionDefinitionRequest(putDeploymentOptionDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.PutDeploymentOptionDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDeploymentOptionDefinition`: PutDeploymentOptionDefinition200Response
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.PutDeploymentOptionDefinition`: %v\n", resp)
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


## PutVendorProfile

> GetVendorProfiles200ResponseDataInner PutVendorProfile(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).VendorProfilePutRequestBody(vendorProfilePutRequestBody).Execute()

Update a vendor profile



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
	vendorProfilePutRequestBody := *openapiclient.NewVendorProfilePutRequestBody("Cisco IP phones", "Cisco Systems, Inc. IP Phone CP-7940G", "Vendor profile for Cisco IP phones") // VendorProfilePutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VendorProfileResourcesAPI.PutVendorProfile(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).VendorProfilePutRequestBody(vendorProfilePutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VendorProfileResourcesAPI.PutVendorProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVendorProfile`: GetVendorProfiles200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `VendorProfileResourcesAPI.PutVendorProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVendorProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **vendorProfilePutRequestBody** | [**VendorProfilePutRequestBody**](VendorProfilePutRequestBody.md) |  | 

### Return type

[**GetVendorProfiles200ResponseDataInner**](GetVendorProfiles200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

