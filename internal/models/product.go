package models

type Product struct {
	Id          int     `json:"id" db:"id"`
	Name        *string `json:"product_name" db:"product_name"`
	Description *string `json:"product_desc" db:"product_desc"`
	Price       *int    `json:"price" db:"price"`
	Quantity    *int    `json:"quantity" db:"quantity"`
	Discount    *int    `json:"discount" db:"discount"`
}

type Discount struct {
	Id          int    `json:"discount_id" db:"discount_id"`
	Rate        int    `json:"discount_rate" db:"discount_rate"`
	Description string `json:"description" db:"description"`
	FlashSale   string `json:"is_flash_sale" db:"is_flash_sale"`
}

type ProductImages struct {
	Images    int    `json:"product_images_id" db:"product_images_id"`
	ProductId int    `json:"product_id" db:"product_id"`
	Path      string `json:"path" db:"path"`
}

type ProductSize struct {
	Size      int    `json:"product_size_id" db:"product_size_id"`
	ProductId int    `json:"product_id" db:"product_id"`
	SizeName  string `json:"size_name" db:"size_name"`
	SizePrice int    `json:"size_price" db:"size_price"`
}

type Variant struct {
	VariantId   int    `json:"product_variant_id" db:"product_variant_id"`
	ProductId   int    `json:"product_id" db:"product_id"`
	VariantName string `json:"variant_name" db:"variant_name"`
	Price       string `json:"add_price" db:"add_price"`
}

type ProductHome struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"product_name" db:"product_name"`
	Description string  `json:"product_desc" db:"product_desc"`
	Price       int     `json:"price" db:"price"`
	Ratings     *int    `json:"ratings" db:"ratings"`
	Path        *string `json:"path" db:"path"`
}

type ReviewProduct struct {
	UserName string  `json:"full_name" db:"full_name"`
	Pictures *string `json:"pictures" db:"pictures"`
	Message  *string `json:"messages" db:"messages"`
	Rating   *int    `json:"ratings" db:"ratings"`
}

// browse product
type BrowseProduct struct {
	Id          int     `json:"id" db:"id"`
	Name        *string `json:"product_name" db:"product_name"`
	Description *string `json:"product_desc" db:"product_desc"`
	Price       *int    `json:"price" db:"price"`
	Quantity    *int    `json:"quantity" db:"quantity"`
	Discount    *int    `json:"discount" db:"discount"`
	IsFlashSale *bool   `json:"is_flash_sale" db:"is_flash_sale"`
}

// detailproduct
// 1. images
// 2. is flash sale
// 3. title = done
// 4. price and oldprice = done
// 5. review star = done
// 6. desc = done
// 7. size = done
// 8. variant = done
type DetailProduct struct {
	Id           int     `json:"id" db:"id"`
	Name         *string `json:"product_name" db:"product_name"`
	Description  *string `json:"product_desc" db:"product_desc"`
	Price        *int    `json:"price" db:"price"`
	Quantity     *int    `json:"quantity" db:"quantity"`
	Discount     *int    `json:"discount" db:"discount"`
	DiscountRate *int    `json:"discount_rate" db:"discount_rate"`
	IsFlashSale  *bool   `json:"is_flash_sale" db:"is_flash_sale"`
	Images       string  `json:"path" db:"path"`
	Sizes        string  `json:"sizes" db:"sizes"`
	Variants     string  `json:"variants" db:"variants"`
	Rating       *int    `json:"ratings" db:"ratings"`
}
