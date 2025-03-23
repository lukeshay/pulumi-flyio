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

// A Fly.io IP address allocation for an application.
type IP struct {
	pulumi.CustomResourceState

	// The allocated IP address.
	Address pulumi.StringOutput `pulumi:"address"`
	// The application the IP address is allocated for.
	App pulumi.StringOutput `pulumi:"app"`
	// When the IP address was allocated.
	CreatedAt time.TimeOutput `pulumi:"createdAt"`
	// The Fly.io IP address ID.
	FlyId pulumi.StringPtrOutput `pulumi:"flyId"`
	// The input arguments used to allocate the IP address.
	Input IPArgsTypeOutput `pulumi:"input"`
	// The network the IP address belongs to.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The region the IP address is allocated in.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The type of IP address (v4 or v6).
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIP registers a new resource with the given unique name, arguments, and options.
func NewIP(ctx *pulumi.Context,
	name string, args *IPArgs, opts ...pulumi.ResourceOption) (*IP, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddrType == nil {
		return nil, errors.New("invalid value for required argument 'AddrType'")
	}
	if args.App == nil {
		return nil, errors.New("invalid value for required argument 'App'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IP
	err := ctx.RegisterResource("flyio:index:IP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIP gets an existing IP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IPState, opts ...pulumi.ResourceOption) (*IP, error) {
	var resource IP
	err := ctx.ReadResource("flyio:index:IP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IP resources.
type ipState struct {
}

type IPState struct {
}

func (IPState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipState)(nil)).Elem()
}

type ipArgs struct {
	// The type of IP address (v4, v6, shared_v4, or private_v6).
	AddrType string `pulumi:"addrType"`
	// The name of the Fly.io application to allocate the IP address for.
	App string `pulumi:"app"`
	// The network to allocate the IP address in.
	Network *string `pulumi:"network"`
	// The region to allocate the IP address in. This is required for non-shared IP addresses.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a IP resource.
type IPArgs struct {
	// The type of IP address (v4, v6, shared_v4, or private_v6).
	AddrType pulumi.StringInput
	// The name of the Fly.io application to allocate the IP address for.
	App pulumi.StringInput
	// The network to allocate the IP address in.
	Network pulumi.StringPtrInput
	// The region to allocate the IP address in. This is required for non-shared IP addresses.
	Region pulumi.StringPtrInput
}

func (IPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipArgs)(nil)).Elem()
}

type IPInput interface {
	pulumi.Input

	ToIPOutput() IPOutput
	ToIPOutputWithContext(ctx context.Context) IPOutput
}

func (*IP) ElementType() reflect.Type {
	return reflect.TypeOf((**IP)(nil)).Elem()
}

func (i *IP) ToIPOutput() IPOutput {
	return i.ToIPOutputWithContext(context.Background())
}

func (i *IP) ToIPOutputWithContext(ctx context.Context) IPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPOutput)
}

// IPArrayInput is an input type that accepts IPArray and IPArrayOutput values.
// You can construct a concrete instance of `IPArrayInput` via:
//
//	IPArray{ IPArgs{...} }
type IPArrayInput interface {
	pulumi.Input

	ToIPArrayOutput() IPArrayOutput
	ToIPArrayOutputWithContext(context.Context) IPArrayOutput
}

type IPArray []IPInput

func (IPArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IP)(nil)).Elem()
}

func (i IPArray) ToIPArrayOutput() IPArrayOutput {
	return i.ToIPArrayOutputWithContext(context.Background())
}

func (i IPArray) ToIPArrayOutputWithContext(ctx context.Context) IPArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPArrayOutput)
}

// IPMapInput is an input type that accepts IPMap and IPMapOutput values.
// You can construct a concrete instance of `IPMapInput` via:
//
//	IPMap{ "key": IPArgs{...} }
type IPMapInput interface {
	pulumi.Input

	ToIPMapOutput() IPMapOutput
	ToIPMapOutputWithContext(context.Context) IPMapOutput
}

type IPMap map[string]IPInput

func (IPMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IP)(nil)).Elem()
}

func (i IPMap) ToIPMapOutput() IPMapOutput {
	return i.ToIPMapOutputWithContext(context.Background())
}

func (i IPMap) ToIPMapOutputWithContext(ctx context.Context) IPMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPMapOutput)
}

type IPOutput struct{ *pulumi.OutputState }

func (IPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IP)(nil)).Elem()
}

func (o IPOutput) ToIPOutput() IPOutput {
	return o
}

func (o IPOutput) ToIPOutputWithContext(ctx context.Context) IPOutput {
	return o
}

// The allocated IP address.
func (o IPOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The application the IP address is allocated for.
func (o IPOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.App }).(pulumi.StringOutput)
}

// When the IP address was allocated.
func (o IPOutput) CreatedAt() time.TimeOutput {
	return o.ApplyT(func(v *IP) time.TimeOutput { return v.CreatedAt }).(time.TimeOutput)
}

// The Fly.io IP address ID.
func (o IPOutput) FlyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IP) pulumi.StringPtrOutput { return v.FlyId }).(pulumi.StringPtrOutput)
}

// The input arguments used to allocate the IP address.
func (o IPOutput) Input() IPArgsTypeOutput {
	return o.ApplyT(func(v *IP) IPArgsTypeOutput { return v.Input }).(IPArgsTypeOutput)
}

// The network the IP address belongs to.
func (o IPOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IP) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// The region the IP address is allocated in.
func (o IPOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IP) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The type of IP address (v4 or v6).
func (o IPOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IP) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type IPArrayOutput struct{ *pulumi.OutputState }

func (IPArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IP)(nil)).Elem()
}

func (o IPArrayOutput) ToIPArrayOutput() IPArrayOutput {
	return o
}

func (o IPArrayOutput) ToIPArrayOutputWithContext(ctx context.Context) IPArrayOutput {
	return o
}

func (o IPArrayOutput) Index(i pulumi.IntInput) IPOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IP {
		return vs[0].([]*IP)[vs[1].(int)]
	}).(IPOutput)
}

type IPMapOutput struct{ *pulumi.OutputState }

func (IPMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IP)(nil)).Elem()
}

func (o IPMapOutput) ToIPMapOutput() IPMapOutput {
	return o
}

func (o IPMapOutput) ToIPMapOutputWithContext(ctx context.Context) IPMapOutput {
	return o
}

func (o IPMapOutput) MapIndex(k pulumi.StringInput) IPOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IP {
		return vs[0].(map[string]*IP)[vs[1].(string)]
	}).(IPOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IPInput)(nil)).Elem(), &IP{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPArrayInput)(nil)).Elem(), IPArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IPMapInput)(nil)).Elem(), IPMap{})
	pulumi.RegisterOutputType(IPOutput{})
	pulumi.RegisterOutputType(IPArrayOutput{})
	pulumi.RegisterOutputType(IPMapOutput{})
}
