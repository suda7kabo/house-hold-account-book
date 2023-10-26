# house-hold-account-book

curl -X 'POST' 'http://localhost:1323/expenses' -H 'accept: application/json' -H 'Content-Type: application/json' -d '{"name":"通信費"}'

 docker exec -it household-account-book-db-1 sh
  mysql -uapp-user -ptest -D household_account_book
  describe expenses;
  select * from expenses;
  