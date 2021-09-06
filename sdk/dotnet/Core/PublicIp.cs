// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core
{
    /// <summary>
    /// This resource provides the Public Ip resource in Oracle Cloud Infrastructure Core service.
    /// 
    /// Creates a public IP. Use the `lifetime` property to specify whether it's an ephemeral or
    /// reserved public IP. For information about limits on how many you can create, see
    /// [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
    /// 
    /// * **For an ephemeral public IP assigned to a private IP:** You must also specify a `privateIpId`
    /// with the OCID of the primary private IP you want to assign the public IP to. The public IP is
    /// created in the same availability domain as the private IP. An ephemeral public IP must always be
    /// assigned to a private IP, and only to the *primary* private IP on a VNIC, not a secondary
    /// private IP. Exception: If you create a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/), Oracle
    /// automatically assigns the NAT gateway a regional ephemeral public IP that you cannot remove.
    /// 
    /// * **For a reserved public IP:** You may also optionally assign the public IP to a private
    /// IP by specifying `privateIpId`. Or you can later assign the public IP with
    /// [UpdatePublicIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/UpdatePublicIp).
    /// 
    /// **Note:** When assigning a public IP to a private IP, the private IP must not already have
    /// a public IP with `lifecycleState` = ASSIGNING or ASSIGNED. If it does, an error is returned.
    /// 
    /// Also, for reserved public IPs, the optional assignment part of this operation is
    /// asynchronous. Poll the public IP's `lifecycleState` to determine if the assignment
    /// succeeded.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Oci = Pulumi.Oci;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testPublicIp = new Oci.Core.PublicIp("testPublicIp", new Oci.Core.PublicIpArgs
    ///         {
    ///             CompartmentId = @var.Compartment_id,
    ///             Lifetime = @var.Public_ip_lifetime,
    ///             DefinedTags = 
    ///             {
    ///                 { "Operations.CostCenter", "42" },
    ///             },
    ///             DisplayName = @var.Public_ip_display_name,
    ///             FreeformTags = 
    ///             {
    ///                 { "Department", "Finance" },
    ///             },
    ///             PrivateIpId = oci_core_private_ip.Test_private_ip.Id,
    ///             PublicIpPoolId = oci_core_public_ip_pool.Test_public_ip_pool.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// PublicIps can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:core/publicIp:PublicIp test_public_ip "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:core/publicIp:PublicIp")]
    public partial class PublicIp : Pulumi.CustomResource
    {
        /// <summary>
        /// The OCID of the entity the public IP is assigned to, or in the process of being assigned to.
        /// </summary>
        [Output("assignedEntityId")]
        public Output<string> AssignedEntityId { get; private set; } = null!;

        /// <summary>
        /// The type of entity the public IP is assigned to, or in the process of being assigned to.
        /// </summary>
        [Output("assignedEntityType")]
        public Output<string> AssignedEntityType { get; private set; } = null!;

        /// <summary>
        /// The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the availability domain of the assigned private IP.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        [Output("availabilityDomain")]
        public Output<string> AvailabilityDomain { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// The public IP address of the `publicIp` object.  Example: `203.0.113.2`
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
        /// </summary>
        [Output("lifetime")]
        public Output<string> Lifetime { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the private IP to assign the public IP to.
        /// </summary>
        [Output("privateIpId")]
        public Output<string?> PrivateIpId { get; private set; } = null!;

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
        /// </summary>
        [Output("publicIpPoolId")]
        public Output<string> PublicIpPoolId { get; private set; } = null!;

        /// <summary>
        /// Whether the public IP is regional or specific to a particular availability domain.
        /// * `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs and ephemeral public IPs assigned to a regional entity have `scope` = `REGION`.
        /// * `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// The public IP's current state.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The date and time the public IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;


        /// <summary>
        /// Create a PublicIp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicIp(string name, PublicIpArgs args, CustomResourceOptions? options = null)
            : base("oci:core/publicIp:PublicIp", name, args ?? new PublicIpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicIp(string name, Input<string> id, PublicIpState? state = null, CustomResourceOptions? options = null)
            : base("oci:core/publicIp:PublicIp", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PublicIp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicIp Get(string name, Input<string> id, PublicIpState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicIp(name, id, state, options);
        }
    }

    public sealed class PublicIpArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
        /// </summary>
        [Input("compartmentId", required: true)]
        public Input<string> CompartmentId { get; set; } = null!;

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
        /// </summary>
        [Input("lifetime", required: true)]
        public Input<string> Lifetime { get; set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the private IP to assign the public IP to.
        /// </summary>
        [Input("privateIpId")]
        public Input<string>? PrivateIpId { get; set; }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
        /// </summary>
        [Input("publicIpPoolId")]
        public Input<string>? PublicIpPoolId { get; set; }

        public PublicIpArgs()
        {
        }
    }

    public sealed class PublicIpState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The OCID of the entity the public IP is assigned to, or in the process of being assigned to.
        /// </summary>
        [Input("assignedEntityId")]
        public Input<string>? AssignedEntityId { get; set; }

        /// <summary>
        /// The type of entity the public IP is assigned to, or in the process of being assigned to.
        /// </summary>
        [Input("assignedEntityType")]
        public Input<string>? AssignedEntityType { get; set; }

        /// <summary>
        /// The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the availability domain of the assigned private IP.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        [Input("availabilityDomain")]
        public Input<string>? AvailabilityDomain { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment OCID.
        /// </summary>
        [Input("compartmentId")]
        public Input<string>? CompartmentId { get; set; }

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        /// <summary>
        /// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// The public IP address of the `publicIp` object.  Example: `203.0.113.2`
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
        /// </summary>
        [Input("lifetime")]
        public Input<string>? Lifetime { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the private IP to assign the public IP to.
        /// </summary>
        [Input("privateIpId")]
        public Input<string>? PrivateIpId { get; set; }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
        /// </summary>
        [Input("publicIpPoolId")]
        public Input<string>? PublicIpPoolId { get; set; }

        /// <summary>
        /// Whether the public IP is regional or specific to a particular availability domain.
        /// * `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs and ephemeral public IPs assigned to a regional entity have `scope` = `REGION`.
        /// * `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The public IP's current state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The date and time the public IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        public PublicIpState()
        {
        }
    }
}