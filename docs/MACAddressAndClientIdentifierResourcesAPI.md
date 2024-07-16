# \MACAddressAndClientIdentifierResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#DeleteMacAddress) | **Delete** /api/v2/macAddresses/{id} | Delete a MAC address
[**DeleteMacPool**](MACAddressAndClientIdentifierResourcesAPI.md#DeleteMacPool) | **Delete** /api/v2/macPools/{id} | Delete a MAC pool
[**DeleteMacPoolMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#DeleteMacPoolMacAddress) | **Delete** /api/v2/macPools/{collectionId}/macAddresses/{id} | Unlink a MAC address from a MAC pool
[**GetClientIdentifier**](MACAddressAndClientIdentifierResourcesAPI.md#GetClientIdentifier) | **Get** /api/v2/clientIdentifiers/{id} | Retrieve a single client identifier
[**GetClientIdentifiers**](MACAddressAndClientIdentifierResourcesAPI.md#GetClientIdentifiers) | **Get** /api/v2/clientIdentifiers | Retrieve a collection of client identifiers
[**GetConfigurationClientIdentifiers**](MACAddressAndClientIdentifierResourcesAPI.md#GetConfigurationClientIdentifiers) | **Get** /api/v2/configurations/{collectionId}/clientIdentifiers | Retrieve a collection of client identifiers
[**GetConfigurationMacAddresses**](MACAddressAndClientIdentifierResourcesAPI.md#GetConfigurationMacAddresses) | **Get** /api/v2/configurations/{collectionId}/macAddresses | Retrieve a collection of MAC addresses
[**GetConfigurationMacPools**](MACAddressAndClientIdentifierResourcesAPI.md#GetConfigurationMacPools) | **Get** /api/v2/configurations/{collectionId}/macPools | Retrieve a collection of MAC pools
[**GetMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#GetMacAddress) | **Get** /api/v2/macAddresses/{id} | Retrieve a single MAC address
[**GetMacAddresses**](MACAddressAndClientIdentifierResourcesAPI.md#GetMacAddresses) | **Get** /api/v2/macAddresses | Retrieve a collection of MAC addresses
[**GetMacPool**](MACAddressAndClientIdentifierResourcesAPI.md#GetMacPool) | **Get** /api/v2/macPools/{id} | Retrieve a single MAC pool
[**GetMacPoolMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#GetMacPoolMacAddress) | **Get** /api/v2/macPools/{collectionId}/macAddresses/{id} | Retrieve a single linked MAC address
[**GetMacPoolMacAddresses**](MACAddressAndClientIdentifierResourcesAPI.md#GetMacPoolMacAddresses) | **Get** /api/v2/macPools/{collectionId}/macAddresses | Retrieve a collection of linked MAC addresses
[**GetMacPools**](MACAddressAndClientIdentifierResourcesAPI.md#GetMacPools) | **Get** /api/v2/macPools | Retrieve a collection of MAC pools
[**PostConfigurationMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#PostConfigurationMacAddress) | **Post** /api/v2/configurations/{collectionId}/macAddresses | Create a new MAC address
[**PostConfigurationMacPool**](MACAddressAndClientIdentifierResourcesAPI.md#PostConfigurationMacPool) | **Post** /api/v2/configurations/{collectionId}/macPools | Create a new MAC pool
[**PostMacPoolMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#PostMacPoolMacAddress) | **Post** /api/v2/macPools/{collectionId}/macAddresses | Link a MAC address to a MAC pool
[**PutMacAddress**](MACAddressAndClientIdentifierResourcesAPI.md#PutMacAddress) | **Put** /api/v2/macAddresses/{id} | Update a MAC address
[**PutMacPool**](MACAddressAndClientIdentifierResourcesAPI.md#PutMacPool) | **Put** /api/v2/macPools/{id} | Update a MAC pool



## DeleteMacAddress

> DeleteMacAddress(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a MAC address



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
	r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.DeleteMacAddress(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.DeleteMacAddress``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMacAddressRequest struct via the builder pattern


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


## DeleteMacPool

> DeleteMacPool(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a MAC pool



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
	r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.DeleteMacPool(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.DeleteMacPool``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMacPoolRequest struct via the builder pattern


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


## DeleteMacPoolMacAddress

> DeleteMacPoolMacAddress(ctx, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Unlink a MAC address from a MAC pool



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
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.DeleteMacPoolMacAddress(context.Background(), collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.DeleteMacPoolMacAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMacPoolMacAddressRequest struct via the builder pattern


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


## GetClientIdentifier

> GetClientIdentifiers200ResponseDataInner GetClientIdentifier(ctx, id).Fields(fields).Execute()

Retrieve a single client identifier



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifier(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientIdentifier`: GetClientIdentifiers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetClientIdentifiers200ResponseDataInner**](GetClientIdentifiers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientIdentifiers

> GetClientIdentifiers200Response GetClientIdentifiers(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of client identifiers



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifiers(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientIdentifiers`: GetClientIdentifiers200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetClientIdentifiers200Response**](GetClientIdentifiers200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationClientIdentifiers

> GetClientIdentifiers200Response GetConfigurationClientIdentifiers(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of client identifiers



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetConfigurationClientIdentifiers(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetConfigurationClientIdentifiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationClientIdentifiers`: GetClientIdentifiers200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetConfigurationClientIdentifiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationClientIdentifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetClientIdentifiers200Response**](GetClientIdentifiers200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationMacAddresses

> GetMacAddresses200Response GetConfigurationMacAddresses(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of MAC addresses



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacAddresses(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationMacAddresses`: GetMacAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationMacAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetMacAddresses200Response**](GetMacAddresses200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationMacPools

> GetMacPools200Response GetConfigurationMacPools(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of MAC pools



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacPools(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationMacPools`: GetMacPools200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationMacPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetMacPools200Response**](GetMacPools200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacAddress

> GetMacAddresses200ResponseDataInner GetMacAddress(ctx, id).Fields(fields).Execute()

Retrieve a single MAC address



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacAddress(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetMacAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacAddress`: GetMacAddresses200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetMacAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetMacAddresses200ResponseDataInner**](GetMacAddresses200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacAddresses

> GetMacAddresses200Response GetMacAddresses(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of MAC addresses



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacAddresses(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetMacAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacAddresses`: GetMacAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetMacAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMacAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetMacAddresses200Response**](GetMacAddresses200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacPool

> GetMacPools200ResponseDataInner GetMacPool(ctx, id).Fields(fields).Execute()

Retrieve a single MAC pool



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPool(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetMacPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacPool`: GetMacPools200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetMacPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetMacPools200ResponseDataInner**](GetMacPools200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacPoolMacAddress

> GetMacAddresses200ResponseDataInner GetMacPoolMacAddress(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single linked MAC address



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddress(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacPoolMacAddress`: GetMacAddresses200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacPoolMacAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetMacAddresses200ResponseDataInner**](GetMacAddresses200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacPoolMacAddresses

> GetMacAddresses200Response GetMacPoolMacAddresses(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of linked MAC addresses



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddresses(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacPoolMacAddresses`: GetMacAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMacPoolMacAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetMacAddresses200Response**](GetMacAddresses200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMacPools

> GetMacPools200Response GetMacPools(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of MAC pools



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
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPools(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.GetMacPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMacPools`: GetMacPools200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.GetMacPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMacPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetMacPools200Response**](GetMacPools200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConfigurationMacAddress

> GetMacAddresses200ResponseDataInner PostConfigurationMacAddress(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).MACAddressPostRequestBody(mACAddressPostRequestBody).Execute()

Create a new MAC address



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
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	mACAddressPostRequestBody := *openapiclient.NewMACAddressPostRequestBody("00-00-5E-00-53-AF") // MACAddressPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacAddress(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).MACAddressPostRequestBody(mACAddressPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConfigurationMacAddress`: GetMacAddresses200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConfigurationMacAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **mACAddressPostRequestBody** | [**MACAddressPostRequestBody**](MACAddressPostRequestBody.md) |  | 

### Return type

[**GetMacAddresses200ResponseDataInner**](GetMacAddresses200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConfigurationMacPool

> PutMacPool200Response PostConfigurationMacPool(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).MACPoolPostRequestBody(mACPoolPostRequestBody).Execute()

Create a new MAC pool



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
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	mACPoolPostRequestBody := *openapiclient.NewMACPoolPostRequestBody("Type_example", "name") // MACPoolPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacPool(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).MACPoolPostRequestBody(mACPoolPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConfigurationMacPool`: PutMacPool200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConfigurationMacPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **mACPoolPostRequestBody** | [**MACPoolPostRequestBody**](MACPoolPostRequestBody.md) |  | 

### Return type

[**PutMacPool200Response**](PutMacPool200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMacPoolMacAddress

> GetMacAddresses200ResponseDataInner PostMacPoolMacAddress(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedMACAddressPostRequestBody(linkedMACAddressPostRequestBody).Execute()

Link a MAC address to a MAC pool



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
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)
	linkedMACAddressPostRequestBody := *openapiclient.NewLinkedMACAddressPostRequestBody(int64(103307)) // LinkedMACAddressPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PostMacPoolMacAddress(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedMACAddressPostRequestBody(linkedMACAddressPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.PostMacPoolMacAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMacPoolMacAddress`: GetMacAddresses200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.PostMacPoolMacAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMacPoolMacAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **linkedMACAddressPostRequestBody** | [**LinkedMACAddressPostRequestBody**](LinkedMACAddressPostRequestBody.md) |  | 

### Return type

[**GetMacAddresses200ResponseDataInner**](GetMacAddresses200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMacAddress

> GetMacAddresses200ResponseDataInner PutMacAddress(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).MACAddressPutRequestBody(mACAddressPutRequestBody).Execute()

Update a MAC address



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
	mACAddressPutRequestBody := *openapiclient.NewMACAddressPutRequestBody("00-00-5E-00-53-AF") // MACAddressPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PutMacAddress(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).MACAddressPutRequestBody(mACAddressPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.PutMacAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMacAddress`: GetMacAddresses200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.PutMacAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMacAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **mACAddressPutRequestBody** | [**MACAddressPutRequestBody**](MACAddressPutRequestBody.md) |  | 

### Return type

[**GetMacAddresses200ResponseDataInner**](GetMacAddresses200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMacPool

> PutMacPool200Response PutMacPool(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).MACPoolPutRequestBody(mACPoolPutRequestBody).Execute()

Update a MAC pool



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
	mACPoolPutRequestBody := *openapiclient.NewMACPoolPutRequestBody("Type_example", "name", false) // MACPoolPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PutMacPool(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).MACPoolPutRequestBody(mACPoolPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MACAddressAndClientIdentifierResourcesAPI.PutMacPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMacPool`: PutMacPool200Response
	fmt.Fprintf(os.Stdout, "Response from `MACAddressAndClientIdentifierResourcesAPI.PutMacPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMacPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **mACPoolPutRequestBody** | [**MACPoolPutRequestBody**](MACPoolPutRequestBody.md) |  | 

### Return type

[**PutMacPool200Response**](PutMacPool200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

