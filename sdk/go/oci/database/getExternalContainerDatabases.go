// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of External Container Databases in Oracle Cloud Infrastructure Database service.
//
// Gets a list of the external container databases in the specified compartment.
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
// 		opt0 := _var.External_container_database_display_name
// 		opt1 := _var.External_container_database_state
// 		_, err := database.GetExternalContainerDatabases(ctx, &database.GetExternalContainerDatabasesArgs{
// 			CompartmentId: _var.Compartment_id,
// 			DisplayName:   &opt0,
// 			State:         &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetExternalContainerDatabases(ctx *pulumi.Context, args *GetExternalContainerDatabasesArgs, opts ...pulumi.InvokeOption) (*GetExternalContainerDatabasesResult, error) {
	var rv GetExternalContainerDatabasesResult
	err := ctx.Invoke("oci:database/getExternalContainerDatabases:getExternalContainerDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getExternalContainerDatabases.
type GetExternalContainerDatabasesArgs struct {
	// The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId string `pulumi:"compartmentId"`
	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string                               `pulumi:"displayName"`
	Filters     []GetExternalContainerDatabasesFilter `pulumi:"filters"`
	// A filter to return only resources that match the specified lifecycle state.
	State *string `pulumi:"state"`
}

// A collection of values returned by getExternalContainerDatabases.
type GetExternalContainerDatabasesResult struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The user-friendly name for the external database. The name does not have to be unique.
	DisplayName *string `pulumi:"displayName"`
	// The list of external_container_databases.
	ExternalContainerDatabases []GetExternalContainerDatabasesExternalContainerDatabase `pulumi:"externalContainerDatabases"`
	Filters                    []GetExternalContainerDatabasesFilter                    `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The current state of the Oracle Cloud Infrastructure external database resource.
	State *string `pulumi:"state"`
}
