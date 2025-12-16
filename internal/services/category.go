package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/marcelobracet/grpc/internal/database"
	"github.com/marcelobracet/grpc/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

// DeleteCategory implements pb.CategoryServiceServer.
func (s *CategoryService) DeleteCategory(context.Context, *pb.DeleteCategoryRequest) (*pb.Category, error) {
	panic("unimplemented")
}

// GetCategory implements pb.CategoryServiceServer.
func (s *CategoryService) GetCategory(context.Context, *pb.GetCategoryRequest) (*pb.Category, error) {
	panic("unimplemented")
}

// ListCategories implements pb.CategoryServiceServer.
func (s *CategoryService) ListCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {
	rows, err := s.CategoryDB.ListCategories()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CategoryList{
		Categories: rows,
	}, nil
}

func (s *CategoryService) ListCategoriesByName(ctx context.Context, req *pb.ListCategoriesByNameRequest) (*pb.CategoryList, error) {
	categories, err := s.CategoryDB.ListCategoriesByName(req.Name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.CategoryList{
		Categories: categories,
	}, nil
}

// UpdateCategory implements pb.CategoryServiceServer.
func (s *CategoryService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.Category, error) {
	err := s.CategoryDB.UpdateCategory(req.Id, req.Name, req.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Category{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
	}, nil
}

// mustEmbedUnimplementedCategoryServiceServer implements pb.CategoryServiceServer.
func (s *CategoryService) mustEmbedUnimplementedCategoryServiceServer() {
	panic("unimplemented")
}

func NewCategory(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.Category, error) {
	err := s.CategoryDB.CreateCategory(req.Name, req.Description)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Category{
		Id:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
	}, nil
}
