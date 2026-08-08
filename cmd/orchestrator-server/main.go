// Copyright (c) 2026 Viva Las Mesh Authors
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ipfs/go-datastore"
	"github.com/opd-ai/vivalasmesh/internal/daemon"
	"github.com/opd-ai/vivalasmesh/internal/sync"
	"github.com/opd-ai/vivalasmesh/internal/transport"
)

func main() {
	// Create an in-memory datastore for CRDT state.
	ds := datastore.NewMapDatastore()
	// Wrap it with the sync package's Datastore.
	syncStore := sync.NewDatastore(ds)

	// Create daemon config with empty transports (for now).
	config := daemon.DaemonConfig{
		NodeID:       "orchestrator-server-001",
		ListenAddrs:  []net.Addr{},            // TODO: implement actual listeners
		Transports:   []transport.Transport{}, // No transports yet
		Datastore:    syncStore,
		SyncInterval: 5 * time.Second,
	}

	// Create and start the daemon.
	d := daemon.NewDaemon(config)
	if err := d.Start(); err != nil {
		log.Fatalf("Failed to start daemon: %v", err)
	}
	defer d.Stop()

	// Wait for shutdown signal.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("Shutting down orchestrator server...")
}
