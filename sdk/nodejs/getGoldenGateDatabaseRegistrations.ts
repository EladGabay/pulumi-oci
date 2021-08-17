// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Database Registrations in Oracle Cloud Infrastructure Golden Gate service.
 *
 * Lists the DatabaseRegistrations in the compartment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testDatabaseRegistrations = oci.GetGoldenGateDatabaseRegistrations({
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.database_registration_display_name,
 *     state: _var.database_registration_state,
 * });
 * ```
 */
export function getGoldenGateDatabaseRegistrations(args: GetGoldenGateDatabaseRegistrationsArgs, opts?: pulumi.InvokeOptions): Promise<GetGoldenGateDatabaseRegistrationsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getGoldenGateDatabaseRegistrations:GetGoldenGateDatabaseRegistrations", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "filters": args.filters,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking GetGoldenGateDatabaseRegistrations.
 */
export interface GetGoldenGateDatabaseRegistrationsArgs {
    /**
     * The ID of the compartment in which to list resources.
     */
    compartmentId: string;
    /**
     * A filter to return only the resources that match the entire 'displayName' given.
     */
    displayName?: string;
    filters?: inputs.GetGoldenGateDatabaseRegistrationsFilter[];
    /**
     * A filter to return only the resources that match the 'lifecycleState' given.
     */
    state?: string;
}

/**
 * A collection of values returned by GetGoldenGateDatabaseRegistrations.
 */
export interface GetGoldenGateDatabaseRegistrationsResult {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
     */
    readonly compartmentId: string;
    /**
     * The list of database_registration_collection.
     */
    readonly databaseRegistrationCollections: outputs.GetGoldenGateDatabaseRegistrationsDatabaseRegistrationCollection[];
    /**
     * An object's Display Name.
     */
    readonly displayName?: string;
    readonly filters?: outputs.GetGoldenGateDatabaseRegistrationsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Possible lifecycle states.
     */
    readonly state?: string;
}