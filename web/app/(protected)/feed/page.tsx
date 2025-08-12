import FeedPost from "@/components/feedPost";
import { FeedService } from "@/lib/service/server/feedService";
import { PostService } from "@/lib/service/server/postService";

export default async function FeedPage() {
  const posts = await FeedService.getFeed("", "", 10); // todo: Handle cursors
  return (
    <div className="flex flex-col gap-12">
      <h2 className="font-bold text-2xl mb-4">Feed</h2>
      {posts.posts.map(async (p) => {
        const liked = await PostService.myLikedPost(p.id.toString());
        return <FeedPost key={p.id} post={p} liked={liked.liked} />;
      })}
    </div>
  );
}
