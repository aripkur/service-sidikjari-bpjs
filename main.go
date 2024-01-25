package main

import (
	"errors"
	"time"
	"unicode/utf16"

	"github.com/GuoFlight/gkeybd"
	"github.com/gofiber/fiber/v2"
	"github.com/lxn/win"
	"github.com/micmonay/keybd_event"
)

type OpenRequest struct {
	AppName     string `json:"app_name"`
	NoIdentitas string `json:"no_indetitas"`
}
type CloseRequest struct {
	AppName string `json:"app_name"`
}

func main() {
	app := fiber.New()

	app.Post("/open", func(c *fiber.Ctx) error {
		var request OpenRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(422).JSON(fiber.Map{"message": "Bad Request"})
		}
		targetAppNameUTF16 := utf16.Encode([]rune(request.AppName + "\x00"))
		appHandle, err := findWindow(targetAppNameUTF16)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}

		err = openWindow(appHandle)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}

		time.Sleep(2 * time.Second)

		err = tapCtrlA()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}

		time.Sleep(1 * time.Second)

		err = gkeybd.TypeStr(request.NoIdentitas)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}

		return c.Status(200).JSON(fiber.Map{"message": "success"})
	})

	app.Post("/close", func(c *fiber.Ctx) error {
		var request CloseRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(422).JSON(fiber.Map{"message": "Bad Request"})
		}
		targetAppNameUTF16 := utf16.Encode([]rune(request.AppName + "\x00"))
		appHandle, err := findWindow(targetAppNameUTF16)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}

		err = closeWindow(appHandle)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{"message": "success"})
	})

	err := app.Listen(":3001")
	if err != nil {
		panic(err)
	}
}

func findWindow(targetAppNameUTF16 []uint16) (win.HWND, error) {
	hwnd := win.FindWindow(nil, &targetAppNameUTF16[0])
	if hwnd == 0 {
		return 0, errors.New("aplikasi tidak ditemukan")
	}
	return hwnd, nil
}
func openWindow(appHandle win.HWND) error {
	if !win.ShowWindow(appHandle, win.SW_RESTORE) {
		return errors.New("tidak bisa membuka aplikasi")
	}
	if !win.SetForegroundWindow(appHandle) {
		return errors.New("tidak bisa membuka aplikasi")
	}
	return nil
}

func closeWindow(appHandle win.HWND) error {
	if !win.ShowWindow(appHandle, win.SW_MINIMIZE) {
		return errors.New("tidak bisa menutup aplikasi")
	}
	return nil
}

func tapCtrlA() error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}
	kb.HasCTRL(true)
	kb.SetKeys(keybd_event.VK_A)
	err = kb.Launching()
	if err != nil {
		return err
	}
	return nil
}
