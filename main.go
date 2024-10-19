package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)



type model struct {
    choices []string
    cursor int 
    selected map[int]struct{}
}

func initial_Model() model {
    return model{
        choices : []string{"SHADY_KNIGHT", "DOOM ETERNAL","KILL KNIGHT"},
        selected : make(map[int]struct{}),
    } 
}

func (m model) Init () tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

    switch msg := msg.(type){

    case tea.KeyMsg :
        switch msg.String(){

            case "ctrl+c", "q": 
                return m, tea.Quit
            case "up","w":
                if m.cursor > 0 {
                    m.cursor --
                }
            case "down","s":
                if m.cursor < len(m.choices)-1{
                    m.cursor ++
                }
            case " ", "enter":
                _, ok := m.selected[m.cursor]
                if ok {
                    delete(m.selected, m.cursor)
                } else {
                    m.selected[m.cursor] = struct{}{}
                }
        } 
    }

    return m, nil
}

func (m model) View () string {
    s := "\n What game do you want to rate?\n \n"

    for i, choice := range m.choices{
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }
        s += fmt.Sprintf("%s %s %s \n", cursor,checked, choice)
    }

    s += fmt.Sprintf(" \n Press q to quit. \n")

    return s
}

func main() {
    p := tea.NewProgram(initial_Model())

    if _, err := p.Run(); err != nil {
        fmt.Printf("Exiting program, error is %s", err)
        os.Exit(1)
    }
}