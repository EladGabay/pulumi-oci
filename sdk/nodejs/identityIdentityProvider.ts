// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource provides the Identity Provider resource in Oracle Cloud Infrastructure Identity service.
 *
 * Creates a new identity provider in your tenancy. For more information, see
 * [Identity Providers and Federation](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/federation.htm).
 *
 * You must specify your tenancy's OCID as the compartment ID in the request object.
 * Remember that the tenancy is simply the root compartment. For information about
 * OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
 *
 * You must also specify a *name* for the `IdentityProvider`, which must be unique
 * across all `IdentityProvider` objects in your tenancy and cannot be changed.
 *
 * You must also specify a *description* for the `IdentityProvider` (although
 * it can be an empty string). It does not have to be unique, and you can change
 * it anytime with
 * [UpdateIdentityProvider](https://docs.cloud.oracle.com/iaas/api/#/en/identity/20160918/IdentityProvider/UpdateIdentityProvider).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testIdentityProvider = new oci.IdentityIdentityProvider("testIdentityProvider", {
 *     compartmentId: _var.tenancy_ocid,
 *     description: _var.identity_provider_description,
 *     metadata: _var.identity_provider_metadata,
 *     metadataUrl: _var.identity_provider_metadata_url,
 *     productType: _var.identity_provider_product_type,
 *     protocol: _var.identity_provider_protocol,
 *     definedTags: {
 *         "Operations.CostCenter": "42",
 *     },
 *     freeformAttributes: _var.identity_provider_freeform_attributes,
 *     freeformTags: {
 *         Department: "Finance",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * IdentityProviders can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import oci:index/identityIdentityProvider:IdentityIdentityProvider test_identity_provider "id"
 * ```
 */
export class IdentityIdentityProvider extends pulumi.CustomResource {
    /**
     * Get an existing IdentityIdentityProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityIdentityProviderState, opts?: pulumi.CustomResourceOptions): IdentityIdentityProvider {
        return new IdentityIdentityProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'oci:index/identityIdentityProvider:IdentityIdentityProvider';

    /**
     * Returns true if the given object is an instance of IdentityIdentityProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityIdentityProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityIdentityProvider.__pulumiType;
    }

    /**
     * The OCID of your tenancy.
     */
    public readonly compartmentId!: pulumi.Output<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    public readonly definedTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
     */
    public readonly freeformAttributes!: pulumi.Output<{[key: string]: any}>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    public readonly freeformTags!: pulumi.Output<{[key: string]: any}>;
    /**
     * The detailed status of INACTIVE lifecycleState.
     */
    public /*out*/ readonly inactiveState!: pulumi.Output<string>;
    /**
     * (Updatable) The XML that contains the information required for federating.
     */
    public readonly metadata!: pulumi.Output<string>;
    /**
     * (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
     */
    public readonly metadataUrl!: pulumi.Output<string>;
    /**
     * The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
     */
    public readonly productType!: pulumi.Output<string>;
    /**
     * (Updatable) The protocol used for federation.  Example: `SAML2`
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The URL to redirect federated users to for authentication with the identity provider.
     */
    public /*out*/ readonly redirectUrl!: pulumi.Output<string>;
    /**
     * The identity provider's signing certificate used by the IAM Service to validate the SAML2 token.
     */
    public /*out*/ readonly signingCertificate!: pulumi.Output<string>;
    /**
     * The current state.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
     */
    public /*out*/ readonly timeCreated!: pulumi.Output<string>;

    /**
     * Create a IdentityIdentityProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityIdentityProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityIdentityProviderArgs | IdentityIdentityProviderState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentityIdentityProviderState | undefined;
            inputs["compartmentId"] = state ? state.compartmentId : undefined;
            inputs["definedTags"] = state ? state.definedTags : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["freeformAttributes"] = state ? state.freeformAttributes : undefined;
            inputs["freeformTags"] = state ? state.freeformTags : undefined;
            inputs["inactiveState"] = state ? state.inactiveState : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["metadataUrl"] = state ? state.metadataUrl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["productType"] = state ? state.productType : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["redirectUrl"] = state ? state.redirectUrl : undefined;
            inputs["signingCertificate"] = state ? state.signingCertificate : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["timeCreated"] = state ? state.timeCreated : undefined;
        } else {
            const args = argsOrState as IdentityIdentityProviderArgs | undefined;
            if ((!args || args.compartmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'compartmentId'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.metadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadata'");
            }
            if ((!args || args.metadataUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadataUrl'");
            }
            if ((!args || args.productType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productType'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            inputs["compartmentId"] = args ? args.compartmentId : undefined;
            inputs["definedTags"] = args ? args.definedTags : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["freeformAttributes"] = args ? args.freeformAttributes : undefined;
            inputs["freeformTags"] = args ? args.freeformTags : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["metadataUrl"] = args ? args.metadataUrl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["productType"] = args ? args.productType : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["inactiveState"] = undefined /*out*/;
            inputs["redirectUrl"] = undefined /*out*/;
            inputs["signingCertificate"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["timeCreated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(IdentityIdentityProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityIdentityProvider resources.
 */
export interface IdentityIdentityProviderState {
    /**
     * The OCID of your tenancy.
     */
    compartmentId?: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
     */
    description?: pulumi.Input<string>;
    /**
     * (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
     */
    freeformAttributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The detailed status of INACTIVE lifecycleState.
     */
    inactiveState?: pulumi.Input<string>;
    /**
     * (Updatable) The XML that contains the information required for federating.
     */
    metadata?: pulumi.Input<string>;
    /**
     * (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
     */
    metadataUrl?: pulumi.Input<string>;
    /**
     * The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
     */
    productType?: pulumi.Input<string>;
    /**
     * (Updatable) The protocol used for federation.  Example: `SAML2`
     */
    protocol?: pulumi.Input<string>;
    /**
     * The URL to redirect federated users to for authentication with the identity provider.
     */
    redirectUrl?: pulumi.Input<string>;
    /**
     * The identity provider's signing certificate used by the IAM Service to validate the SAML2 token.
     */
    signingCertificate?: pulumi.Input<string>;
    /**
     * The current state.
     */
    state?: pulumi.Input<string>;
    /**
     * Date and time the `IdentityProvider` was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
     */
    timeCreated?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityIdentityProvider resource.
 */
export interface IdentityIdentityProviderArgs {
    /**
     * The OCID of your tenancy.
     */
    compartmentId: pulumi.Input<string>;
    /**
     * (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
     */
    definedTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) The description you assign to the `IdentityProvider` during creation. Does not have to be unique, and it's changeable.
     */
    description: pulumi.Input<string>;
    /**
     * (Updatable) Extra name value pairs associated with this identity provider. Example: `{"clientId": "appSf3kdjf3"}`
     */
    freeformAttributes?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
     */
    freeformTags?: pulumi.Input<{[key: string]: any}>;
    /**
     * (Updatable) The XML that contains the information required for federating.
     */
    metadata: pulumi.Input<string>;
    /**
     * (Updatable) The URL for retrieving the identity provider's metadata, which contains information required for federating.
     */
    metadataUrl: pulumi.Input<string>;
    /**
     * The name you assign to the `IdentityProvider` during creation. The name must be unique across all `IdentityProvider` objects in the tenancy and cannot be changed.
     */
    name?: pulumi.Input<string>;
    /**
     * The identity provider service or product. Supported identity providers are Oracle Identity Cloud Service (IDCS) and Microsoft Active Directory Federation Services (ADFS).  Example: `IDCS`
     */
    productType: pulumi.Input<string>;
    /**
     * (Updatable) The protocol used for federation.  Example: `SAML2`
     */
    protocol: pulumi.Input<string>;
}