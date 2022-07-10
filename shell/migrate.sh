cd ~/myappTest
migrate -database="mysql://myapp:admin@tcp(127.0.0.1:3306)/myapp" -path=db/migrate/$1 up
# migrate -database="mysql://myapp:admin@tcp(10.0.2.10:3306)/myapp?charset=utf8mb4" -path=db/migrate/$1 up

