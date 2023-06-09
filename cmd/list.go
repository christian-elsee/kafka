package cmd

import (
	"fmt"
	"context"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	"github.com/christian-elsee/kafka/pkg"

)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List topics",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := pkg.InterruptContext(context.Background())
		bs    := pkg.Must(cmd.Flags().GetString("bootstrap-server"))
		logf  := log.WithFields(log.Fields{
	  	"trace": pkg.Trace("listCmd.Run", "cmd/list"),
	  	"bs": bs,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")
	  defer stop()

	  admin, err := pkg.Administrator(bs)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create administrative client")
	  }
	  logf.Info("Created administrative client")

	  topics, err := pkg.GetTopics(ctx, admin)
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to get topics")
	  }
	  logf.WithFields(log.Fields{
	  	"topics": topics,
	  }).Info("Succeeded getting topics")

	  for _, v := range topics {
	  	logf.WithFields(log.Fields{
	  		"topic": v,
	  	}).Debug("Iterate listing topic")
	  	fmt.Println(v)
	  }
	},
}
func init() {
	topicCmd.AddCommand(listCmd)
}
