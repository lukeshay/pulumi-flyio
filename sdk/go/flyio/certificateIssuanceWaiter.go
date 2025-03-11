// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flyio

import (
	"context"
	"reflect"

	"errors"
	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/internal"
	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/time"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A component that waits for a Fly.io SSL/TLS certificate to be fully issued.
type CertificateIssuanceWaiter struct {
	pulumi.CustomResourceState

	// The name of the Fly app.
	App pulumi.StringOutput `pulumi:"app"`
	// The certificate authority used.
	CertificateAuthority pulumi.StringOutput `pulumi:"certificateAuthority"`
	// The Fly.io certificate ID.
	CertificateId pulumi.StringOutput `pulumi:"certificateId"`
	// The status of the certificate.
	ClientStatus pulumi.StringOutput `pulumi:"clientStatus"`
	// Expiration time for the ECDSA certificate.
	EcdsaExpiresAt time.TimePtrOutput `pulumi:"ecdsaExpiresAt"`
	// The hostname for the certificate.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The input arguments used for the certificate issuance waiter.
	Input CertificateIssuanceWaiterArgsTypeOutput `pulumi:"input"`
	// Whether the certificate is fully issued (has both ECDSA and RSA certificates).
	IsFullyIssued pulumi.BoolOutput `pulumi:"isFullyIssued"`
	// When the certificate was fully issued.
	IssuedAt time.TimeOutput `pulumi:"issuedAt"`
	// Expiration time for the RSA certificate.
	RsaExpiresAt time.TimePtrOutput `pulumi:"rsaExpiresAt"`
	// The timeout duration used for waiting.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
}

// NewCertificateIssuanceWaiter registers a new resource with the given unique name, arguments, and options.
func NewCertificateIssuanceWaiter(ctx *pulumi.Context,
	name string, args *CertificateIssuanceWaiterArgs, opts ...pulumi.ResourceOption) (*CertificateIssuanceWaiter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.App == nil {
		return nil, errors.New("invalid value for required argument 'App'")
	}
	if args.Hostname == nil {
		return nil, errors.New("invalid value for required argument 'Hostname'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CertificateIssuanceWaiter
	err := ctx.RegisterResource("flyio:index:CertificateIssuanceWaiter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateIssuanceWaiter gets an existing CertificateIssuanceWaiter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateIssuanceWaiter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateIssuanceWaiterState, opts ...pulumi.ResourceOption) (*CertificateIssuanceWaiter, error) {
	var resource CertificateIssuanceWaiter
	err := ctx.ReadResource("flyio:index:CertificateIssuanceWaiter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateIssuanceWaiter resources.
type certificateIssuanceWaiterState struct {
}

type CertificateIssuanceWaiterState struct {
}

func (CertificateIssuanceWaiterState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateIssuanceWaiterState)(nil)).Elem()
}

type certificateIssuanceWaiterArgs struct {
	// The name of the Fly app that the certificate is associated with.
	App string `pulumi:"app"`
	// The hostname for the certificate (e.g., example.com).
	Hostname string `pulumi:"hostname"`
	// The maximum time to wait for the certificate to be fully issued (in seconds).
	Timeout int `pulumi:"timeout"`
}

// The set of arguments for constructing a CertificateIssuanceWaiter resource.
type CertificateIssuanceWaiterArgs struct {
	// The name of the Fly app that the certificate is associated with.
	App pulumi.StringInput
	// The hostname for the certificate (e.g., example.com).
	Hostname pulumi.StringInput
	// The maximum time to wait for the certificate to be fully issued (in seconds).
	Timeout pulumi.IntInput
}

func (CertificateIssuanceWaiterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateIssuanceWaiterArgs)(nil)).Elem()
}

type CertificateIssuanceWaiterInput interface {
	pulumi.Input

	ToCertificateIssuanceWaiterOutput() CertificateIssuanceWaiterOutput
	ToCertificateIssuanceWaiterOutputWithContext(ctx context.Context) CertificateIssuanceWaiterOutput
}

func (*CertificateIssuanceWaiter) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateIssuanceWaiter)(nil)).Elem()
}

func (i *CertificateIssuanceWaiter) ToCertificateIssuanceWaiterOutput() CertificateIssuanceWaiterOutput {
	return i.ToCertificateIssuanceWaiterOutputWithContext(context.Background())
}

func (i *CertificateIssuanceWaiter) ToCertificateIssuanceWaiterOutputWithContext(ctx context.Context) CertificateIssuanceWaiterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateIssuanceWaiterOutput)
}

// CertificateIssuanceWaiterArrayInput is an input type that accepts CertificateIssuanceWaiterArray and CertificateIssuanceWaiterArrayOutput values.
// You can construct a concrete instance of `CertificateIssuanceWaiterArrayInput` via:
//
//	CertificateIssuanceWaiterArray{ CertificateIssuanceWaiterArgs{...} }
type CertificateIssuanceWaiterArrayInput interface {
	pulumi.Input

	ToCertificateIssuanceWaiterArrayOutput() CertificateIssuanceWaiterArrayOutput
	ToCertificateIssuanceWaiterArrayOutputWithContext(context.Context) CertificateIssuanceWaiterArrayOutput
}

type CertificateIssuanceWaiterArray []CertificateIssuanceWaiterInput

func (CertificateIssuanceWaiterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateIssuanceWaiter)(nil)).Elem()
}

func (i CertificateIssuanceWaiterArray) ToCertificateIssuanceWaiterArrayOutput() CertificateIssuanceWaiterArrayOutput {
	return i.ToCertificateIssuanceWaiterArrayOutputWithContext(context.Background())
}

func (i CertificateIssuanceWaiterArray) ToCertificateIssuanceWaiterArrayOutputWithContext(ctx context.Context) CertificateIssuanceWaiterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateIssuanceWaiterArrayOutput)
}

// CertificateIssuanceWaiterMapInput is an input type that accepts CertificateIssuanceWaiterMap and CertificateIssuanceWaiterMapOutput values.
// You can construct a concrete instance of `CertificateIssuanceWaiterMapInput` via:
//
//	CertificateIssuanceWaiterMap{ "key": CertificateIssuanceWaiterArgs{...} }
type CertificateIssuanceWaiterMapInput interface {
	pulumi.Input

	ToCertificateIssuanceWaiterMapOutput() CertificateIssuanceWaiterMapOutput
	ToCertificateIssuanceWaiterMapOutputWithContext(context.Context) CertificateIssuanceWaiterMapOutput
}

type CertificateIssuanceWaiterMap map[string]CertificateIssuanceWaiterInput

func (CertificateIssuanceWaiterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateIssuanceWaiter)(nil)).Elem()
}

func (i CertificateIssuanceWaiterMap) ToCertificateIssuanceWaiterMapOutput() CertificateIssuanceWaiterMapOutput {
	return i.ToCertificateIssuanceWaiterMapOutputWithContext(context.Background())
}

func (i CertificateIssuanceWaiterMap) ToCertificateIssuanceWaiterMapOutputWithContext(ctx context.Context) CertificateIssuanceWaiterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateIssuanceWaiterMapOutput)
}

type CertificateIssuanceWaiterOutput struct{ *pulumi.OutputState }

func (CertificateIssuanceWaiterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateIssuanceWaiter)(nil)).Elem()
}

func (o CertificateIssuanceWaiterOutput) ToCertificateIssuanceWaiterOutput() CertificateIssuanceWaiterOutput {
	return o
}

func (o CertificateIssuanceWaiterOutput) ToCertificateIssuanceWaiterOutputWithContext(ctx context.Context) CertificateIssuanceWaiterOutput {
	return o
}

// The name of the Fly app.
func (o CertificateIssuanceWaiterOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.StringOutput { return v.App }).(pulumi.StringOutput)
}

// The certificate authority used.
func (o CertificateIssuanceWaiterOutput) CertificateAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.StringOutput { return v.CertificateAuthority }).(pulumi.StringOutput)
}

// The Fly.io certificate ID.
func (o CertificateIssuanceWaiterOutput) CertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.StringOutput { return v.CertificateId }).(pulumi.StringOutput)
}

// The status of the certificate.
func (o CertificateIssuanceWaiterOutput) ClientStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.StringOutput { return v.ClientStatus }).(pulumi.StringOutput)
}

// Expiration time for the ECDSA certificate.
func (o CertificateIssuanceWaiterOutput) EcdsaExpiresAt() time.TimePtrOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) time.TimePtrOutput { return v.EcdsaExpiresAt }).(time.TimePtrOutput)
}

// The hostname for the certificate.
func (o CertificateIssuanceWaiterOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// The input arguments used for the certificate issuance waiter.
func (o CertificateIssuanceWaiterOutput) Input() CertificateIssuanceWaiterArgsTypeOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) CertificateIssuanceWaiterArgsTypeOutput { return v.Input }).(CertificateIssuanceWaiterArgsTypeOutput)
}

// Whether the certificate is fully issued (has both ECDSA and RSA certificates).
func (o CertificateIssuanceWaiterOutput) IsFullyIssued() pulumi.BoolOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.BoolOutput { return v.IsFullyIssued }).(pulumi.BoolOutput)
}

// When the certificate was fully issued.
func (o CertificateIssuanceWaiterOutput) IssuedAt() time.TimeOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) time.TimeOutput { return v.IssuedAt }).(time.TimeOutput)
}

// Expiration time for the RSA certificate.
func (o CertificateIssuanceWaiterOutput) RsaExpiresAt() time.TimePtrOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) time.TimePtrOutput { return v.RsaExpiresAt }).(time.TimePtrOutput)
}

// The timeout duration used for waiting.
func (o CertificateIssuanceWaiterOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateIssuanceWaiter) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

type CertificateIssuanceWaiterArrayOutput struct{ *pulumi.OutputState }

func (CertificateIssuanceWaiterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertificateIssuanceWaiter)(nil)).Elem()
}

func (o CertificateIssuanceWaiterArrayOutput) ToCertificateIssuanceWaiterArrayOutput() CertificateIssuanceWaiterArrayOutput {
	return o
}

func (o CertificateIssuanceWaiterArrayOutput) ToCertificateIssuanceWaiterArrayOutputWithContext(ctx context.Context) CertificateIssuanceWaiterArrayOutput {
	return o
}

func (o CertificateIssuanceWaiterArrayOutput) Index(i pulumi.IntInput) CertificateIssuanceWaiterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertificateIssuanceWaiter {
		return vs[0].([]*CertificateIssuanceWaiter)[vs[1].(int)]
	}).(CertificateIssuanceWaiterOutput)
}

type CertificateIssuanceWaiterMapOutput struct{ *pulumi.OutputState }

func (CertificateIssuanceWaiterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertificateIssuanceWaiter)(nil)).Elem()
}

func (o CertificateIssuanceWaiterMapOutput) ToCertificateIssuanceWaiterMapOutput() CertificateIssuanceWaiterMapOutput {
	return o
}

func (o CertificateIssuanceWaiterMapOutput) ToCertificateIssuanceWaiterMapOutputWithContext(ctx context.Context) CertificateIssuanceWaiterMapOutput {
	return o
}

func (o CertificateIssuanceWaiterMapOutput) MapIndex(k pulumi.StringInput) CertificateIssuanceWaiterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertificateIssuanceWaiter {
		return vs[0].(map[string]*CertificateIssuanceWaiter)[vs[1].(string)]
	}).(CertificateIssuanceWaiterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateIssuanceWaiterInput)(nil)).Elem(), &CertificateIssuanceWaiter{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateIssuanceWaiterArrayInput)(nil)).Elem(), CertificateIssuanceWaiterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateIssuanceWaiterMapInput)(nil)).Elem(), CertificateIssuanceWaiterMap{})
	pulumi.RegisterOutputType(CertificateIssuanceWaiterOutput{})
	pulumi.RegisterOutputType(CertificateIssuanceWaiterArrayOutput{})
	pulumi.RegisterOutputType(CertificateIssuanceWaiterMapOutput{})
}
