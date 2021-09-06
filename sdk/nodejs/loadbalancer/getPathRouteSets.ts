// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Path Route Sets in Oracle Cloud Infrastructure Load Balancer service.
 *
 * Lists all path route sets associated with the specified load balancer.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testPathRouteSets = oci.loadbalancer.getPathRouteSets({
 *     loadBalancerId: oci_load_balancer_load_balancer.test_load_balancer.id,
 * });
 * ```
 */
export function getPathRouteSets(args: GetPathRouteSetsArgs, opts?: pulumi.InvokeOptions): Promise<GetPathRouteSetsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:loadbalancer/getPathRouteSets:getPathRouteSets", {
        "filters": args.filters,
        "loadBalancerId": args.loadBalancerId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPathRouteSets.
 */
export interface GetPathRouteSetsArgs {
    filters?: inputs.loadbalancer.GetPathRouteSetsFilter[];
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the load balancer associated with the path route sets to retrieve.
     */
    loadBalancerId: string;
}

/**
 * A collection of values returned by getPathRouteSets.
 */
export interface GetPathRouteSetsResult {
    readonly filters?: outputs.loadbalancer.GetPathRouteSetsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loadBalancerId: string;
    /**
     * The list of path_route_sets.
     */
    readonly pathRouteSets: outputs.loadbalancer.GetPathRouteSetsPathRouteSet[];
}