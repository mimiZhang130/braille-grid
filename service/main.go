package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"go.bug.st/serial"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	ports        []string
	baudRate     textinput.Model
	cursor       int
	state        int
	selectedPort string
	host         string
	err          error
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "9600"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	return model{
		baudRate: ti,
		err:      nil,
		ports:    ports,
		state:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	state := m.state
	var cmd tea.Cmd
	if state == 0 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.ports)-1 {
					m.cursor++
				}
			case "enter", " ":
				m.selectedPort = m.ports[m.cursor]
				m.state = 1
			}
		}
	} else if state == 1 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter", " ":
				m.state = 2
			}
		}
		m.baudRate, cmd = m.baudRate.Update(msg)
	} else if state == 2 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "enter", " ":
				baud, err := strconv.Atoi(m.baudRate.Value())
				if err != nil {
					log.Fatal(err)
				}
				mode := &serial.Mode{
					BaudRate: baud,
					Parity:   serial.NoParity,
					DataBits: 8,
					StopBits: serial.OneStopBit,
				}
				port, err := serial.Open(m.selectedPort, mode)
				if err != nil {
					log.Fatal(err)
				}
				m.state = 3

				host := "localhost:3000"
				m.host = host

				go func() {
					msgch := make(chan []byte, 10)
					http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
						echo(w, r, msgch)
					})
					log.Fatal(http.ListenAndServe(":3000", nil))
					go func() {
						for msg := range msgch {
							_, err = port.Write(msg)
                            if (err != nil) {
                                log.Fatal(err);
                            }
						}
					}()
				}()

			}
		}
	} else if state == 3 {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		}
	}
	return m, cmd
}

func (m model) View() string {
	s := ""
	state := m.state
	if state == 0 {
		s += "Please select the serial port to connect to?\n\n"
		for i, choice := range m.ports {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
	} else if state == 1 {
		s += fmt.Sprintf(
			"Baud rate?\n\n%s",
			m.baudRate.View(),
		) + "\n"
	} else if state == 2 {
		s += fmt.Sprintf("Press enter to start http server for the port %s and at baud rate of %s:\n", m.selectedPort, m.baudRate.Value())
	} else if state == 3 {
		s += fmt.Sprintf("Serving http server at %s", m.host)
	}
	s += "\nPress q to quit.\n"
	return s
}

func echo(w http.ResponseWriter, r *http.Request, msgch chan []byte) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading request body:", err)
	}
	defer r.Body.Close()
	io.WriteString(w, "ok!")
	msgch <- body
}
