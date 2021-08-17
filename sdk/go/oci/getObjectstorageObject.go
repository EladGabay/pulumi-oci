// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package oci

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source provides details about a specific Object resource in Oracle Cloud Infrastructure Object Storage service.
//
// Gets the metadata and body of an object.
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
// 		opt0 := _var.Object_http_response_cache_control
// 		opt1 := _var.Object_http_response_content_disposition
// 		opt2 := _var.Object_http_response_content_encoding
// 		opt3 := _var.Object_http_response_content_language
// 		opt4 := _var.Object_http_response_content_type
// 		opt5 := _var.Object_http_response_expires
// 		opt6 := oci_objectstorage_version.Test_version.Id
// 		_, err := oci.GetObjectstorageObject(ctx, &GetObjectstorageObjectArgs{
// 			Bucket:                         _var.Object_bucket,
// 			Namespace:                      _var.Object_namespace,
// 			Object:                         _var.Object_object,
// 			HttpResponseCacheControl:       &opt0,
// 			HttpResponseContentDisposition: &opt1,
// 			HttpResponseContentEncoding:    &opt2,
// 			HttpResponseContentLanguage:    &opt3,
// 			HttpResponseContentType:        &opt4,
// 			HttpResponseExpires:            &opt5,
// 			VersionId:                      &opt6,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupObjectstorageObject(ctx *pulumi.Context, args *LookupObjectstorageObjectArgs, opts ...pulumi.InvokeOption) (*LookupObjectstorageObjectResult, error) {
	var rv LookupObjectstorageObjectResult
	err := ctx.Invoke("oci:index/getObjectstorageObject:GetObjectstorageObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking GetObjectstorageObject.
type LookupObjectstorageObjectArgs struct {
	Base64EncodeContent *bool `pulumi:"base64EncodeContent"`
	// The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
	Bucket string `pulumi:"bucket"`
	// The limit of the content length of the object body to download from the object store. The default is 1Mb.
	ContentLengthLimit *int `pulumi:"contentLengthLimit"`
	// Specify this query parameter to override the Cache-Control response header in the GetObject response.
	HttpResponseCacheControl *string `pulumi:"httpResponseCacheControl"`
	// Specify this query parameter to override the value of the Content-Disposition response header in the GetObject response.
	HttpResponseContentDisposition *string `pulumi:"httpResponseContentDisposition"`
	// Specify this query parameter to override the Content-Encoding response header in the GetObject response.
	HttpResponseContentEncoding *string `pulumi:"httpResponseContentEncoding"`
	// Specify this query parameter to override the Content-Language response header in the GetObject response.
	HttpResponseContentLanguage *string `pulumi:"httpResponseContentLanguage"`
	// Specify this query parameter to override the Content-Type response header in the GetObject response.
	HttpResponseContentType *string `pulumi:"httpResponseContentType"`
	// Specify this query parameter to override the Expires response header in the GetObject response.
	HttpResponseExpires *string `pulumi:"httpResponseExpires"`
	// The Object Storage namespace used for the request.
	Namespace string `pulumi:"namespace"`
	// The name of the object. Avoid entering confidential information. Example: `test/object1.log`
	Object string `pulumi:"object"`
	// VersionId used to identify a particular version of the object
	VersionId *string `pulumi:"versionId"`
}

// A collection of values returned by GetObjectstorageObject.
type LookupObjectstorageObjectResult struct {
	Base64EncodeContent *bool `pulumi:"base64EncodeContent"`
	// The name of the bucket. Avoid entering confidential information. Example: `my-new-bucket1`
	Bucket       string `pulumi:"bucket"`
	CacheControl string `pulumi:"cacheControl"`
	// The object to upload to the object store.
	Content            string `pulumi:"content"`
	ContentDisposition string `pulumi:"contentDisposition"`
	// The content encoding of the object.
	ContentEncoding string `pulumi:"contentEncoding"`
	// The content language of the object.
	ContentLanguage string `pulumi:"contentLanguage"`
	// The content length of the body.
	ContentLength      string `pulumi:"contentLength"`
	ContentLengthLimit *int   `pulumi:"contentLengthLimit"`
	// The base-64 encoded MD5 hash of the body.
	ContentMd5 string `pulumi:"contentMd5"`
	// The content type of the object.  Defaults to 'application/octet-stream' if not overridden during the PutObject call.
	ContentType                    string  `pulumi:"contentType"`
	HttpResponseCacheControl       *string `pulumi:"httpResponseCacheControl"`
	HttpResponseContentDisposition *string `pulumi:"httpResponseContentDisposition"`
	HttpResponseContentEncoding    *string `pulumi:"httpResponseContentEncoding"`
	HttpResponseContentLanguage    *string `pulumi:"httpResponseContentLanguage"`
	HttpResponseContentType        *string `pulumi:"httpResponseContentType"`
	HttpResponseExpires            *string `pulumi:"httpResponseExpires"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Optional user-defined metadata key and value. Note: Metadata keys are case-insensitive and all returned keys will be lower case.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The top-level namespace used for the request.
	Namespace string `pulumi:"namespace"`
	// The name of the object. Avoid entering confidential information. Example: `test/object1.log`
	Object string `pulumi:"object"`
	// The storage tier that the object is stored in.
	StorageTier string `pulumi:"storageTier"`
	VersionId   string `pulumi:"versionId"`
}