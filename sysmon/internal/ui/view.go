package ui

import "fmt"



func (m Model) View() (string){
	titleBox := TitleStyle.Render("SYSTEM MONITOR DASHBOARD")
	container := ContainerStyle.Width(m.Width).Height(m.Height)
	print := fmt.Sprintf("CPU: %.1f%%, MEMORY USED: %.1f%%", m.CPUUsage, m.MemoryUsage)
	fmt.Println(print)
	return container.Render(titleBox)
}