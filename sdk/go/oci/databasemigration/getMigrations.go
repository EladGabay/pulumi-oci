// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Migrations in Oracle Cloud Infrastructure Database Migration service.
//
// List all Migrations.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/databasemigration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := _var.Migration_display_name
// 		opt1 := _var.Migration_lifecycle_details
// 		opt2 := _var.Migration_state
// 		_, err := databasemigration.GetMigrations(ctx, &databasemigration.GetMigrationsArgs{
// 			CompartmentId:    _var.Compartment_id,
// 			DisplayName:      &opt0,
// 			LifecycleDetails: &opt1,
// 			State:            &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetMigrations(ctx *pulumi.Context, args *GetMigrationsArgs, opts ...pulumi.InvokeOption) (*GetMigrationsResult, error) {
	var rv GetMigrationsResult
	err := ctx.Invoke("oci:databasemigration/getMigrations:getMigrations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMigrations.
type GetMigrationsArgs struct {
	// The ID of the compartment in which to list resources.
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only resources that match the entire display name given.
	DisplayName *string               `pulumi:"displayName"`
	Filters     []GetMigrationsFilter `pulumi:"filters"`
	// The lifecycle detailed status of the Migration.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The current state of the Database Migration Deployment.
	State *string `pulumi:"state"`
}

// A collection of values returned by getMigrations.
type GetMigrationsResult struct {
	// OCID of the compartment where the secret containing the credentials will be created.
	CompartmentId string `pulumi:"compartmentId"`
	// Migration Display Name
	DisplayName *string               `pulumi:"displayName"`
	Filters     []GetMigrationsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Additional status related to the execution and current state of the Migration.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The list of migration_collection.
	MigrationCollections []GetMigrationsMigrationCollection `pulumi:"migrationCollections"`
	// The current state of the Migration Resource.
	State *string `pulumi:"state"`
}