package types

type UserStore interface {
	GetUserByEmail(user UserLogin) (*User, error)
	GetUserById(user User) (*User, error)
	CreateUser(user User, token string, hashedPassword string) error
}
type CommandeStore interface {
	GetAllCommandes() (*[]Commande, error)
	GetCommandeById(id string) (*Commande, error)
	CreateCommande(idCommande, idCustomer string) error
	// UpdateCommande(commande Commande) error
	// DeleteCommande(commande Commande) error
	// GetCommandeByUser(idUser string) (*[]Commande, error)
	InsertProductINCommande(product Product, idCommande string) (*CommandeProduct, error)
}

type CommandeProduct struct {
	IdCommande string `json:"idCommande"`
	IdProduct  string `json:"idProduct"`
}

type Commande struct {
	IdCommande string `json:"idCommande"`
	IdCustomer string `json:"idCustomer"`
}
type CommandeCreate struct {
	IdCustomer string `json:"idCustomer"`
}

type User struct {
	Id           string `json:"idProfile"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Address      string `json:"adress"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	LastLogin    string `json:"lastLogin"`
	CreatedAt    string `json:"createdAt"`
	Refreshtoken string `json:"refreshToken"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Address   string `json:"adress"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type ProductStore interface {
	GetProductById(id string) (*Product, error)
	GetAllProducts() (*[]Product, error)
	CreateProduct(product ProductCreate, idProduct string) error
	// UpdateProduct(product Product) error
	// DeleteProduct(product Product) error
	GetProductByCategorie(idCategorie string) (*[]Product, error)
}

type Product struct {
	IdProduct   string `json:"idProduct"`
	NameProduct string `json:"nameProduct"`
	Price       string `json:"price"`
	Description string `json:"description"`
	IdCategorie string `json:"idCategorie"`
}

type ProductCreate struct {
	NameProduct string `json:"nameProduct"`
	Price       string `json:"price"`
	Description string `json:"description"`
	IdCategorie string `json:"idCategorie"`
}

type Categorie struct {
	IdCategorie   string `json:"idCategorie"`
	NameCategorie string `json:"nameCategorie"`
}
