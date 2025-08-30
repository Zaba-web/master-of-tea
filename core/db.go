package core

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const DB_NAME = "tea.db"

var db, _ = gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})

type Tea struct {
	gorm.Model
	Name         string
	Year         int
	Stock        float32
	PricePerGram float32
	Tags         []*Tag `gorm:"many2many:tea_tags;"`
}

type Tag struct {
	gorm.Model
	Name  string `gorm:"index"`
	Color uint8
	Teas  []*Tea `gorm:"many2many:tea_tags;"`
}

type Brew struct {
	gorm.Model
	TeaId  uint
	Tea    Tea
	Weight float32
}

func init_db() {
	db.AutoMigrate(&Tea{}, &Tag{}, &Brew{})
}

func saveTea(t *Tea) {
	result := db.Save(&t)
	result.Last(&t)
}

func assignTags(t *Tea, tagNames *[]string) {
	for _, tagName := range *tagNames {
		tag := &Tag{}
		db.Where("name = ?", tagName).Find(tag)

		if tag.ID == 0 {
			tag.Name = strings.TrimSpace(strings.ToLower(tagName))
			tag.Color = uint8(rand.IntN(7-0) + 0)

			result := db.Save(&tag)
			result.Last(&tag)
		}

		t.Tags = append(t.Tags, tag)
		db.Save(t)
	}
}

func getNumberOfTeas() int64 {
	var count int64
	db.Model(&Tea{}).Count(&count)

	return count
}

func getTotalWeight() float32 {
	var weight float32
	db.Model(&Tea{}).Select("SUM(stock)").Row().Scan(&weight)

	return weight
}

func getTotalPrice() float32 {
	var teas []Tea
	db.Find(&teas)

	var totalPrice float32 = 0.00

	for _, tea := range teas {
		totalPrice += tea.Stock * tea.PricePerGram
	}

	return totalPrice
}

func getAllTeas() []Tea {
	var teas []Tea
	db.Model(&Tea{}).Preload("Tags").Find(&teas)

	return teas
}

func GetTeasWithFilters(filters [][]string, sorting []string) []Tea {
	var teas []Tea
	query := db.Model(&Tea{})

	for _, filter := range filters {
		if numVal, ok := ParseNumber(filter[1]); ok {
			query.Where(fmt.Sprintf("%s %d ?", filter[0], numVal), filter[2])
		} else {
			query.Where(fmt.Sprintf("%s %s ?", filter[0], filter[1]), filter[2])
		}
	}

	if len(sorting) > 1 {
		query.Order(fmt.Sprintf("%s %s", sorting[0], sorting[1]))
	}

	query.Preload("Tags").Find(&teas)

	return teas
}

func deleteTea(id int) {
	db.Delete(&Tea{}, id)
}

func ParseNumber(s string) (any, bool) {
	// try int first
	if i, err := strconv.Atoi(s); err == nil {
		return i, true
	}
	// then try float
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f, true
	}

	return nil, false
}
