package products

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product type representing the goods in a our datastore
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
	ModifiedOn  string  `json:"-"`
}

// FromJSON unmarsal a json object to a Product type
func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	prodlist = append(prodlist, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProductByID(id)
	if err != nil {
		return err
	}

	p.ID = id
	prodlist[pos] = p
	return nil
}

func findProductByID(id int) (*Product, int, error) {
	for pos, prod := range prodlist {
		if prod.ID == id {
			return prod, pos, nil
		}
	}
	return nil, 0, fmt.Errorf("product with %d not found", id)
}

func getNextID() int {
	last := prodlist[len(prodlist)-1]
	return last.ID + 1
}

// Products is a slice multiple goods
type Products []*Product

// ToJSON encodes a slice of Products to JSON format
func (ps *Products) ToJSON(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(ps)
}

// GetProducts returns a list of products as a slice
// from the datastore. Currently only those from memory.
// Should implement a proper database in the future.
func GetProducts() Products {
	return prodlist
}

var prodlist = []*Product{
	&Product{
		ID:          1,
		Name:        "Feldschlossen IPA",
		Description: "schweizerische bier indian pale ale brew",
		Price:       2.33,
		SKU:         "123flipa",
		CreatedOn:   time.Now().UTC().String(),
		ModifiedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Weizenbier",
		Description: "schweizerische weizenbier",
		Price:       2.33,
		SKU:         "143flwei",
		CreatedOn:   time.Now().UTC().String(),
		ModifiedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "Moehrenbreu",
		Description: "oesterreicische kellerbier",
		Price:       2.33,
		SKU:         "124oemkb",
		CreatedOn:   time.Now().UTC().String(),
		ModifiedOn:  time.Now().UTC().String(),
	},
}
