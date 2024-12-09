import React from "react";
import LogoutButton from "../ui/logoutButton";

export default function Header() {
  return (
    <header className="bg-stone-900 text-white py-4 px-8 flex justify-between items-center">
      <h1 className="text-xl font-bold">Terra</h1>
      <LogoutButton />
    </header>
  );
}
