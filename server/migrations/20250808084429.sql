ALTER TABLE public.gyms ENABLE ROW LEVEL SECURITY;

-- Policy for SELECT: A user can SELECT all gyms
CREATE POLICY "Users can read all gyms"
ON public.gyms
AS PERMISSIVE
FOR SELECT
TO backend_service_role
USING (true);

-- No permissions allowed for create, update and delete
