// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Db System Shapes in Oracle Cloud Infrastructure Database service.
//
// Gets a list of the shapes that can be used to launch a new DB system. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
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
// 		opt0 := _var.Db_system_shape_availability_domain
// 		_, err := oci.GetDatabaseDbSystemShapes(ctx, &GetDatabaseDbSystemShapesArgs{
// 			CompartmentId:      _var.Compartment_id,
// 			AvailabilityDomain: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatabaseDbSystemShapes(ctx *pulumi.Context, args *GetDatabaseDbSystemShapesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseDbSystemShapesResult, error) {
	var rv GetDatabaseDbSystemShapesResult
	err := ctx.Invoke("oci:index/getDatabaseDbSystemShapes:GetDatabaseDbSystemShapes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseDbSystemShapes.
type GetDatabaseDbSystemShapesArgs struct {
	// The name of the Availability Domain.
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId string                            `pulumi:"compartmentId"`
	Filters       []GetDatabaseDbSystemShapesFilter `pulumi:"filters"`
}

// A collection of values returned by GetDatabaseDbSystemShapes.
type GetDatabaseDbSystemShapesResult struct {
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	CompartmentId      string  `pulumi:"compartmentId"`
	// The list of db_system_shapes.
	DbSystemShapes []GetDatabaseDbSystemShapesDbSystemShape `pulumi:"dbSystemShapes"`
	Filters        []GetDatabaseDbSystemShapesFilter        `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}