// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Exadata Iorm Config resource in Oracle Cloud Infrastructure Database service.
//
// Gets the IORM configuration settings for the specified cloud Exadata DB system.
// All Exadata service instances have default IORM settings.
//
// **Note:** Deprecated for Exadata Cloud Service systems. Use the [new resource model APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem.htm#exaflexsystem_topic-resource_model) instead.
//
// For Exadata Cloud Service instances, support for this API will end on May 15th, 2021. See [Switching an Exadata DB System to the New Resource Model and APIs](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/exaflexsystem_topic-resource_model_conversion.htm) for details on converting existing Exadata DB systems to the new resource model.
//
// The [GetCloudVmClusterIormConfig](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudVmCluster/GetCloudVmClusterIormConfig/) API is used for this operation with Exadata systems using the
// new resource model.
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
// 		_, err := oci.GetDatabaseExadataIormConfig(ctx, &GetDatabaseExadataIormConfigArgs{
// 			DbSystemId: oci_database_db_system.Test_db_system.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatabaseExadataIormConfig(ctx *pulumi.Context, args *LookupDatabaseExadataIormConfigArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseExadataIormConfigResult, error) {
	var rv LookupDatabaseExadataIormConfigResult
	err := ctx.Invoke("oci:index/getDatabaseExadataIormConfig:GetDatabaseExadataIormConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseExadataIormConfig.
type LookupDatabaseExadataIormConfigArgs struct {
	// The DB system [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId string `pulumi:"dbSystemId"`
}

// A collection of values returned by GetDatabaseExadataIormConfig.
type LookupDatabaseExadataIormConfigResult struct {
	// An array of IORM settings for all the database in the Exadata DB system.
	DbPlans    []GetDatabaseExadataIormConfigDbPlan `pulumi:"dbPlans"`
	DbSystemId string                               `pulumi:"dbSystemId"`
	Id         string                               `pulumi:"id"`
	// Additional information about the current `lifecycleState`.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// The current value for the IORM objective. The default is `AUTO`.
	Objective string `pulumi:"objective"`
	// The current state of IORM configuration for the Exadata DB system.
	State string `pulumi:"state"`
}