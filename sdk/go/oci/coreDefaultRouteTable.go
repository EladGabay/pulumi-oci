// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CoreDefaultRouteTable struct {
	pulumi.CustomResourceState

	CompartmentId           pulumi.StringOutput                       `pulumi:"compartmentId"`
	DefinedTags             pulumi.MapOutput                          `pulumi:"definedTags"`
	DisplayName             pulumi.StringOutput                       `pulumi:"displayName"`
	FreeformTags            pulumi.MapOutput                          `pulumi:"freeformTags"`
	ManageDefaultResourceId pulumi.StringOutput                       `pulumi:"manageDefaultResourceId"`
	RouteRules              CoreDefaultRouteTableRouteRuleArrayOutput `pulumi:"routeRules"`
	State                   pulumi.StringOutput                       `pulumi:"state"`
	TimeCreated             pulumi.StringOutput                       `pulumi:"timeCreated"`
}

// NewCoreDefaultRouteTable registers a new resource with the given unique name, arguments, and options.
func NewCoreDefaultRouteTable(ctx *pulumi.Context,
	name string, args *CoreDefaultRouteTableArgs, opts ...pulumi.ResourceOption) (*CoreDefaultRouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManageDefaultResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ManageDefaultResourceId'")
	}
	var resource CoreDefaultRouteTable
	err := ctx.RegisterResource("oci:index/coreDefaultRouteTable:CoreDefaultRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreDefaultRouteTable gets an existing CoreDefaultRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreDefaultRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreDefaultRouteTableState, opts ...pulumi.ResourceOption) (*CoreDefaultRouteTable, error) {
	var resource CoreDefaultRouteTable
	err := ctx.ReadResource("oci:index/coreDefaultRouteTable:CoreDefaultRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreDefaultRouteTable resources.
type coreDefaultRouteTableState struct {
	CompartmentId           *string                          `pulumi:"compartmentId"`
	DefinedTags             map[string]interface{}           `pulumi:"definedTags"`
	DisplayName             *string                          `pulumi:"displayName"`
	FreeformTags            map[string]interface{}           `pulumi:"freeformTags"`
	ManageDefaultResourceId *string                          `pulumi:"manageDefaultResourceId"`
	RouteRules              []CoreDefaultRouteTableRouteRule `pulumi:"routeRules"`
	State                   *string                          `pulumi:"state"`
	TimeCreated             *string                          `pulumi:"timeCreated"`
}

type CoreDefaultRouteTableState struct {
	CompartmentId           pulumi.StringPtrInput
	DefinedTags             pulumi.MapInput
	DisplayName             pulumi.StringPtrInput
	FreeformTags            pulumi.MapInput
	ManageDefaultResourceId pulumi.StringPtrInput
	RouteRules              CoreDefaultRouteTableRouteRuleArrayInput
	State                   pulumi.StringPtrInput
	TimeCreated             pulumi.StringPtrInput
}

func (CoreDefaultRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDefaultRouteTableState)(nil)).Elem()
}

type coreDefaultRouteTableArgs struct {
	CompartmentId           *string                          `pulumi:"compartmentId"`
	DefinedTags             map[string]interface{}           `pulumi:"definedTags"`
	DisplayName             *string                          `pulumi:"displayName"`
	FreeformTags            map[string]interface{}           `pulumi:"freeformTags"`
	ManageDefaultResourceId string                           `pulumi:"manageDefaultResourceId"`
	RouteRules              []CoreDefaultRouteTableRouteRule `pulumi:"routeRules"`
}

// The set of arguments for constructing a CoreDefaultRouteTable resource.
type CoreDefaultRouteTableArgs struct {
	CompartmentId           pulumi.StringPtrInput
	DefinedTags             pulumi.MapInput
	DisplayName             pulumi.StringPtrInput
	FreeformTags            pulumi.MapInput
	ManageDefaultResourceId pulumi.StringInput
	RouteRules              CoreDefaultRouteTableRouteRuleArrayInput
}

func (CoreDefaultRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreDefaultRouteTableArgs)(nil)).Elem()
}

type CoreDefaultRouteTableInput interface {
	pulumi.Input

	ToCoreDefaultRouteTableOutput() CoreDefaultRouteTableOutput
	ToCoreDefaultRouteTableOutputWithContext(ctx context.Context) CoreDefaultRouteTableOutput
}

func (*CoreDefaultRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreDefaultRouteTable)(nil))
}

func (i *CoreDefaultRouteTable) ToCoreDefaultRouteTableOutput() CoreDefaultRouteTableOutput {
	return i.ToCoreDefaultRouteTableOutputWithContext(context.Background())
}

func (i *CoreDefaultRouteTable) ToCoreDefaultRouteTableOutputWithContext(ctx context.Context) CoreDefaultRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefaultRouteTableOutput)
}

func (i *CoreDefaultRouteTable) ToCoreDefaultRouteTablePtrOutput() CoreDefaultRouteTablePtrOutput {
	return i.ToCoreDefaultRouteTablePtrOutputWithContext(context.Background())
}

func (i *CoreDefaultRouteTable) ToCoreDefaultRouteTablePtrOutputWithContext(ctx context.Context) CoreDefaultRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefaultRouteTablePtrOutput)
}

type CoreDefaultRouteTablePtrInput interface {
	pulumi.Input

	ToCoreDefaultRouteTablePtrOutput() CoreDefaultRouteTablePtrOutput
	ToCoreDefaultRouteTablePtrOutputWithContext(ctx context.Context) CoreDefaultRouteTablePtrOutput
}

type coreDefaultRouteTablePtrType CoreDefaultRouteTableArgs

func (*coreDefaultRouteTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDefaultRouteTable)(nil))
}

func (i *coreDefaultRouteTablePtrType) ToCoreDefaultRouteTablePtrOutput() CoreDefaultRouteTablePtrOutput {
	return i.ToCoreDefaultRouteTablePtrOutputWithContext(context.Background())
}

func (i *coreDefaultRouteTablePtrType) ToCoreDefaultRouteTablePtrOutputWithContext(ctx context.Context) CoreDefaultRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefaultRouteTablePtrOutput)
}

// CoreDefaultRouteTableArrayInput is an input type that accepts CoreDefaultRouteTableArray and CoreDefaultRouteTableArrayOutput values.
// You can construct a concrete instance of `CoreDefaultRouteTableArrayInput` via:
//
//          CoreDefaultRouteTableArray{ CoreDefaultRouteTableArgs{...} }
type CoreDefaultRouteTableArrayInput interface {
	pulumi.Input

	ToCoreDefaultRouteTableArrayOutput() CoreDefaultRouteTableArrayOutput
	ToCoreDefaultRouteTableArrayOutputWithContext(context.Context) CoreDefaultRouteTableArrayOutput
}

type CoreDefaultRouteTableArray []CoreDefaultRouteTableInput

func (CoreDefaultRouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreDefaultRouteTable)(nil)).Elem()
}

func (i CoreDefaultRouteTableArray) ToCoreDefaultRouteTableArrayOutput() CoreDefaultRouteTableArrayOutput {
	return i.ToCoreDefaultRouteTableArrayOutputWithContext(context.Background())
}

func (i CoreDefaultRouteTableArray) ToCoreDefaultRouteTableArrayOutputWithContext(ctx context.Context) CoreDefaultRouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefaultRouteTableArrayOutput)
}

// CoreDefaultRouteTableMapInput is an input type that accepts CoreDefaultRouteTableMap and CoreDefaultRouteTableMapOutput values.
// You can construct a concrete instance of `CoreDefaultRouteTableMapInput` via:
//
//          CoreDefaultRouteTableMap{ "key": CoreDefaultRouteTableArgs{...} }
type CoreDefaultRouteTableMapInput interface {
	pulumi.Input

	ToCoreDefaultRouteTableMapOutput() CoreDefaultRouteTableMapOutput
	ToCoreDefaultRouteTableMapOutputWithContext(context.Context) CoreDefaultRouteTableMapOutput
}

type CoreDefaultRouteTableMap map[string]CoreDefaultRouteTableInput

func (CoreDefaultRouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreDefaultRouteTable)(nil)).Elem()
}

func (i CoreDefaultRouteTableMap) ToCoreDefaultRouteTableMapOutput() CoreDefaultRouteTableMapOutput {
	return i.ToCoreDefaultRouteTableMapOutputWithContext(context.Background())
}

func (i CoreDefaultRouteTableMap) ToCoreDefaultRouteTableMapOutputWithContext(ctx context.Context) CoreDefaultRouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreDefaultRouteTableMapOutput)
}

type CoreDefaultRouteTableOutput struct {
	*pulumi.OutputState
}

func (CoreDefaultRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreDefaultRouteTable)(nil))
}

func (o CoreDefaultRouteTableOutput) ToCoreDefaultRouteTableOutput() CoreDefaultRouteTableOutput {
	return o
}

func (o CoreDefaultRouteTableOutput) ToCoreDefaultRouteTableOutputWithContext(ctx context.Context) CoreDefaultRouteTableOutput {
	return o
}

func (o CoreDefaultRouteTableOutput) ToCoreDefaultRouteTablePtrOutput() CoreDefaultRouteTablePtrOutput {
	return o.ToCoreDefaultRouteTablePtrOutputWithContext(context.Background())
}

func (o CoreDefaultRouteTableOutput) ToCoreDefaultRouteTablePtrOutputWithContext(ctx context.Context) CoreDefaultRouteTablePtrOutput {
	return o.ApplyT(func(v CoreDefaultRouteTable) *CoreDefaultRouteTable {
		return &v
	}).(CoreDefaultRouteTablePtrOutput)
}

type CoreDefaultRouteTablePtrOutput struct {
	*pulumi.OutputState
}

func (CoreDefaultRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreDefaultRouteTable)(nil))
}

func (o CoreDefaultRouteTablePtrOutput) ToCoreDefaultRouteTablePtrOutput() CoreDefaultRouteTablePtrOutput {
	return o
}

func (o CoreDefaultRouteTablePtrOutput) ToCoreDefaultRouteTablePtrOutputWithContext(ctx context.Context) CoreDefaultRouteTablePtrOutput {
	return o
}

type CoreDefaultRouteTableArrayOutput struct{ *pulumi.OutputState }

func (CoreDefaultRouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CoreDefaultRouteTable)(nil))
}

func (o CoreDefaultRouteTableArrayOutput) ToCoreDefaultRouteTableArrayOutput() CoreDefaultRouteTableArrayOutput {
	return o
}

func (o CoreDefaultRouteTableArrayOutput) ToCoreDefaultRouteTableArrayOutputWithContext(ctx context.Context) CoreDefaultRouteTableArrayOutput {
	return o
}

func (o CoreDefaultRouteTableArrayOutput) Index(i pulumi.IntInput) CoreDefaultRouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CoreDefaultRouteTable {
		return vs[0].([]CoreDefaultRouteTable)[vs[1].(int)]
	}).(CoreDefaultRouteTableOutput)
}

type CoreDefaultRouteTableMapOutput struct{ *pulumi.OutputState }

func (CoreDefaultRouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CoreDefaultRouteTable)(nil))
}

func (o CoreDefaultRouteTableMapOutput) ToCoreDefaultRouteTableMapOutput() CoreDefaultRouteTableMapOutput {
	return o
}

func (o CoreDefaultRouteTableMapOutput) ToCoreDefaultRouteTableMapOutputWithContext(ctx context.Context) CoreDefaultRouteTableMapOutput {
	return o
}

func (o CoreDefaultRouteTableMapOutput) MapIndex(k pulumi.StringInput) CoreDefaultRouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CoreDefaultRouteTable {
		return vs[0].(map[string]CoreDefaultRouteTable)[vs[1].(string)]
	}).(CoreDefaultRouteTableOutput)
}

func init() {
	pulumi.RegisterOutputType(CoreDefaultRouteTableOutput{})
	pulumi.RegisterOutputType(CoreDefaultRouteTablePtrOutput{})
	pulumi.RegisterOutputType(CoreDefaultRouteTableArrayOutput{})
	pulumi.RegisterOutputType(CoreDefaultRouteTableMapOutput{})
}