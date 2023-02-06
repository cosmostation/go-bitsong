package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/app"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	dbm "github.com/tendermint/tm-db"
	"io"
	"os"
	"path/filepath"
)

func openDB(rootDir string) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	return sdk.NewLevelDB("application", dataDir)
}

func openTraceWriter(traceWriterFile string) (w io.Writer, err error) {
	if traceWriterFile == "" {
		return
	}
	return os.OpenFile(
		traceWriterFile,
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		0666,
	)
}

// CustomExportCmd dump selected modules state to JSON.
func CustomExportCmd(appExporter app.AppExporter, defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "custom-export",
		Short: "Custom export state to JSON",
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			homeDir, _ := cmd.Flags().GetString(flags.FlagHome)
			config.SetRoot(homeDir)

			if _, err := os.Stat(config.GenesisFile()); os.IsNotExist(err) {
				return err
			}

			db, err := openDB(config.RootDir)
			if err != nil {
				return err
			}

			if appExporter == nil {
				return fmt.Errorf("app exporter not defined")
			}

			height, _ := cmd.Flags().GetInt64(server.FlagHeight)
			traceWriterFile, _ := cmd.Flags().GetString("trace-store")
			traceWriter, err := openTraceWriter(traceWriterFile)
			if err != nil {
				return err
			}

			exported, err := appExporter(serverCtx.Logger, db, traceWriter, height, serverCtx.Viper)
			if err != nil {
				return fmt.Errorf("error exporting state: %v", err)
			}

			doc := map[string]interface{}{
				"app_state": exported.AppState,
				"height":    exported.Height,
			}

			// NOTE: Tendermint uses a custom JSON decoder for GenesisDoc
			// (except for stuff inside AppState). Inside AppState, we're free
			// to encode as protobuf or amino.
			encoded, err := json.Marshal(doc)
			if err != nil {
				return err
			}

			cmd.Println(string(encoded))
			return nil
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")
	cmd.Flags().Int64(server.FlagHeight, -1, "Export state from a particular height (-1 means latest height)")

	return cmd
}
