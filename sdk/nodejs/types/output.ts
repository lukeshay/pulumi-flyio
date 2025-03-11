// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AppArgs {
    /**
     * Whether to enable subdomains for the application.
     */
    enableSubdomains?: boolean;
    /**
     * The name of the Fly.io application.
     */
    name: string;
    /**
     * The network the application belongs to.
     */
    network?: string;
    /**
     * The organization the application belongs to.
     */
    org: string;
}

export interface CertificateArgs {
    /**
     * The name of the Fly app to add the certificate to.
     */
    app: string;
    /**
     * The hostname for the certificate (e.g., example.com).
     */
    hostname: string;
}

export interface CertificateIssuanceWaiterArgs {
    /**
     * The name of the Fly app that the certificate is associated with.
     */
    app: string;
    /**
     * The hostname for the certificate (e.g., example.com).
     */
    hostname: string;
    /**
     * The maximum time to wait for the certificate to be fully issued. Formatted like 5s, 5m, etc.
     */
    timeout: number;
}

export interface IPArgs {
    /**
     * The type of IP address (v4 or v6).
     */
    addrType: string;
    /**
     * The name of the Fly.io application to allocate the IP address for.
     */
    app: string;
    /**
     * The network to allocate the IP address in.
     */
    network?: string;
    /**
     * The region to allocate the IP address in.
     */
    region: string;
}

export interface MachineArgs {
    /**
     * The Fly.io application to deploy the machine to.
     */
    app: string;
    config: outputs.flyio.FlyMachineConfig;
    /**
     * The deployment strategy for the machine.
     */
    deploymentStrategy?: string;
    leaseTtl?: number;
    lsvd?: boolean;
    name?: string;
    region?: string;
    skipLaunch?: boolean;
    skipServiceRegistration?: boolean;
}

export interface VolumeArgs {
    /**
     * The Fly.io application to attach the volume to.
     */
    app: string;
    /**
     * Whether to enable automatic backups for the volume.
     */
    autoBackupEnabled?: boolean;
    compute?: outputs.flyio.FlyMachineGuest;
    computeImage?: string;
    encrypted?: boolean;
    fstype?: string;
    name?: string;
    region?: string;
    requireUniqueZone?: boolean;
    sizeGb?: number;
    snapshotId?: string;
    snapshotRetention?: number;
    sourceVolumeId?: string;
    uniqueZoneAppWide?: boolean;
}

export namespace flyio {
    export interface CheckStatus {
        name?: string;
        output?: string;
        status?: string;
        updatedAt?: string;
    }

    export interface FlyContainerConfig {
        cmd?: string[];
        dependsOn?: outputs.flyio.FlyContainerDependency[];
        entrypoint?: string[];
        env?: {[key: string]: string};
        envFrom?: outputs.flyio.FlyEnvFrom[];
        exec?: string[];
        files?: outputs.flyio.FlyFile[];
        healthchecks?: outputs.flyio.FlyContainerHealthcheck[];
        image?: string;
        mounts?: outputs.flyio.FlyContainerMount[];
        name?: string;
        restart?: outputs.flyio.FlyMachineRestart;
        secrets?: outputs.flyio.FlyMachineSecret[];
        stop?: outputs.flyio.FlyStopConfig;
        user?: string;
    }

    export interface FlyContainerDependency {
        condition?: string;
        name?: string;
    }

    export interface FlyContainerHealthcheck {
        exec?: outputs.flyio.FlyExecHealthcheck;
        failureThreshold?: number;
        gracePeriod?: number;
        http?: outputs.flyio.FlyHTTPHealthcheck;
        interval?: number;
        kind?: string;
        name?: string;
        successThreshold?: number;
        tcp?: outputs.flyio.FlyTCPHealthcheck;
        timeout?: number;
        unhealthy?: string;
    }

    export interface FlyContainerMount {
        name?: string;
        path?: string;
    }

    export interface FlyDNSConfig {
        dnsForwardRules?: outputs.flyio.FlyDnsForwardRule[];
        hostname?: string;
        hostnameFqdn?: string;
        nameservers?: string[];
        options?: outputs.flyio.FlyDnsOption[];
        searches?: string[];
        skipRegistration?: boolean;
    }

    export interface FlyDnsForwardRule {
        addr?: string;
        basename?: string;
    }

    export interface FlyDnsOption {
        name?: string;
        value?: string;
    }

    export interface FlyEnvFrom {
        envVar?: string;
        fieldRef?: string;
    }

    export interface FlyExecHealthcheck {
        command?: string[];
    }

    export interface FlyFile {
        guestPath?: string;
        mode?: number;
        rawValue?: string;
        secretName?: string;
    }

    export interface FlyHTTPHealthcheck {
        headers?: outputs.flyio.FlyMachineHTTPHeader[];
        method?: string;
        path?: string;
        port?: number;
        scheme?: string;
        tlsServerName?: string;
        tlsSkipVerify?: boolean;
    }

    export interface FlyHTTPOptions {
        compress?: boolean;
        h2Backend?: boolean;
        headersReadTimeout?: number;
        idleTimeout?: number;
        response?: outputs.flyio.FlyHTTPResponseOptions;
    }

    export interface FlyHTTPResponseOptions {
        headers?: {[key: string]: {[key: string]: any}};
        pristine?: boolean;
    }

    export interface FlyMachineCheck {
        gracePeriod?: string;
        headers?: outputs.flyio.FlyMachineHTTPHeader[];
        interval?: string;
        kind?: string;
        method?: string;
        path?: string;
        port?: number;
        protocol?: string;
        timeout?: string;
        tlsServerName?: string;
        tlsSkipVerify?: boolean;
        type?: string;
    }

    export interface FlyMachineConfig {
        autoDestroy?: boolean;
        checks?: {[key: string]: outputs.flyio.FlyMachineCheck};
        containers?: outputs.flyio.FlyContainerConfig[];
        dns?: outputs.flyio.FlyDNSConfig;
        env?: {[key: string]: string};
        files?: outputs.flyio.FlyFile[];
        guest?: outputs.flyio.FlyMachineGuest;
        image: string;
        init?: outputs.flyio.FlyMachineInit;
        metadata?: {[key: string]: string};
        metrics?: outputs.flyio.FlyMachineMetrics;
        mounts?: outputs.flyio.FlyMachineMount[];
        processes?: outputs.flyio.FlyMachineProcess[];
        restart?: outputs.flyio.FlyMachineRestart;
        schedule?: string;
        services?: outputs.flyio.FlyMachineService[];
        standbys?: string[];
        statics?: outputs.flyio.FlyStatic[];
        stopConfig?: outputs.flyio.FlyStopConfig;
        volumes?: outputs.flyio.FlyVolumeConfig[];
    }

    export interface FlyMachineGuest {
        cpuKind: string;
        cpus: number;
        gpuKind?: string;
        gpus?: number;
        hostDedicationId?: string;
        kernelArgs?: string[];
        memoryMb: number;
    }

    export interface FlyMachineHTTPHeader {
        name?: string;
        values?: string[];
    }

    export interface FlyMachineInit {
        cmd?: string[];
        entrypoint?: string[];
        exec?: string[];
        kernelArgs?: string[];
        swapSizeMb?: number;
        tty?: boolean;
    }

    export interface FlyMachineMetrics {
        https?: boolean;
        path?: string;
        port?: number;
    }

    export interface FlyMachineMount {
        addSizeGb?: number;
        encrypted?: boolean;
        extendThresholdPercent?: number;
        name?: string;
        path?: string;
        sizeGb?: number;
        sizeGbLimit?: number;
        volume?: string;
    }

    export interface FlyMachinePort {
        endPort?: number;
        forceHttps?: boolean;
        handlers?: string[];
        httpOptions?: outputs.flyio.FlyHTTPOptions;
        port?: number;
        proxyProtoOptions?: outputs.flyio.FlyProxyProtoOptions;
        startPort?: number;
        tlsOptions?: outputs.flyio.FlyTLSOptions;
    }

    export interface FlyMachineProcess {
        cmd?: string[];
        entrypoint?: string[];
        env?: {[key: string]: string};
        envFrom?: outputs.flyio.FlyEnvFrom[];
        exec?: string[];
        ignoreAppSecrets?: boolean;
        secrets?: outputs.flyio.FlyMachineSecret[];
        user?: string;
    }

    export interface FlyMachineRestart {
        gpuBidPrice?: number;
        maxRetries?: number;
        policy?: string;
    }

    export interface FlyMachineSecret {
        envVar?: string;
        name?: string;
    }

    export interface FlyMachineService {
        autostart?: boolean;
        autostop?: string;
        checks?: outputs.flyio.FlyMachineCheck[];
        concurrency?: outputs.flyio.FlyMachineServiceConcurrency;
        forceInstanceDescription?: string;
        forceInstanceKey?: string;
        internalPort?: number;
        minMachinesRunning?: number;
        ports?: outputs.flyio.FlyMachinePort[];
        protocol?: string;
    }

    export interface FlyMachineServiceConcurrency {
        hardLimit?: number;
        softLimit?: number;
        type?: string;
    }

    export interface FlyProxyProtoOptions {
        version?: string;
    }

    export interface FlyStatic {
        guestPath: string;
        indexDocument?: string;
        tigrisBucket?: string;
        urlPrefix: string;
    }

    export interface FlyStopConfig {
        signal?: string;
        timeout?: string;
    }

    export interface FlyTCPHealthcheck {
        port?: number;
    }

    export interface FlyTLSOptions {
        alpn?: string[];
        defaultSelfSigned?: boolean;
        versions?: string[];
    }

    export interface FlyTempDirVolume {
        sizeMb?: number;
        storageType?: string;
    }

    export interface FlyVolumeConfig {
        name?: string;
        tempDir?: outputs.flyio.FlyTempDirVolume;
    }

    export interface ImageRef {
        digest?: string;
        labels?: {[key: string]: string};
        registry?: string;
        repository?: string;
        tag?: string;
    }

    export interface MachineEvent {
        flyId: string;
        request?: {[key: string]: any};
        source?: string;
        status?: string;
        timestamp?: number;
        type?: string;
    }

}

export namespace time {
    export interface Time {
    }

}
