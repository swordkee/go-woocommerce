package entity

// ProductVariationAttribute product variation attribute properties
type ProductVariationAttribute struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Option string `json:"option"`
}

// ProductVariation product variation properties
type ProductVariation struct {
	ID                int               `json:"id"`
	DateCreate        string            `json:"date_create,omitempty"`
	DateCreateGMT     string            `json:"date_create_gmt,omitempty"`
	DateModified      string            `json:"date_modified,omitempty"`
	DateModifiedGMT   string            `json:"date_modified_gmt,omitempty"`
	Description       string            `json:"description"`
	Permalink         string            `json:"permalink"`
	SKU               string            `json:"sku"`
	Price             string            `json:"price"`
	RegularPrice      string            `json:"regular_price"`
	SalePrice         string            `json:"sale_price"`
	DateOnSaleFrom    string            `json:"date_on_sale_from"`
	DateOnSaleFromGMT string            `json:"date_on_sale_from_gmt"`
	DateOnSaleTo      string            `json:"date_on_sale_to"`
	DateOnSaleToGMT   string            `json:"date_on_sale_to_gmt"`
	OnSale            bool              `json:"on_sale"`
	Status            string            `json:"status"`
	Purchasable       bool              `json:"purchasable"`
	Virtual           bool              `json:"virtual"`
	Downloadable      bool              `json:"downloadable"`
	Downloads         []ProductDownload `json:"downloads"`
	DownloadLimit     int               `json:"download_limit"`
	DownloadExpiry    int               `json:"download_expiry"`
	TaxStatus         string            `json:"tax_status"`
	TaxClass          string            `json:"tax_class"`
	ManageStock       bool              `json:"manage_stock"`
	StockQuantity     int               `json:"stock_quantity"`
	StockStatus       string            `json:"stock_status"`
	Backorders        string            `json:"backorders"`
	BackordersAllowed bool              `json:"backorders_allowed"`
	Backordered       bool              `json:"backordered"`
	Weight            string            `json:"weight"`
	// ProductDimension        *request.VariationDimensionsRequest `json:"dimensions"`
	ShippingClass   string                      `json:"shipping_class"`
	ShippingClassId int                         `json:"shipping_class_id"`
	Image           *ProductImage               `json:"image"`
	Attributes      []ProductVariationAttribute `json:"attributes"`
	MenuOrder       int                         `json:"menu_order"`
	MetaData        []Meta                      `json:"meta_data"`
}
