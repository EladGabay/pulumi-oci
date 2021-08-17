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
    /// This resource provides the Certificate resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.
    /// 
    /// Allows an SSL certificate to be added to a WAAS policy. The Web Application Firewall terminates SSL connections to inspect requests in runtime, and then re-encrypts requests before sending them to the origin for fulfillment.
    /// 
    /// For more information, see [WAF Settings](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/wafsettings.htm).
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
    ///         var testCertificate = new Oci.WaasCertificate("testCertificate", new Oci.WaasCertificateArgs
    ///         {
    ///             CertificateData = @var.Certificate_certificate_data,
    ///             CompartmentId = @var.Compartment_id,
    ///             PrivateKeyData = @var.Certificate_private_key_data,
    ///             DefinedTags = 
    ///             {
    ///                 { "Operations.CostCenter", "42" },
    ///             },
    ///             DisplayName = @var.Certificate_display_name,
    ///             FreeformTags = 
    ///             {
    ///                 { "Department", "Finance" },
    ///             },
    ///             IsTrustVerificationDisabled = @var.Certificate_is_trust_verification_disabled,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Import is not supported for this resource.
    /// </summary>
    [OciResourceType("oci:index/waasCertificate:WaasCertificate")]
    public partial class WaasCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// The data of the SSL certificate.
        /// </summary>
        [Output("certificateData")]
        public Output<string> CertificateData { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the SSL certificate.
        /// </summary>
        [Output("compartmentId")]
        public Output<string> CompartmentId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        [Output("definedTags")]
        public Output<ImmutableDictionary<string, object>> DefinedTags { get; private set; } = null!;

        /// <summary>
        /// (Updatable) A user-friendly name for the SSL certificate. The name can be changed and does not need to be unique.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Additional attributes associated with users or public keys for managing relationships between Certificate Authorities.
        /// </summary>
        [Output("extensions")]
        public Output<ImmutableArray<Outputs.WaasCertificateExtension>> Extensions { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        [Output("freeformTags")]
        public Output<ImmutableDictionary<string, object>> FreeformTags { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if the SSL certificate is self-signed.
        /// </summary>
        [Output("isTrustVerificationDisabled")]
        public Output<bool> IsTrustVerificationDisabled { get; private set; } = null!;

        [Output("issuedBy")]
        public Output<string> IssuedBy { get; private set; } = null!;

        /// <summary>
        /// The issuer of the certificate.
        /// </summary>
        [Output("issuerName")]
        public Output<Outputs.WaasCertificateIssuerName> IssuerName { get; private set; } = null!;

        /// <summary>
        /// The private key of the SSL certificate.
        /// </summary>
        [Output("privateKeyData")]
        public Output<string> PrivateKeyData { get; private set; } = null!;

        /// <summary>
        /// Information about the public key and the algorithm used by the public key.
        /// </summary>
        [Output("publicKeyInfo")]
        public Output<Outputs.WaasCertificatePublicKeyInfo> PublicKeyInfo { get; private set; } = null!;

        /// <summary>
        /// A unique, positive integer assigned by the Certificate Authority (CA). The issuer name and serial number identify a unique certificate.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// The identifier for the cryptographic algorithm used by the Certificate Authority (CA) to sign this certificate.
        /// </summary>
        [Output("signatureAlgorithm")]
        public Output<string> SignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The current lifecycle state of the SSL certificate.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The entity to be secured by the certificate.
        /// </summary>
        [Output("subjectName")]
        public Output<Outputs.WaasCertificateSubjectName> SubjectName { get; private set; } = null!;

        /// <summary>
        /// The date and time the certificate was created, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Output("timeNotValidAfter")]
        public Output<string> TimeNotValidAfter { get; private set; } = null!;

        /// <summary>
        /// The date and time the certificate will become valid, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Output("timeNotValidBefore")]
        public Output<string> TimeNotValidBefore { get; private set; } = null!;

        /// <summary>
        /// The version of the encoded certificate.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a WaasCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WaasCertificate(string name, WaasCertificateArgs args, CustomResourceOptions? options = null)
            : base("oci:index/waasCertificate:WaasCertificate", name, args ?? new WaasCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WaasCertificate(string name, Input<string> id, WaasCertificateState? state = null, CustomResourceOptions? options = null)
            : base("oci:index/waasCertificate:WaasCertificate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WaasCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WaasCertificate Get(string name, Input<string> id, WaasCertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new WaasCertificate(name, id, state, options);
        }
    }

    public sealed class WaasCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data of the SSL certificate.
        /// </summary>
        [Input("certificateData", required: true)]
        public Input<string> CertificateData { get; set; } = null!;

        /// <summary>
        /// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the SSL certificate.
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
        /// (Updatable) A user-friendly name for the SSL certificate. The name can be changed and does not need to be unique.
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
        /// Set to `true` if the SSL certificate is self-signed.
        /// </summary>
        [Input("isTrustVerificationDisabled")]
        public Input<bool>? IsTrustVerificationDisabled { get; set; }

        /// <summary>
        /// The private key of the SSL certificate.
        /// </summary>
        [Input("privateKeyData", required: true)]
        public Input<string> PrivateKeyData { get; set; } = null!;

        public WaasCertificateArgs()
        {
        }
    }

    public sealed class WaasCertificateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The data of the SSL certificate.
        /// </summary>
        [Input("certificateData")]
        public Input<string>? CertificateData { get; set; }

        /// <summary>
        /// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the SSL certificate.
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
        /// (Updatable) A user-friendly name for the SSL certificate. The name can be changed and does not need to be unique.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("extensions")]
        private InputList<Inputs.WaasCertificateExtensionGetArgs>? _extensions;

        /// <summary>
        /// Additional attributes associated with users or public keys for managing relationships between Certificate Authorities.
        /// </summary>
        public InputList<Inputs.WaasCertificateExtensionGetArgs> Extensions
        {
            get => _extensions ?? (_extensions = new InputList<Inputs.WaasCertificateExtensionGetArgs>());
            set => _extensions = value;
        }

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
        /// Set to `true` if the SSL certificate is self-signed.
        /// </summary>
        [Input("isTrustVerificationDisabled")]
        public Input<bool>? IsTrustVerificationDisabled { get; set; }

        [Input("issuedBy")]
        public Input<string>? IssuedBy { get; set; }

        /// <summary>
        /// The issuer of the certificate.
        /// </summary>
        [Input("issuerName")]
        public Input<Inputs.WaasCertificateIssuerNameGetArgs>? IssuerName { get; set; }

        /// <summary>
        /// The private key of the SSL certificate.
        /// </summary>
        [Input("privateKeyData")]
        public Input<string>? PrivateKeyData { get; set; }

        /// <summary>
        /// Information about the public key and the algorithm used by the public key.
        /// </summary>
        [Input("publicKeyInfo")]
        public Input<Inputs.WaasCertificatePublicKeyInfoGetArgs>? PublicKeyInfo { get; set; }

        /// <summary>
        /// A unique, positive integer assigned by the Certificate Authority (CA). The issuer name and serial number identify a unique certificate.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        /// <summary>
        /// The identifier for the cryptographic algorithm used by the Certificate Authority (CA) to sign this certificate.
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        /// <summary>
        /// The current lifecycle state of the SSL certificate.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The entity to be secured by the certificate.
        /// </summary>
        [Input("subjectName")]
        public Input<Inputs.WaasCertificateSubjectNameGetArgs>? SubjectName { get; set; }

        /// <summary>
        /// The date and time the certificate was created, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// The date and time the certificate will expire, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Input("timeNotValidAfter")]
        public Input<string>? TimeNotValidAfter { get; set; }

        /// <summary>
        /// The date and time the certificate will become valid, expressed in RFC 3339 timestamp format.
        /// </summary>
        [Input("timeNotValidBefore")]
        public Input<string>? TimeNotValidBefore { get; set; }

        /// <summary>
        /// The version of the encoded certificate.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public WaasCertificateState()
        {
        }
    }
}