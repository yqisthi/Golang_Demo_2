- ganti gorm.opennya sesuai database pribadi
- table secara default masih kosong bisa diisi dulu pake create user:
  GET http://localhost:8080/users buat dapetin semua data
  POST http://localhost:8080/users buat create data
  POST http://localhost:8080/login buat login

required package:
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get -u github.com/jinzhu/gorm/dialects/postgres

go run main.go
