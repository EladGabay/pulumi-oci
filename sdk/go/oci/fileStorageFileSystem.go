// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the File System resource in Oracle Cloud Infrastructure File Storage service.
//
// Creates a new file system in the specified compartment and
// availability domain. Instances can mount file systems in
// another availability domain, but doing so might increase
// latency when compared to mounting instances in the same
// availability domain.
//
// After you create a file system, you can associate it with a mount
// target. Instances can then mount the file system by connecting to the
// mount target's IP address. You can associate a file system with
// more than one mount target at a time.
//
// For information about access control and compartments, see
// [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
//
// For information about Network Security Groups access control, see
// [Network Security Groups](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm).
//
// For information about availability domains, see [Regions and
// Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
// To get a list of availability domains, use the
// `ListAvailabilityDomains` operation in the Identity and Access
// Management Service API.
//
// All Oracle Cloud Infrastructure resources, including
// file systems, get an Oracle-assigned, unique ID called an Oracle
// Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).\
// When you create a resource, you can find its OCID in the response.
// You can also retrieve a resource's OCID by using a List API operation on that resource
// type or by viewing the resource in the Console.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := oci.NewFileStorageFileSystem(ctx, "testFileSystem", &oci.FileStorageFileSystemArgs{
// 			AvailabilityDomain: pulumi.Any(_var.File_system_availability_domain),
// 			CompartmentId:      pulumi.Any(_var.Compartment_id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName: pulumi.Any(_var.File_system_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			KmsKeyId:         pulumi.Any(oci_kms_key.Test_key.Id),
// 			SourceSnapshotId: pulumi.Any(oci_file_storage_snapshot.Test_snapshot.Id),
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
// FileSystems can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:index/fileStorageFileSystem:FileStorageFileSystem test_file_system "id"
// ```
type FileStorageFileSystem struct {
	pulumi.CustomResourceState

	// The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringOutput `pulumi:"availabilityDomain"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system`
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// Specifies whether the file system has been cloned. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	IsCloneParent pulumi.BoolOutput `pulumi:"isCloneParent"`
	// Specifies whether the data has finished copying from the source to the clone. Hydration can take up to several hours to complete depending on the size of the source. The source and clone remain available during hydration, but there may be some performance impact. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm#hydration).
	IsHydrated pulumi.BoolOutput `pulumi:"isHydrated"`
	// (Updatable) The OCID of KMS key used to encrypt the encryption keys associated with this file system. May be unset as a blank or deleted from the configuration to remove the KMS key.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Additional information about the current 'lifecycleState'.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. For more information, see [File System Usage and Metering](https://docs.cloud.oracle.com/iaas/Content/File/Concepts/FSutilization.htm).
	MeteredBytes pulumi.StringOutput `pulumi:"meteredBytes"`
	// Source information for the file system.
	SourceDetails FileStorageFileSystemSourceDetailsOutput `pulumi:"sourceDetails"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	SourceSnapshotId pulumi.StringOutput `pulumi:"sourceSnapshotId"`
	// The current state of the file system.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewFileStorageFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileStorageFileSystem(ctx *pulumi.Context,
	name string, args *FileStorageFileSystemArgs, opts ...pulumi.ResourceOption) (*FileStorageFileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityDomain == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityDomain'")
	}
	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	var resource FileStorageFileSystem
	err := ctx.RegisterResource("oci:index/fileStorageFileSystem:FileStorageFileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileStorageFileSystem gets an existing FileStorageFileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileStorageFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileStorageFileSystemState, opts ...pulumi.ResourceOption) (*FileStorageFileSystem, error) {
	var resource FileStorageFileSystem
	err := ctx.ReadResource("oci:index/fileStorageFileSystem:FileStorageFileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileStorageFileSystem resources.
type fileStorageFileSystemState struct {
	// The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system`
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Specifies whether the file system has been cloned. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	IsCloneParent *bool `pulumi:"isCloneParent"`
	// Specifies whether the data has finished copying from the source to the clone. Hydration can take up to several hours to complete depending on the size of the source. The source and clone remain available during hydration, but there may be some performance impact. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm#hydration).
	IsHydrated *bool `pulumi:"isHydrated"`
	// (Updatable) The OCID of KMS key used to encrypt the encryption keys associated with this file system. May be unset as a blank or deleted from the configuration to remove the KMS key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Additional information about the current 'lifecycleState'.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. For more information, see [File System Usage and Metering](https://docs.cloud.oracle.com/iaas/Content/File/Concepts/FSutilization.htm).
	MeteredBytes *string `pulumi:"meteredBytes"`
	// Source information for the file system.
	SourceDetails *FileStorageFileSystemSourceDetails `pulumi:"sourceDetails"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
	// The current state of the file system.
	State *string `pulumi:"state"`
	// The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
}

type FileStorageFileSystemState struct {
	// The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system`
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// Specifies whether the file system has been cloned. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	IsCloneParent pulumi.BoolPtrInput
	// Specifies whether the data has finished copying from the source to the clone. Hydration can take up to several hours to complete depending on the size of the source. The source and clone remain available during hydration, but there may be some performance impact. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm#hydration).
	IsHydrated pulumi.BoolPtrInput
	// (Updatable) The OCID of KMS key used to encrypt the encryption keys associated with this file system. May be unset as a blank or deleted from the configuration to remove the KMS key.
	KmsKeyId pulumi.StringPtrInput
	// Additional information about the current 'lifecycleState'.
	LifecycleDetails pulumi.StringPtrInput
	// The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. For more information, see [File System Usage and Metering](https://docs.cloud.oracle.com/iaas/Content/File/Concepts/FSutilization.htm).
	MeteredBytes pulumi.StringPtrInput
	// Source information for the file system.
	SourceDetails FileStorageFileSystemSourceDetailsPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	SourceSnapshotId pulumi.StringPtrInput
	// The current state of the file system.
	State pulumi.StringPtrInput
	// The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
}

func (FileStorageFileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileStorageFileSystemState)(nil)).Elem()
}

type fileStorageFileSystemArgs struct {
	// The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system`
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The OCID of KMS key used to encrypt the encryption keys associated with this file system. May be unset as a blank or deleted from the configuration to remove the KMS key.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	SourceSnapshotId *string `pulumi:"sourceSnapshotId"`
}

// The set of arguments for constructing a FileStorageFileSystem resource.
type FileStorageFileSystemArgs struct {
	// The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system`
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The OCID of KMS key used to encrypt the encryption keys associated with this file system. May be unset as a blank or deleted from the configuration to remove the KMS key.
	KmsKeyId pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningafilesystem.htm).
	SourceSnapshotId pulumi.StringPtrInput
}

func (FileStorageFileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileStorageFileSystemArgs)(nil)).Elem()
}

type FileStorageFileSystemInput interface {
	pulumi.Input

	ToFileStorageFileSystemOutput() FileStorageFileSystemOutput
	ToFileStorageFileSystemOutputWithContext(ctx context.Context) FileStorageFileSystemOutput
}

func (*FileStorageFileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((*FileStorageFileSystem)(nil))
}

func (i *FileStorageFileSystem) ToFileStorageFileSystemOutput() FileStorageFileSystemOutput {
	return i.ToFileStorageFileSystemOutputWithContext(context.Background())
}

func (i *FileStorageFileSystem) ToFileStorageFileSystemOutputWithContext(ctx context.Context) FileStorageFileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileStorageFileSystemOutput)
}

func (i *FileStorageFileSystem) ToFileStorageFileSystemPtrOutput() FileStorageFileSystemPtrOutput {
	return i.ToFileStorageFileSystemPtrOutputWithContext(context.Background())
}

func (i *FileStorageFileSystem) ToFileStorageFileSystemPtrOutputWithContext(ctx context.Context) FileStorageFileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileStorageFileSystemPtrOutput)
}

type FileStorageFileSystemPtrInput interface {
	pulumi.Input

	ToFileStorageFileSystemPtrOutput() FileStorageFileSystemPtrOutput
	ToFileStorageFileSystemPtrOutputWithContext(ctx context.Context) FileStorageFileSystemPtrOutput
}

type fileStorageFileSystemPtrType FileStorageFileSystemArgs

func (*fileStorageFileSystemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileStorageFileSystem)(nil))
}

func (i *fileStorageFileSystemPtrType) ToFileStorageFileSystemPtrOutput() FileStorageFileSystemPtrOutput {
	return i.ToFileStorageFileSystemPtrOutputWithContext(context.Background())
}

func (i *fileStorageFileSystemPtrType) ToFileStorageFileSystemPtrOutputWithContext(ctx context.Context) FileStorageFileSystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileStorageFileSystemPtrOutput)
}

// FileStorageFileSystemArrayInput is an input type that accepts FileStorageFileSystemArray and FileStorageFileSystemArrayOutput values.
// You can construct a concrete instance of `FileStorageFileSystemArrayInput` via:
//
//          FileStorageFileSystemArray{ FileStorageFileSystemArgs{...} }
type FileStorageFileSystemArrayInput interface {
	pulumi.Input

	ToFileStorageFileSystemArrayOutput() FileStorageFileSystemArrayOutput
	ToFileStorageFileSystemArrayOutputWithContext(context.Context) FileStorageFileSystemArrayOutput
}

type FileStorageFileSystemArray []FileStorageFileSystemInput

func (FileStorageFileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileStorageFileSystem)(nil)).Elem()
}

func (i FileStorageFileSystemArray) ToFileStorageFileSystemArrayOutput() FileStorageFileSystemArrayOutput {
	return i.ToFileStorageFileSystemArrayOutputWithContext(context.Background())
}

func (i FileStorageFileSystemArray) ToFileStorageFileSystemArrayOutputWithContext(ctx context.Context) FileStorageFileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileStorageFileSystemArrayOutput)
}

// FileStorageFileSystemMapInput is an input type that accepts FileStorageFileSystemMap and FileStorageFileSystemMapOutput values.
// You can construct a concrete instance of `FileStorageFileSystemMapInput` via:
//
//          FileStorageFileSystemMap{ "key": FileStorageFileSystemArgs{...} }
type FileStorageFileSystemMapInput interface {
	pulumi.Input

	ToFileStorageFileSystemMapOutput() FileStorageFileSystemMapOutput
	ToFileStorageFileSystemMapOutputWithContext(context.Context) FileStorageFileSystemMapOutput
}

type FileStorageFileSystemMap map[string]FileStorageFileSystemInput

func (FileStorageFileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileStorageFileSystem)(nil)).Elem()
}

func (i FileStorageFileSystemMap) ToFileStorageFileSystemMapOutput() FileStorageFileSystemMapOutput {
	return i.ToFileStorageFileSystemMapOutputWithContext(context.Background())
}

func (i FileStorageFileSystemMap) ToFileStorageFileSystemMapOutputWithContext(ctx context.Context) FileStorageFileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileStorageFileSystemMapOutput)
}

type FileStorageFileSystemOutput struct {
	*pulumi.OutputState
}

func (FileStorageFileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileStorageFileSystem)(nil))
}

func (o FileStorageFileSystemOutput) ToFileStorageFileSystemOutput() FileStorageFileSystemOutput {
	return o
}

func (o FileStorageFileSystemOutput) ToFileStorageFileSystemOutputWithContext(ctx context.Context) FileStorageFileSystemOutput {
	return o
}

func (o FileStorageFileSystemOutput) ToFileStorageFileSystemPtrOutput() FileStorageFileSystemPtrOutput {
	return o.ToFileStorageFileSystemPtrOutputWithContext(context.Background())
}

func (o FileStorageFileSystemOutput) ToFileStorageFileSystemPtrOutputWithContext(ctx context.Context) FileStorageFileSystemPtrOutput {
	return o.ApplyT(func(v FileStorageFileSystem) *FileStorageFileSystem {
		return &v
	}).(FileStorageFileSystemPtrOutput)
}

type FileStorageFileSystemPtrOutput struct {
	*pulumi.OutputState
}

func (FileStorageFileSystemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileStorageFileSystem)(nil))
}

func (o FileStorageFileSystemPtrOutput) ToFileStorageFileSystemPtrOutput() FileStorageFileSystemPtrOutput {
	return o
}

func (o FileStorageFileSystemPtrOutput) ToFileStorageFileSystemPtrOutputWithContext(ctx context.Context) FileStorageFileSystemPtrOutput {
	return o
}

type FileStorageFileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileStorageFileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FileStorageFileSystem)(nil))
}

func (o FileStorageFileSystemArrayOutput) ToFileStorageFileSystemArrayOutput() FileStorageFileSystemArrayOutput {
	return o
}

func (o FileStorageFileSystemArrayOutput) ToFileStorageFileSystemArrayOutputWithContext(ctx context.Context) FileStorageFileSystemArrayOutput {
	return o
}

func (o FileStorageFileSystemArrayOutput) Index(i pulumi.IntInput) FileStorageFileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FileStorageFileSystem {
		return vs[0].([]FileStorageFileSystem)[vs[1].(int)]
	}).(FileStorageFileSystemOutput)
}

type FileStorageFileSystemMapOutput struct{ *pulumi.OutputState }

func (FileStorageFileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FileStorageFileSystem)(nil))
}

func (o FileStorageFileSystemMapOutput) ToFileStorageFileSystemMapOutput() FileStorageFileSystemMapOutput {
	return o
}

func (o FileStorageFileSystemMapOutput) ToFileStorageFileSystemMapOutputWithContext(ctx context.Context) FileStorageFileSystemMapOutput {
	return o
}

func (o FileStorageFileSystemMapOutput) MapIndex(k pulumi.StringInput) FileStorageFileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FileStorageFileSystem {
		return vs[0].(map[string]FileStorageFileSystem)[vs[1].(string)]
	}).(FileStorageFileSystemOutput)
}

func init() {
	pulumi.RegisterOutputType(FileStorageFileSystemOutput{})
	pulumi.RegisterOutputType(FileStorageFileSystemPtrOutput{})
	pulumi.RegisterOutputType(FileStorageFileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileStorageFileSystemMapOutput{})
}