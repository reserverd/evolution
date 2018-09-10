
main: main.go src/web/web.go src/db/db.go
	go build main.go
clean:
	rm -f main
