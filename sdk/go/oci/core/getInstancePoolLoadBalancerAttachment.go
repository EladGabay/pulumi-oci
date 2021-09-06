// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Instance Pool Load Balancer Attachment resource in Oracle Cloud Infrastructure Core service.
//
// Gets information about a load balancer that is attached to the specified instance pool.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.GetInstancePoolLoadBalancerAttachment(ctx, &core.GetInstancePoolLoadBalancerAttachmentArgs{
// 			InstancePoolId:                       oci_core_instance_pool.Test_instance_pool.Id,
// 			InstancePoolLoadBalancerAttachmentId: oci_core_instance_pool_load_balancer_attachment.Test_instance_pool_load_balancer_attachment.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstancePoolLoadBalancerAttachment(ctx *pulumi.Context, args *GetInstancePoolLoadBalancerAttachmentArgs, opts ...pulumi.InvokeOption) (*GetInstancePoolLoadBalancerAttachmentResult, error) {
	var rv GetInstancePoolLoadBalancerAttachmentResult
	err := ctx.Invoke("oci:core/getInstancePoolLoadBalancerAttachment:getInstancePoolLoadBalancerAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstancePoolLoadBalancerAttachment.
type GetInstancePoolLoadBalancerAttachmentArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
	InstancePoolId string `pulumi:"instancePoolId"`
	// The OCID of the load balancer attachment.
	InstancePoolLoadBalancerAttachmentId string `pulumi:"instancePoolLoadBalancerAttachmentId"`
}

// A collection of values returned by getInstancePoolLoadBalancerAttachment.
type GetInstancePoolLoadBalancerAttachmentResult struct {
	// The name of the backend set on the load balancer.
	BackendSetName string `pulumi:"backendSetName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool of the load balancer attachment.
	InstancePoolId                       string `pulumi:"instancePoolId"`
	InstancePoolLoadBalancerAttachmentId string `pulumi:"instancePoolLoadBalancerAttachmentId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer attached to the instance pool.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The port value used for the backends.
	Port int `pulumi:"port"`
	// The status of the interaction between the instance pool and the load balancer.
	State string `pulumi:"state"`
	// Indicates which VNIC on each instance in the instance pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated with the instance pool.
	VnicSelection string `pulumi:"vnicSelection"`
}