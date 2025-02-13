// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package core

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Drg Route Table resource in Oracle Cloud Infrastructure Core service.
//
// Creates a new DRG route table for the specified DRG. Assign the DRG route table to a DRG attachment
// using the `UpdateDrgAttachment` or `CreateDrgAttachment` operations.
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
// 		_, err := core.NewDrgRouteTable(ctx, "testDrgRouteTable", &core.DrgRouteTableArgs{
// 			DrgId: pulumi.Any(oci_core_drg.Test_drg.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			DisplayName: pulumi.Any(_var.Drg_route_table_display_name),
// 			FreeformTags: pulumi.AnyMap{
// 				"Department": pulumi.Any("Finance"),
// 			},
// 			ImportDrgRouteDistributionId: pulumi.Any(oci_core_drg_route_distribution.Test_drg_route_distribution.Id),
// 			IsEcmpEnabled:                pulumi.Any(_var.Drg_route_table_is_ecmp_enabled),
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
// DrgRouteTables can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:core/drgRouteTable:DrgRouteTable test_drg_route_table "id"
// ```
type DrgRouteTable struct {
	pulumi.CustomResourceState

	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table is always in the same compartment as the DRG.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to.
	DrgId pulumi.StringOutput `pulumi:"drgId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements through referenced attachments are inserted into the DRG route table.
	ImportDrgRouteDistributionId pulumi.StringOutput `pulumi:"importDrgRouteDistributionId"`
	// (Updatable) If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises networks, enable ECMP on the DRG route table.
	IsEcmpEnabled pulumi.BoolOutput `pulumi:"isEcmpEnabled"`
	// (Updatable) An optional property when flipped disables the import of route Distribution by setting importDrgRouteDistributionId to null.
	RemoveImportTrigger pulumi.BoolPtrOutput `pulumi:"removeImportTrigger"`
	// The DRG route table's current state.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the DRG route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
}

// NewDrgRouteTable registers a new resource with the given unique name, arguments, and options.
func NewDrgRouteTable(ctx *pulumi.Context,
	name string, args *DrgRouteTableArgs, opts ...pulumi.ResourceOption) (*DrgRouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DrgId == nil {
		return nil, errors.New("invalid value for required argument 'DrgId'")
	}
	var resource DrgRouteTable
	err := ctx.RegisterResource("oci:core/drgRouteTable:DrgRouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDrgRouteTable gets an existing DrgRouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDrgRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DrgRouteTableState, opts ...pulumi.ResourceOption) (*DrgRouteTable, error) {
	var resource DrgRouteTable
	err := ctx.ReadResource("oci:core/drgRouteTable:DrgRouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DrgRouteTable resources.
type drgRouteTableState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table is always in the same compartment as the DRG.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to.
	DrgId *string `pulumi:"drgId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements through referenced attachments are inserted into the DRG route table.
	ImportDrgRouteDistributionId *string `pulumi:"importDrgRouteDistributionId"`
	// (Updatable) If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises networks, enable ECMP on the DRG route table.
	IsEcmpEnabled *bool `pulumi:"isEcmpEnabled"`
	// (Updatable) An optional property when flipped disables the import of route Distribution by setting importDrgRouteDistributionId to null.
	RemoveImportTrigger *bool `pulumi:"removeImportTrigger"`
	// The DRG route table's current state.
	State *string `pulumi:"state"`
	// The date and time the DRG route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
}

type DrgRouteTableState struct {
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the DRG is in. The DRG route table is always in the same compartment as the DRG.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to.
	DrgId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements through referenced attachments are inserted into the DRG route table.
	ImportDrgRouteDistributionId pulumi.StringPtrInput
	// (Updatable) If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises networks, enable ECMP on the DRG route table.
	IsEcmpEnabled pulumi.BoolPtrInput
	// (Updatable) An optional property when flipped disables the import of route Distribution by setting importDrgRouteDistributionId to null.
	RemoveImportTrigger pulumi.BoolPtrInput
	// The DRG route table's current state.
	State pulumi.StringPtrInput
	// The date and time the DRG route table was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
}

func (DrgRouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*drgRouteTableState)(nil)).Elem()
}

type drgRouteTableArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to.
	DrgId string `pulumi:"drgId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements through referenced attachments are inserted into the DRG route table.
	ImportDrgRouteDistributionId *string `pulumi:"importDrgRouteDistributionId"`
	// (Updatable) If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises networks, enable ECMP on the DRG route table.
	IsEcmpEnabled *bool `pulumi:"isEcmpEnabled"`
	// (Updatable) An optional property when flipped disables the import of route Distribution by setting importDrgRouteDistributionId to null.
	RemoveImportTrigger *bool `pulumi:"removeImportTrigger"`
}

// The set of arguments for constructing a DrgRouteTable resource.
type DrgRouteTableArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG the DRG route table belongs to.
	DrgId pulumi.StringInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the import route distribution used to specify how incoming route advertisements through referenced attachments are inserted into the DRG route table.
	ImportDrgRouteDistributionId pulumi.StringPtrInput
	// (Updatable) If you want traffic to be routed using ECMP across your virtual circuits or IPSec tunnels to your on-premises networks, enable ECMP on the DRG route table.
	IsEcmpEnabled pulumi.BoolPtrInput
	// (Updatable) An optional property when flipped disables the import of route Distribution by setting importDrgRouteDistributionId to null.
	RemoveImportTrigger pulumi.BoolPtrInput
}

func (DrgRouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*drgRouteTableArgs)(nil)).Elem()
}

type DrgRouteTableInput interface {
	pulumi.Input

	ToDrgRouteTableOutput() DrgRouteTableOutput
	ToDrgRouteTableOutputWithContext(ctx context.Context) DrgRouteTableOutput
}

func (*DrgRouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*DrgRouteTable)(nil))
}

func (i *DrgRouteTable) ToDrgRouteTableOutput() DrgRouteTableOutput {
	return i.ToDrgRouteTableOutputWithContext(context.Background())
}

func (i *DrgRouteTable) ToDrgRouteTableOutputWithContext(ctx context.Context) DrgRouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgRouteTableOutput)
}

func (i *DrgRouteTable) ToDrgRouteTablePtrOutput() DrgRouteTablePtrOutput {
	return i.ToDrgRouteTablePtrOutputWithContext(context.Background())
}

func (i *DrgRouteTable) ToDrgRouteTablePtrOutputWithContext(ctx context.Context) DrgRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgRouteTablePtrOutput)
}

type DrgRouteTablePtrInput interface {
	pulumi.Input

	ToDrgRouteTablePtrOutput() DrgRouteTablePtrOutput
	ToDrgRouteTablePtrOutputWithContext(ctx context.Context) DrgRouteTablePtrOutput
}

type drgRouteTablePtrType DrgRouteTableArgs

func (*drgRouteTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DrgRouteTable)(nil))
}

func (i *drgRouteTablePtrType) ToDrgRouteTablePtrOutput() DrgRouteTablePtrOutput {
	return i.ToDrgRouteTablePtrOutputWithContext(context.Background())
}

func (i *drgRouteTablePtrType) ToDrgRouteTablePtrOutputWithContext(ctx context.Context) DrgRouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgRouteTablePtrOutput)
}

// DrgRouteTableArrayInput is an input type that accepts DrgRouteTableArray and DrgRouteTableArrayOutput values.
// You can construct a concrete instance of `DrgRouteTableArrayInput` via:
//
//          DrgRouteTableArray{ DrgRouteTableArgs{...} }
type DrgRouteTableArrayInput interface {
	pulumi.Input

	ToDrgRouteTableArrayOutput() DrgRouteTableArrayOutput
	ToDrgRouteTableArrayOutputWithContext(context.Context) DrgRouteTableArrayOutput
}

type DrgRouteTableArray []DrgRouteTableInput

func (DrgRouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DrgRouteTable)(nil)).Elem()
}

func (i DrgRouteTableArray) ToDrgRouteTableArrayOutput() DrgRouteTableArrayOutput {
	return i.ToDrgRouteTableArrayOutputWithContext(context.Background())
}

func (i DrgRouteTableArray) ToDrgRouteTableArrayOutputWithContext(ctx context.Context) DrgRouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgRouteTableArrayOutput)
}

// DrgRouteTableMapInput is an input type that accepts DrgRouteTableMap and DrgRouteTableMapOutput values.
// You can construct a concrete instance of `DrgRouteTableMapInput` via:
//
//          DrgRouteTableMap{ "key": DrgRouteTableArgs{...} }
type DrgRouteTableMapInput interface {
	pulumi.Input

	ToDrgRouteTableMapOutput() DrgRouteTableMapOutput
	ToDrgRouteTableMapOutputWithContext(context.Context) DrgRouteTableMapOutput
}

type DrgRouteTableMap map[string]DrgRouteTableInput

func (DrgRouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DrgRouteTable)(nil)).Elem()
}

func (i DrgRouteTableMap) ToDrgRouteTableMapOutput() DrgRouteTableMapOutput {
	return i.ToDrgRouteTableMapOutputWithContext(context.Background())
}

func (i DrgRouteTableMap) ToDrgRouteTableMapOutputWithContext(ctx context.Context) DrgRouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DrgRouteTableMapOutput)
}

type DrgRouteTableOutput struct {
	*pulumi.OutputState
}

func (DrgRouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DrgRouteTable)(nil))
}

func (o DrgRouteTableOutput) ToDrgRouteTableOutput() DrgRouteTableOutput {
	return o
}

func (o DrgRouteTableOutput) ToDrgRouteTableOutputWithContext(ctx context.Context) DrgRouteTableOutput {
	return o
}

func (o DrgRouteTableOutput) ToDrgRouteTablePtrOutput() DrgRouteTablePtrOutput {
	return o.ToDrgRouteTablePtrOutputWithContext(context.Background())
}

func (o DrgRouteTableOutput) ToDrgRouteTablePtrOutputWithContext(ctx context.Context) DrgRouteTablePtrOutput {
	return o.ApplyT(func(v DrgRouteTable) *DrgRouteTable {
		return &v
	}).(DrgRouteTablePtrOutput)
}

type DrgRouteTablePtrOutput struct {
	*pulumi.OutputState
}

func (DrgRouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DrgRouteTable)(nil))
}

func (o DrgRouteTablePtrOutput) ToDrgRouteTablePtrOutput() DrgRouteTablePtrOutput {
	return o
}

func (o DrgRouteTablePtrOutput) ToDrgRouteTablePtrOutputWithContext(ctx context.Context) DrgRouteTablePtrOutput {
	return o
}

type DrgRouteTableArrayOutput struct{ *pulumi.OutputState }

func (DrgRouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DrgRouteTable)(nil))
}

func (o DrgRouteTableArrayOutput) ToDrgRouteTableArrayOutput() DrgRouteTableArrayOutput {
	return o
}

func (o DrgRouteTableArrayOutput) ToDrgRouteTableArrayOutputWithContext(ctx context.Context) DrgRouteTableArrayOutput {
	return o
}

func (o DrgRouteTableArrayOutput) Index(i pulumi.IntInput) DrgRouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DrgRouteTable {
		return vs[0].([]DrgRouteTable)[vs[1].(int)]
	}).(DrgRouteTableOutput)
}

type DrgRouteTableMapOutput struct{ *pulumi.OutputState }

func (DrgRouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DrgRouteTable)(nil))
}

func (o DrgRouteTableMapOutput) ToDrgRouteTableMapOutput() DrgRouteTableMapOutput {
	return o
}

func (o DrgRouteTableMapOutput) ToDrgRouteTableMapOutputWithContext(ctx context.Context) DrgRouteTableMapOutput {
	return o
}

func (o DrgRouteTableMapOutput) MapIndex(k pulumi.StringInput) DrgRouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DrgRouteTable {
		return vs[0].(map[string]DrgRouteTable)[vs[1].(string)]
	}).(DrgRouteTableOutput)
}

func init() {
	pulumi.RegisterOutputType(DrgRouteTableOutput{})
	pulumi.RegisterOutputType(DrgRouteTablePtrOutput{})
	pulumi.RegisterOutputType(DrgRouteTableArrayOutput{})
	pulumi.RegisterOutputType(DrgRouteTableMapOutput{})
}
