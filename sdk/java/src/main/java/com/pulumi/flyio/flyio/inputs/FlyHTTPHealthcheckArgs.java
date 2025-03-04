// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.flyio.flyio.inputs.FlyMachineHTTPHeaderArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlyHTTPHealthcheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlyHTTPHealthcheckArgs Empty = new FlyHTTPHealthcheckArgs();

    @Import(name="headers")
    private @Nullable Output<List<FlyMachineHTTPHeaderArgs>> headers;

    public Optional<Output<List<FlyMachineHTTPHeaderArgs>>> headers() {
        return Optional.ofNullable(this.headers);
    }

    @Import(name="method")
    private @Nullable Output<String> method;

    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
    }

    @Import(name="path")
    private @Nullable Output<String> path;

    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    @Import(name="port")
    private @Nullable Output<Integer> port;

    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    @Import(name="scheme")
    private @Nullable Output<String> scheme;

    public Optional<Output<String>> scheme() {
        return Optional.ofNullable(this.scheme);
    }

    @Import(name="tlsServerName")
    private @Nullable Output<String> tlsServerName;

    public Optional<Output<String>> tlsServerName() {
        return Optional.ofNullable(this.tlsServerName);
    }

    @Import(name="tlsSkipVerify")
    private @Nullable Output<Boolean> tlsSkipVerify;

    public Optional<Output<Boolean>> tlsSkipVerify() {
        return Optional.ofNullable(this.tlsSkipVerify);
    }

    private FlyHTTPHealthcheckArgs() {}

    private FlyHTTPHealthcheckArgs(FlyHTTPHealthcheckArgs $) {
        this.headers = $.headers;
        this.method = $.method;
        this.path = $.path;
        this.port = $.port;
        this.scheme = $.scheme;
        this.tlsServerName = $.tlsServerName;
        this.tlsSkipVerify = $.tlsSkipVerify;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlyHTTPHealthcheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlyHTTPHealthcheckArgs $;

        public Builder() {
            $ = new FlyHTTPHealthcheckArgs();
        }

        public Builder(FlyHTTPHealthcheckArgs defaults) {
            $ = new FlyHTTPHealthcheckArgs(Objects.requireNonNull(defaults));
        }

        public Builder headers(@Nullable Output<List<FlyMachineHTTPHeaderArgs>> headers) {
            $.headers = headers;
            return this;
        }

        public Builder headers(List<FlyMachineHTTPHeaderArgs> headers) {
            return headers(Output.of(headers));
        }

        public Builder headers(FlyMachineHTTPHeaderArgs... headers) {
            return headers(List.of(headers));
        }

        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        public Builder method(String method) {
            return method(Output.of(method));
        }

        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        public Builder path(String path) {
            return path(Output.of(path));
        }

        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public Builder scheme(@Nullable Output<String> scheme) {
            $.scheme = scheme;
            return this;
        }

        public Builder scheme(String scheme) {
            return scheme(Output.of(scheme));
        }

        public Builder tlsServerName(@Nullable Output<String> tlsServerName) {
            $.tlsServerName = tlsServerName;
            return this;
        }

        public Builder tlsServerName(String tlsServerName) {
            return tlsServerName(Output.of(tlsServerName));
        }

        public Builder tlsSkipVerify(@Nullable Output<Boolean> tlsSkipVerify) {
            $.tlsSkipVerify = tlsSkipVerify;
            return this;
        }

        public Builder tlsSkipVerify(Boolean tlsSkipVerify) {
            return tlsSkipVerify(Output.of(tlsSkipVerify));
        }

        public FlyHTTPHealthcheckArgs build() {
            return $;
        }
    }

}
