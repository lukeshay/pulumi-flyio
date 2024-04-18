import * as pulumi from "@pulumi/pulumi";
import * as flyio from "@pulumi/flyio";

const appNameResource = new flyio.Random("appName", {length: 24});
const app = new flyio.App("app", {
    appName: pulumi.interpolate`pulumi-${appNameResource.result}`,
    orgSlug: "luke-shay",
});
const machineSea1 = new flyio.Machine("machine-sea-1", {
    name: "machine-sea-1",
    updateStrategy: "bluegreen",
    region: "sea",
    appName: app.name,
    waitForChecks: 60000,
    waitForUpdate: 1000,
    config: {
        image: "nginx:latest",
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
            autostop: true,
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
    appName: app.name,
    waitForChecks: 60000,
    config: {
        image: "nginx:latest",
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
            autostop: true,
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
    updateStrategy: "bluegreen",
    region: "iad",
    appName: app.name,
    waitForUpdate: 1000,
    waitForChecks: 60000,
    config: {
        image: "nginx:latest",
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
            autostop: true,
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
    appName: app.name,
    waitForChecks: 60000,
    config: {
        image: "nginx:latest",
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
            autostop: true,
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
export const appName = {
    value: app.name,
};
