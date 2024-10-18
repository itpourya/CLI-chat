package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
)

type model struct {
	conn     net.Conn
	messages []string
	input    string
	err      error
}

func main() {
	conn, err := net.Dial(hostType, addr+":"+port)
	if err != nil {
		log.Fatal("Failed to connect on ", addr+":"+port)
	}
	defer conn.Close()

	log.Info("Connected to the server at " + addr + ":" + port)

	huh.NewInput().Title("What’s your name?").Value(&name).Run()
	username := name

	_, err = conn.Write([]byte(username + "\n"))
	if err != nil {
		log.Error("Failed to send username ", "error", err)
		return
	}

	m := model{conn: conn}
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}

func (m model) Init() tea.Cmd {
	return m.readMessage
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.conn.Write([]byte(m.input + "\n"))
			timestamp := time.Now().Format("15:04")
			m.messages = append(m.messages, fmt.Sprintf("[%s] You: %s", timestamp, strings.TrimSpace(m.input)))
			m.input = ""
			return m, nil
		default:
			m.input += msg.String()
			return m, nil
		}
	case string:
		m.messages = append(m.messages, msg)
		return m, m.readMessage
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("✨ Welcome to GoChat! ✨")

	// Messages section
	messages := strings.Join(m.messages, "\n")
	messageBox := borderStyle.Render(messages)

	// Input section
	inputBox := inputStyle.Render("You: " + m.input)

	// Return final view
	return fmt.Sprintf("%s\n%s\n\n%s", s, messageBox, inputBox)
}

func (m model) readMessage() tea.Msg {
	message, err := bufio.NewReader(m.conn).ReadString('\n')
	if err != nil {
		m.err = err
		return tea.Quit
	}
	timestamp := time.Now().Format("15:04")
	return fmt.Sprintf("[%s] %s", timestamp, strings.TrimSpace(message))
}
