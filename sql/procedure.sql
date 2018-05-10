CREATE OR REPLACE FUNCTION MERGE_USERS
  (
    _master_id bigint,
    _slave_id bigint
  )
  RETURNS void AS
$BODY$
DECLARE
  rec RECORD;
BEGIN

  FOR rec IN
  SELECT
    up.id,
    up.allow_for_matching,
    up.source_id,
    up.user_id,
    up.parameter_id,
    up.created_at,
    up.value,
    up.updated_at,
    up.actual_create_date
  FROM user_parameters up
  WHERE up.user_id = _slave_id
  LOOP

    DELETE FROM user_parameters WHERE id = rec.id;
    INSERT INTO
      user_parameters(
        "allow_for_matching",
        "source_id",
        "user_id",
        "parameter_id",
        "created_at",
        "value",
        "updated_at",
        "actual_create_date"
      )
    VALUES
      (
        rec.allow_for_matching,
        rec.source_id,
        _master_id,
        rec.parameter_id,
        rec.created_at,
        rec.value,
        rec.updated_at,
        rec.actual_create_date
      )
    ON CONFLICT
    ("value", "parameter_id", "user_id", "source_id")
      DO NOTHING
  ;

  END LOOP;

  DELETE FROM users where id = _slave_id;
END
$BODY$
LANGUAGE 'plpgsql'