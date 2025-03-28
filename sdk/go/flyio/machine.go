// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flyio

import (
	"context"
	"reflect"

	"errors"
	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/flyio"
	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Fly.io machine represents a VM instance that runs your application.
type Machine struct {
	pulumi.CustomResourceState

	// The Fly.io application the machine belongs to.
	App       pulumi.StringOutput          `pulumi:"app"`
	Checks    flyio.CheckStatusArrayOutput `pulumi:"checks"`
	Config    flyio.FlyMachineConfigOutput `pulumi:"config"`
	CreatedAt pulumi.StringPtrOutput       `pulumi:"createdAt"`
	// The deployment strategy used for the machine.
	DeploymentStrategy pulumi.StringPtrOutput          `pulumi:"deploymentStrategy"`
	Events             flyio.MachineEventArrayOutput   `pulumi:"events"`
	FlyId              pulumi.StringOutput             `pulumi:"flyId"`
	HostStatus         pulumi.StringPtrOutput          `pulumi:"hostStatus"`
	ImageRef           flyio.ImageRefPtrOutput         `pulumi:"imageRef"`
	IncompleteConfig   flyio.FlyMachineConfigPtrOutput `pulumi:"incompleteConfig"`
	// The input arguments used to create the machine.
	Input      MachineArgsTypeOutput  `pulumi:"input"`
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Nonce      pulumi.StringPtrOutput `pulumi:"nonce"`
	PrivateIp  pulumi.StringPtrOutput `pulumi:"privateIp"`
	Region     pulumi.StringPtrOutput `pulumi:"region"`
	State      pulumi.StringOutput    `pulumi:"state"`
	UpdatedAt  pulumi.StringPtrOutput `pulumi:"updatedAt"`
}

// NewMachine registers a new resource with the given unique name, arguments, and options.
func NewMachine(ctx *pulumi.Context,
	name string, args *MachineArgs, opts ...pulumi.ResourceOption) (*Machine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.App == nil {
		return nil, errors.New("invalid value for required argument 'App'")
	}
	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Machine
	err := ctx.RegisterResource("flyio:index:Machine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachine gets an existing Machine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineState, opts ...pulumi.ResourceOption) (*Machine, error) {
	var resource Machine
	err := ctx.ReadResource("flyio:index:Machine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Machine resources.
type machineState struct {
}

type MachineState struct {
}

func (MachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineState)(nil)).Elem()
}

type machineArgs struct {
	// The Fly.io application to deploy the machine to.
	App    string                 `pulumi:"app"`
	Config flyio.FlyMachineConfig `pulumi:"config"`
	// The deployment strategy for the machine.
	DeploymentStrategy      *string `pulumi:"deploymentStrategy"`
	LeaseTtl                *int    `pulumi:"leaseTtl"`
	Lsvd                    *bool   `pulumi:"lsvd"`
	Name                    *string `pulumi:"name"`
	Region                  *string `pulumi:"region"`
	SkipLaunch              *bool   `pulumi:"skipLaunch"`
	SkipServiceRegistration *bool   `pulumi:"skipServiceRegistration"`
}

// The set of arguments for constructing a Machine resource.
type MachineArgs struct {
	// The Fly.io application to deploy the machine to.
	App    pulumi.StringInput
	Config flyio.FlyMachineConfigInput
	// The deployment strategy for the machine.
	DeploymentStrategy      pulumi.StringPtrInput
	LeaseTtl                pulumi.IntPtrInput
	Lsvd                    pulumi.BoolPtrInput
	Name                    pulumi.StringPtrInput
	Region                  pulumi.StringPtrInput
	SkipLaunch              pulumi.BoolPtrInput
	SkipServiceRegistration pulumi.BoolPtrInput
}

func (MachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineArgs)(nil)).Elem()
}

type MachineInput interface {
	pulumi.Input

	ToMachineOutput() MachineOutput
	ToMachineOutputWithContext(ctx context.Context) MachineOutput
}

func (*Machine) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (i *Machine) ToMachineOutput() MachineOutput {
	return i.ToMachineOutputWithContext(context.Background())
}

func (i *Machine) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineOutput)
}

// MachineArrayInput is an input type that accepts MachineArray and MachineArrayOutput values.
// You can construct a concrete instance of `MachineArrayInput` via:
//
//	MachineArray{ MachineArgs{...} }
type MachineArrayInput interface {
	pulumi.Input

	ToMachineArrayOutput() MachineArrayOutput
	ToMachineArrayOutputWithContext(context.Context) MachineArrayOutput
}

type MachineArray []MachineInput

func (MachineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Machine)(nil)).Elem()
}

func (i MachineArray) ToMachineArrayOutput() MachineArrayOutput {
	return i.ToMachineArrayOutputWithContext(context.Background())
}

func (i MachineArray) ToMachineArrayOutputWithContext(ctx context.Context) MachineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineArrayOutput)
}

// MachineMapInput is an input type that accepts MachineMap and MachineMapOutput values.
// You can construct a concrete instance of `MachineMapInput` via:
//
//	MachineMap{ "key": MachineArgs{...} }
type MachineMapInput interface {
	pulumi.Input

	ToMachineMapOutput() MachineMapOutput
	ToMachineMapOutputWithContext(context.Context) MachineMapOutput
}

type MachineMap map[string]MachineInput

func (MachineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Machine)(nil)).Elem()
}

func (i MachineMap) ToMachineMapOutput() MachineMapOutput {
	return i.ToMachineMapOutputWithContext(context.Background())
}

func (i MachineMap) ToMachineMapOutputWithContext(ctx context.Context) MachineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineMapOutput)
}

type MachineOutput struct{ *pulumi.OutputState }

func (MachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Machine)(nil)).Elem()
}

func (o MachineOutput) ToMachineOutput() MachineOutput {
	return o
}

func (o MachineOutput) ToMachineOutputWithContext(ctx context.Context) MachineOutput {
	return o
}

// The Fly.io application the machine belongs to.
func (o MachineOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.App }).(pulumi.StringOutput)
}

func (o MachineOutput) Checks() flyio.CheckStatusArrayOutput {
	return o.ApplyT(func(v *Machine) flyio.CheckStatusArrayOutput { return v.Checks }).(flyio.CheckStatusArrayOutput)
}

func (o MachineOutput) Config() flyio.FlyMachineConfigOutput {
	return o.ApplyT(func(v *Machine) flyio.FlyMachineConfigOutput { return v.Config }).(flyio.FlyMachineConfigOutput)
}

func (o MachineOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The deployment strategy used for the machine.
func (o MachineOutput) DeploymentStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.DeploymentStrategy }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) Events() flyio.MachineEventArrayOutput {
	return o.ApplyT(func(v *Machine) flyio.MachineEventArrayOutput { return v.Events }).(flyio.MachineEventArrayOutput)
}

func (o MachineOutput) FlyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.FlyId }).(pulumi.StringOutput)
}

func (o MachineOutput) HostStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.HostStatus }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) ImageRef() flyio.ImageRefPtrOutput {
	return o.ApplyT(func(v *Machine) flyio.ImageRefPtrOutput { return v.ImageRef }).(flyio.ImageRefPtrOutput)
}

func (o MachineOutput) IncompleteConfig() flyio.FlyMachineConfigPtrOutput {
	return o.ApplyT(func(v *Machine) flyio.FlyMachineConfigPtrOutput { return v.IncompleteConfig }).(flyio.FlyMachineConfigPtrOutput)
}

// The input arguments used to create the machine.
func (o MachineOutput) Input() MachineArgsTypeOutput {
	return o.ApplyT(func(v *Machine) MachineArgsTypeOutput { return v.Input }).(MachineArgsTypeOutput)
}

func (o MachineOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineOutput) Nonce() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.Nonce }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) PrivateIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.PrivateIp }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o MachineOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o MachineOutput) UpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Machine) pulumi.StringPtrOutput { return v.UpdatedAt }).(pulumi.StringPtrOutput)
}

type MachineArrayOutput struct{ *pulumi.OutputState }

func (MachineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Machine)(nil)).Elem()
}

func (o MachineArrayOutput) ToMachineArrayOutput() MachineArrayOutput {
	return o
}

func (o MachineArrayOutput) ToMachineArrayOutputWithContext(ctx context.Context) MachineArrayOutput {
	return o
}

func (o MachineArrayOutput) Index(i pulumi.IntInput) MachineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Machine {
		return vs[0].([]*Machine)[vs[1].(int)]
	}).(MachineOutput)
}

type MachineMapOutput struct{ *pulumi.OutputState }

func (MachineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Machine)(nil)).Elem()
}

func (o MachineMapOutput) ToMachineMapOutput() MachineMapOutput {
	return o
}

func (o MachineMapOutput) ToMachineMapOutputWithContext(ctx context.Context) MachineMapOutput {
	return o
}

func (o MachineMapOutput) MapIndex(k pulumi.StringInput) MachineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Machine {
		return vs[0].(map[string]*Machine)[vs[1].(string)]
	}).(MachineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MachineInput)(nil)).Elem(), &Machine{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineArrayInput)(nil)).Elem(), MachineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MachineMapInput)(nil)).Elem(), MachineMap{})
	pulumi.RegisterOutputType(MachineOutput{})
	pulumi.RegisterOutputType(MachineArrayOutput{})
	pulumi.RegisterOutputType(MachineMapOutput{})
}
