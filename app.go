package main

import (
	"context"
	"nsmid/nvidiasmi"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	go a.startAndListenForUpdates()
}

func (a *App) startAndListenForUpdates() {
	go nvidiasmi.HandleNvidiaSMI()
	for {
		<-nvidiasmi.UpdateChan
		a.updateFrontendData()
	}
}

func (a *App) updateFrontendData() {
	data := nvidiasmi.GetData()
	runtime.EventsEmit(a.ctx, "gpuDataUpdated", data)
}
