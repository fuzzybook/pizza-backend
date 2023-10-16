package models

import (
	"context"
	"fmt"
	"pizza-backend/common"
	"pizza-backend/jwt"

	"strconv"

	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userId"`
	Title     string    `json:"title"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	ZipCode   string    `json:"zipCode"`
	Country   string    `json:"country"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserBio struct {
	ID        int         `json:"id" gorm:"primaryKey"`
	UserID    int         `json:"userId"`
	Avatar    string      `json:"avatar"`
	Bio       string      `json:"bio" gorm:"type:text"`
	Socials   UserSocials `json:"socials" gorm:"type:text"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

type UserPro struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userId"`
	Ipi       string    `json:"ipi"`
	Promember string    `json:"promember"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserBank struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"userId"`
	Name        string    `json:"name"`
	Bicswift    string    `json:"bicswift"`
	Account     string    `json:"account"`
	AccountName string    `json:"accountName"`
	Iban        string    `json:"iban"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type User struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Email       string         `json:"email"  gorm:"unique"`
	Name        string         `json:"name"`
	Password    string         `json:"password"`
	Roles       UserRoles      `json:"roles" gorm:"type:text"`
	Status      UserStatus     `json:"status" gorm:"type:text"`
	Types       UserType       `json:"types"  gorm:"type:text"`
	ActivatedAt *time.Time     `json:"activatedAt"`
	UUID        string         `json:"uuid"`
	PastelleID  string         `json:"pastelleid"`
	Details     *UserDetails   `json:"details" `
	Bio         *UserBio       `json:"bio" `
	Pro         *UserPro       `json:"pro" `
	Bank        *UserBank      `json:"bank" `
	Session     *Session       `json:"session" `
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type UserBridge struct {
	ID        int `json:"id" gorm:"primaryKey"`
	UserId    int `json:"userId"`
	OldUserId int `json:"oldUserId"`
}

type Session struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	UserID        int            `json:"userId"`
	RecoveryToken string         `json:"recoveryToken"`
	Roles         string         `json:"roles"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type NewUser struct {
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Roles    []UserRole `json:"roles"`
	Type     UserType   `json:"types"`
	Password string     `json:"password"`
}

type UpdateUserRoles struct {
	UserId int        `json:"userId"`
	Roles  []UserRole `json:"roles"`
}

type UpdateUserPassword struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPages struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"pagesize"`
	SortBy     string `json:"sortBy"`
	Descending bool   `json:"descending"`
}

// user actions

func CreateUser(ctx context.Context, input NewUser) (*User, error) {
	context := common.GetContext(ctx)
	user := &User{
		Email:    input.Email,
		Name:     input.Name,
		Roles:    input.Roles,
		Types:    input.Type,
		Password: HashPassword(input.Password),
	}
	err := context.Database.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Login(ctx context.Context, input UserLogin) (*Session, error) {
	context := common.GetContext(ctx)
	var user *User
	err := context.Database.Where("email = ?", input.Email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, fmt.Errorf("user not foundr %s", input.Email)
	}
	// create session
	session := &Session{}
	err = context.Database.Where("user_id = ?", user.ID).Find(&session).Error
	if err != nil {
		return nil, fmt.Errorf("user not foundr %s", input.Email)
	}

	if err := ComparePassword(user.Password, input.Password); err != nil {
		return nil, err
	}

	token, err := jwt.JwtGenerate(ctx, strconv.Itoa(user.ID))
	if err != nil {
		return nil, err
	}
	session.RecoveryToken = token
	if session.ID == 0 {
		session.UserID = user.ID
		err = context.Database.Create(&session).Error
		if err != nil {
			return nil, err
		}
		return session, nil
	}
	err = context.Database.Save(&session).Error
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (u *User) GetById(db *gorm.DB, id int) error {
	err := db.Where("id = ?", id).Find(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(ctx context.Context, userId int) (*User, error) {
	context := common.GetContext(ctx)
	user := &User{}
	err := context.Database.Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateRoles(ctx context.Context, input UpdateUserRoles) (*User, error) {
	context := common.GetContext(ctx)
	user := &User{}
	err := context.Database.Where("id = ?", input.UserId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	user.Roles = input.Roles
	err = context.Database.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdatePassword(ctx context.Context, input UpdateUserPassword) (*User, error) {
	context := common.GetContext(ctx)
	user := &User{}

	err := context.Database.Where("id = ?", input.UserId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	user.Password = HashPassword(input.Password)
	err = context.Database.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserDetails(ctx context.Context, userId int) (*UserDetails, error) {
	context := common.GetContext(ctx)
	var details *UserDetails
	err := context.Database.Where("id = ?", userId).Find(&details).Error
	if err != nil {
		return nil, err
	}
	if details.ID == 0 {
		return nil, nil
	}
	return details, nil
}

func GetUserBio(ctx context.Context, userId int) (*UserBio, error) {
	context := common.GetContext(ctx)
	var bio *UserBio
	err := context.Database.Where("id = ?", userId).Find(&bio).Error
	if err != nil {
		return nil, err
	}
	if bio.ID == 0 {
		return nil, nil
	}
	return bio, nil
}

func GetUserSession(ctx context.Context, id int) (*Session, error) {
	context := common.GetContext(ctx)
	var session *Session
	err := context.Database.Where("id = ?", id).Find(&session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}

func GetUserPro(ctx context.Context, proId int) (*UserPro, error) {
	context := common.GetContext(ctx)
	var pro *UserPro
	err := context.Database.Where("id = ?", proId).Find(&pro).Error
	if err != nil {
		return nil, err
	}
	if pro.ID == 0 {
		return nil, nil
	}
	return pro, nil
}

func GetUserBank(ctx context.Context, bankId int) (*UserBank, error) {
	context := common.GetContext(ctx)
	var bank *UserBank
	err := context.Database.Where("id = ?", bankId).Find(&bank).Error
	if err != nil {
		return nil, err
	}
	if bank.ID == 0 {
		return nil, nil
	}
	return bank, nil
}
