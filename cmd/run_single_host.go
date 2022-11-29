package cmd

import (
	"fmt"
	"log"
	"net/http"

	"pronxy/proxy"

	"github.com/spf13/cobra"
)

var (
	target string
	host   string
)

func init() {
	runSingleHostCmd.PersistentFlags().StringVar(&target, "target", "https://ewanvalentine.io", "Target URL")
	runSingleHostCmd.PersistentFlags().StringVar(&host, "host", "localhost:8080", "Host to run the proxy on")
	rootCmd.AddCommand(runSingleHostCmd)
}

var runSingleHostCmd = &cobra.Command{
	Use: "run-single-host",
	Run: runSingleHost,
}

func runSingleHost(cmd *cobra.Command, args []string) {
	proxyInstance, err := proxy.NewSingleHostProxy(target)
	if err != nil {
		fmt.Errorf("error starting proxy: %v", err)
		return
	}

	http.HandleFunc("/", proxy.ProxyRequestHandler(proxyInstance))

	log.Fatal(http.ListenAndServe(host, nil))
}
