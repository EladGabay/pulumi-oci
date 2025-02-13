// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Backup Destinations in Oracle Cloud Infrastructure Database service.
//
// Gets a list of backup destinations in the specified compartment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/database"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Backup_destination_type
// 		_, err := database.GetBackupDestinations(ctx, &database.GetBackupDestinationsArgs{
// 			CompartmentId: _var.Compartment_id,
// 			Type:          &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetBackupDestinations(ctx *pulumi.Context, args *GetBackupDestinationsArgs, opts ...pulumi.InvokeOption) (*GetBackupDestinationsResult, error) {
	var rv GetBackupDestinationsResult
	err := ctx.Invoke("oci:database/getBackupDestinations:getBackupDestinations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBackupDestinations.
type GetBackupDestinationsArgs struct {
	// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId string                        `pulumi:"compartmentId"`
	Filters       []GetBackupDestinationsFilter `pulumi:"filters"`
	// A filter to return only resources that match the given type of the Backup Destination.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getBackupDestinations.
type GetBackupDestinationsResult struct {
	// The list of backup_destinations.
	BackupDestinations []GetBackupDestinationsBackupDestination `pulumi:"backupDestinations"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string                        `pulumi:"compartmentId"`
	Filters       []GetBackupDestinationsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Type of the backup destination.
	Type *string `pulumi:"type"`
}
