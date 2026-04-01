package ProjectLibrary

import (
	"Poligon/ProjectLibrary/http"
	"Poligon/ProjectLibrary/library"
	"fmt"
)

func MainLibrary() {
	lib := library.NewLibrary()
	httpHandlers := http.NewHTTPHadlersBook(lib)
	httpServer := http.NewHTTPServer(httpHandlers)

	fmt.Println("Start server", httpServer)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server ", err)
	}
}
