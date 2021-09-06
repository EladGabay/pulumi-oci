// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package objectstorage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Bucket resource in Oracle Cloud Infrastructure Object Storage service.
//
// Creates a bucket in the given namespace with a bucket name and optional user-defined metadata. Avoid entering
// confidential information in bucket names.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/objectstorage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := objectstorage.NewObjectstorageBucket(ctx, "testBucket", &objectstorage.ObjectstorageBucketArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			Namespace:     pulumi.Any(_var.Bucket_namespace),
// 			AccessType:    pulumi.Any(_var.Bucket_access_type),
// 			AutoTiering:   pulumi.Any(_var.Bucket_auto_tiering),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			KmsKeyId:            pulumi.Any(oci_kms_key.Test_key.Id),
// 			Metadata:            pulumi.Any(_var.Bucket_metadata),
// 			ObjectEventsEnabled: pulumi.Any(_var.Bucket_object_events_enabled),
// 			StorageTier:         pulumi.Any(_var.Bucket_storage_tier),
// 			RetentionRules: objectstorage.ObjectstorageBucketRetentionRuleArray{
// 				&objectstorage.ObjectstorageBucketRetentionRuleArgs{
// 					DisplayName: pulumi.Any(_var.Retention_rule_display_name),
// 					Duration: &objectstorage.ObjectstorageBucketRetentionRuleDurationArgs{
// 						TimeAmount: pulumi.Any(_var.Retention_rule_duration_time_amount),
// 						TimeUnit:   pulumi.Any(_var.Retention_rule_duration_time_unit),
// 					},
// 					TimeRuleLocked: pulumi.Any(_var.Retention_rule_time_rule_locked),
// 				},
// 			},
// 			Versioning: pulumi.Any(_var.Bucket_versioning),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Buckets can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:objectstorage/objectstorageBucket:ObjectstorageBucket test_bucket "n/{namespaceName}/b/{bucketName}"
// ```
type ObjectstorageBucket struct {
	pulumi.CustomResourceState

	// (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations.
	AccessType pulumi.StringPtrOutput `pulumi:"accessType"`
	// The approximate number of objects in the bucket. Count statistics are reported periodically. You will see a lag between what is displayed and the actual object count.
	ApproximateCount pulumi.StringOutput `pulumi:"approximateCount"`
	// The approximate total size in bytes of all objects in the bucket. Size statistics are reported periodically. You will see a lag between what is displayed and the actual size of the bucket.
	ApproximateSize pulumi.StringOutput `pulumi:"approximateSize"`
	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering `Disabled`. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to `InfrequentAccess` are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering pulumi.StringOutput `pulumi:"autoTiering"`
	// The OCID of the bucket which is a Oracle assigned unique identifier for this resource type (bucket). `bucketId` cannot be used for bucket lookup.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// (Updatable) The ID of the compartment in which to create the bucket.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the bucket.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// The entity tag (ETag) for the bucket.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Whether or not this bucket is read only. By default, `isReadOnly` is set to `false`. This will be set to 'true' when this bucket is configured as a destination in a replication policy.
	IsReadOnly pulumi.BoolOutput `pulumi:"isReadOnly"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata pulumi.MapOutput `pulumi:"metadata"`
	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name pulumi.StringOutput `pulumi:"name"`
	// The Object Storage namespace used for the request.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
	ObjectEventsEnabled pulumi.BoolOutput `pulumi:"objectEventsEnabled"`
	// The entity tag (ETag) for the live object lifecycle policy on the bucket.
	ObjectLifecyclePolicyEtag pulumi.StringOutput `pulumi:"objectLifecyclePolicyEtag"`
	// Whether or not this bucket is a replication source. By default, `replicationEnabled` is set to `false`. This will be set to 'true' when you create a replication policy for the bucket.
	ReplicationEnabled pulumi.BoolOutput `pulumi:"replicationEnabled"`
	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules ObjectstorageBucketRetentionRuleArrayOutput `pulumi:"retentionRules"`
	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier pulumi.StringOutput `pulumi:"storageTier"`
	// The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning `Disabled`. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning pulumi.StringOutput `pulumi:"versioning"`
}

// NewObjectstorageBucket registers a new resource with the given unique name, arguments, and options.
func NewObjectstorageBucket(ctx *pulumi.Context,
	name string, args *ObjectstorageBucketArgs, opts ...pulumi.ResourceOption) (*ObjectstorageBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	var resource ObjectstorageBucket
	err := ctx.RegisterResource("oci:objectstorage/objectstorageBucket:ObjectstorageBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectstorageBucket gets an existing ObjectstorageBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectstorageBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectstorageBucketState, opts ...pulumi.ResourceOption) (*ObjectstorageBucket, error) {
	var resource ObjectstorageBucket
	err := ctx.ReadResource("oci:objectstorage/objectstorageBucket:ObjectstorageBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectstorageBucket resources.
type objectstorageBucketState struct {
	// (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations.
	AccessType *string `pulumi:"accessType"`
	// The approximate number of objects in the bucket. Count statistics are reported periodically. You will see a lag between what is displayed and the actual object count.
	ApproximateCount *string `pulumi:"approximateCount"`
	// The approximate total size in bytes of all objects in the bucket. Size statistics are reported periodically. You will see a lag between what is displayed and the actual size of the bucket.
	ApproximateSize *string `pulumi:"approximateSize"`
	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering `Disabled`. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to `InfrequentAccess` are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering *string `pulumi:"autoTiering"`
	// The OCID of the bucket which is a Oracle assigned unique identifier for this resource type (bucket). `bucketId` cannot be used for bucket lookup.
	BucketId *string `pulumi:"bucketId"`
	// (Updatable) The ID of the compartment in which to create the bucket.
	CompartmentId *string `pulumi:"compartmentId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the bucket.
	CreatedBy *string `pulumi:"createdBy"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The entity tag (ETag) for the bucket.
	Etag *string `pulumi:"etag"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Whether or not this bucket is read only. By default, `isReadOnly` is set to `false`. This will be set to 'true' when this bucket is configured as a destination in a replication policy.
	IsReadOnly *bool `pulumi:"isReadOnly"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name *string `pulumi:"name"`
	// The Object Storage namespace used for the request.
	Namespace *string `pulumi:"namespace"`
	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
	ObjectEventsEnabled *bool `pulumi:"objectEventsEnabled"`
	// The entity tag (ETag) for the live object lifecycle policy on the bucket.
	ObjectLifecyclePolicyEtag *string `pulumi:"objectLifecyclePolicyEtag"`
	// Whether or not this bucket is a replication source. By default, `replicationEnabled` is set to `false`. This will be set to 'true' when you create a replication policy for the bucket.
	ReplicationEnabled *bool `pulumi:"replicationEnabled"`
	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules []ObjectstorageBucketRetentionRule `pulumi:"retentionRules"`
	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier *string `pulumi:"storageTier"`
	// The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
	TimeCreated *string `pulumi:"timeCreated"`
	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning `Disabled`. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning *string `pulumi:"versioning"`
}

type ObjectstorageBucketState struct {
	// (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations.
	AccessType pulumi.StringPtrInput
	// The approximate number of objects in the bucket. Count statistics are reported periodically. You will see a lag between what is displayed and the actual object count.
	ApproximateCount pulumi.StringPtrInput
	// The approximate total size in bytes of all objects in the bucket. Size statistics are reported periodically. You will see a lag between what is displayed and the actual size of the bucket.
	ApproximateSize pulumi.StringPtrInput
	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering `Disabled`. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to `InfrequentAccess` are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering pulumi.StringPtrInput
	// The OCID of the bucket which is a Oracle assigned unique identifier for this resource type (bucket). `bucketId` cannot be used for bucket lookup.
	BucketId pulumi.StringPtrInput
	// (Updatable) The ID of the compartment in which to create the bucket.
	CompartmentId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the bucket.
	CreatedBy pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// The entity tag (ETag) for the bucket.
	Etag pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// Whether or not this bucket is read only. By default, `isReadOnly` is set to `false`. This will be set to 'true' when this bucket is configured as a destination in a replication policy.
	IsReadOnly pulumi.BoolPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	KmsKeyId pulumi.StringPtrInput
	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata pulumi.MapInput
	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name pulumi.StringPtrInput
	// The Object Storage namespace used for the request.
	Namespace pulumi.StringPtrInput
	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
	ObjectEventsEnabled pulumi.BoolPtrInput
	// The entity tag (ETag) for the live object lifecycle policy on the bucket.
	ObjectLifecyclePolicyEtag pulumi.StringPtrInput
	// Whether or not this bucket is a replication source. By default, `replicationEnabled` is set to `false`. This will be set to 'true' when you create a replication policy for the bucket.
	ReplicationEnabled pulumi.BoolPtrInput
	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules ObjectstorageBucketRetentionRuleArrayInput
	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier pulumi.StringPtrInput
	// The date and time the bucket was created, as described in [RFC 2616](https://tools.ietf.org/html/rfc2616#section-14.29).
	TimeCreated pulumi.StringPtrInput
	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning `Disabled`. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning pulumi.StringPtrInput
}

func (ObjectstorageBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectstorageBucketState)(nil)).Elem()
}

type objectstorageBucketArgs struct {
	// (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations.
	AccessType *string `pulumi:"accessType"`
	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering `Disabled`. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to `InfrequentAccess` are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering *string `pulumi:"autoTiering"`
	// (Updatable) The ID of the compartment in which to create the bucket.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name *string `pulumi:"name"`
	// The Object Storage namespace used for the request.
	Namespace string `pulumi:"namespace"`
	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
	ObjectEventsEnabled *bool `pulumi:"objectEventsEnabled"`
	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules []ObjectstorageBucketRetentionRule `pulumi:"retentionRules"`
	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier *string `pulumi:"storageTier"`
	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning `Disabled`. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning *string `pulumi:"versioning"`
}

// The set of arguments for constructing a ObjectstorageBucket resource.
type ObjectstorageBucketArgs struct {
	// (Updatable) The type of public access enabled on this bucket. A bucket is set to `NoPublicAccess` by default, which only allows an authenticated caller to access the bucket and its contents. When `ObjectRead` is enabled on the bucket, public access is allowed for the `GetObject`, `HeadObject`, and `ListObjects` operations. When `ObjectReadWithoutList` is enabled on the bucket, public access is allowed for the `GetObject` and `HeadObject` operations.
	AccessType pulumi.StringPtrInput
	// (Updatable) Set the auto tiering status on the bucket. By default, a bucket is created with auto tiering `Disabled`. Use this option to enable auto tiering during bucket creation. Objects in a bucket with auto tiering set to `InfrequentAccess` are transitioned automatically between the 'Standard' and 'InfrequentAccess' tiers based on the access pattern of the objects.
	AutoTiering pulumi.StringPtrInput
	// (Updatable) The ID of the compartment in which to create the bucket.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.
	KmsKeyId pulumi.StringPtrInput
	// (Updatable) Arbitrary string, up to 4KB, of keys and values for user-defined metadata.
	Metadata pulumi.MapInput
	// The name of the bucket. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. Bucket names must be unique within an Object Storage namespace. Avoid entering confidential information. example: Example: my-new-bucket1
	Name pulumi.StringPtrInput
	// The Object Storage namespace used for the request.
	Namespace pulumi.StringInput
	// (Updatable) Whether or not events are emitted for object state changes in this bucket. By default, `objectEventsEnabled` is set to `false`. Set `objectEventsEnabled` to `true` to emit events for object state changes. For more information about events, see [Overview of Events](https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
	ObjectEventsEnabled pulumi.BoolPtrInput
	// (Updatable) Creates a new retention rule in the specified bucket. The new rule will take effect typically within 30 seconds. Note that a maximum of 100 rules are supported on a bucket.
	RetentionRules ObjectstorageBucketRetentionRuleArrayInput
	// The type of storage tier of this bucket. A bucket is set to 'Standard' tier by default, which means the bucket will be put in the standard storage tier. When 'Archive' tier type is set explicitly, the bucket is put in the Archive Storage tier. The 'storageTier' property is immutable after bucket is created.
	StorageTier pulumi.StringPtrInput
	// (Updatable) Set the versioning status on the bucket. By default, a bucket is created with versioning `Disabled`. Use this option to enable versioning during bucket creation. Objects in a version enabled bucket are protected from overwrites and deletions. Previous versions of the same object will be available in the bucket. Allowed Create values: Enabled, Disabled. Allowed Update values: Enabled, Suspended.
	Versioning pulumi.StringPtrInput
}

func (ObjectstorageBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectstorageBucketArgs)(nil)).Elem()
}

type ObjectstorageBucketInput interface {
	pulumi.Input

	ToObjectstorageBucketOutput() ObjectstorageBucketOutput
	ToObjectstorageBucketOutputWithContext(ctx context.Context) ObjectstorageBucketOutput
}

func (*ObjectstorageBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectstorageBucket)(nil))
}

func (i *ObjectstorageBucket) ToObjectstorageBucketOutput() ObjectstorageBucketOutput {
	return i.ToObjectstorageBucketOutputWithContext(context.Background())
}

func (i *ObjectstorageBucket) ToObjectstorageBucketOutputWithContext(ctx context.Context) ObjectstorageBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectstorageBucketOutput)
}

func (i *ObjectstorageBucket) ToObjectstorageBucketPtrOutput() ObjectstorageBucketPtrOutput {
	return i.ToObjectstorageBucketPtrOutputWithContext(context.Background())
}

func (i *ObjectstorageBucket) ToObjectstorageBucketPtrOutputWithContext(ctx context.Context) ObjectstorageBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectstorageBucketPtrOutput)
}

type ObjectstorageBucketPtrInput interface {
	pulumi.Input

	ToObjectstorageBucketPtrOutput() ObjectstorageBucketPtrOutput
	ToObjectstorageBucketPtrOutputWithContext(ctx context.Context) ObjectstorageBucketPtrOutput
}

type objectstorageBucketPtrType ObjectstorageBucketArgs

func (*objectstorageBucketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectstorageBucket)(nil))
}

func (i *objectstorageBucketPtrType) ToObjectstorageBucketPtrOutput() ObjectstorageBucketPtrOutput {
	return i.ToObjectstorageBucketPtrOutputWithContext(context.Background())
}

func (i *objectstorageBucketPtrType) ToObjectstorageBucketPtrOutputWithContext(ctx context.Context) ObjectstorageBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectstorageBucketPtrOutput)
}

// ObjectstorageBucketArrayInput is an input type that accepts ObjectstorageBucketArray and ObjectstorageBucketArrayOutput values.
// You can construct a concrete instance of `ObjectstorageBucketArrayInput` via:
//
//          ObjectstorageBucketArray{ ObjectstorageBucketArgs{...} }
type ObjectstorageBucketArrayInput interface {
	pulumi.Input

	ToObjectstorageBucketArrayOutput() ObjectstorageBucketArrayOutput
	ToObjectstorageBucketArrayOutputWithContext(context.Context) ObjectstorageBucketArrayOutput
}

type ObjectstorageBucketArray []ObjectstorageBucketInput

func (ObjectstorageBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ObjectstorageBucket)(nil)).Elem()
}

func (i ObjectstorageBucketArray) ToObjectstorageBucketArrayOutput() ObjectstorageBucketArrayOutput {
	return i.ToObjectstorageBucketArrayOutputWithContext(context.Background())
}

func (i ObjectstorageBucketArray) ToObjectstorageBucketArrayOutputWithContext(ctx context.Context) ObjectstorageBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectstorageBucketArrayOutput)
}

// ObjectstorageBucketMapInput is an input type that accepts ObjectstorageBucketMap and ObjectstorageBucketMapOutput values.
// You can construct a concrete instance of `ObjectstorageBucketMapInput` via:
//
//          ObjectstorageBucketMap{ "key": ObjectstorageBucketArgs{...} }
type ObjectstorageBucketMapInput interface {
	pulumi.Input

	ToObjectstorageBucketMapOutput() ObjectstorageBucketMapOutput
	ToObjectstorageBucketMapOutputWithContext(context.Context) ObjectstorageBucketMapOutput
}

type ObjectstorageBucketMap map[string]ObjectstorageBucketInput

func (ObjectstorageBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ObjectstorageBucket)(nil)).Elem()
}

func (i ObjectstorageBucketMap) ToObjectstorageBucketMapOutput() ObjectstorageBucketMapOutput {
	return i.ToObjectstorageBucketMapOutputWithContext(context.Background())
}

func (i ObjectstorageBucketMap) ToObjectstorageBucketMapOutputWithContext(ctx context.Context) ObjectstorageBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectstorageBucketMapOutput)
}

type ObjectstorageBucketOutput struct {
	*pulumi.OutputState
}

func (ObjectstorageBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ObjectstorageBucket)(nil))
}

func (o ObjectstorageBucketOutput) ToObjectstorageBucketOutput() ObjectstorageBucketOutput {
	return o
}

func (o ObjectstorageBucketOutput) ToObjectstorageBucketOutputWithContext(ctx context.Context) ObjectstorageBucketOutput {
	return o
}

func (o ObjectstorageBucketOutput) ToObjectstorageBucketPtrOutput() ObjectstorageBucketPtrOutput {
	return o.ToObjectstorageBucketPtrOutputWithContext(context.Background())
}

func (o ObjectstorageBucketOutput) ToObjectstorageBucketPtrOutputWithContext(ctx context.Context) ObjectstorageBucketPtrOutput {
	return o.ApplyT(func(v ObjectstorageBucket) *ObjectstorageBucket {
		return &v
	}).(ObjectstorageBucketPtrOutput)
}

type ObjectstorageBucketPtrOutput struct {
	*pulumi.OutputState
}

func (ObjectstorageBucketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectstorageBucket)(nil))
}

func (o ObjectstorageBucketPtrOutput) ToObjectstorageBucketPtrOutput() ObjectstorageBucketPtrOutput {
	return o
}

func (o ObjectstorageBucketPtrOutput) ToObjectstorageBucketPtrOutputWithContext(ctx context.Context) ObjectstorageBucketPtrOutput {
	return o
}

type ObjectstorageBucketArrayOutput struct{ *pulumi.OutputState }

func (ObjectstorageBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ObjectstorageBucket)(nil))
}

func (o ObjectstorageBucketArrayOutput) ToObjectstorageBucketArrayOutput() ObjectstorageBucketArrayOutput {
	return o
}

func (o ObjectstorageBucketArrayOutput) ToObjectstorageBucketArrayOutputWithContext(ctx context.Context) ObjectstorageBucketArrayOutput {
	return o
}

func (o ObjectstorageBucketArrayOutput) Index(i pulumi.IntInput) ObjectstorageBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ObjectstorageBucket {
		return vs[0].([]ObjectstorageBucket)[vs[1].(int)]
	}).(ObjectstorageBucketOutput)
}

type ObjectstorageBucketMapOutput struct{ *pulumi.OutputState }

func (ObjectstorageBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ObjectstorageBucket)(nil))
}

func (o ObjectstorageBucketMapOutput) ToObjectstorageBucketMapOutput() ObjectstorageBucketMapOutput {
	return o
}

func (o ObjectstorageBucketMapOutput) ToObjectstorageBucketMapOutputWithContext(ctx context.Context) ObjectstorageBucketMapOutput {
	return o
}

func (o ObjectstorageBucketMapOutput) MapIndex(k pulumi.StringInput) ObjectstorageBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ObjectstorageBucket {
		return vs[0].(map[string]ObjectstorageBucket)[vs[1].(string)]
	}).(ObjectstorageBucketOutput)
}

func init() {
	pulumi.RegisterOutputType(ObjectstorageBucketOutput{})
	pulumi.RegisterOutputType(ObjectstorageBucketPtrOutput{})
	pulumi.RegisterOutputType(ObjectstorageBucketArrayOutput{})
	pulumi.RegisterOutputType(ObjectstorageBucketMapOutput{})
}