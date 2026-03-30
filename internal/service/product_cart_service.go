package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
)

type ProductCartService struct {
	ProductCartRepo *repository.ProductCartRepository
}

func NewProductCartService(repo *repository.ProductCartRepository) *ProductCartService {
	return &ProductCartService{
		ProductCartRepo: repo,
	}
}

func (s *ProductCartService) AddCart(ctx context.Context, userID int, productID int, quantity int, CartItem models.CartItem) error {
	return s.ProductCartRepo.AddCart(ctx, userID, productID, CartItem.ProductName, CartItem.BasePrice, CartItem.Quantity, CartItem.VariantName, CartItem.SizeName)
	// return s.ProductCartRepo.AddCart(ctx, userID, productID, CartItem.ProductName, CartItem.BasePrice, CartItem.Quantity, CartItem.VariantName, CartItem.SizeName)
}

// CartItem.VariantName, CartItem.VariantName
// func (s *ProductCartService) AddCart(ctx context.Context, req AddToCartRequest) error {

// 	//  1. VALIDATION
// 	if req.Quantity <= 0 {
// 		return fmt.Errorf("quantity must be greater than 0")
// 	}

// 	//  2. GET PRODUCT (source of truth)
// 	product, err := s.ProductRepo.GetByID(ctx, req.ProductID)
// 	if err != nil {
// 		return fmt.Errorf("product not found")
// 	}

// 	//  3. CHECK STOCK
// 	if req.Quantity > product.Quantity {
// 		return fmt.Errorf("insufficient stock")
// 	}

// 	//  4. HITUNG HARGA (optional, kalau ada variant/size pricing)
// 	basePrice := product.Price

// 	// :
// 	//  logic variant_price + size_price kalau ada

// 	//  5. BUILD CART ITEM (INTERNAL, bukan dari client)
// 	item := models.CartItem{
// 		ProductName: product.Name,
// 		BasePrice:   basePrice,
// 		VariantName: req.VariantName,
// 		SizeName:    req.SizeName,
// 	}

// 	// ✅ 6. CALL REPOSITORY
// 	return s.ProductCartRepo.AddCart(
// 		ctx,
// 		req.UserID,
// 		req.ProductID,
// 		req.Quantity,
// 		item,
// 	)
// }
