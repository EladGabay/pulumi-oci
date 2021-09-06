// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Service Gateways in Oracle Cloud Infrastructure Core service.
 *
 * Lists the service gateways in the specified compartment. You may optionally specify a VCN OCID
 * to filter the results by VCN.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testServiceGateways = oci.core.getServiceGateways({
 *     compartmentId: _var.compartment_id,
 *     state: _var.service_gateway_state,
 *     vcnId: oci_core_vcn.test_vcn.id,
 * });
 * ```
 */
export function getServiceGateways(args: GetServiceGatewaysArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceGatewaysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:core/getServiceGateways:getServiceGateways", {
        "compartmentId": args.compartmentId,
        "filters": args.filters,
        "state": args.state,
        "vcnId": args.vcnId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceGateways.
 */
export interface GetServiceGatewaysArgs {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
     */
    compartmentId: string;
    filters?: inputs.core.GetServiceGatewaysFilter[];
    /**
     * A filter to return only resources that match the given lifecycle state. The state value is case-insensitive.
     */
    state?: string;
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.
     */
    vcnId?: string;
}

/**
 * A collection of values returned by getServiceGateways.
 */
export interface GetServiceGatewaysResult {
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the service gateway.
     */
    readonly compartmentId: string;
    readonly filters?: outputs.core.GetServiceGatewaysFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of service_gateways.
     */
    readonly serviceGateways: outputs.core.GetServiceGatewaysServiceGateway[];
    /**
     * The service gateway's current state.
     */
    readonly state?: string;
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the service gateway belongs to.
     */
    readonly vcnId?: string;
}