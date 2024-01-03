package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"pizza-backend/common"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type MenuIngredient struct {
	ID         int               `json:"id" gorm:"primaryKey"`
	Text       string            `json:"text"`
	Available  int               `json:"available"`
	Priority   int               `json:"priority"`
	Categories MenuDataCategoies `json:"categories" gorm:"type:text"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
}

type MenuDough struct {
	ID         int               `json:"id" gorm:"primaryKey"`
	Text       string            `json:"text"`
	Available  int               `json:"available"`
	Priority   int               `json:"priority"`
	Categories MenuDataCategoies `json:"categories" gorm:"type:text"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
}

type MenuCondiment struct {
	ID         int               `json:"id" gorm:"primaryKey"`
	Text       string            `json:"text"`
	Available  int               `json:"available"`
	Priority   int               `json:"priority"`
	Categories MenuDataCategoies `json:"categories" gorm:"type:text"`
	CreatedAt  time.Time         `json:"createdAt"`
	UpdatedAt  time.Time         `json:"updatedAt"`
}

type UpdateIngredient struct {
	ID         int               `json:"id"`
	Available  int               `json:"available"`
	Priority   int               `json:"priority"`
	Categories MenuDataCategoies `json:"categories" `
	Text       string            `json:"text"`
}
type DeleteIngredient struct {
	ID int `json:"id"`
}

type UpdateDough struct {
	ID         int               `json:"id"`
	Available  int               `json:"available"`
	Priority   int               `json:"priority"`
	Categories MenuDataCategoies `json:"categories" gorm:"type:text"`
	Text       string            `json:"text"`
}
type DeleteDough struct {
	ID int `json:"id"`
}

type UpdateCondiment struct {
	ID         int               `json:"id"`
	Available  int               `json:"available"`
	Priority   int               `json:"priority"`
	Categories MenuDataCategoies `json:"categories" gorm:"type:text"`
	Text       string            `json:"text"`
}
type DeleteCondiment struct {
	ID int `json:"id"`
}

type UpdateCategory struct {
	ID       int    `json:"id"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Image    string `json:"image"`
}

type DeleteCategory struct {
	ID int `json:"id"`
}

type UpdateCategoryItem struct {
	ID            int    `json:"id"`
	CategoryRefer int    `json:"CategoryRefer"`
	Data          string `json:"data"`
}

type MenuCategory struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Uuid      string     `json:"uuid"`
	Category  string     `json:"category"`
	Title     string     `json:"title"`
	Image     string     `json:"image"`
	Items     []MenuItem `json:"items" gorm:"foreignKey:CategoryRefer"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type MenuItem struct {
	ID            int    `json:"id" gorm:"primaryKey"`
	CategoryRefer int    `json:"CategoryRefer"`
	Uuid          string `json:"uuid"`
	Promo         bool   `json:"promo"`
	// Dirty         bool         `json:"dirty"`
	Data MenuItemData `json:"data" gorm:"type:bytes;serializer:gob"`
}

type MenuItemData struct {
	Doughs      []string    `json:"doughs" gorm:"type:bytes;serializer:gob"`
	Condiments  []string    `json:"condiments" gorm:"type:bytes;serializer:gob"`
	Ingredients []string    `json:"ingredients" gorm:"type:bytes;serializer:gob"`
	Extra       []MenuExtra `json:"extra" gorm:"type:bytes;serializer:gob"`
	Title       string      `json:"title"`
	Text        string      `json:"text"`
	Image       string      `json:"image"`
	Price       string      `json:"price"`
}
type MenuExtra struct {
	Take        bool   `json:"take"`
	Text        string `json:"text"`
	Price       string `json:"price"`
	Alternative string `json:"alternative"`
}

// The `File` type, represents the response of uploading a file.
type File struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
}

// The `UploadFile` type, represents the request for uploading a file with certain payload.
type UploadFile struct {
	ID     int            `json:"id"`
	Domain string         `json:"domain"`
	Uuid   string         `json:"uuid"`
	File   graphql.Upload `json:"file"`
}

type SaveImage struct {
	Id    int    `json:"id"`
	Uuid  string `json:"uuid"`
	Image string `json:"image"`
}

type MenuOrder struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Data      string `json:"data"`
	Confirmed bool
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MenuDataIngredient struct {
	Take bool   `json:"take"`
	Text string `json:"text"`
}

type MenuDataResult struct {
	Uid         string               `json:"uid"`
	Selected    bool                 `json:"selected"`
	Tipo        string               `json:"tipo"`
	Title       string               `json:"titlr"`
	Price       string               `json:"price"`
	Quantity    int                  `json:"quantity"`
	Total       int                  `json:"total"`
	Dought      string               `json:"dought"`
	Extra       []MenuExtra          `json:"extra"`
	Ingredients []MenuDataIngredient `json:"ingredients"`
	Condiments  []MenuDataIngredient `json:"condiments"`
}

type MenuTimes struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Times     string    `json:"times"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StringArray []string

func (e *StringArray) Scan(value interface{}) error {
	s := strings.Split(value.(string), ",")
	*e = []string{}
	for _, v := range s {
		*e = append(*e, v)
	}
	return nil
}

func (j StringArray) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return strings.Join(j.Strings(), ","), nil
}

func (e StringArray) Strings() []string {
	result := []string{}
	for _, v := range e {
		result = append(result, v)
	}
	return result
}

func DeleteMenuIngredient(ctx context.Context, input DeleteIngredient) (bool, error) {
	context := common.GetContext(ctx)
	context.Database.Delete(&MenuIngredient{}, input.ID)
	return true, nil
}

func UpdateMenuIngredient(ctx context.Context, input UpdateIngredient) (*MenuIngredient, error) {
	context := common.GetContext(ctx)
	menuIngredient := MenuIngredient{}

	if input.ID == 0 {
		menuIngredient.Text = input.Text
		menuIngredient.Priority = input.Priority
		menuIngredient.Available = 1
		menuIngredient.Categories = input.Categories

		err := context.Database.Save(&menuIngredient).Error
		if err != nil {
			return nil, err
		}
		return &menuIngredient, nil
	}

	err := context.Database.Where("id = ?", input.ID).Find(&menuIngredient).Error
	if err != nil {
		return nil, fmt.Errorf("%d not fouund", input.ID)
	}
	if menuIngredient.ID == 0 {
		return nil, fmt.Errorf("%d not fouund", input.ID)
	}
	menuIngredient.Categories = input.Categories
	menuIngredient.Text = input.Text
	menuIngredient.Priority = input.Priority
	menuIngredient.Available = input.Available
	err = context.Database.Save(&menuIngredient).Error
	if err != nil {
		return nil, err
	}
	return &menuIngredient, nil
}

func UpdateMenuIngredients(ctx context.Context, input []*UpdateIngredient) ([]*MenuIngredient, error) {
	context := common.GetContext(ctx)
	result := []*MenuIngredient{}

	for _, v := range input {
		_, err := UpdateMenuIngredient(ctx, *v)
		if err != nil {
			return nil, err
		}
	}

	err := context.Database.Order("priority").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteMenuDough(ctx context.Context, input DeleteDough) (bool, error) {
	context := common.GetContext(ctx)
	context.Database.Delete(&MenuDough{}, input.ID)
	return true, nil
}

func UpdateMenuDough(ctx context.Context, input UpdateDough) (*MenuDough, error) {
	context := common.GetContext(ctx)
	menuDough := MenuDough{}

	if input.ID == 0 {
		menuDough.Text = input.Text
		menuDough.Priority = input.Priority
		menuDough.Available = 1
		menuDough.Categories = input.Categories
		err := context.Database.Save(&menuDough).Error
		if err != nil {
			return nil, err
		}
		return &menuDough, nil
	}

	err := context.Database.Where("id = ?", input.ID).Find(&menuDough).Error
	if err != nil {
		return nil, fmt.Errorf("%d not fouund", input.ID)
	}
	if menuDough.ID == 0 {
		return nil, fmt.Errorf("%d not fouund", input.ID)
	}
	menuDough.Categories = input.Categories
	menuDough.Text = input.Text
	menuDough.Priority = input.Priority
	menuDough.Available = input.Available
	err = context.Database.Save(&menuDough).Error
	if err != nil {
		return nil, err
	}
	return &menuDough, nil
}

func UpdateMenuDoughs(ctx context.Context, input []*UpdateDough) ([]*MenuDough, error) {
	context := common.GetContext(ctx)
	result := []*MenuDough{}

	for _, v := range input {
		_, err := UpdateMenuDough(ctx, *v)
		if err != nil {
			return nil, err
		}
	}

	err := context.Database.Order("priority").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteMenuCondiment(ctx context.Context, input DeleteCondiment) (bool, error) {
	context := common.GetContext(ctx)
	context.Database.Delete(&MenuCondiment{}, input.ID)
	return true, nil
}

func UpdateMenuCondiment(ctx context.Context, input UpdateCondiment) (*MenuCondiment, error) {
	context := common.GetContext(ctx)
	menuCondiment := MenuCondiment{}

	if input.ID == 0 {
		menuCondiment.Text = input.Text
		menuCondiment.Priority = input.Priority
		menuCondiment.Available = 1
		menuCondiment.Categories = input.Categories
		err := context.Database.Save(&menuCondiment).Error
		if err != nil {
			return nil, err
		}
		return &menuCondiment, nil
	}
	err := context.Database.Where("id = ?", input.ID).Find(&menuCondiment).Error
	if err != nil {
		return nil, err
	}
	if menuCondiment.ID == 0 {
		return nil, fmt.Errorf("%d not fouund", input.ID)
	}
	menuCondiment.Categories = input.Categories
	menuCondiment.Text = input.Text
	menuCondiment.Priority = input.Priority
	menuCondiment.Available = input.Available
	err = context.Database.Save(&menuCondiment).Error
	if err != nil {
		return nil, err
	}
	return &menuCondiment, nil
}

func UpdateMenuCondiments(ctx context.Context, input []*UpdateCondiment) ([]*MenuCondiment, error) {
	context := common.GetContext(ctx)
	result := []*MenuCondiment{}

	for _, v := range input {
		_, err := UpdateMenuCondiment(ctx, *v)
		if err != nil {
			return nil, err
		}
	}

	err := context.Database.Order("priority").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateMenuCategory(ctx context.Context, input *UpdateCategory) (*MenuCategory, error) {
	context := common.GetContext(ctx)
	menuCategory := MenuCategory{}

	err := context.Database.Where("id = ?", input.ID).Find(&menuCategory).Error
	if err != nil {
		return nil, err
	}

	menuCategory.Category = input.Category
	menuCategory.Title = input.Title
	menuCategory.Image = input.Image

	if menuCategory.ID == 0 {
		menuCategory.Uuid = uuid.NewString()
		err = context.Database.Save(&menuCategory).Error
		if err != nil {
			return nil, err
		}
		return &menuCategory, nil
	}
	if menuCategory.Uuid == "" {
		menuCategory.Uuid = uuid.NewString()
	}
	context.Database.Model(&menuCategory).Updates(MenuCategory{Uuid: menuCategory.Uuid, Category: input.Category, Title: input.Title, Image: input.Image})
	return &menuCategory, nil
}

func GetMenuCategories(ctx context.Context, obj *Menu) ([]*MenuCategory, error) {
	context := common.GetContext(ctx)
	menuCategories := []*MenuCategory{}
	err := context.Database.Preload(clause.Associations).Preload("Items." + clause.Associations).Find(&menuCategories).Error
	if err != nil {
		return nil, err
	}
	return menuCategories, nil
}

func UpdateMenuCategoryItem(ctx context.Context, input *UpdateCategoryItem) (*MenuItem, error) {
	context := common.GetContext(ctx)
	menuItem := MenuItem{}

	err := context.Database.Where("id = ? AND category_refer = ?", input.ID, input.CategoryRefer).Find(&menuItem).Error
	if err != nil {
		return nil, err
	}

	if menuItem.ID == 0 {
		if input.CategoryRefer == 0 {
			return nil, fmt.Errorf("wong Category ID")
		}

		m := make(map[string]string)
		menuItem.Uuid = uuid.NewString()
		if err := json.Unmarshal([]byte(input.Data), &m); err != nil {
			menuItem.Data.Doughs = []string{}
			menuItem.Data.Condiments = []string{}
			menuItem.Data.Ingredients = []string{}
			menuItem.Data.Extra = []MenuExtra{}
			menuItem.Data.Title = ""
			menuItem.Data.Text = ""
			menuItem.Data.Image = ""
			menuItem.Data.Price = "0.00"
		} else {
			menuItem.Data.Title = m["title"]
			menuItem.Data.Text = m["text"]
			menuItem.Data.Image = m["image"]
			menuItem.Data.Price = m["price"]
		}

		menuItem.CategoryRefer = input.CategoryRefer
		err = context.Database.Save(&menuItem).Error
		if err != nil {
			return nil, err
		}
		return &menuItem, nil
	}

	if err := json.Unmarshal([]byte(input.Data), &menuItem); err != nil {
		fmt.Println(err)
		return &MenuItem{}, err
	}
	if menuItem.Uuid == "" {
		menuItem.Uuid = uuid.NewString()
	}
	context.Database.Model(&menuItem).Updates(&menuItem)
	return &menuItem, nil
}

func DeleteMenuCategoryItem(ctx context.Context, input *UpdateCategoryItem) (bool, error) {
	context := common.GetContext(ctx)
	menuItem := MenuItem{}
	err := context.Database.Where("id = ? AND category_refer = ?", input.ID, input.CategoryRefer).Find(&menuItem).Error
	if err != nil {
		return false, err
	}
	if menuItem.ID == 0 {
		return false, nil
	}
	context.Database.Where("id = ? AND category_refer = ?", input.ID, input.CategoryRefer).Delete(&menuItem)
	return true, nil
}

func DeleteMenuCategory(ctx context.Context, input DeleteCategory) (bool, error) {
	context := common.GetContext(ctx)
	menuCategory := MenuCategory{}
	err := context.Database.Where("id = ? ", input.ID).Find(&menuCategory).Error
	if err != nil {
		return false, err
	}
	if menuCategory.ID == 0 {
		return false, nil
	}
	context.Database.Where("id = ? ", input.ID).Delete(&menuCategory)
	return true, nil
}

func UpdateMenuItemPromo(ctx context.Context, promo bool, id int) (bool, error) {
	context := common.GetContext(ctx)
	menuItem := MenuItem{}

	err := context.Database.Where("id = ?", id).Find(&menuItem).Error
	if err != nil {
		return false, err
	}

	if menuItem.ID == 0 {
		return false, fmt.Errorf("wong Category ID")
	}

	menuItem.Promo = promo

	err = context.Database.Save(&menuItem).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

type __Time struct {
	Active       bool
	Start        int
	End          int
	AfterTime    int
	HasAfterTime bool
}

type TimeZone = struct {
	Id     int
	Label  string
	Height int
	Start  int
	End    int
	Times  map[string]__Time
}

func SaveTimes(ctx context.Context, times string) (bool, error) {
	context := common.GetContext(ctx)
	menuTimes := MenuTimes{}

	err := context.Database.Where("id = ?", 1).Find(&menuTimes).Error
	if err != nil {
		return false, err
	}

	menuTimes.Times = times
	err = context.Database.Save(&menuTimes).Error
	if err != nil {
		return false, err
	}
	return true, nil

}

func TodayTimes(ctx context.Context) (string, error) {
	context := common.GetContext(ctx)
	menuTimes := MenuTimes{}

	err := context.Database.Where("id = ?", 1).Find(&menuTimes).Error
	if err != nil {
		return "", err
	}

	if menuTimes.ID == 0 {
		return "", nil
	}

	x := map[int]TimeZone{}
	json.Unmarshal([]byte(menuTimes.Times), &x)
	fmt.Println(x)

	x2 := map[int]TimeZone{}

	for t := 1; t < len(x)+1; t++ {
		tz := TimeZone{}
		tz.Id = x[t].Id
		tz.Start = x[t].Start
		tz.End = x[t].End
		tz.Label = x[t].Label
		tz.Times = make(map[string]__Time)
		x2[t] = tz

	}

	for i := 0; i < 2; i++ {
		year, month, day := time.Now().AddDate(0, 0, i).Date()
		date := fmt.Sprintf("%d-%02d-%02d", year, int(month), day)
		for t := 1; t < len(x); t++ {
			val, ok := x[t].Times[date]
			if ok {
				x2[t].Times[date] = val
			}
		}
	}

	result, err := json.Marshal(x2)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func WeekTimes(ctx context.Context) (string, error) {
	context := common.GetContext(ctx)
	menuTimes := MenuTimes{}

	err := context.Database.Where("id = ?", 1).Find(&menuTimes).Error
	if err != nil {
		return "", err
	}

	if menuTimes.ID == 0 {
		return "", nil
	}

	x := map[int]TimeZone{}
	json.Unmarshal([]byte(menuTimes.Times), &x)
	fmt.Println(x)

	x2 := map[int]TimeZone{}

	for t := 1; t < len(x)+1; t++ {
		tz := TimeZone{}
		tz.Id = x[t].Id
		tz.Start = x[t].Start
		tz.End = x[t].End
		tz.Label = x[t].Label
		tz.Times = make(map[string]__Time)
		x2[t] = tz

	}

	for i := 0; i < 7; i++ {
		year, month, day := time.Now().AddDate(0, 0, i).Date()
		date := fmt.Sprintf("%d-%02d-%02d", year, int(month), day)
		for t := 1; t < len(x); t++ {
			val, ok := x[t].Times[date]
			if ok {
				x2[t].Times[date] = val
			}
		}
	}

	return menuTimes.Times, nil
}
