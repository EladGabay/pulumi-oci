// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci
{
    public static class GetLogAnalyticsNamespace
    {
        /// <summary>
        /// This data source provides details about a specific Namespace resource in Oracle Cloud Infrastructure Log Analytics service.
        /// 
        /// This API gets the namespace details of a tenancy already onboarded in Logging Analytics Application
        /// 
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
        ///         var testNamespace = Output.Create(Oci.GetLogAnalyticsNamespace.InvokeAsync(new Oci.GetLogAnalyticsNamespaceArgs
        ///         {
        ///             Namespace = @var.Namespace_namespace,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLogAnalyticsNamespaceResult> InvokeAsync(GetLogAnalyticsNamespaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLogAnalyticsNamespaceResult>("oci:index/getLogAnalyticsNamespace:GetLogAnalyticsNamespace", args ?? new GetLogAnalyticsNamespaceArgs(), options.WithVersion());
    }


    public sealed class GetLogAnalyticsNamespaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Logging Analytics namespace used for the request.
        /// </summary>
        [Input("namespace", required: true)]
        public string Namespace { get; set; } = null!;

        public GetLogAnalyticsNamespaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLogAnalyticsNamespaceResult
    {
        /// <summary>
        /// The is the tenancy ID
        /// </summary>
        public readonly string CompartmentId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// This indicates if the tenancy is onboarded to Logging Analytics
        /// </summary>
        public readonly bool IsOnboarded;
        /// <summary>
        /// This is the namespace name of a tenancy
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private GetLogAnalyticsNamespaceResult(
            string compartmentId,

            string id,

            bool isOnboarded,

            string @namespace)
        {
            CompartmentId = compartmentId;
            Id = id;
            IsOnboarded = isOnboarded;
            Namespace = @namespace;
        }
    }
}