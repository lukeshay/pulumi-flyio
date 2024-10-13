package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.flyio.Random;
import com.pulumi.flyio.RandomArgs;
import com.pulumi.flyio.App;
import com.pulumi.flyio.AppArgs;
import com.pulumi.flyio.Machine;
import com.pulumi.flyio.MachineArgs;
import com.pulumi.flyio.flyio.inputs.FlyMachineConfigArgs;
import com.pulumi.flyio.flyio.inputs.FlyMachineGuestArgs;
import com.pulumi.flyio.Volume;
import com.pulumi.flyio.VolumeArgs;
import com.pulumi.flyio.IP;
import com.pulumi.flyio.IPArgs;
import com.pulumi.flyio.Certificate;
import com.pulumi.flyio.CertificateArgs;
import com.pulumi.flyio.WireGuardPeer;
import com.pulumi.flyio.WireGuardPeerArgs;
import com.pulumi.flyio.WireGuardToken;
import com.pulumi.flyio.WireGuardTokenArgs;
import com.pulumi.resources.CustomResourceOptions;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var appNameResource = new Random("appNameResource", RandomArgs.builder()
            .length(24)
            .build());

        var app = new App("app", AppArgs.builder()
            .name(appNameResource.result().applyValue(result -> String.format("pulumi-%s", result)))
            .org("luke-shay")
            .network("pulumi-flyio")
            .build());

        var machineSea1 = new Machine("machineSea1", MachineArgs.builder()
            .name("machine-sea-1")
            .region("sea")
            .app(app.name())
            .deploymentStrategy("bluegreen")
            .config(FlyMachineConfigArgs.builder()
                .image("nginxdemos/hello:0.4")
                .guest(FlyMachineGuestArgs.builder()
                    .cpus(1)
                    .cpuKind("shared")
                    .memoryMb(256)
                    .build())
                .services(FlyMachineServiceArgs.builder()
                    .internalPort(80)
                    .ports(                    
                        FlyMachinePortArgs.builder()
                            .port(80)
                            .handlers("http")
                            .forceHttps(true)
                            .build(),
                        FlyMachinePortArgs.builder()
                            .port(443)
                            .handlers(                            
                                "tls",
                                "http")
                            .build())
                    .protocol("tcp")
                    .checks(FlyMachineCheckArgs.builder()
                        .interval("10s")
                        .gracePeriod("5s")
                        .method("get")
                        .path("/")
                        .protocol("http")
                        .timeout("2s")
                        .tlsSkipVerify(true)
                        .build())
                    .autostop("suspend")
                    .autostart(true)
                    .concurrency(FlyMachineServiceConcurrencyArgs.builder()
                        .softLimit(2000)
                        .hardLimit(2500)
                        .build())
                    .build())
                .build())
            .build());

        var machineSea2 = new Machine("machineSea2", MachineArgs.builder()
            .name("machine-sea-2")
            .region("sea")
            .app(app.name())
            .config(FlyMachineConfigArgs.builder()
                .image("nginxdemos/hello:latest")
                .guest(FlyMachineGuestArgs.builder()
                    .cpus(1)
                    .cpuKind("shared")
                    .memoryMb(256)
                    .build())
                .services(FlyMachineServiceArgs.builder()
                    .internalPort(80)
                    .ports(                    
                        FlyMachinePortArgs.builder()
                            .port(80)
                            .handlers("http")
                            .forceHttps(true)
                            .build(),
                        FlyMachinePortArgs.builder()
                            .port(443)
                            .handlers(                            
                                "tls",
                                "http")
                            .build())
                    .protocol("tcp")
                    .checks(FlyMachineCheckArgs.builder()
                        .interval("10s")
                        .gracePeriod("5s")
                        .method("get")
                        .path("/")
                        .protocol("http")
                        .timeout("2s")
                        .tlsSkipVerify(true)
                        .build())
                    .autostop("suspend")
                    .autostart(true)
                    .concurrency(FlyMachineServiceConcurrencyArgs.builder()
                        .softLimit(2000)
                        .hardLimit(2500)
                        .build())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(machineSea1)
                .build());

        var machineIad1 = new Machine("machineIad1", MachineArgs.builder()
            .name("machine-iad-1")
            .region("iad")
            .app(app.name())
            .config(FlyMachineConfigArgs.builder()
                .image("nginxdemos/hello:latest")
                .guest(FlyMachineGuestArgs.builder()
                    .cpus(1)
                    .cpuKind("shared")
                    .memoryMb(512)
                    .build())
                .services(FlyMachineServiceArgs.builder()
                    .internalPort(80)
                    .ports(                    
                        FlyMachinePortArgs.builder()
                            .port(80)
                            .handlers("http")
                            .forceHttps(true)
                            .build(),
                        FlyMachinePortArgs.builder()
                            .port(443)
                            .handlers(                            
                                "tls",
                                "http")
                            .build())
                    .protocol("tcp")
                    .checks(FlyMachineCheckArgs.builder()
                        .interval("10s")
                        .gracePeriod("5s")
                        .method("get")
                        .path("/")
                        .protocol("http")
                        .timeout("2s")
                        .tlsSkipVerify(true)
                        .build())
                    .autostop("suspend")
                    .autostart(true)
                    .concurrency(FlyMachineServiceConcurrencyArgs.builder()
                        .softLimit(2000)
                        .hardLimit(2500)
                        .build())
                    .build())
                .build())
            .build());

        var machineIad2 = new Machine("machineIad2", MachineArgs.builder()
            .name("machine-iad-2")
            .region("iad")
            .app(app.name())
            .config(FlyMachineConfigArgs.builder()
                .image("nginxdemos/hello:latest")
                .guest(FlyMachineGuestArgs.builder()
                    .cpus(1)
                    .cpuKind("shared")
                    .memoryMb(256)
                    .build())
                .services(FlyMachineServiceArgs.builder()
                    .internalPort(80)
                    .ports(                    
                        FlyMachinePortArgs.builder()
                            .port(80)
                            .handlers("http")
                            .forceHttps(true)
                            .build(),
                        FlyMachinePortArgs.builder()
                            .port(443)
                            .handlers(                            
                                "tls",
                                "http")
                            .build())
                    .protocol("tcp")
                    .checks(FlyMachineCheckArgs.builder()
                        .interval("10s")
                        .gracePeriod("5s")
                        .method("get")
                        .path("/")
                        .protocol("http")
                        .timeout("2s")
                        .tlsSkipVerify(true)
                        .build())
                    .autostop("suspend")
                    .autostart(true)
                    .concurrency(FlyMachineServiceConcurrencyArgs.builder()
                        .softLimit(2000)
                        .hardLimit(2500)
                        .build())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(machineIad1)
                .build());

        var volumeIad = new Volume("volumeIad", VolumeArgs.builder()
            .name("volume_iad")
            .autoBackupEnabled(true)
            .region("iad")
            .app(app.name())
            .sizeGb(5)
            .build(), CustomResourceOptions.builder()
                .dependsOn(machineIad1)
                .build());

        var volumeSea = new Volume("volumeSea", VolumeArgs.builder()
            .name("volume_sea")
            .region("sea")
            .app(app.name())
            .sizeGb(5)
            .build(), CustomResourceOptions.builder()
                .dependsOn(machineSea1)
                .build());

        var ipv4 = new IP("ipv4", IPArgs.builder()
            .region("sea")
            .app(app.name())
            .addrType("v4")
            .build());

        var ipv6 = new IP("ipv6", IPArgs.builder()
            .region("sea")
            .app(app.name())
            .addrType("v6")
            .build());

        var privateipv6 = new IP("privateipv6", IPArgs.builder()
            .region("sea")
            .app(app.name())
            .addrType("private_v6")
            .network("pulumi-flyio")
            .build());

        var certificate = new Certificate("certificate", CertificateArgs.builder()
            .app(app.name())
            .hostname("pulumi-flyio.lshay.land")
            .build());

        var wireguardPeer = new WireGuardPeer("wireguardPeer", WireGuardPeerArgs.builder()
            .org("luke-shay")
            .build());

        var wireguardToken = new WireGuardToken("wireguardToken", WireGuardTokenArgs.builder()
            .org("luke-shay")
            .build());

        ctx.export("appName", Map.of("value", app.name()));
    }
}
