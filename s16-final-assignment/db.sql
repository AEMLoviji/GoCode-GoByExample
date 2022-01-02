-- Execute below scripts to 
-- 1. create schema, tables in your MySql instance
-- 2. and fill the tables with data 

create schema `go_db1`;

CREATE TABLE house (
	id SERIAL PRIMARY KEY,
	price FLOAT,
	min_down FLOAT,
	prop_tax FLOAT,
	m_cost FLOAT
); 

INSERT INTO `go_db1`.`house` (`id`, `price`, `min_down`, `prop_tax`, `m_cost`)
VALUES ('1', '2000000', '400000', '9200', '1800');
INSERT INTO `go_db1`.`house` (`id`, `price`, `min_down`, `prop_tax`, `m_cost`)
VALUES ('2', '1750000', '350000', '8050', '1800');
INSERT INTO `go_db1`.`house` (`id`, `price`, `min_down`, `prop_tax`, `m_cost`)
VALUES ('3', '1500000', '300000', '6900', '1800');
INSERT INTO `go_db1`.`house` (`id`, `price`, `min_down`, `prop_tax`, `m_cost`)
VALUES ('4', '1250000', '250000', '5750', '1800'); 

CREATE TABLE customer (
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	credit_score INT,
	salary FLOAT,
	downpayment FLOAT,
	house_id INT
); 

INSERT INTO `go_db1`.`customer` (`id`, `first_name`, `last_name`, `credit_score`,
`salary`, `downpayment`, `house_id`) VALUES ('1', 'Nick', 'White', '670', '253000',
'320000', '3');
INSERT INTO `go_db1`.`customer` (`id`, `first_name`, `last_name`, `credit_score`,
`salary`, `downpayment`, `house_id`) VALUES ('2', 'Tom', 'Jones', '770', '264000',
'800000', '1');
INSERT INTO `go_db1`.`customer` (`id`, `first_name`, `last_name`, `credit_score`,
`salary`, `downpayment`, `house_id`) VALUES ('3', 'Nicole', 'Freedman', '750',
'203000', '850000', '2');
INSERT INTO `go_db1`.`customer` (`id`, `first_name`, `last_name`, `credit_score`,
`salary`, `downpayment`, `house_id`) VALUES ('4', 'June', 'Green', '700', '148000',
'600000', '4'); 