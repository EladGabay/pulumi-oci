// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Remote Peering Connections in Oracle Cloud Infrastructure Core service.
//
// Lists the remote peering connections (RPCs) for the specified DRG and compartment
// (the RPC's compartment).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := oci_core_drg.Test_drg.Id
// 		_, err := core.GetRemotePeeringConnections(ctx, &core.GetRemotePeeringConnectionsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DrgId:         &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetRemotePeeringConnections(ctx *pulumi.Context, args *GetRemotePeeringConnectionsArgs, opts ...pulumi.InvokeOption) (*GetRemotePeeringConnectionsResult, error) {
	var rv GetRemotePeeringConnectionsResult
	err := ctx.Invoke("oci:core/getRemotePeeringConnections:getRemotePeeringConnections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemotePeeringConnections.
type GetRemotePeeringConnectionsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG.
	DrgId   *string                             `pulumi:"drgId"`
	Filters []GetRemotePeeringConnectionsFilter `pulumi:"filters"`
}

// A collection of values returned by getRemotePeeringConnections.
type GetRemotePeeringConnectionsResult struct {
	// The OCID of the compartment that contains the RPC.
	CompartmentId string `pulumi:"compartmentId"`
	// The OCID of the DRG that this RPC belongs to.
	DrgId   *string                             `pulumi:"drgId"`
	Filters []GetRemotePeeringConnectionsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of remote_peering_connections.
	RemotePeeringConnections []GetRemotePeeringConnectionsRemotePeeringConnection `pulumi:"remotePeeringConnections"`
}
