FROM postgres:15

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

COPY 00_create_replication_user.sql /docker-entrypoint-initdb.d/

RUN chown postgres:postgres /docker-entrypoint-initdb.d/00_create_replication_user.sql
RUN chmod 755 /docker-entrypoint-initdb.d/00_create_replication_user.sql

EXPOSE 5432
