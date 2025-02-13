// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Migration resource in Oracle Cloud Infrastructure Database Migration service.
//
// Display Migration details.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/databasemigration"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := databasemigration.LookupMigration(ctx, &databasemigration.LookupMigrationArgs{
// 			MigrationId: oci_database_migration_migration.Test_migration.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMigration(ctx *pulumi.Context, args *LookupMigrationArgs, opts ...pulumi.InvokeOption) (*LookupMigrationResult, error) {
	var rv LookupMigrationResult
	err := ctx.Invoke("oci:databasemigration/getMigration:getMigration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMigration.
type LookupMigrationArgs struct {
	// The OCID of the job
	MigrationId string `pulumi:"migrationId"`
}

// A collection of values returned by getMigration.
type LookupMigrationResult struct {
	// The OCID of the registered On-Prem ODMS Agent. Required for Offline Migrations.
	AgentId string `pulumi:"agentId"`
	// OCID of the compartment where the secret containing the credentials will be created.
	CompartmentId string `pulumi:"compartmentId"`
	// OCID of the Secret in the Oracle Cloud Infrastructure vault containing the Migration credentials. Used to store Golden Gate admin user credentials.
	CredentialsSecretId string `pulumi:"credentialsSecretId"`
	// Data Transfer Medium details for the Migration.
	DataTransferMediumDetails GetMigrationDataTransferMediumDetails `pulumi:"dataTransferMediumDetails"`
	// Optional settings for Datapump Export and Import jobs
	DatapumpSettings GetMigrationDatapumpSettings `pulumi:"datapumpSettings"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Migration Display Name
	DisplayName string `pulumi:"displayName"`
	// Database objects to exclude from migration.
	ExcludeObjects []GetMigrationExcludeObject `pulumi:"excludeObjects"`
	// OCID of the current ODMS Job in execution for the Migration, if any.
	ExecutingJobId string `pulumi:"executingJobId"`
	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Details about Oracle GoldenGate Microservices.
	GoldenGateDetails GetMigrationGoldenGateDetails `pulumi:"goldenGateDetails"`
	// The OCID of the resource
	Id string `pulumi:"id"`
	// Additional status related to the execution and current state of the Migration.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	MigrationId      string `pulumi:"migrationId"`
	// The OCID of the Source Container Database Connection.
	SourceContainerDatabaseConnectionId string `pulumi:"sourceContainerDatabaseConnectionId"`
	// The OCID of the Source Database Connection.
	SourceDatabaseConnectionId string `pulumi:"sourceDatabaseConnectionId"`
	// The current state of the Migration Resource.
	State string `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The OCID of the Target Database Connection.
	TargetDatabaseConnectionId string `pulumi:"targetDatabaseConnectionId"`
	// The time the Migration was created. An RFC3339 formatted datetime string.
	TimeCreated string `pulumi:"timeCreated"`
	// The time of last Migration. An RFC3339 formatted datetime string.
	TimeLastMigration string `pulumi:"timeLastMigration"`
	// The time of the last Migration details update. An RFC3339 formatted datetime string.
	TimeUpdated string `pulumi:"timeUpdated"`
	// Migration type.
	Type string `pulumi:"type"`
	// Oracle Cloud Infrastructure Vault details to store migration and connection credentials secrets
	VaultDetails GetMigrationVaultDetails `pulumi:"vaultDetails"`
	// Name of a migration phase. The Job will wait after executing this phase until the Resume Job endpoint is called.
	WaitAfter string `pulumi:"waitAfter"`
}
