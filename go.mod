module github.com/levinholsety/console-go

go 1.13

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20191202143827-86a70503ff7e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20191128015809-6d18c012aee9
)

require golang.org/x/sys v0.0.0-20191128015809-6d18c012aee9
