package ui

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type TickMsg time.Time

func DoTick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg { return TickMsg(t) })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	//key press
	case tea.KeyMsg:
		//which keys pressed?
		switch msg.String() {

		// exit program
		case "q", "ctrl+c":
			return m, tea.Quit

		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case TickMsg:
		percentCpu, errCpu := cpu.Percent(0, false)
		percentMem, errMem := mem.VirtualMemory()
		if errCpu!= nil && len(percentCpu) == 0{
			m.CPUUsage = 0
			m.err = errCpu
			fmt.Println("fatal:", m.err)
			os.Exit(1)
		} else if errMem != nil{
			m.MemoryUsage = 0
			m.err = errMem
			fmt.Println("fatal:", m.err)
			os.Exit(1)
		}
		m.CPUUsage = percentCpu[0]
		m.MemoryUsage = percentMem.UsedPercent
	}
	return m, DoTick()
}
