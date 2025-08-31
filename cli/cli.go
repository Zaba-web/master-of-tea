package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Zaba-web/master-of-tea/core"
)

func InitCli(app *core.TeaMaster) {
	clear()
	printHeader()
	init_commands(app)

	var arguments = os.Args[1:]

	if len(arguments) < 1 {
		printErr("Уведіть аргументи")
		return
	}

	var command = arguments[0]
	var command_ags []string

	if len(arguments) > 1 {
		command_ags = arguments[1:]
	}

	var code = exec_command(command, &command_ags)
	handleCommandCode(code)
}

func printErr(err_text string) {
	fmt.Println(colorRed + err_text + colorNone)
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func printHeader() {
	headerText := `                                                                                           
@@@@@@@@@@    @@@@@@    @@@@@@   @@@@@@@  @@@@@@@@  @@@@@@@       @@@@@@   @@@@@@@@     @@@@@@@  @@@@@@@@   @@@@@@   
@@@@@@@@@@@  @@@@@@@@  @@@@@@@   @@@@@@@  @@@@@@@@  @@@@@@@@     @@@@@@@@  @@@@@@@@     @@@@@@@  @@@@@@@@  @@@@@@@@  
@@! @@! @@!  @@!  @@@  !@@         @@!    @@!       @@!  @@@     @@!  @@@  @@!            @@!    @@!       @@!  @@@  
!@! !@! !@!  !@!  @!@  !@!         !@!    !@!       !@!  @!@     !@!  @!@  !@!            !@!    !@!       !@!  @!@  
@!! !!@ @!@  @!@!@!@!  !!@@!!      @!!    @!!!:!    @!@!!@!      @!@  !@!  @!!!:!         @!!    @!!!:!    @!@!@!@!  
!@!   ! !@!  !!!@!!!!   !!@!!!     !!!    !!!!!:    !!@!@!       !@!  !!!  !!!!!:         !!!    !!!!!:    !!!@!!!!  
!!:     !!:  !!:  !!!       !:!    !!:    !!:       !!: :!!      !!:  !!!  !!:            !!:    !!:       !!:  !!!  
:!:     :!:  :!:  !:!      !:!     :!:    :!:       :!:  !:!     :!:  !:!  :!:            :!:    :!:       :!:  !:!  
:::     ::   ::   :::  :::: ::      ::     :: ::::  ::   :::     ::::: ::   ::             ::     :: ::::  ::   :::  
 :      :     :   : :  :: : :       :     : :: ::    :   : :      : :  :    :              :     : :: ::    :   : :  
                                                                                                                     `

	fmt.Println(headerText)
}

func printTabulated(key string, value string) {
	fmt.Println(key + "\t" + value)
}

func printTeaList(teas []core.Tea) {
	fmt.Println("\n\nID\tРік\tК-сть\tЦіна\tНазва\t\t\tТеги")
	for _, tea := range teas {
		tagList := ""
		for _, tag := range tea.Tags {
			tagList = tagList + COLORS[tag.Color] + tag.Name + colorNone + " "
		}

		stock := strconv.FormatFloat(float64(tea.Stock), 'f', 2, 64)

		if tea.Stock < 25.00 {
			stock = BG_RED + stock + colorNone
		} else if tea.Stock >= 25 && tea.Stock < 75 {
			stock = colorBlack + "" + BG_YELLOW + stock + colorNone
		} else {
			stock = colorBlack + "" + GB_GREEN + stock + colorNone
		}

		fmt.Printf("%d\t%d\t%s\t%.2f\t%s\t\t%s\n", tea.ID, tea.Year, stock, tea.PricePerGram, tea.Name, tagList)
	}
}

func handleCommandCode(code int) {
	switch code {
	case COMMAND_NOT_FOUND:
		printErr("Команду не знайдено")
	case NOT_ENOUGH_ARUMENTS:
		printErr("Недостатньо аргументів")
	}
}
