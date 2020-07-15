package config

import (
    "fmt"
    "io/ioutil"

    "github.com/imdario/mergo"
    "github.com/urfave/cli"
    "gopkg.in/yaml.v2"
)


// Rocket Pool config
type RocketPoolConfig struct {
    Rocketpool struct {
        StorageAddress string           `yaml:"storageAddress,omitempty"`
    }                                   `yaml:"rocketpool,omitempty"`
    Smartnode struct {
        PasswordPath string             `yaml:"passwordPath,omitempty"`
        NodeKeychainPath string         `yaml:"nodeKeychainPath,omitempty"`
        ValidatorKeychainPath string    `yaml:"validatorKeychainPath,omitempty"`
    }                                   `yaml:"smartnode,omitempty"`
    Chains struct {
        Eth1 chain                      `yaml:"eth1,omitempty"`
        Eth2 chain                      `yaml:"eth2,omitempty"`
    }                                   `yaml:"chains,omitempty"`
}
type chain struct {
    Provider string                     `yaml:"provider,omitempty"`
    Client struct {
        Options []clientOption          `yaml:"options,omitempty"`
        Selected string                 `yaml:"selected,omitempty"`
        Params []userParam              `yaml:"params,omitempty"`
    }                                   `yaml:"client,omitempty"`
}
type clientOption struct {
    ID string                           `yaml:"id,omitempty"`
    Name string                         `yaml:"name,omitempty"`
    Image string                        `yaml:"image,omitempty"`
    Params []clientParam                `yaml:"params,omitempty"`
}
type clientParam struct {
    Name string                         `yaml:"name,omitempty"`
    Env string                          `yaml:"env,omitempty"`
    Required bool                       `yaml:"required,omitempty"`
    Regex string                        `yaml:"regex,omitempty"`
}
type userParam struct {
    Env string                          `yaml:"env,omitempty"`
    Value string                        `yaml:"value"`
}


// Get the selected clients from a config
func (config *RocketPoolConfig) GetSelectedEth1Client() *clientOption {
    return config.Chains.Eth1.GetSelectedClient()
}
func (config *RocketPoolConfig) GetSelectedEth2Client() *clientOption {
    return config.Chains.Eth2.GetSelectedClient()
}
func (chain *chain) GetSelectedClient() *clientOption {
    for _, option := range chain.Client.Options {
        if option.ID == chain.Client.Selected {
            return &option
        }
    }
    return nil
}


// Load Rocket Pool config from files
// Returns global config and merged config
func Load(c *cli.Context) (RocketPoolConfig, RocketPoolConfig, error) {

    // Load configs
    globalConfig, err := loadFile(c.GlobalString("config"))
    if err != nil {
        return RocketPoolConfig{}, RocketPoolConfig{}, err
    }
    userConfig, err := loadFile(c.GlobalString("settings"))
    if err != nil {
        return RocketPoolConfig{}, RocketPoolConfig{}, err
    }
    cliConfig := getCliConfig(c)

    // Merge
    mergedConfig := mergeConfigs(&globalConfig, &userConfig, &cliConfig)

    // Return
    return globalConfig, mergedConfig, nil

}


// Load Rocket Pool config from a file
func loadFile(path string) (RocketPoolConfig, error) {

    // Read file
    // Squelch not found errors; config files are optional
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        return RocketPoolConfig{}, nil
    }

    // Parse config
    var config RocketPoolConfig
    if err := yaml.Unmarshal(bytes, &config); err != nil {
        return RocketPoolConfig{}, fmt.Errorf("Could not parse config file at %s: %w", path, err)
    }

    // Return
    return config, nil

}


// Create Rocket Pool config from CLI arguments
func getCliConfig(c *cli.Context) RocketPoolConfig {
    var config RocketPoolConfig
    config.Rocketpool.StorageAddress = c.GlobalString("storageAddress")
    config.Smartnode.PasswordPath = c.GlobalString("password")
    config.Smartnode.NodeKeychainPath = c.GlobalString("nodeKeychain")
    config.Smartnode.ValidatorKeychainPath = c.GlobalString("validatorKeychain")
    config.Chains.Eth1.Provider = c.GlobalString("eth1Provider")
    config.Chains.Eth2.Provider = c.GlobalString("eth2Provider")
    return config
}


// Merge Rocket Pool configs
func mergeConfigs(configs ...*RocketPoolConfig) RocketPoolConfig {
    var merged RocketPoolConfig
    for i := len(configs) - 1; i >= 0; i-- {
        mergo.Merge(&merged, configs[i])
    }
    return merged
}
