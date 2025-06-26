ALTER TABLE public.users ENABLE ROW LEVEL SECURITY;

-- Policy for SELECT: A user can SELECT all other profiles.
CREATE POLICY "Users can read all users"
ON public.users
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- Policy for UPDATE: A user can UPDATE their OWN profile.
CREATE POLICY "Users can update their own profile"
ON public.users
FOR UPDATE
TO backend_service_role
USING ((current_setting('app.current_user_id'::text, true))::uuid = supabase_id)
WITH CHECK ((current_setting('app.current_user_id'::text, true))::uuid = supabase_id);

-- Policy for DELETE: A user can DELETE their OWN profile.
CREATE POLICY "Users can delete their own profile"
ON public.users
FOR DELETE
TO backend_service_role
USING ((current_setting('app.current_user_id'::text, true))::uuid = supabase_id)
