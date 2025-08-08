ALTER TABLE public.media ENABLE ROW LEVEL SECURITY;

-- Policy for INSERT: A user can create their own media
CREATE POLICY "Users can create their own media"
ON public.media
FOR INSERT
TO backend_service_role
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for SELECT: A user can SELECT all media
CREATE POLICY "Users can read all media"
ON public.media
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- Policy for UPDATE: A user can UPDATE their OWN media
CREATE POLICY "Users can update their own media"
ON public.media
FOR UPDATE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
WITH CHECK (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id);

-- Policy for DELETE: A user can DELETE their OWN media
CREATE POLICY "Users can delete their own media"
ON public.media
FOR DELETE
TO backend_service_role
USING (((SELECT current_setting('app.current_user_id'::text, true)))::uuid = user_id)
