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
    public sealed class GetCloudGuardTargetsTargetCollectionItemTargetDetectorRecipeEffectiveDetectorRuleDetailsResult
    {
        /// <summary>
        /// Condition group corresponding to each compartment
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudGuardTargetsTargetCollectionItemTargetDetectorRecipeEffectiveDetectorRuleDetailsConditionGroupResult> ConditionGroups;
        /// <summary>
        /// ResponderRule configurations
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCloudGuardTargetsTargetCollectionItemTargetDetectorRecipeEffectiveDetectorRuleDetailsConfigurationResult> Configurations;
        /// <summary>
        /// configuration allowed or not
        /// </summary>
        public readonly bool IsConfigurationAllowed;
        /// <summary>
        /// Identifies state for ResponderRule
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// user defined labels for a detector rule
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// The Risk Level
        /// </summary>
        public readonly string RiskLevel;

        [OutputConstructor]
        private GetCloudGuardTargetsTargetCollectionItemTargetDetectorRecipeEffectiveDetectorRuleDetailsResult(
            ImmutableArray<Outputs.GetCloudGuardTargetsTargetCollectionItemTargetDetectorRecipeEffectiveDetectorRuleDetailsConditionGroupResult> conditionGroups,

            ImmutableArray<Outputs.GetCloudGuardTargetsTargetCollectionItemTargetDetectorRecipeEffectiveDetectorRuleDetailsConfigurationResult> configurations,

            bool isConfigurationAllowed,

            bool isEnabled,

            ImmutableArray<string> labels,

            string riskLevel)
        {
            ConditionGroups = conditionGroups;
            Configurations = configurations;
            IsConfigurationAllowed = isConfigurationAllowed;
            IsEnabled = isEnabled;
            Labels = labels;
            RiskLevel = riskLevel;
        }
    }
}