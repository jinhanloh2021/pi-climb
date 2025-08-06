import FeedPost from "@/components/feedPost";
import { FeedService } from "@/lib/service/feedService";

export default async function FeedPage() {
  const res = await FeedService.getFeed("", "", 10);
  return (
    <div className="flex flex-col gap-12 lg:w-[50%]">
      <h2 className="font-bold text-2xl mb-4">Feed</h2>
      {res.posts.map((p) => (
        <FeedPost key={p.id} post={p} />
      ))}
    </div>
  );
}
