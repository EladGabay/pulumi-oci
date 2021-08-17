// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Container Configuration resource in Oracle Cloud Infrastructure Artifacts service.
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
// 		_, err := oci.NewArtifactsContainerConfiguration(ctx, "testContainerConfiguration", nil)
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
// ContainerConfiguration can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:index/artifactsContainerConfiguration:ArtifactsContainerConfiguration test_container_configuration "container/configuration/compartmentId/{compartmentId}"
// ```
type ArtifactsContainerConfiguration struct {
	pulumi.CustomResourceState

	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment.
	IsRepositoryCreatedOnFirstPush pulumi.BoolOutput `pulumi:"isRepositoryCreatedOnFirstPush"`
	// The tenancy namespace used in the container repository path.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
}

// NewArtifactsContainerConfiguration registers a new resource with the given unique name, arguments, and options.
func NewArtifactsContainerConfiguration(ctx *pulumi.Context,
	name string, args *ArtifactsContainerConfigurationArgs, opts ...pulumi.ResourceOption) (*ArtifactsContainerConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.IsRepositoryCreatedOnFirstPush == nil {
		return nil, errors.New("invalid value for required argument 'IsRepositoryCreatedOnFirstPush'")
	}
	var resource ArtifactsContainerConfiguration
	err := ctx.RegisterResource("oci:index/artifactsContainerConfiguration:ArtifactsContainerConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetArtifactsContainerConfiguration gets an existing ArtifactsContainerConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetArtifactsContainerConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactsContainerConfigurationState, opts ...pulumi.ResourceOption) (*ArtifactsContainerConfiguration, error) {
	var resource ArtifactsContainerConfiguration
	err := ctx.ReadResource("oci:index/artifactsContainerConfiguration:ArtifactsContainerConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ArtifactsContainerConfiguration resources.
type artifactsContainerConfigurationState struct {
	CompartmentId *string `pulumi:"compartmentId"`
	// Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment.
	IsRepositoryCreatedOnFirstPush *bool `pulumi:"isRepositoryCreatedOnFirstPush"`
	// The tenancy namespace used in the container repository path.
	Namespace *string `pulumi:"namespace"`
}

type ArtifactsContainerConfigurationState struct {
	CompartmentId pulumi.StringPtrInput
	// Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment.
	IsRepositoryCreatedOnFirstPush pulumi.BoolPtrInput
	// The tenancy namespace used in the container repository path.
	Namespace pulumi.StringPtrInput
}

func (ArtifactsContainerConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactsContainerConfigurationState)(nil)).Elem()
}

type artifactsContainerConfigurationArgs struct {
	CompartmentId string `pulumi:"compartmentId"`
	// Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment.
	IsRepositoryCreatedOnFirstPush bool `pulumi:"isRepositoryCreatedOnFirstPush"`
}

// The set of arguments for constructing a ArtifactsContainerConfiguration resource.
type ArtifactsContainerConfigurationArgs struct {
	CompartmentId pulumi.StringInput
	// Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment.
	IsRepositoryCreatedOnFirstPush pulumi.BoolInput
}

func (ArtifactsContainerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactsContainerConfigurationArgs)(nil)).Elem()
}

type ArtifactsContainerConfigurationInput interface {
	pulumi.Input

	ToArtifactsContainerConfigurationOutput() ArtifactsContainerConfigurationOutput
	ToArtifactsContainerConfigurationOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationOutput
}

func (*ArtifactsContainerConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactsContainerConfiguration)(nil))
}

func (i *ArtifactsContainerConfiguration) ToArtifactsContainerConfigurationOutput() ArtifactsContainerConfigurationOutput {
	return i.ToArtifactsContainerConfigurationOutputWithContext(context.Background())
}

func (i *ArtifactsContainerConfiguration) ToArtifactsContainerConfigurationOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactsContainerConfigurationOutput)
}

func (i *ArtifactsContainerConfiguration) ToArtifactsContainerConfigurationPtrOutput() ArtifactsContainerConfigurationPtrOutput {
	return i.ToArtifactsContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i *ArtifactsContainerConfiguration) ToArtifactsContainerConfigurationPtrOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactsContainerConfigurationPtrOutput)
}

type ArtifactsContainerConfigurationPtrInput interface {
	pulumi.Input

	ToArtifactsContainerConfigurationPtrOutput() ArtifactsContainerConfigurationPtrOutput
	ToArtifactsContainerConfigurationPtrOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationPtrOutput
}

type artifactsContainerConfigurationPtrType ArtifactsContainerConfigurationArgs

func (*artifactsContainerConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactsContainerConfiguration)(nil))
}

func (i *artifactsContainerConfigurationPtrType) ToArtifactsContainerConfigurationPtrOutput() ArtifactsContainerConfigurationPtrOutput {
	return i.ToArtifactsContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i *artifactsContainerConfigurationPtrType) ToArtifactsContainerConfigurationPtrOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactsContainerConfigurationPtrOutput)
}

// ArtifactsContainerConfigurationArrayInput is an input type that accepts ArtifactsContainerConfigurationArray and ArtifactsContainerConfigurationArrayOutput values.
// You can construct a concrete instance of `ArtifactsContainerConfigurationArrayInput` via:
//
//          ArtifactsContainerConfigurationArray{ ArtifactsContainerConfigurationArgs{...} }
type ArtifactsContainerConfigurationArrayInput interface {
	pulumi.Input

	ToArtifactsContainerConfigurationArrayOutput() ArtifactsContainerConfigurationArrayOutput
	ToArtifactsContainerConfigurationArrayOutputWithContext(context.Context) ArtifactsContainerConfigurationArrayOutput
}

type ArtifactsContainerConfigurationArray []ArtifactsContainerConfigurationInput

func (ArtifactsContainerConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ArtifactsContainerConfiguration)(nil)).Elem()
}

func (i ArtifactsContainerConfigurationArray) ToArtifactsContainerConfigurationArrayOutput() ArtifactsContainerConfigurationArrayOutput {
	return i.ToArtifactsContainerConfigurationArrayOutputWithContext(context.Background())
}

func (i ArtifactsContainerConfigurationArray) ToArtifactsContainerConfigurationArrayOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactsContainerConfigurationArrayOutput)
}

// ArtifactsContainerConfigurationMapInput is an input type that accepts ArtifactsContainerConfigurationMap and ArtifactsContainerConfigurationMapOutput values.
// You can construct a concrete instance of `ArtifactsContainerConfigurationMapInput` via:
//
//          ArtifactsContainerConfigurationMap{ "key": ArtifactsContainerConfigurationArgs{...} }
type ArtifactsContainerConfigurationMapInput interface {
	pulumi.Input

	ToArtifactsContainerConfigurationMapOutput() ArtifactsContainerConfigurationMapOutput
	ToArtifactsContainerConfigurationMapOutputWithContext(context.Context) ArtifactsContainerConfigurationMapOutput
}

type ArtifactsContainerConfigurationMap map[string]ArtifactsContainerConfigurationInput

func (ArtifactsContainerConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ArtifactsContainerConfiguration)(nil)).Elem()
}

func (i ArtifactsContainerConfigurationMap) ToArtifactsContainerConfigurationMapOutput() ArtifactsContainerConfigurationMapOutput {
	return i.ToArtifactsContainerConfigurationMapOutputWithContext(context.Background())
}

func (i ArtifactsContainerConfigurationMap) ToArtifactsContainerConfigurationMapOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactsContainerConfigurationMapOutput)
}

type ArtifactsContainerConfigurationOutput struct {
	*pulumi.OutputState
}

func (ArtifactsContainerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArtifactsContainerConfiguration)(nil))
}

func (o ArtifactsContainerConfigurationOutput) ToArtifactsContainerConfigurationOutput() ArtifactsContainerConfigurationOutput {
	return o
}

func (o ArtifactsContainerConfigurationOutput) ToArtifactsContainerConfigurationOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationOutput {
	return o
}

func (o ArtifactsContainerConfigurationOutput) ToArtifactsContainerConfigurationPtrOutput() ArtifactsContainerConfigurationPtrOutput {
	return o.ToArtifactsContainerConfigurationPtrOutputWithContext(context.Background())
}

func (o ArtifactsContainerConfigurationOutput) ToArtifactsContainerConfigurationPtrOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationPtrOutput {
	return o.ApplyT(func(v ArtifactsContainerConfiguration) *ArtifactsContainerConfiguration {
		return &v
	}).(ArtifactsContainerConfigurationPtrOutput)
}

type ArtifactsContainerConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (ArtifactsContainerConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactsContainerConfiguration)(nil))
}

func (o ArtifactsContainerConfigurationPtrOutput) ToArtifactsContainerConfigurationPtrOutput() ArtifactsContainerConfigurationPtrOutput {
	return o
}

func (o ArtifactsContainerConfigurationPtrOutput) ToArtifactsContainerConfigurationPtrOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationPtrOutput {
	return o
}

type ArtifactsContainerConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ArtifactsContainerConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ArtifactsContainerConfiguration)(nil))
}

func (o ArtifactsContainerConfigurationArrayOutput) ToArtifactsContainerConfigurationArrayOutput() ArtifactsContainerConfigurationArrayOutput {
	return o
}

func (o ArtifactsContainerConfigurationArrayOutput) ToArtifactsContainerConfigurationArrayOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationArrayOutput {
	return o
}

func (o ArtifactsContainerConfigurationArrayOutput) Index(i pulumi.IntInput) ArtifactsContainerConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ArtifactsContainerConfiguration {
		return vs[0].([]ArtifactsContainerConfiguration)[vs[1].(int)]
	}).(ArtifactsContainerConfigurationOutput)
}

type ArtifactsContainerConfigurationMapOutput struct{ *pulumi.OutputState }

func (ArtifactsContainerConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ArtifactsContainerConfiguration)(nil))
}

func (o ArtifactsContainerConfigurationMapOutput) ToArtifactsContainerConfigurationMapOutput() ArtifactsContainerConfigurationMapOutput {
	return o
}

func (o ArtifactsContainerConfigurationMapOutput) ToArtifactsContainerConfigurationMapOutputWithContext(ctx context.Context) ArtifactsContainerConfigurationMapOutput {
	return o
}

func (o ArtifactsContainerConfigurationMapOutput) MapIndex(k pulumi.StringInput) ArtifactsContainerConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ArtifactsContainerConfiguration {
		return vs[0].(map[string]ArtifactsContainerConfiguration)[vs[1].(string)]
	}).(ArtifactsContainerConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(ArtifactsContainerConfigurationOutput{})
	pulumi.RegisterOutputType(ArtifactsContainerConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ArtifactsContainerConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ArtifactsContainerConfigurationMapOutput{})
}