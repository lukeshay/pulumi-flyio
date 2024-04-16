import * as flyio from "pulumi-flyio"
import * as pulumi from "@pulumi/pulumi"
import * as docker from "@pulumi/docker"

const name = new flyio.Random("appName", { length: 24 })
const appName = pulumi.interpolate`pulumi-${name.result}`
const app = new flyio.App("app", {
  appName,
  orgSlug: "luke-shay",
})

const image = new docker.Image("image", {
  build: {
    context: ".",
    dockerfile: "Dockerfile",
    platform: "linux/amd64",
  },
  imageName: pulumi.interpolate`registry.fly.io/${appName}:latest`,
  buildOnPreview: true,
})

const createMachineConfig = (region: string): flyio.MachineArgs => ({
  region,
  appName: app.name,
  waitForChecks: 60_000,
  config: {
    image: image.imageName,
    guest: {
      cpus: 1,
      cpuKind: "shared",
      memoryMb: 256,
    },
    init: {
      cmd: ["-enable-health"]
    },
    services: [
      {
        internalPort: 8043,
        ports: [
          {
            port: 80,
            handlers: ["http"],
            forceHttps: true,
          },
          {
            port: 443,
            handlers: ["tls", "http"],
          },
        ],
        protocol: "tcp",
        checks: [
          {
            interval: "10s",
            gracePeriod: "5s",
            method: "get",
            path: "/",
            protocol: "http",
            timeout: "2s",
            tlsSkipVerify: true,
          },
        ],
      },
    ],
  },
})

const machine = new flyio.Machine("machine", createMachineConfig("dfw"))

export const output = {
  appId: app.flyId,
  appName: app.name,
  machineId: machine.flyId,
  machineName: machine.name,
  machineChecks: machine.checks,
}
