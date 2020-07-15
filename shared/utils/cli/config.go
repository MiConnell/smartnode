package cli

import (
    "github.com/urfave/cli"
)


// Configure application
func Configure(app *cli.App) {
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "config",
            Usage: "Rocket Pool service global config absolute `path`",
            Value: "/.rocketpool/config.yml",
        },
        cli.StringFlag{
            Name:  "settings",
            Usage: "Rocket Pool service user config absolute `path`",
            Value: "/.rocketpool/settings.yml",
        },
        cli.StringFlag{
            Name:  "storageAddress",
            Usage: "Rocket Pool storage contract `address`",
        },
        cli.StringFlag{
            Name:  "password",
            Usage: "Rocket Pool service keychain password absolute `path`",
        },
        cli.StringFlag{
            Name:  "nodeKeychain",
            Usage: "Node account keychain absolute `path`",
        },
        cli.StringFlag{
            Name:  "validatorKeychain",
            Usage: "Validator keychain absolute `path`",
        },
        cli.StringFlag{
            Name:  "eth1Provider",
            Usage: "Eth 1.0 provider `address`",
        },
        cli.StringFlag{
            Name:  "eth2Provider",
            Usage: "Eth 2.0 provider `address`",
        },
    }
}
