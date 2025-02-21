# Go Auth

<h1 align="center">
  <a href="https://gofiber.io">
    <picture>
      <source height="125" media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo-dark.svg">
      <img height="125" alt="Fiber" src="https://raw.githubusercontent.com/gofiber/docs/master/static/img/logo.svg">
    </picture>
  </a>
</h1>
<p align="center">
  <em><b>Fiber</b> is an <a href="https://github.com/expressjs/express">Express</a> inspired <b>web framework</b> built on top of <a href="https://github.com/valyala/fasthttp">Fasthttp</a>, the <b>fastest</b> HTTP engine for <a href="https://go.dev/doc/">Go</a>. Designed to <b>ease</b> things up for <b>fast</b> development with <a href="https://docs.gofiber.io/#zero-allocation"><b>zero memory allocation</b></a> and <b>performance</b> in mind.</em>
</p>

This project implements the roles and permissions for user that is assigned to a particular role, i am very suprised with how gorm query works it is very different from ORM package like TypeORM or Prisma that i have worked on. I think gorm is a very interesting ORM Mapping from golang and studying this code has been a challenging task for me, but still in the end i managed to slightly master some of these ORM query that will going to be useful in the future.

## Installation

```bash
go mod init nameof/yourproject
go get -u github.com/gofiber/fiber/v3
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/google/uuid
go get -u github.com/joho/godotenv
go get -u golang.org/x/crypto
go get -u gorm.io/driver/postgres
go get -u gorm.io/gorm
```
