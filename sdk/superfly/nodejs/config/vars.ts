// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("flyio");

/**
 * API key for the Fly.io API.
 */
export declare const token: string;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token") ?? (utilities.getEnv("FLY_API_TOKEN", "FLY_TOKEN", "FLY_API_KEY", "FLY_KEY") || "");
    },
    enumerable: true,
});

