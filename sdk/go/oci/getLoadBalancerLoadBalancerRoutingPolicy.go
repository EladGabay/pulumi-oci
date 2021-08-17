// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Load Balancer Routing Policy resource in Oracle Cloud Infrastructure Load Balancer service.
//
// Gets the specified routing policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.GetLoadBalancerLoadBalancerRoutingPolicy(ctx, &GetLoadBalancerLoadBalancerRoutingPolicyArgs{
// 			LoadBalancerId:    oci_load_balancer_load_balancer.Test_load_balancer.Id,
// 			RoutingPolicyName: oci_load_balancer_routing_policy.Test_routing_policy.Name,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupLoadBalancerLoadBalancerRoutingPolicy(ctx *pulumi.Context, args *LookupLoadBalancerLoadBalancerRoutingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupLoadBalancerLoadBalancerRoutingPolicyResult, error) {
	var rv LookupLoadBalancerLoadBalancerRoutingPolicyResult
	err := ctx.Invoke("oci:index/getLoadBalancerLoadBalancerRoutingPolicy:GetLoadBalancerLoadBalancerRoutingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetLoadBalancerLoadBalancerRoutingPolicy.
type LookupLoadBalancerLoadBalancerRoutingPolicyArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the specified load balancer.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The name of the routing policy to retrieve.  Example: `exampleRoutingPolicy`
	RoutingPolicyName string `pulumi:"routingPolicyName"`
}

// A collection of values returned by GetLoadBalancerLoadBalancerRoutingPolicy.
type LookupLoadBalancerLoadBalancerRoutingPolicyResult struct {
	// The version of the language in which `condition` of `rules` are composed.
	ConditionLanguageVersion string `pulumi:"conditionLanguageVersion"`
	Id                       string `pulumi:"id"`
	LoadBalancerId           string `pulumi:"loadBalancerId"`
	// A unique name for the routing policy rule. Avoid entering confidential information.
	Name              string `pulumi:"name"`
	RoutingPolicyName string `pulumi:"routingPolicyName"`
	// The ordered list of routing rules.
	Rules []GetLoadBalancerLoadBalancerRoutingPolicyRule `pulumi:"rules"`
	State string                                         `pulumi:"state"`
}