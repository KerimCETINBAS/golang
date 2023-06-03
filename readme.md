#### [go to index](https://github.com/KerimCETINBAS/golang) - [previous lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_22) - [next lesson](https://github.com/KerimCETINBAS/golang/tree/lesson_24)

&#10;

# Lesson 23

# Basic crud app with gin framework

### start project

```bash
go run .
```

### Create user

```bash
post http://localhost:8080/users  #{ "name": "new user"}
```

### Get users

```bash
get http://localhost:8080/users
```

### Get user by id

```bash
get http://localhost:8080/users/:id
```

### update user by id

```bash
patch http://localhost:8080/users/:id  #{ "name": "new name"}
```

### delete user by id

```bash
delete http://localhost:8080/users/:id
```
