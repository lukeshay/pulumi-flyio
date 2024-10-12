using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Flyio = Pulumi.Flyio;

return await Deployment.RunAsync(() => 
{
    var appNameResource = new Flyio.Random("appName", new()
    {
        Length = 24,
    });

    var app = new Flyio.App("app", new()
    {
        Name = appNameResource.Result.Apply(result => $"pulumi-{result}"),
        Org = "luke-shay",
    });

    var machineSea1 = new Flyio.Machine("machine-sea-1", new()
    {
        Name = "machine-sea-1",
        Region = "sea",
        App = app.Name,
        DeploymentStrategy = "bluegreen",
        Config = new Flyio.Flyio.Inputs.FlyMachineConfigArgs
        {
            Image = "nginxdemos/hello:0.4",
            Guest = new Flyio.Flyio.Inputs.FlyMachineGuestArgs
            {
                Cpus = 1,
                CpuKind = "shared",
                MemoryMb = 256,
            },
            Services = new[]
            {
                new Flyio.Flyio.Inputs.FlyMachineServiceArgs
                {
                    InternalPort = 80,
                    Ports = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 80,
                            Handlers = new[]
                            {
                                "http",
                            },
                            ForceHttps = true,
                        },
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 443,
                            Handlers = new[]
                            {
                                "tls",
                                "http",
                            },
                        },
                    },
                    Protocol = "tcp",
                    Checks = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachineCheckArgs
                        {
                            Interval = "10s",
                            GracePeriod = "5s",
                            Method = "get",
                            Path = "/",
                            Protocol = "http",
                            Timeout = "2s",
                            TlsSkipVerify = true,
                        },
                    },
                    Autostop = "suspend",
                    Autostart = true,
                    Concurrency = new Flyio.Flyio.Inputs.FlyMachineServiceConcurrencyArgs
                    {
                        SoftLimit = 2000,
                        HardLimit = 2500,
                    },
                },
            },
        },
    });

    var machineSea2 = new Flyio.Machine("machine-sea-2", new()
    {
        Name = "machine-sea-2",
        Region = "sea",
        App = app.Name,
        Config = new Flyio.Flyio.Inputs.FlyMachineConfigArgs
        {
            Image = "nginxdemos/hello:latest",
            Guest = new Flyio.Flyio.Inputs.FlyMachineGuestArgs
            {
                Cpus = 1,
                CpuKind = "shared",
                MemoryMb = 256,
            },
            Services = new[]
            {
                new Flyio.Flyio.Inputs.FlyMachineServiceArgs
                {
                    InternalPort = 80,
                    Ports = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 80,
                            Handlers = new[]
                            {
                                "http",
                            },
                            ForceHttps = true,
                        },
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 443,
                            Handlers = new[]
                            {
                                "tls",
                                "http",
                            },
                        },
                    },
                    Protocol = "tcp",
                    Checks = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachineCheckArgs
                        {
                            Interval = "10s",
                            GracePeriod = "5s",
                            Method = "get",
                            Path = "/",
                            Protocol = "http",
                            Timeout = "2s",
                            TlsSkipVerify = true,
                        },
                    },
                    Autostop = "suspend",
                    Autostart = true,
                    Concurrency = new Flyio.Flyio.Inputs.FlyMachineServiceConcurrencyArgs
                    {
                        SoftLimit = 2000,
                        HardLimit = 2500,
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            machineSea1,
        },
    });

    var machineIad1 = new Flyio.Machine("machine-iad-1", new()
    {
        Name = "machine-iad-1",
        Region = "iad",
        App = app.Name,
        Config = new Flyio.Flyio.Inputs.FlyMachineConfigArgs
        {
            Image = "nginxdemos/hello:latest",
            Guest = new Flyio.Flyio.Inputs.FlyMachineGuestArgs
            {
                Cpus = 1,
                CpuKind = "shared",
                MemoryMb = 512,
            },
            Services = new[]
            {
                new Flyio.Flyio.Inputs.FlyMachineServiceArgs
                {
                    InternalPort = 80,
                    Ports = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 80,
                            Handlers = new[]
                            {
                                "http",
                            },
                            ForceHttps = true,
                        },
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 443,
                            Handlers = new[]
                            {
                                "tls",
                                "http",
                            },
                        },
                    },
                    Protocol = "tcp",
                    Checks = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachineCheckArgs
                        {
                            Interval = "10s",
                            GracePeriod = "5s",
                            Method = "get",
                            Path = "/",
                            Protocol = "http",
                            Timeout = "2s",
                            TlsSkipVerify = true,
                        },
                    },
                    Autostop = "suspend",
                    Autostart = true,
                    Concurrency = new Flyio.Flyio.Inputs.FlyMachineServiceConcurrencyArgs
                    {
                        SoftLimit = 2000,
                        HardLimit = 2500,
                    },
                },
            },
        },
    });

    var machineIad2 = new Flyio.Machine("machine-iad-2", new()
    {
        Name = "machine-iad-2",
        Region = "iad",
        App = app.Name,
        Config = new Flyio.Flyio.Inputs.FlyMachineConfigArgs
        {
            Image = "nginxdemos/hello:latest",
            Guest = new Flyio.Flyio.Inputs.FlyMachineGuestArgs
            {
                Cpus = 1,
                CpuKind = "shared",
                MemoryMb = 256,
            },
            Services = new[]
            {
                new Flyio.Flyio.Inputs.FlyMachineServiceArgs
                {
                    InternalPort = 80,
                    Ports = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 80,
                            Handlers = new[]
                            {
                                "http",
                            },
                            ForceHttps = true,
                        },
                        new Flyio.Flyio.Inputs.FlyMachinePortArgs
                        {
                            Port = 443,
                            Handlers = new[]
                            {
                                "tls",
                                "http",
                            },
                        },
                    },
                    Protocol = "tcp",
                    Checks = new[]
                    {
                        new Flyio.Flyio.Inputs.FlyMachineCheckArgs
                        {
                            Interval = "10s",
                            GracePeriod = "5s",
                            Method = "get",
                            Path = "/",
                            Protocol = "http",
                            Timeout = "2s",
                            TlsSkipVerify = true,
                        },
                    },
                    Autostop = "suspend",
                    Autostart = true,
                    Concurrency = new Flyio.Flyio.Inputs.FlyMachineServiceConcurrencyArgs
                    {
                        SoftLimit = 2000,
                        HardLimit = 2500,
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            machineIad1,
        },
    });

    var volumeIad = new Flyio.Volume("volume-iad", new()
    {
        Name = "volume_iad",
        AutoBackupEnabled = true,
        Region = "iad",
        App = app.Name,
        SizeGb = 5,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            machineIad1,
        },
    });

    var volumeSea = new Flyio.Volume("volume-sea", new()
    {
        Name = "volume_sea",
        Region = "sea",
        App = app.Name,
        SizeGb = 5,
    }, new CustomResourceOptions
    {
        DependsOn =
        {
            machineSea1,
        },
    });

    var ipv4 = new Flyio.IP("ipv4", new()
    {
        Region = "sea",
        App = app.Name,
        AddrType = "v4",
    });

    var ipv6 = new Flyio.IP("ipv6", new()
    {
        Region = "sea",
        App = app.Name,
        AddrType = "v6",
    });

    var certificate = new Flyio.Certificate("certificate", new()
    {
        App = app.Name,
        Hostname = "pulumi-flyio.lshay.land",
    });

    return new Dictionary<string, object?>
    {
        ["appName"] = 
        {
            { "value", app.Name },
        },
    };
});

