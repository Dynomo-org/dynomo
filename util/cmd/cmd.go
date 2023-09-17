package cmd

import (
	"context"
	"fmt"
	"os/exec"
)

func ExecCommandWithContext(ctx context.Context, command string) error {
	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return err
	}

	return err
}
