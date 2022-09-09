package controller

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tes/config"
	"tes/entity"

	"github.com/labstack/echo/v4"
	// "gorm.io/gorm"
)

var database = &config.Database{
	Host:     os.Getenv("HOST"),
	User:     os.Getenv("USER"),
	Password: os.Getenv("PASSWORD"),
	Dbname:   os.Getenv("DBNAME"),
	Port:     os.Getenv("PORT"),
	Sslmode:  os.Getenv("SSLMODE"),
	TimeZone: os.Getenv("TIMEZONE"),
}

func GetDetailProduk(c echo.Context) (err error) {

	DB, err := config.ConnectDB(database)
	var produk entity.Produk

	detail := c.Param("nama_produk")
	result := DB.Where("nama_produk = ?", detail).Find(&produk)
	errDB := result.Error

	DP := entity.DetailProduk{
		Nama_DP:      produk.Nama_produk,
		Foto_DP:      produk.Foto,
		Harga_DP:     produk.Harga,
		Deskripsi_DP: produk.Deskripsi,
	}

	if err != nil {
		log.Println("cannot Listing produk", errDB)
		return c.JSON(http.StatusBadRequest, errDB)
	}

	return c.JSON(http.StatusAccepted, DP)
}

func GetListProduk(c echo.Context) (err error) {
	DB, err := config.ConnectDB(database)
	var produk []entity.Produk

	result := DB.Find(&produk)
	errDB := result.Error
	temp := entity.ListProduk{}
	LP := []entity.ListProduk{}

	if err != nil {
		log.Println("cannot get list of product", errDB)
		return c.JSON(http.StatusBadRequest, errDB)
	}

	for key := 0; key < len(produk); key++ {
		temp.Nama_LP = produk[key].Nama_produk
		temp.Foto_LP = produk[key].Foto
		temp.Harga_LP = produk[key].Harga
		temp.Kategori_LP = produk[key].Kategori

		LP = append(LP, temp)
	}

	return c.JSON(http.StatusAccepted, LP)
}

func GetByKategori(c echo.Context) (err error) {
	DB, err := config.ConnectDB(database)
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
	DB, err := config.ConnectDB(database)
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
	DB, err := config.ConnectDB(database)

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

func ClearTroli(c echo.Context) (err error) {

	Cart = nil

	return c.JSON(http.StatusAccepted, "Troli dikosongkan")
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
		Status:        "Transaksi Sukses",
		Total_troli:   Total,
		Bayar_troli:   pay,
		Kembali_troli: chg,
	}
	if pay < Total {
		return c.JSON(http.StatusBadRequest, "Pembayaran tidak cukup")
	}

	if Total == 0 {
		return c.JSON(http.StatusBadRequest, "Troli masih kosong")
	}

	return c.JSON(http.StatusAccepted, result)

}
