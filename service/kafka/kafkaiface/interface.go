// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kafkaiface provides an interface to enable mocking the Managed Streaming for Kafka service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kafkaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kafka"
)

// KafkaAPI provides an interface to enable mocking the
// kafka.Kafka service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Managed Streaming for Kafka.
//	func myFunc(svc kafkaiface.KafkaAPI) bool {
//	    // Make svc.BatchAssociateScramSecret request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := kafka.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockKafkaClient struct {
//	    kafkaiface.KafkaAPI
//	}
//	func (m *mockKafkaClient) BatchAssociateScramSecret(input *kafka.BatchAssociateScramSecretInput) (*kafka.BatchAssociateScramSecretOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockKafkaClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KafkaAPI interface {
	BatchAssociateScramSecret(*kafka.BatchAssociateScramSecretInput) (*kafka.BatchAssociateScramSecretOutput, error)
	BatchAssociateScramSecretWithContext(aws.Context, *kafka.BatchAssociateScramSecretInput, ...request.Option) (*kafka.BatchAssociateScramSecretOutput, error)
	BatchAssociateScramSecretRequest(*kafka.BatchAssociateScramSecretInput) (*request.Request, *kafka.BatchAssociateScramSecretOutput)

	BatchDisassociateScramSecret(*kafka.BatchDisassociateScramSecretInput) (*kafka.BatchDisassociateScramSecretOutput, error)
	BatchDisassociateScramSecretWithContext(aws.Context, *kafka.BatchDisassociateScramSecretInput, ...request.Option) (*kafka.BatchDisassociateScramSecretOutput, error)
	BatchDisassociateScramSecretRequest(*kafka.BatchDisassociateScramSecretInput) (*request.Request, *kafka.BatchDisassociateScramSecretOutput)

	CreateCluster(*kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *kafka.CreateClusterInput, ...request.Option) (*kafka.CreateClusterOutput, error)
	CreateClusterRequest(*kafka.CreateClusterInput) (*request.Request, *kafka.CreateClusterOutput)

	CreateClusterV2(*kafka.CreateClusterV2Input) (*kafka.CreateClusterV2Output, error)
	CreateClusterV2WithContext(aws.Context, *kafka.CreateClusterV2Input, ...request.Option) (*kafka.CreateClusterV2Output, error)
	CreateClusterV2Request(*kafka.CreateClusterV2Input) (*request.Request, *kafka.CreateClusterV2Output)

	CreateConfiguration(*kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error)
	CreateConfigurationWithContext(aws.Context, *kafka.CreateConfigurationInput, ...request.Option) (*kafka.CreateConfigurationOutput, error)
	CreateConfigurationRequest(*kafka.CreateConfigurationInput) (*request.Request, *kafka.CreateConfigurationOutput)

	CreateVpcConnection(*kafka.CreateVpcConnectionInput) (*kafka.CreateVpcConnectionOutput, error)
	CreateVpcConnectionWithContext(aws.Context, *kafka.CreateVpcConnectionInput, ...request.Option) (*kafka.CreateVpcConnectionOutput, error)
	CreateVpcConnectionRequest(*kafka.CreateVpcConnectionInput) (*request.Request, *kafka.CreateVpcConnectionOutput)

	DeleteCluster(*kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *kafka.DeleteClusterInput, ...request.Option) (*kafka.DeleteClusterOutput, error)
	DeleteClusterRequest(*kafka.DeleteClusterInput) (*request.Request, *kafka.DeleteClusterOutput)

	DeleteClusterPolicy(*kafka.DeleteClusterPolicyInput) (*kafka.DeleteClusterPolicyOutput, error)
	DeleteClusterPolicyWithContext(aws.Context, *kafka.DeleteClusterPolicyInput, ...request.Option) (*kafka.DeleteClusterPolicyOutput, error)
	DeleteClusterPolicyRequest(*kafka.DeleteClusterPolicyInput) (*request.Request, *kafka.DeleteClusterPolicyOutput)

	DeleteConfiguration(*kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error)
	DeleteConfigurationWithContext(aws.Context, *kafka.DeleteConfigurationInput, ...request.Option) (*kafka.DeleteConfigurationOutput, error)
	DeleteConfigurationRequest(*kafka.DeleteConfigurationInput) (*request.Request, *kafka.DeleteConfigurationOutput)

	DeleteVpcConnection(*kafka.DeleteVpcConnectionInput) (*kafka.DeleteVpcConnectionOutput, error)
	DeleteVpcConnectionWithContext(aws.Context, *kafka.DeleteVpcConnectionInput, ...request.Option) (*kafka.DeleteVpcConnectionOutput, error)
	DeleteVpcConnectionRequest(*kafka.DeleteVpcConnectionInput) (*request.Request, *kafka.DeleteVpcConnectionOutput)

	DescribeCluster(*kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error)
	DescribeClusterWithContext(aws.Context, *kafka.DescribeClusterInput, ...request.Option) (*kafka.DescribeClusterOutput, error)
	DescribeClusterRequest(*kafka.DescribeClusterInput) (*request.Request, *kafka.DescribeClusterOutput)

	DescribeClusterOperation(*kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterOperationWithContext(aws.Context, *kafka.DescribeClusterOperationInput, ...request.Option) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterOperationRequest(*kafka.DescribeClusterOperationInput) (*request.Request, *kafka.DescribeClusterOperationOutput)

	DescribeClusterOperationV2(*kafka.DescribeClusterOperationV2Input) (*kafka.DescribeClusterOperationV2Output, error)
	DescribeClusterOperationV2WithContext(aws.Context, *kafka.DescribeClusterOperationV2Input, ...request.Option) (*kafka.DescribeClusterOperationV2Output, error)
	DescribeClusterOperationV2Request(*kafka.DescribeClusterOperationV2Input) (*request.Request, *kafka.DescribeClusterOperationV2Output)

	DescribeClusterV2(*kafka.DescribeClusterV2Input) (*kafka.DescribeClusterV2Output, error)
	DescribeClusterV2WithContext(aws.Context, *kafka.DescribeClusterV2Input, ...request.Option) (*kafka.DescribeClusterV2Output, error)
	DescribeClusterV2Request(*kafka.DescribeClusterV2Input) (*request.Request, *kafka.DescribeClusterV2Output)

	DescribeConfiguration(*kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationWithContext(aws.Context, *kafka.DescribeConfigurationInput, ...request.Option) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationRequest(*kafka.DescribeConfigurationInput) (*request.Request, *kafka.DescribeConfigurationOutput)

	DescribeConfigurationRevision(*kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error)
	DescribeConfigurationRevisionWithContext(aws.Context, *kafka.DescribeConfigurationRevisionInput, ...request.Option) (*kafka.DescribeConfigurationRevisionOutput, error)
	DescribeConfigurationRevisionRequest(*kafka.DescribeConfigurationRevisionInput) (*request.Request, *kafka.DescribeConfigurationRevisionOutput)

	DescribeVpcConnection(*kafka.DescribeVpcConnectionInput) (*kafka.DescribeVpcConnectionOutput, error)
	DescribeVpcConnectionWithContext(aws.Context, *kafka.DescribeVpcConnectionInput, ...request.Option) (*kafka.DescribeVpcConnectionOutput, error)
	DescribeVpcConnectionRequest(*kafka.DescribeVpcConnectionInput) (*request.Request, *kafka.DescribeVpcConnectionOutput)

	GetBootstrapBrokers(*kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersWithContext(aws.Context, *kafka.GetBootstrapBrokersInput, ...request.Option) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersRequest(*kafka.GetBootstrapBrokersInput) (*request.Request, *kafka.GetBootstrapBrokersOutput)

	GetClusterPolicy(*kafka.GetClusterPolicyInput) (*kafka.GetClusterPolicyOutput, error)
	GetClusterPolicyWithContext(aws.Context, *kafka.GetClusterPolicyInput, ...request.Option) (*kafka.GetClusterPolicyOutput, error)
	GetClusterPolicyRequest(*kafka.GetClusterPolicyInput) (*request.Request, *kafka.GetClusterPolicyOutput)

	GetCompatibleKafkaVersions(*kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	GetCompatibleKafkaVersionsWithContext(aws.Context, *kafka.GetCompatibleKafkaVersionsInput, ...request.Option) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	GetCompatibleKafkaVersionsRequest(*kafka.GetCompatibleKafkaVersionsInput) (*request.Request, *kafka.GetCompatibleKafkaVersionsOutput)

	ListClientVpcConnections(*kafka.ListClientVpcConnectionsInput) (*kafka.ListClientVpcConnectionsOutput, error)
	ListClientVpcConnectionsWithContext(aws.Context, *kafka.ListClientVpcConnectionsInput, ...request.Option) (*kafka.ListClientVpcConnectionsOutput, error)
	ListClientVpcConnectionsRequest(*kafka.ListClientVpcConnectionsInput) (*request.Request, *kafka.ListClientVpcConnectionsOutput)

	ListClientVpcConnectionsPages(*kafka.ListClientVpcConnectionsInput, func(*kafka.ListClientVpcConnectionsOutput, bool) bool) error
	ListClientVpcConnectionsPagesWithContext(aws.Context, *kafka.ListClientVpcConnectionsInput, func(*kafka.ListClientVpcConnectionsOutput, bool) bool, ...request.Option) error

	ListClusterOperations(*kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error)
	ListClusterOperationsWithContext(aws.Context, *kafka.ListClusterOperationsInput, ...request.Option) (*kafka.ListClusterOperationsOutput, error)
	ListClusterOperationsRequest(*kafka.ListClusterOperationsInput) (*request.Request, *kafka.ListClusterOperationsOutput)

	ListClusterOperationsPages(*kafka.ListClusterOperationsInput, func(*kafka.ListClusterOperationsOutput, bool) bool) error
	ListClusterOperationsPagesWithContext(aws.Context, *kafka.ListClusterOperationsInput, func(*kafka.ListClusterOperationsOutput, bool) bool, ...request.Option) error

	ListClusterOperationsV2(*kafka.ListClusterOperationsV2Input) (*kafka.ListClusterOperationsV2Output, error)
	ListClusterOperationsV2WithContext(aws.Context, *kafka.ListClusterOperationsV2Input, ...request.Option) (*kafka.ListClusterOperationsV2Output, error)
	ListClusterOperationsV2Request(*kafka.ListClusterOperationsV2Input) (*request.Request, *kafka.ListClusterOperationsV2Output)

	ListClusterOperationsV2Pages(*kafka.ListClusterOperationsV2Input, func(*kafka.ListClusterOperationsV2Output, bool) bool) error
	ListClusterOperationsV2PagesWithContext(aws.Context, *kafka.ListClusterOperationsV2Input, func(*kafka.ListClusterOperationsV2Output, bool) bool, ...request.Option) error

	ListClusters(*kafka.ListClustersInput) (*kafka.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *kafka.ListClustersInput, ...request.Option) (*kafka.ListClustersOutput, error)
	ListClustersRequest(*kafka.ListClustersInput) (*request.Request, *kafka.ListClustersOutput)

	ListClustersPages(*kafka.ListClustersInput, func(*kafka.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *kafka.ListClustersInput, func(*kafka.ListClustersOutput, bool) bool, ...request.Option) error

	ListClustersV2(*kafka.ListClustersV2Input) (*kafka.ListClustersV2Output, error)
	ListClustersV2WithContext(aws.Context, *kafka.ListClustersV2Input, ...request.Option) (*kafka.ListClustersV2Output, error)
	ListClustersV2Request(*kafka.ListClustersV2Input) (*request.Request, *kafka.ListClustersV2Output)

	ListClustersV2Pages(*kafka.ListClustersV2Input, func(*kafka.ListClustersV2Output, bool) bool) error
	ListClustersV2PagesWithContext(aws.Context, *kafka.ListClustersV2Input, func(*kafka.ListClustersV2Output, bool) bool, ...request.Option) error

	ListConfigurationRevisions(*kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurationRevisionsWithContext(aws.Context, *kafka.ListConfigurationRevisionsInput, ...request.Option) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurationRevisionsRequest(*kafka.ListConfigurationRevisionsInput) (*request.Request, *kafka.ListConfigurationRevisionsOutput)

	ListConfigurationRevisionsPages(*kafka.ListConfigurationRevisionsInput, func(*kafka.ListConfigurationRevisionsOutput, bool) bool) error
	ListConfigurationRevisionsPagesWithContext(aws.Context, *kafka.ListConfigurationRevisionsInput, func(*kafka.ListConfigurationRevisionsOutput, bool) bool, ...request.Option) error

	ListConfigurations(*kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error)
	ListConfigurationsWithContext(aws.Context, *kafka.ListConfigurationsInput, ...request.Option) (*kafka.ListConfigurationsOutput, error)
	ListConfigurationsRequest(*kafka.ListConfigurationsInput) (*request.Request, *kafka.ListConfigurationsOutput)

	ListConfigurationsPages(*kafka.ListConfigurationsInput, func(*kafka.ListConfigurationsOutput, bool) bool) error
	ListConfigurationsPagesWithContext(aws.Context, *kafka.ListConfigurationsInput, func(*kafka.ListConfigurationsOutput, bool) bool, ...request.Option) error

	ListKafkaVersions(*kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error)
	ListKafkaVersionsWithContext(aws.Context, *kafka.ListKafkaVersionsInput, ...request.Option) (*kafka.ListKafkaVersionsOutput, error)
	ListKafkaVersionsRequest(*kafka.ListKafkaVersionsInput) (*request.Request, *kafka.ListKafkaVersionsOutput)

	ListKafkaVersionsPages(*kafka.ListKafkaVersionsInput, func(*kafka.ListKafkaVersionsOutput, bool) bool) error
	ListKafkaVersionsPagesWithContext(aws.Context, *kafka.ListKafkaVersionsInput, func(*kafka.ListKafkaVersionsOutput, bool) bool, ...request.Option) error

	ListNodes(*kafka.ListNodesInput) (*kafka.ListNodesOutput, error)
	ListNodesWithContext(aws.Context, *kafka.ListNodesInput, ...request.Option) (*kafka.ListNodesOutput, error)
	ListNodesRequest(*kafka.ListNodesInput) (*request.Request, *kafka.ListNodesOutput)

	ListNodesPages(*kafka.ListNodesInput, func(*kafka.ListNodesOutput, bool) bool) error
	ListNodesPagesWithContext(aws.Context, *kafka.ListNodesInput, func(*kafka.ListNodesOutput, bool) bool, ...request.Option) error

	ListScramSecrets(*kafka.ListScramSecretsInput) (*kafka.ListScramSecretsOutput, error)
	ListScramSecretsWithContext(aws.Context, *kafka.ListScramSecretsInput, ...request.Option) (*kafka.ListScramSecretsOutput, error)
	ListScramSecretsRequest(*kafka.ListScramSecretsInput) (*request.Request, *kafka.ListScramSecretsOutput)

	ListScramSecretsPages(*kafka.ListScramSecretsInput, func(*kafka.ListScramSecretsOutput, bool) bool) error
	ListScramSecretsPagesWithContext(aws.Context, *kafka.ListScramSecretsInput, func(*kafka.ListScramSecretsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *kafka.ListTagsForResourceInput, ...request.Option) (*kafka.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*kafka.ListTagsForResourceInput) (*request.Request, *kafka.ListTagsForResourceOutput)

	ListVpcConnections(*kafka.ListVpcConnectionsInput) (*kafka.ListVpcConnectionsOutput, error)
	ListVpcConnectionsWithContext(aws.Context, *kafka.ListVpcConnectionsInput, ...request.Option) (*kafka.ListVpcConnectionsOutput, error)
	ListVpcConnectionsRequest(*kafka.ListVpcConnectionsInput) (*request.Request, *kafka.ListVpcConnectionsOutput)

	ListVpcConnectionsPages(*kafka.ListVpcConnectionsInput, func(*kafka.ListVpcConnectionsOutput, bool) bool) error
	ListVpcConnectionsPagesWithContext(aws.Context, *kafka.ListVpcConnectionsInput, func(*kafka.ListVpcConnectionsOutput, bool) bool, ...request.Option) error

	PutClusterPolicy(*kafka.PutClusterPolicyInput) (*kafka.PutClusterPolicyOutput, error)
	PutClusterPolicyWithContext(aws.Context, *kafka.PutClusterPolicyInput, ...request.Option) (*kafka.PutClusterPolicyOutput, error)
	PutClusterPolicyRequest(*kafka.PutClusterPolicyInput) (*request.Request, *kafka.PutClusterPolicyOutput)

	RebootBroker(*kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error)
	RebootBrokerWithContext(aws.Context, *kafka.RebootBrokerInput, ...request.Option) (*kafka.RebootBrokerOutput, error)
	RebootBrokerRequest(*kafka.RebootBrokerInput) (*request.Request, *kafka.RebootBrokerOutput)

	RejectClientVpcConnection(*kafka.RejectClientVpcConnectionInput) (*kafka.RejectClientVpcConnectionOutput, error)
	RejectClientVpcConnectionWithContext(aws.Context, *kafka.RejectClientVpcConnectionInput, ...request.Option) (*kafka.RejectClientVpcConnectionOutput, error)
	RejectClientVpcConnectionRequest(*kafka.RejectClientVpcConnectionInput) (*request.Request, *kafka.RejectClientVpcConnectionOutput)

	TagResource(*kafka.TagResourceInput) (*kafka.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *kafka.TagResourceInput, ...request.Option) (*kafka.TagResourceOutput, error)
	TagResourceRequest(*kafka.TagResourceInput) (*request.Request, *kafka.TagResourceOutput)

	UntagResource(*kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *kafka.UntagResourceInput, ...request.Option) (*kafka.UntagResourceOutput, error)
	UntagResourceRequest(*kafka.UntagResourceInput) (*request.Request, *kafka.UntagResourceOutput)

	UpdateBrokerCount(*kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error)
	UpdateBrokerCountWithContext(aws.Context, *kafka.UpdateBrokerCountInput, ...request.Option) (*kafka.UpdateBrokerCountOutput, error)
	UpdateBrokerCountRequest(*kafka.UpdateBrokerCountInput) (*request.Request, *kafka.UpdateBrokerCountOutput)

	UpdateBrokerStorage(*kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error)
	UpdateBrokerStorageWithContext(aws.Context, *kafka.UpdateBrokerStorageInput, ...request.Option) (*kafka.UpdateBrokerStorageOutput, error)
	UpdateBrokerStorageRequest(*kafka.UpdateBrokerStorageInput) (*request.Request, *kafka.UpdateBrokerStorageOutput)

	UpdateBrokerType(*kafka.UpdateBrokerTypeInput) (*kafka.UpdateBrokerTypeOutput, error)
	UpdateBrokerTypeWithContext(aws.Context, *kafka.UpdateBrokerTypeInput, ...request.Option) (*kafka.UpdateBrokerTypeOutput, error)
	UpdateBrokerTypeRequest(*kafka.UpdateBrokerTypeInput) (*request.Request, *kafka.UpdateBrokerTypeOutput)

	UpdateClusterConfiguration(*kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error)
	UpdateClusterConfigurationWithContext(aws.Context, *kafka.UpdateClusterConfigurationInput, ...request.Option) (*kafka.UpdateClusterConfigurationOutput, error)
	UpdateClusterConfigurationRequest(*kafka.UpdateClusterConfigurationInput) (*request.Request, *kafka.UpdateClusterConfigurationOutput)

	UpdateClusterKafkaVersion(*kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error)
	UpdateClusterKafkaVersionWithContext(aws.Context, *kafka.UpdateClusterKafkaVersionInput, ...request.Option) (*kafka.UpdateClusterKafkaVersionOutput, error)
	UpdateClusterKafkaVersionRequest(*kafka.UpdateClusterKafkaVersionInput) (*request.Request, *kafka.UpdateClusterKafkaVersionOutput)

	UpdateConfiguration(*kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error)
	UpdateConfigurationWithContext(aws.Context, *kafka.UpdateConfigurationInput, ...request.Option) (*kafka.UpdateConfigurationOutput, error)
	UpdateConfigurationRequest(*kafka.UpdateConfigurationInput) (*request.Request, *kafka.UpdateConfigurationOutput)

	UpdateConnectivity(*kafka.UpdateConnectivityInput) (*kafka.UpdateConnectivityOutput, error)
	UpdateConnectivityWithContext(aws.Context, *kafka.UpdateConnectivityInput, ...request.Option) (*kafka.UpdateConnectivityOutput, error)
	UpdateConnectivityRequest(*kafka.UpdateConnectivityInput) (*request.Request, *kafka.UpdateConnectivityOutput)

	UpdateMonitoring(*kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error)
	UpdateMonitoringWithContext(aws.Context, *kafka.UpdateMonitoringInput, ...request.Option) (*kafka.UpdateMonitoringOutput, error)
	UpdateMonitoringRequest(*kafka.UpdateMonitoringInput) (*request.Request, *kafka.UpdateMonitoringOutput)

	UpdateSecurity(*kafka.UpdateSecurityInput) (*kafka.UpdateSecurityOutput, error)
	UpdateSecurityWithContext(aws.Context, *kafka.UpdateSecurityInput, ...request.Option) (*kafka.UpdateSecurityOutput, error)
	UpdateSecurityRequest(*kafka.UpdateSecurityInput) (*request.Request, *kafka.UpdateSecurityOutput)

	UpdateStorage(*kafka.UpdateStorageInput) (*kafka.UpdateStorageOutput, error)
	UpdateStorageWithContext(aws.Context, *kafka.UpdateStorageInput, ...request.Option) (*kafka.UpdateStorageOutput, error)
	UpdateStorageRequest(*kafka.UpdateStorageInput) (*request.Request, *kafka.UpdateStorageOutput)
}

var _ KafkaAPI = (*kafka.Kafka)(nil)
