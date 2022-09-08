package entity

type Produk struct {
	Id_produk   int    `json:"id" gorm:"primaryKey"`
	Nama_produk string `json:"nama_produk"`
	Harga       int    `json:"harga"`
	Kategori    string `json:"kategori"`
	Deskripsi   string `json:"deskripsi"`
	Foto        string `json:"link_foto"`
}

type Troli struct {
	Id_produk   int    `json:"id_troli"`
	Nama_produk string `json:"name_troli"`
	Harga       int    `json:"harga_troli"`
	Kategori    string `json:"kategori_troli"`
	Deskripsi   string `json:"deskripsi_troli"`
	Foto        string `json:"foto_troli"`
}

type Payment struct {
	Total_troli   int `json:"total"`
	Bayar_troli   int `json:"bayar"`
	Kembali_troli int `json:"kembali"`
}
