# \AdminResourcesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccessRight**](AdminResourcesAPI.md#DeleteAccessRight) | **Delete** /api/v2/accessRights/{id} | Delete an access right
[**DeleteAuthenticator**](AdminResourcesAPI.md#DeleteAuthenticator) | **Delete** /api/v2/authenticators/{id} | Delete an authenticator
[**DeleteCollectionUserDefinedLink**](AdminResourcesAPI.md#DeleteCollectionUserDefinedLink) | **Delete** /api/v2/{collection}/{collectionId}/userDefinedLinks/{id} | Delete a user-defined link
[**DeleteGroup**](AdminResourcesAPI.md#DeleteGroup) | **Delete** /api/v2/groups/{id} | Delete a group
[**DeleteGroupUser**](AdminResourcesAPI.md#DeleteGroupUser) | **Delete** /api/v2/groups/{collectionId}/users/{id} | Unlink a user from a group
[**DeleteImport**](AdminResourcesAPI.md#DeleteImport) | **Delete** /api/v2/imports/{id} | Delete an import resource
[**DeleteResponsePolicyImport**](AdminResourcesAPI.md#DeleteResponsePolicyImport) | **Delete** /api/v2/responsePolicies/{collectionId}/imports/{id} | Delete a response policy item import resource
[**DeleteSigningPolicy**](AdminResourcesAPI.md#DeleteSigningPolicy) | **Delete** /api/v2/signingPolicies/{id} | Delete a signing policy
[**DeleteTrustRelationship**](AdminResourcesAPI.md#DeleteTrustRelationship) | **Delete** /api/v2/trustRelationships/{id} | Delete a trust relationship
[**DeleteUser**](AdminResourcesAPI.md#DeleteUser) | **Delete** /api/v2/users/{id} | Delete a user
[**DeleteUserDefinedFieldDefinition**](AdminResourcesAPI.md#DeleteUserDefinedFieldDefinition) | **Delete** /api/v2/userDefinedFieldDefinitions/{id} | Delete a user-defined field definition
[**DeleteUserDefinedLinkDefinition**](AdminResourcesAPI.md#DeleteUserDefinedLinkDefinition) | **Delete** /api/v2/userDefinedLinkDefinitions/{id} | Delete a user-defined link definition
[**DeleteUserGroup**](AdminResourcesAPI.md#DeleteUserGroup) | **Delete** /api/v2/users/{collectionId}/groups/{id} | Unlink a group from a user
[**DeleteWorkflowRequest**](AdminResourcesAPI.md#DeleteWorkflowRequest) | **Delete** /api/v2/workflowRequests/{id} | Delete a workflow request
[**GetAccessRight**](AdminResourcesAPI.md#GetAccessRight) | **Get** /api/v2/accessRights/{id} | Retrieve a single access right
[**GetAccessRights**](AdminResourcesAPI.md#GetAccessRights) | **Get** /api/v2/accessRights | Retrieve a collection of access rights
[**GetAuthenticator**](AdminResourcesAPI.md#GetAuthenticator) | **Get** /api/v2/authenticators/{id} | Retrieve a single authenticator
[**GetAuthenticators**](AdminResourcesAPI.md#GetAuthenticators) | **Get** /api/v2/authenticators | Retrieve a collection of authenticators
[**GetCollectionAccessRights**](AdminResourcesAPI.md#GetCollectionAccessRights) | **Get** /api/v2/{collection}/{collectionId}/accessRights | Retrieve a collection of access rights
[**GetCollectionTransactions**](AdminResourcesAPI.md#GetCollectionTransactions) | **Get** /api/v2/{collection}/{collectionId}/transactions | Retrieve a collection of transactions
[**GetCollectionUserDefinedLink**](AdminResourcesAPI.md#GetCollectionUserDefinedLink) | **Get** /api/v2/{collection}/{collectionId}/userDefinedLinks/{id} | Retrieve a single user-defined link
[**GetCollectionUserDefinedLinks**](AdminResourcesAPI.md#GetCollectionUserDefinedLinks) | **Get** /api/v2/{collection}/{collectionId}/userDefinedLinks | Retrieve a collection of user-defined links
[**GetCollectionWorkflowRequests**](AdminResourcesAPI.md#GetCollectionWorkflowRequests) | **Get** /api/v2/{collection}/{collectionId}/workflowRequests | Retrieve a collection of workflow requests
[**GetDatabase**](AdminResourcesAPI.md#GetDatabase) | **Get** /api/v2/databases/{id} | Retrieve a single database resource
[**GetDatabases**](AdminResourcesAPI.md#GetDatabases) | **Get** /api/v2/databases | Retrieve a collection of database resources
[**GetEvent**](AdminResourcesAPI.md#GetEvent) | **Get** /api/v2/events/{id} | Retrieve a single event
[**GetEvents**](AdminResourcesAPI.md#GetEvents) | **Get** /api/v2/events | Retrieve a collection of events
[**GetGroup**](AdminResourcesAPI.md#GetGroup) | **Get** /api/v2/groups/{id} | Retrieve a single group
[**GetGroupUser**](AdminResourcesAPI.md#GetGroupUser) | **Get** /api/v2/groups/{collectionId}/users/{id} | Retrieve a single linked user
[**GetGroupUsers**](AdminResourcesAPI.md#GetGroupUsers) | **Get** /api/v2/groups/{collectionId}/users | Retrieve a collection of linked users
[**GetGroups**](AdminResourcesAPI.md#GetGroups) | **Get** /api/v2/groups | Retrieve a collection of groups
[**GetImport**](AdminResourcesAPI.md#GetImport) | **Get** /api/v2/imports/{id} | Retrieve a single resource import
[**GetImportLog**](AdminResourcesAPI.md#GetImportLog) | **Get** /api/v2/imports/{collectionId}/logs/{id} | Retrieve a log resource
[**GetImportLogs**](AdminResourcesAPI.md#GetImportLogs) | **Get** /api/v2/imports/{collectionId}/logs | Retrieve a collection of log resources
[**GetImports**](AdminResourcesAPI.md#GetImports) | **Get** /api/v2/imports | Retrieve a collection of resource imports
[**GetResources**](AdminResourcesAPI.md#GetResources) | **Get** /api/v2 | Retrieve a collection of resources
[**GetResponsePolicyImport**](AdminResourcesAPI.md#GetResponsePolicyImport) | **Get** /api/v2/responsePolicies/{collectionId}/imports/{id} | Retrieve a single response policy item import
[**GetResponsePolicyImports**](AdminResourcesAPI.md#GetResponsePolicyImports) | **Get** /api/v2/responsePolicies/{collectionId}/imports | Retrieve a collection of response policy item imports
[**GetRoot**](AdminResourcesAPI.md#GetRoot) | **Get** /api/v2/{id} | Retrieve the root resource
[**GetSession**](AdminResourcesAPI.md#GetSession) | **Get** /api/v2/sessions/{id} | Retrieve a single session
[**GetSessionCurrent**](AdminResourcesAPI.md#GetSessionCurrent) | **Get** /api/v2/sessions/current | Retrieve the current session
[**GetSessions**](AdminResourcesAPI.md#GetSessions) | **Get** /api/v2/sessions | Retrieve a collection of sessions
[**GetSetting**](AdminResourcesAPI.md#GetSetting) | **Get** /api/v2/settings/{id} | Retrieve a single settings resource
[**GetSettings**](AdminResourcesAPI.md#GetSettings) | **Get** /api/v2/settings | Retrieve a collection of settings resources
[**GetSigningPolicies**](AdminResourcesAPI.md#GetSigningPolicies) | **Get** /api/v2/signingPolicies | Retrieve a collection of signing policies
[**GetSigningPolicy**](AdminResourcesAPI.md#GetSigningPolicy) | **Get** /api/v2/signingPolicies/{id} | Retrieve a single signing policy
[**GetTransaction**](AdminResourcesAPI.md#GetTransaction) | **Get** /api/v2/transactions/{id} | Retrieve a single transaction
[**GetTransactionOperation**](AdminResourcesAPI.md#GetTransactionOperation) | **Get** /api/v2/transactions/{collectionId}/operations/{id} | Retrieve a single operation
[**GetTransactionOperations**](AdminResourcesAPI.md#GetTransactionOperations) | **Get** /api/v2/transactions/{collectionId}/operations | Retrieve a collection of operations
[**GetTransactions**](AdminResourcesAPI.md#GetTransactions) | **Get** /api/v2/transactions | Retrieve a collection of transactions
[**GetTrustRelationship**](AdminResourcesAPI.md#GetTrustRelationship) | **Get** /api/v2/trustRelationships/{id} | Retrieve a single trust relationship
[**GetTrustRelationships**](AdminResourcesAPI.md#GetTrustRelationships) | **Get** /api/v2/trustRelationships | Retrieve a collection of trust relationships
[**GetUser**](AdminResourcesAPI.md#GetUser) | **Get** /api/v2/users/{id} | Retrieve a single user
[**GetUserDefinedFieldDefinition**](AdminResourcesAPI.md#GetUserDefinedFieldDefinition) | **Get** /api/v2/userDefinedFieldDefinitions/{id} | Retrieve a single user-defined field defintion
[**GetUserDefinedFieldDefinitions**](AdminResourcesAPI.md#GetUserDefinedFieldDefinitions) | **Get** /api/v2/userDefinedFieldDefinitions | Retrieve a collection of user-defined field definitions
[**GetUserDefinedLinkDefinition**](AdminResourcesAPI.md#GetUserDefinedLinkDefinition) | **Get** /api/v2/userDefinedLinkDefinitions/{id} | Retrieve a single user-defined link definition
[**GetUserDefinedLinkDefinitions**](AdminResourcesAPI.md#GetUserDefinedLinkDefinitions) | **Get** /api/v2/userDefinedLinkDefinitions | Retrieve a collection of user-defined link definitions
[**GetUserGroup**](AdminResourcesAPI.md#GetUserGroup) | **Get** /api/v2/users/{collectionId}/groups/{id} | Retrieve a single linked group
[**GetUserGroups**](AdminResourcesAPI.md#GetUserGroups) | **Get** /api/v2/users/{collectionId}/groups | Retrieve a collection of linked groups
[**GetUserSession**](AdminResourcesAPI.md#GetUserSession) | **Get** /api/v2/users/{collectionId}/sessions/{id} | Retrieve a single session
[**GetUserSessions**](AdminResourcesAPI.md#GetUserSessions) | **Get** /api/v2/users/{collectionId}/sessions | Retrieve a collection of user sessions
[**GetUsers**](AdminResourcesAPI.md#GetUsers) | **Get** /api/v2/users | Retrieve a collection of users
[**GetWorkflowRequest**](AdminResourcesAPI.md#GetWorkflowRequest) | **Get** /api/v2/workflowRequests/{id} | Retrieve a single workflow request
[**GetWorkflowRequests**](AdminResourcesAPI.md#GetWorkflowRequests) | **Get** /api/v2/workflowRequests | Retrieve a collection of workflow requests
[**PatchDatabase**](AdminResourcesAPI.md#PatchDatabase) | **Patch** /api/v2/databases/{id} | Partially update an Address Manager database
[**PatchImport**](AdminResourcesAPI.md#PatchImport) | **Patch** /api/v2/imports/{id} | Partially update a import resource
[**PatchSession**](AdminResourcesAPI.md#PatchSession) | **Patch** /api/v2/sessions/{id} | Partially update a session
[**PatchSessionCurrent**](AdminResourcesAPI.md#PatchSessionCurrent) | **Patch** /api/v2/sessions/current | Partially update the current session
[**PatchTrustRelationship**](AdminResourcesAPI.md#PatchTrustRelationship) | **Patch** /api/v2/trustRelationships/{id} | Partially update a trust relationship
[**PatchUser**](AdminResourcesAPI.md#PatchUser) | **Patch** /api/v2/users/{id} | Partially update a user
[**PatchWorkflowRequest**](AdminResourcesAPI.md#PatchWorkflowRequest) | **Patch** /api/v2/workflowRequests/{id} | Partially update a workflow request
[**PostAccessRight**](AdminResourcesAPI.md#PostAccessRight) | **Post** /api/v2/accessRights | Create a new access right
[**PostAuthenticator**](AdminResourcesAPI.md#PostAuthenticator) | **Post** /api/v2/authenticators | Configure a new authenticator
[**PostCollectionUserDefinedLink**](AdminResourcesAPI.md#PostCollectionUserDefinedLink) | **Post** /api/v2/{collection}/{collectionId}/userDefinedLinks | Create a new user-defined link
[**PostGroup**](AdminResourcesAPI.md#PostGroup) | **Post** /api/v2/groups | Create a new group
[**PostGroupUser**](AdminResourcesAPI.md#PostGroupUser) | **Post** /api/v2/groups/{collectionId}/users | Link a user to a group
[**PostImport**](AdminResourcesAPI.md#PostImport) | **Post** /api/v2/imports | Create a new resource import
[**PostResponsePolicyImport**](AdminResourcesAPI.md#PostResponsePolicyImport) | **Post** /api/v2/responsePolicies/{collectionId}/imports | Create a new response policy item import
[**PostSession**](AdminResourcesAPI.md#PostSession) | **Post** /api/v2/sessions | Create a new session
[**PostSigningPolicy**](AdminResourcesAPI.md#PostSigningPolicy) | **Post** /api/v2/signingPolicies | Create a new signing policy
[**PostTransaction**](AdminResourcesAPI.md#PostTransaction) | **Post** /api/v2/transactions | Create a new transaction
[**PostTrustRelationship**](AdminResourcesAPI.md#PostTrustRelationship) | **Post** /api/v2/trustRelationships | Create a trust relationship
[**PostUser**](AdminResourcesAPI.md#PostUser) | **Post** /api/v2/users | Create a new user
[**PostUserDefinedFieldDefinition**](AdminResourcesAPI.md#PostUserDefinedFieldDefinition) | **Post** /api/v2/userDefinedFieldDefinitions | Create a new user-defined field definition
[**PostUserDefinedLinkDefinition**](AdminResourcesAPI.md#PostUserDefinedLinkDefinition) | **Post** /api/v2/userDefinedLinkDefinitions | Create a new user-defined link definition
[**PostUserGroup**](AdminResourcesAPI.md#PostUserGroup) | **Post** /api/v2/users/{collectionId}/groups | Link a group to a user
[**PutAccessRight**](AdminResourcesAPI.md#PutAccessRight) | **Put** /api/v2/accessRights/{id} | Update an access right
[**PutAuthenticator**](AdminResourcesAPI.md#PutAuthenticator) | **Put** /api/v2/authenticators/{id} | Update an authenticator
[**PutDatabase**](AdminResourcesAPI.md#PutDatabase) | **Put** /api/v2/databases/{id} | Updates a database resource
[**PutGroup**](AdminResourcesAPI.md#PutGroup) | **Put** /api/v2/groups/{id} | Update a group
[**PutSetting**](AdminResourcesAPI.md#PutSetting) | **Put** /api/v2/settings/{id} | Updates a settings resource
[**PutSigningPolicy**](AdminResourcesAPI.md#PutSigningPolicy) | **Put** /api/v2/signingPolicies/{id} | Update a signing policy
[**PutUser**](AdminResourcesAPI.md#PutUser) | **Put** /api/v2/users/{id} | Update a user
[**PutUserDefinedFieldDefinition**](AdminResourcesAPI.md#PutUserDefinedFieldDefinition) | **Put** /api/v2/userDefinedFieldDefinitions/{id} | Update a user-defined field definition
[**PutUserDefinedLinkDefinition**](AdminResourcesAPI.md#PutUserDefinedLinkDefinition) | **Put** /api/v2/userDefinedLinkDefinitions/{id} | Update a user-defined link definition



## DeleteAccessRight

> DeleteAccessRight(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete an access right



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
	r, err := apiClient.AdminResourcesAPI.DeleteAccessRight(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteAccessRight``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAccessRightRequest struct via the builder pattern


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


## DeleteAuthenticator

> DeleteAuthenticator(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete an authenticator



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
	r, err := apiClient.AdminResourcesAPI.DeleteAuthenticator(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteAuthenticator``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAuthenticatorRequest struct via the builder pattern


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


## DeleteCollectionUserDefinedLink

> DeleteCollectionUserDefinedLink(ctx, collection, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a user-defined link



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
	r, err := apiClient.AdminResourcesAPI.DeleteCollectionUserDefinedLink(context.Background(), collection, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteCollectionUserDefinedLink``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCollectionUserDefinedLinkRequest struct via the builder pattern


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


## DeleteGroup

> DeleteGroup(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a group



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
	r, err := apiClient.AdminResourcesAPI.DeleteGroup(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


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


## DeleteGroupUser

> DeleteGroupUser(ctx, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Unlink a user from a group



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
	r, err := apiClient.AdminResourcesAPI.DeleteGroupUser(context.Background(), collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteGroupUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGroupUserRequest struct via the builder pattern


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


## DeleteImport

> DeleteImport(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete an import resource



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
	r, err := apiClient.AdminResourcesAPI.DeleteImport(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteImport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteImportRequest struct via the builder pattern


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


## DeleteResponsePolicyImport

> DeleteResponsePolicyImport(ctx, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a response policy item import resource



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
	r, err := apiClient.AdminResourcesAPI.DeleteResponsePolicyImport(context.Background(), collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteResponsePolicyImport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteResponsePolicyImportRequest struct via the builder pattern


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


## DeleteSigningPolicy

> DeleteSigningPolicy(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a signing policy



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
	r, err := apiClient.AdminResourcesAPI.DeleteSigningPolicy(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteSigningPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSigningPolicyRequest struct via the builder pattern


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


## DeleteTrustRelationship

> DeleteTrustRelationship(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a trust relationship



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
	r, err := apiClient.AdminResourcesAPI.DeleteTrustRelationship(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteTrustRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTrustRelationshipRequest struct via the builder pattern


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


## DeleteUser

> DeleteUser(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a user



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
	r, err := apiClient.AdminResourcesAPI.DeleteUser(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteUser``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


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


## DeleteUserDefinedFieldDefinition

> DeleteUserDefinedFieldDefinition(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a user-defined field definition



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
	r, err := apiClient.AdminResourcesAPI.DeleteUserDefinedFieldDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteUserDefinedFieldDefinition``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserDefinedFieldDefinitionRequest struct via the builder pattern


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


## DeleteUserDefinedLinkDefinition

> DeleteUserDefinedLinkDefinition(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a user-defined link definition



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
	r, err := apiClient.AdminResourcesAPI.DeleteUserDefinedLinkDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteUserDefinedLinkDefinition``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserDefinedLinkDefinitionRequest struct via the builder pattern


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


## DeleteUserGroup

> DeleteUserGroup(ctx, collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Unlink a group from a user



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
	r, err := apiClient.AdminResourcesAPI.DeleteUserGroup(context.Background(), collectionId, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteUserGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUserGroupRequest struct via the builder pattern


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


## DeleteWorkflowRequest

> DeleteWorkflowRequest(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Delete a workflow request



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
	r, err := apiClient.AdminResourcesAPI.DeleteWorkflowRequest(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.DeleteWorkflowRequest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWorkflowRequestRequest struct via the builder pattern


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


## GetAccessRight

> GetAccessRights200ResponseDataInner GetAccessRight(ctx, id).Fields(fields).Execute()

Retrieve a single access right



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
	resp, r, err := apiClient.AdminResourcesAPI.GetAccessRight(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetAccessRight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRight`: GetAccessRights200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetAccessRight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetAccessRights200ResponseDataInner**](GetAccessRights200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRights

> GetAccessRights200Response GetAccessRights(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of access rights



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
	resp, r, err := apiClient.AdminResourcesAPI.GetAccessRights(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetAccessRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRights`: GetAccessRights200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetAccessRights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetAccessRights200Response**](GetAccessRights200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticator

> GetAuthenticators200ResponseDataInner GetAuthenticator(ctx, id).Fields(fields).Execute()

Retrieve a single authenticator



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
	resp, r, err := apiClient.AdminResourcesAPI.GetAuthenticator(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticator`: GetAuthenticators200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetAuthenticators200ResponseDataInner**](GetAuthenticators200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticators

> GetAuthenticators200Response GetAuthenticators(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of authenticators



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
	resp, r, err := apiClient.AdminResourcesAPI.GetAuthenticators(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetAuthenticators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthenticators`: GetAuthenticators200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetAuthenticators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetAuthenticators200Response**](GetAuthenticators200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionAccessRights

> GetAccessRights200Response GetCollectionAccessRights(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of access rights



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
	resp, r, err := apiClient.AdminResourcesAPI.GetCollectionAccessRights(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetCollectionAccessRights``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionAccessRights`: GetAccessRights200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetCollectionAccessRights`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionAccessRightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetAccessRights200Response**](GetAccessRights200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionTransactions

> GetTransactions200Response GetCollectionTransactions(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of transactions



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
	resp, r, err := apiClient.AdminResourcesAPI.GetCollectionTransactions(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetCollectionTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionTransactions`: GetTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetCollectionTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetTransactions200Response**](GetTransactions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionUserDefinedLink

> GetCollectionUserDefinedLinks200ResponseDataInner GetCollectionUserDefinedLink(ctx, collection, collectionId, id).Fields(fields).Execute()

Retrieve a single user-defined link



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
	resp, r, err := apiClient.AdminResourcesAPI.GetCollectionUserDefinedLink(context.Background(), collection, collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetCollectionUserDefinedLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionUserDefinedLink`: GetCollectionUserDefinedLinks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetCollectionUserDefinedLink`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetCollectionUserDefinedLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetCollectionUserDefinedLinks200ResponseDataInner**](GetCollectionUserDefinedLinks200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionUserDefinedLinks

> GetCollectionUserDefinedLinks200Response GetCollectionUserDefinedLinks(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of user-defined links



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
	resp, r, err := apiClient.AdminResourcesAPI.GetCollectionUserDefinedLinks(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetCollectionUserDefinedLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionUserDefinedLinks`: GetCollectionUserDefinedLinks200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetCollectionUserDefinedLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionUserDefinedLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetCollectionUserDefinedLinks200Response**](GetCollectionUserDefinedLinks200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionWorkflowRequests

> GetWorkflowRequests200Response GetCollectionWorkflowRequests(ctx, collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of workflow requests



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
	resp, r, err := apiClient.AdminResourcesAPI.GetCollectionWorkflowRequests(context.Background(), collection, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetCollectionWorkflowRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionWorkflowRequests`: GetWorkflowRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetCollectionWorkflowRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionWorkflowRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetWorkflowRequests200Response**](GetWorkflowRequests200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabase

> GetDatabases200ResponseDataInner GetDatabase(ctx, id).Fields(fields).Execute()

Retrieve a single database resource



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
	resp, r, err := apiClient.AdminResourcesAPI.GetDatabase(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabase`: GetDatabases200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetDatabases200ResponseDataInner**](GetDatabases200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabases

> GetDatabases200Response GetDatabases(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of database resources



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
	resp, r, err := apiClient.AdminResourcesAPI.GetDatabases(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetDatabases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabases`: GetDatabases200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetDatabases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetDatabases200Response**](GetDatabases200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> GetEvents200ResponseDataInner GetEvent(ctx, id).Fields(fields).Execute()

Retrieve a single event



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
	resp, r, err := apiClient.AdminResourcesAPI.GetEvent(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: GetEvents200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetEvents200ResponseDataInner**](GetEvents200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> GetEvents200Response GetEvents(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of events



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
	resp, r, err := apiClient.AdminResourcesAPI.GetEvents(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: GetEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetEvents200Response**](GetEvents200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> GetGroups200ResponseDataInner GetGroup(ctx, id).Fields(fields).Execute()

Retrieve a single group



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
	resp, r, err := apiClient.AdminResourcesAPI.GetGroup(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: GetGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetGroups200ResponseDataInner**](GetGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupUser

> GetUsers200ResponseDataInner GetGroupUser(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single linked user



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
	resp, r, err := apiClient.AdminResourcesAPI.GetGroupUser(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupUser`: GetUsers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetGroupUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetUsers200ResponseDataInner**](GetUsers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroupUsers

> GetUsers200Response GetGroupUsers(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of linked users



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
	resp, r, err := apiClient.AdminResourcesAPI.GetGroupUsers(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupUsers`: GetUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetUsers200Response**](GetUsers200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GetGroups200Response GetGroups(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of groups



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
	resp, r, err := apiClient.AdminResourcesAPI.GetGroups(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroups`: GetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetGroups200Response**](GetGroups200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImport

> GetImports200ResponseDataInner GetImport(ctx, id).Fields(fields).Execute()

Retrieve a single resource import



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
	resp, r, err := apiClient.AdminResourcesAPI.GetImport(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImport`: GetImports200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetImports200ResponseDataInner**](GetImports200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportLog

> GetImportLogs200ResponseDataInner GetImportLog(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a log resource



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
	resp, r, err := apiClient.AdminResourcesAPI.GetImportLog(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetImportLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportLog`: GetImportLogs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetImportLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetImportLogs200ResponseDataInner**](GetImportLogs200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportLogs

> GetImportLogs200Response GetImportLogs(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of log resources



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
	resp, r, err := apiClient.AdminResourcesAPI.GetImportLogs(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetImportLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImportLogs`: GetImportLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetImportLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetImportLogs200Response**](GetImportLogs200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImports

> GetImports200Response GetImports(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of resource imports



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
	resp, r, err := apiClient.AdminResourcesAPI.GetImports(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImports`: GetImports200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetImports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetImports200Response**](GetImports200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResources

> GetResources200Response GetResources(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of resources



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
	resp, r, err := apiClient.AdminResourcesAPI.GetResources(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResources`: GetResources200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetResources200Response**](GetResources200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponsePolicyImport

> GetResponsePolicyImports200ResponseDataInner GetResponsePolicyImport(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single response policy item import



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
	resp, r, err := apiClient.AdminResourcesAPI.GetResponsePolicyImport(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetResponsePolicyImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponsePolicyImport`: GetResponsePolicyImports200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetResponsePolicyImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponsePolicyImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetResponsePolicyImports200ResponseDataInner**](GetResponsePolicyImports200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResponsePolicyImports

> GetResponsePolicyImports200Response GetResponsePolicyImports(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of response policy item imports



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
	resp, r, err := apiClient.AdminResourcesAPI.GetResponsePolicyImports(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetResponsePolicyImports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResponsePolicyImports`: GetResponsePolicyImports200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetResponsePolicyImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResponsePolicyImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetResponsePolicyImports200Response**](GetResponsePolicyImports200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoot

> GetRoot200Response GetRoot(ctx, id).Fields(fields).Execute()

Retrieve the root resource



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
	resp, r, err := apiClient.AdminResourcesAPI.GetRoot(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetRoot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoot`: GetRoot200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetRoot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetRoot200Response**](GetRoot200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSession

> GetSessions200ResponseDataInner GetSession(ctx, id).Fields(fields).Execute()

Retrieve a single session



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSession(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSession`: GetSessions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetSessions200ResponseDataInner**](GetSessions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionCurrent

> GetSessions200Response GetSessionCurrent(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve the current session



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSessionCurrent(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSessionCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessionCurrent`: GetSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSessionCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetSessions200Response**](GetSessions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessions

> GetSessions200Response GetSessions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of sessions



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSessions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessions`: GetSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetSessions200Response**](GetSessions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetting

> GetSettings200ResponseDataInner GetSetting(ctx, id).Fields(fields).Execute()

Retrieve a single settings resource



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSetting(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetting`: GetSettings200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetSettings200ResponseDataInner**](GetSettings200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings

> GetSettings200Response GetSettings(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of settings resources



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSettings(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: GetSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetSettings200Response**](GetSettings200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSigningPolicies

> GetSigningPolicies200Response GetSigningPolicies(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of signing policies



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSigningPolicies(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSigningPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSigningPolicies`: GetSigningPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSigningPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetSigningPolicies200Response**](GetSigningPolicies200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSigningPolicy

> GetSigningPolicies200ResponseDataInner GetSigningPolicy(ctx, id).Fields(fields).Execute()

Retrieve a single signing policy



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
	resp, r, err := apiClient.AdminResourcesAPI.GetSigningPolicy(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetSigningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSigningPolicy`: GetSigningPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetSigningPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetSigningPolicies200ResponseDataInner**](GetSigningPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> GetTransactions200ResponseDataInner GetTransaction(ctx, id).Fields(fields).Execute()

Retrieve a single transaction



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
	resp, r, err := apiClient.AdminResourcesAPI.GetTransaction(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransaction`: GetTransactions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetTransactions200ResponseDataInner**](GetTransactions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionOperation

> GetTransactionOperations200ResponseDataInner GetTransactionOperation(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single operation



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
	resp, r, err := apiClient.AdminResourcesAPI.GetTransactionOperation(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetTransactionOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionOperation`: GetTransactionOperations200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetTransactionOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetTransactionOperations200ResponseDataInner**](GetTransactionOperations200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionOperations

> GetTransactionOperations200Response GetTransactionOperations(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of operations



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
	resp, r, err := apiClient.AdminResourcesAPI.GetTransactionOperations(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetTransactionOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactionOperations`: GetTransactionOperations200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetTransactionOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetTransactionOperations200Response**](GetTransactionOperations200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactions

> GetTransactions200Response GetTransactions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of transactions



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
	resp, r, err := apiClient.AdminResourcesAPI.GetTransactions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransactions`: GetTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetTransactions200Response**](GetTransactions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustRelationship

> GetTrustRelationships200ResponseDataInner GetTrustRelationship(ctx, id).Fields(fields).Execute()

Retrieve a single trust relationship



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
	resp, r, err := apiClient.AdminResourcesAPI.GetTrustRelationship(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustRelationship`: GetTrustRelationships200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetTrustRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetTrustRelationships200ResponseDataInner**](GetTrustRelationships200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrustRelationships

> GetTrustRelationships200Response GetTrustRelationships(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of trust relationships



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
	resp, r, err := apiClient.AdminResourcesAPI.GetTrustRelationships(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetTrustRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustRelationships`: GetTrustRelationships200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetTrustRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetTrustRelationships200Response**](GetTrustRelationships200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUsers200ResponseDataInner GetUser(ctx, id).Fields(fields).Execute()

Retrieve a single user



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUser(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUsers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetUsers200ResponseDataInner**](GetUsers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDefinedFieldDefinition

> GetUserDefinedFieldDefinitions200ResponseDataInner GetUserDefinedFieldDefinition(ctx, id).Fields(fields).Execute()

Retrieve a single user-defined field defintion



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserDefinedFieldDefinition(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserDefinedFieldDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDefinedFieldDefinition`: GetUserDefinedFieldDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserDefinedFieldDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDefinedFieldDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetUserDefinedFieldDefinitions200ResponseDataInner**](GetUserDefinedFieldDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDefinedFieldDefinitions

> GetUserDefinedFieldDefinitions200Response GetUserDefinedFieldDefinitions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of user-defined field definitions



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserDefinedFieldDefinitions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserDefinedFieldDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDefinedFieldDefinitions`: GetUserDefinedFieldDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserDefinedFieldDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDefinedFieldDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetUserDefinedFieldDefinitions200Response**](GetUserDefinedFieldDefinitions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDefinedLinkDefinition

> GetUserDefinedLinkDefinitions200ResponseDataInner GetUserDefinedLinkDefinition(ctx, id).Fields(fields).Execute()

Retrieve a single user-defined link definition



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserDefinedLinkDefinition(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserDefinedLinkDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDefinedLinkDefinition`: GetUserDefinedLinkDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserDefinedLinkDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDefinedLinkDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetUserDefinedLinkDefinitions200ResponseDataInner**](GetUserDefinedLinkDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDefinedLinkDefinitions

> GetUserDefinedLinkDefinitions200Response GetUserDefinedLinkDefinitions(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of user-defined link definitions



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserDefinedLinkDefinitions(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserDefinedLinkDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDefinedLinkDefinitions`: GetUserDefinedLinkDefinitions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserDefinedLinkDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDefinedLinkDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetUserDefinedLinkDefinitions200Response**](GetUserDefinedLinkDefinitions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroup

> GetGroups200ResponseDataInner GetUserGroup(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single linked group



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserGroup(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroup`: GetGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetGroups200ResponseDataInner**](GetGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserGroups

> GetGroups200Response GetUserGroups(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of linked groups



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserGroups(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserGroups`: GetGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetGroups200Response**](GetGroups200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSession

> GetSessions200ResponseDataInner GetUserSession(ctx, collectionId, id).Fields(fields).Execute()

Retrieve a single session



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserSession(context.Background(), collectionId, id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSession`: GetSessions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetSessions200ResponseDataInner**](GetSessions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSessions

> GetSessions200Response GetUserSessions(ctx, collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of user sessions



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUserSessions(context.Background(), collectionId).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUserSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSessions`: GetSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUserSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetSessions200Response**](GetSessions200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> GetUsers200Response GetUsers(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of users



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
	resp, r, err := apiClient.AdminResourcesAPI.GetUsers(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: GetUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetUsers200Response**](GetUsers200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowRequest

> GetWorkflowRequests200ResponseDataInner GetWorkflowRequest(ctx, id).Fields(fields).Execute()

Retrieve a single workflow request



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
	resp, r, err := apiClient.AdminResourcesAPI.GetWorkflowRequest(context.Background(), id).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetWorkflowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowRequest`: GetWorkflowRequests200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetWorkflowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | The subset of response resource field paths. | 

### Return type

[**GetWorkflowRequests200ResponseDataInner**](GetWorkflowRequests200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowRequests

> GetWorkflowRequests200Response GetWorkflowRequests(ctx).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()

Retrieve a collection of workflow requests



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
	resp, r, err := apiClient.AdminResourcesAPI.GetWorkflowRequests(context.Background()).Fields(fields).Limit(limit).Offset(offset).Filter(filter).OrderBy(orderBy).Total(total).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.GetWorkflowRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowRequests`: GetWorkflowRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.GetWorkflowRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | The subset of response resource field paths. | 
 **limit** | **int32** | The maximum number of resources returned in the response. | 
 **offset** | **int32** | The offset of the first resource returned in the response. | 
 **filter** | **string** | The query filter string. | 
 **orderBy** | **string** | The field path to order the response resources by. | 
 **total** | **bool** | Return the total number of resources matching the query. | 

### Return type

[**GetWorkflowRequests200Response**](GetWorkflowRequests200Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, text/csv, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDatabase

> GetDatabases200ResponseDataInner PatchDatabase(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchDatabaseRequest(patchDatabaseRequest).Execute()

Partially update an Address Manager database



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
	patchDatabaseRequest := *openapiclient.NewPatchDatabaseRequest() // PatchDatabaseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchDatabase(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchDatabaseRequest(patchDatabaseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDatabase`: GetDatabases200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchDatabaseRequest** | [**PatchDatabaseRequest**](PatchDatabaseRequest.md) |  | 

### Return type

[**GetDatabases200ResponseDataInner**](GetDatabases200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchImport

> GetImports200ResponseDataInner PatchImport(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchImportRequest(patchImportRequest).Execute()

Partially update a import resource



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
	patchImportRequest := *openapiclient.NewPatchImportRequest() // PatchImportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchImport(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchImportRequest(patchImportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchImport`: GetImports200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchImportRequest** | [**PatchImportRequest**](PatchImportRequest.md) |  | 

### Return type

[**GetImports200ResponseDataInner**](GetImports200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSession

> GetSessions200ResponseDataInner PatchSession(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchSessionCurrentRequest(patchSessionCurrentRequest).Execute()

Partially update a session



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
	patchSessionCurrentRequest := *openapiclient.NewPatchSessionCurrentRequest() // PatchSessionCurrentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchSession(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchSessionCurrentRequest(patchSessionCurrentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSession`: GetSessions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchSessionCurrentRequest** | [**PatchSessionCurrentRequest**](PatchSessionCurrentRequest.md) |  | 

### Return type

[**GetSessions200ResponseDataInner**](GetSessions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSessionCurrent

> GetSessions200ResponseDataInner PatchSessionCurrent(ctx).XBcnChangeControlComment(xBcnChangeControlComment).PatchSessionCurrentRequest(patchSessionCurrentRequest).Execute()

Partially update the current session



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
	patchSessionCurrentRequest := *openapiclient.NewPatchSessionCurrentRequest() // PatchSessionCurrentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchSessionCurrent(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).PatchSessionCurrentRequest(patchSessionCurrentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchSessionCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSessionCurrent`: GetSessions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchSessionCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchSessionCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchSessionCurrentRequest** | [**PatchSessionCurrentRequest**](PatchSessionCurrentRequest.md) |  | 

### Return type

[**GetSessions200ResponseDataInner**](GetSessions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchTrustRelationship

> GetTrustRelationships200ResponseDataInner PatchTrustRelationship(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchTrustRelationshipRequest(patchTrustRelationshipRequest).Execute()

Partially update a trust relationship



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
	patchTrustRelationshipRequest := *openapiclient.NewPatchTrustRelationshipRequest() // PatchTrustRelationshipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchTrustRelationship(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchTrustRelationshipRequest(patchTrustRelationshipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchTrustRelationship`: GetTrustRelationships200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchTrustRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchTrustRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchTrustRelationshipRequest** | [**PatchTrustRelationshipRequest**](PatchTrustRelationshipRequest.md) |  | 

### Return type

[**GetTrustRelationships200ResponseDataInner**](GetTrustRelationships200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUser

> GetUsers200ResponseDataInner PatchUser(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchUserRequest(patchUserRequest).Execute()

Partially update a user



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
	patchUserRequest := *openapiclient.NewPatchUserRequest() // PatchUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchUser(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchUserRequest(patchUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchUser`: GetUsers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchUserRequest** | [**PatchUserRequest**](PatchUserRequest.md) |  | 

### Return type

[**GetUsers200ResponseDataInner**](GetUsers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkflowRequest

> GetWorkflowRequests200ResponseDataInner PatchWorkflowRequest(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PatchWorkflowRequestRequest(patchWorkflowRequestRequest).Execute()

Partially update a workflow request



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
	patchWorkflowRequestRequest := *openapiclient.NewPatchWorkflowRequestRequest() // PatchWorkflowRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PatchWorkflowRequest(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PatchWorkflowRequestRequest(patchWorkflowRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PatchWorkflowRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkflowRequest`: GetWorkflowRequests200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PatchWorkflowRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkflowRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **patchWorkflowRequestRequest** | [**PatchWorkflowRequestRequest**](PatchWorkflowRequestRequest.md) |  | 

### Return type

[**GetWorkflowRequests200ResponseDataInner**](GetWorkflowRequests200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccessRight

> GetAccessRights200ResponseDataInner PostAccessRight(ctx).XBcnChangeControlComment(xBcnChangeControlComment).AccessRightPostRequestBody(accessRightPostRequestBody).Execute()

Create a new access right



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
	accessRightPostRequestBody := *openapiclient.NewAccessRightPostRequestBody("DefaultAccessLevel_example", *openapiclient.NewInlinedUserScope()) // AccessRightPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostAccessRight(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).AccessRightPostRequestBody(accessRightPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostAccessRight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccessRight`: GetAccessRights200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostAccessRight`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessRightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **accessRightPostRequestBody** | [**AccessRightPostRequestBody**](AccessRightPostRequestBody.md) |  | 

### Return type

[**GetAccessRights200ResponseDataInner**](GetAccessRights200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthenticator

> PostAuthenticator201Response PostAuthenticator(ctx).XBcnChangeControlComment(xBcnChangeControlComment).PostAuthenticatorRequest(postAuthenticatorRequest).Execute()

Configure a new authenticator



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
	postAuthenticatorRequest := openapiclient.postAuthenticator_request{KerberosAuthenticatorPostRequestBody: openapiclient.NewKerberosAuthenticatorPostRequestBody("Type_example", "LDAP authenticator", "https://hostname.com", "DOMAIN.COM")} // PostAuthenticatorRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostAuthenticator(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).PostAuthenticatorRequest(postAuthenticatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthenticator`: PostAuthenticator201Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostAuthenticator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **postAuthenticatorRequest** | [**PostAuthenticatorRequest**](PostAuthenticatorRequest.md) |  | 

### Return type

[**PostAuthenticator201Response**](PostAuthenticator201Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCollectionUserDefinedLink

> GetCollectionUserDefinedLinks200ResponseDataInner PostCollectionUserDefinedLink(ctx, collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedLinkedResourcePostRequestBody(userDefinedLinkedResourcePostRequestBody).Execute()

Create a new user-defined link



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
	userDefinedLinkedResourcePostRequestBody := *openapiclient.NewUserDefinedLinkedResourcePostRequestBody(int64(103307), *openapiclient.NewUserDefinedLinkedResourceLinkDefinition(int64(103307), "UDL IP4Addr to IP4Block")) // UserDefinedLinkedResourcePostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostCollectionUserDefinedLink(context.Background(), collection, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedLinkedResourcePostRequestBody(userDefinedLinkedResourcePostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostCollectionUserDefinedLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCollectionUserDefinedLink`: GetCollectionUserDefinedLinks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostCollectionUserDefinedLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** | The name of the collection. | 
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCollectionUserDefinedLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userDefinedLinkedResourcePostRequestBody** | [**UserDefinedLinkedResourcePostRequestBody**](UserDefinedLinkedResourcePostRequestBody.md) |  | 

### Return type

[**GetCollectionUserDefinedLinks200ResponseDataInner**](GetCollectionUserDefinedLinks200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGroup

> GetGroups200ResponseDataInner PostGroup(ctx).XBcnChangeControlComment(xBcnChangeControlComment).UserGroupPostRequestBody(userGroupPostRequestBody).Execute()

Create a new group



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
	userGroupPostRequestBody := *openapiclient.NewUserGroupPostRequestBody("Test group") // UserGroupPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostGroup(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).UserGroupPostRequestBody(userGroupPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGroup`: GetGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userGroupPostRequestBody** | [**UserGroupPostRequestBody**](UserGroupPostRequestBody.md) |  | 

### Return type

[**GetGroups200ResponseDataInner**](GetGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGroupUser

> GetUsers200ResponseDataInner PostGroupUser(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedUserPostRequestBody(linkedUserPostRequestBody).Execute()

Link a user to a group



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
	linkedUserPostRequestBody := *openapiclient.NewLinkedUserPostRequestBody(int64(103307)) // LinkedUserPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostGroupUser(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedUserPostRequestBody(linkedUserPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostGroupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGroupUser`: GetUsers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostGroupUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGroupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **linkedUserPostRequestBody** | [**LinkedUserPostRequestBody**](LinkedUserPostRequestBody.md) |  | 

### Return type

[**GetUsers200ResponseDataInner**](GetUsers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImport

> GetImports200ResponseDataInner PostImport(ctx).FileInfo(fileInfo).FileData(fileData).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Create a new resource import



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
	fileInfo := *openapiclient.NewResourceImport() // ResourceImport | 
	fileData := os.NewFile(1234, "some_file") // *os.File | The import file contents
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostImport(context.Background()).FileInfo(fileInfo).FileData(fileData).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostImport`: GetImports200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileInfo** | [**ResourceImport**](ResourceImport.md) |  | 
 **fileData** | ***os.File** | The import file contents | 
 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

[**GetImports200ResponseDataInner**](GetImports200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResponsePolicyImport

> GetResponsePolicyImports200ResponseDataInner PostResponsePolicyImport(ctx, collectionId).FileInfo(fileInfo).FileData(fileData).XBcnChangeControlComment(xBcnChangeControlComment).Execute()

Create a new response policy item import



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
	fileInfo := *openapiclient.NewResponsePolicyItemImport() // ResponsePolicyItemImport | 
	fileData := os.NewFile(1234, "some_file") // *os.File | 
	xBcnChangeControlComment := "xBcnChangeControlComment_example" // string | The change control comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostResponsePolicyImport(context.Background(), collectionId).FileInfo(fileInfo).FileData(fileData).XBcnChangeControlComment(xBcnChangeControlComment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostResponsePolicyImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResponsePolicyImport`: GetResponsePolicyImports200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostResponsePolicyImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResponsePolicyImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileInfo** | [**ResponsePolicyItemImport**](ResponsePolicyItemImport.md) |  | 
 **fileData** | ***os.File** |  | 
 **xBcnChangeControlComment** | **string** | The change control comment. | 

### Return type

[**GetResponsePolicyImports200ResponseDataInner**](GetResponsePolicyImports200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSession

> GetSessions200ResponseDataInner PostSession(ctx).XBcnChangeControlComment(xBcnChangeControlComment).UserSessionPostRequestBody(userSessionPostRequestBody).Execute()

Create a new session



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
	userSessionPostRequestBody := *openapiclient.NewUserSessionPostRequestBody("admin", "password") // UserSessionPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostSession(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).UserSessionPostRequestBody(userSessionPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSession`: GetSessions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userSessionPostRequestBody** | [**UserSessionPostRequestBody**](UserSessionPostRequestBody.md) |  | 

### Return type

[**GetSessions200ResponseDataInner**](GetSessions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSigningPolicy

> GetSigningPolicies200ResponseDataInner PostSigningPolicy(ctx).XBcnChangeControlComment(xBcnChangeControlComment).DNSSECSigningPolicyPostRequestBody(dNSSECSigningPolicyPostRequestBody).Execute()

Create a new signing policy



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
	dNSSECSigningPolicyPostRequestBody := *openapiclient.NewDNSSECSigningPolicyPostRequestBody("name", "KeyProvider_example", "SignatureDigestAlgorithm_example", "PT240H", "PT48H", "ZoneSigningKeyAlgorithm_example", int32(384), "PT720H", "PT168H", "ZoneSigningKeyRolloverMethod_example", "KeySigningKeyAlgorithm_example", int32(384), "PT8640H", "PT336H", "KeySigningKeyRolloverMethod_example") // DNSSECSigningPolicyPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostSigningPolicy(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).DNSSECSigningPolicyPostRequestBody(dNSSECSigningPolicyPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostSigningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSigningPolicy`: GetSigningPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostSigningPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSigningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **dNSSECSigningPolicyPostRequestBody** | [**DNSSECSigningPolicyPostRequestBody**](DNSSECSigningPolicyPostRequestBody.md) |  | 

### Return type

[**GetSigningPolicies200ResponseDataInner**](GetSigningPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTransaction

> GetTransactions200ResponseDataInner PostTransaction(ctx).XBcnChangeControlComment(xBcnChangeControlComment).TransactionPostRequestBody(transactionPostRequestBody).Execute()

Create a new transaction



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
	transactionPostRequestBody := *openapiclient.NewTransactionPostRequestBody("TransactionType_example", int64(3149)) // TransactionPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostTransaction(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).TransactionPostRequestBody(transactionPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTransaction`: GetTransactions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **transactionPostRequestBody** | [**TransactionPostRequestBody**](TransactionPostRequestBody.md) |  | 

### Return type

[**GetTransactions200ResponseDataInner**](GetTransactions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTrustRelationship

> GetTrustRelationships200ResponseDataInner PostTrustRelationship(ctx).XBcnChangeControlComment(xBcnChangeControlComment).TrustRelationshipPostRequestBody(trustRelationshipPostRequestBody).Execute()

Create a trust relationship



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
	trustRelationshipPostRequestBody := *openapiclient.NewTrustRelationshipPostRequestBody("192.168.1.1", "apiUser", "apiPassword") // TrustRelationshipPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostTrustRelationship(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).TrustRelationshipPostRequestBody(trustRelationshipPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTrustRelationship`: GetTrustRelationships200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostTrustRelationship`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTrustRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **trustRelationshipPostRequestBody** | [**TrustRelationshipPostRequestBody**](TrustRelationshipPostRequestBody.md) |  | 

### Return type

[**GetTrustRelationships200ResponseDataInner**](GetTrustRelationships200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUser

> GetUsers200ResponseDataInner PostUser(ctx).XBcnChangeControlComment(xBcnChangeControlComment).UserPostRequestBody(userPostRequestBody).Execute()

Create a new user



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
	userPostRequestBody := *openapiclient.NewUserPostRequestBody("admin", "user@example.com", "SecurityPrivilege_example", "AccessType_example") // UserPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostUser(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).UserPostRequestBody(userPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUser`: GetUsers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userPostRequestBody** | [**UserPostRequestBody**](UserPostRequestBody.md) |  | 

### Return type

[**GetUsers200ResponseDataInner**](GetUsers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserDefinedFieldDefinition

> GetUserDefinedFieldDefinitions200ResponseDataInner PostUserDefinedFieldDefinition(ctx).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedFieldDefinitionPostRequestBody(userDefinedFieldDefinitionPostRequestBody).Execute()

Create a new user-defined field definition



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
	userDefinedFieldDefinitionPostRequestBody := *openapiclient.NewUserDefinedFieldDefinitionPostRequestBody("BlueCatGateway", "BlueCat Gateway", "ResourceType_example", "DataType_example") // UserDefinedFieldDefinitionPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostUserDefinedFieldDefinition(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedFieldDefinitionPostRequestBody(userDefinedFieldDefinitionPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostUserDefinedFieldDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUserDefinedFieldDefinition`: GetUserDefinedFieldDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostUserDefinedFieldDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserDefinedFieldDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userDefinedFieldDefinitionPostRequestBody** | [**UserDefinedFieldDefinitionPostRequestBody**](UserDefinedFieldDefinitionPostRequestBody.md) |  | 

### Return type

[**GetUserDefinedFieldDefinitions200ResponseDataInner**](GetUserDefinedFieldDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserDefinedLinkDefinition

> GetUserDefinedLinkDefinitions200ResponseDataInner PostUserDefinedLinkDefinition(ctx).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedLinkDefinitionPostRequestBody(userDefinedLinkDefinitionPostRequestBody).Execute()

Create a new user-defined link definition



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
	userDefinedLinkDefinitionPostRequestBody := *openapiclient.NewUserDefinedLinkDefinitionPostRequestBody("UDL IP4Addr to IP4Block", "UDL IP4Addr to IP4Block", []string{"SourceTypes_example"}, []string{"DestinationTypes_example"}) // UserDefinedLinkDefinitionPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostUserDefinedLinkDefinition(context.Background()).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedLinkDefinitionPostRequestBody(userDefinedLinkDefinitionPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostUserDefinedLinkDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUserDefinedLinkDefinition`: GetUserDefinedLinkDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostUserDefinedLinkDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserDefinedLinkDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userDefinedLinkDefinitionPostRequestBody** | [**UserDefinedLinkDefinitionPostRequestBody**](UserDefinedLinkDefinitionPostRequestBody.md) |  | 

### Return type

[**GetUserDefinedLinkDefinitions200ResponseDataInner**](GetUserDefinedLinkDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserGroup

> GetGroups200ResponseDataInner PostUserGroup(ctx, collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedUserGroupPostRequestBody(linkedUserGroupPostRequestBody).Execute()

Link a group to a user



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
	linkedUserGroupPostRequestBody := *openapiclient.NewLinkedUserGroupPostRequestBody(int64(103307)) // LinkedUserGroupPostRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PostUserGroup(context.Background(), collectionId).XBcnChangeControlComment(xBcnChangeControlComment).LinkedUserGroupPostRequestBody(linkedUserGroupPostRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PostUserGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUserGroup`: GetGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PostUserGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUserGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **linkedUserGroupPostRequestBody** | [**LinkedUserGroupPostRequestBody**](LinkedUserGroupPostRequestBody.md) |  | 

### Return type

[**GetGroups200ResponseDataInner**](GetGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccessRight

> GetAccessRights200ResponseDataInner PutAccessRight(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).AccessRightPutRequestBody(accessRightPutRequestBody).Execute()

Update an access right



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
	accessRightPutRequestBody := *openapiclient.NewAccessRightPutRequestBody("DefaultAccessLevel_example", false, false, "WorkflowLevel_example", []openapiclient.AccessRightOverrideBean{*openapiclient.NewAccessRightOverrideBean()}, false) // AccessRightPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutAccessRight(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).AccessRightPutRequestBody(accessRightPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutAccessRight``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAccessRight`: GetAccessRights200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutAccessRight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccessRightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **accessRightPutRequestBody** | [**AccessRightPutRequestBody**](AccessRightPutRequestBody.md) |  | 

### Return type

[**GetAccessRights200ResponseDataInner**](GetAccessRights200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAuthenticator

> PostAuthenticator201Response PutAuthenticator(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PutAuthenticatorRequest(putAuthenticatorRequest).Execute()

Update an authenticator



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
	putAuthenticatorRequest := openapiclient.putAuthenticator_request{KerberosAuthenticatorPutRequestBody: openapiclient.NewKerberosAuthenticatorPutRequestBody("Type_example", "LDAP authenticator", "https://hostname.com", "DOMAIN.COM")} // PutAuthenticatorRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutAuthenticator(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutAuthenticatorRequest(putAuthenticatorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutAuthenticator``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAuthenticator`: PostAuthenticator201Response
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutAuthenticator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAuthenticatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **putAuthenticatorRequest** | [**PutAuthenticatorRequest**](PutAuthenticatorRequest.md) |  | 

### Return type

[**PostAuthenticator201Response**](PostAuthenticator201Response.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDatabase

> GetDatabases200ResponseDataInner PutDatabase(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).AddressManagerDatabasePutRequestBody(addressManagerDatabasePutRequestBody).Execute()

Updates a database resource



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
	addressManagerDatabasePutRequestBody := *openapiclient.NewAddressManagerDatabasePutRequestBody("Type_example", false) // AddressManagerDatabasePutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutDatabase(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).AddressManagerDatabasePutRequestBody(addressManagerDatabasePutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDatabase`: GetDatabases200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **addressManagerDatabasePutRequestBody** | [**AddressManagerDatabasePutRequestBody**](AddressManagerDatabasePutRequestBody.md) |  | 

### Return type

[**GetDatabases200ResponseDataInner**](GetDatabases200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutGroup

> GetGroups200ResponseDataInner PutGroup(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).UserGroupPutRequestBody(userGroupPutRequestBody).Execute()

Update a group



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
	userGroupPutRequestBody := *openapiclient.NewUserGroupPutRequestBody("Test group", "GroupType_example", false) // UserGroupPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutGroup(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).UserGroupPutRequestBody(userGroupPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutGroup`: GetGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userGroupPutRequestBody** | [**UserGroupPutRequestBody**](UserGroupPutRequestBody.md) |  | 

### Return type

[**GetGroups200ResponseDataInner**](GetGroups200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSetting

> GetSettings200ResponseDataInner PutSetting(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).PutSettingRequest(putSettingRequest).Execute()

Updates a settings resource



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
	putSettingRequest := openapiclient.putSetting_request{AuditDataSettingsPutRequestBody: openapiclient.NewAuditDataSettingsPutRequestBody("Type_example", false)} // PutSettingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutSetting(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).PutSettingRequest(putSettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSetting`: GetSettings200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **putSettingRequest** | [**PutSettingRequest**](PutSettingRequest.md) |  | 

### Return type

[**GetSettings200ResponseDataInner**](GetSettings200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSigningPolicy

> GetSigningPolicies200ResponseDataInner PutSigningPolicy(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).DNSSECSigningPolicyPutRequestBody(dNSSECSigningPolicyPutRequestBody).Execute()

Update a signing policy



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
	dNSSECSigningPolicyPutRequestBody := *openapiclient.NewDNSSECSigningPolicyPutRequestBody("name", "KeyProvider_example", "SignatureDigestAlgorithm_example", "PT240H", "PT48H", "ZoneSigningKeyAlgorithm_example", int32(384), "PT720H", "PT168H", "ZoneSigningKeyRolloverMethod_example", "KeySigningKeyAlgorithm_example", int32(384), "PT8640H", "PT336H", "KeySigningKeyRolloverMethod_example") // DNSSECSigningPolicyPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutSigningPolicy(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).DNSSECSigningPolicyPutRequestBody(dNSSECSigningPolicyPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutSigningPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSigningPolicy`: GetSigningPolicies200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutSigningPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSigningPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **dNSSECSigningPolicyPutRequestBody** | [**DNSSECSigningPolicyPutRequestBody**](DNSSECSigningPolicyPutRequestBody.md) |  | 

### Return type

[**GetSigningPolicies200ResponseDataInner**](GetSigningPolicies200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUser

> GetUsers200ResponseDataInner PutUser(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).UserPutRequestBody(userPutRequestBody).Execute()

Update a user



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
	userPutRequestBody := *openapiclient.NewUserPutRequestBody(*openapiclient.NewInlinedAuthenticator(), "user@example.com", "SecurityPrivilege_example", "HistoryPrivilege_example", "AccessType_example", false, false, false) // UserPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutUser(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).UserPutRequestBody(userPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUser`: GetUsers200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userPutRequestBody** | [**UserPutRequestBody**](UserPutRequestBody.md) |  | 

### Return type

[**GetUsers200ResponseDataInner**](GetUsers200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUserDefinedFieldDefinition

> GetUserDefinedFieldDefinitions200ResponseDataInner PutUserDefinedFieldDefinition(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedFieldDefinitionPutRequestBody(userDefinedFieldDefinitionPutRequestBody).Execute()

Update a user-defined field definition



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
	userDefinedFieldDefinitionPutRequestBody := *openapiclient.NewUserDefinedFieldDefinitionPutRequestBody("BlueCat Gateway", "ResourceType_example") // UserDefinedFieldDefinitionPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutUserDefinedFieldDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedFieldDefinitionPutRequestBody(userDefinedFieldDefinitionPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutUserDefinedFieldDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUserDefinedFieldDefinition`: GetUserDefinedFieldDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutUserDefinedFieldDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserDefinedFieldDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userDefinedFieldDefinitionPutRequestBody** | [**UserDefinedFieldDefinitionPutRequestBody**](UserDefinedFieldDefinitionPutRequestBody.md) |  | 

### Return type

[**GetUserDefinedFieldDefinitions200ResponseDataInner**](GetUserDefinedFieldDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUserDefinedLinkDefinition

> GetUserDefinedLinkDefinitions200ResponseDataInner PutUserDefinedLinkDefinition(ctx, id).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedLinkDefinitionPutRequestBody(userDefinedLinkDefinitionPutRequestBody).Execute()

Update a user-defined link definition



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
	userDefinedLinkDefinitionPutRequestBody := *openapiclient.NewUserDefinedLinkDefinitionPutRequestBody("UDL IP4Addr to IP4Block", []string{"SourceTypes_example"}, []string{"DestinationTypes_example"}) // UserDefinedLinkDefinitionPutRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminResourcesAPI.PutUserDefinedLinkDefinition(context.Background(), id).XBcnChangeControlComment(xBcnChangeControlComment).UserDefinedLinkDefinitionPutRequestBody(userDefinedLinkDefinitionPutRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminResourcesAPI.PutUserDefinedLinkDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutUserDefinedLinkDefinition`: GetUserDefinedLinkDefinitions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `AdminResourcesAPI.PutUserDefinedLinkDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserDefinedLinkDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xBcnChangeControlComment** | **string** | The change control comment. | 
 **userDefinedLinkDefinitionPutRequestBody** | [**UserDefinedLinkDefinitionPutRequestBody**](UserDefinedLinkDefinitionPutRequestBody.md) |  | 

### Return type

[**GetUserDefinedLinkDefinitions200ResponseDataInner**](GetUserDefinedLinkDefinitions200ResponseDataInner.md)

### Authorization

[bearerToken](../README.md#bearerToken), [basicAuthentication](../README.md#basicAuthentication)

### HTTP request headers

- **Content-Type**: application/hal+json, application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

