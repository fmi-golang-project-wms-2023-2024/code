"use client";

import { useState, useEffect } from "react";
import { toast } from "react-toastify";
import { User, getUsers, deleteUser } from "@/src/api/api";

export default function UserListPage() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const fetchedUsers = await getUsers();
        console.log(fetchUsers)
        setUsers(fetchedUsers);
        setError("");
      } catch (error) {
        setError("Error fetching users. Please try again.");
        toast.error("Error deleting users. Please try again.");
      } finally {
        setLoading(false);
      }
    };

    fetchUsers();
  }, []);

  const handleDeleteUser = async (userId: string) => {
    try {
      await deleteUser(userId);
      setUsers((prevUsers) => prevUsers.filter((user) => user.id !== userId));
      toast.success("User deleted successfully!");
    } catch (error) {
      setError("Error deleting user. Please try again.");
      toast.error("Error deleting user. Please try again.");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-96">
        <h1 className="text-3xl font-semibold mb-6">User List</h1>
        {loading && <p className="text-gray-500">Loading users...</p>}
        {users.length > 0 && (
          <table className="w-full">
            <thead>
              <tr>
                <th className="text-left">Name</th>
                <th className="text-left">Username</th>
                <th className="text-right">Actions</th>
              </tr>
            </thead>
            <tbody>
              {users.map((user) => (
                <tr key={user.id} className="border-b">
                  <td className="py-2">{user.full_name}</td>
                  <td className="py-2 text-gray-200">{user.username}</td>
                  <td className="text-right">
                    <button
                      onClick={() => handleDeleteUser(user.id)}
                      className="text-red-500 hover:text-red-700 transition duration-300"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        className="h-6 w-6"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth="2"
                          d="M6 18L18 6M6 6l12 12"
                        />
                      </svg>
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
        {users.length === 0 && !loading && !error && (
          <p className="text-gray-500">No users available.</p>
        )}
      </div>
    </div>
  );
}
