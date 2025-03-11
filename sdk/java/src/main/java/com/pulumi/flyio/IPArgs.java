// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IPArgs extends com.pulumi.resources.ResourceArgs {

    public static final IPArgs Empty = new IPArgs();

    /**
     * The type of IP address (v4 or v6).
     * 
     */
    @Import(name="addrType", required=true)
    private Output<String> addrType;

    /**
     * @return The type of IP address (v4 or v6).
     * 
     */
    public Output<String> addrType() {
        return this.addrType;
    }

    /**
     * The name of the Fly.io application to allocate the IP address for.
     * 
     */
    @Import(name="app", required=true)
    private Output<String> app;

    /**
     * @return The name of the Fly.io application to allocate the IP address for.
     * 
     */
    public Output<String> app() {
        return this.app;
    }

    /**
     * The network to allocate the IP address in.
     * 
     */
    @Import(name="network")
    private @Nullable Output<String> network;

    /**
     * @return The network to allocate the IP address in.
     * 
     */
    public Optional<Output<String>> network() {
        return Optional.ofNullable(this.network);
    }

    /**
     * The region to allocate the IP address in.
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return The region to allocate the IP address in.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    private IPArgs() {}

    private IPArgs(IPArgs $) {
        this.addrType = $.addrType;
        this.app = $.app;
        this.network = $.network;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IPArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IPArgs $;

        public Builder() {
            $ = new IPArgs();
        }

        public Builder(IPArgs defaults) {
            $ = new IPArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addrType The type of IP address (v4 or v6).
         * 
         * @return builder
         * 
         */
        public Builder addrType(Output<String> addrType) {
            $.addrType = addrType;
            return this;
        }

        /**
         * @param addrType The type of IP address (v4 or v6).
         * 
         * @return builder
         * 
         */
        public Builder addrType(String addrType) {
            return addrType(Output.of(addrType));
        }

        /**
         * @param app The name of the Fly.io application to allocate the IP address for.
         * 
         * @return builder
         * 
         */
        public Builder app(Output<String> app) {
            $.app = app;
            return this;
        }

        /**
         * @param app The name of the Fly.io application to allocate the IP address for.
         * 
         * @return builder
         * 
         */
        public Builder app(String app) {
            return app(Output.of(app));
        }

        /**
         * @param network The network to allocate the IP address in.
         * 
         * @return builder
         * 
         */
        public Builder network(@Nullable Output<String> network) {
            $.network = network;
            return this;
        }

        /**
         * @param network The network to allocate the IP address in.
         * 
         * @return builder
         * 
         */
        public Builder network(String network) {
            return network(Output.of(network));
        }

        /**
         * @param region The region to allocate the IP address in.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region to allocate the IP address in.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        public IPArgs build() {
            if ($.addrType == null) {
                throw new MissingRequiredPropertyException("IPArgs", "addrType");
            }
            if ($.app == null) {
                throw new MissingRequiredPropertyException("IPArgs", "app");
            }
            if ($.region == null) {
                throw new MissingRequiredPropertyException("IPArgs", "region");
            }
            return $;
        }
    }

}
