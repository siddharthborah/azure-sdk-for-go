package storage

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
	"net/http"
	"net/url"
)

// UsageOperations Client
type UsageOperationsClient struct {
	StorageManagementClient
}

func NewUsageOperationsClient(subscriptionId string) UsageOperationsClient {
	return NewUsageOperationsClientWithBaseUri(DefaultBaseUri, subscriptionId)
}

func NewUsageOperationsClientWithBaseUri(baseUri string, subscriptionId string) UsageOperationsClient {
	return UsageOperationsClient{NewWithBaseUri(baseUri, subscriptionId)}
}

// List gets the current usage count and the limit for the resources under the
// subscription.
func (client UsageOperationsClient) List() (result UsageListResult, ae autorest.Error) {
	req, err := client.NewListRequest()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storage.UsageOperationsClient", "List", "Failure creating request")
	}

	req, err = autorest.Prepare(
		req,
		client.WithAuthorization(),
		client.WithInspection())
	if err != nil {
		return result, autorest.NewErrorWithError(err, "storage.UsageOperationsClient", "List", "Failure preparing request")
	}

	resp, err := autorest.SendWithSender(client, req)

	if err == nil {
		err = autorest.Respond(
			resp,
			client.ByInspecting(),
			autorest.WithErrorUnlessOK(),
			autorest.ByUnmarshallingJSON(&result))
		if err != nil {
			ae = autorest.NewErrorWithError(err, "storage.UsageOperationsClient", "List", "Failure responding to request")
		}
	} else {
		ae = autorest.NewErrorWithError(err, "storage.UsageOperationsClient", "List", "Failure sending request")
	}

	autorest.Respond(resp,
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}

	return
}

// Create the List request.
func (client UsageOperationsClient) NewListRequest() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionId),
	}

	queryParameters := map[string]interface{}{
		"api-version": ApiVersion,
	}

	return autorest.DecoratePreparer(
		client.ListRequestPreparer(),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters)).Prepare(&http.Request{})
}

// Create a Preparer by which to prepare the List request.
func (client UsageOperationsClient) ListRequestPreparer() autorest.Preparer {
	return autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseUri),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/usages"))
}
