// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * A Fly.io SSL/TLS certificate for an app's domain.
 */
export class Certificate extends pulumi.CustomResource {
    /**
     * Get an existing Certificate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Certificate {
        return new Certificate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'flyio:index:Certificate';

    /**
     * Returns true if the given object is an instance of Certificate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Certificate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Certificate.__pulumiType;
    }

    /**
     * Whether ACME ALPN verification is configured.
     */
    public /*out*/ readonly acmeAlpnConfigured!: pulumi.Output<boolean>;
    /**
     * Whether ACME DNS verification is configured.
     */
    public /*out*/ readonly acmeDnsConfigured!: pulumi.Output<boolean>;
    /**
     * The name of the Fly app.
     */
    public readonly app!: pulumi.Output<string>;
    /**
     * The certificate authority used.
     */
    public /*out*/ readonly certificateAuthority!: pulumi.Output<string>;
    /**
     * The status of the certificate.
     */
    public /*out*/ readonly clientStatus!: pulumi.Output<string>;
    /**
     * Whether the certificate is fully configured.
     */
    public /*out*/ readonly configured!: pulumi.Output<boolean>;
    /**
     * When the certificate was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<outputs.time.Time>;
    /**
     * The DNS provider for the hostname.
     */
    public /*out*/ readonly dnsProvider!: pulumi.Output<string>;
    /**
     * Hostname for DNS validation.
     */
    public /*out*/ readonly dnsValidationHostname!: pulumi.Output<string>;
    /**
     * Instructions for DNS validation.
     */
    public /*out*/ readonly dnsValidationInstructions!: pulumi.Output<string>;
    /**
     * Target for DNS validation.
     */
    public /*out*/ readonly dnsValidationTarget!: pulumi.Output<string>;
    /**
     * The Fly.io certificate ID.
     */
    public /*out*/ readonly flyId!: pulumi.Output<string>;
    /**
     * The hostname for the certificate.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The input arguments used to create the certificate.
     */
    public /*out*/ readonly input!: pulumi.Output<outputs.CertificateArgs>;
    /**
     * Whether the hostname is an apex domain.
     */
    public /*out*/ readonly isApex!: pulumi.Output<boolean>;
    /**
     * Whether the certificate is a wildcard certificate.
     */
    public /*out*/ readonly isWildcard!: pulumi.Output<boolean>;
    /**
     * The source of the certificate.
     */
    public /*out*/ readonly source!: pulumi.Output<string>;

    /**
     * Create a Certificate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.app === undefined) && !opts.urn) {
                throw new Error("Missing required property 'app'");
            }
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            resourceInputs["app"] = args ? args.app : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["acmeAlpnConfigured"] = undefined /*out*/;
            resourceInputs["acmeDnsConfigured"] = undefined /*out*/;
            resourceInputs["certificateAuthority"] = undefined /*out*/;
            resourceInputs["clientStatus"] = undefined /*out*/;
            resourceInputs["configured"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dnsProvider"] = undefined /*out*/;
            resourceInputs["dnsValidationHostname"] = undefined /*out*/;
            resourceInputs["dnsValidationInstructions"] = undefined /*out*/;
            resourceInputs["dnsValidationTarget"] = undefined /*out*/;
            resourceInputs["flyId"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["isApex"] = undefined /*out*/;
            resourceInputs["isWildcard"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
        } else {
            resourceInputs["acmeAlpnConfigured"] = undefined /*out*/;
            resourceInputs["acmeDnsConfigured"] = undefined /*out*/;
            resourceInputs["app"] = undefined /*out*/;
            resourceInputs["certificateAuthority"] = undefined /*out*/;
            resourceInputs["clientStatus"] = undefined /*out*/;
            resourceInputs["configured"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["dnsProvider"] = undefined /*out*/;
            resourceInputs["dnsValidationHostname"] = undefined /*out*/;
            resourceInputs["dnsValidationInstructions"] = undefined /*out*/;
            resourceInputs["dnsValidationTarget"] = undefined /*out*/;
            resourceInputs["flyId"] = undefined /*out*/;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["isApex"] = undefined /*out*/;
            resourceInputs["isWildcard"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Certificate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Certificate resource.
 */
export interface CertificateArgs {
    /**
     * The name of the Fly app to add the certificate to.
     */
    app: pulumi.Input<string>;
    /**
     * The hostname for the certificate (e.g., example.com).
     */
    hostname: pulumi.Input<string>;
}
