# apigoevermos
apigoevermos

GO & MongoDB

1. Berdasarkan analisis yang saya dapat dari kasus tersebut adalah, inventaris tidak mempertimbangkan stok barang yang tersedia dengan jumlah barang yang diinput di toko.
2. Solusinya adalah pada aplikasi terebut ketika checkout menghit API stok barang, ketika stok barang sudah kosong maka ada error handling untuk memberitahu konsumen bahwa barang yang ingin dibeli sudah tidak tersedia.

# API

ENDPOINT

1. http://localhost:2021/api/v1/apievermos/evermos/merchant/insert
{
    "merchantname" : "Toko Maju",
    "address" : "Panglima Barat",
    "phone" : "089883938422",
    "owner" : "Andri",
    "isactive" : 1
}
2. http://localhost:2021/api/v1/apievermos/evermos/merchant/list
{
    "id" : "6080368874b4257e562946e8"
}
3. http://localhost:2021/api/v1/apievermos/evermos/product/insert
{
    "merchantid" : "6080368874b4257e562946e8",
    "merchantname" : "Toko Bahagia",
    "productname" : "H & M Sweater",
    "price" : "1240000",
    "category" : "Fashion Wanita",
    "stock": 4,
    "isactive" : 1
}
4. http://localhost:2021/api/v1/apievermos/evermos/product/list
{
    "id" : "60804b2e62db59dfa3da9d5b"
}

insert product in param stock = 0 to check validation 

How to run in MAC OS or Linux  : cd /src/apievermos/ 
                                : go build main.go
                                : ./main 
        
How to run  in Windows  : cd /src/apievermos/ 
                        : go build main.go
                        : main.exe

# MongoDB

