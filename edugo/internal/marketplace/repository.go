package marketplace

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrFileNotFound    = errors.New("file not found")
)

// Repository handles all DB operations for marketplace.
type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

//
// ------------------------------------------------------------
// PRODUCTS
// ------------------------------------------------------------
//

// ListActiveProducts returns all active marketplace products.
func (r *Repository) ListActiveProducts(ctx context.Context) ([]Product, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, type, title, description, price_cents, currency, status, created_at
        FROM products
        WHERE status = 'active'
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ID, &p.Type, &p.Title, &p.Description,
			&p.PriceCents, &p.Currency, &p.Status, &p.CreatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, rows.Err()
}

// GetProductByID loads a single product.
func (r *Repository) GetProductByID(ctx context.Context, productID int) (*Product, error) {
	var p Product

	err := r.db.QueryRow(ctx, `
        SELECT id, type, title, description, price_cents, currency, status, created_at
        FROM products
        WHERE id = $1
    `, productID).Scan(
		&p.ID, &p.Type, &p.Title, &p.Description,
		&p.PriceCents, &p.Currency, &p.Status, &p.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrProductNotFound
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

//
// ------------------------------------------------------------
// FILES
// ------------------------------------------------------------
//

// GetFileByID loads file metadata.
func (r *Repository) GetFileByID(ctx context.Context, fileID int) (*File, error) {
	var f File

	err := r.db.QueryRow(ctx, `
        SELECT id, owner_user_id, original_name, storage_path,
               mime_type, size_bytes, visibility, created_at
        FROM files
        WHERE id = $1
    `, fileID).Scan(
		&f.ID, &f.OwnerUserID, &f.OriginalName, &f.StoragePath,
		&f.MimeType, &f.SizeBytes, &f.Visibility, &f.CreatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrFileNotFound
	}
	if err != nil {
		return nil, err
	}

	return &f, nil
}

// HasUserPurchasedFile checks if a user bought a product containing this file.
func (r *Repository) HasUserPurchasedFile(ctx context.Context, userID, fileID int) (bool, error) {
	var exists bool

	err := r.db.QueryRow(ctx, `
        SELECT EXISTS (
            SELECT 1
            FROM orders o
            JOIN order_items oi ON oi.order_id = o.id
            JOIN product_files pf ON pf.product_id = oi.product_id
            WHERE o.user_id = $1
              AND pf.file_id = $2
              AND o.status = 'paid'
        )
    `, userID, fileID).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

//
// ------------------------------------------------------------
// ORDERS
// ------------------------------------------------------------
//

// CreateOrder inserts an order + items inside a transaction.
func (r *Repository) CreateOrder(ctx context.Context, userID int, items []OrderItem) (*Order, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Calculate total
	var total int
	for _, it := range items {
		total += it.UnitPriceCents * it.Quantity
	}

	// Insert order
	var order Order
	err = tx.QueryRow(ctx, `
        INSERT INTO orders (user_id, total_price_cents, currency, status)
        VALUES ($1, $2, 'ZAR', 'pending')
        RETURNING id, user_id, total_price_cents, currency, status, created_at
    `, userID, total).Scan(
		&order.ID, &order.UserID, &order.TotalPriceCents,
		&order.Currency, &order.Status, &order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Insert items
	for _, it := range items {
		_, err := tx.Exec(ctx, `
            INSERT INTO order_items (order_id, product_id, unit_price_cents, quantity)
            VALUES ($1, $2, $3, $4)
        `, order.ID, it.ProductID, it.UnitPriceCents, it.Quantity)
		if err != nil {
			return nil, err
		}
	}

	// Commit
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &order, nil
}
