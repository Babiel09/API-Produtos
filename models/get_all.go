package models

import "github.com/Babiel09/configs"

func GetAll() (products []Products, err error) {
	//connection /////////////////////
	coon, err := configs.DatabaseConnection()
	if err != nil {
		return
	}
	defer coon.Close()

	//Result //////////////////////////
	result, err := coon.Query("SELECT * FROM products")
	if err != nil {
		return
	}
	for result.Next() {
		var produtos Products
		err := result.Scan(&produtos.Id, &produtos.Name, &produtos.Description, &produtos.Price, &produtos.Stock)
		if err != nil {
			continue
		}
		products = append(products, produtos)

	}

	return
}
