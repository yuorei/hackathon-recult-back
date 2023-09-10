CREATE TABLE USERS (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    hashed_password VARCHAR(100) NOT NULL,
    gender ENUM('MAN','WOMAN','OTHER'),
    affiliation ENUM('STUDENT','OTHER'),
    PRIMARY KEY (id)
);