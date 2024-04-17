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
			cmd: ["-enable-health"],
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
				autostop: true,
				autostart: true,
				concurrency: {
					softLimit: 2000,
					hardLimit: 2500,
				},
			},
		],
	},
})

const machines = [
	{
		id: "1",
		region: "dfw",
	},
	{
		id: "2",
		region: "dfw",
	},
	{
		id: "1",
		region: "sea",
	},
	{
		id: "2",
		region: "sea",
	},
	{
		id: "1",
		region: "iad",
	},
	{
		id: "2",
		region: "iad",
	},
].map(
	(m) =>
		new flyio.Machine(`machine-${m.region}-${m.id}`, {
			...createMachineConfig(m.region),
			name: pulumi.interpolate`machine-${m.region}-${m.id}`,
		}),
)

export const output = {
	appId: app.flyId,
	appName: app.name,
}
