-- Insertion
CREATE OR REPLACE FUNCTION public.handle_follow_insert()
RETURNS trigger
SECURITY DEFINER SET search_path = ''
AS $$
BEGIN
  UPDATE public.users SET follower_count = follower_count + 1 WHERE id = NEW.to_user_id;
  UPDATE public.users SET following_count = following_count + 1 WHERE id = NEW.from_user_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_follow_insert ON public.follows;

CREATE TRIGGER trg_follow_insert
AFTER INSERT ON public.follows
FOR EACH ROW
EXECUTE FUNCTION public.handle_follow_insert();

CREATE OR REPLACE FUNCTION public.handle_soft_delete_follow()
RETURNS trigger
SECURITY DEFINER SET search_path = ''
AS $$
BEGIN
  -- Soft delete
  IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
    UPDATE public.users SET follower_count = follower_count - 1 WHERE id = OLD.to_user_id;
    UPDATE public.users SET following_count = following_count - 1 WHERE id = OLD.from_user_id;

  -- Undo soft delete
  ELSIF OLD.deleted_at IS NOT NULL AND NEW.deleted_at IS NULL THEN
    UPDATE public.users SET follower_count = follower_count + 1 WHERE id = NEW.to_user_id;
    UPDATE public.users SET following_count = following_count + 1 WHERE id = NEW.from_user_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_soft_delete_follow ON public.follows;

CREATE TRIGGER trg_soft_delete_follow
BEFORE UPDATE ON public.follows
FOR EACH ROW
WHEN (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at)
EXECUTE FUNCTION public.handle_soft_delete_follow();

-- Hard delete
CREATE OR REPLACE FUNCTION public.handle_follow_delete()
RETURNS trigger
SECURITY DEFINER SET search_path = ''
AS $$
BEGIN
  UPDATE public.users SET follower_count = follower_count - 1 WHERE id = OLD.to_user_id;
  UPDATE public.users SET following_count = following_count - 1 WHERE id = OLD.from_user_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_follow_delete ON public.follows;

CREATE TRIGGER trg_follow_delete
AFTER DELETE ON public.follows
FOR EACH ROW
EXECUTE FUNCTION public.handle_follow_delete();

