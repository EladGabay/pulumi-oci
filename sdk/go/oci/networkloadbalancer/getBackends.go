// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkloadbalancer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Backends in Oracle Cloud Infrastructure Network Load Balancer service.
//
// Lists the backend servers for a given network load balancer and backend set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/networkloadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkloadbalancer.GetBackends(ctx, &networkloadbalancer.GetBackendsArgs{
// 			BackendSetName:        oci_network_load_balancer_backend_set.Test_backend_set.Name,
// 			NetworkLoadBalancerId: oci_network_load_balancer_network_load_balancer.Test_network_load_balancer.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBackends(ctx *pulumi.Context, args *GetBackendsArgs, opts ...pulumi.InvokeOption) (*GetBackendsResult, error) {
	var rv GetBackendsResult
	err := ctx.Invoke("oci:networkloadbalancer/getBackends:getBackends", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackends.
type GetBackendsArgs struct {
	// The name of the backend set associated with the backend servers.  Example: `exampleBackendSet`
	BackendSetName string              `pulumi:"backendSetName"`
	Filters        []GetBackendsFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId string `pulumi:"networkLoadBalancerId"`
}

// A collection of values returned by getBackends.
type GetBackendsResult struct {
	// The list of backend_collection.
	BackendCollections []GetBackendsBackendCollection `pulumi:"backendCollections"`
	BackendSetName     string                         `pulumi:"backendSetName"`
	Filters            []GetBackendsFilter            `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string `pulumi:"id"`
	NetworkLoadBalancerId string `pulumi:"networkLoadBalancerId"`
}