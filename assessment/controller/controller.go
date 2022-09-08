package controller

import (
	// "tes/config"

	"log"
	"net/http"
	"strconv"
	"strings"
	"tes/config"
	"tes/entity"

	"github.com/labstack/echo/v4"
	// "gorm.io/gorm"
)

func GetAllProduk(c echo.Context) (err error) {
	config.ConnectDB()
	DB := config.DB
	var produk []entity.Produk

	result := DB.Find(&produk)
	errDB := result.Error
	if err != nil {
		log.Println("cannot Listing produk", errDB)
		return c.JSON(http.StatusBadRequest, errDB)
	}

	return c.JSON(http.StatusAccepted, produk)
}

func GetByKategori(c echo.Context) (err error) {
	config.ConnectDB()
	DB := config.DB
	var produk []entity.Produk

	kat := c.Param("kategori")
	result := DB.Where("kategori = ?", kat).Find(&produk)
	errDB := result.Error
	if err != nil {
		log.Println("cannot get produk by kategori", errDB)
		return c.JSON(http.StatusBadRequest, errDB)
	}

	return c.JSON(http.StatusAccepted, produk)
}

func GetByRange(c echo.Context) (err error) {
	config.ConnectDB()
	DB := config.DB
	var produk []entity.Produk

	str := c.Param("harga")
	split := strings.Split(str, ",")
	result := DB.Where("harga >= ? AND harga <= ?", split[0], split[1]).Find(&produk)
	errDB := result.Error

	if err != nil {
		log.Println("error cannot get produk", errDB)
		return c.JSON(http.StatusBadRequest, errDB)
	}

	return c.JSON(http.StatusAccepted, produk)
}

var Cart = []entity.Troli{}

func AddTroli(c echo.Context) (err error) {
	config.ConnectDB()
	DB := config.DB

	var produk entity.Produk

	nama := c.Param("nama_produk")
	x := DB.Where("nama_produk = ?", nama).Find(&produk)
	Cart = append(Cart, entity.Troli(produk))
	errDB := x.Error

	if err != nil {
		log.Println("error cannot add to cart", errDB)
		return c.JSON(http.StatusBadRequest, errDB)
	}

	return c.JSON(http.StatusAccepted, Cart)
}

func Bayar(c echo.Context) (err error) {
	res := c.Param("bayar")
	pay, _ := strconv.Atoi(res)
	var Total int = 0
	for key, _ := range Cart {

		Total = Total + Cart[key].Harga
	}

	chg := pay - Total
	
	result := entity.Payment{

		Total_troli:   Total,
		Bayar_troli:   pay,
		Kembali_troli: chg,
	}
	if pay < Total{
		return c.JSON(http.StatusBadRequest, "Uang tidak cukup")
	}

	return c.JSON(http.StatusAccepted, result)

}
