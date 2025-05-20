package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/panorama32/simplemcp"
)

type RollDiceArguments struct {
	Sides int `json:"sides"`
}

func registerRollDiceTool(h *simplemcp.Handler) {
	h.RegisterTool(&simplemcp.RegisterToolConfig{
		Name:        "roll_dice",
		Description: "Roll a dice",
		Properties: map[string]simplemcp.Property{
			"sides": {
				Type:        simplemcp.PropertyTypeInteger,
				Description: "Number of sides on the dice.",
			},
		},
		ToolFunc: func(
			ctx context.Context,
			arguments json.RawMessage,
		) (simplemcp.CallToolResult, error) {
			var (
				rollDiceArguments = RollDiceArguments{}
				result            = simplemcp.NewCallToolResult()
			)
			err := json.Unmarshal(arguments, &rollDiceArguments)
			if err != nil {
				result.IsError = true
				result.AddTextContent("Invalid parameters")
				return result, err
			}
			diceResult := rand.Intn(rollDiceArguments.Sides) + 1
			resultTxt := strconv.Itoa(diceResult)
			if rollDiceArguments.Sides == 100 {
				if diceResult <= 5 {
					resultTxt = fmt.Sprintf(
						"%d: 決定的成功(クリティカル)",
						diceResult,
					)
				} else if diceResult <= 20 {
					resultTxt = fmt.Sprintf(
						"%d: 成功",
						diceResult,
					)
				} else if diceResult <= 95 {
					resultTxt = fmt.Sprintf(
						"%d: 失敗",
						diceResult,
					)
				} else {
					resultTxt = fmt.Sprintf(
						"%d: 致命的失敗(ファンブル)",
						diceResult,
					)
				}
			}
			result.AddTextContent(resultTxt)
			return result, nil
		},
	})
}

func main() {
	h := simplemcp.NewHandler(&simplemcp.Implementation{
		Name:    "dice_roller",
		Version: "0.0.1",
	})
	registerRollDiceTool(h)
	h.Run(context.Background())
}
