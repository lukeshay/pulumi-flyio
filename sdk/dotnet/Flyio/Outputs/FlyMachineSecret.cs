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
    public sealed class FlyMachineSecret
    {
        public readonly string? EnvVar;
        public readonly string? Name;

        [OutputConstructor]
        private FlyMachineSecret(
            string? envVar,

            string? name)
        {
            EnvVar = envVar;
            Name = name;
        }
    }
}
