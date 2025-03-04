// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.flyio.flyio.outputs.FlyContainerConfig;
import com.pulumi.flyio.flyio.outputs.FlyDNSConfig;
import com.pulumi.flyio.flyio.outputs.FlyFile;
import com.pulumi.flyio.flyio.outputs.FlyMachineCheck;
import com.pulumi.flyio.flyio.outputs.FlyMachineGuest;
import com.pulumi.flyio.flyio.outputs.FlyMachineInit;
import com.pulumi.flyio.flyio.outputs.FlyMachineMetrics;
import com.pulumi.flyio.flyio.outputs.FlyMachineMount;
import com.pulumi.flyio.flyio.outputs.FlyMachineProcess;
import com.pulumi.flyio.flyio.outputs.FlyMachineRestart;
import com.pulumi.flyio.flyio.outputs.FlyMachineService;
import com.pulumi.flyio.flyio.outputs.FlyStatic;
import com.pulumi.flyio.flyio.outputs.FlyStopConfig;
import com.pulumi.flyio.flyio.outputs.FlyVolumeConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FlyMachineConfig {
    private @Nullable Boolean autoDestroy;
    private @Nullable Map<String,FlyMachineCheck> checks;
    private @Nullable List<FlyContainerConfig> containers;
    private @Nullable FlyDNSConfig dns;
    private @Nullable Map<String,String> env;
    private @Nullable List<FlyFile> files;
    private @Nullable FlyMachineGuest guest;
    private String image;
    private @Nullable FlyMachineInit init;
    private @Nullable Map<String,String> metadata;
    private @Nullable FlyMachineMetrics metrics;
    private @Nullable List<FlyMachineMount> mounts;
    private @Nullable List<FlyMachineProcess> processes;
    private @Nullable FlyMachineRestart restart;
    private @Nullable String schedule;
    private @Nullable List<FlyMachineService> services;
    private @Nullable List<String> standbys;
    private @Nullable List<FlyStatic> statics;
    private @Nullable FlyStopConfig stopConfig;
    private @Nullable List<FlyVolumeConfig> volumes;

    private FlyMachineConfig() {}
    public Optional<Boolean> autoDestroy() {
        return Optional.ofNullable(this.autoDestroy);
    }
    public Map<String,FlyMachineCheck> checks() {
        return this.checks == null ? Map.of() : this.checks;
    }
    public List<FlyContainerConfig> containers() {
        return this.containers == null ? List.of() : this.containers;
    }
    public Optional<FlyDNSConfig> dns() {
        return Optional.ofNullable(this.dns);
    }
    public Map<String,String> env() {
        return this.env == null ? Map.of() : this.env;
    }
    public List<FlyFile> files() {
        return this.files == null ? List.of() : this.files;
    }
    public Optional<FlyMachineGuest> guest() {
        return Optional.ofNullable(this.guest);
    }
    public String image() {
        return this.image;
    }
    public Optional<FlyMachineInit> init() {
        return Optional.ofNullable(this.init);
    }
    public Map<String,String> metadata() {
        return this.metadata == null ? Map.of() : this.metadata;
    }
    public Optional<FlyMachineMetrics> metrics() {
        return Optional.ofNullable(this.metrics);
    }
    public List<FlyMachineMount> mounts() {
        return this.mounts == null ? List.of() : this.mounts;
    }
    public List<FlyMachineProcess> processes() {
        return this.processes == null ? List.of() : this.processes;
    }
    public Optional<FlyMachineRestart> restart() {
        return Optional.ofNullable(this.restart);
    }
    public Optional<String> schedule() {
        return Optional.ofNullable(this.schedule);
    }
    public List<FlyMachineService> services() {
        return this.services == null ? List.of() : this.services;
    }
    public List<String> standbys() {
        return this.standbys == null ? List.of() : this.standbys;
    }
    public List<FlyStatic> statics() {
        return this.statics == null ? List.of() : this.statics;
    }
    public Optional<FlyStopConfig> stopConfig() {
        return Optional.ofNullable(this.stopConfig);
    }
    public List<FlyVolumeConfig> volumes() {
        return this.volumes == null ? List.of() : this.volumes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FlyMachineConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean autoDestroy;
        private @Nullable Map<String,FlyMachineCheck> checks;
        private @Nullable List<FlyContainerConfig> containers;
        private @Nullable FlyDNSConfig dns;
        private @Nullable Map<String,String> env;
        private @Nullable List<FlyFile> files;
        private @Nullable FlyMachineGuest guest;
        private String image;
        private @Nullable FlyMachineInit init;
        private @Nullable Map<String,String> metadata;
        private @Nullable FlyMachineMetrics metrics;
        private @Nullable List<FlyMachineMount> mounts;
        private @Nullable List<FlyMachineProcess> processes;
        private @Nullable FlyMachineRestart restart;
        private @Nullable String schedule;
        private @Nullable List<FlyMachineService> services;
        private @Nullable List<String> standbys;
        private @Nullable List<FlyStatic> statics;
        private @Nullable FlyStopConfig stopConfig;
        private @Nullable List<FlyVolumeConfig> volumes;
        public Builder() {}
        public Builder(FlyMachineConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoDestroy = defaults.autoDestroy;
    	      this.checks = defaults.checks;
    	      this.containers = defaults.containers;
    	      this.dns = defaults.dns;
    	      this.env = defaults.env;
    	      this.files = defaults.files;
    	      this.guest = defaults.guest;
    	      this.image = defaults.image;
    	      this.init = defaults.init;
    	      this.metadata = defaults.metadata;
    	      this.metrics = defaults.metrics;
    	      this.mounts = defaults.mounts;
    	      this.processes = defaults.processes;
    	      this.restart = defaults.restart;
    	      this.schedule = defaults.schedule;
    	      this.services = defaults.services;
    	      this.standbys = defaults.standbys;
    	      this.statics = defaults.statics;
    	      this.stopConfig = defaults.stopConfig;
    	      this.volumes = defaults.volumes;
        }

        @CustomType.Setter
        public Builder autoDestroy(@Nullable Boolean autoDestroy) {

            this.autoDestroy = autoDestroy;
            return this;
        }
        @CustomType.Setter
        public Builder checks(@Nullable Map<String,FlyMachineCheck> checks) {

            this.checks = checks;
            return this;
        }
        @CustomType.Setter
        public Builder containers(@Nullable List<FlyContainerConfig> containers) {

            this.containers = containers;
            return this;
        }
        public Builder containers(FlyContainerConfig... containers) {
            return containers(List.of(containers));
        }
        @CustomType.Setter
        public Builder dns(@Nullable FlyDNSConfig dns) {

            this.dns = dns;
            return this;
        }
        @CustomType.Setter
        public Builder env(@Nullable Map<String,String> env) {

            this.env = env;
            return this;
        }
        @CustomType.Setter
        public Builder files(@Nullable List<FlyFile> files) {

            this.files = files;
            return this;
        }
        public Builder files(FlyFile... files) {
            return files(List.of(files));
        }
        @CustomType.Setter
        public Builder guest(@Nullable FlyMachineGuest guest) {

            this.guest = guest;
            return this;
        }
        @CustomType.Setter
        public Builder image(String image) {
            if (image == null) {
              throw new MissingRequiredPropertyException("FlyMachineConfig", "image");
            }
            this.image = image;
            return this;
        }
        @CustomType.Setter
        public Builder init(@Nullable FlyMachineInit init) {

            this.init = init;
            return this;
        }
        @CustomType.Setter
        public Builder metadata(@Nullable Map<String,String> metadata) {

            this.metadata = metadata;
            return this;
        }
        @CustomType.Setter
        public Builder metrics(@Nullable FlyMachineMetrics metrics) {

            this.metrics = metrics;
            return this;
        }
        @CustomType.Setter
        public Builder mounts(@Nullable List<FlyMachineMount> mounts) {

            this.mounts = mounts;
            return this;
        }
        public Builder mounts(FlyMachineMount... mounts) {
            return mounts(List.of(mounts));
        }
        @CustomType.Setter
        public Builder processes(@Nullable List<FlyMachineProcess> processes) {

            this.processes = processes;
            return this;
        }
        public Builder processes(FlyMachineProcess... processes) {
            return processes(List.of(processes));
        }
        @CustomType.Setter
        public Builder restart(@Nullable FlyMachineRestart restart) {

            this.restart = restart;
            return this;
        }
        @CustomType.Setter
        public Builder schedule(@Nullable String schedule) {

            this.schedule = schedule;
            return this;
        }
        @CustomType.Setter
        public Builder services(@Nullable List<FlyMachineService> services) {

            this.services = services;
            return this;
        }
        public Builder services(FlyMachineService... services) {
            return services(List.of(services));
        }
        @CustomType.Setter
        public Builder standbys(@Nullable List<String> standbys) {

            this.standbys = standbys;
            return this;
        }
        public Builder standbys(String... standbys) {
            return standbys(List.of(standbys));
        }
        @CustomType.Setter
        public Builder statics(@Nullable List<FlyStatic> statics) {

            this.statics = statics;
            return this;
        }
        public Builder statics(FlyStatic... statics) {
            return statics(List.of(statics));
        }
        @CustomType.Setter
        public Builder stopConfig(@Nullable FlyStopConfig stopConfig) {

            this.stopConfig = stopConfig;
            return this;
        }
        @CustomType.Setter
        public Builder volumes(@Nullable List<FlyVolumeConfig> volumes) {

            this.volumes = volumes;
            return this;
        }
        public Builder volumes(FlyVolumeConfig... volumes) {
            return volumes(List.of(volumes));
        }
        public FlyMachineConfig build() {
            final var _resultValue = new FlyMachineConfig();
            _resultValue.autoDestroy = autoDestroy;
            _resultValue.checks = checks;
            _resultValue.containers = containers;
            _resultValue.dns = dns;
            _resultValue.env = env;
            _resultValue.files = files;
            _resultValue.guest = guest;
            _resultValue.image = image;
            _resultValue.init = init;
            _resultValue.metadata = metadata;
            _resultValue.metrics = metrics;
            _resultValue.mounts = mounts;
            _resultValue.processes = processes;
            _resultValue.restart = restart;
            _resultValue.schedule = schedule;
            _resultValue.services = services;
            _resultValue.standbys = standbys;
            _resultValue.statics = statics;
            _resultValue.stopConfig = stopConfig;
            _resultValue.volumes = volumes;
            return _resultValue;
        }
    }
}
