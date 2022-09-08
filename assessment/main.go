package main

import (
	"tes/controller"

	"github.com/labstack/echo/v4"
	// "honnef.co/go/tools/config"
)

func main() {

	// config.ConnectDB()
	e := echo.New()

	e.GET("/produk", controller.GetAllProduk)
	e.GET("/produk/:kategori", controller.GetByKategori)
	e.GET("/produk/:harga", controller.GetByRange)
	e.POST("/troli/:nama_produk",controller.AddTroli)
	e.POST("/troli/payment/:bayar",controller.Bayar)

	e.Logger.Fatal(e.Start("127.0.0.1:5002"))
}
