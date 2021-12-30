package cli

import (
    "context"
	

	
    "github.com/spf13/cobra"
    
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/stafiprotocol/stafihub/x/rate/types"
)

func CmdListExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-exchange-rate",
		Short: "list all ExchangeRate",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllExchangeRateRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.ExchangeRateAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-exchange-rate [index]",
		Short: "shows a ExchangeRate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argIndex := args[0]
            
            params := &types.QueryGetExchangeRateRequest{
                Index: argIndex,
                
            }

            res, err := queryClient.ExchangeRate(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
