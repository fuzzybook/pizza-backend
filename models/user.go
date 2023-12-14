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

type User struct {
	ID          int              `json:"id" gorm:"primaryKey"`
	Email       string           `json:"email"  gorm:"unique"`
	Name        string           `json:"name"`
	Password    string           `json:"password"`
	Roles       UserRoles        `json:"roles" gorm:"type:text"`
	Status      UserStatus       `json:"status" gorm:"type:text"`
	Types       UserType         `json:"types"  gorm:"type:text"`
	ActivatedAt *time.Time       `json:"activatedAt"`
	UUID        string           `json:"uuid"`
	Details     *UserDetails     `json:"details" `
	Preferences *UserPreferences `json:"preferences" `
	Avatar      *string          `json:"avatar" `
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt   `gorm:"index"`
}

type UserPreferences struct {
	ID               int            `json:"id" gorm:"primaryKey"`
	UserId           int            `json:"userId"`
	UseIdle          bool           `json:"useIdle"`
	IdleTimeout      int            `json:"idleTimeout"`
	UseIdlePassword  bool           `json:"useIdlePassword"`
	IdlePin          string         `json:"idlePin"`
	UseDirectLogin   bool           `json:"useDirectLogin"`
	UseQuadcodeLogin bool           `json:"useQuadcodeLogin"`
	SendNoticesMail  bool           `json:"sendNoticesMail"`
	Language         string         `json:"language"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
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
	FiredAt       time.Time      `json:"firedAt"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type NewUser struct {
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Roles    []UserRole `json:"roles"`
	Type     UserType   `json:"type"`
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

type LogoutResult struct {
	Ok bool `json:"ok"`
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

func Logout(ctx context.Context) (*LogoutResult, error) {
	context := common.GetContext(ctx)
	var user *User
	var claimUser *ClaimUser
	if claimUser = ForContext(ctx); claimUser == nil {
		return &LogoutResult{Ok: true}, nil
	}

	err := context.Database.Where("id = ?", claimUser.Id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	// create session
	session := &Session{}
	err = context.Database.Where("user_id = ?", user.ID).Find(&session).Error
	if err != nil {
		return &LogoutResult{Ok: true}, nil
	}

	session.FiredAt = time.Now()

	return &LogoutResult{Ok: true}, nil
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
		return nil, fmt.Errorf("user not found %s", input.Email)
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

func GetCompleteUserById(ctx context.Context, userId int) (*User, error) {
	context := common.GetContext(ctx)
	user := &User{}
	err := context.Database.Preload("Details").Preload("Preferences").Where("id = ?", userId).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
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

func GetUserSession(ctx context.Context, id int) (*Session, error) {
	context := common.GetContext(ctx)
	var session *Session
	err := context.Database.Where("id = ?", id).Find(&session).Error
	if err != nil {
		return nil, err
	}
	return session, nil
}

func Promos(ctx context.Context) ([]*MenuItem, error) {
	context := common.GetContext(ctx)
	menuItems := []*MenuItem{}
	err := context.Database.Where("promo = true").Find(&menuItems).Error
	if err != nil {
		return nil, err
	}
	return menuItems, nil
}
