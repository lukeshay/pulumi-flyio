// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlyProxyProtoOptionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlyProxyProtoOptionsArgs Empty = new FlyProxyProtoOptionsArgs();

    @Import(name="version")
    private @Nullable Output<String> version;

    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private FlyProxyProtoOptionsArgs() {}

    private FlyProxyProtoOptionsArgs(FlyProxyProtoOptionsArgs $) {
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlyProxyProtoOptionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlyProxyProtoOptionsArgs $;

        public Builder() {
            $ = new FlyProxyProtoOptionsArgs();
        }

        public Builder(FlyProxyProtoOptionsArgs defaults) {
            $ = new FlyProxyProtoOptionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        public Builder version(String version) {
            return version(Output.of(version));
        }

        public FlyProxyProtoOptionsArgs build() {
            return $;
        }
    }

}
