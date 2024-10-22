CREATE ROLE auth_user WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOREPLICATION
    PASSWORD 'auth_user'
    CONNECTION LIMIT -1;

GRANT SELECT ON racket TO auth_user;
GRANT INSERT ON "user" TO auth_user;