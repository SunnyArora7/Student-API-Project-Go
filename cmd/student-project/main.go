package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"studentPackage/internal/config"
	"studentPackage/internal/config/http/handlers/student"
	sqllite "studentPackage/internal/storage/SqlLite"
	"syscall"
	"time"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqllite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/getStudent/{id}", student.GetStudent(storage))
	router.HandleFunc("GET /api/getStudents", student.GetList(storage))

	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Server running on %s", cfg.Address)
	go func() {
		error := server.ListenAndServe()
		if error != nil {
			log.Fatal("Server not responding", error)
		}
	}()
	<-done

	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	error := server.Shutdown(ctx)
	if error != nil {
		slog.Error("Enable to shut the server down", slog.String("error", error.Error()))
	}
	slog.Info("Server shut down")

}
