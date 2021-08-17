// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Network Security Group Security Rules in Oracle Cloud Infrastructure Core service.
//
// Lists the security rules in the specified network security group.
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
// 		opt0 := _var.Network_security_group_security_rule_direction
// 		_, err := oci.GetCoreNetworkSecurityGroupSecurityRules(ctx, &GetCoreNetworkSecurityGroupSecurityRulesArgs{
// 			NetworkSecurityGroupId: oci_core_network_security_group.Test_network_security_group.Id,
// 			Direction:              &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCoreNetworkSecurityGroupSecurityRules(ctx *pulumi.Context, args *GetCoreNetworkSecurityGroupSecurityRulesArgs, opts ...pulumi.InvokeOption) (*GetCoreNetworkSecurityGroupSecurityRulesResult, error) {
	var rv GetCoreNetworkSecurityGroupSecurityRulesResult
	err := ctx.Invoke("oci:index/getCoreNetworkSecurityGroupSecurityRules:GetCoreNetworkSecurityGroupSecurityRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetCoreNetworkSecurityGroupSecurityRules.
type GetCoreNetworkSecurityGroupSecurityRulesArgs struct {
	// Direction of the security rule. Set to `EGRESS` for rules that allow outbound IP packets, or `INGRESS` for rules that allow inbound IP packets.
	Direction *string                                          `pulumi:"direction"`
	Filters   []GetCoreNetworkSecurityGroupSecurityRulesFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
}

// A collection of values returned by GetCoreNetworkSecurityGroupSecurityRules.
type GetCoreNetworkSecurityGroupSecurityRulesResult struct {
	// Direction of the security rule. Set to `EGRESS` for rules to allow outbound IP packets, or `INGRESS` for rules to allow inbound IP packets.
	Direction *string                                          `pulumi:"direction"`
	Filters   []GetCoreNetworkSecurityGroupSecurityRulesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string `pulumi:"id"`
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The list of security_rules.
	SecurityRules []GetCoreNetworkSecurityGroupSecurityRulesSecurityRule `pulumi:"securityRules"`
}