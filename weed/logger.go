package weed

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)
var DefaultWriter io.Writer = os.Stdout

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	Request *http.Request

	Timestamp  time.Time
	StatusCode int
	Latency    time.Duration
	ClientIP   string
	Method     string
	Path       string
	ErrorMsg   string
}

func (p *LogFormatterParams) StatusColor() string {
	code := p.StatusCode
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return white
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return yellow
	default:
		return red
	}
}

func (p *LogFormatterParams) MethodColor() string {
	method := p.Method
	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	case "PATCH":
		return green
	case "HEAD":
		return magenta
	case "OPTIONS":
		return white
	default:
		return reset
	}
}

func (p *LogFormatterParams) ResetColor() string {
	return reset
}

type LoggerConfig struct {
	Formatter LogFormatter
	Output    io.Writer
}

var defaultLogFormatter = func(param LogFormatterParams) string {
	statusColor := param.StatusColor()
	methodColor := param.MethodColor()
	resetColor := param.ResetColor()
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[Tumbleweed] %v |%s %3d %s| %13v | %15s | %s %-7s %s | %s\n%s",
		param.Timestamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMsg)
}

func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{
		Formatter: defaultLogFormatter,
		Output:    DefaultWriter,
	})
}

func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
	// formatter := conf.Formatter

	return func(w http.ResponseWriter, r *http.Request) {
		// start := time.Now()
		// path := r.URL.Path

	}
}
