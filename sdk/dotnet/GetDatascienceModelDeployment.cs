// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetDatascienceModelDeployment
    {
        /// <summary>
        /// This data source provides details about a specific Model Deployment resource in Oracle Cloud Infrastructure Datascience service.
        /// 
        /// Retrieves the model deployment for the specified `modelDeploymentId`.
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
        ///         var testModelDeployment = Output.Create(Oci.GetDatascienceModelDeployment.InvokeAsync(new Oci.GetDatascienceModelDeploymentArgs
        ///         {
        ///             ModelDeploymentId = oci_datascience_model_deployment.Test_model_deployment.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDatascienceModelDeploymentResult> InvokeAsync(GetDatascienceModelDeploymentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatascienceModelDeploymentResult>("oci:index/getDatascienceModelDeployment:GetDatascienceModelDeployment", args ?? new GetDatascienceModelDeploymentArgs(), options.WithVersion());
    }


    public sealed class GetDatascienceModelDeploymentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
        /// </summary>
        [Input("modelDeploymentId", required: true)]
        public string ModelDeploymentId { get; set; } = null!;

        public GetDatascienceModelDeploymentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatascienceModelDeploymentResult
    {
        /// <summary>
        /// The log details for each category.
        /// </summary>
        public readonly Outputs.GetDatascienceModelDeploymentCategoryLogDetailsResult CategoryLogDetails;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment's compartment.
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model deployment.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> DefinedTags;
        /// <summary>
        /// A short description of the model deployment.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A user-friendly display name for the resource. Does not have to be unique, and can be modified. Avoid entering confidential information. Example: `My ModelDeployment`
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
        /// </summary>
        public readonly ImmutableDictionary<string, object> FreeformTags;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model deployment.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Details about the state of the model deployment.
        /// </summary>
        public readonly string LifecycleDetails;
        /// <summary>
        /// The model deployment configuration details.
        /// </summary>
        public readonly Outputs.GetDatascienceModelDeploymentModelDeploymentConfigurationDetailsResult ModelDeploymentConfigurationDetails;
        public readonly string ModelDeploymentId;
        /// <summary>
        /// The URL to interact with the model deployment.
        /// </summary>
        public readonly string ModelDeploymentUrl;
        /// <summary>
        /// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model deployment.
        /// </summary>
        public readonly string ProjectId;
        /// <summary>
        /// The state of the model deployment.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The date and time the resource was created, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private GetDatascienceModelDeploymentResult(
            Outputs.GetDatascienceModelDeploymentCategoryLogDetailsResult categoryLogDetails,

            string compartmentId,

            string createdBy,

            ImmutableDictionary<string, object> definedTags,

            string description,

            string displayName,

            ImmutableDictionary<string, object> freeformTags,

            string id,

            string lifecycleDetails,

            Outputs.GetDatascienceModelDeploymentModelDeploymentConfigurationDetailsResult modelDeploymentConfigurationDetails,

            string modelDeploymentId,

            string modelDeploymentUrl,

            string projectId,

            string state,

            string timeCreated)
        {
            CategoryLogDetails = categoryLogDetails;
            CompartmentId = compartmentId;
            CreatedBy = createdBy;
            DefinedTags = definedTags;
            Description = description;
            DisplayName = displayName;
            FreeformTags = freeformTags;
            Id = id;
            LifecycleDetails = lifecycleDetails;
            ModelDeploymentConfigurationDetails = modelDeploymentConfigurationDetails;
            ModelDeploymentId = modelDeploymentId;
            ModelDeploymentUrl = modelDeploymentUrl;
            ProjectId = projectId;
            State = state;
            TimeCreated = timeCreated;
        }
    }
}