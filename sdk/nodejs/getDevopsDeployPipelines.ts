// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * This data source provides the list of Deploy Pipelines in Oracle Cloud Infrastructure Devops service.
 *
 * Returns a list of deployment pipelines.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testDeployPipelines = oci.GetDevopsDeployPipelines({
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.deploy_pipeline_display_name,
 *     id: _var.deploy_pipeline_id,
 *     projectId: oci_devops_project.test_project.id,
 *     state: _var.deploy_pipeline_state,
 * });
 * ```
 */
export function getDevopsDeployPipelines(args?: GetDevopsDeployPipelinesArgs, opts?: pulumi.InvokeOptions): Promise<GetDevopsDeployPipelinesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:index/getDevopsDeployPipelines:GetDevopsDeployPipelines", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "filters": args.filters,
        "id": args.id,
        "projectId": args.projectId,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking GetDevopsDeployPipelines.
 */
export interface GetDevopsDeployPipelinesArgs {
    /**
     * The OCID of the compartment in which to list resources.
     */
    compartmentId?: string;
    /**
     * A filter to return only resources that match the entire display name given.
     */
    displayName?: string;
    filters?: inputs.GetDevopsDeployPipelinesFilter[];
    /**
     * Unique identifier or OCID for listing a single resource by ID.
     */
    id?: string;
    /**
     * unique project identifier
     */
    projectId?: string;
    /**
     * A filter to return only DeployPipelines that matches the given lifecycleState.
     */
    state?: string;
}

/**
 * A collection of values returned by GetDevopsDeployPipelines.
 */
export interface GetDevopsDeployPipelinesResult {
    /**
     * The OCID of the compartment where the pipeline is created.
     */
    readonly compartmentId?: string;
    /**
     * The list of deploy_pipeline_collection.
     */
    readonly deployPipelineCollections: outputs.GetDevopsDeployPipelinesDeployPipelineCollection[];
    /**
     * Deployment pipeline display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
     */
    readonly displayName?: string;
    readonly filters?: outputs.GetDevopsDeployPipelinesFilter[];
    /**
     * Unique identifier that is immutable on creation.
     */
    readonly id?: string;
    /**
     * The OCID of a project.
     */
    readonly projectId?: string;
    /**
     * The current state of the deployment pipeline.
     */
    readonly state?: string;
}