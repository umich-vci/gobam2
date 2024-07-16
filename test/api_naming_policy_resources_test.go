/*
BlueCat Address Manager 9.5 RESTful v2 API

Testing NamingPolicyResourcesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gobam2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/umich-vci/gobam2"
)

func Test_gobam2_NamingPolicyResourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NamingPolicyResourcesAPIService DeleteCollectionNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collection string
		var collectionId int64
		var id int64

		httpRes, err := apiClient.NamingPolicyResourcesAPI.DeleteCollectionNamingPolicy(context.Background(), collection, collectionId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService DeleteNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.NamingPolicyResourcesAPI.DeleteNamingPolicy(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService DeleteNamingPolicyRestriction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.NamingPolicyResourcesAPI.DeleteNamingPolicyRestriction(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService DeleteNamingPolicyValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.NamingPolicyResourcesAPI.DeleteNamingPolicyValue(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetCollectionNamingPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collection string
		var collectionId int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetCollectionNamingPolicies(context.Background(), collection, collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetCollectionNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collection string
		var collectionId int64
		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetCollectionNamingPolicy(context.Background(), collection, collectionId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicy(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyPolicyRestriction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64
		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestriction(context.Background(), collectionId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyPolicyRestrictions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyRestrictions(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyPolicyValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64
		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyValue(context.Background(), collectionId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyPolicyValues", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyPolicyValues(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyRestriction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyRestriction(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyRestrictions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyRestrictions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyValue(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService GetNamingPolicyValues", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.GetNamingPolicyValues(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PostCollectionNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collection string
		var collectionId int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PostCollectionNamingPolicy(context.Background(), collection, collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PostNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PostNamingPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PostNamingPolicyRestriction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PostNamingPolicyRestriction(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PostNamingPolicyValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PostNamingPolicyValue(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PutNamingPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PutNamingPolicy(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PutNamingPolicyRestriction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PutNamingPolicyRestriction(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamingPolicyResourcesAPIService PutNamingPolicyValue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.NamingPolicyResourcesAPI.PutNamingPolicyValue(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}