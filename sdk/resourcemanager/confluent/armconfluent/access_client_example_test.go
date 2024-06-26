//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconfluent_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_UsersList.json
func ExampleAccessClient_ListUsers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListUsers(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListUsersSuccessResponse = armconfluent.AccessListUsersSuccessResponse{
	// 	Data: []*armconfluent.UserRecord{
	// 		{
	// 			AuthType: to.Ptr("AUTH_TYPE_SSO"),
	// 			Email: to.Ptr("marty.mcfly@example.com"),
	// 			FullName: to.Ptr("Marty McFly"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("User"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/user=u-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/iam/v2/users/u-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("UserList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/iam/v2/users"),
	// 		Last: to.Ptr("https://api.confluent.cloud/iam/v2/users?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/iam/v2/users?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/iam/v2/users?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_ServiceAccountsList.json
func ExampleAccessClient_ListServiceAccounts() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListServiceAccounts(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListServiceAccountsSuccessResponse = armconfluent.AccessListServiceAccountsSuccessResponse{
	// 	Data: []*armconfluent.ServiceAccountRecord{
	// 		{
	// 			Description: to.Ptr("Doc's repair bot for the DeLorean"),
	// 			DisplayName: to.Ptr("DeLorean_auto_repair"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("ServiceAccount"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/service-account=sa-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts/sa-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("ServiceAccountList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts"),
	// 		Last: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/iam/v2/service-accounts?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_InvitationsList.json
func ExampleAccessClient_ListInvitations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListInvitations(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
			"status":    to.Ptr("INVITE_STATUS_SENT"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListInvitationsSuccessResponse = armconfluent.AccessListInvitationsSuccessResponse{
	// 	Data: []*armconfluent.InvitationRecord{
	// 		{
	// 			AcceptedAt: to.Ptr("2022-07-06T17:21:33Z"),
	// 			AuthType: to.Ptr("AUTH_TYPE_SSO"),
	// 			Email: to.Ptr("johndoe@confluent.io"),
	// 			ExpiresAt: to.Ptr("2022-07-07T17:22:39Z"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("Invitation"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/invitation=i-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/iam/v2/invitations/i-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 			Status: to.Ptr("INVITE_STATUS_SENT"),
	// 	}},
	// 	Kind: to.Ptr("InvitationList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/iam/v2/invitations"),
	// 		Last: to.Ptr("https://api.confluent.cloud/iam/v2/invitations?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/iam/v2/invitations?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/iam/v2/invitations?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_InviteUser.json
func ExampleAccessClient_InviteUser() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().InviteUser(ctx, "myResourceGroup", "myOrganization", armconfluent.AccessInviteUserAccountModel{
		InvitedUserDetails: &armconfluent.AccessInvitedUserDetails{
			AuthType:     to.Ptr("AUTH_TYPE_SSO"),
			InvitedEmail: to.Ptr("user2@onmicrosoft.com"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InvitationRecord = armconfluent.InvitationRecord{
	// 	AcceptedAt: to.Ptr("2022-07-06T17:21:33Z"),
	// 	AuthType: to.Ptr("AUTH_TYPE_SSO"),
	// 	Email: to.Ptr("johndoe@confluent.io"),
	// 	ExpiresAt: to.Ptr("2022-07-07T17:22:39Z"),
	// 	ID: to.Ptr("dlz-f3a90de"),
	// 	Kind: to.Ptr("Invitation"),
	// 	Metadata: &armconfluent.MetadataEntity{
	// 		CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 		DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 		ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/invitation=i-12345"),
	// 		Self: to.Ptr("https://api.confluent.cloud/iam/v2/invitations/i-12345"),
	// 		UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 	},
	// 	Status: to.Ptr("INVITE_STATUS_SENT"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_EnvironmentList.json
func ExampleAccessClient_ListEnvironments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListEnvironments(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListEnvironmentsSuccessResponse = armconfluent.AccessListEnvironmentsSuccessResponse{
	// 	Data: []*armconfluent.EnvironmentRecord{
	// 		{
	// 			DisplayName: to.Ptr("prod-finance01"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("Environment"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=e-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/org/v2/environments/e-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("EnvironmentList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/org/v2/environments"),
	// 		Last: to.Ptr("https://api.confluent.cloud/org/v2/environments?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/org/v2/environments?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/org/v2/environments?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_ClusterList.json
func ExampleAccessClient_ListClusters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListClusters(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListClusterSuccessResponse = armconfluent.AccessListClusterSuccessResponse{
	// 	Data: []*armconfluent.ClusterRecord{
	// 		{
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("Cluster"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-abc123/cloud-cluster=lkc-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters/lkc-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 			Spec: &armconfluent.ClusterSpecEntity{
	// 				APIEndpoint: to.Ptr("https://pkac-00000.us-west-2.aws.confluent.cloud"),
	// 				Availability: to.Ptr("SINGLE_ZONE"),
	// 				Byok: &armconfluent.ClusterByokEntity{
	// 					ID: to.Ptr("cck-00000"),
	// 					Related: to.Ptr("https://api.confluent.cloud/byok/v1/keys/cck-00000"),
	// 					ResourceName: to.Ptr("https://api.confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/key=cck-00000"),
	// 				},
	// 				Cloud: to.Ptr("GCP"),
	// 				Config: &armconfluent.ClusterConfigEntity{
	// 					Kind: to.Ptr("Basic"),
	// 				},
	// 				DisplayName: to.Ptr("ProdKafkaCluster"),
	// 				Environment: &armconfluent.ClusterEnvironmentEntity{
	// 					Environment: to.Ptr("string"),
	// 					ID: to.Ptr("env-00000"),
	// 					Related: to.Ptr("https://api.confluent.cloud/v2/environments/env-00000"),
	// 					ResourceName: to.Ptr("https://api.confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-00000"),
	// 				},
	// 				HTTPEndpoint: to.Ptr("https://lkc-00000-00000.us-central1.gcp.glb.confluent.cloud"),
	// 				KafkaBootstrapEndpoint: to.Ptr("lkc-00000-00000.us-central1.gcp.glb.confluent.cloud:9092"),
	// 				Network: &armconfluent.ClusterNetworkEntity{
	// 					Environment: to.Ptr("string"),
	// 					ID: to.Ptr("n-00000"),
	// 					Related: to.Ptr("https://api.confluent.cloud/networking/v1/networks/n-00000"),
	// 					ResourceName: to.Ptr("https://api.confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/environment=env-abc123/network=n-00000"),
	// 				},
	// 				Region: to.Ptr("us-east4"),
	// 			},
	// 			Status: &armconfluent.ClusterStatusEntity{
	// 				Cku: to.Ptr[int32](2),
	// 				Phase: to.Ptr("PROVISIONED"),
	// 			},
	// 	}},
	// 	Kind: to.Ptr("ClusterList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters"),
	// 		Last: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/cmk/v2/clusters?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/db9788dde7a0c2c0d82e4fdf5f7b4de3843937e3/specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/Access_RoleBindingList.json
func ExampleAccessClient_ListRoleBindings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armconfluent.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccessClient().ListRoleBindings(ctx, "myResourceGroup", "myOrganization", armconfluent.ListAccessRequestModel{
		SearchFilters: map[string]*string{
			"pageSize":  to.Ptr("10"),
			"pageToken": to.Ptr("asc4fts4ft"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessListRoleBindingsSuccessResponse = armconfluent.AccessListRoleBindingsSuccessResponse{
	// 	Data: []*armconfluent.RoleBindingRecord{
	// 		{
	// 			CrnPattern: to.Ptr("crn://confluent.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-aaa1111/cloud-cluster=lkc-1111aaa"),
	// 			ID: to.Ptr("dlz-f3a90de"),
	// 			Kind: to.Ptr("RoleBinding"),
	// 			Metadata: &armconfluent.MetadataEntity{
	// 				CreatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				DeletedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 				ResourceName: to.Ptr("crn://confluent.cloud/organization=9bb441c4-edef-46ac-8a41-c49e44a3fd9a/role-binding=rb-12345"),
	// 				Self: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings/rb-12345"),
	// 				UpdatedAt: to.Ptr("2006-01-02T15:04:05-07:00"),
	// 			},
	// 			Principal: to.Ptr("User:u-111aaa"),
	// 			RoleName: to.Ptr("CloudClusterAdmin"),
	// 	}},
	// 	Kind: to.Ptr("RoleBindingList"),
	// 	Metadata: &armconfluent.ListMetadata{
	// 		First: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings"),
	// 		Last: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=bcAOehAY8F16YD84Z1wT"),
	// 		Next: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=UvmDWOB1iwfAIBPj6EYb"),
	// 		Prev: to.Ptr("https://api.confluent.cloud/iam/v2/role-bindings?page_token=YIXRY97wWYmwzrax4dld"),
	// 		TotalSize: to.Ptr[int32](123),
	// 	},
	// }
}
