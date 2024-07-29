package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/GuoFlight/gkeybd"
	"github.com/gofiber/fiber/v2"
	"github.com/micmonay/keybd_event"
)

type OpenRequest struct {
	AppName  string `json:"app_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	NoBpjs   string `json:"no_bpjs"`
}
type CloseRequest struct {
	AppName string `json:"app_name"`
}

func main() {

	fmt.Println("Letakan file ini sejajar dengan instalan Fingerprint BPJS (After.exe) di C:\\Program Files (x86)\\BPJS Kesehatan\\Aplikasi Sidik Jari BPJS Kesehatan\\ ")
	fmt.Println("code : https://github.com/aripkur/service-sidikjari-bpjs")

	app := fiber.New(fiber.Config{
		AppName: "Restapi Fingerprint BPJS",
	})
	app.Post("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "READY!!!"})
	})
	app.Post("/open", func(c *fiber.Ctx) error {
		var request OpenRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(422).JSON(fiber.Map{"message": "Bad Request"})
		}

		filePath := getExePath(request.AppName)
		cmd := exec.Command(filePath)
		err := cmd.Start()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		time.Sleep(4000 * time.Millisecond)

		gkeybd.TypeStr(request.Username)
		pressTab()

		gkeybd.TypeStr(request.Password)
		pressTab()
		pressEnter()

		time.Sleep(3000 * time.Millisecond)
		gkeybd.TypeStr(request.NoBpjs)

		return c.Status(200).JSON(fiber.Map{"message": "success"})
	})

	app.Post("/close", func(c *fiber.Ctx) error {
		var request CloseRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(422).JSON(fiber.Map{"message": "Bad Request"})
		}
		filePath := getExePath(request.AppName)
		cmd := exec.Command("TASKKILL", "/IM", filepath.Base(filePath), "/F")
		cmd.Run()
		return c.Status(200).JSON(fiber.Map{"message": "success"})
	})

	err := app.Listen(":3005")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}

}

func getExePath(app string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return ""
	}
	return filepath.Join(currentDir, app)
}

func pressTab() error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	kb.SetKeys(keybd_event.VK_TAB)
	err = kb.Launching()
	if err != nil {
		return err
	}
	return nil
}
func pressEnter() error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	kb.SetKeys(keybd_event.VK_ENTER)
	err = kb.Launching()
	if err != nil {
		return err
	}
	return nil
}
