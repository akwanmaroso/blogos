package api

import (
	"fmt"
	"net/http"

	"log"

	"github.com/akwanmaroso/blogos/api/router"
	"github.com/akwanmaroso/blogos/auto"
	"github.com/akwanmaroso/blogos/config"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
