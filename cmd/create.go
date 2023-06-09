package cmd

import (
	"context"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	"github.com/christian-elsee/kafka/pkg"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a topic",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := pkg.InterruptContext(context.Background())
		bs    := pkg.Must(cmd.Flags().GetString("bootstrap-server"))
		topic := pkg.Must(cmd.Flags().GetString("topic"))
		partitions := pkg.Must(cmd.Flags().GetInt("partitions"))
		replicas := pkg.Must(cmd.Flags().GetInt("replication-factor"))

		logf  := log.WithFields(log.Fields{
	  	"trace": pkg.Trace("createCmd.Run", "cmd/create"),
	  	"bs": bs,
	  	"topic": topic,
	  	"partitions": partitions,
	  	"replicas": replicas,
	  })
	  logf.Debug("Enter")
	  defer logf.Debug("Exit")
	  defer stop()

	  // create administrator client to kafka using bootstrap
	  // servers passed as a required flag to this cmd
	  admin, err := pkg.Administrator(bs)
	  if err != nil { 
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create administrative client")
	  }
	  logf.Debug("Created administrative client")

	  err = pkg.CreateTopic(
	  	ctx,
	  	admin,
	  	topic,
	  	partitions,
	  	replicas,
	  )
	  if err != nil {
	  	logf.WithFields(log.Fields{
	  		"error": err,
	  	}).Fatal("Failed to create topic")
	  }
	  logf.Info("Created topic")

	},
}

func init() {
	topicCmd.AddCommand(createCmd)
	createCmd.Flags().IntP("partitions", "p", 1, "Number of partitions")
	createCmd.Flags().IntP("replication-factor", "r", 1, "Number of replicas")
	createCmd.Flags().StringP("topic", "t", "", "The topic to run operation against")
	createCmd.MarkFlagRequired("topic")
}
