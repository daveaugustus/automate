package main

import (
	"io/ioutil"
	"os"

	"github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/components/automate-cli/pkg/status"
	"github.com/chef/automate/components/automate-deployment/pkg/toml"
)

func runInitConfigExistingNodeHACmd() error {
	initConfigHAPath := initConfigHAFlags.path
	if _, err := os.Stat(initConfigHAPath); err == nil {
		writer.Printf("Skipping config initialization. Config already exists at %s\n", initConfigHAPath)
		return nil
	}

	cfg, err := deployment.GenerateInitHAConfig(
		deployment.InitialSecretsKeyFile(initConfigHAFlags.SecretsKeyFile),
		deployment.InitialSecretsStoreFile(initConfigHAFlags.SecretsStoreFile),
		deployment.InitialArchitecture(initConfigHAFlags.Architecture),
		deployment.InitialWorkspacePath(initConfigHAFlags.WorkspacePath),
		deployment.InitialSshUser(initConfigHAFlags.SshUser),
		deployment.InitialSshKeyFile(initConfigHAFlags.SshKeyFile),
		deployment.InitialBackupMount(initConfigHAFlags.BackupMount),
		deployment.InitialAutomateInstanceCount(initConfigHAFlags.AutomateInstanceCount),
		deployment.InitialAutomateConfigFile(initConfigHAFlags.AutomateConfigFile),
		deployment.InitialChefServerInstanceCount(initConfigHAFlags.ChefServerInstanceCount),
		deployment.InitialElasticSearchInstanceCount(initConfigHAFlags.ElasticSearchInstanceCount),
		deployment.InitialPostgresqlInstanceCount(initConfigHAFlags.PostgresqlInstanceCount),
	)

	if err != nil {
		return status.Wrap(err, status.ConfigError, "Generating initial configuration failed")
	}

	t, err := cfg.RenderExistingNodesHaSettings()
	if err != nil {
		return status.Wrap(err, status.ConfigError, "Rendering initial configuration failed")
	}

	// Make sure our user facing config is a valid AutomateConfig
	ac := deployment.NewAutomateConfig()
	err = toml.Unmarshal([]byte(t), ac)
	if err != nil {
		return status.Wrap(err, status.MarshalError, "Marshaling initial configuration failed")
	}

	err = ioutil.WriteFile(initConfigHAPath, []byte(t), 0600)
	if err != nil {
		return status.Wrap(err, status.FileAccessError, "Writing initial configuration failed")
	}
	writer.Println("file generated " + initConfigHAPath)
	return nil
}