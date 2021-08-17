// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides the list of Smtp Credentials in Oracle Cloud Infrastructure Identity service.
//
// Lists the SMTP credentials for the specified user. The returned object contains the credential's OCID,
// the SMTP user name but not the SMTP password. The SMTP password is returned only upon creation.
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
// 		_, err := oci.GetIdentitySmtpCredentials(ctx, &GetIdentitySmtpCredentialsArgs{
// 			UserId: oci_identity_user.Test_user.Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIdentitySmtpCredentials(ctx *pulumi.Context, args *GetIdentitySmtpCredentialsArgs, opts ...pulumi.InvokeOption) (*GetIdentitySmtpCredentialsResult, error) {
	var rv GetIdentitySmtpCredentialsResult
	err := ctx.Invoke("oci:index/getIdentitySmtpCredentials:GetIdentitySmtpCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetIdentitySmtpCredentials.
type GetIdentitySmtpCredentialsArgs struct {
	Filters []GetIdentitySmtpCredentialsFilter `pulumi:"filters"`
	// The OCID of the user.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by GetIdentitySmtpCredentials.
type GetIdentitySmtpCredentialsResult struct {
	Filters []GetIdentitySmtpCredentialsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of smtp_credentials.
	SmtpCredentials []GetIdentitySmtpCredentialsSmtpCredential `pulumi:"smtpCredentials"`
	// The OCID of the user the SMTP credential belongs to.
	UserId string `pulumi:"userId"`
}