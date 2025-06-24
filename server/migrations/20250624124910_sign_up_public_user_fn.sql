-- Create or replace the function that will be triggered.
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
DECLARE
    generated_username VARCHAR(64);
BEGIN
    -- Generate a unique placeholder username
    generated_username := 'user_' || REPLACE(NEW.id::text, '-', '');
    INSERT INTO public.users (
        supabase_id,
        email,
        username,
        created_at,
        updated_at,
        is_public,
        follower_count,
        following_count
    )
    VALUES (
        NEW.id,             -- The UUID from auth.users
        NEW.email,          -- The email from auth.users
        generated_username, -- The generated unique username
        timezone('utc'::text, now()), -- Explicitly set created_at
        timezone('utc'::text, now()), -- Explicitly set updated_at
        TRUE,               -- Default for is_public
        0,                  -- Default for follower_count
        0                   -- Default for following_count
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;
-- SECURITY DEFINER runs function with permissions of function owener, not caller

-- Create the trigger to fire the function automatically on new user sign-ups.
CREATE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW EXECUTE FUNCTION public.handle_new_user();
