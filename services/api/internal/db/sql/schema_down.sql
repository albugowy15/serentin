DROP TABLE IF EXISTS forecasting;
DROP TABLE IF EXISTS logbook;
DROP TABLE IF EXISTS notification_history;
DROP TABLE IF EXISTS physics;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS solutions_to_mbti;
DROP TABLE IF EXISTS solutions;
DROP TABLE IF EXISTS mbti;
DROP TABLE IF EXISTS job_positions;

DROP FUNCTION IF EXISTS public.trigger_set_timestamp();
DROP SEQUENCE IF EXISTS public.logbook_id_sec;
DROP SEQUENCE IF EXISTS public.physics_id_sec;
DROP SEQUENCE IF EXISTS public.users_id_sec;
