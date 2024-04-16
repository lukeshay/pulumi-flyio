// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Outputs
{

    [OutputType]
    public sealed class FlyMachineInit
    {
        public readonly ImmutableArray<string> Cmd;
        public readonly ImmutableArray<string> Entrypoint;
        public readonly ImmutableArray<string> Exec;
        public readonly ImmutableArray<string> KernelArgs;
        public readonly int? SwapSizeMb;
        public readonly bool? Tty;

        [OutputConstructor]
        private FlyMachineInit(
            ImmutableArray<string> cmd,

            ImmutableArray<string> entrypoint,

            ImmutableArray<string> exec,

            ImmutableArray<string> kernelArgs,

            int? swapSizeMb,

            bool? tty)
        {
            Cmd = cmd;
            Entrypoint = entrypoint;
            Exec = exec;
            KernelArgs = kernelArgs;
            SwapSizeMb = swapSizeMb;
            Tty = tty;
        }
    }
}
