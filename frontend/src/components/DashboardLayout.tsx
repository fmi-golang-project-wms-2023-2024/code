"use client";

import { usePathname, useRouter } from "next/navigation";
import Cookies from "js-cookie";
import FeatherIcon from "feather-icons-react";
import React, { ReactNode, useState } from "react";

interface DashboardLayoutProps {
  children: ReactNode;
  sidebarLinks: Array<{ url: string; icon: React.ReactNode; label: string }>;
}

const DashboardLayout: React.FC<DashboardLayoutProps> = ({
  children,
  sidebarLinks,
}) => {
  const pathname = usePathname();
  const router = useRouter();

  const handleLogout = () => {
    Cookies.remove("accessToken");
    Cookies.remove("refreshToken");
    router.push("/login");
  };

  return (
    <div>
      <aside
        id="default-sidebar"
        className={`fixed top-0 left-0 z-40 w-64 h-screen transition-transform`}
        aria-label="Sidebar"
      >
        <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
          <ul className="space-y-2 font-medium">
            <button
              onClick={handleLogout}
              className="inline-flex items-center mt-2 ms-3 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
            >
              <FeatherIcon icon="log-out" className="mr-2" /> Log Out
            </button>
            {sidebarLinks.map((link, index) => (
              <li key={index}>
                <a
                  href={link.url}
                  className={`flex items-center p-2 rounded-lg hover:bg-gray-100 dark:hover:bg-gray-700 group ${
                    pathname == link.url ? "bg-gray-100 dark:bg-gray-700" : ""
                  }`}
                >
                  {link.icon}
                  <span className="ms-3">{link.label}</span>
                </a>
              </li>
            ))}
          </ul>
        </div>
      </aside>

      <div className="p-4 sm:ml-64">{children}</div>
    </div>
  );
};

export default DashboardLayout;
