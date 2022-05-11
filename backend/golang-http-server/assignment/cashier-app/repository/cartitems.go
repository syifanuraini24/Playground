package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type CartItemRepository struct {
	db db.DB
}

const cartItemDbName = "cart_items"

var cartItemColumns = []string{"category", "product_name", "price", "quantity"}

func NewCartItemRepository(db db.DB) CartItemRepository {
	return CartItemRepository{db}
}

func (u *CartItemRepository) LoadOrCreate() ([]CartItem, error) {
	records, err := u.db.Load(cartItemDbName)
	if err != nil {
		records = [][]string{cartItemColumns}
		if err := u.db.Save(cartItemDbName, records); err != nil {
			return nil, err
		}
	}

	result := make([]CartItem, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(records[i][3])
		if err != nil {
			return nil, err
		}

		cartItem := CartItem{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
			Quantity:    qty,
		}
		result = append(result, cartItem)
	}

	return result, nil
}

func (u *CartItemRepository) Save(cartItems []CartItem) error {
	records := [][]string{cartItemColumns}
	for i := 0; i < len(cartItems); i++ {
		records = append(records, []string{
			cartItems[i].Category,
			cartItems[i].ProductName,
			strconv.Itoa(cartItems[i].Price),
			strconv.Itoa(cartItems[i].Quantity),
		})
	}
	return u.db.Save(cartItemDbName, records)
}

func (u *CartItemRepository) SelectAll() ([]CartItem, error) {
	return u.LoadOrCreate()
}

func (u *CartItemRepository) Add(product Product) error {
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	// cek kalau ada product yang sama di cart items
	i, exist := searchCartItem(cartItems, product)

	if exist {
		// tambah quantity
		cartItems[i].Quantity = cartItems[i].Quantity + 1
	} else {
		// add cart item baru
		item := CartItem{
			Category:    product.Category,
			ProductName: product.ProductName,
			Price:       product.Price,
			Quantity:    1,
		}
		cartItems = append(cartItems, item)
	}

	// save ke db
	return u.Save(cartItems)
}

func (u *CartItemRepository) ResetCartItems() error {
	u.db.Delete(cartItemDbName)
	records := [][]string{cartItemColumns}
	return u.db.Save(cartItemDbName, records)
}

func (u *CartItemRepository) TotalPrice() (int, error) {
	// load cart items
	cartItems, err := u.LoadOrCreate()
	if err != nil {
		return 0, err
	}

	// loop cart items, total += item.price * item.quantity
	var total int
	for _, item := range cartItems {
		total = total + (item.Price * item.Quantity)
	}

	return total, nil
}

func searchCartItem(cartItems []CartItem, product Product) (int, bool) {
	for i, item := range cartItems {
		if item.Category == product.Category &&
			item.ProductName == product.ProductName &&
			item.Price == product.Price {
			return i, true
		}
	}
	return -1, false
}
