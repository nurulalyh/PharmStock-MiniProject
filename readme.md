# Pharm Stock - Mini Project MSIB 5 Alterra Academy

## Deskripsi
Pharma Stock adalah API yang dirancang untuk membantu apotek dalam mengelola stok obat dan produk farmasi mereka secara efisien. API ini menyediakan akses terpusat ke data persediaan, memungkinkan penambahan, pengeditan, dan penghapusan data produk. API ini membantu meningkatkan manajemen inventaris apotek, memastikan ketersediaan obat-obatan, dan memenuhi kebutuhan pasien dengan lebih baik.

## Spesifikasi Fitur Product
### Fitur Umum : 
- Melakukan login sebagai administrator atau apoteker menggunakan username dan password.
- Menerapkan pagination untuk memudahkan akses.
- Menerapkan searching agar user dapat mencari data berdasarkan keyword

### Administrator : 
- Dapat menambah, mengubah, dan menghapus seluruh data user.
- Dapat menambah, mengubah, menghapus, dan mengecek produk farmasi yang meliputi nama produk, deskripsi, foto produk, dan sebagainya.
- Dapat menambah, mengubah, menghapus, dan mengecek daftar distributor dari produk-produk farmasi.
- Dapat menambah, mengubah, menghapus, dan mengecek kategori product produk.
- Dapat menambah, mengubah, menghapus, dan mengecek Transaksi barang masuk dan keluar.
- Dapat mengubah status dari request menjadi ditolak, diproses, dan selesai.

### Apoteker : 
- Dapat menambah, mengubah, menghapus, dan mengecek produk farmasi yang meliputi nama produk, deskripsi, foto produk, dan sebagainya.
- Dapat menambah, mengubah, menghapus, dan mengecek Transaksi barang masuk dan keluar.
- Dapat menambah request produk yang tidak ada atau habis.

## Tech Stack
1. App Framework	    : Echo
2. ORM Library		    : GORM
3. Database		        : MySQL
4. Deployment		    : Google Cloud Platform (GCP)
5. Code Structure	    : MVC
6. Authentication		: JWT
7. Containerization	    : Docker
8. Version Control 	    : Git
9. AI Implementation	: OpenAI to generate product description 
10. Other Tools 		: Cloudinary untuk upload foto

## ERD
[ERD Pharm Stock](http://gg.gg/17afbv)
![ERD Pharm Stock-Final Ver (2)](https://github.com/nurulalyh/PharmStock-MiniProject/assets/109571028/503d4836-98c4-4239-b45c-f5a0f01c6741)


## API Documentation
[Postman](https://www.postman.com/cryosat-observer-7678182/workspace/pharm-stock/collection/23286472-4bb5439b-3976-4758-b0da-9b436a924992?action=share&creator=23286472)

<!-- ## Format File .ENV
```
SERVER=
DB_PORT=
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
SECRET=
REF_SECRET=
OPENAI_API_KEY=
``` -->
