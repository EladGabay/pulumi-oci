// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Waas
{
    public static class GetCustomProtectionRule
    {
        /// <summary>
        /// This data source provides details about a specific Custom Protection Rule resource in Oracle Cloud Infrastructure Web Application Acceleration and Security service.
        /// 
        /// Gets the details of a custom protection rule.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Oci = Pulumi.Oci;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testCustomProtectionRule = Output.Create(Oci.Waas.GetCustomProtectionRule.InvokeAsync(new Oci.Waas.GetCustomProtectionRuleArgs
        ///         {
        ///             CustomProtectionRuleId = oci_waas_custom_protection_rule.Test_custom_protection_rule.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCustomProtectionRuleResult> InvokeAsync(GetCustomProtectionRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomProtectionRuleResult>("oci:waas/getCustomProtectionRule:getCustomProtectionRule", args ?? new GetCustomProtectionRuleArgs(), options.WithVersion());
    }


    public sealed class GetCustomProtectionRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom protection rule. This number is generated when the custom protection rule is added to the compartment.
        /// </summary>
        [Input("customProtectionRuleId", required: true)]
        public string CustomProtectionRuleId { get; set; } = null!;

        public GetCustomProtectionRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomProtectionRuleResult
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom protection rule's compartment.
        /// </summary>
        public readonly string CompartmentId;
        public readonly string CustomProtectionRuleId;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// The description of the custom protection rule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The user-friendly name of the custom protection rule.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the custom protection rule.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The auto-generated ID for the custom protection rule. These IDs are referenced in logs.
        /// </summary>
        public readonly ImmutableArray<string> ModSecurityRuleIds;
        /// <summary>
        /// The current lifecycle state of the custom protection rule.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The template text of the custom protection rule. All custom protection rules are expressed in ModSecurity Rule Language.
        /// </summary>
        public readonly string Template;
        /// <summary>
        /// The date and time the protection rule was created, expressed in RFC 3339 timestamp format.
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private GetCustomProtectionRuleResult(
            string compartmentId,

            string customProtectionRuleId,

            ImmutableDictionary<string, object> definedTags,

            string description,

            string displayName,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            ImmutableArray<string> modSecurityRuleIds,

            string state,

            string template,

            string timeCreated)
        {
            CompartmentId = compartmentId;
            CustomProtectionRuleId = customProtectionRuleId;
            DefinedTags = definedTags;
            Description = description;
            DisplayName = displayName;
            FreeformTags = freeformTags;
            Id = id;
            ModSecurityRuleIds = modSecurityRuleIds;
            State = state;
            Template = template;
            TimeCreated = timeCreated;
        }
    }
}