CREATE TABLE `job_positions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `position` varchar(50) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `mbti` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` char(5) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `solutions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `solution` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `solutions_to_mbti` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_mbti` int(11) DEFAULT NULL,
  `id_solution` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_mbti` (`id_mbti`),
  KEY `id_solution` (`id_solution`),
  CONSTRAINT `solutions_to_mbti_ibfk_1` FOREIGN KEY (`id_mbti`) REFERENCES `mbti` (`id`),
  CONSTRAINT `solutions_to_mbti_ibfk_2` FOREIGN KEY (`id_solution`) REFERENCES `solutions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `users` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `fullname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `gender` char(1) NOT NULL,
  `birthdate` date NOT NULL,
  `address` varchar(255) NOT NULL,
  `role` char(20) NOT NULL,
  `id_mbti` int(11) DEFAULT NULL,
  `id_job_position` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `id_mbti` (`id_mbti`),
  KEY `id_job_position` (`id_job_position`),
  KEY `idx_fullname` (`fullname`),
  CONSTRAINT `users_ibfk_1` FOREIGN KEY (`id_mbti`) REFERENCES `mbti` (`id`),
  CONSTRAINT `users_ibfk_2` FOREIGN KEY (`id_job_position`) REFERENCES `job_positions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `physics` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `heart_rate` int(11) NOT NULL,
  `diastolic_blood_pressure` int(11) NOT NULL,
  `sistolic_blood_pressure` int(11) NOT NULL,
  `body_temp` float NOT NULL,
  `stress_level` int(11) DEFAULT NULL,
  `id_user` uuid DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `physics_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `notification_history` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `message` varchar(255) NOT NULL,
  `id_user` uuid DEFAULT NULL,
  `status` char(10) NOT NULL,
  `level` char(10) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `notification_history_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `logbook` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `logs` varchar(255) NOT NULL,
  `id_user` uuid DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `logbook_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `forecasting` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `id_physic` uuid DEFAULT NULL,
  `id_user` uuid DEFAULT NULL,
  `predicted_stress_level` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  KEY `id_physic` (`id_physic`),
  KEY `id_user` (`id_user`),
  CONSTRAINT `forecasting_ibfk_1` FOREIGN KEY (`id_physic`) REFERENCES `physics` (`id`),
  CONSTRAINT `forecasting_ibfk_2` FOREIGN KEY (`id_user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;