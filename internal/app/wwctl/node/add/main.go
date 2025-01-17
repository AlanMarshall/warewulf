package add

import (
	"os"

	"gopkg.in/yaml.v2"

	apinode "github.com/hpcng/warewulf/internal/pkg/api/node"
	"github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) error {
	// remove the default network as the all network values are assigned
	// to this network
	if NetName != "" {
		netDev := *NodeConf.NetDevs["default"]
		NodeConf.NetDevs[NetName] = &netDev
		delete(NodeConf.NetDevs, "default")

	}
	buffer, err := yaml.Marshal(NodeConf)
	if err != nil {
		wwlog.Error("Cant marshall nodeInfo", err)
		os.Exit(1)
	}
	set := wwapiv1.NodeAddParameter{
		NodeConfYaml: string(buffer[:]),
		NodeNames:    args,
	}

	return apinode.NodeAdd(&set)
}
