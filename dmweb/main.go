package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/git-download-manager/gitd-website/dmweb/loggers"
	"github.com/git-download-manager/gitd-website/dmweb/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/peterbourgon/ff/v3"
	"go.uber.org/zap"
)

var ServiceBuild string
var ServiceCommitId string

var ServiceName string = "dmweb"
var ServiceVersion string = "1.0.0"

var (
	fs = flag.NewFlagSet(ServiceName, flag.ExitOnError)

	domainName  = fs.String("domain", "localhost:3001", "domain name")
	httpAddr    = fs.String("http-addr", "localhost:3001", "http server")
	httpPrefork = fs.Bool("http-prefork", false, "http server prefork option")

	// gitdmanager options
	gitdLimitSelectList = fs.Int("gitd-limit-select-list", 5, " select max file/folder list")
	gitdLimitTreeList   = fs.Int("gitd-limit-tree-list", 30, "max number of file return")
	_                   = fs.String("env-file", ".env", "env file")
)

func main() {
	// Zap Logger Init
	loggers.SetupPlainLogger(ServiceName, ServiceVersion)

	// Start Logger: Haydaaa
	loggers.Plain.Info("service started")

	err := ff.Parse(
		fs, os.Args[1:],
		ff.WithConfigFileFlag("env-file"),
		ff.WithConfigFileParser(ff.PlainParser),
		ff.WithEnvVarPrefix(strings.ToUpper(ServiceName)),
	)

	if err != nil {
		loggers.Plain.Fatal("configration error", zap.String("err", err.Error()))
	}

	// Fiber Init
	engine := html.New("./templates", ".tmpl")
	engine.AddFunc("serviceVersion", func() string {
		return ServiceVersion
	}).AddFunc("canonicalUrl", func() string {
		return *domainName
	}).AddFunc("option", func(key string) string {
		option := map[string]string{
			"description":               "Github.com, Bitbucket.org, Gitlab.com, Gitea.com, Gitee.com provides all of the public repos in git services to download selected files and folders as a zip files with a single click, without the need for any API key or token.",
			"twitter":                   "",
			"chrome-store-url":          "https://chrome.google.com/webstore/detail/gitd-download-manager/cbnplpkljokdodpligcaolkmodfondhl",
			"firefox-addons-url":        "https://addons.mozilla.org/en-US/firefox/addon/gitd-download-manager/",
			"microsoft-edge-addons-url": "",
			"analytics-ua":              "",
			"buymeacoffee":              "",
		}
		return option[key]
	}).AddFunc("getLimit", func(key string) int {
		limits := map[string]int{
			"max-select": *gitdLimitSelectList,
			"max-file":   *gitdLimitTreeList,
		}
		return limits[key]
	}) /*.AddFunc("isNotBot", func(userAgent string) bool {
		return true
		// var re = regexp.MustCompile(`(?i)bot|crawl|curl|dataprovider|search|get|Insights|Lighthouse|spider|find|java|majesticsEO|google|yahoo|teoma|contaxe|yandex|libwww-perl|facebookexternalhit`)

		// return len(re.FindStringIndex(userAgent)) == 0
	})*/

	app := fiber.New(fiber.Config{
		Prefork:       *httpPrefork,
		CaseSensitive: true,
		AppName:       ServiceName + " " + ServiceVersion,
		Views:         engine,
		ViewsLayout:   "layouts/base",
		ErrorHandler:  middlewares.ErrorHandler,
	})

	// Assets Files
	app.Static("/assets", "./assets")

	// Middlewares
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1: Best
	}))
	app.Use(middlewares.HeaderConf(ServiceVersion))
	app.Use(fiberlogger.New(fiberlogger.Config{
		Format: "[${ip}]:${port} ${locals:requestid} ${status} - ${latency} ${method} ${path} ${queryParams}\n",
	}))
	app.Use(favicon.New())
	app.Use(recover.New()) // Sometime need to recover all

	// Routes
	setupRoutes(app)

	// 404 Middleware
	app.Use(middlewares.NotFound())

	go func() {
		err = app.Listen(*httpAddr)
		if err != nil {
			loggers.Plain.Fatal("stop.http.server", zap.Error(err))
		}
	}()

	// Listen server quit or something happened and notify channel
	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	<-close

	// shutdown server
	app.ShutdownWithTimeout(3 * time.Second)

	// Bye bye
	loggers.Plain.Info("im shutting down. see you later")
}
