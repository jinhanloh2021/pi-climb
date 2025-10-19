import FeedPost from "@/components/feedpost/feedPost";
import { FeedService } from "@/lib/service/server/feedService";
import { PostService } from "@/lib/service/server/postService";

export default async function FeedPage() {
  // TODO: Handle cursors
  const posts = await FeedService.getFeed("", "", 10);
  return (
    <div className="flex flex-col gap-12">
      <h2 className="font-bold text-2xl mb-4">Feed</h2>
      {posts.posts?.map(async (p) => {
        const liked = await PostService.myLikedPost(p.id.toString());
        return <FeedPost key={p.id} post={p} liked={liked.liked} />;
      })}
      {(!posts.posts || posts.posts.length == 0) && <p>No posts</p>}
    </div>
  );
}
