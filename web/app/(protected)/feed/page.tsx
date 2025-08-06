import { FeedService } from "@/lib/service/feedService";

export default async function FeedPage() {
  const res = await FeedService.getFeed("", "", 10);
  // TODO: Render posts in proper form
  return (
    <div className="flex-1 w-full flex flex-col gap-12">
      <div className="flex flex-col gap-2 items-start">
        <h2 className="font-bold text-2xl mb-4">Feed</h2>
        <pre className="text-xs font-mono p-3 rounded border max-h-32 overflow-auto">
          {JSON.stringify(res, null, 2)}
        </pre>
      </div>
    </div>
  );
}
