// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Instance Console Connection resource in Oracle Cloud Infrastructure Core service.
//
// Creates a new console connection to the specified instance.
// After the console connection has been created and is available,
// you connect to the console using SSH.
//
// For more information about instance console connections, see [Troubleshooting Instances Using Instance Console Connections](https://docs.cloud.oracle.com/iaas/Content/Compute/References/serialconsole.htm).
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
// 		_, err := core.NewInstanceConsoleConnection(ctx, "testInstanceConsoleConnection", &core.InstanceConsoleConnectionArgs{
// 			InstanceId: pulumi.Any(oci_core_instance.Test_instance.Id),
// 			PublicKey:  pulumi.Any(_var.Instance_console_connection_public_key),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
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
// InstanceConsoleConnections can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/instanceConsoleConnection:InstanceConsoleConnection test_instance_console_connection "id"
// ```
type InstanceConsoleConnection struct {
	pulumi.CustomResourceState

	// The OCID of the compartment to contain the console connection.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// The SSH connection string for the console connection.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// The SSH public key fingerprint for the console connection.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// The OCID of the instance to create the console connection to.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The SSH public key used to authenticate the console connection.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The current state of the console connection.
	State pulumi.StringOutput `pulumi:"state"`
	// The SSH connection string for the SSH tunnel used to connect to the console connection over VNC.
	VncConnectionString pulumi.StringOutput `pulumi:"vncConnectionString"`
}

// NewInstanceConsoleConnection registers a new resource with the given unique name, arguments, and options.
func NewInstanceConsoleConnection(ctx *pulumi.Context,
	name string, args *InstanceConsoleConnectionArgs, opts ...pulumi.ResourceOption) (*InstanceConsoleConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	var resource InstanceConsoleConnection
	err := ctx.RegisterResource("oci:core/instanceConsoleConnection:InstanceConsoleConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceConsoleConnection gets an existing InstanceConsoleConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceConsoleConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceConsoleConnectionState, opts ...pulumi.ResourceOption) (*InstanceConsoleConnection, error) {
	var resource InstanceConsoleConnection
	err := ctx.ReadResource("oci:core/instanceConsoleConnection:InstanceConsoleConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceConsoleConnection resources.
type instanceConsoleConnectionState struct {
	// The OCID of the compartment to contain the console connection.
	CompartmentId *string `pulumi:"compartmentId"`
	// The SSH connection string for the console connection.
	ConnectionString *string `pulumi:"connectionString"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// The SSH public key fingerprint for the console connection.
	Fingerprint *string `pulumi:"fingerprint"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of the instance to create the console connection to.
	InstanceId *string `pulumi:"instanceId"`
	// The SSH public key used to authenticate the console connection.
	PublicKey *string `pulumi:"publicKey"`
	// The current state of the console connection.
	State *string `pulumi:"state"`
	// The SSH connection string for the SSH tunnel used to connect to the console connection over VNC.
	VncConnectionString *string `pulumi:"vncConnectionString"`
}

type InstanceConsoleConnectionState struct {
	// The OCID of the compartment to contain the console connection.
	CompartmentId pulumi.StringPtrInput
	// The SSH connection string for the console connection.
	ConnectionString pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// The SSH public key fingerprint for the console connection.
	Fingerprint pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The OCID of the instance to create the console connection to.
	InstanceId pulumi.StringPtrInput
	// The SSH public key used to authenticate the console connection.
	PublicKey pulumi.StringPtrInput
	// The current state of the console connection.
	State pulumi.StringPtrInput
	// The SSH connection string for the SSH tunnel used to connect to the console connection over VNC.
	VncConnectionString pulumi.StringPtrInput
}

func (InstanceConsoleConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceConsoleConnectionState)(nil)).Elem()
}

type instanceConsoleConnectionArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The OCID of the instance to create the console connection to.
	InstanceId string `pulumi:"instanceId"`
	// The SSH public key used to authenticate the console connection.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a InstanceConsoleConnection resource.
type InstanceConsoleConnectionArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The OCID of the instance to create the console connection to.
	InstanceId pulumi.StringInput
	// The SSH public key used to authenticate the console connection.
	PublicKey pulumi.StringInput
}

func (InstanceConsoleConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceConsoleConnectionArgs)(nil)).Elem()
}

type InstanceConsoleConnectionInput interface {
	pulumi.Input

	ToInstanceConsoleConnectionOutput() InstanceConsoleConnectionOutput
	ToInstanceConsoleConnectionOutputWithContext(ctx context.Context) InstanceConsoleConnectionOutput
}

func (*InstanceConsoleConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceConsoleConnection)(nil))
}

func (i *InstanceConsoleConnection) ToInstanceConsoleConnectionOutput() InstanceConsoleConnectionOutput {
	return i.ToInstanceConsoleConnectionOutputWithContext(context.Background())
}

func (i *InstanceConsoleConnection) ToInstanceConsoleConnectionOutputWithContext(ctx context.Context) InstanceConsoleConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConsoleConnectionOutput)
}

func (i *InstanceConsoleConnection) ToInstanceConsoleConnectionPtrOutput() InstanceConsoleConnectionPtrOutput {
	return i.ToInstanceConsoleConnectionPtrOutputWithContext(context.Background())
}

func (i *InstanceConsoleConnection) ToInstanceConsoleConnectionPtrOutputWithContext(ctx context.Context) InstanceConsoleConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConsoleConnectionPtrOutput)
}

type InstanceConsoleConnectionPtrInput interface {
	pulumi.Input

	ToInstanceConsoleConnectionPtrOutput() InstanceConsoleConnectionPtrOutput
	ToInstanceConsoleConnectionPtrOutputWithContext(ctx context.Context) InstanceConsoleConnectionPtrOutput
}

type instanceConsoleConnectionPtrType InstanceConsoleConnectionArgs

func (*instanceConsoleConnectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceConsoleConnection)(nil))
}

func (i *instanceConsoleConnectionPtrType) ToInstanceConsoleConnectionPtrOutput() InstanceConsoleConnectionPtrOutput {
	return i.ToInstanceConsoleConnectionPtrOutputWithContext(context.Background())
}

func (i *instanceConsoleConnectionPtrType) ToInstanceConsoleConnectionPtrOutputWithContext(ctx context.Context) InstanceConsoleConnectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConsoleConnectionPtrOutput)
}

// InstanceConsoleConnectionArrayInput is an input type that accepts InstanceConsoleConnectionArray and InstanceConsoleConnectionArrayOutput values.
// You can construct a concrete instance of `InstanceConsoleConnectionArrayInput` via:
//
//          InstanceConsoleConnectionArray{ InstanceConsoleConnectionArgs{...} }
type InstanceConsoleConnectionArrayInput interface {
	pulumi.Input

	ToInstanceConsoleConnectionArrayOutput() InstanceConsoleConnectionArrayOutput
	ToInstanceConsoleConnectionArrayOutputWithContext(context.Context) InstanceConsoleConnectionArrayOutput
}

type InstanceConsoleConnectionArray []InstanceConsoleConnectionInput

func (InstanceConsoleConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceConsoleConnection)(nil)).Elem()
}

func (i InstanceConsoleConnectionArray) ToInstanceConsoleConnectionArrayOutput() InstanceConsoleConnectionArrayOutput {
	return i.ToInstanceConsoleConnectionArrayOutputWithContext(context.Background())
}

func (i InstanceConsoleConnectionArray) ToInstanceConsoleConnectionArrayOutputWithContext(ctx context.Context) InstanceConsoleConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConsoleConnectionArrayOutput)
}

// InstanceConsoleConnectionMapInput is an input type that accepts InstanceConsoleConnectionMap and InstanceConsoleConnectionMapOutput values.
// You can construct a concrete instance of `InstanceConsoleConnectionMapInput` via:
//
//          InstanceConsoleConnectionMap{ "key": InstanceConsoleConnectionArgs{...} }
type InstanceConsoleConnectionMapInput interface {
	pulumi.Input

	ToInstanceConsoleConnectionMapOutput() InstanceConsoleConnectionMapOutput
	ToInstanceConsoleConnectionMapOutputWithContext(context.Context) InstanceConsoleConnectionMapOutput
}

type InstanceConsoleConnectionMap map[string]InstanceConsoleConnectionInput

func (InstanceConsoleConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceConsoleConnection)(nil)).Elem()
}

func (i InstanceConsoleConnectionMap) ToInstanceConsoleConnectionMapOutput() InstanceConsoleConnectionMapOutput {
	return i.ToInstanceConsoleConnectionMapOutputWithContext(context.Background())
}

func (i InstanceConsoleConnectionMap) ToInstanceConsoleConnectionMapOutputWithContext(ctx context.Context) InstanceConsoleConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceConsoleConnectionMapOutput)
}

type InstanceConsoleConnectionOutput struct {
	*pulumi.OutputState
}

func (InstanceConsoleConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceConsoleConnection)(nil))
}

func (o InstanceConsoleConnectionOutput) ToInstanceConsoleConnectionOutput() InstanceConsoleConnectionOutput {
	return o
}

func (o InstanceConsoleConnectionOutput) ToInstanceConsoleConnectionOutputWithContext(ctx context.Context) InstanceConsoleConnectionOutput {
	return o
}

func (o InstanceConsoleConnectionOutput) ToInstanceConsoleConnectionPtrOutput() InstanceConsoleConnectionPtrOutput {
	return o.ToInstanceConsoleConnectionPtrOutputWithContext(context.Background())
}

func (o InstanceConsoleConnectionOutput) ToInstanceConsoleConnectionPtrOutputWithContext(ctx context.Context) InstanceConsoleConnectionPtrOutput {
	return o.ApplyT(func(v InstanceConsoleConnection) *InstanceConsoleConnection {
		return &v
	}).(InstanceConsoleConnectionPtrOutput)
}

type InstanceConsoleConnectionPtrOutput struct {
	*pulumi.OutputState
}

func (InstanceConsoleConnectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceConsoleConnection)(nil))
}

func (o InstanceConsoleConnectionPtrOutput) ToInstanceConsoleConnectionPtrOutput() InstanceConsoleConnectionPtrOutput {
	return o
}

func (o InstanceConsoleConnectionPtrOutput) ToInstanceConsoleConnectionPtrOutputWithContext(ctx context.Context) InstanceConsoleConnectionPtrOutput {
	return o
}

type InstanceConsoleConnectionArrayOutput struct{ *pulumi.OutputState }

func (InstanceConsoleConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceConsoleConnection)(nil))
}

func (o InstanceConsoleConnectionArrayOutput) ToInstanceConsoleConnectionArrayOutput() InstanceConsoleConnectionArrayOutput {
	return o
}

func (o InstanceConsoleConnectionArrayOutput) ToInstanceConsoleConnectionArrayOutputWithContext(ctx context.Context) InstanceConsoleConnectionArrayOutput {
	return o
}

func (o InstanceConsoleConnectionArrayOutput) Index(i pulumi.IntInput) InstanceConsoleConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceConsoleConnection {
		return vs[0].([]InstanceConsoleConnection)[vs[1].(int)]
	}).(InstanceConsoleConnectionOutput)
}

type InstanceConsoleConnectionMapOutput struct{ *pulumi.OutputState }

func (InstanceConsoleConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]InstanceConsoleConnection)(nil))
}

func (o InstanceConsoleConnectionMapOutput) ToInstanceConsoleConnectionMapOutput() InstanceConsoleConnectionMapOutput {
	return o
}

func (o InstanceConsoleConnectionMapOutput) ToInstanceConsoleConnectionMapOutputWithContext(ctx context.Context) InstanceConsoleConnectionMapOutput {
	return o
}

func (o InstanceConsoleConnectionMapOutput) MapIndex(k pulumi.StringInput) InstanceConsoleConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) InstanceConsoleConnection {
		return vs[0].(map[string]InstanceConsoleConnection)[vs[1].(string)]
	}).(InstanceConsoleConnectionOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceConsoleConnectionOutput{})
	pulumi.RegisterOutputType(InstanceConsoleConnectionPtrOutput{})
	pulumi.RegisterOutputType(InstanceConsoleConnectionArrayOutput{})
	pulumi.RegisterOutputType(InstanceConsoleConnectionMapOutput{})
}
