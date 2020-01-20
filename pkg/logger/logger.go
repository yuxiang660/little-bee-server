package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

// Logger is the alias of logrus logger.
type Logger = logrus.Logger

// Logger Key
const (
	TraceIDKey = "trace_id"
	SpanTitleKey = "span_tigle"
	SpanFunctionKey = "span_function"
	VersionKey = "version"
)

// TraceIDFunc gets trace id for logger.
type TraceIDFunc func() string

var (
	version string
	traceIDFunc TraceIDFunc
)

// SetLevel sets the logger level.
// 1:fatal,2:error,3:warn,4:info,5:debug,6:trace.
func SetLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

// SetVersion sets the project version to logger.
func SetVersion(v string) {
	version = v
}

// SetTraceIDFunc sets callback function for trace id.
func SetTraceIDFunc(fn TraceIDFunc) {
	traceIDFunc = fn
}

type (
	traceIDContextKey struct{}
)

// NewTraceIDContext returns a copy of parent with trace id value.
func NewTraceIDContext(parent context.Context, traceID string) context.Context {
	return context.WithValue(parent, traceIDContextKey{}, traceID)
}

// GetTraceID returns trace id string.
// If the context has trace id, retrieve the string from the context.
// If the context doesn't have trace id, generate a new trace id.
func GetTraceID(ctx context.Context) string {
	v := ctx.Value(traceIDContextKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}

	if traceIDFunc != nil {
		return traceIDFunc()
	}
	return ""
}

type spanOptions struct {
	title string
	funcName string
}

// SpanOption defines function signature to set data in spanOptions.
// Span is a trace unit for a function with a title.
type SpanOption func(*spanOptions)

// SetSpanTitle returns an action to set span title.
func SetSpanTitle(title string) SpanOption {
	return func(o *spanOptions) {
		o.title = title
	}	
}

// SetSpanFuncName returns an action to set span function name.
func SetSpanFuncName(funcName string) SpanOption {
	return func(o *spanOptions) {
		o.funcName = funcName
	}
}

// Entry defines an entry with unified fields for logrus logger.
type Entry struct {
	entry *logrus.Entry
}

// StartSpan retruns an entry of a span logger.
func StartSpan(ctx context.Context, opts ...SpanOption) *Entry {
	if ctx == nil {
		ctx = context.Background()
	}

	var o spanOptions
	for _, opt := range opts {
		opt(&o)
	}

	fields := make(map[string]interface{})
	
	fields[VersionKey] = version

	if v := GetTraceID(ctx); v != ""{
		fields[TraceIDKey] = v
	}
	if v := o.title; v != "" {
		fields[SpanTitleKey] = v
	}
	if v := o.funcName; v != "" {
		fields[SpanFunctionKey] = v
	}

	return &Entry{entry: logrus.WithFields(fields)}
}

// GetStartSpanCall returns a function to start a span logger.
func GetStartSpanCall(ctx context.Context, opts ...SpanOption) func() *Entry {
	return func() *Entry {
		return StartSpan(ctx, opts...)
	}
}

// Fatalf logs fatal message.
func (e *Entry) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

// Errorf logs error message.
func (e *Entry) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

// Warnf logs warning message.
func (e *Entry) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

// Infof logs info level message.
func (e *Entry) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

// Printf logs info level message (same as Infof).
func (e *Entry) Printf(format string, args ...interface{}) {
	e.entry.Printf(format, args...)
}

// Debugf logs debug message.
func (e *Entry) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}
