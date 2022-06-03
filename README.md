# WAPE Microservice

Wape is a microservice that has 2 different apps that allow you to upload your files to the bucket and insert them into DB them after parsing

## ✨ Installation 

First, need to clone the repository


`HTTPS` clone

```bash
git clone https://github.com/metecan/wape.git
```
`SSH` clone

```bash
git clone git@github.com:metecan/wape.git
```

## ✨ Running the apps

Starting docker compose to run the apps

```bash
make up
```
OR 

```bash
docker compose up --build
```

Everything will be up and running.

If you want to all the apps down, just run

```bash
make down
```
OR 

```bash
docker compose down
```

Every container will be downed


## ✨ Usage 

If all containers are up and running, you can use the following steps to upload files to the bucket

1. Open Postman or any other REST client
2. Enter the following URL: http://localhost:8000/ or http://127.0.0.1:8000/

- (If it is working, you can see the following message: "Wape API is alive" and the status code is 200)

3. Send a POST request to the following URL: http://localhost:8000/page or http://127.0.0.1:8000/page
4. Select files using multipart/form-data with the `file` prefix

- (You can upload multiple files at once)

- Example for uploading a file:

![example1](https://res.cloudinary.com/allstar/image/upload/c_scale,w_437/v1654234467/Screenshot_at_Jun_03_08-34-00_zxgyed.png)

5. Send the request
6. Check the response

If it is working, you can see the message like the following message:

![example2](https://res.cloudinary.com/allstar/image/upload/c_scale,w_408/v1654234769/Screenshot_at_Jun_03_08-39-07_yj73ro.png)


File upload is successful.

Then you can use the following steps to insert the files into the DB

Send a GET request with REST client to the following URL: http://localhost:8001/ or http://127.0.0.1:8001/page
Then parser will be triggered to parse the files and insert them into the DB

or you can use terminal to run the following command:

```bash
curl http://localhost:8001/
```

If everything is working, you can check the database to see the parsed file data

![example3](https://res.cloudinary.com/allstar/image/upload/v1654235178/Screenshot_at_Jun_03_08-45-45_zg9bcu.png)

## ✨ Used technologies
- Beanstalkd
- Postgres
- Docker
- Gofiber
- Godotenv
- Supabase Storage for S3 Bucket


### ✨ Folder structure

```bash
WAPE
│
├── docker-compose.yml
├── Makefile
├── .gitignore
├── README.md
├── api/
│   ├── handler/
│   ├── helper/
│   ├── routes/
│   ├── test/
│   ├── utils/
│   ├── main.go
│   ├── Dockerfile
│   ├── .gitignore
│   ├── go.mod
│   ├── go.sum
│   └── .env
├── parser/
│   ├── db/
│   ├── handler/
│   ├── helper/
│   ├── main.go
│   ├── Dockerfile
│   ├── .gitignore
│   ├── go.mod
│   ├── go.sum
│   └── .env

```
✨ 
Last Update: 2022 June 03