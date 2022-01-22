package main

import (
	"bnpl/dao/postgres"
	"bnpl/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	apiPort = 8080
)

func main() {
	var handler, err = handlers.NewHandler(context.Background(), postgres.NewUserManager,
		postgres.NewMerchantManager,
		postgres.NewTransactionManager,
		postgres.NewReportManager)
	if err != nil {
		log.Fatalf("failed to create new handler:%v", err)
	}

	var router = gin.Default()
	// register all routes.
	handlers.Registration(router, &handler)

	var apiAddr = fmt.Sprintf(":%d", apiPort)

	var srv = &http.Server{
		Addr:    apiAddr,
		Handler: router,
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("failed to listen on: %s\n", err)
		}
	}()

	log.Printf("server started successfully on %s", apiAddr)

	wg.Wait()
}
