package entity

type Produk struct {
	Id_produk   int    `gorm:"primaryKey"`
	Nama_produk string `json:"nama_produk"`
	Harga       int    `json:"harga"`
	Kategori    string `json:"kategori"`
	Deskripsi   string 
	Foto        string `json:"link_foto"`
}

type DetailProduk struct {
	Nama_DP      string
	Foto_DP      string
	Harga_DP     int
	Deskripsi_DP string
}

type ListProduk struct {
	Nama_LP     string `json:"nama_LP"`
	Harga_LP    int    `json:"harga_LP"`
	Kategori_LP string `json:"kategori_LP"`
	Foto_LP     string `json:"foto_LP"`
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
	Status        string `json:"status"`
	Total_troli   int    `json:"total"`
	Bayar_troli   int    `json:"bayar"`
	Kembali_troli int    `json:"kembali"`
}
