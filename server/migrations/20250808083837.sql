ALTER TABLE public.posts ENABLE ROW LEVEL SECURITY;

-- Policy for INSERT: A user can create their own post
CREATE POLICY "Users can create their own post"
ON public.posts
FOR INSERT
TO backend_service_role
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for SELECT: A user can SELECT all posts
CREATE POLICY "Users can read all posts"
ON public.posts
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- Policy for UPDATE: A user can UPDATE their OWN post
CREATE POLICY "Users can update their own post"
ON public.posts
FOR UPDATE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for DELETE: A user can DELETE their OWN post
CREATE POLICY "Users can delete their own post"
ON public.posts
FOR DELETE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
