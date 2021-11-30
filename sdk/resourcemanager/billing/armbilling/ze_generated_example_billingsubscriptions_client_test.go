//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByCustomer.json
func ExampleBillingSubscriptionsClient_ListByCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	pager := client.ListByCustomer("<billing-account-name>",
		"<customer-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BillingSubscription.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByBillingAccount.json
func ExampleBillingSubscriptionsClient_ListByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	pager := client.ListByBillingAccount("<billing-account-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BillingSubscription.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByBillingProfile.json
func ExampleBillingSubscriptionsClient_ListByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	pager := client.ListByBillingProfile("<billing-account-name>",
		"<billing-profile-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BillingSubscription.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByInvoiceSection.json
func ExampleBillingSubscriptionsClient_ListByInvoiceSection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	pager := client.ListByInvoiceSection("<billing-account-name>",
		"<billing-profile-name>",
		"<invoice-section-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("BillingSubscription.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscription.json
func ExampleBillingSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<billing-account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingSubscription.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingSubscription.json
func ExampleBillingSubscriptionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<billing-account-name>",
		armbilling.BillingSubscription{
			Properties: &armbilling.BillingSubscriptionProperties{
				CostCenter: to.StringPtr("<cost-center>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingSubscription.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MoveBillingSubscription.json
func ExampleBillingSubscriptionsClient_BeginMove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginMove(ctx,
		"<billing-account-name>",
		armbilling.TransferBillingSubscriptionRequestProperties{
			DestinationInvoiceSectionID: to.StringPtr("<destination-invoice-section-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingSubscription.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ValidateSubscriptionMoveFailure.json
func ExampleBillingSubscriptionsClient_ValidateMove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	_, err = client.ValidateMove(ctx,
		"<billing-account-name>",
		armbilling.TransferBillingSubscriptionRequestProperties{
			DestinationInvoiceSectionID: to.StringPtr("<destination-invoice-section-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}