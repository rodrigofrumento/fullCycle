package database

import (
	"database/sql"

	"github.com/rodrigofrumento/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("Select id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var Product entity.Product
		if err := rows.Scan(&Product.ID, &Product.Name, &Product.Description, &Product.Price, &Product.CategoryID, &Product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, &Product)
	}
	return products, nil
}

func (pd *ProductDB) CreateProduct(Product *entity.Product) (string, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?,?,?,?,?,?)", Product.ID, Product.Name, Product.Description, Product.Price, Product.CategoryID, Product.ImageUrl)
	if err != nil {
		return "", err
	}
	return Product.ID, nil
}

func (pd ProductDB) GetProductById(id string) (*entity.Product, error) {
	var Product entity.Product
	err := pd.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&Product.ID, &Product.Name, &Product.Description, &Product.Price, &Product.CategoryID, &Product.ImageUrl)
	if err != nil {
		return nil, err
	}
	return &Product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("select id, name, description, price, category_id, image_url from products where category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
