// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Instance Configuration resource in Oracle Cloud Infrastructure Core service.
//
// Creates an instance configuration. An instance configuration is a template that defines the
// settings to use when creating Compute instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/core"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := core.NewInstanceConfiguration(ctx, "testInstanceConfiguration", &core.InstanceConfigurationArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName: pulumi.Any(_var.Instance_configuration_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			InstanceDetails: &core.InstanceConfigurationInstanceDetailsArgs{
// 				InstanceType: pulumi.Any(_var.Instance_configuration_instance_details_instance_type),
// 				BlockVolumes: core.InstanceConfigurationInstanceDetailsBlockVolumeArray{
// 					&core.InstanceConfigurationInstanceDetailsBlockVolumeArgs{
// 						AttachDetails: &core.InstanceConfigurationInstanceDetailsBlockVolumeAttachDetailsArgs{
// 							Type:                           pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_type),
// 							Device:                         pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_device),
// 							DisplayName:                    pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_display_name),
// 							IsPvEncryptionInTransitEnabled: pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_is_pv_encryption_in_transit_enabled),
// 							IsReadOnly:                     pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_is_read_only),
// 							IsShareable:                    pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_is_shareable),
// 							UseChap:                        pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_attach_details_use_chap),
// 						},
// 						CreateDetails: &core.InstanceConfigurationInstanceDetailsBlockVolumeCreateDetailsArgs{
// 							AvailabilityDomain: pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_create_details_availability_domain),
// 							BackupPolicyId:     pulumi.Any(data.Oci_core_volume_backup_policies.Test_volume_backup_policies.Volume_backup_policies[0].Id),
// 							CompartmentId:      pulumi.Any(_var.Compartment_id),
// 							DefinedTags: pulumi.AnyMap{
// 								"Operations.CostCenter": pulumi.Any("42"),
// 							},
// 							DisplayName: pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_create_details_display_name),
// 							FreeformTags: pulumi.AnyMap{
// 								"Department": pulumi.Any("Finance"),
// 							},
// 							KmsKeyId:  pulumi.Any(oci_kms_key.Test_key.Id),
// 							SizeInGbs: pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_create_details_size_in_gbs),
// 							SourceDetails: &core.InstanceConfigurationInstanceDetailsBlockVolumeCreateDetailsSourceDetailsArgs{
// 								Type: pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_create_details_source_details_type),
// 								Id:   pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_create_details_source_details_id),
// 							},
// 							VpusPerGb: pulumi.Any(_var.Instance_configuration_instance_details_block_volumes_create_details_vpus_per_gb),
// 						},
// 						VolumeId: pulumi.Any(oci_core_volume.Test_volume.Id),
// 					},
// 				},
// 				LaunchDetails: &core.InstanceConfigurationInstanceDetailsLaunchDetailsArgs{
// 					AgentConfig: &core.InstanceConfigurationInstanceDetailsLaunchDetailsAgentConfigArgs{
// 						AreAllPluginsDisabled: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_agent_config_are_all_plugins_disabled),
// 						IsManagementDisabled:  pulumi.Any(_var.Instance_configuration_instance_details_launch_details_agent_config_is_management_disabled),
// 						IsMonitoringDisabled:  pulumi.Any(_var.Instance_configuration_instance_details_launch_details_agent_config_is_monitoring_disabled),
// 						PluginsConfigs: core.InstanceConfigurationInstanceDetailsLaunchDetailsAgentConfigPluginsConfigArray{
// 							&core.InstanceConfigurationInstanceDetailsLaunchDetailsAgentConfigPluginsConfigArgs{
// 								DesiredState: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_agent_config_plugins_config_desired_state),
// 								Name:         pulumi.Any(_var.Instance_configuration_instance_details_launch_details_agent_config_plugins_config_name),
// 							},
// 						},
// 					},
// 					AvailabilityConfig: &core.InstanceConfigurationInstanceDetailsLaunchDetailsAvailabilityConfigArgs{
// 						RecoveryAction: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_availability_config_recovery_action),
// 					},
// 					AvailabilityDomain:    pulumi.Any(_var.Instance_configuration_instance_details_launch_details_availability_domain),
// 					CapacityReservationId: pulumi.Any(oci_core_capacity_reservation.Test_capacity_reservation.Id),
// 					CompartmentId:         pulumi.Any(_var.Compartment_id),
// 					CreateVnicDetails: &core.InstanceConfigurationInstanceDetailsLaunchDetailsCreateVnicDetailsArgs{
// 						AssignPrivateDnsRecord: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_assign_private_dns_record),
// 						AssignPublicIp:         pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_assign_public_ip),
// 						DefinedTags: pulumi.AnyMap{
// 							"Operations.CostCenter": pulumi.Any("42"),
// 						},
// 						DisplayName: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_display_name),
// 						FreeformTags: pulumi.AnyMap{
// 							"Department": pulumi.Any("Finance"),
// 						},
// 						HostnameLabel:       pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_hostname_label),
// 						NsgIds:              pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_nsg_ids),
// 						PrivateIp:           pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_private_ip),
// 						SkipSourceDestCheck: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_create_vnic_details_skip_source_dest_check),
// 						SubnetId:            pulumi.Any(oci_core_subnet.Test_subnet.Id),
// 					},
// 					DedicatedVmHostId: pulumi.Any(oci_core_dedicated_vm_host.Test_dedicated_vm_host.Id),
// 					DefinedTags: pulumi.AnyMap{
// 						"Operations.CostCenter": pulumi.Any("42"),
// 					},
// 					DisplayName:      pulumi.Any(_var.Instance_configuration_instance_details_launch_details_display_name),
// 					ExtendedMetadata: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_extended_metadata),
// 					FaultDomain:      pulumi.Any(_var.Instance_configuration_instance_details_launch_details_fault_domain),
// 					FreeformTags: pulumi.AnyMap{
// 						"Department": pulumi.Any("Finance"),
// 					},
// 					InstanceOptions: &core.InstanceConfigurationInstanceDetailsLaunchDetailsInstanceOptionsArgs{
// 						AreLegacyImdsEndpointsDisabled: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_instance_options_are_legacy_imds_endpoints_disabled),
// 					},
// 					IpxeScript:                     pulumi.Any(_var.Instance_configuration_instance_details_launch_details_ipxe_script),
// 					IsPvEncryptionInTransitEnabled: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_is_pv_encryption_in_transit_enabled),
// 					LaunchMode:                     pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_mode),
// 					LaunchOptions: &core.InstanceConfigurationInstanceDetailsLaunchDetailsLaunchOptionsArgs{
// 						BootVolumeType:                  pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_options_boot_volume_type),
// 						Firmware:                        pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_options_firmware),
// 						IsConsistentVolumeNamingEnabled: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_options_is_consistent_volume_naming_enabled),
// 						IsPvEncryptionInTransitEnabled:  pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_options_is_pv_encryption_in_transit_enabled),
// 						NetworkType:                     pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_options_network_type),
// 						RemoteDataVolumeType:            pulumi.Any(_var.Instance_configuration_instance_details_launch_details_launch_options_remote_data_volume_type),
// 					},
// 					Metadata: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_metadata),
// 					PlatformConfig: &core.InstanceConfigurationInstanceDetailsLaunchDetailsPlatformConfigArgs{
// 						Type:               pulumi.Any(_var.Instance_configuration_instance_details_launch_details_platform_config_type),
// 						NumaNodesPerSocket: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_platform_config_numa_nodes_per_socket),
// 					},
// 					PreemptibleInstanceConfig: &core.InstanceConfigurationInstanceDetailsLaunchDetailsPreemptibleInstanceConfigArgs{
// 						PreemptionAction: &core.InstanceConfigurationInstanceDetailsLaunchDetailsPreemptibleInstanceConfigPreemptionActionArgs{
// 							Type:               pulumi.Any(_var.Instance_configuration_instance_details_launch_details_preemptible_instance_config_preemption_action_type),
// 							PreserveBootVolume: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_preemptible_instance_config_preemption_action_preserve_boot_volume),
// 						},
// 					},
// 					PreferredMaintenanceAction: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_preferred_maintenance_action),
// 					Shape:                      pulumi.Any(_var.Instance_configuration_instance_details_launch_details_shape),
// 					ShapeConfig: &core.InstanceConfigurationInstanceDetailsLaunchDetailsShapeConfigArgs{
// 						BaselineOcpuUtilization: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_shape_config_baseline_ocpu_utilization),
// 						MemoryInGbs:             pulumi.Any(_var.Instance_configuration_instance_details_launch_details_shape_config_memory_in_gbs),
// 						Ocpus:                   pulumi.Any(_var.Instance_configuration_instance_details_launch_details_shape_config_ocpus),
// 					},
// 					SourceDetails: &core.InstanceConfigurationInstanceDetailsLaunchDetailsSourceDetailsArgs{
// 						SourceType:          pulumi.Any(_var.Instance_configuration_instance_details_launch_details_source_details_source_type),
// 						BootVolumeId:        pulumi.Any(oci_core_boot_volume.Test_boot_volume.Id),
// 						BootVolumeSizeInGbs: pulumi.Any(_var.Instance_configuration_instance_details_launch_details_source_details_boot_volume_size_in_gbs),
// 						ImageId:             pulumi.Any(oci_core_image.Test_image.Id),
// 					},
// 				},
// 				SecondaryVnics: core.InstanceConfigurationInstanceDetailsSecondaryVnicArray{
// 					&core.InstanceConfigurationInstanceDetailsSecondaryVnicArgs{
// 						CreateVnicDetails: &core.InstanceConfigurationInstanceDetailsSecondaryVnicCreateVnicDetailsArgs{
// 							AssignPrivateDnsRecord: pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_private_dns_record),
// 							AssignPublicIp:         pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_assign_public_ip),
// 							DefinedTags: pulumi.AnyMap{
// 								"Operations.CostCenter": pulumi.Any("42"),
// 							},
// 							DisplayName: pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_display_name),
// 							FreeformTags: pulumi.AnyMap{
// 								"Department": pulumi.Any("Finance"),
// 							},
// 							HostnameLabel:       pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_hostname_label),
// 							NsgIds:              pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_nsg_ids),
// 							PrivateIp:           pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_private_ip),
// 							SkipSourceDestCheck: pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_create_vnic_details_skip_source_dest_check),
// 							SubnetId:            pulumi.Any(oci_core_subnet.Test_subnet.Id),
// 						},
// 						DisplayName: pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_display_name),
// 						NicIndex:    pulumi.Any(_var.Instance_configuration_instance_details_secondary_vnics_nic_index),
// 					},
// 				},
// 			},
// 			InstanceId: pulumi.Any(oci_core_instance.Test_instance.Id),
// 			Source:     pulumi.Any(_var.Instance_configuration_source),
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
// InstanceConfigurations can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/instanceConfiguration:InstanceConfiguration test_instance_configuration "id"
// ```
type InstanceConfiguration struct {
	pulumi.CustomResourceState

	// The OCID of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Parameters that were not specified when the instance configuration was created, but that are required to launch an instance from the instance configuration. See the [LaunchInstanceConfiguration](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstanceConfiguration) operation.
	DeferredFields pulumi.StringArrayOutput `pulumi:"deferredFields"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags    pulumi.MapOutput                           `pulumi:"freeformTags"`
	InstanceDetails InstanceConfigurationInstanceDetailsOutput `pulumi:"instanceDetails"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.
	Source pulumi.StringOutput `pulumi:"source"`
	// The date and time the instance configuration was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewInstanceConfiguration registers a new resource with the given unique name, arguments, and options.
func NewInstanceConfiguration(ctx *pulumi.Context,
	name string, args *InstanceConfigurationArgs, opts ...pulumi.ResourceOption) (*InstanceConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	var resource InstanceConfiguration
	err := ctx.RegisterResource("oci:core/instanceConfiguration:InstanceConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceConfiguration gets an existing InstanceConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceConfigurationState, opts ...pulumi.ResourceOption) (*InstanceConfiguration, error) {
	var resource InstanceConfiguration
	err := ctx.ReadResource("oci:core/instanceConfiguration:InstanceConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceConfiguration resources.
type instanceConfigurationState struct {
	// The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// Parameters that were not specified when the instance configuration was created, but that are required to launch an instance from the instance configuration. See the [LaunchInstanceConfiguration](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstanceConfiguration) operation.
	DeferredFields []string `pulumi:"deferredFields"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName *string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags    map[string]interface{}                `pulumi:"freeformTags"`
	InstanceDetails *InstanceConfigurationInstanceDetails `pulumi:"instanceDetails"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration.
	InstanceId *string `pulumi:"instanceId"`
	// The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.
	Source *string `pulumi:"source"`
	// The date and time the instance configuration was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
}

type InstanceConfigurationState struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// Parameters that were not specified when the instance configuration was created, but that are required to launch an instance from the instance configuration. See the [LaunchInstanceConfiguration](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstanceConfiguration) operation.
	DeferredFields pulumi.StringArrayInput
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName pulumi.StringPtrInput
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags    pulumi.MapInput
	InstanceDetails InstanceConfigurationInstanceDetailsPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration.
	InstanceId pulumi.StringPtrInput
	// The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.
	Source pulumi.StringPtrInput
	// The date and time the instance configuration was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
}

func (InstanceConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceConfigurationState)(nil)).Elem()
}

type instanceConfigurationArgs struct {
	// The OCID of the compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName *string `pulumi:"displayName"`
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags    map[string]interface{}                `pulumi:"freeformTags"`
	InstanceDetails *InstanceConfigurationInstanceDetails `pulumi:"instanceDetails"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration.
	InstanceId *string `pulumi:"instanceId"`
	// The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a InstanceConfiguration resource.
type InstanceConfigurationArgs struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringInput
	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed.
	DisplayName pulumi.StringPtrInput
	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags    pulumi.MapInput
	InstanceDetails InstanceConfigurationInstanceDetailsPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance to use to create the instance configuration.
	InstanceId pulumi.StringPtrInput
	// The source of the instance configuration. An instance configuration defines the settings to use when creating Compute instances, including details such as the base image, shape, and metadata. You can also specify the associated resources for the instance, such as block volume attachments and network configuration.
	Source pulumi.StringPtrInput
}

func (InstanceConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceConfigurationArgs)(nil)).Elem()
}

type InstanceConfigurationInput interface {
	pulumi.Input

	ToInstanceConfigurationOutput() InstanceConfigurationOutput
	ToInstanceConfigurationOutputWithContext(ctx context.Context) InstanceConfigurationOutput
}

func (*InstanceConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceConfiguration)(nil))
}

func (i *InstanceConfiguration) ToInstanceConfigurationOutput() InstanceConfigurationOutput {
	return i.ToInstanceConfigurationOutputWithContext(context.Background())
}

func (i *InstanceConfiguration) ToInstanceConfigurationOutputWithContext(ctx context.Context) InstanceConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigurationOutput)
}

func (i *InstanceConfiguration) ToInstanceConfigurationPtrOutput() InstanceConfigurationPtrOutput {
	return i.ToInstanceConfigurationPtrOutputWithContext(context.Background())
}

func (i *InstanceConfiguration) ToInstanceConfigurationPtrOutputWithContext(ctx context.Context) InstanceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigurationPtrOutput)
}

type InstanceConfigurationPtrInput interface {
	pulumi.Input

	ToInstanceConfigurationPtrOutput() InstanceConfigurationPtrOutput
	ToInstanceConfigurationPtrOutputWithContext(ctx context.Context) InstanceConfigurationPtrOutput
}

type instanceConfigurationPtrType InstanceConfigurationArgs

func (*instanceConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceConfiguration)(nil))
}

func (i *instanceConfigurationPtrType) ToInstanceConfigurationPtrOutput() InstanceConfigurationPtrOutput {
	return i.ToInstanceConfigurationPtrOutputWithContext(context.Background())
}

func (i *instanceConfigurationPtrType) ToInstanceConfigurationPtrOutputWithContext(ctx context.Context) InstanceConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigurationPtrOutput)
}

// InstanceConfigurationArrayInput is an input type that accepts InstanceConfigurationArray and InstanceConfigurationArrayOutput values.
// You can construct a concrete instance of `InstanceConfigurationArrayInput` via:
//
//          InstanceConfigurationArray{ InstanceConfigurationArgs{...} }
type InstanceConfigurationArrayInput interface {
	pulumi.Input

	ToInstanceConfigurationArrayOutput() InstanceConfigurationArrayOutput
	ToInstanceConfigurationArrayOutputWithContext(context.Context) InstanceConfigurationArrayOutput
}

type InstanceConfigurationArray []InstanceConfigurationInput

func (InstanceConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceConfiguration)(nil)).Elem()
}

func (i InstanceConfigurationArray) ToInstanceConfigurationArrayOutput() InstanceConfigurationArrayOutput {
	return i.ToInstanceConfigurationArrayOutputWithContext(context.Background())
}

func (i InstanceConfigurationArray) ToInstanceConfigurationArrayOutputWithContext(ctx context.Context) InstanceConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigurationArrayOutput)
}

// InstanceConfigurationMapInput is an input type that accepts InstanceConfigurationMap and InstanceConfigurationMapOutput values.
// You can construct a concrete instance of `InstanceConfigurationMapInput` via:
//
//          InstanceConfigurationMap{ "key": InstanceConfigurationArgs{...} }
type InstanceConfigurationMapInput interface {
	pulumi.Input

	ToInstanceConfigurationMapOutput() InstanceConfigurationMapOutput
	ToInstanceConfigurationMapOutputWithContext(context.Context) InstanceConfigurationMapOutput
}

type InstanceConfigurationMap map[string]InstanceConfigurationInput

func (InstanceConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceConfiguration)(nil)).Elem()
}

func (i InstanceConfigurationMap) ToInstanceConfigurationMapOutput() InstanceConfigurationMapOutput {
	return i.ToInstanceConfigurationMapOutputWithContext(context.Background())
}

func (i InstanceConfigurationMap) ToInstanceConfigurationMapOutputWithContext(ctx context.Context) InstanceConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConfigurationMapOutput)
}

type InstanceConfigurationOutput struct {
	*pulumi.OutputState
}

func (InstanceConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceConfiguration)(nil))
}

func (o InstanceConfigurationOutput) ToInstanceConfigurationOutput() InstanceConfigurationOutput {
	return o
}

func (o InstanceConfigurationOutput) ToInstanceConfigurationOutputWithContext(ctx context.Context) InstanceConfigurationOutput {
	return o
}

func (o InstanceConfigurationOutput) ToInstanceConfigurationPtrOutput() InstanceConfigurationPtrOutput {
	return o.ToInstanceConfigurationPtrOutputWithContext(context.Background())
}

func (o InstanceConfigurationOutput) ToInstanceConfigurationPtrOutputWithContext(ctx context.Context) InstanceConfigurationPtrOutput {
	return o.ApplyT(func(v InstanceConfiguration) *InstanceConfiguration {
		return &v
	}).(InstanceConfigurationPtrOutput)
}

type InstanceConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceConfiguration)(nil))
}

func (o InstanceConfigurationPtrOutput) ToInstanceConfigurationPtrOutput() InstanceConfigurationPtrOutput {
	return o
}

func (o InstanceConfigurationPtrOutput) ToInstanceConfigurationPtrOutputWithContext(ctx context.Context) InstanceConfigurationPtrOutput {
	return o
}

type InstanceConfigurationArrayOutput struct{ *pulumi.OutputState }

func (InstanceConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceConfiguration)(nil))
}

func (o InstanceConfigurationArrayOutput) ToInstanceConfigurationArrayOutput() InstanceConfigurationArrayOutput {
	return o
}

func (o InstanceConfigurationArrayOutput) ToInstanceConfigurationArrayOutputWithContext(ctx context.Context) InstanceConfigurationArrayOutput {
	return o
}

func (o InstanceConfigurationArrayOutput) Index(i pulumi.IntInput) InstanceConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceConfiguration {
		return vs[0].([]InstanceConfiguration)[vs[1].(int)]
	}).(InstanceConfigurationOutput)
}

type InstanceConfigurationMapOutput struct{ *pulumi.OutputState }

func (InstanceConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceConfiguration)(nil))
}

func (o InstanceConfigurationMapOutput) ToInstanceConfigurationMapOutput() InstanceConfigurationMapOutput {
	return o
}

func (o InstanceConfigurationMapOutput) ToInstanceConfigurationMapOutputWithContext(ctx context.Context) InstanceConfigurationMapOutput {
	return o
}

func (o InstanceConfigurationMapOutput) MapIndex(k pulumi.StringInput) InstanceConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceConfiguration {
		return vs[0].(map[string]InstanceConfiguration)[vs[1].(string)]
	}).(InstanceConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceConfigurationOutput{})
	pulumi.RegisterOutputType(InstanceConfigurationPtrOutput{})
	pulumi.RegisterOutputType(InstanceConfigurationArrayOutput{})
	pulumi.RegisterOutputType(InstanceConfigurationMapOutput{})
}
