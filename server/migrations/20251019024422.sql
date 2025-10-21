DROP TRIGGER IF EXISTS on_auth_user_created ON auth.users;

CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER
LANGUAGE plpgsql
SECURITY DEFINER SET search_path = ''
AS $$
DECLARE
    generated_username VARCHAR(64);
BEGIN
    INSERT INTO public.users (
        id,
        email,
        created_at,
        updated_at,
        is_public,
        follower_count,
        following_count
    )
    VALUES (
        NEW.id,             -- The UUID from auth.users
        NEW.email,          -- The email from auth.users
        timezone('utc'::text, now()), -- Explicitly set created_at
        timezone('utc'::text, now()), -- Explicitly set updated_at
        TRUE,               -- Default for is_public
        0,                  -- Default for follower_count
        0                   -- Default for following_count
    );
    RETURN NEW;
END;
$$;

CREATE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW EXECUTE FUNCTION public.handle_new_user();
