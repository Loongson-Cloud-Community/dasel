package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

func runPutIntCommand(fileFlag string, outFlag string, parserFlag string, selectorFlag string, inputValue string) error {
	parser, err := getParser(fileFlag, parserFlag)
	if err != nil {
		return err
	}
	rootNode, err := getRootNode(fileFlag, parser)
	if err != nil {
		return err
	}

	updateValue, err := parseValue(inputValue, "int")
	if err != nil {
		return err
	}

	if err := rootNode.Put(selectorFlag, updateValue); err != nil {
		return fmt.Errorf("could not put value: %w", err)
	}

	if err := writeNodeToOutput(rootNode, parser, fileFlag, outFlag); err != nil {
		return fmt.Errorf("could not write output: %w", err)
	}

	return nil
}

func putIntCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "int -f <file> -s <selector> <value>",
		Short: "Update an int property in the given file.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fileFlag := cmd.Flag("file").Value.String()
			outFlag := cmd.Flag("out").Value.String()
			parserFlag := cmd.Flag("parser").Value.String()
			selectorFlag := cmd.Flag("selector").Value.String()

			return runPutIntCommand(fileFlag, outFlag, parserFlag, selectorFlag, args[0])
		},
	}

	return cmd
}