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
    public sealed class LoggingUnifiedAgentConfigurationServiceConfiguration
    {
        /// <summary>
        /// (Updatable) Type of Unified Agent service configuration.
        /// </summary>
        public readonly string ConfigurationType;
        /// <summary>
        /// (Updatable) Logging destination object.
        /// </summary>
        public readonly Outputs.LoggingUnifiedAgentConfigurationServiceConfigurationDestination Destination;
        /// <summary>
        /// (Updatable)
        /// </summary>
        public readonly ImmutableArray<Outputs.LoggingUnifiedAgentConfigurationServiceConfigurationSource> Sources;

        [OutputConstructor]
        private LoggingUnifiedAgentConfigurationServiceConfiguration(
            string configurationType,

            Outputs.LoggingUnifiedAgentConfigurationServiceConfigurationDestination destination,

            ImmutableArray<Outputs.LoggingUnifiedAgentConfigurationServiceConfigurationSource> sources)
        {
            ConfigurationType = configurationType;
            Destination = destination;
            Sources = sources;
        }
    }
}