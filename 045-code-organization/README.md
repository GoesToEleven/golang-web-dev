# Great reference

[Alex Edwards Code Organization Article](http://www.alexedwards.net/blog/organising-database-access)

# Code organization

How you organize your code depends upon the project and your abilities.

Do not sacrifice simplicity & readability for brevity & cleverness.

Your goal should be to write code which is maintainable. This means that an intermediate developer should be able to sit down, read your code, understand it, and work with it.

# Different Approaches

## One package (main) with package scope variables

## Two packages with package imports

## Three+ packages with package for global variables like DB

![Scope in Go](scope.png)


