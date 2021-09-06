// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core.Inputs
{

    public sealed class VolumeGroupSourceDetailsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type can be one of these values: `volumeGroupBackupId`, `volumeGroupId`, `volumeIds`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The OCID of the volume group backup to restore from.
        /// </summary>
        [Input("volumeGroupBackupId")]
        public Input<string>? VolumeGroupBackupId { get; set; }

        /// <summary>
        /// The OCID of the volume group to clone from.
        /// </summary>
        [Input("volumeGroupId")]
        public Input<string>? VolumeGroupId { get; set; }

        [Input("volumeIds")]
        private InputList<string>? _volumeIds;

        /// <summary>
        /// OCIDs for the volumes in this volume group.
        /// </summary>
        public InputList<string> VolumeIds
        {
            get => _volumeIds ?? (_volumeIds = new InputList<string>());
            set => _volumeIds = value;
        }

        public VolumeGroupSourceDetailsGetArgs()
        {
        }
    }
}