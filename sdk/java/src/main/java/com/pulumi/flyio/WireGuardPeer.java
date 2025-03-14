// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.flyio.Utilities;
import com.pulumi.flyio.WireGuardPeerArgs;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Fly.io WireGuard peer for private network connectivity.
 * 
 */
@ResourceType(type="flyio:index:WireGuardPeer")
public class WireGuardPeer extends com.pulumi.resources.CustomResource {
    /**
     * The endpoint IP address for the WireGuard peer.
     * 
     */
    @Export(name="endpointIp", refs={String.class}, tree="[0]")
    private Output<String> endpointIp;

    /**
     * @return The endpoint IP address for the WireGuard peer.
     * 
     */
    public Output<String> endpointIp() {
        return this.endpointIp;
    }
    /**
     * The name of the WireGuard peer.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the WireGuard peer.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network the WireGuard peer belongs to.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> network;

    /**
     * @return The network the WireGuard peer belongs to.
     * 
     */
    public Output<Optional<String>> network() {
        return Codegen.optional(this.network);
    }
    /**
     * The organization the WireGuard peer belongs to.
     * 
     */
    @Export(name="org", refs={String.class}, tree="[0]")
    private Output<String> org;

    /**
     * @return The organization the WireGuard peer belongs to.
     * 
     */
    public Output<String> org() {
        return this.org;
    }
    /**
     * The IP address assigned to the WireGuard peer.
     * 
     */
    @Export(name="peerIp", refs={String.class}, tree="[0]")
    private Output<String> peerIp;

    /**
     * @return The IP address assigned to the WireGuard peer.
     * 
     */
    public Output<String> peerIp() {
        return this.peerIp;
    }
    /**
     * The private key of the WireGuard peer.
     * 
     */
    @Export(name="privateKey", refs={String.class}, tree="[0]")
    private Output<String> privateKey;

    /**
     * @return The private key of the WireGuard peer.
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * The public key of the WireGuard peer.
     * 
     */
    @Export(name="publicKey", refs={String.class}, tree="[0]")
    private Output<String> publicKey;

    /**
     * @return The public key of the WireGuard peer.
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }
    /**
     * The region the WireGuard peer is in.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region the WireGuard peer is in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The WireGuard configuration for the peer.
     * 
     */
    @Export(name="wireguardConfig", refs={String.class}, tree="[0]")
    private Output<String> wireguardConfig;

    /**
     * @return The WireGuard configuration for the peer.
     * 
     */
    public Output<String> wireguardConfig() {
        return this.wireguardConfig;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public WireGuardPeer(java.lang.String name) {
        this(name, WireGuardPeerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public WireGuardPeer(java.lang.String name, WireGuardPeerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public WireGuardPeer(java.lang.String name, WireGuardPeerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:WireGuardPeer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private WireGuardPeer(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("flyio:index:WireGuardPeer", name, null, makeResourceOptions(options, id), false);
    }

    private static WireGuardPeerArgs makeArgs(WireGuardPeerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? WireGuardPeerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKey",
                "publicKey",
                "wireguardConfig"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static WireGuardPeer get(java.lang.String name, Output<java.lang.String> id, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new WireGuardPeer(name, id, options);
    }
}
