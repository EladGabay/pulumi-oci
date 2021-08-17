// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Database Software Image resource in Oracle Cloud Infrastructure Database service.
//
// Gets information about the specified database software image.
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
// 		_, err := oci.GetDatabaseDatabaseSoftwareImage(ctx, &GetDatabaseDatabaseSoftwareImageArgs{
// 			DatabaseSoftwareImageId: oci_database_database_software_image.Test_database_software_image.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatabaseDatabaseSoftwareImage(ctx *pulumi.Context, args *LookupDatabaseDatabaseSoftwareImageArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseDatabaseSoftwareImageResult, error) {
	var rv LookupDatabaseDatabaseSoftwareImageResult
	err := ctx.Invoke("oci:index/getDatabaseDatabaseSoftwareImage:GetDatabaseDatabaseSoftwareImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseDatabaseSoftwareImage.
type LookupDatabaseDatabaseSoftwareImageArgs struct {
	// The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseSoftwareImageId string `pulumi:"databaseSoftwareImageId"`
}

// A collection of values returned by GetDatabaseDatabaseSoftwareImage.
type LookupDatabaseDatabaseSoftwareImageResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId           string `pulumi:"compartmentId"`
	DatabaseSoftwareImageId string `pulumi:"databaseSoftwareImageId"`
	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageIncludedPatches []string `pulumi:"databaseSoftwareImageIncludedPatches"`
	// List of one-off patches for Database Homes.
	DatabaseSoftwareImageOneOffPatches []string `pulumi:"databaseSoftwareImageOneOffPatches"`
	// The database version with which the database software image is to be built.
	DatabaseVersion string `pulumi:"databaseVersion"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The user-friendly name for the database software image. The name does not have to be unique.
	DisplayName string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database software image.
	Id string `pulumi:"id"`
	// To what shape the image is meant for.
	ImageShapeFamily string `pulumi:"imageShapeFamily"`
	// The type of software image. Can be grid or database.
	ImageType string `pulumi:"imageType"`
	// The patches included in the image and the version of the image
	IncludedPatchesSummary string `pulumi:"includedPatchesSummary"`
	// True if this Database software image is supported for Upgrade.
	IsUpgradeSupported bool `pulumi:"isUpgradeSupported"`
	// Detailed message for the lifecycle state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// output from lsinventory which will get passed as a string
	LsInventory string `pulumi:"lsInventory"`
	// The PSU or PBP or Release Updates. To get a list of supported versions, use the [ListDbVersions](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/DbVersionSummary/ListDbVersions) operation.
	PatchSet       string `pulumi:"patchSet"`
	SourceDbHomeId string `pulumi:"sourceDbHomeId"`
	// The current state of the database software image.
	State string `pulumi:"state"`
	// The date and time the database software image was created.
	TimeCreated string `pulumi:"timeCreated"`
}