package models

//一个功能即在模板里使用又在控制器里使用，就可以抽离出来
import "time"

func UnixTotime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-2 15:04:05")
}
