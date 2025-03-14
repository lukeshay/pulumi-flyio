// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * A component that waits for a Fly.io SSL/TLS certificate to be fully issued.
 */
export class CertificateIssuanceWaiter extends pulumi.CustomResource {
    /**
     * Get an existing CertificateIssuanceWaiter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateIssuanceWaiter {
        return new CertificateIssuanceWaiter(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'flyio:index:CertificateIssuanceWaiter';

    /**
     * Returns true if the given object is an instance of CertificateIssuanceWaiter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateIssuanceWaiter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateIssuanceWaiter.__pulumiType;
    }

    /**
     * The name of the Fly app.
     */
    public readonly app!: pulumi.Output<string>;
    /**
     * The certificate authority used.
     */
    public /*out*/ readonly certificateAuthority!: pulumi.Output<string>;
    /**
     * The Fly.io certificate ID.
     */
    public /*out*/ readonly certificateId!: pulumi.Output<string>;
    /**
     * The status of the certificate.
     */
    public /*out*/ readonly clientStatus!: pulumi.Output<string>;
    /**
     * Expiration time for the ECDSA certificate.
     */
    public /*out*/ readonly ecdsaExpiresAt!: pulumi.Output<outputs.time.Time | undefined>;
    /**
     * The hostname for the certificate.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The input arguments used for the certificate issuance waiter.
     */
    public /*out*/ readonly input!: pulumi.Output<outputs.CertificateIssuanceWaiterArgs>;
    /**
     * Whether the certificate is fully issued (has both ECDSA and RSA certificates).
     */
    public /*out*/ readonly isFullyIssued!: pulumi.Output<boolean>;
    /**
     * When the certificate was fully issued.
     */
    public /*out*/ readonly issuedAt!: pulumi.Output<outputs.time.Time>;
    /**
     * Expiration time for the RSA certificate.
     */
    public /*out*/ readonly rsaExpiresAt!: pulumi.Output<outputs.time.Time | undefined>;
    /**
     * The timeout duration used for waiting.
     */
    public readonly timeout!: pulumi.Output<number>;

    /**
     * Create a CertificateIssuanceWaiter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateIssuanceWaiterArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.app === undefined) && !opts.urn) {
                throw new Error("Missing required property 'app'");
            }
            if ((!args || args.hostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostname'");
            }
            if ((!args || args.timeout === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeout'");
            }
            resourceInputs["app"] = args ? args.app : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["certificateAuthority"] = undefined /*out*/;
            resourceInputs["certificateId"] = undefined /*out*/;
            resourceInputs["clientStatus"] = undefined /*out*/;
            resourceInputs["ecdsaExpiresAt"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["isFullyIssued"] = undefined /*out*/;
            resourceInputs["issuedAt"] = undefined /*out*/;
            resourceInputs["rsaExpiresAt"] = undefined /*out*/;
        } else {
            resourceInputs["app"] = undefined /*out*/;
            resourceInputs["certificateAuthority"] = undefined /*out*/;
            resourceInputs["certificateId"] = undefined /*out*/;
            resourceInputs["clientStatus"] = undefined /*out*/;
            resourceInputs["ecdsaExpiresAt"] = undefined /*out*/;
            resourceInputs["hostname"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["isFullyIssued"] = undefined /*out*/;
            resourceInputs["issuedAt"] = undefined /*out*/;
            resourceInputs["rsaExpiresAt"] = undefined /*out*/;
            resourceInputs["timeout"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CertificateIssuanceWaiter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateIssuanceWaiter resource.
 */
export interface CertificateIssuanceWaiterArgs {
    /**
     * The name of the Fly app that the certificate is associated with.
     */
    app: pulumi.Input<string>;
    /**
     * The hostname for the certificate (e.g., example.com).
     */
    hostname: pulumi.Input<string>;
    /**
     * The maximum time to wait for the certificate to be fully issued. Formatted like 5s, 5m, etc.
     */
    timeout: pulumi.Input<number>;
}
