// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Ping Monitors in Oracle Cloud Infrastructure Health Checks service.
 *
 * Gets a list of configured ping monitors.
 *
 * Results are paginated based on `page` and `limit`.  The `opc-next-page` header provides
 * a URL for fetching the next page.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testPingMonitors = oci.healthchecks.getPingMonitors({
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.ping_monitor_display_name,
 *     homeRegion: _var.ping_monitor_home_region,
 * });
 * ```
 */
export function getPingMonitors(args: GetPingMonitorsArgs, opts?: pulumi.InvokeOptions): Promise<GetPingMonitorsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:healthchecks/getPingMonitors:getPingMonitors", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "filters": args.filters,
        "homeRegion": args.homeRegion,
    }, opts);
}

/**
 * A collection of arguments for invoking getPingMonitors.
 */
export interface GetPingMonitorsArgs {
    /**
     * Filters results by compartment.
     */
    compartmentId: string;
    /**
     * Filters results that exactly match the `displayName` field.
     */
    displayName?: string;
    filters?: inputs.healthchecks.GetPingMonitorsFilter[];
    /**
     * Filters results that match the `homeRegion`.
     */
    homeRegion?: string;
}

/**
 * A collection of values returned by getPingMonitors.
 */
export interface GetPingMonitorsResult {
    /**
     * The OCID of the compartment.
     */
    readonly compartmentId: string;
    /**
     * A user-friendly and mutable name suitable for display in a user interface.
     */
    readonly displayName?: string;
    readonly filters?: outputs.healthchecks.GetPingMonitorsFilter[];
    /**
     * The region where updates must be made and where results must be fetched from.
     */
    readonly homeRegion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of ping_monitors.
     */
    readonly pingMonitors: outputs.healthchecks.GetPingMonitorsPingMonitor[];
}