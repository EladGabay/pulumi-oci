// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This resource provides the Object Lifecycle Policy resource in Oracle Cloud Infrastructure Object Storage service.
 *
 * Creates or replaces the object lifecycle policy for the bucket.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testObjectLifecyclePolicy = new oci.ObjectstorageObjectLifecyclePolicy("testObjectLifecyclePolicy", {
 *     bucket: _var.object_lifecycle_policy_bucket,
 *     namespace: _var.object_lifecycle_policy_namespace,
 *     rules: [{
 *         action: _var.object_lifecycle_policy_rules_action,
 *         isEnabled: _var.object_lifecycle_policy_rules_is_enabled,
 *         name: _var.object_lifecycle_policy_rules_name,
 *         timeAmount: _var.object_lifecycle_policy_rules_time_amount,
 *         timeUnit: _var.object_lifecycle_policy_rules_time_unit,
 *         objectNameFilter: {
 *             exclusionPatterns: _var.object_lifecycle_policy_rules_object_name_filter_exclusion_patterns,
 *             inclusionPatterns: _var.object_lifecycle_policy_rules_object_name_filter_inclusion_patterns,
 *             inclusionPrefixes: _var.object_lifecycle_policy_rules_object_name_filter_inclusion_prefixes,
 *         },
 *         target: _var.object_lifecycle_policy_rules_target,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ObjectLifecyclePolicies can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:index/objectstorageObjectLifecyclePolicy:ObjectstorageObjectLifecyclePolicy test_object_lifecycle_policy "n/{namespaceName}/b/{bucketName}/l"
 * ```
 */
export class ObjectstorageObjectLifecyclePolicy extends pulumi.CustomResource {
    /**
     * Get an existing ObjectstorageObjectLifecyclePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ObjectstorageObjectLifecyclePolicyState, opts?: pulumi.CustomResourceOptions): ObjectstorageObjectLifecyclePolicy {
        return new ObjectstorageObjectLifecyclePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/objectstorageObjectLifecyclePolicy:ObjectstorageObjectLifecyclePolicy';

    /**
     * Returns true if the given object is an instance of ObjectstorageObjectLifecyclePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObjectstorageObjectLifecyclePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObjectstorageObjectLifecyclePolicy.__pulumiType;
    }

    /**
     * The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The Object Storage namespace used for the request.
     */
    public readonly namespace!: pulumi.Output<string>;
    /**
     * (Updatable) The bucket's set of lifecycle policy rules.
     */
    public readonly rules!: pulumi.Output<outputs.ObjectstorageObjectLifecyclePolicyRule[]>;
    /**
     * The date and time the object lifecycle policy was created, as described in [RFC 3339](https://tools.ietf.org/html/rfc3339).
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;

    /**
     * Create a ObjectstorageObjectLifecyclePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObjectstorageObjectLifecyclePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ObjectstorageObjectLifecyclePolicyArgs | ObjectstorageObjectLifecyclePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ObjectstorageObjectLifecyclePolicyState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["namespace"] = state ? state.namespace : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
        } else {
            const args = argsOrState as ObjectstorageObjectLifecyclePolicyArgs | undefined;
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            if ((!args || args.namespace === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespace'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["namespace"] = args ? args.namespace : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["timeCreated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ObjectstorageObjectLifecyclePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ObjectstorageObjectLifecyclePolicy resources.
 */
export interface ObjectstorageObjectLifecyclePolicyState {
    /**
     * The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
     */
    bucket?: pulumi.Input<string>;
    /**
     * The Object Storage namespace used for the request.
     */
    namespace?: pulumi.Input<string>;
    /**
     * (Updatable) The bucket's set of lifecycle policy rules.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ObjectstorageObjectLifecyclePolicyRule>[]>;
    /**
     * The date and time the object lifecycle policy was created, as described in [RFC 3339](https://tools.ietf.org/html/rfc3339).
     */
    timeCreated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ObjectstorageObjectLifecyclePolicy resource.
 */
export interface ObjectstorageObjectLifecyclePolicyArgs {
    /**
     * The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
     */
    bucket: pulumi.Input<string>;
    /**
     * The Object Storage namespace used for the request.
     */
    namespace: pulumi.Input<string>;
    /**
     * (Updatable) The bucket's set of lifecycle policy rules.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ObjectstorageObjectLifecyclePolicyRule>[]>;
}