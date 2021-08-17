// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Database resource in Oracle Cloud Infrastructure Database service.
//
// Gets information about the specified database.
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
// 		_, err := oci.GetDatabaseDatabase(ctx, &GetDatabaseDatabaseArgs{
// 			DatabaseId: _var.Database_id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDatabaseDatabase(ctx *pulumi.Context, args *LookupDatabaseDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseDatabaseResult, error) {
	var rv LookupDatabaseDatabaseResult
	err := ctx.Invoke("oci:index/getDatabaseDatabase:GetDatabaseDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetDatabaseDatabase.
type LookupDatabaseDatabaseArgs struct {
	// The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseId string `pulumi:"databaseId"`
}

// A collection of values returned by GetDatabaseDatabase.
type LookupDatabaseDatabaseResult struct {
	// The character set for the database.
	CharacterSet string `pulumi:"characterSet"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// The Connection strings used to connect to the Oracle Database.
	ConnectionStrings GetDatabaseDatabaseConnectionStrings `pulumi:"connectionStrings"`
	Database          GetDatabaseDatabaseDatabase          `pulumi:"database"`
	DatabaseId        string                               `pulumi:"databaseId"`
	// The database software image [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	DatabaseSoftwareImageId string `pulumi:"databaseSoftwareImageId"`
	// Backup Options To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator. If you're an administrator who needs to write policies to give users access, see [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
	DbBackupConfig GetDatabaseDatabaseDbBackupConfig `pulumi:"dbBackupConfig"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId string `pulumi:"dbHomeId"`
	// The database name.
	DbName string `pulumi:"dbName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId string `pulumi:"dbSystemId"`
	// A system-generated name for the database to ensure uniqueness within an Oracle Data Guard group (a primary database and its standby databases). The unique name cannot be changed.
	DbUniqueName string `pulumi:"dbUniqueName"`
	DbVersion    string `pulumi:"dbVersion"`
	// The database workload type.
	DbWorkload string `pulumi:"dbWorkload"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	Id string `pulumi:"id"`
	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId        string `pulumi:"kmsKeyId"`
	KmsKeyMigration bool   `pulumi:"kmsKeyMigration"`
	KmsKeyRotation  int    `pulumi:"kmsKeyRotation"`
	KmsKeyVersionId string `pulumi:"kmsKeyVersionId"`
	// The date and time when the latest database backup was created.
	LastBackupTimestamp string `pulumi:"lastBackupTimestamp"`
	// Additional information about the current lifecycle state.
	LifecycleDetails string `pulumi:"lifecycleDetails"`
	// The national character set for the database.
	NcharacterSet string `pulumi:"ncharacterSet"`
	// The name of the pluggable database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. Pluggable database should not be same as database name.
	PdbName string `pulumi:"pdbName"`
	Source  string `pulumi:"source"`
	// Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system, as described in [RFC 3339](https://tools.ietf.org/rfc/rfc3339)
	SourceDatabasePointInTimeRecoveryTimestamp string `pulumi:"sourceDatabasePointInTimeRecoveryTimestamp"`
	// The current state of the database.
	State string `pulumi:"state"`
	// The date and time the database was created.
	TimeCreated string `pulumi:"timeCreated"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
	VmClusterId string `pulumi:"vmClusterId"`
}