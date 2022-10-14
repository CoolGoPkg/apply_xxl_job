package main

import (
	"CoolGoPkg/apply_xxl_job/conf"
	"CoolGoPkg/apply_xxl_job/xxl_job"
)

func main() {
	conf.InitConf("conf/config.yaml")
	xxl_job.StartXXLJobClient(conf.Config)
}
