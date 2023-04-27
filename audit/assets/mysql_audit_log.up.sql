SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0;
SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE =
        'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Table `audit_log`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `audit_log`
(
    `id`            INT                                 NOT NULL AUTO_INCREMENT,
    `user_id`       INT                                 NOT NULL,
    `resource_id`   INT                                 NOT NULL,
    `action_type`   ENUM ("create", "update", "delete") NOT NULL,
    `resource_type` VARCHAR(45)                         NULL,
    `resource`      JSON                                NULL,
    `resource_prev` JSON                                NULL,
    `time`          DATETIME                            NOT NULL,
    PRIMARY KEY (`id`, `user_id`),
    INDEX `fk_audit_log_users1_idx` (`user_id` ASC),
    CONSTRAINT `fk_audit_log_users1`
        FOREIGN KEY (`user_id`)
            REFERENCES `users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;


SET SQL_MODE = @OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;
