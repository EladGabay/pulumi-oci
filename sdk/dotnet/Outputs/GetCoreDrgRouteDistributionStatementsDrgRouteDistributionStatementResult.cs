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
    public sealed class GetCoreDrgRouteDistributionStatementsDrgRouteDistributionStatementResult
    {
        /// <summary>
        /// `ACCEPT` indicates the route should be imported or exported as-is.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// The Oracle-assigned ID of the route distribution statement.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The action is applied only if all of the match criteria is met. If there are no match criteria in a statement, any input is considered a match and the action is applied.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCoreDrgRouteDistributionStatementsDrgRouteDistributionStatementMatchCriteriaResult> MatchCriterias;
        /// <summary>
        /// This field specifies the priority of each statement in a route distribution. Priorities must be unique within a particular route distribution. The priority will be represented as a number between 0 and 65535 where a lower number indicates a higher priority. When a route is processed, statements are applied in the order defined by their priority. The first matching rule dictates the action that will be taken on the route.
        /// </summary>
        public readonly int Priority;

        [OutputConstructor]
        private GetCoreDrgRouteDistributionStatementsDrgRouteDistributionStatementResult(
            string action,

            string id,

            ImmutableArray<Outputs.GetCoreDrgRouteDistributionStatementsDrgRouteDistributionStatementMatchCriteriaResult> matchCriterias,

            int priority)
        {
            Action = action;
            Id = id;
            MatchCriterias = matchCriterias;
            Priority = priority;
        }
    }
}