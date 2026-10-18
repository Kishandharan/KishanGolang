package main
import (
	"fmt"
	tea "charm.land/bubbletea/v2"
	lg "charm.land/lipgloss/v2"
)

type model struct{
	message string
}

var style1 = lg.NewStyle().Foreground(lg.Color("#00d5ff")).Bold(true).Italic(true);

func (m model) Init() tea.Cmd{
	return nil;
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	event, ok := msg.(tea.KeyPressMsg);

	if ok{
		if event.String() == "q"{
			return m, tea.Quit;
		}else if event.String() == "a" {
			m.message = "Bye";
		}else if event.String() == "b"{
			m.message = "Bubbletea is amazing";
		}
	}

	return m, nil
}

func (m model) View() tea.View{
	view1 := tea.NewView(style1.Render(m.message));
	view2 := tea.NewView(style1.Render(m.message));
	view1.AltScreen = true;
	view2.AltScreen = true;
	return view1;
}

func main(){
	p := tea.NewProgram(model{message: "Hi bro!"});

	_, err := p.Run();

	if err != nil{
		fmt.Println("WE GOT AN ERROR! AAAAH");
	}
}
