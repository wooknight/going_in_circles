FROM postgres:alpine

COPY title.basics.tsv.gz /docker-entrypoint-initdb.d/
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=title.basics