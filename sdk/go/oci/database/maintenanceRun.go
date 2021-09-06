// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package database

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Maintenance Run resource in Oracle Cloud Infrastructure Database service.
//
// Updates the properties of a maintenance run, such as the state of a maintenance run.
//
// ## Import
//
// MaintenanceRuns can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:database/maintenanceRun:MaintenanceRun test_maintenance_run "id"
// ```
type MaintenanceRun struct {
	pulumi.CustomResourceState

	// The OCID of the compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Description of the maintenance run.
	Description pulumi.StringOutput `pulumi:"description"`
	// The user-friendly name for the maintenance run.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) If `FALSE`, skips the maintenance run.
	IsEnabled pulumi.BoolOutput `pulumi:"isEnabled"`
	// (Updatable) If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled pulumi.BoolOutput `pulumi:"isPatchNowEnabled"`
	// Additional information about the current lifecycle state.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// The maintenance run OCID.
	MaintenanceRunId pulumi.StringOutput `pulumi:"maintenanceRunId"`
	// Maintenance sub-type.
	MaintenanceSubtype pulumi.StringOutput `pulumi:"maintenanceSubtype"`
	// Maintenance type.
	MaintenanceType pulumi.StringOutput `pulumi:"maintenanceType"`
	// Contain the patch failure count.
	PatchFailureCount pulumi.IntOutput `pulumi:"patchFailureCount"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId pulumi.StringOutput `pulumi:"patchId"`
	// (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode pulumi.StringOutput `pulumi:"patchingMode"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
	PeerMaintenanceRunId pulumi.StringOutput `pulumi:"peerMaintenanceRunId"`
	// The current state of the maintenance run. For Autonomous Database on shared Exadata infrastructure, valid states are IN_PROGRESS, SUCCEEDED and FAILED.
	State pulumi.StringOutput `pulumi:"state"`
	// The ID of the target resource on which the maintenance run occurs.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
	// The type of the target resource on which the maintenance run occurs.
	TargetResourceType pulumi.StringOutput `pulumi:"targetResourceType"`
	// The date and time the maintenance run was completed.
	TimeEnded pulumi.StringOutput `pulumi:"timeEnded"`
	// (Updatable) The scheduled date and time of the maintenance run to update.
	TimeScheduled pulumi.StringOutput `pulumi:"timeScheduled"`
	// The date and time the maintenance run starts.
	TimeStarted pulumi.StringOutput `pulumi:"timeStarted"`
}

// NewMaintenanceRun registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceRun(ctx *pulumi.Context,
	name string, args *MaintenanceRunArgs, opts ...pulumi.ResourceOption) (*MaintenanceRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaintenanceRunId == nil {
		return nil, errors.New("invalid value for required argument 'MaintenanceRunId'")
	}
	var resource MaintenanceRun
	err := ctx.RegisterResource("oci:database/maintenanceRun:MaintenanceRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceRun gets an existing MaintenanceRun resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceRunState, opts ...pulumi.ResourceOption) (*MaintenanceRun, error) {
	var resource MaintenanceRun
	err := ctx.ReadResource("oci:database/maintenanceRun:MaintenanceRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceRun resources.
type maintenanceRunState struct {
	// The OCID of the compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// Description of the maintenance run.
	Description *string `pulumi:"description"`
	// The user-friendly name for the maintenance run.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) If `FALSE`, skips the maintenance run.
	IsEnabled *bool `pulumi:"isEnabled"`
	// (Updatable) If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled *bool `pulumi:"isPatchNowEnabled"`
	// Additional information about the current lifecycle state.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The maintenance run OCID.
	MaintenanceRunId *string `pulumi:"maintenanceRunId"`
	// Maintenance sub-type.
	MaintenanceSubtype *string `pulumi:"maintenanceSubtype"`
	// Maintenance type.
	MaintenanceType *string `pulumi:"maintenanceType"`
	// Contain the patch failure count.
	PatchFailureCount *int `pulumi:"patchFailureCount"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId *string `pulumi:"patchId"`
	// (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode *string `pulumi:"patchingMode"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
	PeerMaintenanceRunId *string `pulumi:"peerMaintenanceRunId"`
	// The current state of the maintenance run. For Autonomous Database on shared Exadata infrastructure, valid states are IN_PROGRESS, SUCCEEDED and FAILED.
	State *string `pulumi:"state"`
	// The ID of the target resource on which the maintenance run occurs.
	TargetResourceId *string `pulumi:"targetResourceId"`
	// The type of the target resource on which the maintenance run occurs.
	TargetResourceType *string `pulumi:"targetResourceType"`
	// The date and time the maintenance run was completed.
	TimeEnded *string `pulumi:"timeEnded"`
	// (Updatable) The scheduled date and time of the maintenance run to update.
	TimeScheduled *string `pulumi:"timeScheduled"`
	// The date and time the maintenance run starts.
	TimeStarted *string `pulumi:"timeStarted"`
}

type MaintenanceRunState struct {
	// The OCID of the compartment.
	CompartmentId pulumi.StringPtrInput
	// Description of the maintenance run.
	Description pulumi.StringPtrInput
	// The user-friendly name for the maintenance run.
	DisplayName pulumi.StringPtrInput
	// (Updatable) If `FALSE`, skips the maintenance run.
	IsEnabled pulumi.BoolPtrInput
	// (Updatable) If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled pulumi.BoolPtrInput
	// Additional information about the current lifecycle state.
	LifecycleDetails pulumi.StringPtrInput
	// The maintenance run OCID.
	MaintenanceRunId pulumi.StringPtrInput
	// Maintenance sub-type.
	MaintenanceSubtype pulumi.StringPtrInput
	// Maintenance type.
	MaintenanceType pulumi.StringPtrInput
	// Contain the patch failure count.
	PatchFailureCount pulumi.IntPtrInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId pulumi.StringPtrInput
	// (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance run for the Autonomous Data Guard association's peer container database.
	PeerMaintenanceRunId pulumi.StringPtrInput
	// The current state of the maintenance run. For Autonomous Database on shared Exadata infrastructure, valid states are IN_PROGRESS, SUCCEEDED and FAILED.
	State pulumi.StringPtrInput
	// The ID of the target resource on which the maintenance run occurs.
	TargetResourceId pulumi.StringPtrInput
	// The type of the target resource on which the maintenance run occurs.
	TargetResourceType pulumi.StringPtrInput
	// The date and time the maintenance run was completed.
	TimeEnded pulumi.StringPtrInput
	// (Updatable) The scheduled date and time of the maintenance run to update.
	TimeScheduled pulumi.StringPtrInput
	// The date and time the maintenance run starts.
	TimeStarted pulumi.StringPtrInput
}

func (MaintenanceRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceRunState)(nil)).Elem()
}

type maintenanceRunArgs struct {
	// (Updatable) If `FALSE`, skips the maintenance run.
	IsEnabled *bool `pulumi:"isEnabled"`
	// (Updatable) If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled *bool `pulumi:"isPatchNowEnabled"`
	// The maintenance run OCID.
	MaintenanceRunId string `pulumi:"maintenanceRunId"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId *string `pulumi:"patchId"`
	// (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode *string `pulumi:"patchingMode"`
	// (Updatable) The scheduled date and time of the maintenance run to update.
	TimeScheduled *string `pulumi:"timeScheduled"`
}

// The set of arguments for constructing a MaintenanceRun resource.
type MaintenanceRunArgs struct {
	// (Updatable) If `FALSE`, skips the maintenance run.
	IsEnabled pulumi.BoolPtrInput
	// (Updatable) If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled pulumi.BoolPtrInput
	// The maintenance run OCID.
	MaintenanceRunId pulumi.StringInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId pulumi.StringPtrInput
	// (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	PatchingMode pulumi.StringPtrInput
	// (Updatable) The scheduled date and time of the maintenance run to update.
	TimeScheduled pulumi.StringPtrInput
}

func (MaintenanceRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceRunArgs)(nil)).Elem()
}

type MaintenanceRunInput interface {
	pulumi.Input

	ToMaintenanceRunOutput() MaintenanceRunOutput
	ToMaintenanceRunOutputWithContext(ctx context.Context) MaintenanceRunOutput
}

func (*MaintenanceRun) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceRun)(nil))
}

func (i *MaintenanceRun) ToMaintenanceRunOutput() MaintenanceRunOutput {
	return i.ToMaintenanceRunOutputWithContext(context.Background())
}

func (i *MaintenanceRun) ToMaintenanceRunOutputWithContext(ctx context.Context) MaintenanceRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceRunOutput)
}

func (i *MaintenanceRun) ToMaintenanceRunPtrOutput() MaintenanceRunPtrOutput {
	return i.ToMaintenanceRunPtrOutputWithContext(context.Background())
}

func (i *MaintenanceRun) ToMaintenanceRunPtrOutputWithContext(ctx context.Context) MaintenanceRunPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceRunPtrOutput)
}

type MaintenanceRunPtrInput interface {
	pulumi.Input

	ToMaintenanceRunPtrOutput() MaintenanceRunPtrOutput
	ToMaintenanceRunPtrOutputWithContext(ctx context.Context) MaintenanceRunPtrOutput
}

type maintenanceRunPtrType MaintenanceRunArgs

func (*maintenanceRunPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceRun)(nil))
}

func (i *maintenanceRunPtrType) ToMaintenanceRunPtrOutput() MaintenanceRunPtrOutput {
	return i.ToMaintenanceRunPtrOutputWithContext(context.Background())
}

func (i *maintenanceRunPtrType) ToMaintenanceRunPtrOutputWithContext(ctx context.Context) MaintenanceRunPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceRunPtrOutput)
}

// MaintenanceRunArrayInput is an input type that accepts MaintenanceRunArray and MaintenanceRunArrayOutput values.
// You can construct a concrete instance of `MaintenanceRunArrayInput` via:
//
//          MaintenanceRunArray{ MaintenanceRunArgs{...} }
type MaintenanceRunArrayInput interface {
	pulumi.Input

	ToMaintenanceRunArrayOutput() MaintenanceRunArrayOutput
	ToMaintenanceRunArrayOutputWithContext(context.Context) MaintenanceRunArrayOutput
}

type MaintenanceRunArray []MaintenanceRunInput

func (MaintenanceRunArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceRun)(nil)).Elem()
}

func (i MaintenanceRunArray) ToMaintenanceRunArrayOutput() MaintenanceRunArrayOutput {
	return i.ToMaintenanceRunArrayOutputWithContext(context.Background())
}

func (i MaintenanceRunArray) ToMaintenanceRunArrayOutputWithContext(ctx context.Context) MaintenanceRunArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceRunArrayOutput)
}

// MaintenanceRunMapInput is an input type that accepts MaintenanceRunMap and MaintenanceRunMapOutput values.
// You can construct a concrete instance of `MaintenanceRunMapInput` via:
//
//          MaintenanceRunMap{ "key": MaintenanceRunArgs{...} }
type MaintenanceRunMapInput interface {
	pulumi.Input

	ToMaintenanceRunMapOutput() MaintenanceRunMapOutput
	ToMaintenanceRunMapOutputWithContext(context.Context) MaintenanceRunMapOutput
}

type MaintenanceRunMap map[string]MaintenanceRunInput

func (MaintenanceRunMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceRun)(nil)).Elem()
}

func (i MaintenanceRunMap) ToMaintenanceRunMapOutput() MaintenanceRunMapOutput {
	return i.ToMaintenanceRunMapOutputWithContext(context.Background())
}

func (i MaintenanceRunMap) ToMaintenanceRunMapOutputWithContext(ctx context.Context) MaintenanceRunMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceRunMapOutput)
}

type MaintenanceRunOutput struct {
	*pulumi.OutputState
}

func (MaintenanceRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceRun)(nil))
}

func (o MaintenanceRunOutput) ToMaintenanceRunOutput() MaintenanceRunOutput {
	return o
}

func (o MaintenanceRunOutput) ToMaintenanceRunOutputWithContext(ctx context.Context) MaintenanceRunOutput {
	return o
}

func (o MaintenanceRunOutput) ToMaintenanceRunPtrOutput() MaintenanceRunPtrOutput {
	return o.ToMaintenanceRunPtrOutputWithContext(context.Background())
}

func (o MaintenanceRunOutput) ToMaintenanceRunPtrOutputWithContext(ctx context.Context) MaintenanceRunPtrOutput {
	return o.ApplyT(func(v MaintenanceRun) *MaintenanceRun {
		return &v
	}).(MaintenanceRunPtrOutput)
}

type MaintenanceRunPtrOutput struct {
	*pulumi.OutputState
}

func (MaintenanceRunPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceRun)(nil))
}

func (o MaintenanceRunPtrOutput) ToMaintenanceRunPtrOutput() MaintenanceRunPtrOutput {
	return o
}

func (o MaintenanceRunPtrOutput) ToMaintenanceRunPtrOutputWithContext(ctx context.Context) MaintenanceRunPtrOutput {
	return o
}

type MaintenanceRunArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceRunArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MaintenanceRun)(nil))
}

func (o MaintenanceRunArrayOutput) ToMaintenanceRunArrayOutput() MaintenanceRunArrayOutput {
	return o
}

func (o MaintenanceRunArrayOutput) ToMaintenanceRunArrayOutputWithContext(ctx context.Context) MaintenanceRunArrayOutput {
	return o
}

func (o MaintenanceRunArrayOutput) Index(i pulumi.IntInput) MaintenanceRunOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MaintenanceRun {
		return vs[0].([]MaintenanceRun)[vs[1].(int)]
	}).(MaintenanceRunOutput)
}

type MaintenanceRunMapOutput struct{ *pulumi.OutputState }

func (MaintenanceRunMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MaintenanceRun)(nil))
}

func (o MaintenanceRunMapOutput) ToMaintenanceRunMapOutput() MaintenanceRunMapOutput {
	return o
}

func (o MaintenanceRunMapOutput) ToMaintenanceRunMapOutputWithContext(ctx context.Context) MaintenanceRunMapOutput {
	return o
}

func (o MaintenanceRunMapOutput) MapIndex(k pulumi.StringInput) MaintenanceRunOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MaintenanceRun {
		return vs[0].(map[string]MaintenanceRun)[vs[1].(string)]
	}).(MaintenanceRunOutput)
}

func init() {
	pulumi.RegisterOutputType(MaintenanceRunOutput{})
	pulumi.RegisterOutputType(MaintenanceRunPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceRunArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceRunMapOutput{})
}