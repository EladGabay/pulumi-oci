// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Ipv6s in Oracle Cloud Infrastructure Core service.
//
// Lists the [IPv6](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Ipv6/) objects based
// on one of these filters:
//
//   * Subnet OCID.
//   * VNIC OCID.
//   * Both IPv6 address and subnet OCID: This lets you get an `Ipv6` object based on its private
//       IPv6 address (for example, 2001:0db8:0123:1111:abcd:ef01:2345:6789) and not its OCID. For comparison,
//       [GetIpv6](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Ipv6/GetIpv6) requires the OCID.
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
// 		opt0 := _var.Ipv6_ip_address
// 		opt1 := oci_core_subnet.Test_subnet.Id
// 		opt2 := oci_core_vnic_attachment.Test_vnic_attachment.Id
// 		_, err := core.GetIpv6s(ctx, &core.GetIpv6sArgs{
// 			IpAddress: &opt0,
// 			SubnetId:  &opt1,
// 			VnicId:    &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIpv6s(ctx *pulumi.Context, args *GetIpv6sArgs, opts ...pulumi.InvokeOption) (*GetIpv6sResult, error) {
	var rv GetIpv6sResult
	err := ctx.Invoke("oci:core/getIpv6s:getIpv6s", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv6s.
type GetIpv6sArgs struct {
	Filters []GetIpv6sFilter `pulumi:"filters"`
	// An IP address. This could be either IPv4 or IPv6, depending on the resource. Example: `10.0.3.3`
	IpAddress *string `pulumi:"ipAddress"`
	// The OCID of the subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The OCID of the VNIC.
	VnicId *string `pulumi:"vnicId"`
}

// A collection of values returned by getIpv6s.
type GetIpv6sResult struct {
	Filters []GetIpv6sFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IPv6 address of the `IPv6` object. The address is within the IPv6 CIDR block of the VNIC's subnet (see the `ipv6CidrBlock` attribute for the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/) object.  Example: `2001:0db8:0123:1111:abcd:ef01:2345:6789`
	IpAddress *string `pulumi:"ipAddress"`
	// The list of ipv6s.
	Ipv6s []GetIpv6sIpv6 `pulumi:"ipv6s"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
	SubnetId *string `pulumi:"subnetId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC the IPv6 is assigned to. The VNIC and IPv6 must be in the same subnet.
	VnicId *string `pulumi:"vnicId"`
}
