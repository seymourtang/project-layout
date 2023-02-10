package repository

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type StudentDTO struct {
	ID   string
	Name string
	Age  int
	Addr string
}

var _ Student = (*studentImpl)(nil)

type studentImpl struct {
	db *gorm.DB
}

func NewStudentImpl(db *gorm.DB) *studentImpl {
	return &studentImpl{db: db}
}

func (s *studentImpl) Get(ctx context.Context, ID string) (*StudentDTO, error) {
	log.Println("student Impl")
	return &StudentDTO{}, nil
}
