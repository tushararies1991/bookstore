
Create Table with unique email id and id as primary key
```
CREATE TABLE users_db.users (
          id BIGINT(20) NOT NULL AUTO_INCREMENT, 
          first_name VARCHAR(45) NULL, 
          last_name VARCHAR(45) NULL, 
          email VARCHAR(45) NOT NULL, 
          phone VARCHAR(45) NULL, 
          created_at VARCHAR(45) NULL,
          PRIMARY KEY (id),
          UNIQUE INDEX email_UNIQUE (email ASC)
      );

Alter Table, updating column type
```
ALTER TABLE users_db.users
MODIFY created_at DATETIME NOT NULL AFTER email;

Alter Table, adding column type
```
ALTER TABLE users_db.users
ADD status VARCHAR(45) NOT NULL AFTER phone,
ADD password VARCHAR(32) NOT NULL;