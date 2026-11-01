# qry
Qry builds your SQL query that won't let you cry.

A small SQL query builder for Go. It generates SQL strings for `SELECT`,
`INSERT`, `UPDATE`, `DELETE`, `CREATE`, `DROP`, and `ALTER`; executing the
statements and binding parameters is left to `database/sql`.

## Requirements
- Go 1.27 or newer

## Installation
```sh
go get github.com/gowok/qry
```

## Usage
### Select
```go
q := qry.Select().
	From("table").
	Where("a = 1 AND b = ?").
	SQL()

fmt.Println(q)
// output: SELECT * FROM table WHERE a = 1 AND b = ?
```

Columns, distinct rows, joins, grouping, ordering and paging:
```go
q := qry.Select("a", qry.Count("id")).
	Distinct().
	From("orders o").
	Join("customers c", "c.id = o.customer_id").
	AndWhere("o.paid = ?").
	GroupBy("a").
	OrderBy("a").
	Desc("a").
	Limit(10).
	Offset(20).
	SQL()
```

`Prefix` and `Suffix` wrap the generated statement:
```go
qry.Select("a").Prefix("WITH cte AS (SELECT 1)").SQL() // WITH cte AS (SELECT 1) SELECT a
qry.Select("1").Suffix("FOR UPDATE").SQL()             // SELECT 1 FOR UPDATE
```

### Insert
```go
q := qry.Insert("table")
q.Column("a", "b", "c")
q.Values("1", "?", "?")

fmt.Println(q.SQL())
// output: INSERT INTO table(a, b, c) VALUES(1, ?, ?)
```

Columns and values can also be appended as pairs with `Set`:
```go
q := qry.Insert("table")
q.Set("a", "1")
q.Set("b", "?")
q.Suffix("RETURNING a")

fmt.Println(q.SQL())
// output: INSERT INTO table(a, b) VALUES(1, ?) RETURNING a
```

### Update
```go
q := qry.Update("table")
q.Set("a", 1)
q.Set("b", 2)
q.Where("c = ?")

fmt.Println(q.SQL())
// output: UPDATE table SET a = 1, b = 2 WHERE c = ?
```

### Delete
```go
q := qry.Delete("table").
	Where("a = 1 AND b = ?").
	SQL()

fmt.Println(q)
// output: DELETE FROM table WHERE a = 1 AND b = ?
```

### Create
```go
q := qry.Create().
	Table("users").
	IfNotExists().
	Columns(
		"id INT PRIMARY KEY",
		"email TEXT NOT NULL",
	).
	SQL()

fmt.Println(q)
// output: CREATE TABLE IF NOT EXISTS users (id INT PRIMARY KEY, email TEXT NOT NULL)
```

`Create().Database("company")` builds the same statement for a database.

### Drop
```go
qry.Drop().Table("users").SQL()      // DROP TABLE users
qry.Drop().Database("company").SQL() // DROP DATABASE company
```

### Alter
```go
qry.Alter().Table("user").RenameTo("users").SQL()           // ALTER TABLE user RENAME TO users
qry.Alter().Table("users").AddColumn("address TEXT").SQL()  // ALTER TABLE users ADD COLUMN address TEXT
```

### Aggregates
`qry.Avg`, `qry.Count`, `qry.Max`, `qry.Min`, and `qry.Sum` return the
aggregate expression for a column:
```go
qry.Count("id")   // COUNT(id)
qry.Sum("price")  // SUM(price)
```

### Utilities
```go
qry.SQLString("alex") // 'alex'
```

Options are functional:
```go
o := qry.NewOption()
qry.WithDriver("mysql")(o)
```

## Development
```sh
go vet ./...
go test ./...
```
