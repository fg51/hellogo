# hellogo
golang sample


## project

### correct
```sh
root/
  a.go (pkg: main)
  b.go (pkg: main)
  sub/
    x.go (pkg: sub)
    y.go (pkg: sub)
```

### incorrect
```sh
root/
  a.go (pkg: main)
  b.go (pkg: foo !!!)
  sub/
    x.go (pkg: sub)
    y.go (pkg: bar !!!)
```

## boolean

- && and
- || or
- !  not

```go
true && true
true && false
true || true
true || false
!true
```



