"use client";

import { PostService } from "@/lib/service/postService";

type Props = { postID: string };

export default function LikeButton({ postID }: Props) {
  const likePost = async () => {
    await PostService.likePost(postID);
  };
  return (
    <button className="bg-white text-black" onClick={likePost}>
      Like
    </button>
  );
}
