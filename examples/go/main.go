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
			AppName: appNameResource.Result.ApplyT(func(result string) (string, error) {
				return fmt.Sprintf("pulumi-%v", result), nil
			}).(pulumi.StringOutput),
			OrgSlug: pulumi.String("luke-shay"),
		})
		if err != nil {
			return err
		}
		machineSea1, err := flyio.NewMachine(ctx, "machine-sea-1", &flyio.MachineArgs{
			Name:           pulumi.String("machine-sea-1"),
			UpdateStrategy: pulumi.String("bluegreen"),
			Region:         pulumi.String("sea"),
			AppName:        app.Name,
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
						Autostop:  pulumi.Bool(true),
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
			Name:    pulumi.String("machine-sea-2"),
			Region:  pulumi.String("sea"),
			AppName: app.Name,
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
						Autostop:  pulumi.Bool(true),
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
			Name:           pulumi.String("machine-iad-1"),
			UpdateStrategy: pulumi.String("bluegreen"),
			Region:         pulumi.String("iad"),
			AppName:        app.Name,
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
						Autostop:  pulumi.Bool(true),
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
			Name:    pulumi.String("machine-iad-2"),
			Region:  pulumi.String("iad"),
			AppName: app.Name,
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
						Autostop:  pulumi.Bool(true),
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
		_, err = flyio.NewVolume(ctx, "volume-sea", &flyio.VolumeArgs{
			Name:    pulumi.String("volume_sea"),
			Region:  pulumi.String("sea"),
			AppName: app.Name,
			SizeGb:  pulumi.Int(5),
		})
		if err != nil {
			return err
		}
		_, err = flyio.NewVolume(ctx, "volume-iad", &flyio.VolumeArgs{
			Name:    pulumi.String("volume_iad"),
			Region:  pulumi.String("iad"),
			AppName: app.Name,
			SizeGb:  pulumi.Int(5),
		})
		if err != nil {
			return err
		}
		ctx.Export("appName", map[string]interface{}{
			"value": app.Name,
		})
		return nil
	})
}
