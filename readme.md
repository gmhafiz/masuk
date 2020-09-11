# Protobuf

    build/proto-gen.sh

# Migration

    migrate create -ext sql -dir database/migrations -format unix create_users_table

    migrate -database "postgres://127.0.0.1/{{database}}?sslmode=disable&user={{user}}&password={{password}}" -path database/migrations up
    