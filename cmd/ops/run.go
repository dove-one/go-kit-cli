package ops

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"go-kit-cli/cmd/ops/cicd/drone"
	"go-kit-cli/cmd/ops/cicd/gogs"
	"go-kit-cli/cmd/ops/cicd/harbor"
	"go-kit-cli/cmd/ops/initGo"
	"go-kit-cli/cmd/ops/micro/consul"
	"go-kit-cli/cmd/ops/micro/jaeger"
	"go-kit-cli/cmd/ops/micro/konga"
	"go-kit-cli/cmd/ops/micro/nacos"
	"go-kit-cli/cmd/ops/service/docker"
	"go-kit-cli/cmd/ops/service/es"
	"go-kit-cli/cmd/ops/service/mysql"
	"go-kit-cli/cmd/ops/service/rabbitmq"
	"go-kit-cli/cmd/ops/service/redis"
)

func RunOps(args []string) {
	serviceName := args[0]

	// 基础服务
	{
		if serviceName == "mysql" {
			mysql.RunMysql()
			return
		}

		if serviceName == "redis" {
			redis.RunRedis()
			return
		}

		if serviceName == "rabbit" || serviceName == "rabbitmq" {
			rabbitmq.RunRabbit()
			return
		}
		if serviceName == "es" || serviceName == "elasticsearch" {
			es.RunEs()
			return
		}
	}

	// micro service
	{
		if serviceName == "nacos" {
			nacos.RunNacos()
			return
		}
		if serviceName == "jaeger" {
			jaeger.RunJaeger()
			return
		}
		if serviceName == "kong" {
			konga.RunKonga()
			return
		}
		if serviceName == "consul" {
			consul.RunConsul()
			return
		}
		if serviceName == "docker" {
			docker.RunDocker()
			return
		}
		if serviceName == "docker-compose" {
			docker.RunDockerCompose()
			return
		}
	}

	// cicd
	{
		if serviceName == "gogs" {
			gogs.RunGogs()
			return
		}
		if serviceName == "harbor" {
			harbor.RunHarbor()
			return
		}
		if serviceName == "drone" {
			drone.RunDrone()
			return
		}
	}

	// 项目环境
	{
		if serviceName == "go" {
			initGo.RunGo()
			return
		}
	}

	mlog.Print("The Command not found")
}
