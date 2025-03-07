// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.flyio.flyio.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.flyio.flyio.inputs.FlyEnvFromArgs;
import com.pulumi.flyio.flyio.inputs.FlyMachineSecretArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FlyMachineProcessArgs extends com.pulumi.resources.ResourceArgs {

    public static final FlyMachineProcessArgs Empty = new FlyMachineProcessArgs();

    @Import(name="cmd")
    private @Nullable Output<List<String>> cmd;

    public Optional<Output<List<String>>> cmd() {
        return Optional.ofNullable(this.cmd);
    }

    @Import(name="entrypoint")
    private @Nullable Output<List<String>> entrypoint;

    public Optional<Output<List<String>>> entrypoint() {
        return Optional.ofNullable(this.entrypoint);
    }

    @Import(name="env")
    private @Nullable Output<Map<String,String>> env;

    public Optional<Output<Map<String,String>>> env() {
        return Optional.ofNullable(this.env);
    }

    @Import(name="envFrom")
    private @Nullable Output<List<FlyEnvFromArgs>> envFrom;

    public Optional<Output<List<FlyEnvFromArgs>>> envFrom() {
        return Optional.ofNullable(this.envFrom);
    }

    @Import(name="exec")
    private @Nullable Output<List<String>> exec;

    public Optional<Output<List<String>>> exec() {
        return Optional.ofNullable(this.exec);
    }

    @Import(name="ignoreAppSecrets")
    private @Nullable Output<Boolean> ignoreAppSecrets;

    public Optional<Output<Boolean>> ignoreAppSecrets() {
        return Optional.ofNullable(this.ignoreAppSecrets);
    }

    @Import(name="secrets")
    private @Nullable Output<List<FlyMachineSecretArgs>> secrets;

    public Optional<Output<List<FlyMachineSecretArgs>>> secrets() {
        return Optional.ofNullable(this.secrets);
    }

    @Import(name="user")
    private @Nullable Output<String> user;

    public Optional<Output<String>> user() {
        return Optional.ofNullable(this.user);
    }

    private FlyMachineProcessArgs() {}

    private FlyMachineProcessArgs(FlyMachineProcessArgs $) {
        this.cmd = $.cmd;
        this.entrypoint = $.entrypoint;
        this.env = $.env;
        this.envFrom = $.envFrom;
        this.exec = $.exec;
        this.ignoreAppSecrets = $.ignoreAppSecrets;
        this.secrets = $.secrets;
        this.user = $.user;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FlyMachineProcessArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FlyMachineProcessArgs $;

        public Builder() {
            $ = new FlyMachineProcessArgs();
        }

        public Builder(FlyMachineProcessArgs defaults) {
            $ = new FlyMachineProcessArgs(Objects.requireNonNull(defaults));
        }

        public Builder cmd(@Nullable Output<List<String>> cmd) {
            $.cmd = cmd;
            return this;
        }

        public Builder cmd(List<String> cmd) {
            return cmd(Output.of(cmd));
        }

        public Builder cmd(String... cmd) {
            return cmd(List.of(cmd));
        }

        public Builder entrypoint(@Nullable Output<List<String>> entrypoint) {
            $.entrypoint = entrypoint;
            return this;
        }

        public Builder entrypoint(List<String> entrypoint) {
            return entrypoint(Output.of(entrypoint));
        }

        public Builder entrypoint(String... entrypoint) {
            return entrypoint(List.of(entrypoint));
        }

        public Builder env(@Nullable Output<Map<String,String>> env) {
            $.env = env;
            return this;
        }

        public Builder env(Map<String,String> env) {
            return env(Output.of(env));
        }

        public Builder envFrom(@Nullable Output<List<FlyEnvFromArgs>> envFrom) {
            $.envFrom = envFrom;
            return this;
        }

        public Builder envFrom(List<FlyEnvFromArgs> envFrom) {
            return envFrom(Output.of(envFrom));
        }

        public Builder envFrom(FlyEnvFromArgs... envFrom) {
            return envFrom(List.of(envFrom));
        }

        public Builder exec(@Nullable Output<List<String>> exec) {
            $.exec = exec;
            return this;
        }

        public Builder exec(List<String> exec) {
            return exec(Output.of(exec));
        }

        public Builder exec(String... exec) {
            return exec(List.of(exec));
        }

        public Builder ignoreAppSecrets(@Nullable Output<Boolean> ignoreAppSecrets) {
            $.ignoreAppSecrets = ignoreAppSecrets;
            return this;
        }

        public Builder ignoreAppSecrets(Boolean ignoreAppSecrets) {
            return ignoreAppSecrets(Output.of(ignoreAppSecrets));
        }

        public Builder secrets(@Nullable Output<List<FlyMachineSecretArgs>> secrets) {
            $.secrets = secrets;
            return this;
        }

        public Builder secrets(List<FlyMachineSecretArgs> secrets) {
            return secrets(Output.of(secrets));
        }

        public Builder secrets(FlyMachineSecretArgs... secrets) {
            return secrets(List.of(secrets));
        }

        public Builder user(@Nullable Output<String> user) {
            $.user = user;
            return this;
        }

        public Builder user(String user) {
            return user(Output.of(user));
        }

        public FlyMachineProcessArgs build() {
            return $;
        }
    }

}
