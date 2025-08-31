package cli

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/Zaba-web/master-of-tea/core"
)

const TEA_ADD_CMD = "tea:add"
const STATS = "stats"
const DELETE = "tea:delete"
const ADD_TAG = "tea:add:tag"
const HELP_CMD = "help"
const BREW = "tea:brew"
const TEA_IMPORT = "tea:import"

var app *core.TeaMaster

var commands = map[string]func(*[]string) int{
	TEA_ADD_CMD: tea_add_cmd,
	STATS:       tea_stats_cmd,
	DELETE:      tea_delete_cmd,
	ADD_TAG:     add_tag_cmd,
	HELP_CMD:    help_cmd,
	BREW:        brew_cmd,
	TEA_IMPORT:  tea_import_cmd,
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
				case "t":
					tags := strings.Split(val, ",")
					productIds := app.GetProductIdsByTags(tags)

					current_filter = append(current_filter, "id")
					current_filter = append(current_filter, "IN")

					strs := make([]string, len(productIds))
					for i, v := range productIds {
						strs[i] = strconv.Itoa(v)
					}

					current_filter = append(current_filter, strings.Join(strs, ","))

					filters = append(filters, current_filter)
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

func add_tag_cmd(args *[]string) int {
	if len(*args) == 0 {
		return NOT_ENOUGH_ARUMENTS
	}

	var teaId int
	var tagName string

	for _, a := range *args {
		if key, val, ok := strings.Cut(a, "="); ok {
			switch key {
			case "tea":
				bigIntId, _ := strconv.ParseInt(val, 10, 32)
				teaId = int(bigIntId)
			case "tag":
				tagName = val
			}
		}
	}

	app.AddTag(teaId, tagName)

	return 0
}

func brew_cmd(args *[]string) int {
	var teaId int
	var weight float32

	for _, a := range *args {
		if key, val, ok := strings.Cut(a, "="); ok {
			switch key {
			case "t":
				bigIntId, _ := strconv.ParseInt(val, 10, 32)
				teaId = int(bigIntId)
			case "w":
				weight_64, _ := strconv.ParseFloat(val, 32)
				weight = float32(weight_64)
			}
		}
	}

	app.Brew(teaId, weight)

	return 0
}

func help_cmd(args *[]string) int {
	printTabulated("Додати чай:", "tea:add n=назва y=рік s=к-сть в наявності p=ціна за грам t=теги")
	printTabulated("Статистика:", "stats (f=поле|умова|значення) (s=поле|напрямок) (t=тег1,тег2...)")
	printTabulated("Видалити чай:", "tea:delete id")
	printTabulated("Додати тег:", "tea:tag:add tea=ід tag=тег")
	printTabulated("Імпорт чаю:", "tea:import шлях до файлу")
	printTabulated("Заварювання:", "tea:brew t=ід чаю w=маса")
	return 0
}
func tea_import_cmd(args *[]string) int {
	if len(*args) == 0 {
		return NOT_ENOUGH_ARUMENTS
	}

	filepath := (*args)[0]

	f, err := os.Open(filepath)

	if err != nil {
		printErr("Не вдалося відкрити файл " + filepath)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		printErr("Помилка при зчитуванні файлу:" + err.Error())
	}

	for idx, line := range records {
		if idx == 0 {
			continue
		}

		var dataFormatted []string

		for idx, datum := range line {
			switch idx {
			case 0: // name
				dataFormatted = append(dataFormatted, "n="+datum)
			case 1: // year
				dataFormatted = append(dataFormatted, "y="+datum)
			case 2: // stock
				dataFormatted = append(dataFormatted, "s="+datum)
			case 3: // price
				dataFormatted = append(dataFormatted, "p="+datum)
			case 4: // tags
				dataFormatted = append(dataFormatted, "t="+datum)
			}
		}

		tea_add_cmd(&dataFormatted)
	}

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
