// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Volume extends pulumi.CustomResource {
    /**
     * Get an existing Volume resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Volume {
        return new Volume(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'flyio:index:Volume';

    /**
     * Returns true if the given object is an instance of Volume.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Volume {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Volume.__pulumiType;
    }

    public readonly appName!: pulumi.Output<string>;
    public /*out*/ readonly attachedAllocId!: pulumi.Output<string | undefined>;
    public /*out*/ readonly attachedMachineId!: pulumi.Output<string | undefined>;
    public readonly autoBackupEnabled!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly blockSize!: pulumi.Output<number | undefined>;
    public /*out*/ readonly blocks!: pulumi.Output<number | undefined>;
    public /*out*/ readonly blocksAvail!: pulumi.Output<number | undefined>;
    public /*out*/ readonly blocksFree!: pulumi.Output<number | undefined>;
    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    public readonly encrypted!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly flyId!: pulumi.Output<string>;
    public readonly fstype!: pulumi.Output<string | undefined>;
    public /*out*/ readonly hostStatus!: pulumi.Output<string | undefined>;
    public /*out*/ readonly input!: pulumi.Output<outputs.VolumeArgs>;
    public readonly name!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string | undefined>;
    public readonly sizeGb!: pulumi.Output<number | undefined>;
    public readonly snapshotRetention!: pulumi.Output<number | undefined>;
    public /*out*/ readonly state!: pulumi.Output<string | undefined>;
    public /*out*/ readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Volume resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appName'");
            }
            resourceInputs["appName"] = args ? args.appName : undefined;
            resourceInputs["autoBackupEnabled"] = args ? args.autoBackupEnabled : undefined;
            resourceInputs["compute"] = args ? args.compute : undefined;
            resourceInputs["computeImage"] = args ? args.computeImage : undefined;
            resourceInputs["encrypted"] = args ? args.encrypted : undefined;
            resourceInputs["fstype"] = args ? args.fstype : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requireUniqueZone"] = args ? args.requireUniqueZone : undefined;
            resourceInputs["sizeGb"] = args ? args.sizeGb : undefined;
            resourceInputs["snapshotId"] = args ? args.snapshotId : undefined;
            resourceInputs["snapshotRetention"] = args ? args.snapshotRetention : undefined;
            resourceInputs["sourceVolumeId"] = args ? args.sourceVolumeId : undefined;
            resourceInputs["attachedAllocId"] = undefined /*out*/;
            resourceInputs["attachedMachineId"] = undefined /*out*/;
            resourceInputs["blockSize"] = undefined /*out*/;
            resourceInputs["blocks"] = undefined /*out*/;
            resourceInputs["blocksAvail"] = undefined /*out*/;
            resourceInputs["blocksFree"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["flyId"] = undefined /*out*/;
            resourceInputs["hostStatus"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        } else {
            resourceInputs["appName"] = undefined /*out*/;
            resourceInputs["attachedAllocId"] = undefined /*out*/;
            resourceInputs["attachedMachineId"] = undefined /*out*/;
            resourceInputs["autoBackupEnabled"] = undefined /*out*/;
            resourceInputs["blockSize"] = undefined /*out*/;
            resourceInputs["blocks"] = undefined /*out*/;
            resourceInputs["blocksAvail"] = undefined /*out*/;
            resourceInputs["blocksFree"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["encrypted"] = undefined /*out*/;
            resourceInputs["flyId"] = undefined /*out*/;
            resourceInputs["fstype"] = undefined /*out*/;
            resourceInputs["hostStatus"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["sizeGb"] = undefined /*out*/;
            resourceInputs["snapshotRetention"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Volume.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Volume resource.
 */
export interface VolumeArgs {
    appName: pulumi.Input<string>;
    autoBackupEnabled?: pulumi.Input<boolean>;
    compute?: pulumi.Input<inputs.flyio.FlyMachineGuestArgs>;
    computeImage?: pulumi.Input<string>;
    encrypted?: pulumi.Input<boolean>;
    fstype?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    requireUniqueZone?: pulumi.Input<boolean>;
    sizeGb?: pulumi.Input<number>;
    snapshotId?: pulumi.Input<string>;
    snapshotRetention?: pulumi.Input<number>;
    sourceVolumeId?: pulumi.Input<string>;
}
