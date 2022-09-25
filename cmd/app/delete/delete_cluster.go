package delete

import (
	"context"
	"log"

	"github.com/spf13/cobra"

	"github.com/GreptimeTeam/gtctl/pkg/cluster"
)

type deleteOptions struct {
	ClusterName  string
	Namespace    string
	TearDownEtcd bool
}

func NewDeleteClusterCommand() *cobra.Command {
	var options deleteOptions

	cmd := &cobra.Command{
		Use:   "cluster",
		Short: "Delete a GreptimeDB cluster.",
		Long:  `Delete a GreptimeDB cluster.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Printf("Deleting cluster %s in %s...\n", options.ClusterName, options.Namespace)

			manager, err := cluster.NewClusterManager()
			if err != nil {
				return err
			}

			ctx := context.TODO()
			_, err = manager.GetCluster(ctx, options.ClusterName, options.Namespace)
			if err != nil {
				return err
			}

			if err := manager.DeleteCluster(ctx, options.ClusterName, options.Namespace, options.TearDownEtcd); err != nil {
				return err
			}

			// TODO(zyy17): Should we wait until the cluster is actually deleted?
			log.Printf("GreptimeDB Cluster %s in %s is Deleted!\n", options.ClusterName, options.Namespace)

			return nil
		},
	}

	cmd.Flags().StringVarP(&options.ClusterName, "cluster-name", "n", "greptimedb", "Name of GreptimeDB cluster.")
	cmd.Flags().StringVar(&options.Namespace, "namespace", "default", "Namespace of GreptimeDB cluster.")
	cmd.Flags().BoolVar(&options.TearDownEtcd, "tear-down-etcd", false, "Tear down etcd cluster.")

	return cmd
}
