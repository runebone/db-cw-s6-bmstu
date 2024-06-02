CREATE OR REPLACE FUNCTION adjust_trust_level()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.status IS NULL THEN
        IF NEW.status = 'approved' THEN
            UPDATE user_data
            SET trust_level = trust_level + 5
            WHERE id = NEW.done_by;
        ELSE
            UPDATE user_data
            SET trust_level = trust_level - 4
            WHERE id = NEW.done_by;
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_struct_annotation
AFTER UPDATE ON struct_annotation
FOR EACH ROW
EXECUTE FUNCTION adjust_trust_level();

CREATE TRIGGER trigger_term_annotation
AFTER UPDATE ON term_annotation
FOR EACH ROW
EXECUTE FUNCTION adjust_trust_level();
