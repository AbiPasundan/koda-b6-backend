package models

type Product struct {
	Id          int     `db:"id"`
	Name        string  `db:"product_name"`
	Description *string `db:"product_desc"`
	Price       *int    `db:"price"`
	Quantity    *int    `db:"quantity"`
	Discount    *int    `db:"discount"`
}

type Discount struct {
	Id          int    `db:"discount_id"`
	Rate        int    `db:"discount_rate"`
	Description string `db:"description"`
	FlashSale   string `db:"is_flash_sale"`
}

type Category struct {
	Id   int    `db:"category_id"`
	Name string `db:"category_name"`
}

type ProductImages struct {
	Images    int    `db:"product_images_id"`
	ProductId int    `db:"product_id"`
	Path      string `db:"path"`
}

type ProductSize struct {
	Size      int    `db:"product_size_id"`
	ProductId int    `db:"product_id"`
	SizeName  string `db:"size_name"`
	SizePrice int    `db:"size_price"`
}

type Variant struct {
	VariantId   int    `db:"product_variant_id"`
	ProductId   int    `db:"product_id"`
	VariantName string `db:"variant_name"`
	Price       string `db:"add_price"`
}
