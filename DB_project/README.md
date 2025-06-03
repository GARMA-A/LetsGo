## how to config DB mysql with your go code 



```sh
sudo apt install mysql-server
```
### this command will install mysql server on your system

```sh
$ sudo mysql
mysql>
```
### this command will open mysql shell

```sql
-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Switch to using the `snippetbox` database.
USE snippetbox;
```
### this command will create a new database called `snippetbox` and switch to it

```sql
-- Create a `snippets` table.
CREATE TABLE snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);
```
### this command will create a new table called `snippets` with the specified columns and an index on the `created` column

```sql
-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```
### this command will insert some dummy records into the `snippets` table

```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
-- Important: Make sure to swap 'pass' with a password of your own choosing.
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
FLUSH PRIVILEGES;
```
### this command will create a new user called `web` with the specified privileges on the `snippetbox` database and set a password for the user

```sh
$ mysql -D snippetbox -u web -p
Enter password:
mysql>
```
### this command will connect to the `snippetbox` database using the `web` user and prompt for the password

```sh
mysql> SELECT id, title, expires FROM snippets;
+----+------------------------+---------------------+
| id | title
| expires
|
+----+------------------------+---------------------+
| 1 | An old silent pond
| 2025-03-18 10:00:26 |
| 2 | Over the wintry forest | 2025-03-18 10:00:26 |
| 3 | First autumn morning
| 2024-03-25 10:00:26 |
+----+------------------------+---------------------+
3 rows in set (0.00 sec)
mysql> DROP TABLE snippets;
ERROR 1142 (42000): DROP command denied to user 'web'@'localhost' for table 'snippets'
```
### this command will attempt to select data from the `snippets` table and then drop the table, which will fail due to insufficient privileges for the `web` user

## now wel will use the `DATABASE` from the golang code

```sh
  go get github.com/go-sql-driver/mysql@v1
```
### this command will install the MySQL driver for Go

```go




