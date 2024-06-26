//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnotificationhubs

// ClientCheckNotificationHubAvailabilityResponse contains the response from method Client.CheckNotificationHubAvailability.
type ClientCheckNotificationHubAvailabilityResponse struct {
	// Description of a CheckAvailability resource.
	CheckAvailabilityResult
}

// ClientCreateOrUpdateAuthorizationRuleResponse contains the response from method Client.CreateOrUpdateAuthorizationRule.
type ClientCreateOrUpdateAuthorizationRuleResponse struct {
	// Description of a Namespace AuthorizationRules.
	SharedAccessAuthorizationRuleResource
}

// ClientCreateOrUpdateResponse contains the response from method Client.CreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	// Description of a NotificationHub Resource.
	NotificationHubResource
}

// ClientDebugSendResponse contains the response from method Client.DebugSend.
type ClientDebugSendResponse struct {
	// Description of a NotificationHub Resource.
	DebugSendResponse
}

// ClientDeleteAuthorizationRuleResponse contains the response from method Client.DeleteAuthorizationRule.
type ClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// ClientDeleteResponse contains the response from method Client.Delete.
type ClientDeleteResponse struct {
	// placeholder for future response values
}

// ClientGetAuthorizationRuleResponse contains the response from method Client.GetAuthorizationRule.
type ClientGetAuthorizationRuleResponse struct {
	// Description of a Namespace AuthorizationRules.
	SharedAccessAuthorizationRuleResource
}

// ClientGetPnsCredentialsResponse contains the response from method Client.GetPnsCredentials.
type ClientGetPnsCredentialsResponse struct {
	// Description of a NotificationHub PNS Credentials.
	PnsCredentialsResource
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// Description of a NotificationHub Resource.
	NotificationHubResource
}

// ClientListAuthorizationRulesResponse contains the response from method Client.NewListAuthorizationRulesPager.
type ClientListAuthorizationRulesResponse struct {
	// The response of the List Namespace operation.
	SharedAccessAuthorizationRuleListResult
}

// ClientListKeysResponse contains the response from method Client.ListKeys.
type ClientListKeysResponse struct {
	// Namespace/NotificationHub Connection String
	ResourceListKeys
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// The response of the List NotificationHub operation.
	NotificationHubListResult
}

// ClientPatchResponse contains the response from method Client.Patch.
type ClientPatchResponse struct {
	// Description of a NotificationHub Resource.
	NotificationHubResource
}

// ClientRegenerateKeysResponse contains the response from method Client.RegenerateKeys.
type ClientRegenerateKeysResponse struct {
	// Namespace/NotificationHub Connection String
	ResourceListKeys
}

// NamespacesClientCheckAvailabilityResponse contains the response from method NamespacesClient.CheckAvailability.
type NamespacesClientCheckAvailabilityResponse struct {
	// Description of a CheckAvailability resource.
	CheckAvailabilityResult
}

// NamespacesClientCreateOrUpdateAuthorizationRuleResponse contains the response from method NamespacesClient.CreateOrUpdateAuthorizationRule.
type NamespacesClientCreateOrUpdateAuthorizationRuleResponse struct {
	// Description of a Namespace AuthorizationRules.
	SharedAccessAuthorizationRuleResource
}

// NamespacesClientCreateOrUpdateResponse contains the response from method NamespacesClient.CreateOrUpdate.
type NamespacesClientCreateOrUpdateResponse struct {
	// Description of a Namespace resource.
	NamespaceResource
}

// NamespacesClientDeleteAuthorizationRuleResponse contains the response from method NamespacesClient.DeleteAuthorizationRule.
type NamespacesClientDeleteAuthorizationRuleResponse struct {
	// placeholder for future response values
}

// NamespacesClientDeleteResponse contains the response from method NamespacesClient.BeginDelete.
type NamespacesClientDeleteResponse struct {
	// placeholder for future response values
}

// NamespacesClientGetAuthorizationRuleResponse contains the response from method NamespacesClient.GetAuthorizationRule.
type NamespacesClientGetAuthorizationRuleResponse struct {
	// Description of a Namespace AuthorizationRules.
	SharedAccessAuthorizationRuleResource
}

// NamespacesClientGetResponse contains the response from method NamespacesClient.Get.
type NamespacesClientGetResponse struct {
	// Description of a Namespace resource.
	NamespaceResource
}

// NamespacesClientListAllResponse contains the response from method NamespacesClient.NewListAllPager.
type NamespacesClientListAllResponse struct {
	// The response of the List Namespace operation.
	NamespaceListResult
}

// NamespacesClientListAuthorizationRulesResponse contains the response from method NamespacesClient.NewListAuthorizationRulesPager.
type NamespacesClientListAuthorizationRulesResponse struct {
	// The response of the List Namespace operation.
	SharedAccessAuthorizationRuleListResult
}

// NamespacesClientListKeysResponse contains the response from method NamespacesClient.ListKeys.
type NamespacesClientListKeysResponse struct {
	// Namespace/NotificationHub Connection String
	ResourceListKeys
}

// NamespacesClientListResponse contains the response from method NamespacesClient.NewListPager.
type NamespacesClientListResponse struct {
	// The response of the List Namespace operation.
	NamespaceListResult
}

// NamespacesClientPatchResponse contains the response from method NamespacesClient.Patch.
type NamespacesClientPatchResponse struct {
	// Description of a Namespace resource.
	NamespaceResource
}

// NamespacesClientRegenerateKeysResponse contains the response from method NamespacesClient.RegenerateKeys.
type NamespacesClientRegenerateKeysResponse struct {
	// Namespace/NotificationHub Connection String
	ResourceListKeys
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list NotificationHubs operations. It contains a list of operations and a URL link to get the next
	// set of results.
	OperationListResult
}
