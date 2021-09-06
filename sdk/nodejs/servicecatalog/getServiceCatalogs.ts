// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the list of Service Catalogs in Oracle Cloud Infrastructure Service Catalog service.
 *
 * Lists all the service catalogs in the given compartment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as oci from "@pulumi/oci";
 *
 * const testServiceCatalogs = oci.servicecatalog.getServiceCatalogs({
 *     compartmentId: _var.compartment_id,
 *     displayName: _var.service_catalog_display_name,
 *     serviceCatalogId: oci_service_catalog_service_catalog.test_service_catalog.id,
 * });
 * ```
 */
export function getServiceCatalogs(args: GetServiceCatalogsArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceCatalogsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("oci:servicecatalog/getServiceCatalogs:getServiceCatalogs", {
        "compartmentId": args.compartmentId,
        "displayName": args.displayName,
        "filters": args.filters,
        "serviceCatalogId": args.serviceCatalogId,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceCatalogs.
 */
export interface GetServiceCatalogsArgs {
    /**
     * The unique identifier for the compartment.
     */
    compartmentId: string;
    /**
     * Exact match name filter.
     */
    displayName?: string;
    filters?: inputs.servicecatalog.GetServiceCatalogsFilter[];
    /**
     * The unique identifier for the service catalog.
     */
    serviceCatalogId?: string;
}

/**
 * A collection of values returned by getServiceCatalogs.
 */
export interface GetServiceCatalogsResult {
    /**
     * The Compartment id where the service catalog exists
     */
    readonly compartmentId: string;
    /**
     * The name of the service catalog.
     */
    readonly displayName?: string;
    readonly filters?: outputs.servicecatalog.GetServiceCatalogsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of service_catalog_collection.
     */
    readonly serviceCatalogCollections: outputs.servicecatalog.GetServiceCatalogsServiceCatalogCollection[];
    readonly serviceCatalogId?: string;
}