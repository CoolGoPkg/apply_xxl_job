package xxl_job

import (
	"CoolGoPkg/apply_xxl_job/conf"
	"CoolGoPkg/apply_xxl_job/xxl_job/job"
	"fmt"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"log"
	"strconv"
)

func StartXXLJobClient(conf *conf.DemoConfig) {

	exec := xxl.NewExecutor(
		xxl.RegistryKey(conf.XXLJobConf.AppName), //执行器名称
		xxl.ServerAddr(conf.XXLJobConf.AdminAddress),
		xxl.AccessToken(conf.XXLJobConf.Token),                     //请求令牌(默认为空)
		xxl.ExecutorPort(strconv.Itoa(conf.XXLJobConf.ClientPort)), //默认9999（非必填）
		xxl.SetLogger(&logger{}),                                   //自定义日志
		//xxl.ExecutorIp("127.0.0.1"),      //可自动获取
	)
	exec.Init()
	//设置日志查看handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
			FromLineNum: req.FromLineNum,
			ToLineNum:   2,
			LogContent:  "这个是自定义日志handler",
			IsEnd:       true,
		}}
	})
	//注册任务handler
	exec.RegTask("TestJob1", job.TestJob1)
	exec.RegTask("TestJob2", job.TestJob2)
	log.Fatal(exec.Run())
}

//xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}
