module pingumail

go 1.24.0

replace server => ./server

replace client => ./client

require (
	client v0.0.0-00010101000000-000000000000
	github.com/urfave/cli/v2 v2.27.7
	golang.org/x/term v0.37.0
	server v0.0.0-00010101000000-000000000000
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/crypto v0.45.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
)
