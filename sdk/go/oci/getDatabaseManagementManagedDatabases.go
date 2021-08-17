// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Managed Databases in Oracle Cloud Infrastructure Database Management service.
//
// Gets the Managed Database for a specific ID or the list of Managed Databases in a specific compartment.
// Managed Databases can also be filtered based on the name parameter. Only one of the parameters, ID or name
// should be provided. If none of these parameters is provided, all the Managed Databases in the compartment are listed.
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
// 		opt0 := _var.Managed_database_id
// 		opt1 := _var.Managed_database_name
// 		_, err := oci.GetDatabaseManagementManagedDatabases(ctx, &GetDatabaseManagementManagedDatabasesArgs{
// 			CompartmentId: _var.Compartment_id,
// 			Id:            &opt0,
// 			Name:          &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDatabaseManagementManagedDatabases(ctx *pulumi.Context, args *GetDatabaseManagementManagedDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabaseManagementManagedDatabasesResult, error) {
	var rv GetDatabaseManagementManagedDatabasesResult
	err := ctx.Invoke("oci:index/getDatabaseManagementManagedDatabases:GetDatabaseManagementManagedDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseManagementManagedDatabases.
type GetDatabaseManagementManagedDatabasesArgs struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string                                        `pulumi:"compartmentId"`
	Filters       []GetDatabaseManagementManagedDatabasesFilter `pulumi:"filters"`
	// The identifier of the resource.
	Id *string `pulumi:"id"`
	// A filter to return only resources that match the entire name.
	Name *string `pulumi:"name"`
}

// A collection of values returned by GetDatabaseManagementManagedDatabases.
type GetDatabaseManagementManagedDatabasesResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides.
	CompartmentId string                                        `pulumi:"compartmentId"`
	Filters       []GetDatabaseManagementManagedDatabasesFilter `pulumi:"filters"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	Id *string `pulumi:"id"`
	// The list of managed_database_collection.
	ManagedDatabaseCollections []GetDatabaseManagementManagedDatabasesManagedDatabaseCollection `pulumi:"managedDatabaseCollections"`
	// The name of the Managed Database.
	Name *string `pulumi:"name"`
}