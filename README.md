https://github.com/dirkarnez/vivom

```
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  PrepareStmt: true,
})

db.Raw("select sum(age) from users where role = ?", "admin").Scan(&age)

```
- [A Non-Intrusive Transaction Management Lib in Go — How it Works? | by Jin Feng | Medium](https://medium.com/@jfeng45/a-non-intrusive-transaction-management-lib-in-go-how-it-works-51d4b2ede8af)
- Use golang parser
  - [mockc/internal/mockc/parser.go at master · KimMachineGun/mockc](https://github.com/KimMachineGun/mockc/blob/master/internal/mockc/parser.go)
