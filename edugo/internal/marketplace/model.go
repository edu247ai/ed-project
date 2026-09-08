package marketplace

import "time"

// ------------------------------------------------------------
// PRODUCT MODELS
// ------------------------------------------------------------

type ProductType string

const (
	ProductTypeEbook       ProductType = "ebook"
	ProductTypeAudiobook   ProductType = "audiobook"
	ProductTypeCoffeeAudio ProductType = "coffee_audio"
)

type Product struct {
	ID          int         `json:"id"`
	Type        ProductType `json:"type"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	PriceCents  int         `json:"price_cents"`
	Currency    string      `json:"currency"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
}

// ------------------------------------------------------------
// FILE MODELS
// ------------------------------------------------------------

type FileVisibility string

const (
	VisibilityPrivate       FileVisibility = "private"
	VisibilityCourseOnly    FileVisibility = "course_only"
	VisibilityPurchasedOnly FileVisibility = "purchased_only"
	VisibilityPublic        FileVisibility = "public"
)

type File struct {
	ID           int            `json:"id"`
	OwnerUserID  int            `json:"owner_user_id"`
	OriginalName string         `json:"original_name"`
	StoragePath  string         `json:"storage_path"`
	MimeType     string         `json:"mime_type"`
	SizeBytes    int64          `json:"size_bytes"`
	Visibility   FileVisibility `json:"visibility"`
	CreatedAt    time.Time      `json:"created_at"`
}

// ------------------------------------------------------------
// ORDER MODELS
// ------------------------------------------------------------

type Order struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	TotalPriceCents int       `json:"total_price_cents"`
	Currency        string    `json:"currency"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

type OrderItem struct {
	ID             int `json:"id"`
	OrderID        int `json:"order_id"`
	ProductID      int `json:"product_id"`
	UnitPriceCents int `json:"unit_price_cents"`
	Quantity       int `json:"quantity"`
}

// ------------------------------------------------------------
// SERVICE INPUT MODELS
// ------------------------------------------------------------

type OrderItemInput struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
