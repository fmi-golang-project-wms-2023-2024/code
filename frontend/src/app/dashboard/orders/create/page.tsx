"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { toast } from "react-toastify";
import {
  createOrder,
  listProducts,
  Order,
  OrderStatus,
  OrderLine,
  Product,
} from "@/src/api/api";

const CreateOrderPage: React.FC = () => {
  const router = useRouter();
  const [formData, setFormData] = useState<Order>({
    recipient_full_name: "",
    email_address: "",
    delivery_address: "",
    phone: "",
    status: OrderStatus.ORDER_STATUS_UNSPECIFIED,
    order_lines: [],
  });
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [selectedProducts, setSelectedProducts] = useState<{
    [key: string]: number;
  }>({});

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await listProducts();
        setProducts(response.products);
      } catch (error) {
        console.error("Error fetching products:", error);
      }
    };

    fetchProducts();
  }, []);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleProductSelection = (productId: string, quantity: number) => {
    setSelectedProducts((prevSelectedProducts) => ({
      ...prevSelectedProducts,
      [productId]: quantity,
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      // Create order_lines array based on selected products and quantities
      const orderLines: OrderLine[] = Object.entries(selectedProducts).map(
        ([productId, quantity]) => ({
          product_id: productId,
          quantity,
          price: products.find((product) => product.id === productId)
            ?.price as string,
        })
      );

      await createOrder({
        ...formData,
        order_lines: orderLines,
      });
      toast.success("Order created successfully!");

      setFormData({
        recipient_full_name: "",
        email_address: "",
        delivery_address: "",
        phone: "",
        status: 0,
        order_lines: [],
      });

      setSelectedProducts({}); // Reset selected products and quantities
      router.push("/dashboard/orders/list");
    } catch (error) {
      console.error("Error creating order:", error);
      toast.error("Error creating order. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-96">
        <h1 className={`text-3xl font-semibold mb-6`}>Create Order</h1>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label
              htmlFor="recipient_full_name"
              className="block text-gray-200"
            >
              Recipient Full Name:
            </label>
            <input
              type="text"
              id="recipient_full_name"
              name="recipient_full_name"
              value={formData.recipient_full_name}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="email_address" className="block text-gray-200">
              Email Address:
            </label>
            <input
              type="email"
              id="email_address"
              name="email_address"
              value={formData.email_address}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="delivery_address" className="block text-gray-200">
              Delivery Address:
            </label>
            <input
              type="text"
              id="delivery_address"
              name="delivery_address"
              value={formData.delivery_address}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="phone" className="block text-gray-200">
              Phone:
            </label>
            <input
              type="tel"
              id="phone"
              name="phone"
              value={formData.phone}
              onChange={handleInputChange}
              className="w-full border-gray-300 text-black rounded-md p-2"
              required
            />
          </div>
          <div className="mb-4">
            <label htmlFor="product" className="block text-gray-200">
              Select Products:
            </label>
            {products.map((product) => (
              <div key={product.id} className="flex items-center mb-2">
                <input
                  type="number"
                  min="0"
                  id={`quantity-${product.id}`}
                  name={`quantity-${product.id}`}
                  value={selectedProducts[product.id] || 0}
                  onChange={(e) =>
                    handleProductSelection(product.id, Number(e.target.value))
                  }
                  className="w-16 mr-2 border-gray-300 text-black rounded-md p-2"
                />
                <label htmlFor={`quantity-${product.id}`}>
                  {product.title} - ${product.price}
                </label>
              </div>
            ))}
          </div>
          <button
            type="submit"
            className={`bg-blue-500 text-white rounded-md p-2 hover:bg-blue-200 transition duration-300 ${
              loading ? "opacity-50 cursor-not-allowed" : ""
            }`}
            disabled={loading}
          >
            {loading ? "Creating..." : "Create Order"}
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateOrderPage;
