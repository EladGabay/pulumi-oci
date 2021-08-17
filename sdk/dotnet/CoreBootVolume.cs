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
    /// This resource provides the Boot Volume resource in Oracle Cloud Infrastructure Core service.
    /// 
    /// Creates a new boot volume in the specified compartment from an existing boot volume or a boot volume backup.
    /// For general information about boot volumes, see [Boot Volumes](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumes.htm).
    /// You may optionally specify a *display name* for the volume, which is simply a friendly name or
    /// description. It does not have to be unique, and you can change it. Avoid entering confidential information.
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
    ///         var testBootVolume = new Oci.CoreBootVolume("testBootVolume", new Oci.CoreBootVolumeArgs
    ///         {
    ///             CompartmentId = @var.Compartment_id,
    ///             SourceDetails = new Oci.Inputs.CoreBootVolumeSourceDetailsArgs
    ///             {
    ///                 Id = @var.Boot_volume_source_details_id,
    ///                 Type = @var.Boot_volume_source_details_type,
    ///             },
    ///             AvailabilityDomain = @var.Boot_volume_availability_domain,
    ///             BackupPolicyId = data.Oci_core_volume_backup_policies.Test_volume_backup_policies.Volume_backup_policies[0].Id,
    ///             BootVolumeReplicas = 
    ///             {
    ///                 new Oci.Inputs.CoreBootVolumeBootVolumeReplicaArgs
    ///                 {
    ///                     AvailabilityDomain = @var.Boot_volume_boot_volume_replicas_availability_domain,
    ///                     DisplayName = @var.Boot_volume_boot_volume_replicas_display_name,
    ///                 },
    ///             },
    ///             DefinedTags = 
    ///             {
    ///                 { "Operations.CostCenter", "42" },
    ///             },
    ///             DisplayName = @var.Boot_volume_display_name,
    ///             FreeformTags = 
    ///             {
    ///                 { "Department", "Finance" },
    ///             },
    ///             IsAutoTuneEnabled = @var.Boot_volume_is_auto_tune_enabled,
    ///             KmsKeyId = oci_kms_key.Test_key.Id,
    ///             SizeInGbs = @var.Boot_volume_size_in_gbs,
    ///             VpusPerGb = @var.Boot_volume_vpus_per_gb,
    ///             BootVolumeReplicasDeletion = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// BootVolumes can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import oci:index/coreBootVolume:CoreBootVolume test_boot_volume "id"
    /// ```
    /// </summary>
    [OciResourceType("oci:index/coreBootVolume:CoreBootVolume")]
    public partial class CoreBootVolume : Pulumi.CustomResource
    {
        /// <summary>
        /// The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle.
        /// </summary>
        [Output("autoTunedVpusPerGb")]
        public Output<string> AutoTunedVpusPerGb { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        [Output("availabilityDomain")]
        public Output<string> AvailabilityDomain { get; private set; } = null!;

        /// <summary>
        /// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
        /// </summary>
        [Output("backupPolicyId")]
        public Output<string> BackupPolicyId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
        /// </summary>
        [Output("bootVolumeReplicas")]
        public Output<ImmutableArray<Outputs.CoreBootVolumeBootVolumeReplica>> BootVolumeReplicas { get; private set; } = null!;

        /// <summary>
        /// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `boot_volume_replicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `boot_volume_replicas` again.
        /// </summary>
        [Output("bootVolumeReplicasDeletion")]
        public Output<bool?> BootVolumeReplicasDeletion { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the compartment that contains the boot volume.
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
        /// The image OCID used to create the boot volume.
        /// </summary>
        [Output("imageId")]
        public Output<string> ImageId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
        /// </summary>
        [Output("isAutoTuneEnabled")]
        public Output<bool> IsAutoTuneEnabled { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
        /// </summary>
        [Output("isHydrated")]
        public Output<bool> IsHydrated { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The size of the volume in GBs.
        /// </summary>
        [Output("sizeInGbs")]
        public Output<string> SizeInGbs { get; private set; } = null!;

        /// <summary>
        /// The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`.
        /// </summary>
        [Output("sizeInMbs")]
        public Output<string> SizeInMbs { get; private set; } = null!;

        [Output("sourceDetails")]
        public Output<Outputs.CoreBootVolumeSourceDetails> SourceDetails { get; private set; } = null!;

        /// <summary>
        /// The current state of a boot volume.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        [Output("systemTags")]
        public Output<ImmutableDictionary<string, object>> SystemTags { get; private set; } = null!;

        /// <summary>
        /// The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        /// </summary>
        [Output("timeCreated")]
        public Output<string> TimeCreated { get; private set; } = null!;

        /// <summary>
        /// The OCID of the source volume group.
        /// </summary>
        [Output("volumeGroupId")]
        public Output<string> VolumeGroupId { get; private set; } = null!;

        /// <summary>
        /// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
        /// </summary>
        [Output("vpusPerGb")]
        public Output<string> VpusPerGb { get; private set; } = null!;


        /// <summary>
        /// Create a CoreBootVolume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CoreBootVolume(string name, CoreBootVolumeArgs args, CustomResourceOptions? options = null)
            : base("oci:index/coreBootVolume:CoreBootVolume", name, args ?? new CoreBootVolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CoreBootVolume(string name, Input<string> id, CoreBootVolumeState? state = null, CustomResourceOptions? options = null)
            : base("oci:index/coreBootVolume:CoreBootVolume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CoreBootVolume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CoreBootVolume Get(string name, Input<string> id, CoreBootVolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new CoreBootVolume(name, id, state, options);
        }
    }

    public sealed class CoreBootVolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        [Input("availabilityDomain", required: true)]
        public Input<string> AvailabilityDomain { get; set; } = null!;

        /// <summary>
        /// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
        /// </summary>
        [Input("backupPolicyId")]
        public Input<string>? BackupPolicyId { get; set; }

        [Input("bootVolumeReplicas")]
        private InputList<Inputs.CoreBootVolumeBootVolumeReplicaArgs>? _bootVolumeReplicas;

        /// <summary>
        /// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
        /// </summary>
        public InputList<Inputs.CoreBootVolumeBootVolumeReplicaArgs> BootVolumeReplicas
        {
            get => _bootVolumeReplicas ?? (_bootVolumeReplicas = new InputList<Inputs.CoreBootVolumeBootVolumeReplicaArgs>());
            set => _bootVolumeReplicas = value;
        }

        /// <summary>
        /// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `boot_volume_replicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `boot_volume_replicas` again.
        /// </summary>
        [Input("bootVolumeReplicasDeletion")]
        public Input<bool>? BootVolumeReplicasDeletion { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the compartment that contains the boot volume.
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
        /// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
        /// </summary>
        [Input("isAutoTuneEnabled")]
        public Input<bool>? IsAutoTuneEnabled { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// (Updatable) The size of the volume in GBs.
        /// </summary>
        [Input("sizeInGbs")]
        public Input<string>? SizeInGbs { get; set; }

        [Input("sourceDetails", required: true)]
        public Input<Inputs.CoreBootVolumeSourceDetailsArgs> SourceDetails { get; set; } = null!;

        /// <summary>
        /// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
        /// </summary>
        [Input("vpusPerGb")]
        public Input<string>? VpusPerGb { get; set; }

        public CoreBootVolumeArgs()
        {
        }
    }

    public sealed class CoreBootVolumeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle.
        /// </summary>
        [Input("autoTunedVpusPerGb")]
        public Input<string>? AutoTunedVpusPerGb { get; set; }

        /// <summary>
        /// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
        /// </summary>
        [Input("availabilityDomain")]
        public Input<string>? AvailabilityDomain { get; set; }

        /// <summary>
        /// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
        /// </summary>
        [Input("backupPolicyId")]
        public Input<string>? BackupPolicyId { get; set; }

        [Input("bootVolumeReplicas")]
        private InputList<Inputs.CoreBootVolumeBootVolumeReplicaGetArgs>? _bootVolumeReplicas;

        /// <summary>
        /// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
        /// </summary>
        public InputList<Inputs.CoreBootVolumeBootVolumeReplicaGetArgs> BootVolumeReplicas
        {
            get => _bootVolumeReplicas ?? (_bootVolumeReplicas = new InputList<Inputs.CoreBootVolumeBootVolumeReplicaGetArgs>());
            set => _bootVolumeReplicas = value;
        }

        /// <summary>
        /// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `boot_volume_replicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `boot_volume_replicas` again.
        /// </summary>
        [Input("bootVolumeReplicasDeletion")]
        public Input<bool>? BootVolumeReplicasDeletion { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the compartment that contains the boot volume.
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
        /// The image OCID used to create the boot volume.
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
        /// </summary>
        [Input("isAutoTuneEnabled")]
        public Input<bool>? IsAutoTuneEnabled { get; set; }

        /// <summary>
        /// Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
        /// </summary>
        [Input("isHydrated")]
        public Input<bool>? IsHydrated { get; set; }

        /// <summary>
        /// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// (Updatable) The size of the volume in GBs.
        /// </summary>
        [Input("sizeInGbs")]
        public Input<string>? SizeInGbs { get; set; }

        /// <summary>
        /// The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`.
        /// </summary>
        [Input("sizeInMbs")]
        public Input<string>? SizeInMbs { get; set; }

        [Input("sourceDetails")]
        public Input<Inputs.CoreBootVolumeSourceDetailsGetArgs>? SourceDetails { get; set; }

        /// <summary>
        /// The current state of a boot volume.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("systemTags")]
        private InputMap<object>? _systemTags;

        /// <summary>
        /// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
        /// </summary>
        public InputMap<object> SystemTags
        {
            get => _systemTags ?? (_systemTags = new InputMap<object>());
            set => _systemTags = value;
        }

        /// <summary>
        /// The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
        /// </summary>
        [Input("timeCreated")]
        public Input<string>? TimeCreated { get; set; }

        /// <summary>
        /// The OCID of the source volume group.
        /// </summary>
        [Input("volumeGroupId")]
        public Input<string>? VolumeGroupId { get; set; }

        /// <summary>
        /// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
        /// </summary>
        [Input("vpusPerGb")]
        public Input<string>? VpusPerGb { get; set; }

        public CoreBootVolumeState()
        {
        }
    }
}