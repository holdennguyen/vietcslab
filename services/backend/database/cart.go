package database

import "errors"

var (

	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't decode the product")
	ErrUserIdIsNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("can't add this product to the cart")
	ErrCantRemoveCartItem = errors.New("can't remove this item from the cart")
	ErrCantGetItem = errors.New("can't get the item to the cart")
	ErrCantBuyCartItem = errors.New("can't purchase")

)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}