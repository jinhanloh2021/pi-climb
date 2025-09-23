"use client";

import { Input } from "@/components/ui/input";
import { CreatePostRequest } from "@/lib/api/types";
import { PostService } from "@/lib/service/client/postService";
import { createClient } from "@/lib/supabase/client";
import { SubmitHandler, useForm } from "react-hook-form";
import { v4 as uuidV4 } from "uuid";

type Inputs = {
  caption: string;
  holdColour: string;
  grade: string;
  gymId: number;
  media: FileList;
};

export default function NewPost() {
  const { register, handleSubmit } = useForm<Inputs>();
  const supabase = createClient();
  const onSubmit: SubmitHandler<Inputs> = async (formData) => {
    // TODO: Error handling, input validation
    // Check valid media, valid filename, file ext, valid user and userId
    // Support for multiple media upload and video

    // upload to supabase
    console.dir(formData);
    const {
      data: { user },
    } = await supabase.auth.getUser();
    const fileName = formData.media[0].name;
    const genFileName = `${uuidV4()}.${fileName.split(".").pop()}`;
    const storageKey = `${user?.id}/${genFileName}`;

    const { data, error } = await supabase.storage
      .from("media")
      .upload(storageKey, formData.media[0]);
    if (error) {
      console.error(error);
      return;
    }
    console.dir(data);

    const post: CreatePostRequest = {
      caption: formData.caption,
      hold_colour: formData.holdColour,
      grade: formData.grade,
      media: [
        {
          storage_key: storageKey,
          bucket: "media",
          original_name: fileName,
          file_size: formData.media[0].size,
          mime_type: formData.media[0].type,
          order: 0,
          width: 0,
          height: 0,
        },
      ],
      // gym_id: formData.gymId, // currently not supported
    };
    console.dir(post);

    const createdPost = await PostService.createPost(post);
    console.dir(createdPost);
    return;
  };

  return (
    <div className="w-full lg:px-[25%]">
      New Post
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-5">
        <Input
          {...register("media")}
          type="file"
          className="hover:cursor-pointer text-secondary-foreground"
        />
        <Input // should use text area
          {...register("caption", { required: true })}
          placeholder="Caption"
          className="w-full"
          autoComplete="off"
        />
        <Input
          {...register("holdColour")}
          placeholder="Hold Colour"
          className="w-full"
          autoComplete="off"
        />
        <Input
          {...register("grade")}
          placeholder="Grade"
          className="w-full"
          autoComplete="off"
        />
        <Input
          {...register("gymId", { valueAsNumber: true })}
          placeholder="Gym"
          className="w-full"
          autoComplete="off"
        />
        <Input
          type="submit"
          value={"Post"}
          className="border-none hover:cursor-pointer hover:bg-neutral-900 w-min"
        />
      </form>
    </div>
  );
}
