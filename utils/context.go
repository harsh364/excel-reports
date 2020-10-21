package utils

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	ReportName string `short:"r"`
}

type Context struct {
	Options Options
	args    []string
}

func Init() (*Context, error) {
	var context Context
	args, options, err := getUserOptions()
	if err != nil {
		return &context, err
	}

	context.args = args
	context.Options = options

	return &context, nil
}

func (ctx *Context) NextArg() string {
	if len(ctx.args) == 0 {
		return ""
	}
	arg := ctx.args[0]
	ctx.args = ctx.args[1:]
	return arg
}

func getUserOptions() ([]string, Options, error) {
	var opts Options
	args, err := flags.Parse(&opts)
	return args, opts, err
}
