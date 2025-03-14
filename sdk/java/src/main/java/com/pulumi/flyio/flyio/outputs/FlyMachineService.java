// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.flyio.flyio.outputs.FlyMachineCheck;
import com.pulumi.flyio.flyio.outputs.FlyMachinePort;
import com.pulumi.flyio.flyio.outputs.FlyMachineServiceConcurrency;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FlyMachineService {
    private @Nullable Boolean autostart;
    private @Nullable String autostop;
    private @Nullable List<FlyMachineCheck> checks;
    private @Nullable FlyMachineServiceConcurrency concurrency;
    private @Nullable String forceInstanceDescription;
    private @Nullable String forceInstanceKey;
    private @Nullable Integer internalPort;
    private @Nullable Integer minMachinesRunning;
    private @Nullable List<FlyMachinePort> ports;
    private @Nullable String protocol;

    private FlyMachineService() {}
    public Optional<Boolean> autostart() {
        return Optional.ofNullable(this.autostart);
    }
    public Optional<String> autostop() {
        return Optional.ofNullable(this.autostop);
    }
    public List<FlyMachineCheck> checks() {
        return this.checks == null ? List.of() : this.checks;
    }
    public Optional<FlyMachineServiceConcurrency> concurrency() {
        return Optional.ofNullable(this.concurrency);
    }
    public Optional<String> forceInstanceDescription() {
        return Optional.ofNullable(this.forceInstanceDescription);
    }
    public Optional<String> forceInstanceKey() {
        return Optional.ofNullable(this.forceInstanceKey);
    }
    public Optional<Integer> internalPort() {
        return Optional.ofNullable(this.internalPort);
    }
    public Optional<Integer> minMachinesRunning() {
        return Optional.ofNullable(this.minMachinesRunning);
    }
    public List<FlyMachinePort> ports() {
        return this.ports == null ? List.of() : this.ports;
    }
    public Optional<String> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FlyMachineService defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean autostart;
        private @Nullable String autostop;
        private @Nullable List<FlyMachineCheck> checks;
        private @Nullable FlyMachineServiceConcurrency concurrency;
        private @Nullable String forceInstanceDescription;
        private @Nullable String forceInstanceKey;
        private @Nullable Integer internalPort;
        private @Nullable Integer minMachinesRunning;
        private @Nullable List<FlyMachinePort> ports;
        private @Nullable String protocol;
        public Builder() {}
        public Builder(FlyMachineService defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autostart = defaults.autostart;
    	      this.autostop = defaults.autostop;
    	      this.checks = defaults.checks;
    	      this.concurrency = defaults.concurrency;
    	      this.forceInstanceDescription = defaults.forceInstanceDescription;
    	      this.forceInstanceKey = defaults.forceInstanceKey;
    	      this.internalPort = defaults.internalPort;
    	      this.minMachinesRunning = defaults.minMachinesRunning;
    	      this.ports = defaults.ports;
    	      this.protocol = defaults.protocol;
        }

        @CustomType.Setter
        public Builder autostart(@Nullable Boolean autostart) {

            this.autostart = autostart;
            return this;
        }
        @CustomType.Setter
        public Builder autostop(@Nullable String autostop) {

            this.autostop = autostop;
            return this;
        }
        @CustomType.Setter
        public Builder checks(@Nullable List<FlyMachineCheck> checks) {

            this.checks = checks;
            return this;
        }
        public Builder checks(FlyMachineCheck... checks) {
            return checks(List.of(checks));
        }
        @CustomType.Setter
        public Builder concurrency(@Nullable FlyMachineServiceConcurrency concurrency) {

            this.concurrency = concurrency;
            return this;
        }
        @CustomType.Setter
        public Builder forceInstanceDescription(@Nullable String forceInstanceDescription) {

            this.forceInstanceDescription = forceInstanceDescription;
            return this;
        }
        @CustomType.Setter
        public Builder forceInstanceKey(@Nullable String forceInstanceKey) {

            this.forceInstanceKey = forceInstanceKey;
            return this;
        }
        @CustomType.Setter
        public Builder internalPort(@Nullable Integer internalPort) {

            this.internalPort = internalPort;
            return this;
        }
        @CustomType.Setter
        public Builder minMachinesRunning(@Nullable Integer minMachinesRunning) {

            this.minMachinesRunning = minMachinesRunning;
            return this;
        }
        @CustomType.Setter
        public Builder ports(@Nullable List<FlyMachinePort> ports) {

            this.ports = ports;
            return this;
        }
        public Builder ports(FlyMachinePort... ports) {
            return ports(List.of(ports));
        }
        @CustomType.Setter
        public Builder protocol(@Nullable String protocol) {

            this.protocol = protocol;
            return this;
        }
        public FlyMachineService build() {
            final var _resultValue = new FlyMachineService();
            _resultValue.autostart = autostart;
            _resultValue.autostop = autostop;
            _resultValue.checks = checks;
            _resultValue.concurrency = concurrency;
            _resultValue.forceInstanceDescription = forceInstanceDescription;
            _resultValue.forceInstanceKey = forceInstanceKey;
            _resultValue.internalPort = internalPort;
            _resultValue.minMachinesRunning = minMachinesRunning;
            _resultValue.ports = ports;
            _resultValue.protocol = protocol;
            return _resultValue;
        }
    }
}
