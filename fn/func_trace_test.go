package fn

import (
	"context"
	"testing"

	"github.com/zhwei820/log"
)

func funcTest(ctx context.Context) {
	defer FuncTrace(ctx, "req", []int{1, 2, 3, 4})()
}

func TestFuncTrace(t *testing.T) {
	log.InitLogger("", false, log.EnvDebug, 3)
	funcTest(context.Background())
}
