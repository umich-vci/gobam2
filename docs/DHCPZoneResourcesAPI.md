# \DHCPZoneResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteZoneDeclaration**](DHCPZoneResourcesAPI.md#DeleteZoneDeclaration) | **Delete** /api/v2/zoneDeclarations/{id} | Delete a zone declaration
[**DeleteZoneGroup**](DHCPZoneResourcesAPI.md#DeleteZoneGroup) | **Delete** /api/v2/zoneGroups/{id} | Delete a zone group
[**GetConfigurationZoneGroups**](DHCPZoneResourcesAPI.md#GetConfigurationZoneGroups) | **Get** /api/v2/configurations/{collectionId}/zoneGroups | Retrieve a collection of zone groups
[**GetZoneDeclaration**](DHCPZoneResourcesAPI.md#GetZoneDeclaration) | **Get** /api/v2/zoneDeclarations/{id} | Retrieve a single zone declaration
[**GetZoneDeclarations**](DHCPZoneResourcesAPI.md#GetZoneDeclarations) | **Get** /api/v2/zoneDeclarations | Retrieve a collection of zone declarations
[**GetZoneGroup**](DHCPZoneResourcesAPI.md#GetZoneGroup) | **Get** /api/v2/zoneGroups/{id} | Retrieve a single zone group
[**GetZoneGroupZoneDeclarations**](DHCPZoneResourcesAPI.md#GetZoneGroupZoneDeclarations) | **Get** /api/v2/zoneGroups/{collectionId}/zoneDeclarations | Retrieve a collection of zone declarations
[**GetZoneGroups**](DHCPZoneResourcesAPI.md#GetZoneGroups) | **Get** /api/v2/zoneGroups | Retrieve a collection of zone groups
[**PostConfigurationZoneGroup**](DHCPZoneResourcesAPI.md#PostConfigurationZoneGroup) | **Post** /api/v2/configurations/{collectionId}/zoneGroups | Create a new zone group
[**PostZoneGroupZoneDeclaration**](DHCPZoneResourcesAPI.md#PostZoneGroupZoneDeclaration) | **Post** /api/v2/zoneGroups/{collectionId}/zoneDeclarations | Create a new zone declaration
[**PutZoneDeclaration**](DHCPZoneResourcesAPI.md#PutZoneDeclaration) | **Put** /api/v2/zoneDeclarations/{id} | Update a zone declaration
[**PutZoneGroup**](DHCPZoneResourcesAPI.md#PutZoneGroup) | **Put** /api/v2/zoneGroups/{id} | Update a zone group



## DeleteZoneDeclaration

> DeleteZoneDeclaration(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a zone declaration



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
	r, err := apiClient.DHCPZoneResourcesAPI.DeleteZoneDeclaration(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.DeleteZoneDeclaration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteZoneDeclarationRequest struct via the builder pattern


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


## DeleteZoneGroup

> DeleteZoneGroup(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a zone group



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
	r, err := apiClient.DHCPZoneResourcesAPI.DeleteZoneGroup(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.DeleteZoneGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteZoneGroupRequest struct via the builder pattern


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


## GetConfigurationZoneGroups

> GetZoneGroups200Response GetConfigurationZoneGroups(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of zone groups



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
	resp, r, err := apiClient.DHCPZoneResourcesAPI.GetConfigurationZoneGroups(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.GetConfigurationZoneGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationZoneGroups`: GetZoneGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.GetConfigurationZoneGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationZoneGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetZoneGroups200Response**](GetZoneGroups200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneDeclaration

> GetZoneDeclarations200ResponseDataInner GetZoneDeclaration(ctx, id).Fields(fields).Execute()

Retrieve a single zone declaration



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
	resp, r, err := apiClient.DHCPZoneResourcesAPI.GetZoneDeclaration(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.GetZoneDeclaration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneDeclaration`: GetZoneDeclarations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.GetZoneDeclaration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneDeclarationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetZoneDeclarations200ResponseDataInner**](GetZoneDeclarations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneDeclarations

> GetZoneDeclarations200Response GetZoneDeclarations(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of zone declarations



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
	resp, r, err := apiClient.DHCPZoneResourcesAPI.GetZoneDeclarations(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.GetZoneDeclarations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneDeclarations`: GetZoneDeclarations200Response
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.GetZoneDeclarations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneDeclarationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetZoneDeclarations200Response**](GetZoneDeclarations200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneGroup

> GetZoneGroups200ResponseDataInner GetZoneGroup(ctx, id).Fields(fields).Execute()

Retrieve a single zone group



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
	resp, r, err := apiClient.DHCPZoneResourcesAPI.GetZoneGroup(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.GetZoneGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneGroup`: GetZoneGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.GetZoneGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetZoneGroups200ResponseDataInner**](GetZoneGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneGroupZoneDeclarations

> GetZoneDeclarations200Response GetZoneGroupZoneDeclarations(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of zone declarations



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
	resp, r, err := apiClient.DHCPZoneResourcesAPI.GetZoneGroupZoneDeclarations(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.GetZoneGroupZoneDeclarations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneGroupZoneDeclarations`: GetZoneDeclarations200Response
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.GetZoneGroupZoneDeclarations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneGroupZoneDeclarationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetZoneDeclarations200Response**](GetZoneDeclarations200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetZoneGroups

> GetZoneGroups200Response GetZoneGroups(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of zone groups



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
	resp, r, err := apiClient.DHCPZoneResourcesAPI.GetZoneGroups(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.GetZoneGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetZoneGroups`: GetZoneGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.GetZoneGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetZoneGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetZoneGroups200Response**](GetZoneGroups200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConfigurationZoneGroup

> GetZoneGroups200ResponseDataInner PostConfigurationZoneGroup(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).DHCPZoneGroupPostRequestBody(dHCPZoneGroupPostRequestBody).Execute()

Create a new zone group



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
	dHCPZoneGroupPostRequestBody := *openapiclient.NewDHCPZoneGroupPostRequestBody("Zone Group 1") // DHCPZoneGroupPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPZoneResourcesAPI.PostConfigurationZoneGroup(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).DHCPZoneGroupPostRequestBody(dHCPZoneGroupPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.PostConfigurationZoneGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConfigurationZoneGroup`: GetZoneGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.PostConfigurationZoneGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConfigurationZoneGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **dHCPZoneGroupPostRequestBody** | [**DHCPZoneGroupPostRequestBody**](DHCPZoneGroupPostRequestBody.md) |  | 

### Return type

[**GetZoneGroups200ResponseDataInner**](GetZoneGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostZoneGroupZoneDeclaration

> GetZoneDeclarations200ResponseDataInner PostZoneGroupZoneDeclaration(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostZoneGroupZoneDeclarationRequest(postZoneGroupZoneDeclarationRequest).Execute()

Create a new zone declaration



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
	postZoneGroupZoneDeclarationRequest := openapiclient.postZoneGroupZoneDeclaration_request{DHCPForwardZonePostRequestBody: openapiclient.NewDHCPForwardZonePostRequestBody("Type_example", "10.10.10.10", *openapiclient.NewDHCPForwardZoneAllOfZone(int64(103307), "example.corp"))} // PostZoneGroupZoneDeclarationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPZoneResourcesAPI.PostZoneGroupZoneDeclaration(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).PostZoneGroupZoneDeclarationRequest(postZoneGroupZoneDeclarationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.PostZoneGroupZoneDeclaration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostZoneGroupZoneDeclaration`: GetZoneDeclarations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.PostZoneGroupZoneDeclaration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostZoneGroupZoneDeclarationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **postZoneGroupZoneDeclarationRequest** | [**PostZoneGroupZoneDeclarationRequest**](PostZoneGroupZoneDeclarationRequest.md) |  | 

### Return type

[**GetZoneDeclarations200ResponseDataInner**](GetZoneDeclarations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutZoneDeclaration

> GetZoneDeclarations200ResponseDataInner PutZoneDeclaration(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PutZoneDeclarationRequest(putZoneDeclarationRequest).Execute()

Update a zone declaration



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
	putZoneDeclarationRequest := openapiclient.putZoneDeclaration_request{DHCPForwardZonePutRequestBody: openapiclient.NewDHCPForwardZonePutRequestBody("Type_example", "10.10.10.10", "DynamicUpdateSigningKeyType_example", *openapiclient.NewDHCPForwardZoneAllOfZone(int64(103307), "example.corp"))} // PutZoneDeclarationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPZoneResourcesAPI.PutZoneDeclaration(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutZoneDeclarationRequest(putZoneDeclarationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.PutZoneDeclaration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutZoneDeclaration`: GetZoneDeclarations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.PutZoneDeclaration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneDeclarationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **putZoneDeclarationRequest** | [**PutZoneDeclarationRequest**](PutZoneDeclarationRequest.md) |  | 

### Return type

[**GetZoneDeclarations200ResponseDataInner**](GetZoneDeclarations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutZoneGroup

> GetZoneGroups200ResponseDataInner PutZoneGroup(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).DHCPZoneGroupPutRequestBody(dHCPZoneGroupPutRequestBody).Execute()

Update a zone group



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
	dHCPZoneGroupPutRequestBody := *openapiclient.NewDHCPZoneGroupPutRequestBody() // DHCPZoneGroupPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DHCPZoneResourcesAPI.PutZoneGroup(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).DHCPZoneGroupPutRequestBody(dHCPZoneGroupPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DHCPZoneResourcesAPI.PutZoneGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutZoneGroup`: GetZoneGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `DHCPZoneResourcesAPI.PutZoneGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **dHCPZoneGroupPutRequestBody** | [**DHCPZoneGroupPutRequestBody**](DHCPZoneGroupPutRequestBody.md) |  | 

### Return type

[**GetZoneGroups200ResponseDataInner**](GetZoneGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

