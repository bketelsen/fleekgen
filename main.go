package main

import (
	"archive/tar"
	"log"
	"os"

	"github.com/bketelsen/fleekgen/bling"
	"github.com/bketelsen/fleekgen/devbox"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	nobling, err := bling.NoBling()
	if err != nil {
		panic(err)
	}
	lowbling, err := bling.LowBling()
	if err != nil {
		panic(err)
	}
	defaultbling, err := bling.DefaultBling()
	if err != nil {
		panic(err)
	}
	highbling, err := bling.HighBling()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Devbox!")
	})
	app.Get("/none", func(c *fiber.Ctx) error {
		none := devbox.FromBling(nobling)
		tw := tar.NewWriter(c)
		defer tw.Close()
		files, err := none.Files()
		if err != nil {
			return err
		}
		c.Set("Content-Type", "application/x-tar")
		err = none.Write(files, tw)
		if err != nil {
			return err
		}
		return nil
	})
	app.Get("/low", func(c *fiber.Ctx) error {
		low := devbox.FromBling(lowbling)
		tw := tar.NewWriter(c)
		defer tw.Close()
		files, err := low.Files()
		if err != nil {
			return err
		}
		c.Set("Content-Type", "application/x-tar")
		err = low.Write(files, tw)
		if err != nil {
			return err
		}
		return nil
	})
	app.Get("/default", func(c *fiber.Ctx) error {
		dflt := devbox.FromBling(defaultbling)
		tw := tar.NewWriter(c)
		defer tw.Close()
		files, err := dflt.Files()
		if err != nil {
			return err
		}
		c.Set("Content-Type", "application/x-tar")
		err = dflt.Write(files, tw)
		if err != nil {
			return err
		}
		return nil
	})
	app.Get("/high", func(c *fiber.Ctx) error {
		high := devbox.FromBling(highbling)
		tw := tar.NewWriter(c)
		defer tw.Close()
		files, err := high.Files()
		if err != nil {
			return err
		}
		c.Set("Content-Type", "application/x-tar")
		err = high.Write(files, tw)
		if err != nil {
			return err
		}
		return nil
	})
	// Setup static files
	app.Static("/config", "./static/config")
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
