// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Network Security Group Vnics in Oracle Cloud Infrastructure Core service.
 *
 * Lists the VNICs in the specified network security group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testNetworkSecurityGroupVnics = oci.GetCoreNetworkSecurityGroupVnics({
 *     networkSecurityGroupId: oci_core_network_security_group.test_network_security_group.id,
 * });
 * ```
 */
export function getCoreNetworkSecurityGroupVnics(args: GetCoreNetworkSecurityGroupVnicsArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreNetworkSecurityGroupVnicsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getCoreNetworkSecurityGroupVnics:GetCoreNetworkSecurityGroupVnics", {
        "filters": args.filters,
        "networkSecurityGroupId": args.networkSecurityGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking GetCoreNetworkSecurityGroupVnics.
 */
export interface GetCoreNetworkSecurityGroupVnicsArgs {
    filters?: inputs.GetCoreNetworkSecurityGroupVnicsFilter[];
    /**
     * The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
     */
    networkSecurityGroupId: string;
}

/**
 * A collection of values returned by GetCoreNetworkSecurityGroupVnics.
 */
export interface GetCoreNetworkSecurityGroupVnicsResult {
    readonly filters?: outputs.GetCoreNetworkSecurityGroupVnicsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly networkSecurityGroupId: string;
    /**
     * The list of network_security_group_vnics.
     */
    readonly networkSecurityGroupVnics: outputs.GetCoreNetworkSecurityGroupVnicsNetworkSecurityGroupVnic[];
}