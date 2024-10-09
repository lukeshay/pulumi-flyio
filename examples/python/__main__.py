import pulumi
import pulumi_flyio as flyio

app_name_resource = flyio.Random("appName", length=24)
app = flyio.App("app",
    app_name=app_name_resource.result.apply(lambda result: f"pulumi-{result}"),
    org_slug="luke-shay")
machine_sea1 = flyio.Machine("machine-sea-1",
    name="machine-sea-1",
    update_strategy="bluegreen",
    region="sea",
    app_name=app.name,
    config={
        "image": "nginxdemos/hello:0.4",
        "guest": {
            "cpus": 1,
            "cpu_kind": "shared",
            "memory_mb": 256,
        },
        "services": [{
            "internal_port": 80,
            "ports": [
                {
                    "port": 80,
                    "handlers": ["http"],
                    "force_https": True,
                },
                {
                    "port": 443,
                    "handlers": [
                        "tls",
                        "http",
                    ],
                },
            ],
            "protocol": "tcp",
            "checks": [{
                "interval": "10s",
                "grace_period": "5s",
                "method": "get",
                "path": "/",
                "protocol": "http",
                "timeout": "2s",
                "tls_skip_verify": True,
            }],
            "autostop": "suspend",
            "autostart": True,
            "concurrency": {
                "soft_limit": 2000,
                "hard_limit": 2500,
            },
        }],
    })
machine_sea2 = flyio.Machine("machine-sea-2",
    name="machine-sea-2",
    region="sea",
    app_name=app.name,
    config={
        "image": "nginxdemos/hello:latest",
        "guest": {
            "cpus": 1,
            "cpu_kind": "shared",
            "memory_mb": 256,
        },
        "services": [{
            "internal_port": 80,
            "ports": [
                {
                    "port": 80,
                    "handlers": ["http"],
                    "force_https": True,
                },
                {
                    "port": 443,
                    "handlers": [
                        "tls",
                        "http",
                    ],
                },
            ],
            "protocol": "tcp",
            "checks": [{
                "interval": "10s",
                "grace_period": "5s",
                "method": "get",
                "path": "/",
                "protocol": "http",
                "timeout": "2s",
                "tls_skip_verify": True,
            }],
            "autostop": "suspend",
            "autostart": True,
            "concurrency": {
                "soft_limit": 2000,
                "hard_limit": 2500,
            },
        }],
    },
    opts = pulumi.ResourceOptions(depends_on=[machine_sea1]))
machine_iad1 = flyio.Machine("machine-iad-1",
    name="machine-iad-1",
    update_strategy="bluegreen",
    region="iad",
    app_name=app.name,
    config={
        "image": "nginxdemos/hello:latest",
        "guest": {
            "cpus": 1,
            "cpu_kind": "shared",
            "memory_mb": 512,
        },
        "services": [{
            "internal_port": 80,
            "ports": [
                {
                    "port": 80,
                    "handlers": ["http"],
                    "force_https": True,
                },
                {
                    "port": 443,
                    "handlers": [
                        "tls",
                        "http",
                    ],
                },
            ],
            "protocol": "tcp",
            "checks": [{
                "interval": "10s",
                "grace_period": "5s",
                "method": "get",
                "path": "/",
                "protocol": "http",
                "timeout": "2s",
                "tls_skip_verify": True,
            }],
            "autostop": "suspend",
            "autostart": True,
            "concurrency": {
                "soft_limit": 2000,
                "hard_limit": 2500,
            },
        }],
    })
machine_iad2 = flyio.Machine("machine-iad-2",
    name="machine-iad-2",
    region="iad",
    app_name=app.name,
    config={
        "image": "nginxdemos/hello:latest",
        "guest": {
            "cpus": 1,
            "cpu_kind": "shared",
            "memory_mb": 256,
        },
        "services": [{
            "internal_port": 80,
            "ports": [
                {
                    "port": 80,
                    "handlers": ["http"],
                    "force_https": True,
                },
                {
                    "port": 443,
                    "handlers": [
                        "tls",
                        "http",
                    ],
                },
            ],
            "protocol": "tcp",
            "checks": [{
                "interval": "10s",
                "grace_period": "5s",
                "method": "get",
                "path": "/",
                "protocol": "http",
                "timeout": "2s",
                "tls_skip_verify": True,
            }],
            "autostop": "suspend",
            "autostart": True,
            "concurrency": {
                "soft_limit": 2000,
                "hard_limit": 2500,
            },
        }],
    },
    opts = pulumi.ResourceOptions(depends_on=[machine_iad1]))
volume_iad = flyio.Volume("volume-iad",
    name="volume_iad",
    region="iad",
    app_name=app.name,
    size_gb=5,
    opts = pulumi.ResourceOptions(depends_on=[machine_iad1]))
volume_sea = flyio.Volume("volume-sea",
    name="volume_sea",
    region="sea",
    app_name=app.name,
    size_gb=5,
    opts = pulumi.ResourceOptions(depends_on=[machine_sea1]))
pulumi.export("appName", {
    "value": app.name,
})
