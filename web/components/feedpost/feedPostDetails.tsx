"use client";

import { Post } from "@/lib/api/types";
import { useState } from "react";
import { PostService } from "@/lib/service/client/postService";
import { Input } from "../ui/input";
import { useForm, SubmitHandler } from "react-hook-form";
import { CommentService } from "@/lib/service/client/commentService";
import { HeartIcon } from "lucide-react";
import CommentList from "./commentList";

type Props = {
  post: Post;
  initLiked: boolean;
};
type Inputs = {
  comment: string;
};

export default function FeedPostDetails({ post, initLiked }: Props) {
  const [liked, setLiked] = useState(initLiked);
  const [commentCount, setCommentCount] = useState(post.comment_count);
  const [showComments, setShowComments] = useState(false);
  const postID = post.id.toString();
  const { register, handleSubmit, reset } = useForm<Inputs>();

  let optimisticDiff = 0;
  if (!initLiked && liked) {
    optimisticDiff = 1;
  } else if (initLiked && !liked) {
    optimisticDiff = -1;
  }
  const likePost = async () => {
    setLiked((l) => !l);
    await PostService.likePost(postID);
  };
  const unlikePost = async () => {
    setLiked((l) => !l);
    await PostService.unlikePost(postID);
  };

  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    reset();
    if (document.activeElement instanceof HTMLElement) {
      document.activeElement.blur();
    }
    setCommentCount((c) => c + 1);
    await CommentService.createComment({ text: data.comment }, postID);
  };

  return (
    <div className="flex flex-col justify-start gap-1 mt-2">
      <HeartIcon
        onClick={liked ? unlikePost : likePost}
        className={`hover:cursor-pointer ${liked ? "fill-rose-700 stroke-rose-700" : ""}`}
      />
      <p>{post.like_count + optimisticDiff} Likes</p>
      <div>{post.caption}</div>
      <p
        className="text-muted-foreground hover:cursor-pointer"
        onClick={() => setShowComments((b) => !b)}
      >
        {showComments
          ? `Hide ${commentCount} comments`
          : `View all ${commentCount} comments`}
      </p>
      {showComments && <CommentList postID={postID} />}
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-row ">
        <Input
          {...register("comment", { required: true })}
          placeholder="Add a comment..."
          className="border-none focus-visible:ring-transparent"
          autoComplete="off"
        />
        <Input
          type="submit"
          value={"Post"}
          className="border-none w-auto hover:cursor-pointer hover:bg-neutral-900 text-muted-foreground hover:text-primary"
        />
      </form>
    </div>
  );
}
