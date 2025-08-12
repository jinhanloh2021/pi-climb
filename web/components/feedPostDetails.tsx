"use client";

import { Post } from "@/lib/api/types";
import { useState } from "react";
import { PostService } from "@/lib/service/client/postService";
import { Input } from "./ui/input";
import { useForm, SubmitHandler } from "react-hook-form";
import { CommentService } from "@/lib/service/client/commentService";
import { HeartIcon } from "lucide-react";

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
  const postID = post.id.toString();
  const {
    register,
    handleSubmit,
    reset,
    watch,
    formState: { errors },
  } = useForm<Inputs>();

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
    <>
      <div className="flex flex-row justify-start gap-2 my-2">
        <HeartIcon />
        <button
          className="bg-white text-black"
          onClick={liked ? unlikePost : likePost}
        >
          {liked ? "Unlike" : "Like"}
        </button>
        {/* <button className="bg-white text-black">Reply</button> */}
      </div>
      <p>{post.like_count + optimisticDiff} Likes</p>
      <div>{post.caption}</div>
      <p className="text-muted-foreground hover:cursor-pointer">{`View all ${commentCount} comments`}</p>
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
          className="border-none w-auto hover:cursor-pointer hover:bg-neutral-900 text-muted-foreground"
        />
      </form>
    </>
  );
}
