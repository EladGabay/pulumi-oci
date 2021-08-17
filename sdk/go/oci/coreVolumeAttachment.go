// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource provides the Volume Attachment resource in Oracle Cloud Infrastructure Core service.
//
// Attaches the specified storage volume to the specified instance.
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
// 		_, err := oci.NewCoreVolumeAttachment(ctx, "testVolumeAttachment", &oci.CoreVolumeAttachmentArgs{
// 			AttachmentType:                 pulumi.Any(_var.Volume_attachment_attachment_type),
// 			InstanceId:                     pulumi.Any(oci_core_instance.Test_instance.Id),
// 			VolumeId:                       pulumi.Any(oci_core_volume.Test_volume.Id),
// 			Device:                         pulumi.Any(_var.Volume_attachment_device),
// 			DisplayName:                    pulumi.Any(_var.Volume_attachment_display_name),
// 			EncryptionInTransitType:        pulumi.Any(_var.Volume_attachment_encryption_in_transit_type),
// 			IsPvEncryptionInTransitEnabled: pulumi.Any(_var.Volume_attachment_is_pv_encryption_in_transit_enabled),
// 			IsReadOnly:                     pulumi.Any(_var.Volume_attachment_is_read_only),
// 			IsShareable:                    pulumi.Any(_var.Volume_attachment_is_shareable),
// 			UseChap:                        pulumi.Any(_var.Volume_attachment_use_chap),
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
// VolumeAttachments can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import oci:index/coreVolumeAttachment:CoreVolumeAttachment test_volume_attachment "id"
// ```
type CoreVolumeAttachment struct {
	pulumi.CustomResourceState

	// The type of volume. The only supported values are "iscsi" and "paravirtualized".
	AttachmentType pulumi.StringOutput `pulumi:"attachmentType"`
	// The availability domain of an instance.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringOutput `pulumi:"availabilityDomain"`
	// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)
	ChapSecret pulumi.StringOutput `pulumi:"chapSecret"`
	// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name. See [RFC 1994](https://tools.ietf.org/html/rfc1994) for more on CHAP.  Example: `ocid1.volume.oc1.phx.<unique_ID>`
	ChapUsername pulumi.StringOutput `pulumi:"chapUsername"`
	// The OCID of the compartment.
	//
	// Deprecated: The 'compartment_id' field has been deprecated and may be removed in a future version. Do not use this field.
	CompartmentId pulumi.StringOutput `pulumi:"compartmentId"`
	// The device name.
	Device pulumi.StringOutput `pulumi:"device"`
	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Refer the top-level definition of encryptionInTransitType. The default value is NONE.
	EncryptionInTransitType pulumi.StringOutput `pulumi:"encryptionInTransitType"`
	// The OCID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The volume's iSCSI IP address.  Example: `169.254.2.2`
	Ipv4 pulumi.StringOutput `pulumi:"ipv4"`
	// The target volume's iSCSI Qualified Name in the format defined by [RFC 3720](https://tools.ietf.org/html/rfc3720#page-32).  Example: `iqn.2015-12.com.oracleiaas:40b7ee03-883f-46c6-a951-63d2841d2195`
	Iqn pulumi.StringOutput `pulumi:"iqn"`
	// Whether the attachment is multipath or not.
	IsMultipath pulumi.BoolOutput `pulumi:"isMultipath"`
	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled pulumi.BoolOutput `pulumi:"isPvEncryptionInTransitEnabled"`
	// Whether the attachment was created in read-only mode.
	IsReadOnly pulumi.BoolOutput `pulumi:"isReadOnly"`
	// Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
	IsShareable pulumi.BoolOutput `pulumi:"isShareable"`
	// The iscsi login state of the volume attachment. For a multipath volume attachment, all iscsi sessions need to be all logged-in or logged-out to be in logged-in or logged-out state.
	IscsiLoginState pulumi.StringOutput `pulumi:"iscsiLoginState"`
	// A list of secondary multipath devices
	MultipathDevices CoreVolumeAttachmentMultipathDeviceArrayOutput `pulumi:"multipathDevices"`
	// The volume's iSCSI port, usually port 860 or 3260.  Example: `3260`
	Port pulumi.IntOutput `pulumi:"port"`
	// The current state of the volume attachment.
	State pulumi.StringOutput `pulumi:"state"`
	// The date and time the volume was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap pulumi.BoolOutput `pulumi:"useChap"`
	// The OCID of the volume.
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewCoreVolumeAttachment registers a new resource with the given unique name, arguments, and options.
func NewCoreVolumeAttachment(ctx *pulumi.Context,
	name string, args *CoreVolumeAttachmentArgs, opts ...pulumi.ResourceOption) (*CoreVolumeAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttachmentType == nil {
		return nil, errors.New("invalid value for required argument 'AttachmentType'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.VolumeId == nil {
		return nil, errors.New("invalid value for required argument 'VolumeId'")
	}
	var resource CoreVolumeAttachment
	err := ctx.RegisterResource("oci:index/coreVolumeAttachment:CoreVolumeAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCoreVolumeAttachment gets an existing CoreVolumeAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCoreVolumeAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CoreVolumeAttachmentState, opts ...pulumi.ResourceOption) (*CoreVolumeAttachment, error) {
	var resource CoreVolumeAttachment
	err := ctx.ReadResource("oci:index/coreVolumeAttachment:CoreVolumeAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CoreVolumeAttachment resources.
type coreVolumeAttachmentState struct {
	// The type of volume. The only supported values are "iscsi" and "paravirtualized".
	AttachmentType *string `pulumi:"attachmentType"`
	// The availability domain of an instance.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `pulumi:"availabilityDomain"`
	// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)
	ChapSecret *string `pulumi:"chapSecret"`
	// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name. See [RFC 1994](https://tools.ietf.org/html/rfc1994) for more on CHAP.  Example: `ocid1.volume.oc1.phx.<unique_ID>`
	ChapUsername *string `pulumi:"chapUsername"`
	// The OCID of the compartment.
	//
	// Deprecated: The 'compartment_id' field has been deprecated and may be removed in a future version. Do not use this field.
	CompartmentId *string `pulumi:"compartmentId"`
	// The device name.
	Device *string `pulumi:"device"`
	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// Refer the top-level definition of encryptionInTransitType. The default value is NONE.
	EncryptionInTransitType *string `pulumi:"encryptionInTransitType"`
	// The OCID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The volume's iSCSI IP address.  Example: `169.254.2.2`
	Ipv4 *string `pulumi:"ipv4"`
	// The target volume's iSCSI Qualified Name in the format defined by [RFC 3720](https://tools.ietf.org/html/rfc3720#page-32).  Example: `iqn.2015-12.com.oracleiaas:40b7ee03-883f-46c6-a951-63d2841d2195`
	Iqn *string `pulumi:"iqn"`
	// Whether the attachment is multipath or not.
	IsMultipath *bool `pulumi:"isMultipath"`
	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `pulumi:"isPvEncryptionInTransitEnabled"`
	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `pulumi:"isReadOnly"`
	// Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `pulumi:"isShareable"`
	// The iscsi login state of the volume attachment. For a multipath volume attachment, all iscsi sessions need to be all logged-in or logged-out to be in logged-in or logged-out state.
	IscsiLoginState *string `pulumi:"iscsiLoginState"`
	// A list of secondary multipath devices
	MultipathDevices []CoreVolumeAttachmentMultipathDevice `pulumi:"multipathDevices"`
	// The volume's iSCSI port, usually port 860 or 3260.  Example: `3260`
	Port *int `pulumi:"port"`
	// The current state of the volume attachment.
	State *string `pulumi:"state"`
	// The date and time the volume was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *string `pulumi:"timeCreated"`
	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap *bool `pulumi:"useChap"`
	// The OCID of the volume.
	VolumeId *string `pulumi:"volumeId"`
}

type CoreVolumeAttachmentState struct {
	// The type of volume. The only supported values are "iscsi" and "paravirtualized".
	AttachmentType pulumi.StringPtrInput
	// The availability domain of an instance.  Example: `Uocm:PHX-AD-1`
	AvailabilityDomain pulumi.StringPtrInput
	// The Challenge-Handshake-Authentication-Protocol (CHAP) secret valid for the associated CHAP user name. (Also called the "CHAP password".)
	ChapSecret pulumi.StringPtrInput
	// The volume's system-generated Challenge-Handshake-Authentication-Protocol (CHAP) user name. See [RFC 1994](https://tools.ietf.org/html/rfc1994) for more on CHAP.  Example: `ocid1.volume.oc1.phx.<unique_ID>`
	ChapUsername pulumi.StringPtrInput
	// The OCID of the compartment.
	//
	// Deprecated: The 'compartment_id' field has been deprecated and may be removed in a future version. Do not use this field.
	CompartmentId pulumi.StringPtrInput
	// The device name.
	Device pulumi.StringPtrInput
	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// Refer the top-level definition of encryptionInTransitType. The default value is NONE.
	EncryptionInTransitType pulumi.StringPtrInput
	// The OCID of the instance.
	InstanceId pulumi.StringPtrInput
	// The volume's iSCSI IP address.  Example: `169.254.2.2`
	Ipv4 pulumi.StringPtrInput
	// The target volume's iSCSI Qualified Name in the format defined by [RFC 3720](https://tools.ietf.org/html/rfc3720#page-32).  Example: `iqn.2015-12.com.oracleiaas:40b7ee03-883f-46c6-a951-63d2841d2195`
	Iqn pulumi.StringPtrInput
	// Whether the attachment is multipath or not.
	IsMultipath pulumi.BoolPtrInput
	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled pulumi.BoolPtrInput
	// Whether the attachment was created in read-only mode.
	IsReadOnly pulumi.BoolPtrInput
	// Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
	IsShareable pulumi.BoolPtrInput
	// The iscsi login state of the volume attachment. For a multipath volume attachment, all iscsi sessions need to be all logged-in or logged-out to be in logged-in or logged-out state.
	IscsiLoginState pulumi.StringPtrInput
	// A list of secondary multipath devices
	MultipathDevices CoreVolumeAttachmentMultipathDeviceArrayInput
	// The volume's iSCSI port, usually port 860 or 3260.  Example: `3260`
	Port pulumi.IntPtrInput
	// The current state of the volume attachment.
	State pulumi.StringPtrInput
	// The date and time the volume was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
	TimeCreated pulumi.StringPtrInput
	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap pulumi.BoolPtrInput
	// The OCID of the volume.
	VolumeId pulumi.StringPtrInput
}

func (CoreVolumeAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*coreVolumeAttachmentState)(nil)).Elem()
}

type coreVolumeAttachmentArgs struct {
	// The type of volume. The only supported values are "iscsi" and "paravirtualized".
	AttachmentType string `pulumi:"attachmentType"`
	// The OCID of the compartment.
	//
	// Deprecated: The 'compartment_id' field has been deprecated and may be removed in a future version. Do not use this field.
	CompartmentId *string `pulumi:"compartmentId"`
	// The device name.
	Device *string `pulumi:"device"`
	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `pulumi:"displayName"`
	// Refer the top-level definition of encryptionInTransitType. The default value is NONE.
	EncryptionInTransitType *string `pulumi:"encryptionInTransitType"`
	// The OCID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled *bool `pulumi:"isPvEncryptionInTransitEnabled"`
	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `pulumi:"isReadOnly"`
	// Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `pulumi:"isShareable"`
	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap *bool `pulumi:"useChap"`
	// The OCID of the volume.
	VolumeId string `pulumi:"volumeId"`
}

// The set of arguments for constructing a CoreVolumeAttachment resource.
type CoreVolumeAttachmentArgs struct {
	// The type of volume. The only supported values are "iscsi" and "paravirtualized".
	AttachmentType pulumi.StringInput
	// The OCID of the compartment.
	//
	// Deprecated: The 'compartment_id' field has been deprecated and may be removed in a future version. Do not use this field.
	CompartmentId pulumi.StringPtrInput
	// The device name.
	Device pulumi.StringPtrInput
	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName pulumi.StringPtrInput
	// Refer the top-level definition of encryptionInTransitType. The default value is NONE.
	EncryptionInTransitType pulumi.StringPtrInput
	// The OCID of the instance.
	InstanceId pulumi.StringInput
	// Whether to enable in-transit encryption for the data volume's paravirtualized attachment. The default value is false.
	IsPvEncryptionInTransitEnabled pulumi.BoolPtrInput
	// Whether the attachment was created in read-only mode.
	IsReadOnly pulumi.BoolPtrInput
	// Whether the attachment should be created in shareable mode. If an attachment is created in shareable mode, then other instances can attach the same volume, provided that they also create their attachments in shareable mode. Only certain volume types can be attached in shareable mode. Defaults to false if not specified.
	IsShareable pulumi.BoolPtrInput
	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap pulumi.BoolPtrInput
	// The OCID of the volume.
	VolumeId pulumi.StringInput
}

func (CoreVolumeAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*coreVolumeAttachmentArgs)(nil)).Elem()
}

type CoreVolumeAttachmentInput interface {
	pulumi.Input

	ToCoreVolumeAttachmentOutput() CoreVolumeAttachmentOutput
	ToCoreVolumeAttachmentOutputWithContext(ctx context.Context) CoreVolumeAttachmentOutput
}

func (*CoreVolumeAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreVolumeAttachment)(nil))
}

func (i *CoreVolumeAttachment) ToCoreVolumeAttachmentOutput() CoreVolumeAttachmentOutput {
	return i.ToCoreVolumeAttachmentOutputWithContext(context.Background())
}

func (i *CoreVolumeAttachment) ToCoreVolumeAttachmentOutputWithContext(ctx context.Context) CoreVolumeAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreVolumeAttachmentOutput)
}

func (i *CoreVolumeAttachment) ToCoreVolumeAttachmentPtrOutput() CoreVolumeAttachmentPtrOutput {
	return i.ToCoreVolumeAttachmentPtrOutputWithContext(context.Background())
}

func (i *CoreVolumeAttachment) ToCoreVolumeAttachmentPtrOutputWithContext(ctx context.Context) CoreVolumeAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreVolumeAttachmentPtrOutput)
}

type CoreVolumeAttachmentPtrInput interface {
	pulumi.Input

	ToCoreVolumeAttachmentPtrOutput() CoreVolumeAttachmentPtrOutput
	ToCoreVolumeAttachmentPtrOutputWithContext(ctx context.Context) CoreVolumeAttachmentPtrOutput
}

type coreVolumeAttachmentPtrType CoreVolumeAttachmentArgs

func (*coreVolumeAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreVolumeAttachment)(nil))
}

func (i *coreVolumeAttachmentPtrType) ToCoreVolumeAttachmentPtrOutput() CoreVolumeAttachmentPtrOutput {
	return i.ToCoreVolumeAttachmentPtrOutputWithContext(context.Background())
}

func (i *coreVolumeAttachmentPtrType) ToCoreVolumeAttachmentPtrOutputWithContext(ctx context.Context) CoreVolumeAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreVolumeAttachmentPtrOutput)
}

// CoreVolumeAttachmentArrayInput is an input type that accepts CoreVolumeAttachmentArray and CoreVolumeAttachmentArrayOutput values.
// You can construct a concrete instance of `CoreVolumeAttachmentArrayInput` via:
//
//          CoreVolumeAttachmentArray{ CoreVolumeAttachmentArgs{...} }
type CoreVolumeAttachmentArrayInput interface {
	pulumi.Input

	ToCoreVolumeAttachmentArrayOutput() CoreVolumeAttachmentArrayOutput
	ToCoreVolumeAttachmentArrayOutputWithContext(context.Context) CoreVolumeAttachmentArrayOutput
}

type CoreVolumeAttachmentArray []CoreVolumeAttachmentInput

func (CoreVolumeAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CoreVolumeAttachment)(nil)).Elem()
}

func (i CoreVolumeAttachmentArray) ToCoreVolumeAttachmentArrayOutput() CoreVolumeAttachmentArrayOutput {
	return i.ToCoreVolumeAttachmentArrayOutputWithContext(context.Background())
}

func (i CoreVolumeAttachmentArray) ToCoreVolumeAttachmentArrayOutputWithContext(ctx context.Context) CoreVolumeAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreVolumeAttachmentArrayOutput)
}

// CoreVolumeAttachmentMapInput is an input type that accepts CoreVolumeAttachmentMap and CoreVolumeAttachmentMapOutput values.
// You can construct a concrete instance of `CoreVolumeAttachmentMapInput` via:
//
//          CoreVolumeAttachmentMap{ "key": CoreVolumeAttachmentArgs{...} }
type CoreVolumeAttachmentMapInput interface {
	pulumi.Input

	ToCoreVolumeAttachmentMapOutput() CoreVolumeAttachmentMapOutput
	ToCoreVolumeAttachmentMapOutputWithContext(context.Context) CoreVolumeAttachmentMapOutput
}

type CoreVolumeAttachmentMap map[string]CoreVolumeAttachmentInput

func (CoreVolumeAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CoreVolumeAttachment)(nil)).Elem()
}

func (i CoreVolumeAttachmentMap) ToCoreVolumeAttachmentMapOutput() CoreVolumeAttachmentMapOutput {
	return i.ToCoreVolumeAttachmentMapOutputWithContext(context.Background())
}

func (i CoreVolumeAttachmentMap) ToCoreVolumeAttachmentMapOutputWithContext(ctx context.Context) CoreVolumeAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CoreVolumeAttachmentMapOutput)
}

type CoreVolumeAttachmentOutput struct {
	*pulumi.OutputState
}

func (CoreVolumeAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CoreVolumeAttachment)(nil))
}

func (o CoreVolumeAttachmentOutput) ToCoreVolumeAttachmentOutput() CoreVolumeAttachmentOutput {
	return o
}

func (o CoreVolumeAttachmentOutput) ToCoreVolumeAttachmentOutputWithContext(ctx context.Context) CoreVolumeAttachmentOutput {
	return o
}

func (o CoreVolumeAttachmentOutput) ToCoreVolumeAttachmentPtrOutput() CoreVolumeAttachmentPtrOutput {
	return o.ToCoreVolumeAttachmentPtrOutputWithContext(context.Background())
}

func (o CoreVolumeAttachmentOutput) ToCoreVolumeAttachmentPtrOutputWithContext(ctx context.Context) CoreVolumeAttachmentPtrOutput {
	return o.ApplyT(func(v CoreVolumeAttachment) *CoreVolumeAttachment {
		return &v
	}).(CoreVolumeAttachmentPtrOutput)
}

type CoreVolumeAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (CoreVolumeAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CoreVolumeAttachment)(nil))
}

func (o CoreVolumeAttachmentPtrOutput) ToCoreVolumeAttachmentPtrOutput() CoreVolumeAttachmentPtrOutput {
	return o
}

func (o CoreVolumeAttachmentPtrOutput) ToCoreVolumeAttachmentPtrOutputWithContext(ctx context.Context) CoreVolumeAttachmentPtrOutput {
	return o
}

type CoreVolumeAttachmentArrayOutput struct{ *pulumi.OutputState }

func (CoreVolumeAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CoreVolumeAttachment)(nil))
}

func (o CoreVolumeAttachmentArrayOutput) ToCoreVolumeAttachmentArrayOutput() CoreVolumeAttachmentArrayOutput {
	return o
}

func (o CoreVolumeAttachmentArrayOutput) ToCoreVolumeAttachmentArrayOutputWithContext(ctx context.Context) CoreVolumeAttachmentArrayOutput {
	return o
}

func (o CoreVolumeAttachmentArrayOutput) Index(i pulumi.IntInput) CoreVolumeAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CoreVolumeAttachment {
		return vs[0].([]CoreVolumeAttachment)[vs[1].(int)]
	}).(CoreVolumeAttachmentOutput)
}

type CoreVolumeAttachmentMapOutput struct{ *pulumi.OutputState }

func (CoreVolumeAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]CoreVolumeAttachment)(nil))
}

func (o CoreVolumeAttachmentMapOutput) ToCoreVolumeAttachmentMapOutput() CoreVolumeAttachmentMapOutput {
	return o
}

func (o CoreVolumeAttachmentMapOutput) ToCoreVolumeAttachmentMapOutputWithContext(ctx context.Context) CoreVolumeAttachmentMapOutput {
	return o
}

func (o CoreVolumeAttachmentMapOutput) MapIndex(k pulumi.StringInput) CoreVolumeAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) CoreVolumeAttachment {
		return vs[0].(map[string]CoreVolumeAttachment)[vs[1].(string)]
	}).(CoreVolumeAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(CoreVolumeAttachmentOutput{})
	pulumi.RegisterOutputType(CoreVolumeAttachmentPtrOutput{})
	pulumi.RegisterOutputType(CoreVolumeAttachmentArrayOutput{})
	pulumi.RegisterOutputType(CoreVolumeAttachmentMapOutput{})
}