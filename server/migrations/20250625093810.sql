CREATE POLICY "Enable read access for all users"
ON public.users
AS PERMISSIVE
FOR SELECT
TO authenticated
USING (true);
