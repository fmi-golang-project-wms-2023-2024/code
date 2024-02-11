"use client";

import { useState, ChangeEvent, FormEvent } from "react";
import { createUser } from "@/src/api/api";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";

interface UserData {
  username: string;
  password: string;
  full_name: string;
  role: string;
}

export default function CreateUserPage() {
  const [userData, setUserData] = useState<UserData>({
    username: "",
    password: "",
    full_name: "",
    role: "",
  });
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setUserData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleFormSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);

    try {
      await new Promise((resolve) => setTimeout(resolve, 400));
      await createUser({ user: userData });
      toast.success("User created successfully!");
    } catch (error) {
      toast.error("Error creating user. Please try again.");
    } finally {
      setLoading(false);
      router.push("/login");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-96">
        <h1 className={`text-3xl font-semibold mb-6`}>Create User</h1>
        <form onSubmit={handleFormSubmit}>
          {["username", "password", "full_name", "role"].map((field) => (
            <div key={field} className="mb-4">
              <label htmlFor={field} className={`block text-sm font-medium`}>
                {field.charAt(0).toUpperCase() + field.slice(1)}
              </label>
              {field === "role" ? (
                <select
                  id={field}
                  name={field}
                  className={`mt-1 p-2 border rounded-md w-full text-black`}
                  value={userData[field]}
                  onChange={handleChange}
                  required
                >
                  <option value="staff">Staff</option>
                  <option value="admin">Admin</option>
                </select>
              ) : (
                <input
                  type={field === "password" ? "password" : "text"}
                  id={field}
                  name={field}
                  className={`mt-1 p-2 border rounded-md w-full text-black`}
                  //@ts-ignore
                  value={userData[field]}
                  onChange={handleChange}
                  required
                />
              )}
            </div>
          ))}
          <button
            type="submit"
            className={`w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 transition duration-300`}
            disabled={loading}
          >
            {loading ? (
              <div className="flex items-center justify-center">
                <span className="animate-spin mr-2">&#9696;</span>Creating...
              </div>
            ) : (
              "Create User"
            )}
          </button>
        </form>
      </div>
    </div>
  );
}
