// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Core.Outputs
{

    [OutputType]
    public sealed class GetInstanceConfigurationsInstanceConfigurationInstanceDetailsLaunchDetailsPlatformConfigResult
    {
        /// <summary>
        /// The number of NUMA nodes per socket.
        /// </summary>
        public readonly string NumaNodesPerSocket;
        /// <summary>
        /// The type of action to run when the instance is interrupted for eviction.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstanceConfigurationsInstanceConfigurationInstanceDetailsLaunchDetailsPlatformConfigResult(
            string numaNodesPerSocket,

            string type)
        {
            NumaNodesPerSocket = numaNodesPerSocket;
            Type = type;
        }
    }
}