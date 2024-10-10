// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Flyio
{
    [FlyioResourceType("flyio:index:Volume")]
    public partial class Volume : global::Pulumi.CustomResource
    {
        [Output("appName")]
        public Output<string> AppName { get; private set; } = null!;

        [Output("attachedAllocId")]
        public Output<string?> AttachedAllocId { get; private set; } = null!;

        [Output("attachedMachineId")]
        public Output<string?> AttachedMachineId { get; private set; } = null!;

        [Output("autoBackupEnabled")]
        public Output<bool?> AutoBackupEnabled { get; private set; } = null!;

        [Output("blockSize")]
        public Output<int?> BlockSize { get; private set; } = null!;

        [Output("blocks")]
        public Output<int?> Blocks { get; private set; } = null!;

        [Output("blocksAvail")]
        public Output<int?> BlocksAvail { get; private set; } = null!;

        [Output("blocksFree")]
        public Output<int?> BlocksFree { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string?> CreatedAt { get; private set; } = null!;

        [Output("encrypted")]
        public Output<bool?> Encrypted { get; private set; } = null!;

        [Output("flyId")]
        public Output<string> FlyId { get; private set; } = null!;

        [Output("fstype")]
        public Output<string?> Fstype { get; private set; } = null!;

        [Output("hostStatus")]
        public Output<string?> HostStatus { get; private set; } = null!;

        [Output("input")]
        public Output<Outputs.VolumeArgs> Input { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        [Output("sizeGb")]
        public Output<int?> SizeGb { get; private set; } = null!;

        [Output("snapshotRetention")]
        public Output<int?> SnapshotRetention { get; private set; } = null!;

        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        [Output("zone")]
        public Output<string?> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs args, CustomResourceOptions? options = null)
            : base("flyio:index:Volume", name, args ?? new VolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("flyio:index:Volume", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/lukeshay/pulumi-flyio/releases/download/${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, options);
        }
    }

    public sealed class VolumeArgs : global::Pulumi.ResourceArgs
    {
        [Input("appName", required: true)]
        public Input<string> AppName { get; set; } = null!;

        [Input("autoBackupEnabled")]
        public Input<bool>? AutoBackupEnabled { get; set; }

        [Input("compute")]
        public Input<Pulumi.Flyio.Flyio.Inputs.FlyMachineGuestArgs>? Compute { get; set; }

        [Input("computeImage")]
        public Input<string>? ComputeImage { get; set; }

        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("fstype")]
        public Input<string>? Fstype { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("requireUniqueZone")]
        public Input<bool>? RequireUniqueZone { get; set; }

        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("snapshotRetention")]
        public Input<int>? SnapshotRetention { get; set; }

        [Input("sourceVolumeId")]
        public Input<string>? SourceVolumeId { get; set; }

        public VolumeArgs()
        {
        }
        public static new VolumeArgs Empty => new VolumeArgs();
    }
}
