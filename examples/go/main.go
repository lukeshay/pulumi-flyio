package main

import (
	"fmt"

	"github.com/lukeshay/pulumi-flyio/sdk/go/flyio"
	flyioflyio "github.com/lukeshay/pulumi-flyio/sdk/go/flyio/flyio"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		appNameResource, err := flyio.NewRandom(ctx, "appName", &flyio.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		app, err := flyio.NewApp(ctx, "app", &flyio.AppArgs{
			Name: appNameResource.Result.ApplyT(func(result string) (string, error) {
				return fmt.Sprintf("pulumi-%v", result), nil
			}).(pulumi.StringOutput),
			Org:     pulumi.String("luke-shay"),
			Network: pulumi.String("pulumi-flyio"),
		})
		if err != nil {
			return err
		}
		machineSea1, err := flyio.NewMachine(ctx, "machine-sea-1", &flyio.MachineArgs{
			Name:               pulumi.String("machine-sea-1"),
			Region:             pulumi.String("sea"),
			App:                app.Name,
			DeploymentStrategy: pulumi.String("bluegreen"),
			Config: &flyio.FlyMachineConfigArgs{
				Image: pulumi.String("nginxdemos/hello:0.4"),
				Guest: &flyio.FlyMachineGuestArgs{
					Cpus:     pulumi.Int(1),
					CpuKind:  pulumi.String("shared"),
					MemoryMb: pulumi.Int(256),
				},
				Services: flyio.FlyMachineServiceArray{
					&flyio.FlyMachineServiceArgs{
						InternalPort: pulumi.Int(80),
						Ports: flyio.FlyMachinePortArray{
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(80),
								Handlers: pulumi.StringArray{
									pulumi.String("http"),
								},
								ForceHttps: pulumi.Bool(true),
							},
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(443),
								Handlers: pulumi.StringArray{
									pulumi.String("tls"),
									pulumi.String("http"),
								},
							},
						},
						Protocol: pulumi.String("tcp"),
						Checks: flyio.FlyMachineCheckArray{
							&flyio.FlyMachineCheckArgs{
								Interval:      pulumi.String("10s"),
								GracePeriod:   pulumi.String("5s"),
								Method:        pulumi.String("get"),
								Path:          pulumi.String("/"),
								Protocol:      pulumi.String("http"),
								Timeout:       pulumi.String("2s"),
								TlsSkipVerify: pulumi.Bool(true),
							},
						},
						Autostop:  pulumi.String("suspend"),
						Autostart: pulumi.Bool(true),
						Concurrency: &flyio.FlyMachineServiceConcurrencyArgs{
							SoftLimit: pulumi.Int(2000),
							HardLimit: pulumi.Int(2500),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewMachine(ctx, "machine-sea-2", &flyio.MachineArgs{
			Name:   pulumi.String("machine-sea-2"),
			Region: pulumi.String("sea"),
			App:    app.Name,
			Config: &flyio.FlyMachineConfigArgs{
				Image: pulumi.String("nginxdemos/hello:latest"),
				Guest: &flyio.FlyMachineGuestArgs{
					Cpus:     pulumi.Int(1),
					CpuKind:  pulumi.String("shared"),
					MemoryMb: pulumi.Int(256),
				},
				Services: flyio.FlyMachineServiceArray{
					&flyio.FlyMachineServiceArgs{
						InternalPort: pulumi.Int(80),
						Ports: flyio.FlyMachinePortArray{
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(80),
								Handlers: pulumi.StringArray{
									pulumi.String("http"),
								},
								ForceHttps: pulumi.Bool(true),
							},
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(443),
								Handlers: pulumi.StringArray{
									pulumi.String("tls"),
									pulumi.String("http"),
								},
							},
						},
						Protocol: pulumi.String("tcp"),
						Checks: flyio.FlyMachineCheckArray{
							&flyio.FlyMachineCheckArgs{
								Interval:      pulumi.String("10s"),
								GracePeriod:   pulumi.String("5s"),
								Method:        pulumi.String("get"),
								Path:          pulumi.String("/"),
								Protocol:      pulumi.String("http"),
								Timeout:       pulumi.String("2s"),
								TlsSkipVerify: pulumi.Bool(true),
							},
						},
						Autostop:  pulumi.String("suspend"),
						Autostart: pulumi.Bool(true),
						Concurrency: &flyio.FlyMachineServiceConcurrencyArgs{
							SoftLimit: pulumi.Int(2000),
							HardLimit: pulumi.Int(2500),
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			machineSea1,
		}))
		if err != nil {
			return err
		}
		machineIad1, err := flyio.NewMachine(ctx, "machine-iad-1", &flyio.MachineArgs{
			Name:   pulumi.String("machine-iad-1"),
			Region: pulumi.String("iad"),
			App:    app.Name,
			Config: &flyio.FlyMachineConfigArgs{
				Image: pulumi.String("nginxdemos/hello:latest"),
				Guest: &flyio.FlyMachineGuestArgs{
					Cpus:     pulumi.Int(1),
					CpuKind:  pulumi.String("shared"),
					MemoryMb: pulumi.Int(512),
				},
				Services: flyio.FlyMachineServiceArray{
					&flyio.FlyMachineServiceArgs{
						InternalPort: pulumi.Int(80),
						Ports: flyio.FlyMachinePortArray{
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(80),
								Handlers: pulumi.StringArray{
									pulumi.String("http"),
								},
								ForceHttps: pulumi.Bool(true),
							},
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(443),
								Handlers: pulumi.StringArray{
									pulumi.String("tls"),
									pulumi.String("http"),
								},
							},
						},
						Protocol: pulumi.String("tcp"),
						Checks: flyio.FlyMachineCheckArray{
							&flyio.FlyMachineCheckArgs{
								Interval:      pulumi.String("10s"),
								GracePeriod:   pulumi.String("5s"),
								Method:        pulumi.String("get"),
								Path:          pulumi.String("/"),
								Protocol:      pulumi.String("http"),
								Timeout:       pulumi.String("2s"),
								TlsSkipVerify: pulumi.Bool(true),
							},
						},
						Autostop:  pulumi.String("suspend"),
						Autostart: pulumi.Bool(true),
						Concurrency: &flyio.FlyMachineServiceConcurrencyArgs{
							SoftLimit: pulumi.Int(2000),
							HardLimit: pulumi.Int(2500),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewMachine(ctx, "machine-iad-2", &flyio.MachineArgs{
			Name:   pulumi.String("machine-iad-2"),
			Region: pulumi.String("iad"),
			App:    app.Name,
			Config: &flyio.FlyMachineConfigArgs{
				Image: pulumi.String("nginxdemos/hello:latest"),
				Guest: &flyio.FlyMachineGuestArgs{
					Cpus:     pulumi.Int(1),
					CpuKind:  pulumi.String("shared"),
					MemoryMb: pulumi.Int(256),
				},
				Services: flyio.FlyMachineServiceArray{
					&flyio.FlyMachineServiceArgs{
						InternalPort: pulumi.Int(80),
						Ports: flyio.FlyMachinePortArray{
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(80),
								Handlers: pulumi.StringArray{
									pulumi.String("http"),
								},
								ForceHttps: pulumi.Bool(true),
							},
							&flyio.FlyMachinePortArgs{
								Port: pulumi.Int(443),
								Handlers: pulumi.StringArray{
									pulumi.String("tls"),
									pulumi.String("http"),
								},
							},
						},
						Protocol: pulumi.String("tcp"),
						Checks: flyio.FlyMachineCheckArray{
							&flyio.FlyMachineCheckArgs{
								Interval:      pulumi.String("10s"),
								GracePeriod:   pulumi.String("5s"),
								Method:        pulumi.String("get"),
								Path:          pulumi.String("/"),
								Protocol:      pulumi.String("http"),
								Timeout:       pulumi.String("2s"),
								TlsSkipVerify: pulumi.Bool(true),
							},
						},
						Autostop:  pulumi.String("suspend"),
						Autostart: pulumi.Bool(true),
						Concurrency: &flyio.FlyMachineServiceConcurrencyArgs{
							SoftLimit: pulumi.Int(2000),
							HardLimit: pulumi.Int(2500),
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			machineIad1,
		}))
		if err != nil {
			return err
		}
		_, err = flyio.NewVolume(ctx, "volume-iad", &flyio.VolumeArgs{
			Name:              pulumi.String("volume_iad"),
			AutoBackupEnabled: pulumi.Bool(true),
			Region:            pulumi.String("iad"),
			App:               app.Name,
			SizeGb:            pulumi.Int(5),
		}, pulumi.DependsOn([]pulumi.Resource{
			machineIad1,
		}))
		if err != nil {
			return err
		}
		_, err = flyio.NewVolume(ctx, "volume-sea", &flyio.VolumeArgs{
			Name:   pulumi.String("volume_sea"),
			Region: pulumi.String("sea"),
			App:    app.Name,
			SizeGb: pulumi.Int(5),
		}, pulumi.DependsOn([]pulumi.Resource{
			machineSea1,
		}))
		if err != nil {
			return err
		}
		_, err = flyio.NewIP(ctx, "ipv4", &flyio.IPArgs{
			Region:   pulumi.String("sea"),
			App:      app.Name,
			AddrType: pulumi.String("v4"),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewIP(ctx, "ipv6", &flyio.IPArgs{
			Region:   pulumi.String("sea"),
			App:      app.Name,
			AddrType: pulumi.String("v6"),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewIP(ctx, "privateipv6", &flyio.IPArgs{
			Region:   pulumi.String("sea"),
			App:      app.Name,
			AddrType: pulumi.String("private_v6"),
			Network:  pulumi.String("pulumi-flyio"),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewIP(ctx, "sharedipv4", &flyio.IPArgs{
			App:      app.Name,
			AddrType: pulumi.String("shared_v4"),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewCertificate(ctx, "certificate", &flyio.CertificateArgs{
			App:      app.Name,
			Hostname: pulumi.String("pulumi-flyio.lshay.land"),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewWireGuardPeer(ctx, "wireguardPeer", &flyio.WireGuardPeerArgs{
			Org: pulumi.String("luke-shay"),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewWireGuardToken(ctx, "wireguardToken", &flyio.WireGuardTokenArgs{
			Org: pulumi.String("luke-shay"),
		})
		if err != nil {
			return err
		}
		ctx.Export("appName", pulumi.StringMap{
			"value": app.Name,
		})
		return nil
	})
}
