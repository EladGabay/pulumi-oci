// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Network Security Group Vnics in Oracle Cloud Infrastructure Core service.
//
// Lists the VNICs in the specified network security group.
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
// 		_, err := core.GetNetworkSecurityGroupVnics(ctx, &core.GetNetworkSecurityGroupVnicsArgs{
// 			NetworkSecurityGroupId: oci_core_network_security_group.Test_network_security_group.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetNetworkSecurityGroupVnics(ctx *pulumi.Context, args *GetNetworkSecurityGroupVnicsArgs, opts ...pulumi.InvokeOption) (*GetNetworkSecurityGroupVnicsResult, error) {
	var rv GetNetworkSecurityGroupVnicsResult
	err := ctx.Invoke("oci:core/getNetworkSecurityGroupVnics:getNetworkSecurityGroupVnics", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkSecurityGroupVnics.
type GetNetworkSecurityGroupVnicsArgs struct {
	Filters []GetNetworkSecurityGroupVnicsFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
}

// A collection of values returned by getNetworkSecurityGroupVnics.
type GetNetworkSecurityGroupVnicsResult struct {
	Filters []GetNetworkSecurityGroupVnicsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string `pulumi:"id"`
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The list of network_security_group_vnics.
	NetworkSecurityGroupVnics []GetNetworkSecurityGroupVnicsNetworkSecurityGroupVnic `pulumi:"networkSecurityGroupVnics"`
}