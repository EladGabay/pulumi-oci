// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package email

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Dkim resource in Oracle Cloud Infrastructure Email service.
//
// Creates a new DKIM for a email domain.
// This DKIM will sign all approved senders in the tenancy that are in this email domain.
// Best security practices indicate to periodically rotate the DKIM that is doing the signing.
// When a second DKIM is applied, all senders will seamlessly pick up the new key
// without interruption in signing.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-oci/sdk/go/oci/email"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := email.NewDkim(ctx, "testDkim", &email.DkimArgs{
// 			EmailDomainId: pulumi.Any(oci_email_email_domain.Test_email_domain.Id),
// 			DefinedTags: pulumi.AnyMap{
// 				"Operations.CostCenter": pulumi.Any("42"),
// 			},
// 			Description: pulumi.Any(_var.Dkim_description),
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
// Dkims can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:email/dkim:Dkim test_dkim "id"
// ```
type Dkim struct {
	pulumi.CustomResourceState

	// The DNS CNAME record value to provision to the DKIM DNS subdomain, when using the CNAME method for DKIM setup (preferred).
	CnameRecordValue pulumi.StringOutput `pulumi:"cnameRecordValue"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this DKIM.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapOutput `pulumi:"definedTags"`
	// (Updatable) A string that describes the details about the DKIM. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the DNS subdomain that must be provisioned to enable email recipients to verify DKIM signatures. It is usually created with a CNAME record set to the cnameRecordValue
	DnsSubdomainName pulumi.StringOutput `pulumi:"dnsSubdomainName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM.
	EmailDomainId pulumi.StringOutput `pulumi:"emailDomainId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapOutput `pulumi:"freeformTags"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource.
	LifecycleDetails pulumi.StringOutput `pulumi:"lifecycleDetails"`
	// The DKIM selector. This selector is required to be globally unique for this email domain. If you do not provide the selector, we will generate one for you. If you do provide the selector, we suggest adding a short region indicator to differentiate from your signing of emails in other regions you may be subscribed to. Selectors limited to ASCII characters may use alphanumeric, dash ("-"), and dot (".") characters. Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).
	Name pulumi.StringOutput `pulumi:"name"`
	// The current state of the DKIM.
	State pulumi.StringOutput `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapOutput `pulumi:"systemTags"`
	// The time the DKIM was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The time of the last change to the DKIM configuration, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated pulumi.StringOutput `pulumi:"timeUpdated"`
	// The DNS TXT record value to provision to the DKIM DNS subdomain in place of using a CNAME record. This is used in cases where a CNAME can not be used, such as when the cnameRecordValue would exceed the maximum length for a DNS entry. This can also be used by customers who have an existing procedure to directly provision TXT records for DKIM. Be aware that many DNS APIs will require you to break this string into segments of less than 255 characters.
	TxtRecordValue pulumi.StringOutput `pulumi:"txtRecordValue"`
}

// NewDkim registers a new resource with the given unique name, arguments, and options.
func NewDkim(ctx *pulumi.Context,
	name string, args *DkimArgs, opts ...pulumi.ResourceOption) (*Dkim, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EmailDomainId == nil {
		return nil, errors.New("invalid value for required argument 'EmailDomainId'")
	}
	var resource Dkim
	err := ctx.RegisterResource("oci:email/dkim:Dkim", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDkim gets an existing Dkim resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDkim(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DkimState, opts ...pulumi.ResourceOption) (*Dkim, error) {
	var resource Dkim
	err := ctx.ReadResource("oci:email/dkim:Dkim", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dkim resources.
type dkimState struct {
	// The DNS CNAME record value to provision to the DKIM DNS subdomain, when using the CNAME method for DKIM setup (preferred).
	CnameRecordValue *string `pulumi:"cnameRecordValue"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this DKIM.
	CompartmentId *string `pulumi:"compartmentId"`
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A string that describes the details about the DKIM. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description *string `pulumi:"description"`
	// The name of the DNS subdomain that must be provisioned to enable email recipients to verify DKIM signatures. It is usually created with a CNAME record set to the cnameRecordValue
	DnsSubdomainName *string `pulumi:"dnsSubdomainName"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM.
	EmailDomainId *string `pulumi:"emailDomainId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource.
	LifecycleDetails *string `pulumi:"lifecycleDetails"`
	// The DKIM selector. This selector is required to be globally unique for this email domain. If you do not provide the selector, we will generate one for you. If you do provide the selector, we suggest adding a short region indicator to differentiate from your signing of emails in other regions you may be subscribed to. Selectors limited to ASCII characters may use alphanumeric, dash ("-"), and dot (".") characters. Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).
	Name *string `pulumi:"name"`
	// The current state of the DKIM.
	State *string `pulumi:"state"`
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags map[string]interface{} `pulumi:"systemTags"`
	// The time the DKIM was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z`
	TimeCreated *string `pulumi:"timeCreated"`
	// The time of the last change to the DKIM configuration, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated *string `pulumi:"timeUpdated"`
	// The DNS TXT record value to provision to the DKIM DNS subdomain in place of using a CNAME record. This is used in cases where a CNAME can not be used, such as when the cnameRecordValue would exceed the maximum length for a DNS entry. This can also be used by customers who have an existing procedure to directly provision TXT records for DKIM. Be aware that many DNS APIs will require you to break this string into segments of less than 255 characters.
	TxtRecordValue *string `pulumi:"txtRecordValue"`
}

type DkimState struct {
	// The DNS CNAME record value to provision to the DKIM DNS subdomain, when using the CNAME method for DKIM setup (preferred).
	CnameRecordValue pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains this DKIM.
	CompartmentId pulumi.StringPtrInput
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A string that describes the details about the DKIM. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description pulumi.StringPtrInput
	// The name of the DNS subdomain that must be provisioned to enable email recipients to verify DKIM signatures. It is usually created with a CNAME record set to the cnameRecordValue
	DnsSubdomainName pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM.
	EmailDomainId pulumi.StringPtrInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource.
	LifecycleDetails pulumi.StringPtrInput
	// The DKIM selector. This selector is required to be globally unique for this email domain. If you do not provide the selector, we will generate one for you. If you do provide the selector, we suggest adding a short region indicator to differentiate from your signing of emails in other regions you may be subscribed to. Selectors limited to ASCII characters may use alphanumeric, dash ("-"), and dot (".") characters. Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).
	Name pulumi.StringPtrInput
	// The current state of the DKIM.
	State pulumi.StringPtrInput
	// Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
	SystemTags pulumi.MapInput
	// The time the DKIM was created. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".  Example: `2021-02-12T22:47:12.613Z`
	TimeCreated pulumi.StringPtrInput
	// The time of the last change to the DKIM configuration, due to a state change or an update operation. Times are expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format, "YYYY-MM-ddThh:mmZ".
	TimeUpdated pulumi.StringPtrInput
	// The DNS TXT record value to provision to the DKIM DNS subdomain in place of using a CNAME record. This is used in cases where a CNAME can not be used, such as when the cnameRecordValue would exceed the maximum length for a DNS entry. This can also be used by customers who have an existing procedure to directly provision TXT records for DKIM. Be aware that many DNS APIs will require you to break this string into segments of less than 255 characters.
	TxtRecordValue pulumi.StringPtrInput
}

func (DkimState) ElementType() reflect.Type {
	return reflect.TypeOf((*dkimState)(nil)).Elem()
}

type dkimArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags map[string]interface{} `pulumi:"definedTags"`
	// (Updatable) A string that describes the details about the DKIM. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description *string `pulumi:"description"`
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM.
	EmailDomainId string `pulumi:"emailDomainId"`
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags map[string]interface{} `pulumi:"freeformTags"`
	// The DKIM selector. This selector is required to be globally unique for this email domain. If you do not provide the selector, we will generate one for you. If you do provide the selector, we suggest adding a short region indicator to differentiate from your signing of emails in other regions you may be subscribed to. Selectors limited to ASCII characters may use alphanumeric, dash ("-"), and dot (".") characters. Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Dkim resource.
type DkimArgs struct {
	// (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	DefinedTags pulumi.MapInput
	// (Updatable) A string that describes the details about the DKIM. It does not have to be unique, and you can change it. Avoid entering confidential information.
	Description pulumi.StringPtrInput
	// The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the EmailDomain for this DKIM.
	EmailDomainId pulumi.StringInput
	// (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	FreeformTags pulumi.MapInput
	// The DKIM selector. This selector is required to be globally unique for this email domain. If you do not provide the selector, we will generate one for you. If you do provide the selector, we suggest adding a short region indicator to differentiate from your signing of emails in other regions you may be subscribed to. Selectors limited to ASCII characters may use alphanumeric, dash ("-"), and dot (".") characters. Non-ASCII selector names should adopt IDNA2008 normalization (RFC 5891-5892).
	Name pulumi.StringPtrInput
}

func (DkimArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dkimArgs)(nil)).Elem()
}

type DkimInput interface {
	pulumi.Input

	ToDkimOutput() DkimOutput
	ToDkimOutputWithContext(ctx context.Context) DkimOutput
}

func (*Dkim) ElementType() reflect.Type {
	return reflect.TypeOf((*Dkim)(nil))
}

func (i *Dkim) ToDkimOutput() DkimOutput {
	return i.ToDkimOutputWithContext(context.Background())
}

func (i *Dkim) ToDkimOutputWithContext(ctx context.Context) DkimOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DkimOutput)
}

func (i *Dkim) ToDkimPtrOutput() DkimPtrOutput {
	return i.ToDkimPtrOutputWithContext(context.Background())
}

func (i *Dkim) ToDkimPtrOutputWithContext(ctx context.Context) DkimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DkimPtrOutput)
}

type DkimPtrInput interface {
	pulumi.Input

	ToDkimPtrOutput() DkimPtrOutput
	ToDkimPtrOutputWithContext(ctx context.Context) DkimPtrOutput
}

type dkimPtrType DkimArgs

func (*dkimPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dkim)(nil))
}

func (i *dkimPtrType) ToDkimPtrOutput() DkimPtrOutput {
	return i.ToDkimPtrOutputWithContext(context.Background())
}

func (i *dkimPtrType) ToDkimPtrOutputWithContext(ctx context.Context) DkimPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DkimPtrOutput)
}

// DkimArrayInput is an input type that accepts DkimArray and DkimArrayOutput values.
// You can construct a concrete instance of `DkimArrayInput` via:
//
//          DkimArray{ DkimArgs{...} }
type DkimArrayInput interface {
	pulumi.Input

	ToDkimArrayOutput() DkimArrayOutput
	ToDkimArrayOutputWithContext(context.Context) DkimArrayOutput
}

type DkimArray []DkimInput

func (DkimArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dkim)(nil)).Elem()
}

func (i DkimArray) ToDkimArrayOutput() DkimArrayOutput {
	return i.ToDkimArrayOutputWithContext(context.Background())
}

func (i DkimArray) ToDkimArrayOutputWithContext(ctx context.Context) DkimArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DkimArrayOutput)
}

// DkimMapInput is an input type that accepts DkimMap and DkimMapOutput values.
// You can construct a concrete instance of `DkimMapInput` via:
//
//          DkimMap{ "key": DkimArgs{...} }
type DkimMapInput interface {
	pulumi.Input

	ToDkimMapOutput() DkimMapOutput
	ToDkimMapOutputWithContext(context.Context) DkimMapOutput
}

type DkimMap map[string]DkimInput

func (DkimMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dkim)(nil)).Elem()
}

func (i DkimMap) ToDkimMapOutput() DkimMapOutput {
	return i.ToDkimMapOutputWithContext(context.Background())
}

func (i DkimMap) ToDkimMapOutputWithContext(ctx context.Context) DkimMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DkimMapOutput)
}

type DkimOutput struct {
	*pulumi.OutputState
}

func (DkimOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dkim)(nil))
}

func (o DkimOutput) ToDkimOutput() DkimOutput {
	return o
}

func (o DkimOutput) ToDkimOutputWithContext(ctx context.Context) DkimOutput {
	return o
}

func (o DkimOutput) ToDkimPtrOutput() DkimPtrOutput {
	return o.ToDkimPtrOutputWithContext(context.Background())
}

func (o DkimOutput) ToDkimPtrOutputWithContext(ctx context.Context) DkimPtrOutput {
	return o.ApplyT(func(v Dkim) *Dkim {
		return &v
	}).(DkimPtrOutput)
}

type DkimPtrOutput struct {
	*pulumi.OutputState
}

func (DkimPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dkim)(nil))
}

func (o DkimPtrOutput) ToDkimPtrOutput() DkimPtrOutput {
	return o
}

func (o DkimPtrOutput) ToDkimPtrOutputWithContext(ctx context.Context) DkimPtrOutput {
	return o
}

type DkimArrayOutput struct{ *pulumi.OutputState }

func (DkimArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Dkim)(nil))
}

func (o DkimArrayOutput) ToDkimArrayOutput() DkimArrayOutput {
	return o
}

func (o DkimArrayOutput) ToDkimArrayOutputWithContext(ctx context.Context) DkimArrayOutput {
	return o
}

func (o DkimArrayOutput) Index(i pulumi.IntInput) DkimOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Dkim {
		return vs[0].([]Dkim)[vs[1].(int)]
	}).(DkimOutput)
}

type DkimMapOutput struct{ *pulumi.OutputState }

func (DkimMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Dkim)(nil))
}

func (o DkimMapOutput) ToDkimMapOutput() DkimMapOutput {
	return o
}

func (o DkimMapOutput) ToDkimMapOutputWithContext(ctx context.Context) DkimMapOutput {
	return o
}

func (o DkimMapOutput) MapIndex(k pulumi.StringInput) DkimOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Dkim {
		return vs[0].(map[string]Dkim)[vs[1].(string)]
	}).(DkimOutput)
}

func init() {
	pulumi.RegisterOutputType(DkimOutput{})
	pulumi.RegisterOutputType(DkimPtrOutput{})
	pulumi.RegisterOutputType(DkimArrayOutput{})
	pulumi.RegisterOutputType(DkimMapOutput{})
}