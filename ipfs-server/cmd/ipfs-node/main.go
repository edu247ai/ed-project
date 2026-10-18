package main

import (
	"fmt"
	"log"
	"net/http"

	// Core subsystems
	"ipfs-server/api/routes"
	"ipfs-server/config"
	"ipfs-server/monitoring"
	"ipfs-server/node"
	"ipfs-server/node/daemon"
	"ipfs-server/security"
)

func main() {
	//
	// 1. Load configuration (env + secrets + profile)
	//
	cfg, err := config.Load("config/profiles/dev.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	//
	// 2. Initialize security subsystem
	//
	sec, err := security.New(
		cfg.Security.EncryptionKey,
		cfg.Security.SignatureKey,
	)
	if err != nil {
		log.Fatalf("failed to initialize security: %v", err)
	}

	//
	// 3. Initialize monitoring subsystem
	//
	mon := monitoring.New()

	//
	// 4. Initialize node (storage + pipeline + distribution + adapters)
	//
	n := node.NewWithConfig("node-1", cfg, sec, mon)

	//
	// 5. Register API routes
	//
	routes.Register(n)

	//
	// 6. Start API server
	//
	go func() {
		log.Printf("[api] running on :%d", cfg.API.Port)
		if err := http.ListenAndServe(
			":"+fmt.Sprint(cfg.API.Port),
			nil,
		); err != nil {
			log.Fatalf("api server failed: %v", err)
		}
	}()

	//
	// 7. Start dashboard server
	//
	go func() {
		log.Printf("[dashboard] running on :%d", cfg.Dashboard.Port)
		if err := http.ListenAndServe(
			":"+fmt.Sprint(cfg.Dashboard.Port),
			nil,
		); err != nil {
			log.Fatalf("dashboard server failed: %v", err)
		}
	}()

	//
	// 8. Start node daemon (GC, replication, events)
	//
	go func() {
		log.Println("[daemon] starting node daemon...")
		daemon.New(n).Start()
	}()

	//
	// 9. Block forever
	//
	select {}
}
