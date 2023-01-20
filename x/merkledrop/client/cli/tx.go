package cli

import (
	"encoding/json"
	"fmt"
	"github.com/bitsongofficial/go-bitsong/x/merkledrop/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

// NewTxCmd returns the transaction commands for the merkledrop module.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "merkledrop transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdCreate(),
		GetCmdClaim(),
	)

	return txCmd
}

func GetCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [file-json] [out-list-json]",
		Short: "Create a merkledrop from json file",
		Long: `Create a merkledrop from json file
Parameters:
	file-json: input file list
	out-list-json: output list with proofs

Flags:
	denom: the coin denom to distribuite
	start-height: the height when the merkledrop will begin (0 for immediatally)
	end-height: the height when the merkledrop will ends
		`,
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s tx merkledrop create accounts.json out-list.json \
	--denom=ubtsg \
	--start-height=1 \
	--end-height=10 \
	--from=<key-name>

where accounts.json contains
{
	"bitsong10clahhd4g878vzyl69hcnue9uufp5dle4867md": "1000000",
	"bitsong1zm6wlhr622yr9d7hh4t70acdfg6c32kcv34duw": "2000000",
	"bitsong1nzxmsks45e55d5edj4mcd08u8dycaxq5eplakw": "3000000"
}

after the computation the out-list/*.json should be similar to this output
{
  "address": "bitsong107yfv7396n7ket3j0l666trx6ww729q9my785g",
  "index": 3,
  "amount": "1000000000",
  "proof": [
    "228b2d5a49f2724a287443a2ed95d7c663aa40e5df8486efac7f3cdad1b1b0f1",
    "4026fef271e51c652b1e0c98673e4b68beb3bdf9136520ffd8397497efb0c39e",
    "75538c0c43de52a7daa75daaa5029fa8972685849326565722eaa65141c455b9",
    "ca617d2b4256c08bcaecc8a271bbbf9a21325a7bb0a707baf0e7a1a56fa6e972"
  ]
}
{
  "bitsong10clahhd4g878vzyl69hcnue9uufp5dle4867md": {
    "index": 0,
    "amount": "100000",
    "proof": [
      "342cb422e73af25dbb535ea27799d228b9f89a634481cb44325f1b2375ebedc4",
      "b6b9c249fbe8ef1425edd44ae0e1e7f7b4ee26828dfdd00f1b375755eb51550b"
    ]
  },
  "bitsong1nzxmsks45e55d5edj4mcd08u8dycaxq5eplakw": {
    "index": 2,
    "amount": "300000",
    "proof": [
      "c6b063c83b4c971a78466f019e68b90fb97b93f43ecb9d9b29060f54d754c10e"
    ]
  },
  "bitsong1zm6wlhr622yr9d7hh4t70acdfg6c32kcv34duw": {
    "index": 1,
    "amount": "200000",
    "proof": [
      "8c086a5802a9978d1e9fc13259566e3594928703b06c0e845cf45b25936c1fe7",
      "b6b9c249fbe8ef1425edd44ae0e1e7f7b4ee26828dfdd00f1b375755eb51550b"
    ]
  }
}
`,
			version.AppName,
		)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			listBytes, err := ioutil.ReadFile(args[0])
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", args[0], err)
			}

			var stringList map[string]string
			if err := json.Unmarshal(listBytes, &stringList); err != nil {
				return fmt.Errorf("could not unmarshal json: %v", err)
			}

			accMap, err := AccountsFromMap(stringList)
			if err != nil {
				return fmt.Errorf("could not get accounts from map")
			}

			tree, claimInfos, totalAmt, err := CreateDistributionList(accMap)
			if err != nil {
				return fmt.Errorf("could not create distribution list: %v", err)
			}

			tmpDir := os.TempDir()
			merkledropDir := tmpDir + "/merkledrop"

			// Remove the directory if it exists
			if _, err := os.Stat(merkledropDir); !os.IsNotExist(err) {
				err := os.RemoveAll(merkledropDir)
				if err != nil {
					return fmt.Errorf("could not remove directory: %v", err)
				}
			}

			// Create the directory if it does not exist
			if _, err := os.Stat(merkledropDir); os.IsNotExist(err) {
				err := os.Mkdir(merkledropDir, 0755)
				if err != nil {
					return fmt.Errorf("could not create directory: %v", err)
				}
			}

			for i, _ := range claimInfos {
				claimInfo := claimInfos[i]
				claimInfoBytes, err := json.Marshal(claimInfo)
				if err != nil {
					return fmt.Errorf("could not marshal claim info: %v", err)
				}
				filePath := fmt.Sprintf("%s/%s.json", merkledropDir, claimInfo.Address)
				err = ioutil.WriteFile(filePath, claimInfoBytes, 0644)
				if err != nil {
					return fmt.Errorf("could not write file: %v", err)
				}
			}

			ipfsNode, err := cmd.Flags().GetString("ipfs-node")
			ipfs := ipfsapi.NewShell(ipfsNode)
			ipfsDir, err := ipfs.AddDir(merkledropDir)
			if err != nil {
				return fmt.Errorf("could not add directory to ipfs: %v", err)
			}
			// pin uploaded dir to ipfs
			err = ipfs.Pin(ipfsDir)
			if err != nil {
				return fmt.Errorf("could not pin directory to ipfs: %v", err)
			}

			startHeight, endHeight, denom, err := parseGenerateFlags(cmd.Flags())
			if denom == "" {
				return fmt.Errorf("denom cannot be empty")
			}
			merkleRoot := fmt.Sprintf("%x", tree.Root())

			coin, err := sdk.ParseCoinNormalized(fmt.Sprintf("%s%s", totalAmt.String(), denom))
			if err != nil {
				return fmt.Errorf("could not parse coin: %v", err)
			}

			msg := types.NewMsgCreate(clientCtx.GetFromAddress(), merkleRoot, startHeight, endHeight, coin)

			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %v", err)
			}

			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if err != nil {
				return fmt.Errorf("could not broadcast tx: %v", err)
			}

			fmt.Println("IPFS Directory: " + ipfsDir)
			fmt.Println("IPFS Directory Link: https://bas-cdn.com/ipfs/" + ipfsDir)

			return nil
		},
	}

	cmd.Flags().AddFlagSet(FlagsCreate())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseGenerateFlags(flags *flag.FlagSet) (int64, int64, string, error) {
	startHeight, err := flags.GetInt64(FlagStartHeight)
	if err != nil {
		return 0, 0, "", err
	}

	endHeight, err := flags.GetInt64(FlagEndHeight)
	if err != nil {
		return 0, 0, "", err
	}

	denom, err := flags.GetString(FlagDenom)
	if err != nil {
		return 0, 0, "", err
	}

	return startHeight, endHeight, denom, nil
}

func GetCmdClaim() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim [id]",
		Short: "Claim a merkledrop from provided params",
		Long: `Claim a merkledrop from provided params
Parameters:
	id: merkledrop id

Flags:
	proofs: merkle-proofs to claim the merkledrop
	amount: the amount of the merkledrop to claim
	index: the index of the merkledrop to claim
		`,
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s tx merkledrop claim 1 \
	--proofs="a258c32bee9b0bbb7a2d1999ab4698294844e7440aa6dcd067e0d5142fa20522,7f0b92cc8318e4fb0db9052325b474e2eabb80d79e6e1abab92093d3a88fe029" \
	--amount=20000 \
	--index=1
	--from=<key-name>
`,
			version.AppName,
		)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			merkledropId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			proofsStr, err := cmd.Flags().GetString(FlagProofs)
			if err != nil {
				return err
			}
			proofs := []string{}
			if proofsStr != "" {
				proofs = strings.Split(proofsStr, ",")
			}

			amount, err := cmd.Flags().GetInt64(FlagAmount)
			if err != nil {
				return err
			}

			index, err := cmd.Flags().GetUint64(FlagIndex)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaim(index, merkledropId, sdk.NewInt(amount), proofs, clientCtx.GetFromAddress())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagClaimMerkledrop())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdUpdateMerkledropFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-merkledrop-fees [proposal-file]",
		Short: "Submit an update merkledrop fees proposal.",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an update merkledrop fees proposal along with an initial deposit.
The proposal details must be supplied via a JSON file.
Example:
$ %s tx gov submit-proposal update-merkledrop-fees <path/to/proposal.json> --from=<key_or_address>
Where proposal.json contains:
{
  "title": "Update Merkledrop Fees Proposal",
  "description": "update the current fees",
  "creation_fee": "1000000ubtsg",
  "deposit": "500000000ubtsg"
}
`, version.AppName,
			),
		),
		Example: fmt.Sprintf(
			"$ %s tx gov submit-proposal update-merkledrop-fees [proposal-file] "+
				"--from=<key-name> "+
				"--chain-id=<chain-id> "+
				"--fees=<fee>",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := parseUpdateFeesProposal(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			creationFee, err := sdk.ParseCoinNormalized(proposal.CreationFee)
			if err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}

			content := types.NewUpdateFeesProposal(proposal.Title, proposal.Description, creationFee)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}

func parseUpdateFeesProposal(cdc codec.JSONCodec, proposalFile string) (types.UpdateFeesProposalWithDeposit, error) {
	proposal := types.UpdateFeesProposalWithDeposit{}

	contents, err := os.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err = cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
