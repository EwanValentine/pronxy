package cmd

import (
	"log"
	"net/http"
	"net/url"

	"pronxy/filters"
	"pronxy/proxy"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	runReverseCmd.PersistentFlags().StringVar(&target, "target", "https://ewanvalentine.io", "Target URL")
	runReverseCmd.PersistentFlags().StringVar(&host, "host", "localhost:8080", "Host to run the proxy on")
	rootCmd.AddCommand(runReverseCmd)
}

var runReverseCmd = &cobra.Command{
	Use: "run-reverse",
	Run: runFullReverse,
}

func runFullReverse(cmd *cobra.Command, args []string) {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
		return
	}

	logger, _ := zap.NewDevelopment()
	proxyInstance := proxy.NewReverseProxy(
		u,
		logger,
		proxy.WithPreFilters(filters.NewLoggerMiddleware(logger)),
	)

	http.Handle("/", proxyInstance)

	log.Fatal(http.ListenAndServe(host, nil))
}
