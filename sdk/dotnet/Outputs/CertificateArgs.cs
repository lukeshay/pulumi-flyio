// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio.Outputs
{

    [OutputType]
    public sealed class CertificateArgs
    {
        public readonly string App;
        public readonly string Hostname;

        [OutputConstructor]
        private CertificateArgs(
            string app,

            string hostname)
        {
            App = app;
            Hostname = hostname;
        }
    }
}
