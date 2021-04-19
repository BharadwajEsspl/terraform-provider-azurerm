package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// ElasticPoolOperationsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ElasticPoolOperationsClient struct {
	BaseClient
}

// NewElasticPoolOperationsClient creates an instance of the ElasticPoolOperationsClient client.
func NewElasticPoolOperationsClient(subscriptionID string) ElasticPoolOperationsClient {
	return NewElasticPoolOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewElasticPoolOperationsClientWithBaseURI creates an instance of the ElasticPoolOperationsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewElasticPoolOperationsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolOperationsClient {
	return ElasticPoolOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels the asynchronous operation on the elastic pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// operationID - the operation identifier.
func (client ElasticPoolOperationsClient) Cancel(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID uuid.UUID) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ElasticPoolOperationsClient.Cancel")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelPreparer(ctx, resourceGroupName, serverName, elasticPoolName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "Cancel", resp, "Failure responding to request")
		return
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client ElasticPoolOperationsClient) CancelPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, operationID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/operations/{operationId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolOperationsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client ElasticPoolOperationsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByElasticPool gets a list of operations performed on the elastic pool.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ElasticPoolOperationsClient) ListByElasticPool(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPoolOperationListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ElasticPoolOperationsClient.ListByElasticPool")
		defer func() {
			sc := -1
			if result.epolr.Response.Response != nil {
				sc = result.epolr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByElasticPoolNextResults
	req, err := client.ListByElasticPoolPreparer(ctx, resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "ListByElasticPool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByElasticPoolSender(req)
	if err != nil {
		result.epolr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "ListByElasticPool", resp, "Failure sending request")
		return
	}

	result.epolr, err = client.ListByElasticPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "ListByElasticPool", resp, "Failure responding to request")
		return
	}
	if result.epolr.hasNextLink() && result.epolr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByElasticPoolPreparer prepares the ListByElasticPool request.
func (client ElasticPoolOperationsClient) ListByElasticPoolPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-10-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByElasticPoolSender sends the ListByElasticPool request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolOperationsClient) ListByElasticPoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByElasticPoolResponder handles the response to the ListByElasticPool request. The method always
// closes the http.Response Body.
func (client ElasticPoolOperationsClient) ListByElasticPoolResponder(resp *http.Response) (result ElasticPoolOperationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByElasticPoolNextResults retrieves the next set of results, if any.
func (client ElasticPoolOperationsClient) listByElasticPoolNextResults(ctx context.Context, lastResults ElasticPoolOperationListResult) (result ElasticPoolOperationListResult, err error) {
	req, err := lastResults.elasticPoolOperationListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "listByElasticPoolNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByElasticPoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "listByElasticPoolNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByElasticPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolOperationsClient", "listByElasticPoolNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByElasticPoolComplete enumerates all values, automatically crossing page boundaries as required.
func (client ElasticPoolOperationsClient) ListByElasticPoolComplete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPoolOperationListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ElasticPoolOperationsClient.ListByElasticPool")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByElasticPool(ctx, resourceGroupName, serverName, elasticPoolName)
	return
}
