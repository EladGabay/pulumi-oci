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
    public sealed class CloudGuardTargetTargetResponderRecipeResponderRuleDetails
    {
        /// <summary>
        /// (Updatable)
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// (Updatable) Configurations associated with the ResponderRule
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudGuardTargetTargetResponderRecipeResponderRuleDetailsConfiguration> Configurations;
        /// <summary>
        /// Identifies state for ResponderRule
        /// </summary>
        public readonly bool? IsEnabled;
        /// <summary>
        /// (Updatable) Execution Mode for ResponderRule
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private CloudGuardTargetTargetResponderRecipeResponderRuleDetails(
            string? condition,

            ImmutableArray<Outputs.CloudGuardTargetTargetResponderRecipeResponderRuleDetailsConfiguration> configurations,

            bool? isEnabled,

            string? mode)
        {
            Condition = condition;
            Configurations = configurations;
            IsEnabled = isEnabled;
            Mode = mode;
        }
    }
}