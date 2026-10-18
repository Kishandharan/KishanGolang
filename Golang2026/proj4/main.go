package main
import (
	tea "charm.land/bubbletea/v2"
	"fmt"
)

type state struct{
	option string
}

func (s state) Init() tea.Cmd{
	return nil;
}

func (s state) Update(cmd tea.Msg) (tea.Model, tea.Cmd){
	event, ok := cmd.(tea.KeyPressMsg);

	if ok{
		if event.String() == "a"{
			s.option = "1. Option A";
		} else if event.String() == "b" {
			s.option = "2. Option B";
		} else if event.String() == "c" {
			s.option = "3. Option C";
		} else if event.String() == "q"{
			return s, tea.Quit;
		}
	}

	return s, nil;
}

func (s state) View() tea.View{
	view := tea.NewView(s.option);
	return view;
}

func main(){
	p := tea.NewProgram(state{option: "Press a, b, or c to start"});
	_, err := p.Run();
	if err != nil{
		fmt.Println("WE GOT AN ERROR");
	}
}
