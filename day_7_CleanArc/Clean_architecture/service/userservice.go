package service

import (
	"tesla/Entity"
	"tesla/repository"
	"log"
)

type ServiceStudent struct {
	repository1 repository.RepoStudent
}

func NewUserService(repository1 repository.RepoStudent) *ServiceStudent {
	return &ServiceStudent{repository1}
}

func (s *ServiceStudent) GetStudentService() ([]entity.Student, error) {
	users, err := s.repository1.GetStudent()

	if err != nil {
		log.Println("Get Users Service Error", err)
		return nil, err
	}

	return users, nil
}

func (s *ServiceStudent) GetStudentByIdService(id int) (entity.Student, error) {
	user, err := s.repository1.GetById(id)

	if err != nil {
		log.Println("Get User By Id Service Error", err)
		return user, err
	}

	return user, nil
}

func (s *ServiceStudent) CreateStudentService(name string, email string) (entity.Student, error) {
	user, err := s.repository1.CreateStudent(name, email)

	if err != nil {
		log.Println("Create User Service Error", err)
		return user, err
	}

	return user, nil
}

func (s *ServiceStudent) UpdateStudentService(id int, name string, email string) (entity.Student, error) {
	user, err := s.repository1.UpdateStudent(name, email)

	if err != nil {
		log.Println("Update Student Error", err)
		return user, err
	}

	return user, nil
}

func (s *ServiceStudent) DeleteStudentService(id int) (entity.Student, error) {
	user, err := s.repository1.DeleteStudent(id)

	if err != nil {
		log.Println("Delete User Error", err)
		return user, err
	}

	return user, nil
}