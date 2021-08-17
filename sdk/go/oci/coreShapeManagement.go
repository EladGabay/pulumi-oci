// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Shape Management resource in Oracle Cloud Infrastructure Core service.
//
// Add/Remove the specified shape from the compatible shapes list for the image.
type CoreShapeManagement struct {
	pulumi.CustomResourceState

	// The OCID of the compartment containing the image.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// The OCID of the Image to which the shape should be added.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The compatible shape that is to be added to the compatible shapes list for the image.
	ShapeName pulumi.StringOutput `pulumi:"shapeName"`
}

// NewCoreShapeManagement registers a new resource with the given unique name, arguments, and options.
func NewCoreShapeManagement(ctx *pulumi.Context,
	name string, args *CoreShapeManagementArgs, opts ...pulumi.ResourceOption) (*CoreShapeManagement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.ShapeName == nil {
		return nil, errors.New("invalid value for required argument 'ShapeName'")
	}
	var resource CoreShapeManagement
	err := ctx.RegisterResource("oci:index/coreShapeManagement:CoreShapeManagement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreShapeManagement gets an existing CoreShapeManagement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreShapeManagement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreShapeManagementState, opts ...pulumi.ResourceOption) (*CoreShapeManagement, error) {
	var resource CoreShapeManagement
	err := ctx.ReadResource("oci:index/coreShapeManagement:CoreShapeManagement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreShapeManagement resources.
type coreShapeManagementState struct {
	// The OCID of the compartment containing the image.
	CompartmentId *string `pulumi:"compartmentId"`
	// The OCID of the Image to which the shape should be added.
	ImageId *string `pulumi:"imageId"`
	// The compatible shape that is to be added to the compatible shapes list for the image.
	ShapeName *string `pulumi:"shapeName"`
}

type CoreShapeManagementState struct {
	// The OCID of the compartment containing the image.
	CompartmentId pulumi.StringPtrInput
	// The OCID of the Image to which the shape should be added.
	ImageId pulumi.StringPtrInput
	// The compatible shape that is to be added to the compatible shapes list for the image.
	ShapeName pulumi.StringPtrInput
}

func (CoreShapeManagementState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreShapeManagementState)(nil)).Elem()
}

type coreShapeManagementArgs struct {
	// The OCID of the compartment containing the image.
	CompartmentId string `pulumi:"compartmentId"`
	// The OCID of the Image to which the shape should be added.
	ImageId string `pulumi:"imageId"`
	// The compatible shape that is to be added to the compatible shapes list for the image.
	ShapeName string `pulumi:"shapeName"`
}

// The set of arguments for constructing a CoreShapeManagement resource.
type CoreShapeManagementArgs struct {
	// The OCID of the compartment containing the image.
	CompartmentId pulumi.StringInput
	// The OCID of the Image to which the shape should be added.
	ImageId pulumi.StringInput
	// The compatible shape that is to be added to the compatible shapes list for the image.
	ShapeName pulumi.StringInput
}

func (CoreShapeManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreShapeManagementArgs)(nil)).Elem()
}

type CoreShapeManagementInput interface {
	pulumi.Input

	ToCoreShapeManagementOutput() CoreShapeManagementOutput
	ToCoreShapeManagementOutputWithContext(ctx context.Context) CoreShapeManagementOutput
}

func (*CoreShapeManagement) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreShapeManagement)(nil))
}

func (i *CoreShapeManagement) ToCoreShapeManagementOutput() CoreShapeManagementOutput {
	return i.ToCoreShapeManagementOutputWithContext(context.Background())
}

func (i *CoreShapeManagement) ToCoreShapeManagementOutputWithContext(ctx context.Context) CoreShapeManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreShapeManagementOutput)
}

func (i *CoreShapeManagement) ToCoreShapeManagementPtrOutput() CoreShapeManagementPtrOutput {
	return i.ToCoreShapeManagementPtrOutputWithContext(context.Background())
}

func (i *CoreShapeManagement) ToCoreShapeManagementPtrOutputWithContext(ctx context.Context) CoreShapeManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreShapeManagementPtrOutput)
}

type CoreShapeManagementPtrInput interface {
	pulumi.Input

	ToCoreShapeManagementPtrOutput() CoreShapeManagementPtrOutput
	ToCoreShapeManagementPtrOutputWithContext(ctx context.Context) CoreShapeManagementPtrOutput
}

type coreShapeManagementPtrType CoreShapeManagementArgs

func (*coreShapeManagementPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreShapeManagement)(nil))
}

func (i *coreShapeManagementPtrType) ToCoreShapeManagementPtrOutput() CoreShapeManagementPtrOutput {
	return i.ToCoreShapeManagementPtrOutputWithContext(context.Background())
}

func (i *coreShapeManagementPtrType) ToCoreShapeManagementPtrOutputWithContext(ctx context.Context) CoreShapeManagementPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreShapeManagementPtrOutput)
}

// CoreShapeManagementArrayInput is an input type that accepts CoreShapeManagementArray and CoreShapeManagementArrayOutput values.
// You can construct a concrete instance of `CoreShapeManagementArrayInput` via:
//
//          CoreShapeManagementArray{ CoreShapeManagementArgs{...} }
type CoreShapeManagementArrayInput interface {
	pulumi.Input

	ToCoreShapeManagementArrayOutput() CoreShapeManagementArrayOutput
	ToCoreShapeManagementArrayOutputWithContext(context.Context) CoreShapeManagementArrayOutput
}

type CoreShapeManagementArray []CoreShapeManagementInput

func (CoreShapeManagementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreShapeManagement)(nil)).Elem()
}

func (i CoreShapeManagementArray) ToCoreShapeManagementArrayOutput() CoreShapeManagementArrayOutput {
	return i.ToCoreShapeManagementArrayOutputWithContext(context.Background())
}

func (i CoreShapeManagementArray) ToCoreShapeManagementArrayOutputWithContext(ctx context.Context) CoreShapeManagementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreShapeManagementArrayOutput)
}

// CoreShapeManagementMapInput is an input type that accepts CoreShapeManagementMap and CoreShapeManagementMapOutput values.
// You can construct a concrete instance of `CoreShapeManagementMapInput` via:
//
//          CoreShapeManagementMap{ "key": CoreShapeManagementArgs{...} }
type CoreShapeManagementMapInput interface {
	pulumi.Input

	ToCoreShapeManagementMapOutput() CoreShapeManagementMapOutput
	ToCoreShapeManagementMapOutputWithContext(context.Context) CoreShapeManagementMapOutput
}

type CoreShapeManagementMap map[string]CoreShapeManagementInput

func (CoreShapeManagementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreShapeManagement)(nil)).Elem()
}

func (i CoreShapeManagementMap) ToCoreShapeManagementMapOutput() CoreShapeManagementMapOutput {
	return i.ToCoreShapeManagementMapOutputWithContext(context.Background())
}

func (i CoreShapeManagementMap) ToCoreShapeManagementMapOutputWithContext(ctx context.Context) CoreShapeManagementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreShapeManagementMapOutput)
}

type CoreShapeManagementOutput struct {
	*pulumi.OutputState
}

func (CoreShapeManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreShapeManagement)(nil))
}

func (o CoreShapeManagementOutput) ToCoreShapeManagementOutput() CoreShapeManagementOutput {
	return o
}

func (o CoreShapeManagementOutput) ToCoreShapeManagementOutputWithContext(ctx context.Context) CoreShapeManagementOutput {
	return o
}

func (o CoreShapeManagementOutput) ToCoreShapeManagementPtrOutput() CoreShapeManagementPtrOutput {
	return o.ToCoreShapeManagementPtrOutputWithContext(context.Background())
}

func (o CoreShapeManagementOutput) ToCoreShapeManagementPtrOutputWithContext(ctx context.Context) CoreShapeManagementPtrOutput {
	return o.ApplyT(func(v CoreShapeManagement) *CoreShapeManagement {
		return &v
	}).(CoreShapeManagementPtrOutput)
}

type CoreShapeManagementPtrOutput struct {
	*pulumi.OutputState
}

func (CoreShapeManagementPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreShapeManagement)(nil))
}

func (o CoreShapeManagementPtrOutput) ToCoreShapeManagementPtrOutput() CoreShapeManagementPtrOutput {
	return o
}

func (o CoreShapeManagementPtrOutput) ToCoreShapeManagementPtrOutputWithContext(ctx context.Context) CoreShapeManagementPtrOutput {
	return o
}

type CoreShapeManagementArrayOutput struct{ *pulumi.OutputState }

func (CoreShapeManagementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CoreShapeManagement)(nil))
}

func (o CoreShapeManagementArrayOutput) ToCoreShapeManagementArrayOutput() CoreShapeManagementArrayOutput {
	return o
}

func (o CoreShapeManagementArrayOutput) ToCoreShapeManagementArrayOutputWithContext(ctx context.Context) CoreShapeManagementArrayOutput {
	return o
}

func (o CoreShapeManagementArrayOutput) Index(i pulumi.IntInput) CoreShapeManagementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CoreShapeManagement {
		return vs[0].([]CoreShapeManagement)[vs[1].(int)]
	}).(CoreShapeManagementOutput)
}

type CoreShapeManagementMapOutput struct{ *pulumi.OutputState }

func (CoreShapeManagementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CoreShapeManagement)(nil))
}

func (o CoreShapeManagementMapOutput) ToCoreShapeManagementMapOutput() CoreShapeManagementMapOutput {
	return o
}

func (o CoreShapeManagementMapOutput) ToCoreShapeManagementMapOutputWithContext(ctx context.Context) CoreShapeManagementMapOutput {
	return o
}

func (o CoreShapeManagementMapOutput) MapIndex(k pulumi.StringInput) CoreShapeManagementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CoreShapeManagement {
		return vs[0].(map[string]CoreShapeManagement)[vs[1].(string)]
	}).(CoreShapeManagementOutput)
}

func init() {
	pulumi.RegisterOutputType(CoreShapeManagementOutput{})
	pulumi.RegisterOutputType(CoreShapeManagementPtrOutput{})
	pulumi.RegisterOutputType(CoreShapeManagementArrayOutput{})
	pulumi.RegisterOutputType(CoreShapeManagementMapOutput{})
}