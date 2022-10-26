# gen-bcrypt-passwd

## About

A simple golang tool to generate hashed passwords using the [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) library

## Usage

```shell
> go run .
Enter password to hash:
admin

This chart will help you calculate the cost and interations of
generating the hash.

|------|-------------------|
| Cost | Iterations        |
|------|-------------------|
|  2   |    4 iterations   |
|  4   |    16 iterations  |
|  5   |    32 iterations  |
|  6   |    64 iterations  |
|  7   |    128 iterations |
|  8   |    256 iterations |
|  9   |    512 iterations |
| 10   |  1,024 iterations |
| 11   |  2,048 iterations |
| 12   |  4,096 iterations |
| 13   |  8,192 iterations |
| 14   | 16,384 iterations |
| 15   | 32,768 iterations |
| 16   | 65,536 iterations |
|------|-------------------|
Enter cost:
4
Password Hash: $2a$04$t.zOlAavZLUnTZPWlhvrV.3NKSuiRKLAYRMDXuG3zz1y3yRCkRoTO
Password Match: true
```
