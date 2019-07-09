package main

import (
	"os"

	"github.com/jchou8/sling/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.DebugLevel)

	root := &cobra.Command{
		Use:   "sling",
		Short: "chat service",
	}

	root.AddCommand(cmd.RunMigrationsCmd, cmd.RunServerCmd)
	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
