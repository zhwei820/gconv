package fn

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/zhwei820/log"
	"go.uber.org/zap"
)

// getCallerFile 获取调用者源码位置，
func getCallerFile(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip+2, pc)
	file, line := runtime.FuncForPC(pc[0]).FileLine(pc[0])
	rootPath, _ := os.Getwd()
	if rootPath != "" {
		file = strings.ReplaceAll(file, rootPath, "")
		file = strings.TrimLeft(file, "/")
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// getCallerName 获取调用者函数名，
func getCallerName(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip+2, pc)
	f := runtime.FuncForPC(pc[0])
	if strings.Contains(f.Name(), "/") {
		segs := strings.Split(f.Name(), "/")
		if len(segs) >= 1 {
			return segs[len(segs)-1]
		}
	}
	return f.Name()
}

// FuncTrace 记录函数的入口和出口，耗时
func FuncTrace(ctx context.Context, params ...interface{}) func(...interface{}) {
	funName := getCallerName(1)
	callFile := getCallerFile(1)
	params = append(params, "trace_caller", callFile)

	log.DebugZ(ctx, funName+" start", zap.Reflect("params...", params))
	start := time.Now()
	logLevel := log.StrLvlInfo
	return func(results ...interface{}) {
		results = append(results, "trace_caller", callFile, "duration", time.Since(start).Seconds())
		for i, result := range results {
			if result == nil {
				continue
			}
			//zap日志库不会对空接口自动解指针，这里解下指针
			v := reflect.ValueOf(result)
			if v.Kind() == reflect.Ptr && v.Elem().IsValid() {
				results[i] = v.Elem().Interface()
			}
			//如果有error类型，则打印error级别
			if e, ok := results[i].(error); ok && e != nil {
				logLevel = log.StrLvlError
			}
		}
		if logLevel == log.StrLvlError {
			log.ErrorZ(ctx, funName+" done", zap.Reflect("results...", results))
		} else {
			log.InfoZ(ctx, funName+" done", zap.Reflect("results...", results))
		}
	}
}
