package marketplace

import (
	"context"
	"errors"
	"fmt"
)

// ------------------------------------------------------------
// Errors (ONLY declare errors that belong to the service layer)
// ------------------------------------------------------------

var (
	ErrInvalidQuantity = errors.New("invalid quantity")
	ErrOrderEmpty      = errors.New("order contains no items")
)

// Service contains marketplace business logic.
type Service struct {
	repo *Repository
}

// NewService creates a new marketplace service.
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// ------------------------------------------------------------
// PRODUCT LOGIC
// ------------------------------------------------------------

// ListActiveProducts returns all active products.
func (s *Service) ListActiveProducts(ctx context.Context) ([]Product, error) {
	return s.repo.ListActiveProducts(ctx)
}

// GetProduct ensures the product exists and is active.
func (s *Service) GetProduct(ctx context.Context, productID int) (*Product, error) {
	p, err := s.repo.GetProductByID(ctx, productID)
	if err != nil {
		// Use repository-level ErrProductNotFound
		return nil, ErrProductNotFound
	}
	if p.Status != "active" {
		return nil, fmt.Errorf("product %d is not active", productID)
	}
	return p, nil
}

// ------------------------------------------------------------
// ORDER LOGIC
// ------------------------------------------------------------

// CreateOrder validates items, loads product prices,
// calculates totals, and delegates DB writes to the repository.
func (s *Service) CreateOrder(ctx context.Context, userID int, items []OrderItemInput) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrOrderEmpty
	}

	var orderItems []OrderItem
	var total int

	for _, it := range items {
		if it.Quantity <= 0 {
			return nil, ErrInvalidQuantity
		}

		// Load product
		product, err := s.GetProduct(ctx, it.ProductID)
		if err != nil {
			return nil, err
		}

		// Build order item
		oi := OrderItem{
			ProductID:      product.ID,
			UnitPriceCents: product.PriceCents,
			Quantity:       it.Quantity,
		}

		orderItems = append(orderItems, oi)
		total += oi.UnitPriceCents * oi.Quantity
	}

	// Create order in DB
	order, err := s.repo.CreateOrder(ctx, userID, orderItems)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// ------------------------------------------------------------
// FILE ACCESS LOGIC (DOWNLOADS)
// ------------------------------------------------------------

// CanUserAccessFile checks if a user is allowed to download a file.
func (s *Service) CanUserAccessFile(ctx context.Context, userID, fileID int) (bool, error) {
	file, err := s.repo.GetFileByID(ctx, fileID)
	if err != nil {
		return false, err
	}

	switch file.Visibility {
	case VisibilityPublic:
		return true, nil

	case VisibilityPrivate:
		return file.OwnerUserID == userID, nil

	case VisibilityCourseOnly:
		// TODO: check if user is enrolled in the course
		return false, nil

	case VisibilityPurchasedOnly:
		purchased, err := s.repo.HasUserPurchasedFile(ctx, userID, fileID)
		if err != nil {
			return false, err
		}
		return purchased, nil
	}

	return false, nil
}
