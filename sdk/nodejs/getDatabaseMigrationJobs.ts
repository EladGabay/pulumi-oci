// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Jobs in Oracle Cloud Infrastructure Database Migration service.
 *
 * List all the names of the Migration jobs associated to the specified
 * migration site.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testJobs = oci.GetDatabaseMigrationJobs({
 *     migrationId: oci_database_migration_migration.test_migration.id,
 *     displayName: _var.job_display_name,
 *     state: _var.job_state,
 * });
 * ```
 */
export function getDatabaseMigrationJobs(args: GetDatabaseMigrationJobsArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseMigrationJobsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getDatabaseMigrationJobs:GetDatabaseMigrationJobs", {
        "displayName": args.displayName,
        "filters": args.filters,
        "migrationId": args.migrationId,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking GetDatabaseMigrationJobs.
 */
export interface GetDatabaseMigrationJobsArgs {
    /**
     * A filter to return only resources that match the entire display name given.
     */
    displayName?: string;
    filters?: inputs.GetDatabaseMigrationJobsFilter[];
    /**
     * The ID of the migration in which to list resources.
     */
    migrationId: string;
    /**
     * The lifecycle state of the Migration Job.
     */
    state?: string;
}

/**
 * A collection of values returned by GetDatabaseMigrationJobs.
 */
export interface GetDatabaseMigrationJobsResult {
    /**
     * Name of the job.
     */
    readonly displayName?: string;
    readonly filters?: outputs.GetDatabaseMigrationJobsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of job_collection.
     */
    readonly jobCollections: outputs.GetDatabaseMigrationJobsJobCollection[];
    /**
     * The OCID of the Migration that this job belongs to.
     */
    readonly migrationId: string;
    /**
     * The current state of the migration job.
     */
    readonly state?: string;
}