package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/sysComputeBudget"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func Router(result *types.ParsedResult, i int) (types.Action, error) {
	instruction := result.RawTx.Transaction.Message.Instructions[i]
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := decode[0]

	switch discriminator {
	case sysComputeBudget.SetComputeUnitLimitDiscriminator:
		return SetComputeUnitLimitParser(result, i, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     "ComputeBudget",
				InstructionName: "Unknown",
			},
		}, nil
	}
}
