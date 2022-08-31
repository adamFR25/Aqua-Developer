package repository

import (
	"tesla/Entity"
	"log"

	"gorm.io/gorm"
)

type RepoStudent struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *RepoStudent {
	return &RepoStudent{db}
}

func (r *RepoStudent) GetStudent() ([]entity.Student, error) {
	var users []entity.Student

	result := r.db.Order("id asc").Find(&users)

	err := result.Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (r *RepoStudent) GetById(id int) (entity.Student, error) {
	var user entity.Student

	result := r.db.First(&user, id)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *RepoStudent) CreateStudent(name string, email string) (entity.Student, error) {
	user := entity.Student{Name: name, Email: email}

	result := r.db.Create(&user)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *RepoStudent) UpdateStudent(name string, email string) (entity.Student, error) {
	var user entity.Student

	r.db.First(&user, user.ID)

	user.Name = name
	user.Email = email

	result := r.db.Updates(&user)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *RepoStudent) DeleteStudent(id int) (entity.Student, error) {
	var user entity.Student

	r.db.First(&user, id)

	result := r.db.Delete(&user)

	err := result.Error

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}