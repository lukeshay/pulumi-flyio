// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace PulumiFlyio.Flyio.Flyio.Outputs
{

    [OutputType]
    public sealed class FlyContainerHealthcheck
    {
        public readonly Outputs.FlyExecHealthcheck? Exec;
        public readonly int? FailureThreshold;
        public readonly int? GracePeriod;
        public readonly Outputs.FlyHTTPHealthcheck? Http;
        public readonly int? Interval;
        public readonly string? Kind;
        public readonly string? Name;
        public readonly int? SuccessThreshold;
        public readonly Outputs.FlyTCPHealthcheck? Tcp;
        public readonly int? Timeout;
        public readonly string? Unhealthy;

        [OutputConstructor]
        private FlyContainerHealthcheck(
            Outputs.FlyExecHealthcheck? exec,

            int? failureThreshold,

            int? gracePeriod,

            Outputs.FlyHTTPHealthcheck? http,

            int? interval,

            string? kind,

            string? name,

            int? successThreshold,

            Outputs.FlyTCPHealthcheck? tcp,

            int? timeout,

            string? unhealthy)
        {
            Exec = exec;
            FailureThreshold = failureThreshold;
            GracePeriod = gracePeriod;
            Http = http;
            Interval = interval;
            Kind = kind;
            Name = name;
            SuccessThreshold = successThreshold;
            Tcp = tcp;
            Timeout = timeout;
            Unhealthy = unhealthy;
        }
    }
}
