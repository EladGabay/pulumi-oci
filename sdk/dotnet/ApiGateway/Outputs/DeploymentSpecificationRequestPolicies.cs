// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ApiGateway.Outputs
{

    [OutputType]
    public sealed class DeploymentSpecificationRequestPolicies
    {
        /// <summary>
        /// (Updatable) Information on how to authenticate incoming requests.
        /// </summary>
        public readonly Outputs.DeploymentSpecificationRequestPoliciesAuthentication? Authentication;
        /// <summary>
        /// (Updatable) Enable CORS (Cross-Origin-Resource-Sharing) request handling.
        /// </summary>
        public readonly Outputs.DeploymentSpecificationRequestPoliciesCors? Cors;
        /// <summary>
        /// (Updatable) Limit the number of requests that should be handled for the specified window using a specfic key.
        /// </summary>
        public readonly Outputs.DeploymentSpecificationRequestPoliciesRateLimiting? RateLimiting;

        [OutputConstructor]
        private DeploymentSpecificationRequestPolicies(
            Outputs.DeploymentSpecificationRequestPoliciesAuthentication? authentication,

            Outputs.DeploymentSpecificationRequestPoliciesCors? cors,

            Outputs.DeploymentSpecificationRequestPoliciesRateLimiting? rateLimiting)
        {
            Authentication = authentication;
            Cors = cors;
            RateLimiting = rateLimiting;
        }
    }
}