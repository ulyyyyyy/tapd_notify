package middleware

import (
	"github.com/ulyyyyyy/tapd_notify/internal/config"
	"log"
	"os"
	"testing"
)

func TestRocketProducer_StartProducer(t *testing.T) {
	if err := config.Load(); err != nil {
		log.Printf("[configs] load failed: %s\n", err)
		os.Exit(1)
	}
}
