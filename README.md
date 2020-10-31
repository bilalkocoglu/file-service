# go-crud

**Get Started**

`go get -v`

`go build src/main.go`

**Depends**

Mysql \
Default host name: 'localhost' \
Default port: 3306 \
Default user: 'root' \
Default password: 'password' \
Default schema name: 'file-service' 

Minio \
Default URL: 'localhost:9000' \
Default Access Key: 'AKIAIOSFODNN7EXAMPLE'
Default Secret Key: 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
Default Main Bucket Name: 'file-service'

**Postman Collection**

https://www.getpostman.com/collections/5a1f761b87e44176ecdd

**Docker**

`docker build -t file-service .`

`docker run -d -p 9095:9095 --name file-service -e "DATABASE_HOST=localhost" -e "DATABASE_PORT=3306" -e "DATABASE_USER=root" -e "DATABASE_PASSWORD=admin" -e "DATABASE_SCHEMA=file-service" -e "MINIO_URL=localhost:9000" -e "MINIO_ACCESS_KEY=AKIAIOSFODNN7EXAMPLE" -e "MINIO_SECRET_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" -e "MAIN_BUCKET_NAME=file-service" file-service`