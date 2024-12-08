"use client";

import useApiQuery from "@/hooks/useApiQuery";

export default function Dashboard() {
  useApiQuery({ key: "health" });
  return <div>Dashboard</div>;
}
