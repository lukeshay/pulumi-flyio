// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio
{
    /// <summary>
    /// A component that waits for a Fly.io SSL/TLS certificate to be fully issued.
    /// </summary>
    [FlyioResourceType("flyio:index:CertificateIssuanceWaiter")]
    public partial class CertificateIssuanceWaiter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Fly app.
        /// </summary>
        [Output("app")]
        public Output<string> App { get; private set; } = null!;

        /// <summary>
        /// The certificate authority used.
        /// </summary>
        [Output("certificateAuthority")]
        public Output<string> CertificateAuthority { get; private set; } = null!;

        /// <summary>
        /// The Fly.io certificate ID.
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// The status of the certificate.
        /// </summary>
        [Output("clientStatus")]
        public Output<string> ClientStatus { get; private set; } = null!;

        /// <summary>
        /// Expiration time for the ECDSA certificate.
        /// </summary>
        [Output("ecdsaExpiresAt")]
        public Output<PulumiFlyio.Flyio.Time.Outputs.Time?> EcdsaExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The hostname for the certificate.
        /// </summary>
        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        /// <summary>
        /// The input arguments used for the certificate issuance waiter.
        /// </summary>
        [Output("input")]
        public Output<Outputs.CertificateIssuanceWaiterArgs> Input { get; private set; } = null!;

        /// <summary>
        /// Whether the certificate is fully issued (has both ECDSA and RSA certificates).
        /// </summary>
        [Output("isFullyIssued")]
        public Output<bool> IsFullyIssued { get; private set; } = null!;

        /// <summary>
        /// When the certificate was fully issued.
        /// </summary>
        [Output("issuedAt")]
        public Output<PulumiFlyio.Flyio.Time.Outputs.Time> IssuedAt { get; private set; } = null!;

        /// <summary>
        /// Expiration time for the RSA certificate.
        /// </summary>
        [Output("rsaExpiresAt")]
        public Output<PulumiFlyio.Flyio.Time.Outputs.Time?> RsaExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The timeout duration used for waiting.
        /// </summary>
        [Output("timeout")]
        public Output<int> Timeout { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateIssuanceWaiter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateIssuanceWaiter(string name, CertificateIssuanceWaiterArgs args, CustomResourceOptions? options = null)
            : base("flyio:index:CertificateIssuanceWaiter", name, args ?? new CertificateIssuanceWaiterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateIssuanceWaiter(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("flyio:index:CertificateIssuanceWaiter", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lukeshay",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateIssuanceWaiter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateIssuanceWaiter Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateIssuanceWaiter(name, id, options);
        }
    }

    public sealed class CertificateIssuanceWaiterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Fly app that the certificate is associated with.
        /// </summary>
        [Input("app", required: true)]
        public Input<string> App { get; set; } = null!;

        /// <summary>
        /// The hostname for the certificate (e.g., example.com).
        /// </summary>
        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        /// <summary>
        /// The maximum time to wait for the certificate to be fully issued. Formatted like 5s, 5m, etc.
        /// </summary>
        [Input("timeout", required: true)]
        public Input<int> Timeout { get; set; } = null!;

        public CertificateIssuanceWaiterArgs()
        {
        }
        public static new CertificateIssuanceWaiterArgs Empty => new CertificateIssuanceWaiterArgs();
    }
}
