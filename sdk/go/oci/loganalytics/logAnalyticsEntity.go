// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package loganalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Log Analytics Entity resource in Oracle Cloud Infrastructure Log Analytics service.
//
// Create a new log analytics entity.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/v4/go/oci/loganalytics"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := loganalytics.NewLogAnalyticsEntity(ctx, "testLogAnalyticsEntity", &loganalytics.LogAnalyticsEntityArgs{
// 			CompartmentId:   pulumi.Any(_var.Compartment_id),
// 			EntityTypeName:  pulumi.Any(_var.Log_analytics_entity_entity_type_name),
// 			Namespace:       pulumi.Any(_var.Log_analytics_entity_namespace),
// 			CloudResourceId: pulumi.Any(oci_log_analytics_cloud_resource.Test_cloud_resource.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"foo-namespace.bar-key": pulumi.Any("value"),
// 			},
// 			FreeformTags: pulumi.AnyMap{
// 				"bar-key": pulumi.Any("value"),
// 			},
// 			Hostname:          pulumi.Any(_var.Log_analytics_entity_hostname),
// 			ManagementAgentId: pulumi.Any(oci_management_agent_management_agent.Test_management_agent.Id),
// 			Properties:        pulumi.Any(_var.Log_analytics_entity_properties),
// 			SourceId:          pulumi.Any(oci_log_analytics_source.Test_source.Id),
// 			TimezoneRegion:    pulumi.Any(_var.Log_analytics_entity_timezone_region),
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
// LogAnalyticsEntities can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:loganalytics/logAnalyticsEntity:LogAnalyticsEntity test_log_analytics_entity "namespaces/{namespaceName}/logAnalyticsEntities/{logAnalyticsEntityId}"
// ```
type LogAnalyticsEntity struct {
	pulumi.CustomResourceState

	// The Boolean flag to indicate if logs are collected for an entity for log analytics usage.
	AreLogsCollected pulumi.BoolOutput `pulumi:"areLogsCollected"`
	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId pulumi.StringOutput `pulumi:"cloudResourceId"`
	// (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// Internal name for the log analytics entity type.
	EntityTypeInternalName pulumi.StringOutput `pulumi:"entityTypeInternalName"`
	// Log analytics entity type name.
	EntityTypeName pulumi.StringOutput `pulumi:"entityTypeName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// (Updatable) The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// lifecycleDetails has additional information regarding substeps such as management agent plugin deployment.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// Management agent (management-agents resource kind) compartment OCID
	ManagementAgentCompartmentId pulumi.StringOutput `pulumi:"managementAgentCompartmentId"`
	// Management agent (management-agents resource kind) display name
	ManagementAgentDisplayName pulumi.StringOutput `pulumi:"managementAgentDisplayName"`
	// (Updatable) The OCID of the Management Agent.
	ManagementAgentId pulumi.StringOutput `pulumi:"managementAgentId"`
	// (Updatable) Log analytics entity name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Logging Analytics namespace used for the request.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// (Updatable) The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties pulumi.MapOutput `pulumi:"properties"`
	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId pulumi.StringOutput `pulumi:"sourceId"`
	// The current state of the log analytics entity.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
	// (Updatable) The timezone region of the log analytics entity.
	TimezoneRegion pulumi.StringOutput `pulumi:"timezoneRegion"`
}

// NewLogAnalyticsEntity registers a new resource with the given unique name, arguments, and options.
func NewLogAnalyticsEntity(ctx *pulumi.Context,
	name string, args *LogAnalyticsEntityArgs, opts ...pulumi.ResourceOption) (*LogAnalyticsEntity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.EntityTypeName == nil {
		return nil, errors.New("invalid value for required argument 'EntityTypeName'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	var resource LogAnalyticsEntity
	err := ctx.RegisterResource("oci:loganalytics/logAnalyticsEntity:LogAnalyticsEntity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogAnalyticsEntity gets an existing LogAnalyticsEntity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogAnalyticsEntity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogAnalyticsEntityState, opts ...pulumi.ResourceOption) (*LogAnalyticsEntity, error) {
	var resource LogAnalyticsEntity
	err := ctx.ReadResource("oci:loganalytics/logAnalyticsEntity:LogAnalyticsEntity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogAnalyticsEntity resources.
type logAnalyticsEntityState struct {
	// The Boolean flag to indicate if logs are collected for an entity for log analytics usage.
	AreLogsCollected *bool `pulumi:"areLogsCollected"`
	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId *string `pulumi:"cloudResourceId"`
	// (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Internal name for the log analytics entity type.
	EntityTypeInternalName *string `pulumi:"entityTypeInternalName"`
	// Log analytics entity type name.
	EntityTypeName *string `pulumi:"entityTypeName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely.
	Hostname *string `pulumi:"hostname"`
	// lifecycleDetails has additional information regarding substeps such as management agent plugin deployment.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// Management agent (management-agents resource kind) compartment OCID
	ManagementAgentCompartmentId *string `pulumi:"managementAgentCompartmentId"`
	// Management agent (management-agents resource kind) display name
	ManagementAgentDisplayName *string `pulumi:"managementAgentDisplayName"`
	// (Updatable) The OCID of the Management Agent.
	ManagementAgentId *string `pulumi:"managementAgentId"`
	// (Updatable) Log analytics entity name.
	Name *string `pulumi:"name"`
	// The Logging Analytics namespace used for the request.
	Namespace *string `pulumi:"namespace"`
	// (Updatable) The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties map[string]interface{} `pulumi:"properties"`
	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId *string `pulumi:"sourceId"`
	// The current state of the log analytics entity.
	State *string `pulumi:"state"`
	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated *string `pulumi:"timeCreated"`
	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated *string `pulumi:"timeUpdated"`
	// (Updatable) The timezone region of the log analytics entity.
	TimezoneRegion *string `pulumi:"timezoneRegion"`
}

type LogAnalyticsEntityState struct {
	// The Boolean flag to indicate if logs are collected for an entity for log analytics usage.
	AreLogsCollected pulumi.BoolPtrInput
	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId pulumi.StringPtrInput
	// (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// Internal name for the log analytics entity type.
	EntityTypeInternalName pulumi.StringPtrInput
	// Log analytics entity type name.
	EntityTypeName pulumi.StringPtrInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely.
	Hostname pulumi.StringPtrInput
	// lifecycleDetails has additional information regarding substeps such as management agent plugin deployment.
	LifecycleDetails pulumi.StringPtrInput
	// Management agent (management-agents resource kind) compartment OCID
	ManagementAgentCompartmentId pulumi.StringPtrInput
	// Management agent (management-agents resource kind) display name
	ManagementAgentDisplayName pulumi.StringPtrInput
	// (Updatable) The OCID of the Management Agent.
	ManagementAgentId pulumi.StringPtrInput
	// (Updatable) Log analytics entity name.
	Name pulumi.StringPtrInput
	// The Logging Analytics namespace used for the request.
	Namespace pulumi.StringPtrInput
	// (Updatable) The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties pulumi.MapInput
	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId pulumi.StringPtrInput
	// The current state of the log analytics entity.
	State pulumi.StringPtrInput
	// The date and time the resource was created, in the format defined by RFC3339.
	TimeCreated pulumi.StringPtrInput
	// The date and time the resource was last updated, in the format defined by RFC3339.
	TimeUpdated pulumi.StringPtrInput
	// (Updatable) The timezone region of the log analytics entity.
	TimezoneRegion pulumi.StringPtrInput
}

func (LogAnalyticsEntityState) ElementType() reflect.Type {
	return reflect.TypeOf((*logAnalyticsEntityState)(nil)).Elem()
}

type logAnalyticsEntityArgs struct {
	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId *string `pulumi:"cloudResourceId"`
	// (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// Log analytics entity type name.
	EntityTypeName string `pulumi:"entityTypeName"`
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely.
	Hostname *string `pulumi:"hostname"`
	// (Updatable) The OCID of the Management Agent.
	ManagementAgentId *string `pulumi:"managementAgentId"`
	// (Updatable) Log analytics entity name.
	Name *string `pulumi:"name"`
	// The Logging Analytics namespace used for the request.
	Namespace string `pulumi:"namespace"`
	// (Updatable) The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties map[string]interface{} `pulumi:"properties"`
	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId *string `pulumi:"sourceId"`
	// (Updatable) The timezone region of the log analytics entity.
	TimezoneRegion *string `pulumi:"timezoneRegion"`
}

// The set of arguments for constructing a LogAnalyticsEntity resource.
type LogAnalyticsEntityArgs struct {
	// The OCID of the Cloud resource which this entity is a representation of. This may be blank when the entity represents a non-cloud resource that the customer may have on their premises.
	CloudResourceId pulumi.StringPtrInput
	// (Updatable) Compartment Identifier [OCID] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
	DefinedTags pulumi.MapInput
	// Log analytics entity type name.
	EntityTypeName pulumi.StringInput
	// (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The hostname where the entity represented here is actually present. This would be the output one would get if they run `echo $HOSTNAME` on Linux or an equivalent OS command. This may be different from management agents host since logs may be collected remotely.
	Hostname pulumi.StringPtrInput
	// (Updatable) The OCID of the Management Agent.
	ManagementAgentId pulumi.StringPtrInput
	// (Updatable) Log analytics entity name.
	Name pulumi.StringPtrInput
	// The Logging Analytics namespace used for the request.
	Namespace pulumi.StringInput
	// (Updatable) The name/value pairs for parameter values to be used in file patterns specified in log sources.
	Properties pulumi.MapInput
	// This indicates the type of source. It is primarily for Enterprise Manager Repository ID.
	SourceId pulumi.StringPtrInput
	// (Updatable) The timezone region of the log analytics entity.
	TimezoneRegion pulumi.StringPtrInput
}

func (LogAnalyticsEntityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logAnalyticsEntityArgs)(nil)).Elem()
}

type LogAnalyticsEntityInput interface {
	pulumi.Input

	ToLogAnalyticsEntityOutput() LogAnalyticsEntityOutput
	ToLogAnalyticsEntityOutputWithContext(ctx context.Context) LogAnalyticsEntityOutput
}

func (*LogAnalyticsEntity) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsEntity)(nil))
}

func (i *LogAnalyticsEntity) ToLogAnalyticsEntityOutput() LogAnalyticsEntityOutput {
	return i.ToLogAnalyticsEntityOutputWithContext(context.Background())
}

func (i *LogAnalyticsEntity) ToLogAnalyticsEntityOutputWithContext(ctx context.Context) LogAnalyticsEntityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsEntityOutput)
}

func (i *LogAnalyticsEntity) ToLogAnalyticsEntityPtrOutput() LogAnalyticsEntityPtrOutput {
	return i.ToLogAnalyticsEntityPtrOutputWithContext(context.Background())
}

func (i *LogAnalyticsEntity) ToLogAnalyticsEntityPtrOutputWithContext(ctx context.Context) LogAnalyticsEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsEntityPtrOutput)
}

type LogAnalyticsEntityPtrInput interface {
	pulumi.Input

	ToLogAnalyticsEntityPtrOutput() LogAnalyticsEntityPtrOutput
	ToLogAnalyticsEntityPtrOutputWithContext(ctx context.Context) LogAnalyticsEntityPtrOutput
}

type logAnalyticsEntityPtrType LogAnalyticsEntityArgs

func (*logAnalyticsEntityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsEntity)(nil))
}

func (i *logAnalyticsEntityPtrType) ToLogAnalyticsEntityPtrOutput() LogAnalyticsEntityPtrOutput {
	return i.ToLogAnalyticsEntityPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsEntityPtrType) ToLogAnalyticsEntityPtrOutputWithContext(ctx context.Context) LogAnalyticsEntityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsEntityPtrOutput)
}

// LogAnalyticsEntityArrayInput is an input type that accepts LogAnalyticsEntityArray and LogAnalyticsEntityArrayOutput values.
// You can construct a concrete instance of `LogAnalyticsEntityArrayInput` via:
//
//          LogAnalyticsEntityArray{ LogAnalyticsEntityArgs{...} }
type LogAnalyticsEntityArrayInput interface {
	pulumi.Input

	ToLogAnalyticsEntityArrayOutput() LogAnalyticsEntityArrayOutput
	ToLogAnalyticsEntityArrayOutputWithContext(context.Context) LogAnalyticsEntityArrayOutput
}

type LogAnalyticsEntityArray []LogAnalyticsEntityInput

func (LogAnalyticsEntityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogAnalyticsEntity)(nil)).Elem()
}

func (i LogAnalyticsEntityArray) ToLogAnalyticsEntityArrayOutput() LogAnalyticsEntityArrayOutput {
	return i.ToLogAnalyticsEntityArrayOutputWithContext(context.Background())
}

func (i LogAnalyticsEntityArray) ToLogAnalyticsEntityArrayOutputWithContext(ctx context.Context) LogAnalyticsEntityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsEntityArrayOutput)
}

// LogAnalyticsEntityMapInput is an input type that accepts LogAnalyticsEntityMap and LogAnalyticsEntityMapOutput values.
// You can construct a concrete instance of `LogAnalyticsEntityMapInput` via:
//
//          LogAnalyticsEntityMap{ "key": LogAnalyticsEntityArgs{...} }
type LogAnalyticsEntityMapInput interface {
	pulumi.Input

	ToLogAnalyticsEntityMapOutput() LogAnalyticsEntityMapOutput
	ToLogAnalyticsEntityMapOutputWithContext(context.Context) LogAnalyticsEntityMapOutput
}

type LogAnalyticsEntityMap map[string]LogAnalyticsEntityInput

func (LogAnalyticsEntityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogAnalyticsEntity)(nil)).Elem()
}

func (i LogAnalyticsEntityMap) ToLogAnalyticsEntityMapOutput() LogAnalyticsEntityMapOutput {
	return i.ToLogAnalyticsEntityMapOutputWithContext(context.Background())
}

func (i LogAnalyticsEntityMap) ToLogAnalyticsEntityMapOutputWithContext(ctx context.Context) LogAnalyticsEntityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsEntityMapOutput)
}

type LogAnalyticsEntityOutput struct {
	*pulumi.OutputState
}

func (LogAnalyticsEntityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsEntity)(nil))
}

func (o LogAnalyticsEntityOutput) ToLogAnalyticsEntityOutput() LogAnalyticsEntityOutput {
	return o
}

func (o LogAnalyticsEntityOutput) ToLogAnalyticsEntityOutputWithContext(ctx context.Context) LogAnalyticsEntityOutput {
	return o
}

func (o LogAnalyticsEntityOutput) ToLogAnalyticsEntityPtrOutput() LogAnalyticsEntityPtrOutput {
	return o.ToLogAnalyticsEntityPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsEntityOutput) ToLogAnalyticsEntityPtrOutputWithContext(ctx context.Context) LogAnalyticsEntityPtrOutput {
	return o.ApplyT(func(v LogAnalyticsEntity) *LogAnalyticsEntity {
		return &v
	}).(LogAnalyticsEntityPtrOutput)
}

type LogAnalyticsEntityPtrOutput struct {
	*pulumi.OutputState
}

func (LogAnalyticsEntityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsEntity)(nil))
}

func (o LogAnalyticsEntityPtrOutput) ToLogAnalyticsEntityPtrOutput() LogAnalyticsEntityPtrOutput {
	return o
}

func (o LogAnalyticsEntityPtrOutput) ToLogAnalyticsEntityPtrOutputWithContext(ctx context.Context) LogAnalyticsEntityPtrOutput {
	return o
}

type LogAnalyticsEntityArrayOutput struct{ *pulumi.OutputState }

func (LogAnalyticsEntityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogAnalyticsEntity)(nil))
}

func (o LogAnalyticsEntityArrayOutput) ToLogAnalyticsEntityArrayOutput() LogAnalyticsEntityArrayOutput {
	return o
}

func (o LogAnalyticsEntityArrayOutput) ToLogAnalyticsEntityArrayOutputWithContext(ctx context.Context) LogAnalyticsEntityArrayOutput {
	return o
}

func (o LogAnalyticsEntityArrayOutput) Index(i pulumi.IntInput) LogAnalyticsEntityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogAnalyticsEntity {
		return vs[0].([]LogAnalyticsEntity)[vs[1].(int)]
	}).(LogAnalyticsEntityOutput)
}

type LogAnalyticsEntityMapOutput struct{ *pulumi.OutputState }

func (LogAnalyticsEntityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LogAnalyticsEntity)(nil))
}

func (o LogAnalyticsEntityMapOutput) ToLogAnalyticsEntityMapOutput() LogAnalyticsEntityMapOutput {
	return o
}

func (o LogAnalyticsEntityMapOutput) ToLogAnalyticsEntityMapOutputWithContext(ctx context.Context) LogAnalyticsEntityMapOutput {
	return o
}

func (o LogAnalyticsEntityMapOutput) MapIndex(k pulumi.StringInput) LogAnalyticsEntityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LogAnalyticsEntity {
		return vs[0].(map[string]LogAnalyticsEntity)[vs[1].(string)]
	}).(LogAnalyticsEntityOutput)
}

func init() {
	pulumi.RegisterOutputType(LogAnalyticsEntityOutput{})
	pulumi.RegisterOutputType(LogAnalyticsEntityPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsEntityArrayOutput{})
	pulumi.RegisterOutputType(LogAnalyticsEntityMapOutput{})
}
