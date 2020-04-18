package get

import (
	"context"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Kong/kuma/app/kumactl/pkg/output"
	"github.com/Kong/kuma/app/kumactl/pkg/output/printers"
	"github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	rest_types "github.com/Kong/kuma/pkg/core/resources/model/rest"
	"github.com/Kong/kuma/pkg/core/resources/store"
)

func newGetHealthCheckCmd(pctx *getContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "healthcheck NAME",
		Short: "Show a single HealthCheck resource",
		Long:  `Show a single HealthCheck resource.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			rs, err := pctx.CurrentResourceStore()
			if err != nil {
				return err
			}
			name := args[0]
			currentMesh := pctx.CurrentMesh()
			healthcheck := &mesh.HealthCheckResource{}
			if err := rs.Get(context.Background(), healthcheck, store.GetByKey(name, currentMesh)); err != nil {
				if store.IsResourceNotFound(err) {
					return errors.Errorf("No resources found in %s mesh", currentMesh)
				}
				return errors.Wrapf(err, "failed to get mesh %s", currentMesh)
			}
			healthchecks := []*mesh.HealthCheckResource{healthcheck}
			switch format := output.Format(pctx.args.outputFormat); format {
			case output.TableFormat:
				return printHealthChecks(healthchecks, cmd.OutOrStdout())
			default:
				printer, err := printers.NewGenericPrinter(format)
				if err != nil {
					return err
				}
				return printer.Print(rest_types.From.Resource(healthcheck), cmd.OutOrStdout())
			}
		},
	}
	return cmd
}
