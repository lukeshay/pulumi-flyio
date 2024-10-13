// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlyStopConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlyStopConfigArgs Empty = new FlyStopConfigArgs();

    @Import(name="signal")
    private @Nullable Output<String> signal;

    public Optional<Output<String>> signal() {
        return Optional.ofNullable(this.signal);
    }

    @Import(name="timeout")
    private @Nullable Output<String> timeout;

    public Optional<Output<String>> timeout() {
        return Optional.ofNullable(this.timeout);
    }

    private FlyStopConfigArgs() {}

    private FlyStopConfigArgs(FlyStopConfigArgs $) {
        this.signal = $.signal;
        this.timeout = $.timeout;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlyStopConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlyStopConfigArgs $;

        public Builder() {
            $ = new FlyStopConfigArgs();
        }

        public Builder(FlyStopConfigArgs defaults) {
            $ = new FlyStopConfigArgs(Objects.requireNonNull(defaults));
        }

        public Builder signal(@Nullable Output<String> signal) {
            $.signal = signal;
            return this;
        }

        public Builder signal(String signal) {
            return signal(Output.of(signal));
        }

        public Builder timeout(@Nullable Output<String> timeout) {
            $.timeout = timeout;
            return this;
        }

        public Builder timeout(String timeout) {
            return timeout(Output.of(timeout));
        }

        public FlyStopConfigArgs build() {
            return $;
        }
    }

}
