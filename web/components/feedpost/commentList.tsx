"use client";

import { CommentService } from "@/lib/service/client/commentService";
import { useEffect, useState } from "react";
import { Comment } from "@/lib/api/types"; // Assuming a Comment type exists
import { LoaderCircle } from "lucide-react";

type Props = { postID: string };

export default function CommentList({ postID }: Props) {
  const [comments, setComments] = useState<Comment[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // todo: Comment cursor pagination, infinite scrolling
  useEffect(() => {
    const fetchComments = async () => {
      try {
        setIsLoading(true);
        const fetchedComments = await CommentService.getComments(postID);
        setComments(fetchedComments);
      } catch (e: any) {
        setError(`Failed to fetch comments: ${e}`);
      } finally {
        setIsLoading(false);
      }
    };
    fetchComments();
  }, [postID]);

  if (isLoading) {
    return (
      <div className="flex flex-row justify-start align-middle gap-2 text-muted-foreground">
        <LoaderCircle className="animate-[spin_1s_linear_infinite]" />
        <p>Loading comments...</p>
      </div>
    );
  }

  if (error) {
    return <div>{error}</div>;
  }

  if (comments.length === 0) {
    return <div>No comments yet</div>;
  }

  return (
    <div>
      {comments.map((c) => (
        <div key={c.id}>{c.text}</div> // Assuming c.id exists
      ))}
    </div>
  );
}
