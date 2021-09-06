// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.ServiceCatalog
{
    public static class GetServiceCatalogAssociations
    {
        /// <summary>
        /// This data source provides the list of Service Catalog Associations in Oracle Cloud Infrastructure Service Catalog service.
        /// 
        /// Lists all the resource associations for a specific service catalog.
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
        ///         var testServiceCatalogAssociations = Output.Create(Oci.ServiceCatalog.GetServiceCatalogAssociations.InvokeAsync(new Oci.ServiceCatalog.GetServiceCatalogAssociationsArgs
        ///         {
        ///             EntityId = oci_service_catalog_entity.Test_entity.Id,
        ///             EntityType = @var.Service_catalog_association_entity_type,
        ///             ServiceCatalogAssociationId = oci_service_catalog_service_catalog_association.Test_service_catalog_association.Id,
        ///             ServiceCatalogId = oci_service_catalog_service_catalog.Test_service_catalog.Id,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetServiceCatalogAssociationsResult> InvokeAsync(GetServiceCatalogAssociationsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceCatalogAssociationsResult>("oci:servicecatalog/getServiceCatalogAssociations:getServiceCatalogAssociations", args ?? new GetServiceCatalogAssociationsArgs(), options.WithVersion());
    }


    public sealed class GetServiceCatalogAssociationsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique identifier of the entity associated with service catalog.
        /// </summary>
        [Input("entityId")]
        public string? EntityId { get; set; }

        /// <summary>
        /// The type of the application in the service catalog.
        /// </summary>
        [Input("entityType")]
        public string? EntityType { get; set; }

        [Input("filters")]
        private List<Inputs.GetServiceCatalogAssociationsFilterArgs>? _filters;
        public List<Inputs.GetServiceCatalogAssociationsFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetServiceCatalogAssociationsFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The unique identifier for the service catalog association.
        /// </summary>
        [Input("serviceCatalogAssociationId")]
        public string? ServiceCatalogAssociationId { get; set; }

        /// <summary>
        /// The unique identifier for the service catalog.
        /// </summary>
        [Input("serviceCatalogId")]
        public string? ServiceCatalogId { get; set; }

        public GetServiceCatalogAssociationsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceCatalogAssociationsResult
    {
        /// <summary>
        /// Identifier of the entity being associated with service catalog.
        /// </summary>
        public readonly string? EntityId;
        /// <summary>
        /// The type of the entity that is associated with the service catalog.
        /// </summary>
        public readonly string? EntityType;
        public readonly ImmutableArray<Outputs.GetServiceCatalogAssociationsFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of service_catalog_association_collection.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServiceCatalogAssociationsServiceCatalogAssociationCollectionResult> ServiceCatalogAssociationCollections;
        public readonly string? ServiceCatalogAssociationId;
        /// <summary>
        /// Identifier of the service catalog.
        /// </summary>
        public readonly string? ServiceCatalogId;

        [OutputConstructor]
        private GetServiceCatalogAssociationsResult(
            string? entityId,

            string? entityType,

            ImmutableArray<Outputs.GetServiceCatalogAssociationsFilterResult> filters,

            string id,

            ImmutableArray<Outputs.GetServiceCatalogAssociationsServiceCatalogAssociationCollectionResult> serviceCatalogAssociationCollections,

            string? serviceCatalogAssociationId,

            string? serviceCatalogId)
        {
            EntityId = entityId;
            EntityType = entityType;
            Filters = filters;
            Id = id;
            ServiceCatalogAssociationCollections = serviceCatalogAssociationCollections;
            ServiceCatalogAssociationId = serviceCatalogAssociationId;
            ServiceCatalogId = serviceCatalogId;
        }
    }
}