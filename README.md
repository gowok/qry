# qry
Qry built your SQL query that won't let you cry

# Usage
## Select
```go
q := qry.Select().
  From("table").
  Where("a = 1 AND b = ?").
  SQL()

fmt.Println(q)
// output: SELECT * FROM table WHERE a = 1 AND b = ?
```

## Insert
```go
q := qry.Insert("table").
  Column("a", "b", "c").
  Values("1", "?", "?").
  SQL()

fmt.Println(q)
// output: INSERT INTO table(a, b, c) VALUES(1, ?, ?)
```

## Delete
```go
q := qry.Delete("table").
  Where("a = 1 AND b = ?").
  SQL()

fmt.Println(q)
// output: DELETE FROM table WHERE a = 1 AND b = ?
```

## Update
```go
q := qry.Update("table").
  Set("a", 1).
  Set("b", 2).
  Where("c = ?").
  SQL()

fmt.Println(q)
// output: UPDATE table SET a = 1, b = 2 WHERE c = ?
```
