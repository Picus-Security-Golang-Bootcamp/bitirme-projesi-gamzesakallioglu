package productCategory

import (
	"context"
	"errors"

	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/internal/api"
	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/internal/models"
	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/pkg/pagination"
)

type productCategoryService struct {
	repo Repository
}

type Service interface {
	CreateProductCategoryForBulk(ctx context.Context, category *api.ProductCategory) error
	GetProductCategoryByName(ctx context.Context, name *string) (*models.ProductCategory, error)
	GetAllProductCategories(ctx context.Context, page int, pageSize int) (*pagination.Pagination, error)
}

func NewProductCategoryService(repo Repository) Service {
	if repo == nil {
		return nil
	}

	return &productCategoryService{repo: repo}
}

func (p productCategoryService) CreateProductCategoryForBulk(ctx context.Context, category *api.ProductCategory) error {

	categoryToGo := models.ProductCategory{Name: category.Name, Description: category.Description, IsParent: category.IsParent, ParentID: category.ParentID}
	err := p.repo.CreateProductCategoryForBulk(ctx, &categoryToGo)
	if err != nil {
		return err
	}

	return nil
}

func (p productCategoryService) GetProductCategoryByName(ctx context.Context, name *string) (*models.ProductCategory, error) {
	if len(*name) < 1 {
		return nil, errors.New("name cannot be empty")
	}

	category, err := p.repo.GetProductCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (p productCategoryService) GetAllProductCategories(ctx context.Context, page int, pagesize int) (*pagination.Pagination, error) {

	offset := (page - 1) * pagesize
	sorting := "ID desc"
	productCategories, count := p.repo.GetAllProductCategories(ctx, pagesize, offset, sorting)
	productCategoriesApi := ProductCategoriesToResponses(productCategories)

	pagination := pagination.NewPagination(page, pagesize, count, sorting)
	pagination.Rows = productCategoriesApi

	return pagination, nil
}
