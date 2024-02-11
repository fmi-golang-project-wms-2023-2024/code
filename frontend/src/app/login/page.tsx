"use client";

import { useState, ChangeEvent, FormEvent } from "react";
import { loginUser } from "@/src/api/api";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";
import Cookies from "js-cookie";
import Link from "next/link";

interface LoginData {
  username: string;
  password: string;
}

export default function LoginPage() {
  const [loginData, setLoginData] = useState<LoginData>({
    username: "",
    password: "",
  });
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setLoginData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleFormSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);

    try {
      await new Promise((resolve) => setTimeout(resolve, 400));
      const { user, access_token, refresh_token } = await loginUser(
        loginData.username,
        loginData.password
      );

      // Store access and refresh tokens in cookies
      Cookies.set("accessToken", access_token, { path: "/" });
      Cookies.set("refreshToken", refresh_token, { path: "/" });

      toast.success(`Welcome back, ${user.username}!`);
    } catch (error) {
      toast.error("Invalid username or password. Please try again.");
    } finally {
      setLoading(false);
      router.push("/");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-96">
        <h1 className={`text-3xl font-semibold mb-6`}>Log In</h1>
        <form onSubmit={handleFormSubmit}>
          {["username", "password"].map((field) => (
            <div key={field} className="mb-4">
              <label htmlFor={field} className={`block text-sm font-medium `}>
                {field.charAt(0).toUpperCase() + field.slice(1)}
              </label>
              <input
                type={field === "password" ? "password" : "text"}
                id={field}
                name={field}
                className={`mt-1 p-2 border rounded-md w-full text-black`}
                //@ts-ignore
                value={loginData[field]}
                onChange={handleChange}
                required
              />
            </div>
          ))}
          <button
            type="submit"
            className={`w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 transition duration-300`}
            disabled={loading}
          >
            {loading ? (
              <div className="flex items-center justify-center">
                <span className="animate-spin mr-2">&#9696;</span>Logging in...
              </div>
            ) : (
              "Log In"
            )}
          </button>
        </form>
        <div className="mt-4 text-sm text-center">
          <p>
            No account?{" "}
            <Link href="/register" className="text-blue-500 hover:underline">
              Click here to register
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
