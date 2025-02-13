// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Cross Connect Locations in Oracle Cloud Infrastructure Core service.
//
// Lists the available FastConnect locations for cross-connect installation. You need
// this information so you can specify your desired location when you create a cross-connect.
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
// 		_, err := core.GetCrossConnectLocations(ctx, &core.GetCrossConnectLocationsArgs{
// 			CompartmentId: _var.Compartment_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCrossConnectLocations(ctx *pulumi.Context, args *GetCrossConnectLocationsArgs, opts ...pulumi.InvokeOption) (*GetCrossConnectLocationsResult, error) {
	var rv GetCrossConnectLocationsResult
	err := ctx.Invoke("oci:core/getCrossConnectLocations:getCrossConnectLocations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCrossConnectLocations.
type GetCrossConnectLocationsArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string                           `pulumi:"compartmentId"`
	Filters       []GetCrossConnectLocationsFilter `pulumi:"filters"`
}

// A collection of values returned by getCrossConnectLocations.
type GetCrossConnectLocationsResult struct {
	CompartmentId string `pulumi:"compartmentId"`
	// The list of cross_connect_locations.
	CrossConnectLocations []GetCrossConnectLocationsCrossConnectLocation `pulumi:"crossConnectLocations"`
	Filters               []GetCrossConnectLocationsFilter               `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
