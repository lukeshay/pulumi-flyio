// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio.Flyio.Inputs
{

    public sealed class FlyMachineServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("autostart")]
        public Input<bool>? Autostart { get; set; }

        [Input("autostop")]
        public Input<bool>? Autostop { get; set; }

        [Input("checks")]
        private InputList<Inputs.FlyMachineCheckArgs>? _checks;
        public InputList<Inputs.FlyMachineCheckArgs> Checks
        {
            get => _checks ?? (_checks = new InputList<Inputs.FlyMachineCheckArgs>());
            set => _checks = value;
        }

        [Input("concurrency")]
        public Input<Inputs.FlyMachineServiceConcurrencyArgs>? Concurrency { get; set; }

        [Input("forceInstanceDescription")]
        public Input<string>? ForceInstanceDescription { get; set; }

        [Input("forceInstanceKey")]
        public Input<string>? ForceInstanceKey { get; set; }

        [Input("internalPort")]
        public Input<int>? InternalPort { get; set; }

        [Input("minMachinesRunning")]
        public Input<int>? MinMachinesRunning { get; set; }

        [Input("ports")]
        private InputList<Inputs.FlyMachinePortArgs>? _ports;
        public InputList<Inputs.FlyMachinePortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.FlyMachinePortArgs>());
            set => _ports = value;
        }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public FlyMachineServiceArgs()
        {
        }
        public static new FlyMachineServiceArgs Empty => new FlyMachineServiceArgs();
    }
}
