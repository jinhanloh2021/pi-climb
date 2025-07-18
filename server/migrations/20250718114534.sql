-- Create private Media Bucket
INSERT INTO storage.buckets (id, name, public)
VALUES ('media', 'media', false)
ON CONFLICT (id) DO NOTHING;

-- Create public Avatar Bucket
INSERT INTO storage.buckets (id, name, public)
VALUES ('avatar', 'avatar', true)
ON CONFLICT (id) DO NOTHING;

-- MEDIA
create policy "Allow user upload to their media folder"
on storage.objects
for insert
to authenticated
with check (
  bucket_id = 'media'
  AND (storage.foldername(name))[1] = (select auth.uid()::text)
);

create policy "Allow user update their media folder"
on storage.objects
for update
to authenticated
with check (
  bucket_id = 'media'
  AND (storage.foldername(name))[1] = (select auth.uid()::text)
);

create policy "Allow user to delete their media files"
on storage.objects
for delete
to authenticated
using (
  bucket_id = 'media'
  AND owner_id = (select auth.uid()::text)
);

create policy "Allow all authenticated users to view media"
on storage.objects
for select
to authenticated
using (
  bucket_id = 'media'
);

-- AVATAR
create policy "Allow user upload their avatar folder"
on storage.objects
for insert
to authenticated
with check (
  bucket_id = 'avatar'
  AND (storage.foldername(name))[1] = (select auth.uid()::text)
);

create policy "Allow user update their avatar folder"
on storage.objects
for update
to authenticated
with check (
  bucket_id = 'avatar'
  AND (storage.foldername(name))[1] = (select auth.uid()::text)
);

create policy "Allow user to delete their avatar files"
on storage.objects
for delete
to authenticated
using (
  bucket_id = 'avatar'
  AND owner_id = (select auth.uid()::text)
);
