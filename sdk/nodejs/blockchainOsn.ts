// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource provides the Osn resource in Oracle Cloud Infrastructure Blockchain service.
 *
 * Create Blockchain Platform Osn
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testOsn = new oci.BlockchainOsn("testOsn", {
 *     ad: _var.osn_ad,
 *     blockchainPlatformId: oci_blockchain_blockchain_platform.test_blockchain_platform.id,
 *     ocpuAllocationParam: {
 *         ocpuAllocationNumber: _var.osn_ocpu_allocation_param_ocpu_allocation_number,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Osns can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:index/blockchainOsn:BlockchainOsn test_osn "blockchainPlatforms/{blockchainPlatformId}/osns/{osnId}"
 * ```
 */
export class BlockchainOsn extends pulumi.CustomResource {
    /**
     * Get an existing BlockchainOsn resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BlockchainOsnState, opts?: pulumi.CustomResourceOptions): BlockchainOsn {
        return new BlockchainOsn(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/blockchainOsn:BlockchainOsn';

    /**
     * Returns true if the given object is an instance of BlockchainOsn.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BlockchainOsn {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BlockchainOsn.__pulumiType;
    }

    /**
     * Availability Domain to place new OSN
     */
    public readonly ad!: pulumi.Output<string>;
    /**
     * Unique service identifier.
     */
    public readonly blockchainPlatformId!: pulumi.Output<string>;
    /**
     * (Updatable) OCPU allocation parameter
     */
    public readonly ocpuAllocationParam!: pulumi.Output<outputs.BlockchainOsnOcpuAllocationParam>;
    /**
     * OSN identifier
     */
    public /*out*/ readonly osnKey!: pulumi.Output<string>;
    /**
     * The current state of the OSN.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a BlockchainOsn resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlockchainOsnArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BlockchainOsnArgs | BlockchainOsnState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BlockchainOsnState | undefined;
            inputs["ad"] = state ? state.ad : undefined;
            inputs["blockchainPlatformId"] = state ? state.blockchainPlatformId : undefined;
            inputs["ocpuAllocationParam"] = state ? state.ocpuAllocationParam : undefined;
            inputs["osnKey"] = state ? state.osnKey : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as BlockchainOsnArgs | undefined;
            if ((!args || args.ad === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ad'");
            }
            if ((!args || args.blockchainPlatformId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'blockchainPlatformId'");
            }
            inputs["ad"] = args ? args.ad : undefined;
            inputs["blockchainPlatformId"] = args ? args.blockchainPlatformId : undefined;
            inputs["ocpuAllocationParam"] = args ? args.ocpuAllocationParam : undefined;
            inputs["osnKey"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BlockchainOsn.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BlockchainOsn resources.
 */
export interface BlockchainOsnState {
    /**
     * Availability Domain to place new OSN
     */
    ad?: pulumi.Input<string>;
    /**
     * Unique service identifier.
     */
    blockchainPlatformId?: pulumi.Input<string>;
    /**
     * (Updatable) OCPU allocation parameter
     */
    ocpuAllocationParam?: pulumi.Input<inputs.BlockchainOsnOcpuAllocationParam>;
    /**
     * OSN identifier
     */
    osnKey?: pulumi.Input<string>;
    /**
     * The current state of the OSN.
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BlockchainOsn resource.
 */
export interface BlockchainOsnArgs {
    /**
     * Availability Domain to place new OSN
     */
    ad: pulumi.Input<string>;
    /**
     * Unique service identifier.
     */
    blockchainPlatformId: pulumi.Input<string>;
    /**
     * (Updatable) OCPU allocation parameter
     */
    ocpuAllocationParam?: pulumi.Input<inputs.BlockchainOsnOcpuAllocationParam>;
}