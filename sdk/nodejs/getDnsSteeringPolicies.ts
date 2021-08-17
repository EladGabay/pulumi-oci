// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Steering Policies in Oracle Cloud Infrastructure DNS service.
 *
 * Gets a list of all steering policies in the specified compartment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testSteeringPolicies = oci.GetDnsSteeringPolicies({
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.steering_policy_display_name,
 *     displayNameContains: _var.steering_policy_display_name_contains,
 *     healthCheckMonitorId: oci_health_checks_http_monitor.test_http_monitor.id,
 *     id: _var.steering_policy_id,
 *     state: _var.steering_policy_state,
 *     template: _var.steering_policy_template,
 *     timeCreatedGreaterThanOrEqualTo: _var.steering_policy_time_created_greater_than_or_equal_to,
 *     timeCreatedLessThan: _var.steering_policy_time_created_less_than,
 * });
 * ```
 */
export function getDnsSteeringPolicies(args: GetDnsSteeringPoliciesArgs, opts?: pulumi.InvokeOptions): Promise<GetDnsSteeringPoliciesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getDnsSteeringPolicies:GetDnsSteeringPolicies", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "displayNameContains": args.displayNameContains,
        "filters": args.filters,
        "healthCheckMonitorId": args.healthCheckMonitorId,
        "id": args.id,
        "state": args.state,
        "template": args.template,
        "timeCreatedGreaterThanOrEqualTo": args.timeCreatedGreaterThanOrEqualTo,
        "timeCreatedLessThan": args.timeCreatedLessThan,
    }, opts);
}

/**
 * A collection of arguments for invoking GetDnsSteeringPolicies.
 */
export interface GetDnsSteeringPoliciesArgs {
    /**
     * The OCID of the compartment the resource belongs to.
     */
    compartmentId: string;
    /**
     * The displayName of a resource.
     */
    displayName?: string;
    /**
     * The partial displayName of a resource. Will match any resource whose name (case-insensitive) contains the provided value.
     */
    displayNameContains?: string;
    filters?: inputs.GetDnsSteeringPoliciesFilter[];
    /**
     * Search by health check monitor OCID. Will match any resource whose health check monitor ID matches the provided value.
     */
    healthCheckMonitorId?: string;
    /**
     * The OCID of a resource.
     */
    id?: string;
    /**
     * The state of a resource.
     */
    state?: string;
    /**
     * Search by steering template type. Will match any resource whose template type matches the provided value.
     */
    template?: string;
    /**
     * An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created on or after the indicated time.
     */
    timeCreatedGreaterThanOrEqualTo?: string;
    /**
     * An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created before the indicated time.
     */
    timeCreatedLessThan?: string;
}

/**
 * A collection of values returned by GetDnsSteeringPolicies.
 */
export interface GetDnsSteeringPoliciesResult {
    /**
     * The OCID of the compartment containing the steering policy.
     */
    readonly compartmentId: string;
    /**
     * A user-friendly name for the steering policy. Does not have to be unique and can be changed. Avoid entering confidential information.
     */
    readonly displayName?: string;
    readonly displayNameContains?: string;
    readonly filters?: outputs.GetDnsSteeringPoliciesFilter[];
    /**
     * The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with `rdata` matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with `rdata` not matching any monitored endpoint will be assumed healthy.
     */
    readonly healthCheckMonitorId?: string;
    /**
     * The OCID of the resource.
     */
    readonly id?: string;
    /**
     * The current state of the resource.
     */
    readonly state?: string;
    /**
     * The list of steering_policies.
     */
    readonly steeringPolicies: outputs.GetDnsSteeringPoliciesSteeringPolicy[];
    /**
     * A set of predefined rules based on the desired purpose of the steering policy. Each template utilizes Traffic Management's rules in a different order to produce the desired results when answering DNS queries.
     */
    readonly template?: string;
    readonly timeCreatedGreaterThanOrEqualTo?: string;
    readonly timeCreatedLessThan?: string;
}