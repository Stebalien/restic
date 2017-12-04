package advisor

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ConfigurationsClient is the REST APIs for Azure Advisor
type ConfigurationsClient struct {
	ManagementClient
}

// NewConfigurationsClient creates an instance of the ConfigurationsClient client.
func NewConfigurationsClient(subscriptionID string) ConfigurationsClient {
	return NewConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationsClientWithBaseURI creates an instance of the ConfigurationsClient client.
func NewConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationsClient {
	return ConfigurationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateInResourceGroup sends the create in resource group request.
//
// configContract is the Azure Advisor configuration data structure. resourceGroup is the name of the Azure resource
// group.
func (client ConfigurationsClient) CreateInResourceGroup(configContract ConfigData, resourceGroup string) (result ARMErrorResponseBody, err error) {
	req, err := client.CreateInResourceGroupPreparer(configContract, resourceGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateInResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateInResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInResourceGroup", resp, "Failure responding to request")
	}

	return
}

// CreateInResourceGroupPreparer prepares the CreateInResourceGroup request.
func (client ConfigurationsClient) CreateInResourceGroupPreparer(configContract ConfigData, resourceGroup string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroup":  autorest.Encode("path", resourceGroup),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", pathParameters),
		autorest.WithJSON(configContract),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateInResourceGroupSender sends the CreateInResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) CreateInResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateInResourceGroupResponder handles the response to the CreateInResourceGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) CreateInResourceGroupResponder(resp *http.Response) (result ARMErrorResponseBody, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateInSubscription create/Overwrite Azure Advisor configuration and also delete all configurations of contained
// resource groups.
//
// configContract is the Azure Advisor configuration data structure.
func (client ConfigurationsClient) CreateInSubscription(configContract ConfigData) (result ARMErrorResponseBody, err error) {
	req, err := client.CreateInSubscriptionPreparer(configContract)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInSubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateInSubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInSubscription", resp, "Failure sending request")
		return
	}

	result, err = client.CreateInSubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "CreateInSubscription", resp, "Failure responding to request")
	}

	return
}

// CreateInSubscriptionPreparer prepares the CreateInSubscription request.
func (client ConfigurationsClient) CreateInSubscriptionPreparer(configContract ConfigData) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", pathParameters),
		autorest.WithJSON(configContract),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateInSubscriptionSender sends the CreateInSubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) CreateInSubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateInSubscriptionResponder handles the response to the CreateInSubscription request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) CreateInSubscriptionResponder(resp *http.Response) (result ARMErrorResponseBody, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusBadRequest),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup sends the list by resource group request.
//
// resourceGroup is the name of the Azure resource group.
func (client ConfigurationsClient) ListByResourceGroup(resourceGroup string) (result ConfigurationListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ConfigurationsClient) ListByResourceGroupPreparer(resourceGroup string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroup":  autorest.Encode("path", resourceGroup),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Advisor/configurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) ListByResourceGroupResponder(resp *http.Response) (result ConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscription retrieve Azure Advisor configurations and also retrieve configurations of contained resource
// groups.
func (client ConfigurationsClient) ListBySubscription() (result ConfigurationListResult, err error) {
	req, err := client.ListBySubscriptionPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "advisor.ConfigurationsClient", "ListBySubscription", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client ConfigurationsClient) ListBySubscriptionPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-04-19"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/configurations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client ConfigurationsClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client ConfigurationsClient) ListBySubscriptionResponder(resp *http.Response) (result ConfigurationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}