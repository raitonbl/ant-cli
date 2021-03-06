package cmd

import (
	"fmt"
	"github.com/raitonbl/ant/internal"
	"github.com/raitonbl/ant/internal/commands/lint"
	"github.com/thatisuday/commando"
	"os"
)

func AddLintCommand(registry *commando.CommandRegistry) *commando.Command {
	return registry.Register("lint").
		SetShortDescription("validate a specific CLI specification file").
		SetDescription("allows the validation of an CLI specification file").
		AddArgument("file", "the CLI specification file URI", "index.json").
		SetAction(doLint)
}

func doLint(args map[string]commando.ArgValue, _ map[string]commando.FlagValue) {
	uri := args["file"].Value
	ctx, err := internal.GetContext(uri)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	problems, err := lint.Lint(ctx)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if problems != nil && len(problems) > 0 {

		txt := ""
		for index, each := range problems {
			txt += fmt.Sprintf("%d.path:%s\n message:%s", index, each.Path, each.Message)
		}

		fmt.Println(txt)
		fmt.Println("Document isn't valid")
		os.Exit(2)
	}

	fmt.Println("Document is valid")
}
