"use client";

import { PostService } from "@/lib/service/client/postService";
import { useState } from "react";

type Props = { postID: string; initLiked: boolean };

export default function LikeButton({ postID, initLiked }: Props) {
  const [liked, setLiked] = useState(initLiked);
  const likePost = async () => {
    setLiked((l) => !l);
    await PostService.likePost(postID);
  };
  const unlikePost = async () => {
    setLiked((l) => !l);
    await PostService.unlikePost(postID);
  };
  return (
    <button
      className="bg-white text-black"
      onClick={liked ? unlikePost : likePost}
    >
      {liked ? "Unlike" : "Like"}
    </button>
  );
}
