package context

import (
	"context"
	"github.com/995933447/simpletrace"
	"testing"
)

func TestContext_GetSpanId(t *testing.T) {
	ctx := New("test", context.TODO(), "", "")
	t.Logf("t:%s, s:%s, p:%s", ctx.spanId, ctx.spanId, "")
	t.Log(simpletrace.NewSpanId())
	t.Log(simpletrace.NewSpanId())

	childCtx := ctx.NewChildSpan()
	t.Logf("t:%s, s:%s, p:%s", childCtx.traceId, childCtx.spanId, childCtx.parentCtx.spanId)
	t.Logf("t:%s, s:%s, p:%s", childCtx.NewChildSpan().traceId, childCtx.NewChildSpan().spanId, childCtx.NewChildSpan().parentCtx.spanId)
}
