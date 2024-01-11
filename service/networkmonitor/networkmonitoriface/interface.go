// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package networkmonitoriface provides an interface to enable mocking the Amazon CloudWatch Network Monitor service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package networkmonitoriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/networkmonitor"
)

// NetworkMonitorAPI provides an interface to enable mocking the
// networkmonitor.NetworkMonitor service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon CloudWatch Network Monitor.
//	func myFunc(svc networkmonitoriface.NetworkMonitorAPI) bool {
//	    // Make svc.CreateMonitor request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := networkmonitor.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockNetworkMonitorClient struct {
//	    networkmonitoriface.NetworkMonitorAPI
//	}
//	func (m *mockNetworkMonitorClient) CreateMonitor(input *networkmonitor.CreateMonitorInput) (*networkmonitor.CreateMonitorOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockNetworkMonitorClient{}
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
type NetworkMonitorAPI interface {
	CreateMonitor(*networkmonitor.CreateMonitorInput) (*networkmonitor.CreateMonitorOutput, error)
	CreateMonitorWithContext(aws.Context, *networkmonitor.CreateMonitorInput, ...request.Option) (*networkmonitor.CreateMonitorOutput, error)
	CreateMonitorRequest(*networkmonitor.CreateMonitorInput) (*request.Request, *networkmonitor.CreateMonitorOutput)

	CreateProbe(*networkmonitor.CreateProbeInput) (*networkmonitor.CreateProbeOutput, error)
	CreateProbeWithContext(aws.Context, *networkmonitor.CreateProbeInput, ...request.Option) (*networkmonitor.CreateProbeOutput, error)
	CreateProbeRequest(*networkmonitor.CreateProbeInput) (*request.Request, *networkmonitor.CreateProbeOutput)

	DeleteMonitor(*networkmonitor.DeleteMonitorInput) (*networkmonitor.DeleteMonitorOutput, error)
	DeleteMonitorWithContext(aws.Context, *networkmonitor.DeleteMonitorInput, ...request.Option) (*networkmonitor.DeleteMonitorOutput, error)
	DeleteMonitorRequest(*networkmonitor.DeleteMonitorInput) (*request.Request, *networkmonitor.DeleteMonitorOutput)

	DeleteProbe(*networkmonitor.DeleteProbeInput) (*networkmonitor.DeleteProbeOutput, error)
	DeleteProbeWithContext(aws.Context, *networkmonitor.DeleteProbeInput, ...request.Option) (*networkmonitor.DeleteProbeOutput, error)
	DeleteProbeRequest(*networkmonitor.DeleteProbeInput) (*request.Request, *networkmonitor.DeleteProbeOutput)

	GetMonitor(*networkmonitor.GetMonitorInput) (*networkmonitor.GetMonitorOutput, error)
	GetMonitorWithContext(aws.Context, *networkmonitor.GetMonitorInput, ...request.Option) (*networkmonitor.GetMonitorOutput, error)
	GetMonitorRequest(*networkmonitor.GetMonitorInput) (*request.Request, *networkmonitor.GetMonitorOutput)

	GetProbe(*networkmonitor.GetProbeInput) (*networkmonitor.GetProbeOutput, error)
	GetProbeWithContext(aws.Context, *networkmonitor.GetProbeInput, ...request.Option) (*networkmonitor.GetProbeOutput, error)
	GetProbeRequest(*networkmonitor.GetProbeInput) (*request.Request, *networkmonitor.GetProbeOutput)

	ListMonitors(*networkmonitor.ListMonitorsInput) (*networkmonitor.ListMonitorsOutput, error)
	ListMonitorsWithContext(aws.Context, *networkmonitor.ListMonitorsInput, ...request.Option) (*networkmonitor.ListMonitorsOutput, error)
	ListMonitorsRequest(*networkmonitor.ListMonitorsInput) (*request.Request, *networkmonitor.ListMonitorsOutput)

	ListMonitorsPages(*networkmonitor.ListMonitorsInput, func(*networkmonitor.ListMonitorsOutput, bool) bool) error
	ListMonitorsPagesWithContext(aws.Context, *networkmonitor.ListMonitorsInput, func(*networkmonitor.ListMonitorsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*networkmonitor.ListTagsForResourceInput) (*networkmonitor.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *networkmonitor.ListTagsForResourceInput, ...request.Option) (*networkmonitor.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*networkmonitor.ListTagsForResourceInput) (*request.Request, *networkmonitor.ListTagsForResourceOutput)

	TagResource(*networkmonitor.TagResourceInput) (*networkmonitor.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *networkmonitor.TagResourceInput, ...request.Option) (*networkmonitor.TagResourceOutput, error)
	TagResourceRequest(*networkmonitor.TagResourceInput) (*request.Request, *networkmonitor.TagResourceOutput)

	UntagResource(*networkmonitor.UntagResourceInput) (*networkmonitor.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *networkmonitor.UntagResourceInput, ...request.Option) (*networkmonitor.UntagResourceOutput, error)
	UntagResourceRequest(*networkmonitor.UntagResourceInput) (*request.Request, *networkmonitor.UntagResourceOutput)

	UpdateMonitor(*networkmonitor.UpdateMonitorInput) (*networkmonitor.UpdateMonitorOutput, error)
	UpdateMonitorWithContext(aws.Context, *networkmonitor.UpdateMonitorInput, ...request.Option) (*networkmonitor.UpdateMonitorOutput, error)
	UpdateMonitorRequest(*networkmonitor.UpdateMonitorInput) (*request.Request, *networkmonitor.UpdateMonitorOutput)

	UpdateProbe(*networkmonitor.UpdateProbeInput) (*networkmonitor.UpdateProbeOutput, error)
	UpdateProbeWithContext(aws.Context, *networkmonitor.UpdateProbeInput, ...request.Option) (*networkmonitor.UpdateProbeOutput, error)
	UpdateProbeRequest(*networkmonitor.UpdateProbeInput) (*request.Request, *networkmonitor.UpdateProbeOutput)
}

var _ NetworkMonitorAPI = (*networkmonitor.NetworkMonitor)(nil)