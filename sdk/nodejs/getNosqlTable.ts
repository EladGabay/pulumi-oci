// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides details about a specific Table resource in Oracle Cloud Infrastructure NoSQL Database service.
 *
 * Get table info by identifier.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testTable = oci.GetNosqlTable({
 *     tableNameOrId: oci_nosql_table_name_or.test_table_name_or.id,
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getNosqlTable(args: GetNosqlTableArgs, opts?: pulumi.InvokeOptions): Promise<GetNosqlTableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getNosqlTable:GetNosqlTable", {
        "compartmentId": args.compartmentId,
        "tableNameOrId": args.tableNameOrId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetNosqlTable.
 */
export interface GetNosqlTableArgs {
    /**
     * The ID of a table's compartment. When a table is identified by name, the compartmentId is often needed to provide context for interpreting the name.
     */
    compartmentId: string;
    /**
     * A table name within the compartment, or a table OCID.
     */
    tableNameOrId: string;
}

/**
 * A collection of values returned by GetNosqlTable.
 */
export interface GetNosqlTableResult {
    /**
     * Compartment Identifier.
     */
    readonly compartmentId: string;
    /**
     * A DDL statement representing the schema.
     */
    readonly ddlStatement: string;
    /**
     * Defined tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"foo-namespace": {"bar-key": "value"}}`
     */
    readonly definedTags: {[key: string]: any};
    /**
     * Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
     */
    readonly freeformTags: {[key: string]: any};
    /**
     * Unique identifier that is immutable.
     */
    readonly id: string;
    /**
     * True if this table can be reclaimed after an idle period.
     */
    readonly isAutoReclaimable: boolean;
    /**
     * A message describing the current state in more detail.
     */
    readonly lifecycleDetails: string;
    /**
     * The column name.
     */
    readonly name: string;
    /**
     * The table schema information as a JSON object.
     */
    readonly schema: outputs.GetNosqlTableSchema;
    /**
     * The state of a table.
     */
    readonly state: string;
    /**
     * Read-only system tag. These predefined keys are scoped to namespaces.  At present the only supported namespace is `"orcl-cloud"`; and the only key in that namespace is `"free-tier-retained"`. Example: `{"orcl-cloud"": {"free-tier-retained": "true"}}`
     */
    readonly systemTags: {[key: string]: any};
    /**
     * Throughput and storage limits configuration of a table.
     */
    readonly tableLimits: outputs.GetNosqlTableTableLimits;
    readonly tableNameOrId: string;
    /**
     * The time the the table was created. An RFC3339 formatted datetime string.
     */
    readonly timeCreated: string;
    /**
     * If lifecycleState is INACTIVE, indicates when this table will be automatically removed. An RFC3339 formatted datetime string.
     */
    readonly timeOfExpiration: string;
    /**
     * The time the the table's metadata was last updated. An RFC3339 formatted datetime string.
     */
    readonly timeUpdated: string;
}