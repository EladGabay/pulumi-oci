// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    /// <summary>
    /// This resource provides the Tsig Key resource in Oracle Cloud Infrastructure DNS service.
    /// 
    /// Creates a new TSIG key in the specified compartment. There is no
    /// `opc-retry-token` header since TSIG key names must be globally unique.
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
    ///         var testTsigKey = new Oci.DnsTsigKey("testTsigKey", new Oci.DnsTsigKeyArgs
    ///         {
    ///             Algorithm = @var.Tsig_key_algorithm,
    ///             CompartmentId = @var.Compartment_id,
    ///             Secret = @var.Tsig_key_secret,
    ///             DefinedTags = @var.Tsig_key_defined_tags,
    ///             FreeformTags = @var.Tsig_key_freeform_tags,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// TsigKeys can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:index/dnsTsigKey:DnsTsigKey test_tsig_key "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:index/dnsTsigKey:DnsTsigKey")]
    public partial class DnsTsigKey : Pulumi.CustomResource
    {
        /// <summary>
        /// TSIG key algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2).
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the compartment containing the TSIG key.
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// A globally unique domain name identifying the key for a given pair of hosts.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A base64 string encoding the binary shared secret.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;

        /// <summary>
        /// The canonical absolute URL of the resource.
        /// </summary>
        [Output("self")]
        public Output<string> Self { get; private set; } = null!;

        /// <summary>
        /// The current state of the resource.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The date and time the resource was created, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The date and time the resource was last updated, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Output("timeUpdated")]
        public Output<string> TimeUpdated { get; private set; } = null!;


        /// <summary>
        /// Create a DnsTsigKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsTsigKey(string name, DnsTsigKeyArgs args, CustomResourceOptions? options = null)
            : base("oci:index/dnsTsigKey:DnsTsigKey", name, args ?? new DnsTsigKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsTsigKey(string name, Input<string> id, DnsTsigKeyState? state = null, CustomResourceOptions? options = null)
            : base("oci:index/dnsTsigKey:DnsTsigKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DnsTsigKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsTsigKey Get(string name, Input<string> id, DnsTsigKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsTsigKey(name, id, state, options);
        }
    }

    public sealed class DnsTsigKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// TSIG key algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2).
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the compartment containing the TSIG key.
        /// </summary>
        [Input("compartmentId", required: true)]
        public Input<string> CompartmentId { get; set; } = null!;

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// A globally unique domain name identifying the key for a given pair of hosts.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A base64 string encoding the binary shared secret.
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        public DnsTsigKeyArgs()
        {
        }
    }

    public sealed class DnsTsigKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// TSIG key algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256, hmac-sha512. For more information on these algorithms, see [RFC 4635](https://tools.ietf.org/html/rfc4635#section-2).
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the compartment containing the TSIG key.
        /// </summary>
        [Input("compartmentId")]
        public Input<string>? CompartmentId { get; set; }

        [Input("definedTags")]
        private InputMap<object>? _definedTags;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public InputMap<object> DefinedTags
        {
            get => _definedTags ?? (_definedTags = new InputMap<object>());
            set => _definedTags = value;
        }

        [Input("freeformTags")]
        private InputMap<object>? _freeformTags;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
        /// </summary>
        public InputMap<object> FreeformTags
        {
            get => _freeformTags ?? (_freeformTags = new InputMap<object>());
            set => _freeformTags = value;
        }

        /// <summary>
        /// A globally unique domain name identifying the key for a given pair of hosts.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A base64 string encoding the binary shared secret.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        /// <summary>
        /// The canonical absolute URL of the resource.
        /// </summary>
        [Input("self")]
        public Input<string>? Self { get; set; }

        /// <summary>
        /// The current state of the resource.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The date and time the resource was created, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// The date and time the resource was last updated, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Input("timeUpdated")]
        public Input<string>? TimeUpdated { get; set; }

        public DnsTsigKeyState()
        {
        }
    }
}