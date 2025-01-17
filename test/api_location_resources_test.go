/*
BlueCat Address Manager 9.5 RESTful v2 API

Testing LocationResourcesAPIService

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

func Test_gobam2_LocationResourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LocationResourcesAPIService DeleteLocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.LocationResourcesAPI.DeleteLocation(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService DeleteLocationAnnotatedResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64
		var id int64

		httpRes, err := apiClient.LocationResourcesAPI.DeleteLocationAnnotatedResource(context.Background(), collectionId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService GetLocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.GetLocation(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService GetLocationAnnotatedResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64
		var id int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.GetLocationAnnotatedResource(context.Background(), collectionId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService GetLocationAnnotatedResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.GetLocationAnnotatedResources(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService GetLocationLocations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.GetLocationLocations(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService GetLocations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocationResourcesAPI.GetLocations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService PostLocationAnnotatedResource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.PostLocationAnnotatedResource(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService PostLocationLocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.PostLocationLocation(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationResourcesAPIService PutLocation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.LocationResourcesAPI.PutLocation(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
