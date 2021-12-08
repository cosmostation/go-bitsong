package cmd

import (
	"github.com/bitsongofficial/go-bitsong/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRootCmdConfig(t *testing.T) {

	rootCmd, _ := NewRootCmd()
	rootCmd.SetArgs([]string{
		"config",          // Test the config cmd
		"keyring-backend", // key
		"test",            // value
	})

	require.NoError(t, svrcmd.Execute(rootCmd, app.DefaultNodeHome))
}
