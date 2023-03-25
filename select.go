package bprompt

import (
	"os"

	"github.com/fatih/color"
)

func Select(p string,  c ...any) *Selector{
	return &Selector{
		Writer: os.Stdout,
		List: c,
		Prompt: p,
		PromptStyle: []color.Attribute{
			color.FgRed,
		},
		
	}
}
