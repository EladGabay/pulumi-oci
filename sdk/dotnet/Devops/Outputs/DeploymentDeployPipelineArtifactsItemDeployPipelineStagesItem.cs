// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Devops.Outputs
{

    [OutputType]
    public sealed class DeploymentDeployPipelineArtifactsItemDeployPipelineStagesItem
    {
        /// <summary>
        /// Specifies the OCID of the stage to be redeployed.
        /// </summary>
        public readonly string? DeployStageId;
        /// <summary>
        /// (Updatable) Deployment display name. Avoid entering confidential information.
        /// </summary>
        public readonly string? DisplayName;

        [OutputConstructor]
        private DeploymentDeployPipelineArtifactsItemDeployPipelineStagesItem(
            string? deployStageId,

            string? displayName)
        {
            DeployStageId = deployStageId;
            DisplayName = displayName;
        }
    }
}