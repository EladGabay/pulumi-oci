// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the View resource in Oracle Cloud Infrastructure DNS service.
//
// Creates a new view in the specified compartment. Requires a `PRIVATE` scope query parameter.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/dns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dns.NewView(ctx, "testView", &dns.ViewArgs{
// 			CompartmentId: pulumi.Any(_var.Compartment_id),
// 			Scope:         pulumi.String("PRIVATE"),
// 			DefinedTags:   pulumi.Any(_var.View_defined_tags),
// 			DisplayName:   pulumi.Any(_var.View_display_name),
// 			FreeformTags:  pulumi.Any(_var.View_freeform_tags),
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
// For legacy Views that were created without using `scope`, these Views can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:dns/view:View test_view "id"
// ```
//
//  For Views created using `scope`, these Views can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:dns/view:View test_view "viewId/{viewId}/scope/{scope}"
// ```
type View struct {
	pulumi.CustomResourceState

	// (Updatable) The OCID of the owning compartment.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) The display name of the view.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected pulumi.BoolOutput `pulumi:"isProtected"`
	// Value must be `PRIVATE` when creating a view for private zones.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// The canonical absolute URL of the resource.
	Self pulumi.StringOutput `pulumi:"self"`
	// The current state of the resource.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
}

// NewView registers a new resource with the given unique name, arguments, and options.
func NewView(ctx *pulumi.Context,
	name string, args *ViewArgs, opts ...pulumi.ResourceOption) (*View, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CompartmentId == nil {
		return nil, errors.New("invalid value for required argument 'CompartmentId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	var resource View
	err := ctx.RegisterResource("oci:dns/view:View", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetView gets an existing View resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetView(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ViewState, opts ...pulumi.ResourceOption) (*View, error) {
	var resource View
	err := ctx.ReadResource("oci:dns/view:View", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering View resources.
type viewState struct {
	// (Updatable) The OCID of the owning compartment.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The display name of the view.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `pulumi:"isProtected"`
	// Value must be `PRIVATE` when creating a view for private zones.
	Scope *string `pulumi:"scope"`
	// The canonical absolute URL of the resource.
	Self *string `pulumi:"self"`
	// The current state of the resource.
	State *string `pulumi:"state"`
	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated *string `pulumi:"timeCreated"`
	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeUpdated *string `pulumi:"timeUpdated"`
}

type ViewState struct {
	// (Updatable) The OCID of the owning compartment.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// (Updatable) The display name of the view.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags pulumi.MapInput
	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected pulumi.BoolPtrInput
	// Value must be `PRIVATE` when creating a view for private zones.
	Scope pulumi.StringPtrInput
	// The canonical absolute URL of the resource.
	Self pulumi.StringPtrInput
	// The current state of the resource.
	State pulumi.StringPtrInput
	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeCreated pulumi.StringPtrInput
	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ" format with a Z offset, as defined by RFC 3339.
	TimeUpdated pulumi.StringPtrInput
}

func (ViewState) ElementType() reflect.Type {
	return reflect.TypeOf((*viewState)(nil)).Elem()
}

type viewArgs struct {
	// (Updatable) The OCID of the owning compartment.
	CompartmentId string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) The display name of the view.
	DisplayName *string `pulumi:"displayName"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// Value must be `PRIVATE` when creating a view for private zones.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a View resource.
type ViewArgs struct {
	// (Updatable) The OCID of the owning compartment.
	CompartmentId pulumi.StringInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	DefinedTags pulumi.MapInput
	// (Updatable) The display name of the view.
	DisplayName pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	FreeformTags pulumi.MapInput
	// Value must be `PRIVATE` when creating a view for private zones.
	Scope pulumi.StringInput
}

func (ViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*viewArgs)(nil)).Elem()
}

type ViewInput interface {
	pulumi.Input

	ToViewOutput() ViewOutput
	ToViewOutputWithContext(ctx context.Context) ViewOutput
}

func (*View) ElementType() reflect.Type {
	return reflect.TypeOf((*View)(nil))
}

func (i *View) ToViewOutput() ViewOutput {
	return i.ToViewOutputWithContext(context.Background())
}

func (i *View) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewOutput)
}

func (i *View) ToViewPtrOutput() ViewPtrOutput {
	return i.ToViewPtrOutputWithContext(context.Background())
}

func (i *View) ToViewPtrOutputWithContext(ctx context.Context) ViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewPtrOutput)
}

type ViewPtrInput interface {
	pulumi.Input

	ToViewPtrOutput() ViewPtrOutput
	ToViewPtrOutputWithContext(ctx context.Context) ViewPtrOutput
}

type viewPtrType ViewArgs

func (*viewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil))
}

func (i *viewPtrType) ToViewPtrOutput() ViewPtrOutput {
	return i.ToViewPtrOutputWithContext(context.Background())
}

func (i *viewPtrType) ToViewPtrOutputWithContext(ctx context.Context) ViewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewPtrOutput)
}

// ViewArrayInput is an input type that accepts ViewArray and ViewArrayOutput values.
// You can construct a concrete instance of `ViewArrayInput` via:
//
//          ViewArray{ ViewArgs{...} }
type ViewArrayInput interface {
	pulumi.Input

	ToViewArrayOutput() ViewArrayOutput
	ToViewArrayOutputWithContext(context.Context) ViewArrayOutput
}

type ViewArray []ViewInput

func (ViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*View)(nil)).Elem()
}

func (i ViewArray) ToViewArrayOutput() ViewArrayOutput {
	return i.ToViewArrayOutputWithContext(context.Background())
}

func (i ViewArray) ToViewArrayOutputWithContext(ctx context.Context) ViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewArrayOutput)
}

// ViewMapInput is an input type that accepts ViewMap and ViewMapOutput values.
// You can construct a concrete instance of `ViewMapInput` via:
//
//          ViewMap{ "key": ViewArgs{...} }
type ViewMapInput interface {
	pulumi.Input

	ToViewMapOutput() ViewMapOutput
	ToViewMapOutputWithContext(context.Context) ViewMapOutput
}

type ViewMap map[string]ViewInput

func (ViewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*View)(nil)).Elem()
}

func (i ViewMap) ToViewMapOutput() ViewMapOutput {
	return i.ToViewMapOutputWithContext(context.Background())
}

func (i ViewMap) ToViewMapOutputWithContext(ctx context.Context) ViewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ViewMapOutput)
}

type ViewOutput struct {
	*pulumi.OutputState
}

func (ViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*View)(nil))
}

func (o ViewOutput) ToViewOutput() ViewOutput {
	return o
}

func (o ViewOutput) ToViewOutputWithContext(ctx context.Context) ViewOutput {
	return o
}

func (o ViewOutput) ToViewPtrOutput() ViewPtrOutput {
	return o.ToViewPtrOutputWithContext(context.Background())
}

func (o ViewOutput) ToViewPtrOutputWithContext(ctx context.Context) ViewPtrOutput {
	return o.ApplyT(func(v View) *View {
		return &v
	}).(ViewPtrOutput)
}

type ViewPtrOutput struct {
	*pulumi.OutputState
}

func (ViewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**View)(nil))
}

func (o ViewPtrOutput) ToViewPtrOutput() ViewPtrOutput {
	return o
}

func (o ViewPtrOutput) ToViewPtrOutputWithContext(ctx context.Context) ViewPtrOutput {
	return o
}

type ViewArrayOutput struct{ *pulumi.OutputState }

func (ViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]View)(nil))
}

func (o ViewArrayOutput) ToViewArrayOutput() ViewArrayOutput {
	return o
}

func (o ViewArrayOutput) ToViewArrayOutputWithContext(ctx context.Context) ViewArrayOutput {
	return o
}

func (o ViewArrayOutput) Index(i pulumi.IntInput) ViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) View {
		return vs[0].([]View)[vs[1].(int)]
	}).(ViewOutput)
}

type ViewMapOutput struct{ *pulumi.OutputState }

func (ViewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]View)(nil))
}

func (o ViewMapOutput) ToViewMapOutput() ViewMapOutput {
	return o
}

func (o ViewMapOutput) ToViewMapOutputWithContext(ctx context.Context) ViewMapOutput {
	return o
}

func (o ViewMapOutput) MapIndex(k pulumi.StringInput) ViewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) View {
		return vs[0].(map[string]View)[vs[1].(string)]
	}).(ViewOutput)
}

func init() {
	pulumi.RegisterOutputType(ViewOutput{})
	pulumi.RegisterOutputType(ViewPtrOutput{})
	pulumi.RegisterOutputType(ViewArrayOutput{})
	pulumi.RegisterOutputType(ViewMapOutput{})
}