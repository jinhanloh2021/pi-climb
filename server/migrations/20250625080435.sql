ALTER TABLE public.users ENABLE ROW LEVEL SECURITY;

-- Policy for SELECT: A user can SELECT their OWN profile.
CREATE POLICY "Users can select their own profile"
ON public.users
FOR SELECT
USING ((select auth.uid()) = supabase_id);

-- Policy for UPDATE: A user can UPDATE their OWN profile.
CREATE POLICY "Users can update their own profile"
ON public.users
FOR UPDATE
USING ((select auth.uid()) = supabase_id)
WITH CHECK ((select auth.uid()) = supabase_id);

-- Policy for DELETE: A user can DELETE their OWN profile.
CREATE POLICY "Users can delete their own profile"
ON public.users
FOR DELETE
USING ((select auth.uid()) = supabase_id);
