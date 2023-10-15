CREATE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE job_positions (
  id SERIAL PRIMARY KEY,
  position VARCHAR(50) NOT NULL,
  description VARCHAR(255)
);

CREATE TABLE mbti (
  id SERIAL PRIMARY KEY,
  personality CHAR(5) UNIQUE NOT NULL,
  description VARCHAR(255) DEFAULT NULL
);

CREATE TABLE solutions (
  id SERIAL PRIMARY KEY,
  solution VARCHAR(255) NOT NULL
);

CREATE TABLE solutions_to_mbti (
  id SERIAL PRIMARY KEY,
  id_mbti INT DEFAULT NULL,
  id_solution INT DEFAULT NULL,
  FOREIGN KEY (id_mbti) REFERENCES mbti (id),
  FOREIGN KEY (id_solution) REFERENCES solutions (id)
);

CREATE SEQUENCE users_id_seq
    START 3000;

CREATE TABLE users (
  id INT default nextval('users_id_seq') PRIMARY KEY,
  fullname VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  gender CHAR(1) NOT NULL,
  birthdate DATE NOT NULL,
  address VARCHAR(255) NOT NULL,
  role CHAR(20) NOT NULL,
  id_mbti INT,
  id_job_position INT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (id_mbti) REFERENCES mbti (id),
  FOREIGN KEY (id_job_position) REFERENCES job_positions (id)
);

CREATE TRIGGER set_timestamp
BEFORE
UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE SEQUENCE physics_id_sec
    START 100000;

CREATE TABLE physics (
  id INT default nextval('physics_id_sec') PRIMARY KEY,
  heart_rate int NOT NULL,
  diastolic_blood_pressure int NOT NULL,
  systolic_blood_pressure int NOT NULL,
  body_temp float NOT NULL,
  oxygen_saturation int NOT NULL,
  stress_level int DEFAULT NULL,
  id_user INT DEFAULT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp default now(),
  FOREIGN KEY (id_user) REFERENCES users (id)
);

CREATE TRIGGER set_timestamp
BEFORE
UPDATE ON physics
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE notification_history (
  id SERIAL PRIMARY KEY,
  message varchar(100) NOT NULL,
  id_user INT DEFAULT NULL,
  status char(10) NOT NULL,
  level char(10) DEFAULT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  FOREIGN KEY (id_user) REFERENCES users (id)
);

CREATE TRIGGER set_timestamp
BEFORE
UPDATE ON notification_history
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

create sequence logbook_id_sec start 20000;

CREATE TABLE logbook (
  id INT default nextval('logbook_id_sec') PRIMARY KEY,
  logs varchar(150) NOT NULL,
  id_user INT DEFAULT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  FOREIGN KEY (id_user) REFERENCES users (id)
);

CREATE TRIGGER set_timestamp
BEFORE
UPDATE ON logbook
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE forecasting (
  id SERIAL PRIMARY KEY,
  id_physic INT DEFAULT NULL,
  id_user INT DEFAULT NULL,
  predicted_stress_level int NOT NULL,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  FOREIGN KEY (id_physic) REFERENCES physics (id),
  FOREIGN KEY (id_user) REFERENCES users (id)
);

CREATE TRIGGER set_timestamp
BEFORE
UPDATE ON forecasting
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();