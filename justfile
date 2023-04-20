export PATH := "./graphiql/node_modules/.bin:" + env_var('PATH')

run-local:
    go build && godotenv ./gqlgen-defer-demo

build-graphiql:
    cd graphiql && esbuild --bundle --outdir=output/ --loader:.js=jsx app.js --sourcemap

gqlgen:
    go run github.com/99designs/gqlgen generate
    gofumpt -w graphql.go
