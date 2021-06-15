-- CREATE TABLE IF NOT EXISTS `product`(
--    `product_id` INT UNSIGNED AUTO_INCREMENT,
--    `product_name` VARCHAR(100) NOT NULL,
--    `product_number` INT NOT NULL,
--    PRIMARY KEY ( `product_id` )
-- )ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- -- 
CREATE TABLE IF NOT EXISTS `sec_kill`(
   `sec_kill_id` INT UNSIGNED AUTO_INCREMENT,
   `product` INT,
   `start_time` DATE,
    `end_time` DATE,
   PRIMARY KEY ( `sec_kill_id` ),
   foreign key(product) references product(product_id) on delete cascade
)ENGINE=InnoDB DEFAULT CHARSET=utf8;