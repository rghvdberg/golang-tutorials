# Setup MariaDB with docker compose

1. In a terminal navigate into the **mariadb-docker** directory and start the database.

    ```
    docker compose up -d
    ```

2. Connect to MariaDB from the MySQL/MariaDB command line client and fill the database
    * adapted from <https://github.com/docker-library/docs/blob/master/mariadb/README.md#connect-to-mariadb-from-the-mysqlmariadb-command-line-client>
      * Find the network to use.

      ```
      $ docker network ls | grep mariadb  
      a8a2c26f1718   mariadb-docker_default   bridge    local
      ```

      * Find the running maria database name.

      ```
      $ docker ps --format "{{.Names}}" | grep mariadb
      mariadb-docker-mariadb-1
      ```

      * Use a temporary mariadb container to connect to our running mariadb instance (as root) and import the data from the **create-tables.sql** file.
      We'll use the password from the **docker-compse.yml** file

      ```
      docker run --network mariadb-docker_default --rm mariadb mysql -hmariadb-docker-mariadb-1 -uroot -pmy-secret-pw < create-tables.sql
      ```
