package service

import (
	"time"

	"go-user-api/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"

	"context"

	db "go-user-api/db/sqlc"
)

type UserService struct {
	Repo *repository.UserRepository
}
type UserServiceInterface interface {
	CreateUser(name, dob string) (db.User, error)
	GetUserWithAge(id int32) (map[string]interface{}, error)
	ListUsers(page, limit int) ([]map[string]interface{}, error)
	UpdateUser(id int32, name string, dob string) (map[string]interface{}, error)
	DeleteUser(id int32) error
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}
func pgDateToTime(d pgtype.Date) time.Time {
	return d.Time
}

// calculate age from DOB
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func (s *UserService) GetUserWithAge(id int32) (map[string]interface{}, error) {
	user, err := s.Repo.GetUser(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	dobTime := pgDateToTime(user.Dob)

	return map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
		"dob":  dobTime.Format("2006-01-02"),
		"age":  calculateAge(dobTime),
	}, nil
}

func (s *UserService) CreateUser(name string, dobStr string) (db.User, error) {

	dob, err := parseDOB(dobStr)
	if err != nil {
		return db.User{}, err
	}
	return s.Repo.CreateUser(context.TODO(), name, dob)
}
func parseDOB(dateStr string) (pgtype.Date, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return pgtype.Date{}, err
	}

	return pgtype.Date{
		Time:  t,
		Valid: true,
	}, nil
}

func (s *UserService) ListUsers(page, limit int) ([]map[string]interface{}, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, err := s.Repo.ListUsersPaginated(
		context.TODO(),
		db.ListUsersPaginatedParams{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}

	for _, u := range users {

		dobTime := pgDateToTime(u.Dob)

		result = append(result, map[string]interface{}{
			"id":   u.ID,
			"name": u.Name,
			"dob":  dobTime.Format("2006-01-02"),
			"age":  calculateAge(dobTime),
		})
	}

	return result, nil
}

func (s *UserService) UpdateUser(id int32, name string, dobStr string) (map[string]interface{}, error) {

	// get existing user first
	user, err := s.Repo.GetUser(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	// default values (keep old if not provided)
	if name == "" {
		name = user.Name
	}

	dob := user.Dob

	if dobStr != "" {
		parsedDob, err := parseDOB(dobStr)
		if err != nil {
			return nil, err
		}
		dob = parsedDob
	}

	// update DB
	updatedUser, err := s.Repo.UpdateUser(context.TODO(), id, name, dob)
	if err != nil {
		return nil, err
	}

	// response with age
	dobTime := pgDateToTime(updatedUser.Dob)

	return map[string]interface{}{
		"id":   updatedUser.ID,
		"name": updatedUser.Name,
		"dob":  dobTime.Format("2006-01-02"),
		"age":  calculateAge(dobTime),
	}, nil
}

func (s *UserService) DeleteUser(id int32) error {

	// check if user exists first (optional but good practice)
	_, err := s.Repo.GetUser(context.TODO(), id)
	if err != nil {
		return err
	}

	return s.Repo.DeleteUser(context.TODO(), id)
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
