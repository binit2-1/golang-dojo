package ui

import "fmt"



func (m Model) View() (string){
	s := fmt.Sprintf("CPU Usage %v", m.CPUUsage )
	return s
}