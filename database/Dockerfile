FROM mysql:8
ADD mysqld.cnf /etc/mysql/conf.d/
ADD docker-entrypoint-initdb.d /docker-entrypoint-initdb.d/
ADD seed /seed/
