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
