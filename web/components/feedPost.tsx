import { Media, Post } from "@/lib/api/types";
import { createClient } from "@/lib/supabase/server";
import Image from "next/image";
import FeedPostDetails from "./feedPostDetails";

type Props = {
  post: Post;
  liked: boolean;
};

export default async function FeedPost({ post, liked }: Props) {
  if (!post.media || post.media.length === 0) {
    return <div>no media content</div>;
  }

  const media: Media = post.media[0];
  const supabase = await createClient();
  // todo: Server-side infinite scrolling
  const { data, error } = await supabase.storage
    .from(media.bucket)
    .createSignedUrl(`${media.storage_key}`, 300);
  if (error) {
    return <div>Error loading media</div>;
  }

  return (
    <div>
      <p>{post.user?.username ?? post.user_id}</p>
      <Image src={data?.signedUrl ?? ""} width={500} height={500} alt="" />
      <FeedPostDetails post={post} initLiked={liked} />
    </div>
  );
}
