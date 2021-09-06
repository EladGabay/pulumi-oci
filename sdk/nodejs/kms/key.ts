// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource provides the Key resource in Oracle Cloud Infrastructure Kms service.
 *
 * Creates a new master encryption key.
 *
 * As a management operation, this call is subject to a Key Management limit that applies to the total
 * number of requests across all management write operations. Key Management might throttle this call
 * to reject an otherwise valid request when the total rate of management write operations exceeds 10
 * requests per second for a given tenancy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testKey = new oci.kms.Key("testKey", {
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.key_display_name,
 *     keyShape: {
 *         algorithm: _var.key_key_shape_algorithm,
 *         length: _var.key_key_shape_length,
 *         curveId: oci_kms_curve.test_curve.id,
 *     },
 *     managementEndpoint: _var.key_management_endpoint,
 *     definedTags: {
 *         "Operations.CostCenter": "42",
 *     },
 *     freeformTags: {
 *         Department: "Finance",
 *     },
 *     protectionMode: _var.key_protection_mode,
 * });
 * ```
 *
 * ## Import
 *
 * Keys can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:kms/key:Key test_key "managementEndpoint/{managementEndpoint}/keys/{keyId}"
 * ```
 */
export class Key extends pulumi.CustomResource {
    /**
     * Get an existing Key resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeyState, opts?: pulumi.CustomResourceOptions): Key {
        return new Key(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:kms/key:Key';

    /**
     * Returns true if the given object is an instance of Key.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Key {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Key.__pulumiType;
    }

    /**
     * (Updatable) The OCID of the compartment where you want to create the master encryption key.
     */
    public readonly compartmentId!: pulumi.Output<string>;
    /**
     * The OCID of the key version used in cryptographic operations. During key rotation, the service might be in a transitional state where this or a newer key version are used intermittently. The `currentKeyVersion` property is updated when the service is guaranteed to use the new key version for all subsequent encryption operations.
     */
    public /*out*/ readonly currentKeyVersion!: pulumi.Output<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    public readonly definedTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * (Updatable) Desired state of the key. Possible values : `ENABLED` or `DISABLED`
     */
    public readonly desiredState!: pulumi.Output<string>;
    /**
     * (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    public readonly freeformTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * A boolean that will be true when key is primary, and will be false when key is a replica from a primary key.
     */
    public /*out*/ readonly isPrimary!: pulumi.Output<boolean>;
    /**
     * The cryptographic properties of a key.
     */
    public readonly keyShape!: pulumi.Output<outputs.kms.KeyKeyShape>;
    /**
     * The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
     */
    public readonly managementEndpoint!: pulumi.Output<string>;
    /**
     * The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists  on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,  a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported.
     */
    public readonly protectionMode!: pulumi.Output<string>;
    /**
     * Key replica details
     */
    public /*out*/ readonly replicaDetails!: pulumi.Output<outputs.kms.KeyReplicaDetails>;
    /**
     * (Updatable) Details where key was backed up.
     */
    public readonly restoreFromFile!: pulumi.Output<outputs.kms.KeyRestoreFromFile | undefined>;
    /**
     * (Updatable) Details where key was backed up
     */
    public readonly restoreFromObjectStore!: pulumi.Output<outputs.kms.KeyRestoreFromObjectStore | undefined>;
    /**
     * (Updatable) An optional property when flipped triggers restore from restore option provided in config file.
     */
    public readonly restoreTrigger!: pulumi.Output<boolean | undefined>;
    /**
     * The OCID of the key from which this key was restored.
     */
    public /*out*/ readonly restoredFromKeyId!: pulumi.Output<string>;
    /**
     * The key's current lifecycle state.  Example: `ENABLED`
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z`
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * (Updatable) An optional property for the deletion time of the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`
     */
    public readonly timeOfDeletion!: pulumi.Output<string>;
    /**
     * The OCID of the vault that contains this key.
     */
    public /*out*/ readonly vaultId!: pulumi.Output<string>;

    /**
     * Create a Key resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeyArgs | KeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeyState | undefined;
            inputs["compartmentId"] = state ? state.compartmentId : undefined;
            inputs["currentKeyVersion"] = state ? state.currentKeyVersion : undefined;
            inputs["definedTags"] = state ? state.definedTags : undefined;
            inputs["desiredState"] = state ? state.desiredState : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["freeformTags"] = state ? state.freeformTags : undefined;
            inputs["isPrimary"] = state ? state.isPrimary : undefined;
            inputs["keyShape"] = state ? state.keyShape : undefined;
            inputs["managementEndpoint"] = state ? state.managementEndpoint : undefined;
            inputs["protectionMode"] = state ? state.protectionMode : undefined;
            inputs["replicaDetails"] = state ? state.replicaDetails : undefined;
            inputs["restoreFromFile"] = state ? state.restoreFromFile : undefined;
            inputs["restoreFromObjectStore"] = state ? state.restoreFromObjectStore : undefined;
            inputs["restoreTrigger"] = state ? state.restoreTrigger : undefined;
            inputs["restoredFromKeyId"] = state ? state.restoredFromKeyId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
            inputs["timeOfDeletion"] = state ? state.timeOfDeletion : undefined;
            inputs["vaultId"] = state ? state.vaultId : undefined;
        } else {
            const args = argsOrState as KeyArgs | undefined;
            if ((!args || args.compartmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compartmentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.keyShape === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyShape'");
            }
            if ((!args || args.managementEndpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'managementEndpoint'");
            }
            inputs["compartmentId"] = args ? args.compartmentId : undefined;
            inputs["definedTags"] = args ? args.definedTags : undefined;
            inputs["desiredState"] = args ? args.desiredState : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["freeformTags"] = args ? args.freeformTags : undefined;
            inputs["keyShape"] = args ? args.keyShape : undefined;
            inputs["managementEndpoint"] = args ? args.managementEndpoint : undefined;
            inputs["protectionMode"] = args ? args.protectionMode : undefined;
            inputs["restoreFromFile"] = args ? args.restoreFromFile : undefined;
            inputs["restoreFromObjectStore"] = args ? args.restoreFromObjectStore : undefined;
            inputs["restoreTrigger"] = args ? args.restoreTrigger : undefined;
            inputs["timeOfDeletion"] = args ? args.timeOfDeletion : undefined;
            inputs["currentKeyVersion"] = undefined /*out*/;
            inputs["isPrimary"] = undefined /*out*/;
            inputs["replicaDetails"] = undefined /*out*/;
            inputs["restoredFromKeyId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["vaultId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Key.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Key resources.
 */
export interface KeyState {
    /**
     * (Updatable) The OCID of the compartment where you want to create the master encryption key.
     */
    compartmentId?: pulumi.Input<string>;
    /**
     * The OCID of the key version used in cryptographic operations. During key rotation, the service might be in a transitional state where this or a newer key version are used intermittently. The `currentKeyVersion` property is updated when the service is guaranteed to use the new key version for all subsequent encryption operations.
     */
    currentKeyVersion?: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) Desired state of the key. Possible values : `ENABLED` or `DISABLED`
     */
    desiredState?: pulumi.Input<string>;
    /**
     * (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information.
     */
    displayName?: pulumi.Input<string>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A boolean that will be true when key is primary, and will be false when key is a replica from a primary key.
     */
    isPrimary?: pulumi.Input<boolean>;
    /**
     * The cryptographic properties of a key.
     */
    keyShape?: pulumi.Input<inputs.kms.KeyKeyShape>;
    /**
     * The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
     */
    managementEndpoint?: pulumi.Input<string>;
    /**
     * The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists  on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,  a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported.
     */
    protectionMode?: pulumi.Input<string>;
    /**
     * Key replica details
     */
    replicaDetails?: pulumi.Input<inputs.kms.KeyReplicaDetails>;
    /**
     * (Updatable) Details where key was backed up.
     */
    restoreFromFile?: pulumi.Input<inputs.kms.KeyRestoreFromFile>;
    /**
     * (Updatable) Details where key was backed up
     */
    restoreFromObjectStore?: pulumi.Input<inputs.kms.KeyRestoreFromObjectStore>;
    /**
     * (Updatable) An optional property when flipped triggers restore from restore option provided in config file.
     */
    restoreTrigger?: pulumi.Input<boolean>;
    /**
     * The OCID of the key from which this key was restored.
     */
    restoredFromKeyId?: pulumi.Input<string>;
    /**
     * The key's current lifecycle state.  Example: `ENABLED`
     */
    state?: pulumi.Input<string>;
    /**
     * The date and time the key was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.  Example: `2018-04-03T21:10:29.600Z`
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * (Updatable) An optional property for the deletion time of the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`
     */
    timeOfDeletion?: pulumi.Input<string>;
    /**
     * The OCID of the vault that contains this key.
     */
    vaultId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Key resource.
 */
export interface KeyArgs {
    /**
     * (Updatable) The OCID of the compartment where you want to create the master encryption key.
     */
    compartmentId: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) Desired state of the key. Possible values : `ENABLED` or `DISABLED`
     */
    desiredState?: pulumi.Input<string>;
    /**
     * (Updatable) A user-friendly name for the key. It does not have to be unique, and it is changeable. Avoid entering confidential information.
     */
    displayName: pulumi.Input<string>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The cryptographic properties of a key.
     */
    keyShape: pulumi.Input<inputs.kms.KeyKeyShape>;
    /**
     * The service endpoint to perform management operations against. Management operations include 'Create,' 'Update,' 'List,' 'Get,' and 'Delete' operations. See Vault Management endpoint.
     */
    managementEndpoint: pulumi.Input<string>;
    /**
     * The key's protection mode indicates how the key persists and where cryptographic operations that use the key are performed. A protection mode of `HSM` means that the key persists on a hardware security module (HSM) and all cryptographic operations are performed inside the HSM. A protection mode of `SOFTWARE` means that the key persists on the server, protected by the vault's RSA wrapping key which persists  on the HSM. All cryptographic operations that use a key with a protection mode of `SOFTWARE` are performed on the server. By default,  a key's protection mode is set to `HSM`. You can't change a key's protection mode after the key is created or imported.
     */
    protectionMode?: pulumi.Input<string>;
    /**
     * (Updatable) Details where key was backed up.
     */
    restoreFromFile?: pulumi.Input<inputs.kms.KeyRestoreFromFile>;
    /**
     * (Updatable) Details where key was backed up
     */
    restoreFromObjectStore?: pulumi.Input<inputs.kms.KeyRestoreFromObjectStore>;
    /**
     * (Updatable) An optional property when flipped triggers restore from restore option provided in config file.
     */
    restoreTrigger?: pulumi.Input<boolean>;
    /**
     * (Updatable) An optional property for the deletion time of the key, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z`
     */
    timeOfDeletion?: pulumi.Input<string>;
}