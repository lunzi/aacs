package main

import (
	"context"
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lunzi/aacs/internal/conf"
	"github.com/lunzi/aacs/internal/data/logr"
	"github.com/sirupsen/logrus"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// name is the name of the compiled software.
	name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&name, "name", "", "服务名")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	flag.Parse()
	var logger log.Logger
	if os.Getenv("LOG_MODE") == "production" {
		logger = log.With(logr.NewLogrusLogger(logr.Level(logrus.InfoLevel)),
			"ts", log.DefaultTimestamp,
			"service.id", id,
			"service.name", name,
			"service.version", Version,
			"trace_id", tracing.TraceID(),
			"span_id", tracing.SpanID(),
		)
	} else {
		logger = log.With(log.NewStdLogger(os.Stdout),
			"ts", log.DefaultTimestamp,
			"caller", log.DefaultCaller,
			"service.id", id,
			"service.name", name,
			"service.version", Version,
			"trace_id", tracing.TraceID(),
			"span_id", tracing.SpanID(),
		)
	}
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	oltpCleanup, err := initProvider(context.Background(), bc.Server, logger)
	if err != nil {
		panic(err)
	}
	defer oltpCleanup()
	app, cleanup, err := initApp(context.Background(), bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
