package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"net/http"
	"os"
)

func handleCORS(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "HTTP File Server"
	app.Usage = "Serve files over HTTP from a directory"
	app.Flags = []cli.Flag{
		cli.IntFlag{"port, p", 80, "port number to use"},
		cli.StringFlag{"dir, d", "./", "directory to serve"},
	}
	app.Action = func(c *cli.Context) {
		port := c.Int("port")
		directory := c.String("dir")
		fmt.Println("serving ", directory, " on port ", port)
		fmt.Println("press CTRL-C to exit")
		http.Handle("/", handleCORS(http.FileServer(http.Dir(directory))))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	}
	app.Run(os.Args)
}
