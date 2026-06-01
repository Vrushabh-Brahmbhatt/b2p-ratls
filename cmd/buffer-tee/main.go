package main

import (
	"context"
	"log"
	"os"

	"github.com/yourorg/ratls/pkg/ratls"
)

func main() {
	ctx := context.Background()

	gpuCSAddr := getEnv("GPU_CS_ADDR", "gpu-cs.internal:443")
	audience := getEnv("RATLS_AUDIENCE", "ratls-buffer-tee")
	expectedDigest := getEnv("GPU_CS_IMAGE_DIGEST", "")

	if expectedDigest == "" {
		log.Fatal("buffer-tee: GPU_CS_IMAGE_DIGEST must be set")
	}

	opts := &ratls.VerificationOptions{
		Audience:            audience,
		ExpectedImageDigest: expectedDigest,
	}

	log.Printf("buffer-tee: connecting to GPU CS at %s", gpuCSAddr)

	session, err := ratls.Connect(ctx, gpuCSAddr, opts)
	if err != nil {
		log.Fatalf("buffer-tee: RATLS failed: %v", err)
	}

	log.Printf("buffer-tee: RATLS verified")
	log.Printf("  HWModel:     %s", session.HWModel)
	log.Printf("  ImageDigest: %s", session.ContainerImageDigest)

	if err := uploadModelWeights(ctx, session); err != nil {
		log.Fatalf("buffer-tee: upload model: %v", err)
	}
}

func uploadModelWeights(ctx context.Context, session *ratls.VerifiedSession) error {
	log.Printf("buffer-tee: uploading model weights via verified session")
	// Use session.Client to POST to GPU CS /api/load-model
	return nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
