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

// A Fly.io volume provides persistent storage for your applications.
type Volume struct {
	pulumi.CustomResourceState

	// The Fly.io application the volume is attached to.
	App               pulumi.StringOutput    `pulumi:"app"`
	AttachedAllocId   pulumi.StringPtrOutput `pulumi:"attachedAllocId"`
	AttachedMachineId pulumi.StringPtrOutput `pulumi:"attachedMachineId"`
	AutoBackupEnabled pulumi.BoolPtrOutput   `pulumi:"autoBackupEnabled"`
	BlockSize         pulumi.IntPtrOutput    `pulumi:"blockSize"`
	Blocks            pulumi.IntPtrOutput    `pulumi:"blocks"`
	BlocksAvail       pulumi.IntPtrOutput    `pulumi:"blocksAvail"`
	BlocksFree        pulumi.IntPtrOutput    `pulumi:"blocksFree"`
	CreatedAt         pulumi.StringPtrOutput `pulumi:"createdAt"`
	Encrypted         pulumi.BoolPtrOutput   `pulumi:"encrypted"`
	FlyId             pulumi.StringOutput    `pulumi:"flyId"`
	Fstype            pulumi.StringPtrOutput `pulumi:"fstype"`
	HostStatus        pulumi.StringPtrOutput `pulumi:"hostStatus"`
	// The input arguments used to create the volume.
	Input             VolumeArgsTypeOutput   `pulumi:"input"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	Region            pulumi.StringPtrOutput `pulumi:"region"`
	SizeGb            pulumi.IntPtrOutput    `pulumi:"sizeGb"`
	SnapshotRetention pulumi.IntPtrOutput    `pulumi:"snapshotRetention"`
	State             pulumi.StringPtrOutput `pulumi:"state"`
	Zone              pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.App == nil {
		return nil, errors.New("invalid value for required argument 'App'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("flyio:index:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("flyio:index:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// The Fly.io application to attach the volume to.
	App string `pulumi:"app"`
	// Whether to enable automatic backups for the volume.
	AutoBackupEnabled *bool                  `pulumi:"autoBackupEnabled"`
	Compute           *flyio.FlyMachineGuest `pulumi:"compute"`
	ComputeImage      *string                `pulumi:"computeImage"`
	Encrypted         *bool                  `pulumi:"encrypted"`
	Fstype            *string                `pulumi:"fstype"`
	Name              *string                `pulumi:"name"`
	Region            *string                `pulumi:"region"`
	RequireUniqueZone *bool                  `pulumi:"requireUniqueZone"`
	SizeGb            *int                   `pulumi:"sizeGb"`
	SnapshotId        *string                `pulumi:"snapshotId"`
	SnapshotRetention *int                   `pulumi:"snapshotRetention"`
	SourceVolumeId    *string                `pulumi:"sourceVolumeId"`
	UniqueZoneAppWide *bool                  `pulumi:"uniqueZoneAppWide"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// The Fly.io application to attach the volume to.
	App pulumi.StringInput
	// Whether to enable automatic backups for the volume.
	AutoBackupEnabled pulumi.BoolPtrInput
	Compute           flyio.FlyMachineGuestPtrInput
	ComputeImage      pulumi.StringPtrInput
	Encrypted         pulumi.BoolPtrInput
	Fstype            pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Region            pulumi.StringPtrInput
	RequireUniqueZone pulumi.BoolPtrInput
	SizeGb            pulumi.IntPtrInput
	SnapshotId        pulumi.StringPtrInput
	SnapshotRetention pulumi.IntPtrInput
	SourceVolumeId    pulumi.StringPtrInput
	UniqueZoneAppWide pulumi.BoolPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

// VolumeArrayInput is an input type that accepts VolumeArray and VolumeArrayOutput values.
// You can construct a concrete instance of `VolumeArrayInput` via:
//
//	VolumeArray{ VolumeArgs{...} }
type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

// VolumeMapInput is an input type that accepts VolumeMap and VolumeMapOutput values.
// You can construct a concrete instance of `VolumeMapInput` via:
//
//	VolumeMap{ "key": VolumeArgs{...} }
type VolumeMapInput interface {
	pulumi.Input

	ToVolumeMapOutput() VolumeMapOutput
	ToVolumeMapOutputWithContext(context.Context) VolumeMapOutput
}

type VolumeMap map[string]VolumeInput

func (VolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (i VolumeMap) ToVolumeMapOutput() VolumeMapOutput {
	return i.ToVolumeMapOutputWithContext(context.Background())
}

func (i VolumeMap) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMapOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// The Fly.io application the volume is attached to.
func (o VolumeOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.App }).(pulumi.StringOutput)
}

func (o VolumeOutput) AttachedAllocId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.AttachedAllocId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) AttachedMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.AttachedMachineId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) AutoBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.AutoBackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) BlockSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.BlockSize }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) Blocks() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.Blocks }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) BlocksAvail() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.BlocksAvail }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) BlocksFree() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.BlocksFree }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) FlyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.FlyId }).(pulumi.StringOutput)
}

func (o VolumeOutput) Fstype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Fstype }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) HostStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.HostStatus }).(pulumi.StringPtrOutput)
}

// The input arguments used to create the volume.
func (o VolumeOutput) Input() VolumeArgsTypeOutput {
	return o.ApplyT(func(v *Volume) VolumeArgsTypeOutput { return v.Input }).(VolumeArgsTypeOutput)
}

func (o VolumeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) SizeGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.SizeGb }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) SnapshotRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.SnapshotRetention }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].([]*Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMapOutput struct{ *pulumi.OutputState }

func (VolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (o VolumeMapOutput) ToVolumeMapOutput() VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) MapIndex(k pulumi.StringInput) VolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].(map[string]*Volume)[vs[1].(string)]
	}).(VolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeInput)(nil)).Elem(), &Volume{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArrayInput)(nil)).Elem(), VolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMapInput)(nil)).Elem(), VolumeMap{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMapOutput{})
}
