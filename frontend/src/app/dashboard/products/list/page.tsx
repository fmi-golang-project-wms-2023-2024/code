'use client'

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import Image from 'next/image'
import { listProducts, deleteProduct } from "@/src/api/api";
import { toast } from "react-toastify";

interface Product {
  id: string;
  sku: string;
  title: string;
  price: string;
  image: string;
  quantity: number;
}

const ProductListPage: React.FC = () => {
  const router = useRouter();
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    // Fetch products when the component mounts
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    try {
      const productList = await listProducts();
      setProducts(productList.products);
    } catch (error) {
      console.error("Error fetching products:", error);
      toast.error("Error fetching products. Please try again.");
    }
  };

  const handleDelete = async (productId: string) => {
    setLoading(true);

    try {
      await deleteProduct(productId);
      toast.success("Product deleted successfully!");
      // Update the product list after deletion
      fetchProducts();
    } catch (error) {
      console.error("Error deleting product:", error);
      toast.error("Error deleting product. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-96">
        <h1 className={`text-3xl font-semibold mb-6`}>Product List</h1>
        <div>
          {products.map((product) => (
            <div key={product.id} className="mb-4 flex items-center">
              <div className="flex-shrink-0">
                <img
                  src={product.image}
                  width={100}
                  height={100}
                  style={{objectFit: "contain"}}
                  alt={product.title}
                  className="w-16 h-16 object-cover rounded-md mr-4"
                />
              </div>
              <div>
                <p className="text-lg font-semibold">{product.title}</p>
                <p className="text-gray-500">SKU: {product.sku}</p>
                <p className="text-gray-500">Price: {product.price}</p>
                <p className="text-gray-500">Quantity: {product.quantity}</p>
                <button
                  className="text-red-500 hover:text-red-700"
                  onClick={() => handleDelete(product.id)}
                  disabled={loading}
                >
                  {loading ? "Deleting..." : "Delete"}
                </button>
              </div>
            </div>
          ))}
        </div>
        <button
          onClick={() => router.push("/dashboard/products/create")}
          className="mt-4 bg-blue-500 text-white rounded-md p-2 hover:bg-blue-200 transition duration-300"
        >
          Create Product
        </button>
      </div>
    </div>
  );
};

export default ProductListPage;
