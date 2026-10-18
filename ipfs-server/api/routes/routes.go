package routes

import (
	"net/http"

	"ipfs-server/api/controllers"
	"ipfs-server/api/middleware"
	"ipfs-server/node"
)

func Register(n *node.Node) {
	http.Handle("/", NewHandler(n))
}

func NewHandler(n *node.Node) http.Handler {
	mux := http.NewServeMux()
	upload := controllers.NewUploadController(n)
	download := controllers.NewDownloadController(n)
	replicate := controllers.NewReplicateController(n)

	mux.Handle("/upload",
		middleware.Logging(
			middleware.Auth(
				http.HandlerFunc(upload.Handle),
			),
		),
	)

	mux.Handle("/download",
		middleware.Logging(
			middleware.Auth(
				http.HandlerFunc(download.Handle),
			),
		),
	)

	mux.Handle("/replicate",
		middleware.Logging(
			http.HandlerFunc(replicate.Handle),
		),
	)

	if n.Monitor != nil && n.Monitor.Dashboard != nil {
		mux.Handle("/dashboard", n.Monitor.Dashboard)
	}

	return mux
}
