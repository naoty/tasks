SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `statuses` (
  `status_id` BIGINT (20) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR (255) NOT NULL,
  `position` INT (20) NOT NULL DEFAULT 1,
  PRIMARY KEY (`status_id`)
) ENGINE = InnoDB, DEFAULT CHARACTER SET = utf8mb4, DEFAULT COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE `tasks` (
  `task_id` BIGINT (20) NOT NULL AUTO_INCREMENT,
  `status_id` BIGINT (20) DEFAULT NULL,
  `title` VARCHAR (255) NOT NULL,
  `position` INT (20) NOT NULL,
  PRIMARY KEY (`task_id`),
  INDEX `index_tasks_on_status_id` (`status_id`),
  INDEX `tasks_ibfk_1` (`status_id`),
  CONSTRAINT `tasks_ibfk_1` FOREIGN KEY (`status_id`) REFERENCES `statuses` (`status_id`) ON DELETE SET NULL
) ENGINE = InnoDB, DEFAULT CHARACTER SET = utf8mb4, DEFAULT COLLATE = utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
