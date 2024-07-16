/*
BlueCat Address Manager 9.5 RESTful v2 API

Testing MACAddressAndClientIdentifierResourcesAPIService

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

func Test_gobam2_MACAddressAndClientIdentifierResourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService DeleteMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.DeleteMacAddress(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService DeleteMacPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.DeleteMacPool(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService DeleteMacPoolMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64
		var id int64

		httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.DeleteMacPoolMacAddress(context.Background(), collectionId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetClientIdentifier", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifier(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetClientIdentifiers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetClientIdentifiers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetConfigurationClientIdentifiers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetConfigurationClientIdentifiers(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetConfigurationMacAddresses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacAddresses(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetConfigurationMacPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetConfigurationMacPools(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacAddress(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetMacAddresses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacAddresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetMacPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPool(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetMacPoolMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64
		var id int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddress(context.Background(), collectionId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetMacPoolMacAddresses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPoolMacAddresses(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService GetMacPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.GetMacPools(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService PostConfigurationMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacAddress(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService PostConfigurationMacPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PostConfigurationMacPool(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService PostMacPoolMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PostMacPoolMacAddress(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService PutMacAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PutMacAddress(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MACAddressAndClientIdentifierResourcesAPIService PutMacPool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.MACAddressAndClientIdentifierResourcesAPI.PutMacPool(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
