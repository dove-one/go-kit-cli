package cmd

import (
	"runtime"

	"github.com/gogf/gf-cli/v2/library/mlog"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/spf13/cobra"

	"go-kit-cli/cmd/ops"
	"go-kit-cli/constant"
)

var (
	serviceList = `Install common service, like:
	go-kit-cli run mysql
	go-kit-cli run redis
	go-kit-cli run rabbit
	go-kit-cli run es

	go-kit-cli run consul
	go-kit-cli run jaeger
	go-kit-cli run nacos
	go-kit-cli run kong

	go-kit-cli run gogs
	go-kit-cli run harbor
	go-kit-cli run drone

	go-kit-cli run go
	go-kit-cli run docker
	go-kit-cli run docker-compose
`
)
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Install common service, like go-kit-cli run mysql",
	Long:  serviceList,
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS != "linux" {
			mlog.Print("The run command must be in linux")
			return
		}

		if len(args) == 0 {
			helpRun()
			return
		}

		ops.RunOps(args)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func helpRun() {
	mlog.Print(gstr.TrimLeft(`
USAGE
    go-kit-cli run xx

EXAMPLES
	go-kit-cli run mysql		[Initialize mysql,` + constant.MysqlVersion + `]
	go-kit-cli run redis		[Initialize redis,` + constant.RedisVersion + `]
	go-kit-cli run rabbit		[Initialize rabbit, ` + constant.RabbitmqVersion + `]
	go-kit-cli run es		[Initialize elasticsearch, ` + constant.EsVersion + `]

	go-kit-cli run consul		[Initialize consul,` + constant.ConsulVersion + `]
	go-kit-cli run nacos		[Initialize nacos,` + constant.NacosVersion + `]
	go-kit-cli run jaeger		[Initialize jaeger,` + constant.JaegerVersion + `]
	go-kit-cli run kong		[Initialize kong,` + constant.KongVersion + `]

	go-kit-cli run gogs		[Initialize gogs, ` + constant.GogsVersion + `]
	go-kit-cli run harbor		[Initialize harbor, ` + constant.HarborVersion + `]
	go-kit-cli run drone		[Initialize drone, ` + constant.DroneVersion + `]

	go-kit-cli run go		[Initialize go env, ` + constant.GOVERSION + `]
	go-kit-cli run docker		[Initialize docker, ` + constant.DockerVersion + `]
	go-kit-cli run docker-compose	[Initialize docker-compose, ` + constant.DockerComposeVersion + `]
`))
}
