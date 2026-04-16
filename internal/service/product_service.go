package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
	rdb         *redis.Client
}

func NewProductService(repo *repository.ProductRepository, rdb *redis.Client) *ProductService {
	return &ProductService{
		ProductRepo: repo,
		rdb:         rdb,
	}
}

func (p *ProductService) InvalidateProductCache() {
	ctx := context.Background()
	cacheKey := "product:browseproduct"

	err := p.rdb.Del(ctx, cacheKey).Err()
	if err != nil {
		log.Println("Failed to delete cache:", err)
	} else {
		log.Println("Cache invalidated: products")
	}
}

func (p *ProductService) GetProduct() ([]models.Product, error) {
	return p.ProductRepo.GetAllProduct()
}

func (p *ProductService) GetProductById(id int) (models.Product, error) {
	return p.ProductRepo.GetProductById(id)
}

func (p *ProductService) AddProduct(product models.Product) (models.Product, error) {
	result, err := p.ProductRepo.AddProduct(product)
	if err != nil {
		return result, err
	}

	p.InvalidateProductCache()
	return p.ProductRepo.AddProduct(product)
}

func (p *ProductService) UpdateProductById(id int, product models.Product) (models.Product, error) {
	result, err := p.ProductRepo.AddProduct(product)
	if err != nil {
		return result, err
	}

	p.InvalidateProductCache()
	return p.ProductRepo.UpdateProductById(id, product)
}

func (p *ProductService) DeleteProductById(id int) error {
	err := p.ProductRepo.DeleteProductById(id)
	if err != nil {
		return err
	}

	p.InvalidateProductCache()
	return p.ProductRepo.DeleteProductById(id)
}

// home start
func (p *ProductService) GetProductHome(ctx context.Context) ([]models.ProductHome, error) {
	return p.ProductRepo.GetAllProductHome(ctx)
}

func (p *ProductService) ProductReview(ctx context.Context) ([]models.ReviewProduct, error) {
	return p.ProductRepo.ProductReview(ctx)
}

// home end
// browse product start
func (p *ProductService) BrowseProducts(ctx context.Context) ([]models.BrowseProduct, error) {
	return p.ProductRepo.BrowseProducts()
}

// browse product end

// detail product start
func (p *ProductService) DetailProduct(ctx context.Context, id int) ([]models.DetailProduct, error) {
	return p.ProductRepo.DetailProduct(id)
}

// detail product end
