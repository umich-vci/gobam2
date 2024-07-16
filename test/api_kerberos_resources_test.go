/*
BlueCat Address Manager 9.5 RESTful v2 API

Testing KerberosResourcesAPIService

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

func Test_gobam2_KerberosResourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KerberosResourcesAPIService DeleteKeyDistributionCenter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.KerberosResourcesAPI.DeleteKeyDistributionCenter(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService DeleteRealm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.KerberosResourcesAPI.DeleteRealm(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService DeleteServicePrincipal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		httpRes, err := apiClient.KerberosResourcesAPI.DeleteServicePrincipal(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetConfigurationRealms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetConfigurationRealms(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetKeyDistributionCenter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetKeyDistributionCenter(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetKeyDistributionCenters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetKeyDistributionCenters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetRealm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetRealm(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetRealmKeyDistributionCenters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetRealmKeyDistributionCenters(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetRealmServicePrincipals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetRealmServicePrincipals(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetRealms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetRealms(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetServicePrincipal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetServicePrincipal(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService GetServicePrincipals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KerberosResourcesAPI.GetServicePrincipals(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PatchServicePrincipal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PatchServicePrincipal(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PostConfigurationRealm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PostConfigurationRealm(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PostRealmKeyDistributionCenter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PostRealmKeyDistributionCenter(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PostRealmServicePrincipal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var collectionId int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PostRealmServicePrincipal(context.Background(), collectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PutKeyDistributionCenter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PutKeyDistributionCenter(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PutRealm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PutRealm(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosResourcesAPIService PutServicePrincipal", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id int64

		resp, httpRes, err := apiClient.KerberosResourcesAPI.PutServicePrincipal(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
