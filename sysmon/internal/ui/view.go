package ui



func (m Model) View() (string){
	titleBox := TitleStyle.Render("SYSTEM MONITOR DASHBOARD")
	container := ContainerStyle.Width(m.Width).Height(m.Height)

	return container.Render(titleBox)
}