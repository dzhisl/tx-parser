package parsers

import (
	"github.com/dzhisl/tx-parser/solana/programs/computeBudget"
	"github.com/dzhisl/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitPriceData struct {
	Discriminator uint8
	Units         uint64
}

func SetComputeUnitPriceParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.ComputeBudgetSetComputeUnitPriceAction, error) {
	var data SetComputeUnitPriceData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.ComputeBudgetSetComputeUnitPriceAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     computeBudget.ProgramName,
			InstructionName: "SetComputeUnitPrice",
		},
		MicroLamports: data.Units,
	}
	return &action, nil
}
