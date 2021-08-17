// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Outputs
{

    [OutputType]
    public sealed class GetCoreBootVolumesBootVolumeSourceDetailsResult
    {
        /// <summary>
        /// The OCID of the boot volume replica.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The type can be one of these values: `bootVolume`, `bootVolumeBackup`, `bootVolumeReplica`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCoreBootVolumesBootVolumeSourceDetailsResult(
            string id,

            string type)
        {
            Id = id;
            Type = type;
        }
    }
}