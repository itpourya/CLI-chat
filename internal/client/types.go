package main

import (
	"os"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

var (
	mu         = sync.Mutex{}
	_          = godotenv.Load("../../.env")
	hostType   = os.Getenv("SERVER_TYPE")
	addr       = os.Getenv("SERVER_HOST")
	port       = os.Getenv("SERVER_PORT")
	name       string
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true).
			Underline(true).
			Padding(1, 2)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87")).
			Bold(true)

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFAA")).
			Margin(0, 2)

	systemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#AAAAAA")).
			Italic(true)

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4")).
			Padding(1).
			Width(50)
)
