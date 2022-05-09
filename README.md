Gin Tutorial: https://www.youtube.com/watch?v=bj77B59nkTQ

GROM Tutorial: https://www.youtube.com/watch?v=zTnkskp-xWs

GROM Documentation: https://gorm.io/docs/query.html

## Description

Simple library api that allows:

1. get all books available
2. find book by id
3. borrow a book
4. return a book

## Structure:

book.go - ORM to map with the table \
db.go - connect and query database \
route.go - business logic \
main.go - create gin server and route mapping

## How to run:

1. set up mysql database in localhost:3306
2. create a schema named gin_tutorial with the following table

```
CREATE TABLE `book` (
  `id` varchar(10) NOT NULL,
  `title` varchar(45) DEFAULT NULL,
  `author` varchar(45) DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`id`)
)
```

3. run the command "go run \*.go" in terminal
