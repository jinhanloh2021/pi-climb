ALTER TABLE public.follows ENABLE ROW LEVEL SECURITY;

-- Policy for INSERT: A user can create their own follows
CREATE POLICY "Users can create their own follow"
ON public.follows
FOR INSERT
TO backend_service_role
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = from_user_id);

-- Policy for SELECT: A user can SELECT all follows
CREATE POLICY "Users can read all follows"
ON public.follows
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- Policy for UPDATE: A user can UPDATE their OWN follow
CREATE POLICY "Users can update their own follow"
ON public.follows
FOR UPDATE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = from_user_id)
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = from_user_id);

-- Policy for DELETE: A user can DELETE their OWN follow
CREATE POLICY "Users can delete their own follow"
ON public.follows
FOR DELETE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = from_user_id)
