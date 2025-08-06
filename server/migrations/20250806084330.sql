-- Insertion
CREATE OR REPLACE FUNCTION public.handle_comment_insert()
RETURNS trigger
SECURITY DEFINER SET search_path = ''
AS $$
BEGIN
  UPDATE public.posts SET comment_count = comment_count + 1 WHERE id = NEW.post_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_comment_insert ON public.comments;

CREATE TRIGGER trg_comment_insert
AFTER INSERT ON public.comments
FOR EACH ROW
EXECUTE FUNCTION public.handle_comment_insert();

-- Soft Delete and undo Soft Delete
CREATE OR REPLACE FUNCTION public.handle_soft_delete_comment()
RETURNS trigger
SECURITY DEFINER SET search_path = ''
AS $$
BEGIN
  -- Soft delete
  IF OLD.deleted_at IS NULL AND NEW.deleted_at IS NOT NULL THEN
    UPDATE public.posts SET comment_count = comment_count - 1 WHERE id = OLD.post_id;

  -- Undo soft delete
  ELSIF OLD.deleted_at IS NOT NULL AND NEW.deleted_at IS NULL THEN
    UPDATE public.posts SET comment_count = comment_count + 1 WHERE id = NEW.post_id;
  END IF;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_soft_delete_comment ON public.comments;

CREATE TRIGGER trg_soft_delete_comment
BEFORE UPDATE ON public.comments
FOR EACH ROW
WHEN (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at)
EXECUTE FUNCTION public.handle_soft_delete_comment();

-- Hard delete
CREATE OR REPLACE FUNCTION public.handle_comment_delete()
RETURNS trigger
SECURITY DEFINER SET search_path = ''
AS $$
BEGIN
  UPDATE public.posts SET comment_count = comment_count - 1 WHERE id = OLD.post_id;
  RETURN OLD;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_comment_delete ON public.posts;

CREATE TRIGGER trg_comment_delete
AFTER DELETE ON public.comments
FOR EACH ROW
EXECUTE FUNCTION public.handle_comment_delete();

