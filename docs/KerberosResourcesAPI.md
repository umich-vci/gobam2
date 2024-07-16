# \KerberosResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeyDistributionCenter**](KerberosResourcesAPI.md#DeleteKeyDistributionCenter) | **Delete** /api/v2/keyDistributionCenters/{id} | Delete a key distribution center
[**DeleteRealm**](KerberosResourcesAPI.md#DeleteRealm) | **Delete** /api/v2/realms/{id} | Delete a realm
[**DeleteServicePrincipal**](KerberosResourcesAPI.md#DeleteServicePrincipal) | **Delete** /api/v2/servicePrincipals/{id} | Delete a service principal
[**GetConfigurationRealms**](KerberosResourcesAPI.md#GetConfigurationRealms) | **Get** /api/v2/configurations/{collectionId}/realms | Retrieve a collection of realms
[**GetKeyDistributionCenter**](KerberosResourcesAPI.md#GetKeyDistributionCenter) | **Get** /api/v2/keyDistributionCenters/{id} | Retrieve a single key distribution center
[**GetKeyDistributionCenters**](KerberosResourcesAPI.md#GetKeyDistributionCenters) | **Get** /api/v2/keyDistributionCenters | Retrieve a collection of key distribution centers
[**GetRealm**](KerberosResourcesAPI.md#GetRealm) | **Get** /api/v2/realms/{id} | Retrieve a single realm
[**GetRealmKeyDistributionCenters**](KerberosResourcesAPI.md#GetRealmKeyDistributionCenters) | **Get** /api/v2/realms/{collectionId}/keyDistributionCenters | Retrieve a collection of key distribution centers
[**GetRealmServicePrincipals**](KerberosResourcesAPI.md#GetRealmServicePrincipals) | **Get** /api/v2/realms/{collectionId}/servicePrincipals | Retrieve a collection of service principals
[**GetRealms**](KerberosResourcesAPI.md#GetRealms) | **Get** /api/v2/realms | Retrieve a collection of realms
[**GetServicePrincipal**](KerberosResourcesAPI.md#GetServicePrincipal) | **Get** /api/v2/servicePrincipals/{id} | Retrieve a single service principal
[**GetServicePrincipals**](KerberosResourcesAPI.md#GetServicePrincipals) | **Get** /api/v2/servicePrincipals | Retrieve a collection of service principals
[**PatchServicePrincipal**](KerberosResourcesAPI.md#PatchServicePrincipal) | **Patch** /api/v2/servicePrincipals/{id} | Partially update a service principal
[**PostConfigurationRealm**](KerberosResourcesAPI.md#PostConfigurationRealm) | **Post** /api/v2/configurations/{collectionId}/realms | Create a new Kerberos realm
[**PostRealmKeyDistributionCenter**](KerberosResourcesAPI.md#PostRealmKeyDistributionCenter) | **Post** /api/v2/realms/{collectionId}/keyDistributionCenters | Create a new key distribution center
[**PostRealmServicePrincipal**](KerberosResourcesAPI.md#PostRealmServicePrincipal) | **Post** /api/v2/realms/{collectionId}/servicePrincipals | Create a new service principal
[**PutKeyDistributionCenter**](KerberosResourcesAPI.md#PutKeyDistributionCenter) | **Put** /api/v2/keyDistributionCenters/{id} | Update a key distribution center
[**PutRealm**](KerberosResourcesAPI.md#PutRealm) | **Put** /api/v2/realms/{id} | Update a realm
[**PutServicePrincipal**](KerberosResourcesAPI.md#PutServicePrincipal) | **Put** /api/v2/servicePrincipals/{id} | Update a service principal



## DeleteKeyDistributionCenter

> DeleteKeyDistributionCenter(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a key distribution center



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
	r, err := apiClient.KerberosResourcesAPI.DeleteKeyDistributionCenter(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.DeleteKeyDistributionCenter``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteKeyDistributionCenterRequest struct via the builder pattern


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


## DeleteRealm

> DeleteRealm(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a realm



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
	r, err := apiClient.KerberosResourcesAPI.DeleteRealm(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.DeleteRealm``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRealmRequest struct via the builder pattern


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


## DeleteServicePrincipal

> DeleteServicePrincipal(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a service principal



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
	r, err := apiClient.KerberosResourcesAPI.DeleteServicePrincipal(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.DeleteServicePrincipal``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteServicePrincipalRequest struct via the builder pattern


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


## GetConfigurationRealms

> GetRealms200Response GetConfigurationRealms(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of realms



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetConfigurationRealms(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetConfigurationRealms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationRealms`: GetRealms200Response
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetConfigurationRealms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRealmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetRealms200Response**](GetRealms200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyDistributionCenter

> GetKeyDistributionCenters200ResponseDataInner GetKeyDistributionCenter(ctx, id).Fields(fields).Execute()

Retrieve a single key distribution center



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetKeyDistributionCenter(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetKeyDistributionCenter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyDistributionCenter`: GetKeyDistributionCenters200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetKeyDistributionCenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyDistributionCenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetKeyDistributionCenters200ResponseDataInner**](GetKeyDistributionCenters200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyDistributionCenters

> GetKeyDistributionCenters200Response GetKeyDistributionCenters(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of key distribution centers



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetKeyDistributionCenters(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetKeyDistributionCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyDistributionCenters`: GetKeyDistributionCenters200Response
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetKeyDistributionCenters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyDistributionCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetKeyDistributionCenters200Response**](GetKeyDistributionCenters200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealm

> GetRealms200ResponseDataInner GetRealm(ctx, id).Fields(fields).Execute()

Retrieve a single realm



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetRealm(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealm`: GetRealms200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetRealms200ResponseDataInner**](GetRealms200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealmKeyDistributionCenters

> GetKeyDistributionCenters200Response GetRealmKeyDistributionCenters(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of key distribution centers



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetRealmKeyDistributionCenters(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetRealmKeyDistributionCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealmKeyDistributionCenters`: GetKeyDistributionCenters200Response
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetRealmKeyDistributionCenters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmKeyDistributionCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetKeyDistributionCenters200Response**](GetKeyDistributionCenters200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealmServicePrincipals

> GetServicePrincipals200Response GetRealmServicePrincipals(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of service principals



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetRealmServicePrincipals(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetRealmServicePrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealmServicePrincipals`: GetServicePrincipals200Response
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetRealmServicePrincipals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmServicePrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetServicePrincipals200Response**](GetServicePrincipals200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRealms

> GetRealms200Response GetRealms(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of realms



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetRealms(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetRealms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRealms`: GetRealms200Response
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetRealms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRealmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetRealms200Response**](GetRealms200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePrincipal

> GetServicePrincipals200ResponseDataInner GetServicePrincipal(ctx, id).Fields(fields).Execute()

Retrieve a single service principal



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetServicePrincipal(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetServicePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePrincipal`: GetServicePrincipals200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetServicePrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetServicePrincipals200ResponseDataInner**](GetServicePrincipals200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServicePrincipals

> GetServicePrincipals200Response GetServicePrincipals(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of service principals



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
	resp, r, err := apiClient.KerberosResourcesAPI.GetServicePrincipals(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.GetServicePrincipals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServicePrincipals`: GetServicePrincipals200Response
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.GetServicePrincipals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicePrincipalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetServicePrincipals200Response**](GetServicePrincipals200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServicePrincipal

> GetServicePrincipals200ResponseDataInner PatchServicePrincipal(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchServicePrincipalRequest(patchServicePrincipalRequest).Execute()

Partially update a service principal



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
	patchServicePrincipalRequest := *openapiclient.NewPatchServicePrincipalRequest() // PatchServicePrincipalRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PatchServicePrincipal(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchServicePrincipalRequest(patchServicePrincipalRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PatchServicePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchServicePrincipal`: GetServicePrincipals200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PatchServicePrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchServicePrincipalRequest** | [**PatchServicePrincipalRequest**](PatchServicePrincipalRequest.md) |  | 

### Return type

[**GetServicePrincipals200ResponseDataInner**](GetServicePrincipals200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConfigurationRealm

> GetRealms200ResponseDataInner PostConfigurationRealm(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).KerberosRealmPostRequestBody(kerberosRealmPostRequestBody).Execute()

Create a new Kerberos realm



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
	kerberosRealmPostRequestBody := *openapiclient.NewKerberosRealmPostRequestBody("example.company.net", "EXAMPLE.COMPANY.NET") // KerberosRealmPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PostConfigurationRealm(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).KerberosRealmPostRequestBody(kerberosRealmPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PostConfigurationRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConfigurationRealm`: GetRealms200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PostConfigurationRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConfigurationRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **kerberosRealmPostRequestBody** | [**KerberosRealmPostRequestBody**](KerberosRealmPostRequestBody.md) |  | 

### Return type

[**GetRealms200ResponseDataInner**](GetRealms200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRealmKeyDistributionCenter

> GetKeyDistributionCenters200ResponseDataInner PostRealmKeyDistributionCenter(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).KeyDistributionCenterPostRequestBody(keyDistributionCenterPostRequestBody).Execute()

Create a new key distribution center



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
	keyDistributionCenterPostRequestBody := *openapiclient.NewKeyDistributionCenterPostRequestBody("KDC-1", "192.168.0.10") // KeyDistributionCenterPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PostRealmKeyDistributionCenter(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).KeyDistributionCenterPostRequestBody(keyDistributionCenterPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PostRealmKeyDistributionCenter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRealmKeyDistributionCenter`: GetKeyDistributionCenters200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PostRealmKeyDistributionCenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRealmKeyDistributionCenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **keyDistributionCenterPostRequestBody** | [**KeyDistributionCenterPostRequestBody**](KeyDistributionCenterPostRequestBody.md) |  | 

### Return type

[**GetKeyDistributionCenters200ResponseDataInner**](GetKeyDistributionCenters200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRealmServicePrincipal

> GetServicePrincipals200ResponseDataInner PostRealmServicePrincipal(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).KerberosServicePrincipalPostRequestBody(kerberosServicePrincipalPostRequestBody).Execute()

Create a new service principal



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
	kerberosServicePrincipalPostRequestBody := *openapiclient.NewKerberosServicePrincipalPostRequestBody("DHCP/dhcp1.bcn.com", int32(123), "password") // KerberosServicePrincipalPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PostRealmServicePrincipal(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).KerberosServicePrincipalPostRequestBody(kerberosServicePrincipalPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PostRealmServicePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRealmServicePrincipal`: GetServicePrincipals200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PostRealmServicePrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRealmServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **kerberosServicePrincipalPostRequestBody** | [**KerberosServicePrincipalPostRequestBody**](KerberosServicePrincipalPostRequestBody.md) |  | 

### Return type

[**GetServicePrincipals200ResponseDataInner**](GetServicePrincipals200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutKeyDistributionCenter

> GetKeyDistributionCenters200ResponseDataInner PutKeyDistributionCenter(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).KeyDistributionCenterPutRequestBody(keyDistributionCenterPutRequestBody).Execute()

Update a key distribution center



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
	keyDistributionCenterPutRequestBody := *openapiclient.NewKeyDistributionCenterPutRequestBody("KDC-1", "192.168.0.10", int32(88)) // KeyDistributionCenterPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PutKeyDistributionCenter(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).KeyDistributionCenterPutRequestBody(keyDistributionCenterPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PutKeyDistributionCenter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutKeyDistributionCenter`: GetKeyDistributionCenters200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PutKeyDistributionCenter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutKeyDistributionCenterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **keyDistributionCenterPutRequestBody** | [**KeyDistributionCenterPutRequestBody**](KeyDistributionCenterPutRequestBody.md) |  | 

### Return type

[**GetKeyDistributionCenters200ResponseDataInner**](GetKeyDistributionCenters200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRealm

> GetRealms200ResponseDataInner PutRealm(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).KerberosRealmPutRequestBody(kerberosRealmPutRequestBody).Execute()

Update a realm



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
	kerberosRealmPutRequestBody := *openapiclient.NewKerberosRealmPutRequestBody("example.company.net", "EXAMPLE.COMPANY.NET") // KerberosRealmPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PutRealm(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).KerberosRealmPutRequestBody(kerberosRealmPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PutRealm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRealm`: GetRealms200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PutRealm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRealmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **kerberosRealmPutRequestBody** | [**KerberosRealmPutRequestBody**](KerberosRealmPutRequestBody.md) |  | 

### Return type

[**GetRealms200ResponseDataInner**](GetRealms200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServicePrincipal

> GetServicePrincipals200ResponseDataInner PutServicePrincipal(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).KerberosServicePrincipalPutRequestBody(kerberosServicePrincipalPutRequestBody).Execute()

Update a service principal



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
	kerberosServicePrincipalPutRequestBody := *openapiclient.NewKerberosServicePrincipalPutRequestBody("DHCP/dhcp1.bcn.com", int32(123)) // KerberosServicePrincipalPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KerberosResourcesAPI.PutServicePrincipal(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).KerberosServicePrincipalPutRequestBody(kerberosServicePrincipalPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KerberosResourcesAPI.PutServicePrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutServicePrincipal`: GetServicePrincipals200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `KerberosResourcesAPI.PutServicePrincipal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServicePrincipalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **kerberosServicePrincipalPutRequestBody** | [**KerberosServicePrincipalPutRequestBody**](KerberosServicePrincipalPutRequestBody.md) |  | 

### Return type

[**GetServicePrincipals200ResponseDataInner**](GetServicePrincipals200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

