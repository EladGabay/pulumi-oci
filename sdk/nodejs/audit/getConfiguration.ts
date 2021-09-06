// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides details about a specific Configuration resource in Oracle Cloud Infrastructure Audit service.
 *
 * Get the configuration
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testConfiguration = oci.audit.getConfiguration({
 *     compartmentId: _var.tenancy_ocid,
 * });
 * ```
 */
export function getConfiguration(args: GetConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:audit/getConfiguration:getConfiguration", {
        "compartmentId": args.compartmentId,
    }, opts);
}

/**
 * A collection of arguments for invoking getConfiguration.
 */
export interface GetConfigurationArgs {
    /**
     * ID of the root compartment (tenancy)
     */
    compartmentId: string;
}

/**
 * A collection of values returned by getConfiguration.
 */
export interface GetConfigurationResult {
    readonly compartmentId: string;
    readonly id: string;
    /**
     * The retention period setting, specified in days. The minimum is 90, the maximum 365.  Example: `90`
     */
    readonly retentionPeriodDays: number;
}