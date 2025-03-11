// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * A Fly.io machine represents a VM instance that runs your application.
 */
export class Machine extends pulumi.CustomResource {
    /**
     * Get an existing Machine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Machine {
        return new Machine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'flyio:index:Machine';

    /**
     * Returns true if the given object is an instance of Machine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Machine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Machine.__pulumiType;
    }

    /**
     * The Fly.io application the machine belongs to.
     */
    public readonly app!: pulumi.Output<string>;
    public /*out*/ readonly checks!: pulumi.Output<outputs.flyio.CheckStatus[] | undefined>;
    public readonly config!: pulumi.Output<outputs.flyio.FlyMachineConfig>;
    public /*out*/ readonly createdAt!: pulumi.Output<string | undefined>;
    /**
     * The deployment strategy used for the machine.
     */
    public readonly deploymentStrategy!: pulumi.Output<string | undefined>;
    public /*out*/ readonly events!: pulumi.Output<outputs.flyio.MachineEvent[] | undefined>;
    public /*out*/ readonly flyId!: pulumi.Output<string>;
    public /*out*/ readonly hostStatus!: pulumi.Output<string | undefined>;
    public /*out*/ readonly imageRef!: pulumi.Output<outputs.flyio.ImageRef | undefined>;
    public /*out*/ readonly incompleteConfig!: pulumi.Output<outputs.flyio.FlyMachineConfig | undefined>;
    /**
     * The input arguments used to create the machine.
     */
    public /*out*/ readonly input!: pulumi.Output<outputs.MachineArgs>;
    public /*out*/ readonly instanceId!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly nonce!: pulumi.Output<string | undefined>;
    public /*out*/ readonly privateIp!: pulumi.Output<string | undefined>;
    public readonly region!: pulumi.Output<string | undefined>;
    public /*out*/ readonly state!: pulumi.Output<string>;
    public /*out*/ readonly updatedAt!: pulumi.Output<string | undefined>;

    /**
     * Create a Machine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MachineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.app === undefined) && !opts.urn) {
                throw new Error("Missing required property 'app'");
            }
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            resourceInputs["app"] = args ? args.app : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["deploymentStrategy"] = args ? args.deploymentStrategy : undefined;
            resourceInputs["leaseTtl"] = args ? args.leaseTtl : undefined;
            resourceInputs["lsvd"] = args ? args.lsvd : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["skipLaunch"] = args ? args.skipLaunch : undefined;
            resourceInputs["skipServiceRegistration"] = args ? args.skipServiceRegistration : undefined;
            resourceInputs["checks"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["events"] = undefined /*out*/;
            resourceInputs["flyId"] = undefined /*out*/;
            resourceInputs["hostStatus"] = undefined /*out*/;
            resourceInputs["imageRef"] = undefined /*out*/;
            resourceInputs["incompleteConfig"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["nonce"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["app"] = undefined /*out*/;
            resourceInputs["checks"] = undefined /*out*/;
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["deploymentStrategy"] = undefined /*out*/;
            resourceInputs["events"] = undefined /*out*/;
            resourceInputs["flyId"] = undefined /*out*/;
            resourceInputs["hostStatus"] = undefined /*out*/;
            resourceInputs["imageRef"] = undefined /*out*/;
            resourceInputs["incompleteConfig"] = undefined /*out*/;
            resourceInputs["input"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nonce"] = undefined /*out*/;
            resourceInputs["privateIp"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Machine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Machine resource.
 */
export interface MachineArgs {
    /**
     * The Fly.io application to deploy the machine to.
     */
    app: pulumi.Input<string>;
    config: pulumi.Input<inputs.flyio.FlyMachineConfigArgs>;
    /**
     * The deployment strategy for the machine.
     */
    deploymentStrategy?: pulumi.Input<string>;
    leaseTtl?: pulumi.Input<number>;
    lsvd?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
    skipLaunch?: pulumi.Input<boolean>;
    skipServiceRegistration?: pulumi.Input<boolean>;
}
