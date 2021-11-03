App Hosted on: https://covid-tracker-rest-api.herokuapp.com/

To build locally:
1) Clone the repo from master branch 
2) Install go
3) Run go mod download from the source code folder
4) Create .env file and provide hostname, password, accesstoken to connect to 
    - LocationIQ
    - MongoDB Atlas
    - LocationIQ
    - Redis 

5) Run go build server.go
6) Run go run ./


* If you want to deploy on heroku use prod branch of repo instead
