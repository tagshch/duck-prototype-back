curl -X POST "http://localhost:8080/api/user/add"
curl -X POST "http://localhost:8080/api/user/remove"
curl -X POST "http://localhost:8080/api/user/list"
curl -X POST "http://localhost:8080/api/file/add"
curl -X POST "http://localhost:8080/api/file/remove"
curl -X POST "http://localhost:8080/api/file/list"


# examples

curl -X GET "http://localhost:8081/api/file/list"

"https://storage.yandexcloud.net/duckdbshell/2023_4.json",
"https://storage.yandexcloud.net/duckdbshell/2024_1.json",


curl -X POST "http://localhost:8081/api/file/add" -H 'Content-Type: application/json' -d'
{
  "table" : "2024_2",
  "url": "https://storage.yandexcloud.net/duckdbshell/2024_1.json"
}
' -k


curl -X POST "http://localhost:8081/api/file/remove" -H 'Content-Type: application/json' -d'
{
  "id" : 3
}
' -k