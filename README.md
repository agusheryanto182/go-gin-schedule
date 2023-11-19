Tutorial : 
- create database MYSQL
- Siapkan environment variable dibawah ini:
- MYSQL_DBNAME
- MYSQL_HOST
- MYSQL_USER
- MYSQL_PASSWORD
- MYSQL_PORT
- SERVER_PORT
- SERVER_HOST
- API_KEY
- LOG_LEVEL
- Jalankan : migrate -database 'mysql://USER:PASSWORD@tcp(HOST:PORT)/DBNAME' -path database/migrations up

Setelah itu
- Jalankan : go run .
