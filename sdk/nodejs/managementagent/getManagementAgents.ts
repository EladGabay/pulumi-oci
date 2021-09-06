// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Management Agents in Oracle Cloud Infrastructure Management Agent service.
 *
 * Returns a list of Management Agent.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testManagementAgents = oci.managementagent.getManagementAgents({
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.management_agent_display_name,
 *     platformType: _var.management_agent_platform_type,
 *     pluginName: _var.management_agent_plugin_name,
 *     state: _var.management_agent_state,
 *     version: _var.management_agent_version,
 * });
 * ```
 */
export function getManagementAgents(args: GetManagementAgentsArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementAgentsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:managementagent/getManagementAgents:getManagementAgents", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "filters": args.filters,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getManagementAgents.
 */
export interface GetManagementAgentsArgs {
    /**
     * The ID of the compartment from which the Management Agents to be listed.
     */
    compartmentId: string;
    /**
     * Filter to return only Management Agents having the particular display name.
     */
    displayName?: string;
    filters?: inputs.managementagent.GetManagementAgentsFilter[];
    /**
     * Filter to return only Management Agents in the particular lifecycle state.
     */
    state?: string;
}

/**
 * A collection of values returned by getManagementAgents.
 */
export interface GetManagementAgentsResult {
    /**
     * Compartment Identifier
     */
    readonly compartmentId: string;
    /**
     * Management Agent Name
     */
    readonly displayName?: string;
    readonly filters?: outputs.managementagent.GetManagementAgentsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of management_agents.
     */
    readonly managementAgents: outputs.managementagent.GetManagementAgentsManagementAgent[];
    /**
     * The current state of managementAgent
     */
    readonly state?: string;
}