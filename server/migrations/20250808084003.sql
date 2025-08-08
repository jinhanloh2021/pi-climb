ALTER TABLE public.likes ENABLE ROW LEVEL SECURITY;

-- Policy for INSERT: A user can create their own like
CREATE POLICY "Users can create their own like"
ON public.likes
FOR INSERT
TO backend_service_role
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for SELECT: A user can SELECT all likes
CREATE POLICY "Users can read all likes"
ON public.likes
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- Policy for UPDATE: A user can UPDATE their own like
CREATE POLICY "Users can update their own like"
ON public.likes
FOR UPDATE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for DELETE: A user can DELETE their own like
CREATE POLICY "Users can delete their own like"
ON public.likes
FOR DELETE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
