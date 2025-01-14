/*
BlueCat Address Manager 9.5 RESTful v2 API

Testing DHCPClientClassResourcesAPIService

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

func Test_gobam2_DHCPClientClassResourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DHCPClientClassResourcesAPIService DeleteClientClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.DHCPClientClassResourcesAPI.DeleteClientClass(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService DeleteSubclass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.DHCPClientClassResourcesAPI.DeleteSubclass(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService GetClientClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.GetClientClass(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService GetClientClassSubclasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.GetClientClassSubclasses(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService GetClientClasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.GetClientClasses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService GetConfigurationClientClasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.GetConfigurationClientClasses(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService GetSubclass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.GetSubclass(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService GetSubclasses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.GetSubclasses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService PostClientClassSubclass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.PostClientClassSubclass(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService PostConfigurationClientClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.PostConfigurationClientClass(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService PutClientClass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.PutClientClass(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DHCPClientClassResourcesAPIService PutSubclass", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.DHCPClientClassResourcesAPI.PutSubclass(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
