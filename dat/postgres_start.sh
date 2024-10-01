docker run -d --rm \
    --name local-pg \
    -e POSTGRES_PASSWORD=root \
    -p 5432:5432 \
    postgres
