// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Cross Connect Port Speed Shapes in Oracle Cloud Infrastructure Core service.
 *
 * Lists the available port speeds for cross-connects. You need this information
 * so you can specify your desired port speed (that is, shape) when you create a
 * cross-connect.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testCrossConnectPortSpeedShapes = oci.GetCoreCrossConnectPortSpeedShapes({
 *     compartmentId: _var.compartment_id,
 * });
 * ```
 */
export function getCoreCrossConnectPortSpeedShapes(args: GetCoreCrossConnectPortSpeedShapesArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreCrossConnectPortSpeedShapesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreCrossConnectPortSpeedShapes:GetCoreCrossConnectPortSpeedShapes", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreCrossConnectPortSpeedShapes.
 */
export interface GetCoreCrossConnectPortSpeedShapesArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    filters?: inputs.GetCoreCrossConnectPortSpeedShapesFilter[];
}

/**
 * A collection of values returned by GetCoreCrossConnectPortSpeedShapes.
 */
export interface GetCoreCrossConnectPortSpeedShapesResult {
    readonly compartmentId: string;
    /**
     * The list of cross_connect_port_speed_shapes.
     */
    readonly crossConnectPortSpeedShapes: outputs.GetCoreCrossConnectPortSpeedShapesCrossConnectPortSpeedShape[];
    readonly filters?: outputs.GetCoreCrossConnectPortSpeedShapesFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}