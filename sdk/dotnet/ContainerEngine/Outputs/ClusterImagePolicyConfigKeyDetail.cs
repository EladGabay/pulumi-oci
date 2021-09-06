// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ContainerEngine.Outputs
{

    [OutputType]
    public sealed class ClusterImagePolicyConfigKeyDetail
    {
        /// <summary>
        /// The OCID of the KMS key to be used as the master encryption key for Kubernetes secret encryption. When used, `kubernetesVersion` must be at least `v1.13.0`.
        /// </summary>
        public readonly string? KmsKeyId;

        [OutputConstructor]
        private ClusterImagePolicyConfigKeyDetail(string? kmsKeyId)
        {
            KmsKeyId = kmsKeyId;
        }
    }
}