package main

import (
	"tes/controller"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/detail_produk/:nama_produk", controller.GetDetailProduk)
	e.GET("/list_produk", controller.GetListProduk)
	e.GET("/filter_kategori/:kategori", controller.GetByKategori)
	e.GET("/filter_harga/:harga", controller.GetByRange)
	e.POST("/tambah_troli/:nama_produk", controller.AddTroli)
	e.POST("/clear_troli", controller.ClearTroli)
	e.POST("/troli/payment/:bayar", controller.Bayar)

	e.Logger.Fatal(e.Start("127.0.0.1:5002"))
}
