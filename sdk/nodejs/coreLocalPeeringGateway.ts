// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource provides the Local Peering Gateway resource in Oracle Cloud Infrastructure Core service.
 *
 * Creates a new local peering gateway (LPG) for the specified VCN.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testLocalPeeringGateway = new oci.CoreLocalPeeringGateway("testLocalPeeringGateway", {
 *     compartmentId: _var.compartment_id,
 *     vcnId: oci_core_vcn.test_vcn.id,
 *     definedTags: {
 *         "Operations.CostCenter": "42",
 *     },
 *     displayName: _var.local_peering_gateway_display_name,
 *     freeformTags: {
 *         Department: "Finance",
 *     },
 *     peerId: oci_core_local_peering_gateway.test_local_peering_gateway2.id,
 *     routeTableId: oci_core_route_table.test_route_table.id,
 * });
 * ```
 *
 * ## Import
 *
 * LocalPeeringGateways can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:index/coreLocalPeeringGateway:CoreLocalPeeringGateway test_local_peering_gateway "id"
 * ```
 */
export class CoreLocalPeeringGateway extends pulumi.CustomResource {
    /**
     * Get an existing CoreLocalPeeringGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CoreLocalPeeringGatewayState, opts?: pulumi.CustomResourceOptions): CoreLocalPeeringGateway {
        return new CoreLocalPeeringGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/coreLocalPeeringGateway:CoreLocalPeeringGateway';

    /**
     * Returns true if the given object is an instance of CoreLocalPeeringGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CoreLocalPeeringGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CoreLocalPeeringGateway.__pulumiType;
    }

    /**
     * (Updatable) The OCID of the compartment containing the local peering gateway (LPG).
     */
    public readonly compartmentId!: pulumi.Output<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
     */
    public readonly definedTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
     */
    public readonly freeformTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false`
     */
    public /*out*/ readonly isCrossTenancyPeering!: pulumi.Output<boolean>;
    /**
     * The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN at the other end of the peering from this LPG. See `peerAdvertisedCidrDetails` for the individual CIDRs. The value is `null` if the LPG is not peered.  Example: `192.168.0.0/16`, or if aggregated with `172.16.0.0/24` then `128.0.0.0/1`
     */
    public /*out*/ readonly peerAdvertisedCidr!: pulumi.Output<string>;
    /**
     * The specific ranges of IP addresses available on or via the VCN at the other end of the peering from this LPG. The value is `null` if the LPG is not peered. You can use these as destination CIDRs for route rules to route a subnet's traffic to this LPG.  Example: [`192.168.0.0/16`, `172.16.0.0/24`]
     */
    public /*out*/ readonly peerAdvertisedCidrDetails!: pulumi.Output<string[]>;
    /**
     * The OCID of the LPG you want to peer with. Specifying a peerId connects this local peering gateway (LPG) to another one in the same region. This operation must be called by the VCN administrator who is designated as the *requestor* in the peering relationship. The *acceptor* must implement an Identity and Access Management (IAM) policy that gives the requestor permission to connect to LPGs in the acceptor's compartment. Without that permission, this operation will fail. For more information, see [VCN Peering](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/VCNpeering.htm).
     */
    public readonly peerId!: pulumi.Output<string>;
    /**
     * Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the LPG at the other end of the peering has been deleted.
     */
    public /*out*/ readonly peeringStatus!: pulumi.Output<string>;
    /**
     * Additional information regarding the peering status, if applicable.
     */
    public /*out*/ readonly peeringStatusDetails!: pulumi.Output<string>;
    /**
     * (Updatable) The OCID of the route table the LPG will use.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The LPG's current lifecycle state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The date and time the LPG was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;
    /**
     * The OCID of the VCN the LPG belongs to.
     */
    public readonly vcnId!: pulumi.Output<string>;

    /**
     * Create a CoreLocalPeeringGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CoreLocalPeeringGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CoreLocalPeeringGatewayArgs | CoreLocalPeeringGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CoreLocalPeeringGatewayState | undefined;
            inputs["compartmentId"] = state ? state.compartmentId : undefined;
            inputs["definedTags"] = state ? state.definedTags : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["freeformTags"] = state ? state.freeformTags : undefined;
            inputs["isCrossTenancyPeering"] = state ? state.isCrossTenancyPeering : undefined;
            inputs["peerAdvertisedCidr"] = state ? state.peerAdvertisedCidr : undefined;
            inputs["peerAdvertisedCidrDetails"] = state ? state.peerAdvertisedCidrDetails : undefined;
            inputs["peerId"] = state ? state.peerId : undefined;
            inputs["peeringStatus"] = state ? state.peeringStatus : undefined;
            inputs["peeringStatusDetails"] = state ? state.peeringStatusDetails : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
            inputs["vcnId"] = state ? state.vcnId : undefined;
        } else {
            const args = argsOrState as CoreLocalPeeringGatewayArgs | undefined;
            if ((!args || args.compartmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compartmentId'");
            }
            if ((!args || args.vcnId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vcnId'");
            }
            inputs["compartmentId"] = args ? args.compartmentId : undefined;
            inputs["definedTags"] = args ? args.definedTags : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["freeformTags"] = args ? args.freeformTags : undefined;
            inputs["peerId"] = args ? args.peerId : undefined;
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
            inputs["vcnId"] = args ? args.vcnId : undefined;
            inputs["isCrossTenancyPeering"] = undefined /*out*/;
            inputs["peerAdvertisedCidr"] = undefined /*out*/;
            inputs["peerAdvertisedCidrDetails"] = undefined /*out*/;
            inputs["peeringStatus"] = undefined /*out*/;
            inputs["peeringStatusDetails"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CoreLocalPeeringGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CoreLocalPeeringGateway resources.
 */
export interface CoreLocalPeeringGatewayState {
    /**
     * (Updatable) The OCID of the compartment containing the local peering gateway (LPG).
     */
    compartmentId?: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
     */
    displayName?: pulumi.Input<string>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the VCN at the other end of the peering is in a different tenancy.  Example: `false`
     */
    isCrossTenancyPeering?: pulumi.Input<boolean>;
    /**
     * The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN at the other end of the peering from this LPG. See `peerAdvertisedCidrDetails` for the individual CIDRs. The value is `null` if the LPG is not peered.  Example: `192.168.0.0/16`, or if aggregated with `172.16.0.0/24` then `128.0.0.0/1`
     */
    peerAdvertisedCidr?: pulumi.Input<string>;
    /**
     * The specific ranges of IP addresses available on or via the VCN at the other end of the peering from this LPG. The value is `null` if the LPG is not peered. You can use these as destination CIDRs for route rules to route a subnet's traffic to this LPG.  Example: [`192.168.0.0/16`, `172.16.0.0/24`]
     */
    peerAdvertisedCidrDetails?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The OCID of the LPG you want to peer with. Specifying a peerId connects this local peering gateway (LPG) to another one in the same region. This operation must be called by the VCN administrator who is designated as the *requestor* in the peering relationship. The *acceptor* must implement an Identity and Access Management (IAM) policy that gives the requestor permission to connect to LPGs in the acceptor's compartment. Without that permission, this operation will fail. For more information, see [VCN Peering](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/VCNpeering.htm).
     */
    peerId?: pulumi.Input<string>;
    /**
     * Whether the LPG is peered with another LPG. `NEW` means the LPG has not yet been peered. `PENDING` means the peering is being established. `REVOKED` means the LPG at the other end of the peering has been deleted.
     */
    peeringStatus?: pulumi.Input<string>;
    /**
     * Additional information regarding the peering status, if applicable.
     */
    peeringStatusDetails?: pulumi.Input<string>;
    /**
     * (Updatable) The OCID of the route table the LPG will use.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The LPG's current lifecycle state.
     */
    state?: pulumi.Input<string>;
    /**
     * The date and time the LPG was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * The OCID of the VCN the LPG belongs to.
     */
    vcnId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CoreLocalPeeringGateway resource.
 */
export interface CoreLocalPeeringGatewayArgs {
    /**
     * (Updatable) The OCID of the compartment containing the local peering gateway (LPG).
     */
    compartmentId: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
     */
    displayName?: pulumi.Input<string>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The OCID of the LPG you want to peer with. Specifying a peerId connects this local peering gateway (LPG) to another one in the same region. This operation must be called by the VCN administrator who is designated as the *requestor* in the peering relationship. The *acceptor* must implement an Identity and Access Management (IAM) policy that gives the requestor permission to connect to LPGs in the acceptor's compartment. Without that permission, this operation will fail. For more information, see [VCN Peering](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/VCNpeering.htm).
     */
    peerId?: pulumi.Input<string>;
    /**
     * (Updatable) The OCID of the route table the LPG will use.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The OCID of the VCN the LPG belongs to.
     */
    vcnId: pulumi.Input<string>;
}