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


public final class FlyMachineCheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlyMachineCheckArgs Empty = new FlyMachineCheckArgs();

    @Import(name="gracePeriod")
    private @Nullable Output<String> gracePeriod;

    public Optional<Output<String>> gracePeriod() {
        return Optional.ofNullable(this.gracePeriod);
    }

    @Import(name="headers")
    private @Nullable Output<List<FlyMachineHTTPHeaderArgs>> headers;

    public Optional<Output<List<FlyMachineHTTPHeaderArgs>>> headers() {
        return Optional.ofNullable(this.headers);
    }

    @Import(name="interval")
    private @Nullable Output<String> interval;

    public Optional<Output<String>> interval() {
        return Optional.ofNullable(this.interval);
    }

    @Import(name="kind")
    private @Nullable Output<String> kind;

    public Optional<Output<String>> kind() {
        return Optional.ofNullable(this.kind);
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

    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    @Import(name="timeout")
    private @Nullable Output<String> timeout;

    public Optional<Output<String>> timeout() {
        return Optional.ofNullable(this.timeout);
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

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private FlyMachineCheckArgs() {}

    private FlyMachineCheckArgs(FlyMachineCheckArgs $) {
        this.gracePeriod = $.gracePeriod;
        this.headers = $.headers;
        this.interval = $.interval;
        this.kind = $.kind;
        this.method = $.method;
        this.path = $.path;
        this.port = $.port;
        this.protocol = $.protocol;
        this.timeout = $.timeout;
        this.tlsServerName = $.tlsServerName;
        this.tlsSkipVerify = $.tlsSkipVerify;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlyMachineCheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlyMachineCheckArgs $;

        public Builder() {
            $ = new FlyMachineCheckArgs();
        }

        public Builder(FlyMachineCheckArgs defaults) {
            $ = new FlyMachineCheckArgs(Objects.requireNonNull(defaults));
        }

        public Builder gracePeriod(@Nullable Output<String> gracePeriod) {
            $.gracePeriod = gracePeriod;
            return this;
        }

        public Builder gracePeriod(String gracePeriod) {
            return gracePeriod(Output.of(gracePeriod));
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

        public Builder interval(@Nullable Output<String> interval) {
            $.interval = interval;
            return this;
        }

        public Builder interval(String interval) {
            return interval(Output.of(interval));
        }

        public Builder kind(@Nullable Output<String> kind) {
            $.kind = kind;
            return this;
        }

        public Builder kind(String kind) {
            return kind(Output.of(kind));
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

        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        public Builder timeout(@Nullable Output<String> timeout) {
            $.timeout = timeout;
            return this;
        }

        public Builder timeout(String timeout) {
            return timeout(Output.of(timeout));
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

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public FlyMachineCheckArgs build() {
            return $;
        }
    }

}
