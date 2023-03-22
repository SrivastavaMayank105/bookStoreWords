# bookStoreWords

move to the cmd folder and run the main.go file.
below is the sample request :
curl --location --request POST 'http://localhost:8080/booksData' \
--header 'Authorization: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VyX25hbWUiOiJtYXlhbmsifQ.fPr_amnhISyWTNJF1U5RZjeHROWPpSM3-XqjY68nWiM"' \
--header 'Content-Type: application/json' \
--data-raw '{
    "booktext":"aaa"
}'

1. docker build -t bookstore:v2 .
2. docker tag 
