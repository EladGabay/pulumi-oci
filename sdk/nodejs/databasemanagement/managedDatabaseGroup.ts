// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource provides the Managed Database Group resource in Oracle Cloud Infrastructure Database Management service.
 *
 * Creates a Managed Database Group. The group does not contain any
 * Managed Databases when it is created, and they must be added later.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testManagedDatabaseGroup = new oci.databasemanagement.ManagedDatabaseGroup("testManagedDatabaseGroup", {
 *     compartmentId: _var.compartment_id,
 *     description: _var.managed_database_group_description,
 *     managedDatabases: [{
 *         id: _var.managed_database_id,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ManagedDatabaseGroups can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:databasemanagement/managedDatabaseGroup:ManagedDatabaseGroup test_managed_database_group "id"
 * ```
 */
export class ManagedDatabaseGroup extends pulumi.CustomResource {
    /**
     * Get an existing ManagedDatabaseGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedDatabaseGroupState, opts?: pulumi.CustomResourceOptions): ManagedDatabaseGroup {
        return new ManagedDatabaseGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:databasemanagement/managedDatabaseGroup:ManagedDatabaseGroup';

    /**
     * Returns true if the given object is an instance of ManagedDatabaseGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedDatabaseGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedDatabaseGroup.__pulumiType;
    }

    /**
     * (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides.
     */
    public readonly compartmentId!: pulumi.Output<string>;
    /**
     * (Updatable) The information specified by the user about the Managed Database Group.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * (Updatable) Set of Managed Databases that the user wants to add to the Managed Database Group. Specifying a block will add the Managed Database to Managed Database Group and removing the block will remove Managed Database from the Managed Database Group.
     */
    public readonly managedDatabases!: pulumi.Output<outputs.databasemanagement.ManagedDatabaseGroupManagedDatabase[]>;
    /**
     * The name of the Managed Database Group. Valid characters are uppercase or lowercase letters, numbers, and "_". The name of the Managed Database Group cannot be modified. It must be unique in the compartment and must begin with an alphabetic character.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current lifecycle state of the Managed Database Group.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The date and time the Managed Database Group was created.
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * The date and time the Managed Database Group was last updated.
     */
    public /*out*/ readonly timeUpdated!: pulumi.Output<string>;

    /**
     * Create a ManagedDatabaseGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedDatabaseGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedDatabaseGroupArgs | ManagedDatabaseGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ManagedDatabaseGroupState | undefined;
            inputs["compartmentId"] = state ? state.compartmentId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["managedDatabases"] = state ? state.managedDatabases : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
            inputs["timeUpdated"] = state ? state.timeUpdated : undefined;
        } else {
            const args = argsOrState as ManagedDatabaseGroupArgs | undefined;
            if ((!args || args.compartmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compartmentId'");
            }
            inputs["compartmentId"] = args ? args.compartmentId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["managedDatabases"] = args ? args.managedDatabases : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["state"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
            inputs["timeUpdated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ManagedDatabaseGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedDatabaseGroup resources.
 */
export interface ManagedDatabaseGroupState {
    /**
     * (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides.
     */
    compartmentId?: pulumi.Input<string>;
    /**
     * (Updatable) The information specified by the user about the Managed Database Group.
     */
    description?: pulumi.Input<string>;
    /**
     * (Updatable) Set of Managed Databases that the user wants to add to the Managed Database Group. Specifying a block will add the Managed Database to Managed Database Group and removing the block will remove Managed Database from the Managed Database Group.
     */
    managedDatabases?: pulumi.Input<pulumi.Input<inputs.databasemanagement.ManagedDatabaseGroupManagedDatabase>[]>;
    /**
     * The name of the Managed Database Group. Valid characters are uppercase or lowercase letters, numbers, and "_". The name of the Managed Database Group cannot be modified. It must be unique in the compartment and must begin with an alphabetic character.
     */
    name?: pulumi.Input<string>;
    /**
     * The current lifecycle state of the Managed Database Group.
     */
    state?: pulumi.Input<string>;
    /**
     * The date and time the Managed Database Group was created.
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * The date and time the Managed Database Group was last updated.
     */
    timeUpdated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ManagedDatabaseGroup resource.
 */
export interface ManagedDatabaseGroupArgs {
    /**
     * (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides.
     */
    compartmentId: pulumi.Input<string>;
    /**
     * (Updatable) The information specified by the user about the Managed Database Group.
     */
    description?: pulumi.Input<string>;
    /**
     * (Updatable) Set of Managed Databases that the user wants to add to the Managed Database Group. Specifying a block will add the Managed Database to Managed Database Group and removing the block will remove Managed Database from the Managed Database Group.
     */
    managedDatabases?: pulumi.Input<pulumi.Input<inputs.databasemanagement.ManagedDatabaseGroupManagedDatabase>[]>;
    /**
     * The name of the Managed Database Group. Valid characters are uppercase or lowercase letters, numbers, and "_". The name of the Managed Database Group cannot be modified. It must be unique in the compartment and must begin with an alphabetic character.
     */
    name?: pulumi.Input<string>;
}