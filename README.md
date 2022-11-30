# Clean Hexagonal Architecture SQL Transaction
A sample cli project showcasing how to use transaction with clean &amp; readable code in hexagonal architecture

## 🚀 Quick start
1. Clone the repository
```bash
git clone git@github.com:William9923/clean-architecture.git
```
2. Setup mysql infrastructure (use docker!)
```bash
sh ./scripts/mysql/run-local
```
3. Import the Postman collection in `.api` folder

4. Run the Proof of Concept Code (Transfer from user 1 to user 2 for 20000)
```bash
go run cmd/cli/main.go -from 1 -to 2 -amount 20000
```
5. To stop the infrastructure 
```bash
sh ./scripts/mysql/stop-local
```

## ❌ Prerequisites
- Golang minimum v1.17 (https://golang.org/doc/install)
- Go Modules (https://blog.golang.org/using-go-modules)

## ❤️ Support
If you feel that this repo have helped you provide more example on learning software engineering, then it is enough for me! Wanna contribute more? Please ⭐ this repo so other can see it too!
