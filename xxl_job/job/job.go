package job

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"time"
)

func TestJob1(ctx context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done"
}

func TestJob2(ctx context.Context, param *xxl.RunReq) (msg string) {
	num := 1
	for {

		select {
		case <-ctx.Done():
			fmt.Println("task" + param.ExecutorHandler + "被手动终止")
			return
		default:
			num++
			time.Sleep(10 * time.Second)
			fmt.Println("test one task"+param.ExecutorHandler+" param："+param.ExecutorParams+"执行行", num)
			if num > 10 {
				fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + "执行完毕！")
				return
			}
		}
	}

}
