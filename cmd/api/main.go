package main

import (
	"log"
	"net/http"
	"strings"

	"go-modular-erp/configs"
	"go-modular-erp/internal/modules/users"
	"go-modular-erp/internal/platform/database"
	httpServer "go-modular-erp/internal/platform/http"
)

// Prefix API
func MountPrefix(root *http.ServeMux, prefix string, register func(sub *http.ServeMux)) {
	prefix = "/" + strings.Trim(prefix, "/")

	sub := http.NewServeMux()
	register(sub)
	
	root.Handle(prefix+"/", http.StripPrefix(prefix, sub))
}

func main() {

	dbCfg := configs.LoadDatabase()

	db, err := database.Connect(dbCfg.MySQLDSN())
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()

	userRepo := &users.Repository{
        DB: db,
    }

    userHandler := &users.Handler{
        Repo: userRepo,
    }

	mux := http.NewServeMux()

	MountPrefix(mux, "/api", func(api *http.ServeMux) {
		userHandler.RegisterRoutes(api)
	})

	srv := httpServer.New(":8080", mux)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}