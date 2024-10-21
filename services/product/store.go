package product

import (
	"database/sql"

	"github.com/wael-boudissaa/marquinoBackend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProductById(id string) (*types.Product, error) {
	query := `SELECT * FROM product where idProduct= ?`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	productResult := new(types.Product)
	for rows.Next() {
		productResult, err = s.scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	return productResult, nil
}

func (h *Store) GetAllProducts() (*[]types.Product, error) {
	query := `SELECT * FROM product`
	rows, err := h.db.Query(query)
	if err != nil {
		return nil, err
	}
	products := new([]types.Product)
	product := new(types.Product)
	for rows.Next() {
		err := rows.Scan(&product.IdProduct, &product.NameProduct, &product.Price, &product.Description, &product.IdCategorie)

		if err != nil {
			return nil, err
		}
		*products = append(*products, *product)
	}
	return products, nil
}

func (h *Store) CreateProduct(product types.ProductCreate, idProduct string) error {
	query := `INSERT INTO product (idProduct, nameProduct, price, description, idCategorie,boosted,stock,dateExpiration,createdAt) VALUES (?,  ?, ?, ?,?,?,?,?,now())`
	rows, err := h.db.Query(query, idProduct, product.NameProduct, product.Price, product.Description, product.IdCategorie, product.Boosted, product.Stock, product.DateExpiration)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

// func (h *Store) UpdateProduct(product types.Product) error {
//
// }

func (h *Store) DeleteProduct(product types.Product) error {
	query := `DELETE FROM product WHERE idProduct = ?`
	rows, err := h.db.Query(query, product.IdProduct)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (h *Store) GetProductByCategorie(idCategorie string) (*[]types.Product, error) {
	query := `SELECT * FROM product where idCategorie= ?`
	rows, err := h.db.Query(query, idCategorie)
	if err != nil {
		return nil, err
	}
	products := new([]types.Product)
	product := new(types.Product)
	for rows.Next() {
		err := rows.Scan(&product.IdProduct, &product.NameProduct, &product.Price, &product.Description, &product.IdCategorie)

		if err != nil {
			return nil, err
		}
		*products = append(*products, *product)
	}
	return products, nil
}

func (s *Store) scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	for rows.Next() {
		err := rows.Scan(&product.IdProduct, &product.NameProduct, &product.Price, &product.Description, &product.IdCategorie)
		if err != nil {
			return nil, err
		}
	}
	return product, nil

}
