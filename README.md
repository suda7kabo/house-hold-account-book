# house-hold-account-book

curl -X 'POST' 'http://localhost:1323/expenses' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"name":"通信費"}'
curl -X 'POST' 'http://localhost:1323/expenses' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"name":"生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費生活費"}'

 docker exec -it household-account-book-db-1 sh
  mysql -uapp-user -ptest -D household_account_book
  describe expenses;
  SET NAMES utf8;  select * from expenses;

mkdir -p test/mock/repository
touch test/mock/repository/mockgen.go
  