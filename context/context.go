package context

import (
	"context"
	"wegod/internal/pkg/trace/simpletrace"
)

type Context struct {
	context.Context
	traceId string
	spanId string
	ctx context.Context
	moduleName string
	parentCtx *Context
}

func New(moduleName string, ctx context.Context, traceId, spanId string) *Context {
	if traceId == "" {
		traceId = simpletrace.NewTraceId()
	}
	if spanId == "" {
		spanId = simpletrace.NewSpanId()
	}
	return &Context{
		moduleName: moduleName,
		traceId: traceId,
		spanId: spanId,
		ctx: ctx,
	}
}

func (c *Context) GetTraceId() string {
	return c.traceId
}

func (c *Context) GetSpanId() string {
	return c.spanId
}

func (c *Context) GetModuleName() string {
	return c.moduleName
}

func (c *Context) GetParentCtx() *Context {
	return c.parentCtx
}

func (c *Context) GetParentSpanId() string {
	if c.parentCtx != nil {
		return c.parentCtx.spanId
	}
	return ""
}

func (c *Context) NewChildSpan() *Context {
	return &Context{
		traceId: c.traceId,
		spanId: simpletrace.NewSpanId(),
		parentCtx: c,
		Context: c.Context,
		moduleName: c.moduleName,
	}
}
