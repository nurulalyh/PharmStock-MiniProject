# Pharm Stock - Mini Project MSIB 5 Alterra Academy

## Deskripsi
Pharma Stock adalah API yang dirancang untuk membantu apotek dalam mengelola stok obat dan produk farmasi mereka secara efisien. API ini menyediakan akses terpusat ke data persediaan, memungkinkan penambahan, pengeditan, dan penghapusan data produk. API ini membantu meningkatkan manajemen inventaris apotek, memastikan ketersediaan obat-obatan, dan memenuhi kebutuhan pasien dengan lebih baik.

## Spesifikasi Fitur Product
1. Fitur Umum : 
    - Login 
        
        User dapat melakukan login sebagai administrator atau apoteker dengan memasukkan username dan password.
    - Pagination

        Data yang ditampilkan akan dibagi menjadi beberapa bagian untuk memudahkan akses.
    - Search

        User dapat mencari data sesuai kebutuhan.
    - Testing

        Unit testing dengan coverage 80% pada API memastikan bahwa 80% kode API telah diuji dengan benar.

2. Administrator : 
    - Manajemen Produk

        Dapat menambah, mengubah, menghapus, dan mengecek produk farmasi yang meliputi data, stok, kategori, dan distributor.
    - Manajemen User

        Dapat menambah, mengubah, dan menghapus seluruh data user.
    - Manajemen Request Produk

        Dapat menambah dan mengubah status dari request menjadi ditolak, diproses, dan selesai.

3. Apoteker : 
    - Manajemen Stok
        
        Dapat mencatat stok produk farmasi yang keluar-masuk.
    - Melakukan Request Produk
        
        Dapat menambah request produk yang tidak ada atau habis.

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
[ERD Pharm Stock](http://gg.gg/176p6f)

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