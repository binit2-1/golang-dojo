package main

import (
	"github.com/binit2-1/golang-dojo/sysmon/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"fmt"
	"os"
)

func main(){
	p := tea.NewProgram(ui.Model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil{
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}