package gen

import (
	"github.com/gogf/gf-cli/v2/library/mlog"
	"go-kit-cli/util"
	"os"
)

func Run(host, user, password, port, db, table, serverName, protoName string) {
	// 1. 获取表完整结构信息
	InitDB(host, port, user, password, db)

	genReq := GenInit(serverName, table, protoName)

	mlog.Print("auto gen code start...")
	// 2. 生成项目文件结构
	CreateDir(genReq)

	// 3. 生成 model
	GenModel(genReq)

	// 4. 生成 repository
	GenRepository(genReq)

	// 5. 生成 service
	GenService(genReq)

	// 6. 生成 handler
	GenHandler(genReq)

	// 7. 生成 initialize
	GenInitlialize(genReq)

	// 8. 生成 proto
	GenProto(genReq)

	// 9. 生成 测试代码
	GenT(genReq)

	// 10.生成 web端代码
	GenWeb(genReq)

	// 11.格式化代码
	util.GoFmt(genReq.BaseDir)

	mlog.Print("done!")
}

// 创建需要的文件夹
func CreateDir(req GenReq) {
	os.MkdirAll(req.ModelDir, os.ModePerm)
	os.MkdirAll(req.RepositoryDir, os.ModePerm)
	os.MkdirAll(req.ServiceDir, os.ModePerm)
	os.MkdirAll(req.HandlerDir, os.ModePerm)
	os.MkdirAll(req.InitializeDir, os.ModePerm)
	os.MkdirAll(req.ProtoDir, os.ModePerm)
	os.MkdirAll(req.TestDir, os.ModePerm)
}
