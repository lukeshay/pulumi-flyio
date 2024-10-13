// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flyio

import (
	"context"
	"reflect"

	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/flyio"
	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type AppArgsType struct {
	EnableSubdomains *bool   `pulumi:"enableSubdomains"`
	Name             string  `pulumi:"name"`
	Network          *string `pulumi:"network"`
	Org              string  `pulumi:"org"`
}

type AppArgsTypeOutput struct{ *pulumi.OutputState }

func (AppArgsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppArgsType)(nil)).Elem()
}

func (o AppArgsTypeOutput) ToAppArgsTypeOutput() AppArgsTypeOutput {
	return o
}

func (o AppArgsTypeOutput) ToAppArgsTypeOutputWithContext(ctx context.Context) AppArgsTypeOutput {
	return o
}

func (o AppArgsTypeOutput) EnableSubdomains() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AppArgsType) *bool { return v.EnableSubdomains }).(pulumi.BoolPtrOutput)
}

func (o AppArgsTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppArgsType) string { return v.Name }).(pulumi.StringOutput)
}

func (o AppArgsTypeOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppArgsType) *string { return v.Network }).(pulumi.StringPtrOutput)
}

func (o AppArgsTypeOutput) Org() pulumi.StringOutput {
	return o.ApplyT(func(v AppArgsType) string { return v.Org }).(pulumi.StringOutput)
}

type CertificateArgsType struct {
	App      string `pulumi:"app"`
	Hostname string `pulumi:"hostname"`
}

type CertificateArgsTypeOutput struct{ *pulumi.OutputState }

func (CertificateArgsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateArgsType)(nil)).Elem()
}

func (o CertificateArgsTypeOutput) ToCertificateArgsTypeOutput() CertificateArgsTypeOutput {
	return o
}

func (o CertificateArgsTypeOutput) ToCertificateArgsTypeOutputWithContext(ctx context.Context) CertificateArgsTypeOutput {
	return o
}

func (o CertificateArgsTypeOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateArgsType) string { return v.App }).(pulumi.StringOutput)
}

func (o CertificateArgsTypeOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateArgsType) string { return v.Hostname }).(pulumi.StringOutput)
}

type CertificateStateChecks struct {
	ARecords              []string `pulumi:"aRecords"`
	AaaaRecords           []string `pulumi:"aaaaRecords"`
	CnameRecords          []string `pulumi:"cnameRecords"`
	DnsProvider           string   `pulumi:"dnsProvider"`
	DnsVerificationRecord string   `pulumi:"dnsVerificationRecord"`
	ResolvedAddresses     []string `pulumi:"resolvedAddresses"`
	Soa                   string   `pulumi:"soa"`
}

type CertificateStateChecksOutput struct{ *pulumi.OutputState }

func (CertificateStateChecksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateStateChecks)(nil)).Elem()
}

func (o CertificateStateChecksOutput) ToCertificateStateChecksOutput() CertificateStateChecksOutput {
	return o
}

func (o CertificateStateChecksOutput) ToCertificateStateChecksOutputWithContext(ctx context.Context) CertificateStateChecksOutput {
	return o
}

func (o CertificateStateChecksOutput) ARecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificateStateChecks) []string { return v.ARecords }).(pulumi.StringArrayOutput)
}

func (o CertificateStateChecksOutput) AaaaRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificateStateChecks) []string { return v.AaaaRecords }).(pulumi.StringArrayOutput)
}

func (o CertificateStateChecksOutput) CnameRecords() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificateStateChecks) []string { return v.CnameRecords }).(pulumi.StringArrayOutput)
}

func (o CertificateStateChecksOutput) DnsProvider() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateStateChecks) string { return v.DnsProvider }).(pulumi.StringOutput)
}

func (o CertificateStateChecksOutput) DnsVerificationRecord() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateStateChecks) string { return v.DnsVerificationRecord }).(pulumi.StringOutput)
}

func (o CertificateStateChecksOutput) ResolvedAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CertificateStateChecks) []string { return v.ResolvedAddresses }).(pulumi.StringArrayOutput)
}

func (o CertificateStateChecksOutput) Soa() pulumi.StringOutput {
	return o.ApplyT(func(v CertificateStateChecks) string { return v.Soa }).(pulumi.StringOutput)
}

type IPArgsType struct {
	AddrType string  `pulumi:"addrType"`
	App      string  `pulumi:"app"`
	Network  *string `pulumi:"network"`
	Region   string  `pulumi:"region"`
}

type IPArgsTypeOutput struct{ *pulumi.OutputState }

func (IPArgsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPArgsType)(nil)).Elem()
}

func (o IPArgsTypeOutput) ToIPArgsTypeOutput() IPArgsTypeOutput {
	return o
}

func (o IPArgsTypeOutput) ToIPArgsTypeOutputWithContext(ctx context.Context) IPArgsTypeOutput {
	return o
}

func (o IPArgsTypeOutput) AddrType() pulumi.StringOutput {
	return o.ApplyT(func(v IPArgsType) string { return v.AddrType }).(pulumi.StringOutput)
}

func (o IPArgsTypeOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v IPArgsType) string { return v.App }).(pulumi.StringOutput)
}

func (o IPArgsTypeOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPArgsType) *string { return v.Network }).(pulumi.StringPtrOutput)
}

func (o IPArgsTypeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v IPArgsType) string { return v.Region }).(pulumi.StringOutput)
}

type MachineArgsType struct {
	App                     string                 `pulumi:"app"`
	Config                  flyio.FlyMachineConfig `pulumi:"config"`
	DeploymentStrategy      *string                `pulumi:"deploymentStrategy"`
	LeaseTtl                *int                   `pulumi:"leaseTtl"`
	Lsvd                    *bool                  `pulumi:"lsvd"`
	Name                    *string                `pulumi:"name"`
	Region                  *string                `pulumi:"region"`
	SkipLaunch              *bool                  `pulumi:"skipLaunch"`
	SkipServiceRegistration *bool                  `pulumi:"skipServiceRegistration"`
}

type MachineArgsTypeOutput struct{ *pulumi.OutputState }

func (MachineArgsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineArgsType)(nil)).Elem()
}

func (o MachineArgsTypeOutput) ToMachineArgsTypeOutput() MachineArgsTypeOutput {
	return o
}

func (o MachineArgsTypeOutput) ToMachineArgsTypeOutputWithContext(ctx context.Context) MachineArgsTypeOutput {
	return o
}

func (o MachineArgsTypeOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v MachineArgsType) string { return v.App }).(pulumi.StringOutput)
}

func (o MachineArgsTypeOutput) Config() flyio.FlyMachineConfigOutput {
	return o.ApplyT(func(v MachineArgsType) flyio.FlyMachineConfig { return v.Config }).(flyio.FlyMachineConfigOutput)
}

func (o MachineArgsTypeOutput) DeploymentStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *string { return v.DeploymentStrategy }).(pulumi.StringPtrOutput)
}

func (o MachineArgsTypeOutput) LeaseTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *int { return v.LeaseTtl }).(pulumi.IntPtrOutput)
}

func (o MachineArgsTypeOutput) Lsvd() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *bool { return v.Lsvd }).(pulumi.BoolPtrOutput)
}

func (o MachineArgsTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MachineArgsTypeOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o MachineArgsTypeOutput) SkipLaunch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *bool { return v.SkipLaunch }).(pulumi.BoolPtrOutput)
}

func (o MachineArgsTypeOutput) SkipServiceRegistration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MachineArgsType) *bool { return v.SkipServiceRegistration }).(pulumi.BoolPtrOutput)
}

type VolumeArgsType struct {
	App               string                 `pulumi:"app"`
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
}

type VolumeArgsTypeOutput struct{ *pulumi.OutputState }

func (VolumeArgsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeArgsType)(nil)).Elem()
}

func (o VolumeArgsTypeOutput) ToVolumeArgsTypeOutput() VolumeArgsTypeOutput {
	return o
}

func (o VolumeArgsTypeOutput) ToVolumeArgsTypeOutputWithContext(ctx context.Context) VolumeArgsTypeOutput {
	return o
}

func (o VolumeArgsTypeOutput) App() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeArgsType) string { return v.App }).(pulumi.StringOutput)
}

func (o VolumeArgsTypeOutput) AutoBackupEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *bool { return v.AutoBackupEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeArgsTypeOutput) Compute() flyio.FlyMachineGuestPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *flyio.FlyMachineGuest { return v.Compute }).(flyio.FlyMachineGuestPtrOutput)
}

func (o VolumeArgsTypeOutput) ComputeImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *string { return v.ComputeImage }).(pulumi.StringPtrOutput)
}

func (o VolumeArgsTypeOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *bool { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

func (o VolumeArgsTypeOutput) Fstype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *string { return v.Fstype }).(pulumi.StringPtrOutput)
}

func (o VolumeArgsTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VolumeArgsTypeOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o VolumeArgsTypeOutput) RequireUniqueZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *bool { return v.RequireUniqueZone }).(pulumi.BoolPtrOutput)
}

func (o VolumeArgsTypeOutput) SizeGb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *int { return v.SizeGb }).(pulumi.IntPtrOutput)
}

func (o VolumeArgsTypeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o VolumeArgsTypeOutput) SnapshotRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *int { return v.SnapshotRetention }).(pulumi.IntPtrOutput)
}

func (o VolumeArgsTypeOutput) SourceVolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeArgsType) *string { return v.SourceVolumeId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppArgsTypeOutput{})
	pulumi.RegisterOutputType(CertificateArgsTypeOutput{})
	pulumi.RegisterOutputType(CertificateStateChecksOutput{})
	pulumi.RegisterOutputType(IPArgsTypeOutput{})
	pulumi.RegisterOutputType(MachineArgsTypeOutput{})
	pulumi.RegisterOutputType(VolumeArgsTypeOutput{})
}
