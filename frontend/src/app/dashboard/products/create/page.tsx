"use client";

import { createProduct } from "@/src/api/api";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { toast } from "react-toastify";

interface ProductForm {
  sku: string;
  title: string;
  price: string;
  image: string;
  quantity: number;
}

const CreateProductPage: React.FC = () => {
  const router = useRouter();
  const [formData, setFormData] = useState<ProductForm>({
    sku: "",
    title: "",
    price: "",
    image: "",
    quantity: 0,
  });
  const [loading, setLoading] = useState<boolean>(false);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      await createProduct(formData);
      toast.success("Product created successfully!");

      setFormData({
        sku: "",
        title: "",
        price: "",
        image: "",
        quantity: 0,
      });

      router.push("/dashboard/products/list");
    } catch (error) {
      console.error("Error creating product:", error);
      toast.error("Error creating product. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-96">
        <h1 className={`text-3xl font-semibold mb-6`}>Create Product</h1>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label htmlFor="sku" className="block text-gray-200">
              SKU:
            </label>
            <input
              type="text"
              id="sku"
              name="sku"
              value={formData.sku}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="title" className="block text-gray-200">
              Title:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={formData.title}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="price" className="block text-gray-200">
              Price:
            </label>
            <input
              type="text"
              id="price"
              name="price"
              value={formData.price}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="image" className="block text-gray-200">
              Image URL:
            </label>
            <input
              type="text"
              id="image"
              name="image"
              value={formData.image}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="quantity" className="block text-gray-200">
              Quantity:
            </label>
            <input
              type="number"
              id="quantity"
              name="quantity"
              value={formData.quantity}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <button
            type="submit"
            className={`bg-blue-500 text-white rounded-md p-2 hover:bg-blue-200 transition duration-300 ${
              loading ? "opacity-50 cursor-not-allowed" : ""
            }`}
            disabled={loading}
          >
            {loading ? "Creating..." : "Create Product"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateProductPage;
