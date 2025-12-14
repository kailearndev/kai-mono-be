package app

import (
	"log"
	"os"
)

type Config struct {
	Port              string
	CloudFlyEndpoint  string
	CloudFlyAccessKey string
	CloudFlySecretKey string
	CloudFlyBucket    string
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:              port,
		CloudFlyEndpoint:  os.Getenv("CLOUDFLY_ENDPOINT"),
		CloudFlyAccessKey: os.Getenv("CLOUDFLY_ACCESS_KEY"),
		CloudFlySecretKey: os.Getenv("CLOUDFLY_SECRET_KEY"),
		CloudFlyBucket:    os.Getenv("CLOUDFLY_BUCKET"),
	}
}

func (c *Config) Validate() {
	if c.CloudFlyEndpoint == "" || c.CloudFlyAccessKey == "" ||
		c.CloudFlySecretKey == "" || c.CloudFlyBucket == "" {
		log.Println("⚠️  CloudFly configuration incomplete - upload features may not work")
	}
}
