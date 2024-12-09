import PostForm from "@/components/dashboard/postForm";
import PostTable from "@/components/dashboard/postTable";
import Header from "@/components/header/header";
import { Suspense } from "react";

export default function Home() {
  return (
    <div className="max-h-dvh w-full grow text-white gap-16">
      <Header />
      <div className="flex flex-col md:flex-row w-full">
        <div className="flex-1 h-fit p-4 rounded-lg bg-stone-800 m-4">
          <Suspense>
            <PostTable />
          </Suspense>
        </div>
        <div className="flex-1 h-fit p-4 rounded-lg bg-stone-800 m-4 md:my-4 sm:my-0">
          <PostForm />
        </div>
      </div>
    </div>
  );
}
