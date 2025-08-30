package cli

import (
	"strconv"
	"strings"

	"github.com/Zaba-web/master-of-tea/core"
)

const TEA_ADD_CMD = "tea:add"
const STATS = "stats"
const DELETE = "tea:delete"

var app *core.TeaMaster

var commands = map[string]func(*[]string) int{
	TEA_ADD_CMD: tea_add_cmd,
	STATS:       tea_stats_cmd,
	DELETE:      tea_delete_cmd,
}

func init_commands(_app *core.TeaMaster) {
	app = _app
}

func tea_add_cmd(args *[]string) int {
	if len(*args) == 0 {
		return NOT_ENOUGH_ARUMENTS
	}

	var name string = ""
	var year int = 0
	var stock float32 = 0.00
	var price_per_gram float32 = 0.00

	var tags []string

	for _, a := range *args {
		if key, val, ok := strings.Cut(a, "="); ok {
			switch key {
			case "n":
				name = val
			case "y":
				year_int, _ := strconv.ParseUint(val, 10, 32)
				year = int(year_int)
			case "s":
				stock_64, _ := strconv.ParseFloat(val, 32)
				stock = float32(stock_64)
			case "p":
				price_64, _ := strconv.ParseFloat(val, 32)
				price_per_gram = float32(price_64)
			case "t":
				tags = strings.Split(val, ",")
			}
		}
	}

	app.AddTea(name, year, stock, price_per_gram, &tags)

	return 0
}

func tea_delete_cmd(args *[]string) int {
	if len(*args) == 0 {
		return NOT_ENOUGH_ARUMENTS
	}

	id, _ := strconv.ParseInt((*args)[0], 10, 32)

	app.DeleteTea(int(id))

	return 0
}

func tea_stats_cmd(args *[]string) int {
	var teaList []core.Tea
	var filters [][]string
	var sorting []string

	if len(*args) == 0 {
		teaList = app.GetAllTeas()
	} else {
		for _, a := range *args {
			if key, val, ok := strings.Cut(a, "="); ok {
				var current_filter []string
				switch key {
				case "f":
					current_filter = strings.Split(val, "|")
					filters = append(filters, current_filter)
				case "s":
					sorting = strings.Split(val, "|")
				}
			}
		}

		teaList = app.GetTeasWithFilters(&filters, &sorting)
	}

	stats := app.GetGenetalStats(teaList)
	printTabulated("Всього чаїв:", strconv.FormatInt(int64(stats.NumberOfTeas), 10))
	printTabulated("Заг. маса:", strconv.FormatFloat(float64(stats.TotalWeight), 'f', 2, 64)+" г.")
	printTabulated("Заг. вартість:", strconv.FormatFloat(float64(stats.TotalPrice), 'f', 2, 64)+" грн.")

	printTeaList(teaList)

	return 0
}

func exec_command(command_name string, args *[]string) int {
	for c := range commands {
		if c == command_name {
			return commands[c](args)
		}
	}

	return COMMAND_NOT_FOUND
}
