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
    /// This resource provides the Network Security Group Security Rule resource in Oracle Cloud Infrastructure Core service.
    /// 
    /// Adds a security rule to the specified network security group.
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
    ///         var testNetworkSecurityGroupSecurityRule = new Oci.Core.NetworkSecurityGroupSecurityRule("testNetworkSecurityGroupSecurityRule", new Oci.Core.NetworkSecurityGroupSecurityRuleArgs
    ///         {
    ///             NetworkSecurityGroupId = oci_core_network_security_group.Test_network_security_group.Id,
    ///             Direction = @var.Network_security_group_security_rule_direction,
    ///             Protocol = @var.Network_security_group_security_rule_protocol,
    ///             Description = @var.Network_security_group_security_rule_description,
    ///             Destination = @var.Network_security_group_security_rule_destination,
    ///             DestinationType = @var.Network_security_group_security_rule_destination_type,
    ///             IcmpOptions = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleIcmpOptionsArgs
    ///             {
    ///                 Type = @var.Network_security_group_security_rule_icmp_options_type,
    ///                 Code = @var.Network_security_group_security_rule_icmp_options_code,
    ///             },
    ///             Source = @var.Network_security_group_security_rule_source,
    ///             SourceType = @var.Network_security_group_security_rule_source_type,
    ///             Stateless = @var.Network_security_group_security_rule_stateless,
    ///             TcpOptions = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleTcpOptionsArgs
    ///             {
    ///                 DestinationPortRange = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleTcpOptionsDestinationPortRangeArgs
    ///                 {
    ///                     Max = @var.Network_security_group_security_rule_tcp_options_destination_port_range_max,
    ///                     Min = @var.Network_security_group_security_rule_tcp_options_destination_port_range_min,
    ///                 },
    ///                 SourcePortRange = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleTcpOptionsSourcePortRangeArgs
    ///                 {
    ///                     Max = @var.Network_security_group_security_rule_tcp_options_source_port_range_max,
    ///                     Min = @var.Network_security_group_security_rule_tcp_options_source_port_range_min,
    ///                 },
    ///             },
    ///             UdpOptions = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleUdpOptionsArgs
    ///             {
    ///                 DestinationPortRange = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleUdpOptionsDestinationPortRangeArgs
    ///                 {
    ///                     Max = @var.Network_security_group_security_rule_udp_options_destination_port_range_max,
    ///                     Min = @var.Network_security_group_security_rule_udp_options_destination_port_range_min,
    ///                 },
    ///                 SourcePortRange = new Oci.Core.Inputs.NetworkSecurityGroupSecurityRuleUdpOptionsSourcePortRangeArgs
    ///                 {
    ///                     Max = @var.Network_security_group_security_rule_udp_options_source_port_range_max,
    ///                     Min = @var.Network_security_group_security_rule_udp_options_source_port_range_min,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// NetworkSecurityGroupSecurityRule can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:core/networkSecurityGroupSecurityRule:NetworkSecurityGroupSecurityRule test_network_security_group_security_rule "networkSecurityGroups/{networkSecurityGroupId}/securityRules/{securityRuleId}"
    /// ```
    /// </summary>
    [OciResourceType("oci:core/networkSecurityGroupSecurityRule:NetworkSecurityGroupSecurityRule")]
    public partial class NetworkSecurityGroupSecurityRule : Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of your choice for the rule. Avoid entering confidential information.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Type of destination for the rule. Required if `direction` = `EGRESS`.
        /// </summary>
        [Output("destinationType")]
        public Output<string> DestinationType { get; private set; } = null!;

        /// <summary>
        /// Direction of the security rule. Set to `EGRESS` for rules to allow outbound IP packets, or `INGRESS` for rules to allow inbound IP packets.
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
        /// * [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
        /// * [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)
        /// </summary>
        [Output("icmpOptions")]
        public Output<Outputs.NetworkSecurityGroupSecurityRuleIcmpOptions?> IcmpOptions { get; private set; } = null!;

        /// <summary>
        /// Whether the rule is valid. The value is `True` when the rule is first created. If the rule's `source` or `destination` is a network security group, the value changes to `False` if that network security group is deleted.
        /// </summary>
        [Output("isValid")]
        public Output<bool> IsValid { get; private set; } = null!;

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
        /// </summary>
        [Output("networkSecurityGroupId")]
        public Output<string> NetworkSecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.
        /// </summary>
        [Output("source")]
        public Output<string?> Source { get; private set; } = null!;

        /// <summary>
        /// Type of source for the rule. Required if `direction` = `INGRESS`.
        /// * `CIDR_BLOCK`: If the rule's `source` is an IP address range in CIDR notation.
        /// * `SERVICE_CIDR_BLOCK`: If the rule's `source` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic coming from a particular `Service` through a service gateway).
        /// * `NETWORK_SECURITY_GROUP`: If the rule's `source` is the OCID of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).
        /// </summary>
        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        /// <summary>
        /// A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic.
        /// </summary>
        [Output("stateless")]
        public Output<bool> Stateless { get; private set; } = null!;

        /// <summary>
        /// Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
        /// </summary>
        [Output("tcpOptions")]
        public Output<Outputs.NetworkSecurityGroupSecurityRuleTcpOptions?> TcpOptions { get; private set; } = null!;

        /// <summary>
        /// The date and time the security rule was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
        /// </summary>
        [Output("udpOptions")]
        public Output<Outputs.NetworkSecurityGroupSecurityRuleUdpOptions?> UdpOptions { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkSecurityGroupSecurityRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkSecurityGroupSecurityRule(string name, NetworkSecurityGroupSecurityRuleArgs args, CustomResourceOptions? options = null)
            : base("oci:core/networkSecurityGroupSecurityRule:NetworkSecurityGroupSecurityRule", name, args ?? new NetworkSecurityGroupSecurityRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkSecurityGroupSecurityRule(string name, Input<string> id, NetworkSecurityGroupSecurityRuleState? state = null, CustomResourceOptions? options = null)
            : base("oci:core/networkSecurityGroupSecurityRule:NetworkSecurityGroupSecurityRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkSecurityGroupSecurityRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkSecurityGroupSecurityRule Get(string name, Input<string> id, NetworkSecurityGroupSecurityRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkSecurityGroupSecurityRule(name, id, state, options);
        }
    }

    public sealed class NetworkSecurityGroupSecurityRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of your choice for the rule. Avoid entering confidential information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Type of destination for the rule. Required if `direction` = `EGRESS`.
        /// </summary>
        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        /// <summary>
        /// Direction of the security rule. Set to `EGRESS` for rules to allow outbound IP packets, or `INGRESS` for rules to allow inbound IP packets.
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
        /// * [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
        /// * [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)
        /// </summary>
        [Input("icmpOptions")]
        public Input<Inputs.NetworkSecurityGroupSecurityRuleIcmpOptionsArgs>? IcmpOptions { get; set; }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
        /// </summary>
        [Input("networkSecurityGroupId", required: true)]
        public Input<string> NetworkSecurityGroupId { get; set; } = null!;

        /// <summary>
        /// The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Type of source for the rule. Required if `direction` = `INGRESS`.
        /// * `CIDR_BLOCK`: If the rule's `source` is an IP address range in CIDR notation.
        /// * `SERVICE_CIDR_BLOCK`: If the rule's `source` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic coming from a particular `Service` through a service gateway).
        /// * `NETWORK_SECURITY_GROUP`: If the rule's `source` is the OCID of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic.
        /// </summary>
        [Input("stateless")]
        public Input<bool>? Stateless { get; set; }

        /// <summary>
        /// Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
        /// </summary>
        [Input("tcpOptions")]
        public Input<Inputs.NetworkSecurityGroupSecurityRuleTcpOptionsArgs>? TcpOptions { get; set; }

        /// <summary>
        /// Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
        /// </summary>
        [Input("udpOptions")]
        public Input<Inputs.NetworkSecurityGroupSecurityRuleUdpOptionsArgs>? UdpOptions { get; set; }

        public NetworkSecurityGroupSecurityRuleArgs()
        {
        }
    }

    public sealed class NetworkSecurityGroupSecurityRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of your choice for the rule. Avoid entering confidential information.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// Type of destination for the rule. Required if `direction` = `EGRESS`.
        /// </summary>
        [Input("destinationType")]
        public Input<string>? DestinationType { get; set; }

        /// <summary>
        /// Direction of the security rule. Set to `EGRESS` for rules to allow outbound IP packets, or `INGRESS` for rules to allow inbound IP packets.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
        /// * [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
        /// * [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)
        /// </summary>
        [Input("icmpOptions")]
        public Input<Inputs.NetworkSecurityGroupSecurityRuleIcmpOptionsGetArgs>? IcmpOptions { get; set; }

        /// <summary>
        /// Whether the rule is valid. The value is `True` when the rule is first created. If the rule's `source` or `destination` is a network security group, the value changes to `False` if that network security group is deleted.
        /// </summary>
        [Input("isValid")]
        public Input<bool>? IsValid { get; set; }

        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
        /// </summary>
        [Input("networkSecurityGroupId")]
        public Input<string>? NetworkSecurityGroupId { get; set; }

        /// <summary>
        /// The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58").
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// Type of source for the rule. Required if `direction` = `INGRESS`.
        /// * `CIDR_BLOCK`: If the rule's `source` is an IP address range in CIDR notation.
        /// * `SERVICE_CIDR_BLOCK`: If the rule's `source` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic coming from a particular `Service` through a service gateway).
        /// * `NETWORK_SECURITY_GROUP`: If the rule's `source` is the OCID of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/).
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic.
        /// </summary>
        [Input("stateless")]
        public Input<bool>? Stateless { get; set; }

        /// <summary>
        /// Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed.
        /// </summary>
        [Input("tcpOptions")]
        public Input<Inputs.NetworkSecurityGroupSecurityRuleTcpOptionsGetArgs>? TcpOptions { get; set; }

        /// <summary>
        /// The date and time the security rule was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed.
        /// </summary>
        [Input("udpOptions")]
        public Input<Inputs.NetworkSecurityGroupSecurityRuleUdpOptionsGetArgs>? UdpOptions { get; set; }

        public NetworkSecurityGroupSecurityRuleState()
        {
        }
    }
}