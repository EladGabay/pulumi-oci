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
    public sealed class GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsResult
    {
        /// <summary>
        /// Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy.
        /// </summary>
        public readonly Outputs.GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsFilterHeadersResult FilterHeaders;
        /// <summary>
        /// Rename HTTP headers as they pass through the gateway.
        /// </summary>
        public readonly Outputs.GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsRenameHeadersResult RenameHeaders;
        /// <summary>
        /// Set HTTP headers as they pass through the gateway.
        /// </summary>
        public readonly Outputs.GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsSetHeadersResult SetHeaders;

        [OutputConstructor]
        private GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsResult(
            Outputs.GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsFilterHeadersResult filterHeaders,

            Outputs.GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsRenameHeadersResult renameHeaders,

            Outputs.GetApigatewayDeploymentsDeploymentCollectionSpecificationRouteResponsePoliciesHeaderTransformationsSetHeadersResult setHeaders)
        {
            FilterHeaders = filterHeaders;
            RenameHeaders = renameHeaders;
            SetHeaders = setHeaders;
        }
    }
}