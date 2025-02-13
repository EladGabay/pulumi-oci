// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkloadbalancer

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Backend Sets in Oracle Cloud Infrastructure Network Load Balancer service.
//
// Lists all backend sets associated with a given network load balancer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/networkloadbalancer"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := networkloadbalancer.GetBackendSets(ctx, &networkloadbalancer.GetBackendSetsArgs{
// 			NetworkLoadBalancerId: oci_network_load_balancer_network_load_balancer.Test_network_load_balancer.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBackendSets(ctx *pulumi.Context, args *GetBackendSetsArgs, opts ...pulumi.InvokeOption) (*GetBackendSetsResult, error) {
	var rv GetBackendSetsResult
	err := ctx.Invoke("oci:networkloadbalancer/getBackendSets:getBackendSets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackendSets.
type GetBackendSetsArgs struct {
	Filters []GetBackendSetsFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId string `pulumi:"networkLoadBalancerId"`
}

// A collection of values returned by getBackendSets.
type GetBackendSetsResult struct {
	// The list of backend_set_collection.
	BackendSetCollections []GetBackendSetsBackendSetCollection `pulumi:"backendSetCollections"`
	Filters               []GetBackendSetsFilter               `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string `pulumi:"id"`
	NetworkLoadBalancerId string `pulumi:"networkLoadBalancerId"`
}
