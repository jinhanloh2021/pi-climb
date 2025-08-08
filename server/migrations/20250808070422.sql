ALTER TABLE public.comments ENABLE ROW LEVEL SECURITY;

-- Policy for INSERT: A user can create their own comment
CREATE POLICY "Users can create their own comment"
ON public.comments
FOR INSERT
TO backend_service_role
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for SELECT: A user can SELECT all comments
CREATE POLICY "Users can read all comments"
ON public.comments
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- Policy for UPDATE: A user can UPDATE their OWN comment
CREATE POLICY "Users can update their own comment"
ON public.comments
FOR UPDATE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for DELETE: A user can DELETE their OWN comment
CREATE POLICY "Users can delete their own comment"
ON public.comments
FOR DELETE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
