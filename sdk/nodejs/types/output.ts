// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AppArgs {
    appName: string;
    enableSubdomains?: boolean;
    network?: string;
    orgSlug: string;
}

export interface MachineArgs {
    appName: string;
    config: outputs.flyio.FlyMachineConfig;
    leaseTtl?: number;
    lsvd?: boolean;
    name?: string;
    region?: string;
    skipLaunch?: boolean;
    skipServiceRegistration?: boolean;
    updateStrategy?: string;
    waitForChecks?: number;
    waitForUpdate?: number;
}

export interface VolumeArgs {
    appName: string;
    autoBackupEnabled?: boolean;
    compute?: outputs.flyio.FlyMachineGuest;
    computeImage?: string;
    encrypted?: boolean;
    fstype?: string;
    machinesOnly?: boolean;
    name?: string;
    region?: string;
    requireUniqueZone?: boolean;
    sizeGb?: number;
    snapshotId?: string;
    snapshotRetention?: number;
    sourceVolumeId?: string;
}

export namespace flyio {
    export interface CheckStatus {
        name?: string;
        output?: string;
        status?: string;
        updatedAt?: string;
    }

    export interface FlyDNSConfig {
        dnsForwardRules?: outputs.flyio.FlyDnsForwardRule[];
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

    export interface FlyFile {
        guestPath?: string;
        rawValue?: string;
        secretName?: string;
    }

    export interface FlyHTTPOptions {
        compress?: boolean;
        h2Backend?: boolean;
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
        maxRetries?: number;
        policy?: string;
    }

    export interface FlyMachineSecret {
        envVar?: string;
        name?: string;
    }

    export interface FlyMachineService {
        autostart?: boolean;
        autostop?: boolean;
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
        tigrisBucket?: string;
        urlPrefix: string;
    }

    export interface FlyStopConfig {
        signal?: string;
        timeout?: string;
    }

    export interface FlyTLSOptions {
        alpn?: string[];
        defaultSelfSigned?: boolean;
        versions?: string[];
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

    export interface Organization {
        name: string;
        slug: string;
    }

}
