// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Boot Volume resource in Oracle Cloud Infrastructure Core service.
//
// Creates a new boot volume in the specified compartment from an existing boot volume or a boot volume backup.
// For general information about boot volumes, see [Boot Volumes](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumes.htm).
// You may optionally specify a *display name* for the volume, which is simply a friendly name or
// description. It does not have to be unique, and you can change it. Avoid entering confidential information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.NewBootVolume(ctx, "testBootVolume", &core.BootVolumeArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			SourceDetails: &core.BootVolumeSourceDetailsArgs{
// 				Id:   pulumi.Any(_var.Boot_volume_source_details_id),
// 				Type: pulumi.Any(_var.Boot_volume_source_details_type),
// 			},
// 			AvailabilityDomain: pulumi.Any(_var.Boot_volume_availability_domain),
// 			BackupPolicyId:     pulumi.Any(data.Oci_core_volume_backup_policies.Test_volume_backup_policies.Volume_backup_policies[0].Id),
// 			BootVolumeReplicas: core.BootVolumeBootVolumeReplicaArray{
// 				&core.BootVolumeBootVolumeReplicaArgs{
// 					AvailabilityDomain: pulumi.Any(_var.Boot_volume_boot_volume_replicas_availability_domain),
// 					DisplayName:        pulumi.Any(_var.Boot_volume_boot_volume_replicas_display_name),
// 				},
// 			},
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName: pulumi.Any(_var.Boot_volume_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			IsAutoTuneEnabled:          pulumi.Any(_var.Boot_volume_is_auto_tune_enabled),
// 			KmsKeyId:                   pulumi.Any(oci_kms_key.Test_key.Id),
// 			SizeInGbs:                  pulumi.Any(_var.Boot_volume_size_in_gbs),
// 			VpusPerGb:                  pulumi.Any(_var.Boot_volume_vpus_per_gb),
// 			BootVolumeReplicasDeletion: pulumi.Bool(true),
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
// BootVolumes can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/bootVolume:BootVolume test_boot_volume "id"
// ```
type BootVolume struct {
	pulumi.CustomResourceState

	// The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle.
	AutoTunedVpusPerGb pulumi.StringOutput `pulumi:"autoTunedVpusPerGb"`
	// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringOutput `pulumi:"availabilityDomain"`
	// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
	//
	// Deprecated: The 'backup_policy_id' field has been deprecated. Please use the 'oci_core_volume_backup_policy_assignment' resource instead.
	BackupPolicyId pulumi.StringOutput `pulumi:"backupPolicyId"`
	// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
	BootVolumeReplicas BootVolumeBootVolumeReplicaArrayOutput `pulumi:"bootVolumeReplicas"`
	// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `bootVolumeReplicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `bootVolumeReplicas` again.
	BootVolumeReplicasDeletion pulumi.BoolPtrOutput `pulumi:"bootVolumeReplicasDeletion"`
	// (Updatable) The OCID of the compartment that contains the boot volume.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The image OCID used to create the boot volume.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
	IsAutoTuneEnabled pulumi.BoolOutput `pulumi:"isAutoTuneEnabled"`
	// Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
	IsHydrated pulumi.BoolOutput `pulumi:"isHydrated"`
	// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	KmsKeyId pulumi.StringOutput `pulumi:"kmsKeyId"`
	// (Updatable) The size of the volume in GBs.
	SizeInGbs pulumi.StringOutput `pulumi:"sizeInGbs"`
	// The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `sizeInGbs`.
	SizeInMbs     pulumi.StringOutput           `pulumi:"sizeInMbs"`
	SourceDetails BootVolumeSourceDetailsOutput `pulumi:"sourceDetails"`
	// The current state of a boot volume.
	State pulumi.StringOutput `pulumi:"state"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	SystemTags pulumi.MapOutput `pulumi:"systemTags"`
	// The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The OCID of the source volume group.
	VolumeGroupId pulumi.StringOutput `pulumi:"volumeGroupId"`
	// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
	VpusPerGb pulumi.StringOutput `pulumi:"vpusPerGb"`
}

// NewBootVolume registers a new resource with the given unique name, arguments, and options.
func NewBootVolume(ctx *pulumi.Context,
	name string, args *BootVolumeArgs, opts ...pulumi.ResourceOption) (*BootVolume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityDomain == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityDomain'")
	}
	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.SourceDetails == nil {
		return nil, errors.New("invalid value for required argument 'SourceDetails'")
	}
	var resource BootVolume
	err := ctx.RegisterResource("oci:core/bootVolume:BootVolume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBootVolume gets an existing BootVolume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBootVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BootVolumeState, opts ...pulumi.ResourceOption) (*BootVolume, error) {
	var resource BootVolume
	err := ctx.ReadResource("oci:core/bootVolume:BootVolume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BootVolume resources.
type bootVolumeState struct {
	// The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle.
	AutoTunedVpusPerGb *string `pulumi:"autoTunedVpusPerGb"`
	// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
	//
	// Deprecated: The 'backup_policy_id' field has been deprecated. Please use the 'oci_core_volume_backup_policy_assignment' resource instead.
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
	BootVolumeReplicas []BootVolumeBootVolumeReplica `pulumi:"bootVolumeReplicas"`
	// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `bootVolumeReplicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `bootVolumeReplicas` again.
	BootVolumeReplicasDeletion *bool `pulumi:"bootVolumeReplicasDeletion"`
	// (Updatable) The OCID of the compartment that contains the boot volume.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The image OCID used to create the boot volume.
	ImageId *string `pulumi:"imageId"`
	// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
	IsAutoTuneEnabled *bool `pulumi:"isAutoTuneEnabled"`
	// Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
	IsHydrated *bool `pulumi:"isHydrated"`
	// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// (Updatable) The size of the volume in GBs.
	SizeInGbs *string `pulumi:"sizeInGbs"`
	// The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `sizeInGbs`.
	SizeInMbs     *string                  `pulumi:"sizeInMbs"`
	SourceDetails *BootVolumeSourceDetails `pulumi:"sourceDetails"`
	// The current state of a boot volume.
	State *string `pulumi:"state"`
	// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated *string `pulumi:"timeCreated"`
	// The OCID of the source volume group.
	VolumeGroupId *string `pulumi:"volumeGroupId"`
	// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
	VpusPerGb *string `pulumi:"vpusPerGb"`
}

type BootVolumeState struct {
	// The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle.
	AutoTunedVpusPerGb pulumi.StringPtrInput
	// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringPtrInput
	// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
	//
	// Deprecated: The 'backup_policy_id' field has been deprecated. Please use the 'oci_core_volume_backup_policy_assignment' resource instead.
	BackupPolicyId pulumi.StringPtrInput
	// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
	BootVolumeReplicas BootVolumeBootVolumeReplicaArrayInput
	// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `bootVolumeReplicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `bootVolumeReplicas` again.
	BootVolumeReplicasDeletion pulumi.BoolPtrInput
	// (Updatable) The OCID of the compartment that contains the boot volume.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The image OCID used to create the boot volume.
	ImageId pulumi.StringPtrInput
	// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
	IsAutoTuneEnabled pulumi.BoolPtrInput
	// Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.
	IsHydrated pulumi.BoolPtrInput
	// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	KmsKeyId pulumi.StringPtrInput
	// (Updatable) The size of the volume in GBs.
	SizeInGbs pulumi.StringPtrInput
	// The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `sizeInGbs`.
	SizeInMbs     pulumi.StringPtrInput
	SourceDetails BootVolumeSourceDetailsPtrInput
	// The current state of a boot volume.
	State pulumi.StringPtrInput
	// System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	SystemTags pulumi.MapInput
	// The date and time the boot volume was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
	TimeCreated pulumi.StringPtrInput
	// The OCID of the source volume group.
	VolumeGroupId pulumi.StringPtrInput
	// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
	VpusPerGb pulumi.StringPtrInput
}

func (BootVolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*bootVolumeState)(nil)).Elem()
}

type bootVolumeArgs struct {
	// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain string `pulumi:"availabilityDomain"`
	// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
	//
	// Deprecated: The 'backup_policy_id' field has been deprecated. Please use the 'oci_core_volume_backup_policy_assignment' resource instead.
	BackupPolicyId *string `pulumi:"backupPolicyId"`
	// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
	BootVolumeReplicas []BootVolumeBootVolumeReplica `pulumi:"bootVolumeReplicas"`
	// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `bootVolumeReplicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `bootVolumeReplicas` again.
	BootVolumeReplicasDeletion *bool `pulumi:"bootVolumeReplicasDeletion"`
	// (Updatable) The OCID of the compartment that contains the boot volume.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
	IsAutoTuneEnabled *bool `pulumi:"isAutoTuneEnabled"`
	// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// (Updatable) The size of the volume in GBs.
	SizeInGbs     *string                 `pulumi:"sizeInGbs"`
	SourceDetails BootVolumeSourceDetails `pulumi:"sourceDetails"`
	// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
	VpusPerGb *string `pulumi:"vpusPerGb"`
}

// The set of arguments for constructing a BootVolume resource.
type BootVolumeArgs struct {
	// (Updatable) The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringInput
	// If provided, specifies the ID of the boot volume backup policy to assign to the newly created boot volume. If omitted, no policy will be assigned.
	//
	// Deprecated: The 'backup_policy_id' field has been deprecated. Please use the 'oci_core_volume_backup_policy_assignment' resource instead.
	BackupPolicyId pulumi.StringPtrInput
	// (Updatable) The list of boot volume replicas to be enabled for this boot volume in the specified destination availability domains.
	BootVolumeReplicas BootVolumeBootVolumeReplicaArrayInput
	// (updatable) The boolean value, if you have replicas and want to disable replicas set this argument to true and remove `bootVolumeReplicas` in representation at the same time. If you want to enable a new replicas, remove this argument and use `bootVolumeReplicas` again.
	BootVolumeReplicasDeletion pulumi.BoolPtrInput
	// (Updatable) The OCID of the compartment that contains the boot volume.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) Specifies whether the auto-tune performance is enabled for this boot volume.
	IsAutoTuneEnabled pulumi.BoolPtrInput
	// (Updatable) The OCID of the Key Management key to assign as the master encryption key for the boot volume.
	KmsKeyId pulumi.StringPtrInput
	// (Updatable) The size of the volume in GBs.
	SizeInGbs     pulumi.StringPtrInput
	SourceDetails BootVolumeSourceDetailsInput
	// (Updatable) The number of volume performance units (VPUs) that will be applied to this volume per GB, representing the Block Volume service's elastic performance options. See [Block Volume Elastic Performance](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/blockvolumeelasticperformance.htm) for more information.
	VpusPerGb pulumi.StringPtrInput
}

func (BootVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bootVolumeArgs)(nil)).Elem()
}

type BootVolumeInput interface {
	pulumi.Input

	ToBootVolumeOutput() BootVolumeOutput
	ToBootVolumeOutputWithContext(ctx context.Context) BootVolumeOutput
}

func (*BootVolume) ElementType() reflect.Type {
	return reflect.TypeOf((*BootVolume)(nil))
}

func (i *BootVolume) ToBootVolumeOutput() BootVolumeOutput {
	return i.ToBootVolumeOutputWithContext(context.Background())
}

func (i *BootVolume) ToBootVolumeOutputWithContext(ctx context.Context) BootVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootVolumeOutput)
}

func (i *BootVolume) ToBootVolumePtrOutput() BootVolumePtrOutput {
	return i.ToBootVolumePtrOutputWithContext(context.Background())
}

func (i *BootVolume) ToBootVolumePtrOutputWithContext(ctx context.Context) BootVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootVolumePtrOutput)
}

type BootVolumePtrInput interface {
	pulumi.Input

	ToBootVolumePtrOutput() BootVolumePtrOutput
	ToBootVolumePtrOutputWithContext(ctx context.Context) BootVolumePtrOutput
}

type bootVolumePtrType BootVolumeArgs

func (*bootVolumePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BootVolume)(nil))
}

func (i *bootVolumePtrType) ToBootVolumePtrOutput() BootVolumePtrOutput {
	return i.ToBootVolumePtrOutputWithContext(context.Background())
}

func (i *bootVolumePtrType) ToBootVolumePtrOutputWithContext(ctx context.Context) BootVolumePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootVolumePtrOutput)
}

// BootVolumeArrayInput is an input type that accepts BootVolumeArray and BootVolumeArrayOutput values.
// You can construct a concrete instance of `BootVolumeArrayInput` via:
//
//          BootVolumeArray{ BootVolumeArgs{...} }
type BootVolumeArrayInput interface {
	pulumi.Input

	ToBootVolumeArrayOutput() BootVolumeArrayOutput
	ToBootVolumeArrayOutputWithContext(context.Context) BootVolumeArrayOutput
}

type BootVolumeArray []BootVolumeInput

func (BootVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BootVolume)(nil)).Elem()
}

func (i BootVolumeArray) ToBootVolumeArrayOutput() BootVolumeArrayOutput {
	return i.ToBootVolumeArrayOutputWithContext(context.Background())
}

func (i BootVolumeArray) ToBootVolumeArrayOutputWithContext(ctx context.Context) BootVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootVolumeArrayOutput)
}

// BootVolumeMapInput is an input type that accepts BootVolumeMap and BootVolumeMapOutput values.
// You can construct a concrete instance of `BootVolumeMapInput` via:
//
//          BootVolumeMap{ "key": BootVolumeArgs{...} }
type BootVolumeMapInput interface {
	pulumi.Input

	ToBootVolumeMapOutput() BootVolumeMapOutput
	ToBootVolumeMapOutputWithContext(context.Context) BootVolumeMapOutput
}

type BootVolumeMap map[string]BootVolumeInput

func (BootVolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BootVolume)(nil)).Elem()
}

func (i BootVolumeMap) ToBootVolumeMapOutput() BootVolumeMapOutput {
	return i.ToBootVolumeMapOutputWithContext(context.Background())
}

func (i BootVolumeMap) ToBootVolumeMapOutputWithContext(ctx context.Context) BootVolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BootVolumeMapOutput)
}

type BootVolumeOutput struct {
	*pulumi.OutputState
}

func (BootVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BootVolume)(nil))
}

func (o BootVolumeOutput) ToBootVolumeOutput() BootVolumeOutput {
	return o
}

func (o BootVolumeOutput) ToBootVolumeOutputWithContext(ctx context.Context) BootVolumeOutput {
	return o
}

func (o BootVolumeOutput) ToBootVolumePtrOutput() BootVolumePtrOutput {
	return o.ToBootVolumePtrOutputWithContext(context.Background())
}

func (o BootVolumeOutput) ToBootVolumePtrOutputWithContext(ctx context.Context) BootVolumePtrOutput {
	return o.ApplyT(func(v BootVolume) *BootVolume {
		return &v
	}).(BootVolumePtrOutput)
}

type BootVolumePtrOutput struct {
	*pulumi.OutputState
}

func (BootVolumePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BootVolume)(nil))
}

func (o BootVolumePtrOutput) ToBootVolumePtrOutput() BootVolumePtrOutput {
	return o
}

func (o BootVolumePtrOutput) ToBootVolumePtrOutputWithContext(ctx context.Context) BootVolumePtrOutput {
	return o
}

type BootVolumeArrayOutput struct{ *pulumi.OutputState }

func (BootVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BootVolume)(nil))
}

func (o BootVolumeArrayOutput) ToBootVolumeArrayOutput() BootVolumeArrayOutput {
	return o
}

func (o BootVolumeArrayOutput) ToBootVolumeArrayOutputWithContext(ctx context.Context) BootVolumeArrayOutput {
	return o
}

func (o BootVolumeArrayOutput) Index(i pulumi.IntInput) BootVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BootVolume {
		return vs[0].([]BootVolume)[vs[1].(int)]
	}).(BootVolumeOutput)
}

type BootVolumeMapOutput struct{ *pulumi.OutputState }

func (BootVolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BootVolume)(nil))
}

func (o BootVolumeMapOutput) ToBootVolumeMapOutput() BootVolumeMapOutput {
	return o
}

func (o BootVolumeMapOutput) ToBootVolumeMapOutputWithContext(ctx context.Context) BootVolumeMapOutput {
	return o
}

func (o BootVolumeMapOutput) MapIndex(k pulumi.StringInput) BootVolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BootVolume {
		return vs[0].(map[string]BootVolume)[vs[1].(string)]
	}).(BootVolumeOutput)
}

func init() {
	pulumi.RegisterOutputType(BootVolumeOutput{})
	pulumi.RegisterOutputType(BootVolumePtrOutput{})
	pulumi.RegisterOutputType(BootVolumeArrayOutput{})
	pulumi.RegisterOutputType(BootVolumeMapOutput{})
}