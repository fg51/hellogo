# hellogo
golang sample


## project

### correct
```sh
root/
  sub/
    x.go (pkg: sub)
    y.go (pkg: sub)
  a.go (pkg: main)
  b.go (pkg: main)
```

### incorrect
```sh
root/
  sub/
    x.go (pkg: sub)
    y.go (pkg: bar !!!)
  a.go (pkg: main)
  b.go (pkg: foo !!!)
```

### another correct pattern
```sh
root/
  cmd/
    a.go (pkg: main)
    b.go (pkg: main)
  sub/
    x.go (pkg: sub)
    y.go (pkg: sub)
  l.go (pkg: lib)
  m.go (pkg: lib)
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



