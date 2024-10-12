import * as pulumi from "@pulumi/pulumi";
import * as flyio from "pulumi-flyio";

const appNameResource = new flyio.Random("appName", {length: 24});
const app = new flyio.App("app", {
    name: pulumi.interpolate`pulumi-${appNameResource.result}`,
    org: "luke-shay",
});
const machineSea1 = new flyio.Machine("machine-sea-1", {
    name: "machine-sea-1",
    region: "sea",
    app: app.name,
    deploymentStrategy: "bluegreen",
    config: {
        image: "nginxdemos/hello:0.4",
        guest: {
            cpus: 1,
            cpuKind: "shared",
            memoryMb: 256,
        },
        services: [{
            internalPort: 80,
            ports: [
                {
                    port: 80,
                    handlers: ["http"],
                    forceHttps: true,
                },
                {
                    port: 443,
                    handlers: [
                        "tls",
                        "http",
                    ],
                },
            ],
            protocol: "tcp",
            checks: [{
                interval: "10s",
                gracePeriod: "5s",
                method: "get",
                path: "/",
                protocol: "http",
                timeout: "2s",
                tlsSkipVerify: true,
            }],
            autostop: "suspend",
            autostart: true,
            concurrency: {
                softLimit: 2000,
                hardLimit: 2500,
            },
        }],
    },
});
const machineSea2 = new flyio.Machine("machine-sea-2", {
    name: "machine-sea-2",
    region: "sea",
    app: app.name,
    config: {
        image: "nginxdemos/hello:latest",
        guest: {
            cpus: 1,
            cpuKind: "shared",
            memoryMb: 256,
        },
        services: [{
            internalPort: 80,
            ports: [
                {
                    port: 80,
                    handlers: ["http"],
                    forceHttps: true,
                },
                {
                    port: 443,
                    handlers: [
                        "tls",
                        "http",
                    ],
                },
            ],
            protocol: "tcp",
            checks: [{
                interval: "10s",
                gracePeriod: "5s",
                method: "get",
                path: "/",
                protocol: "http",
                timeout: "2s",
                tlsSkipVerify: true,
            }],
            autostop: "suspend",
            autostart: true,
            concurrency: {
                softLimit: 2000,
                hardLimit: 2500,
            },
        }],
    },
}, {
    dependsOn: [machineSea1],
});
const machineIad1 = new flyio.Machine("machine-iad-1", {
    name: "machine-iad-1",
    region: "iad",
    app: app.name,
    config: {
        image: "nginxdemos/hello:latest",
        guest: {
            cpus: 1,
            cpuKind: "shared",
            memoryMb: 512,
        },
        services: [{
            internalPort: 80,
            ports: [
                {
                    port: 80,
                    handlers: ["http"],
                    forceHttps: true,
                },
                {
                    port: 443,
                    handlers: [
                        "tls",
                        "http",
                    ],
                },
            ],
            protocol: "tcp",
            checks: [{
                interval: "10s",
                gracePeriod: "5s",
                method: "get",
                path: "/",
                protocol: "http",
                timeout: "2s",
                tlsSkipVerify: true,
            }],
            autostop: "suspend",
            autostart: true,
            concurrency: {
                softLimit: 2000,
                hardLimit: 2500,
            },
        }],
    },
});
const machineIad2 = new flyio.Machine("machine-iad-2", {
    name: "machine-iad-2",
    region: "iad",
    app: app.name,
    config: {
        image: "nginxdemos/hello:latest",
        guest: {
            cpus: 1,
            cpuKind: "shared",
            memoryMb: 256,
        },
        services: [{
            internalPort: 80,
            ports: [
                {
                    port: 80,
                    handlers: ["http"],
                    forceHttps: true,
                },
                {
                    port: 443,
                    handlers: [
                        "tls",
                        "http",
                    ],
                },
            ],
            protocol: "tcp",
            checks: [{
                interval: "10s",
                gracePeriod: "5s",
                method: "get",
                path: "/",
                protocol: "http",
                timeout: "2s",
                tlsSkipVerify: true,
            }],
            autostop: "suspend",
            autostart: true,
            concurrency: {
                softLimit: 2000,
                hardLimit: 2500,
            },
        }],
    },
}, {
    dependsOn: [machineIad1],
});
const volumeIad = new flyio.Volume("volume-iad", {
    name: "volume_iad",
    autoBackupEnabled: true,
    region: "iad",
    app: app.name,
    sizeGb: 5,
}, {
    dependsOn: [machineIad1],
});
const volumeSea = new flyio.Volume("volume-sea", {
    name: "volume_sea",
    region: "sea",
    app: app.name,
    sizeGb: 5,
}, {
    dependsOn: [machineSea1],
});
const ipv4 = new flyio.IP("ipv4", {
    region: "sea",
    app: app.name,
    addrType: "v4",
});
const ipv6 = new flyio.IP("ipv6", {
    region: "sea",
    app: app.name,
    addrType: "v6",
});
const certificate = new flyio.Certificate("certificate", {
    app: app.name,
    hostname: "pulumi-flyio.lshay.land",
});
export const appName = {
    value: app.name,
};
