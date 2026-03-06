package models

// id, product_name, product_desc, price, quantity, product_category, discount

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"product_name"`
	Description string `json:"product_desc"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Category    int    `json:"product_category"`
	Discount    int    `json:"discount"`
}

type Discount struct {
	// -- discount_id, discount_rate, description, is_flash_sale
	Id          int    `json:"discount_id"`
	Rate        int    `json:"discount_rate"`
	Description string `json:"description"`
	FlashSale   string `json:"is_flash_sale"`
}

type Category struct {
	// -- category_id, category_name
	Id   int    `json:"category_id"`
	Name string `json:"category_name"`
}

type ProductImages struct {
	// -- product_images_id, product_id, path
	Images    int    `json:"product_images_id"`
	ProductId int    `json:"product_id"`
	Path      string `json:"path"`
}

type ProductSize struct {
	// -- product_size_id, product_id, size_name, size_price
	Size      int    `json:"product_size_id"`
	ProductId int    `json:"product_id"`
	SizeName  string `json:"size_name"`
	SizePrice int    `json:"size_price"`
}

type Variant struct {
	// -- product_variant_id, product_id, variant_name, add_price
	VariantId   int    `json:"product_variant_id"`
	ProductId   int    `json:"product_id"`
	VariantName string `json:"variant_name"`
	Price       string `json:"add_price"`
}
