package main

import (
	"ClassicAddonManager/backend/app"
	"ClassicAddonManager/backend/config"
	"embed"
	"os"

	"github.com/sqweek/dialog"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	err := config.LoadConfig()
	if err != nil {
		dialog.Message("%s", err.Error()).Title("Classic Addon Manager Error").Error()
		os.Exit(1)
	}

	a := app.NewApp()

	err = wails.Run(&options.App{
		Title:         "Classic Addon Manager",
		Width:         950,
		Height:        600,
		MinHeight:     600,
		MinWidth:      950,
		DisableResize: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        a.Startup,
		Bind: []interface{}{
			a,
		},
		Windows: &windows.Options{
			DisablePinchZoom: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
