// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyMachineGuestArgs : global::Pulumi.ResourceArgs
    {
        [Input("cpuKind", required: true)]
        public Input<string> CpuKind { get; set; } = null!;

        [Input("cpus", required: true)]
        public Input<int> Cpus { get; set; } = null!;

        [Input("gpuKind")]
        public Input<string>? GpuKind { get; set; }

        [Input("gpus")]
        public Input<int>? Gpus { get; set; }

        [Input("hostDedicationId")]
        public Input<string>? HostDedicationId { get; set; }

        [Input("kernelArgs")]
        private InputList<string>? _kernelArgs;
        public InputList<string> KernelArgs
        {
            get => _kernelArgs ?? (_kernelArgs = new InputList<string>());
            set => _kernelArgs = value;
        }

        [Input("memoryMb", required: true)]
        public Input<int> MemoryMb { get; set; } = null!;

        public FlyMachineGuestArgs()
        {
        }
        public static new FlyMachineGuestArgs Empty => new FlyMachineGuestArgs();
    }
}