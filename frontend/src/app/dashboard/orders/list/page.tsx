"use client";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import {
  listOrders,
  deleteOrder,
  Order,
  Product,
  listProducts,
  OrderLine,
} from "@/src/api/api";
import { toast } from "react-toastify";

const OrderListPage: React.FC = () => {
  const router = useRouter();
  const [products, setProducts] = useState<Product[]>([]);
  const [orders, setOrders] = useState<Order[]>([]);
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    fetchProducts();
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {
      const ordersList = await listOrders();
      setOrders(ordersList.orders);
    } catch (error) {
      console.error("Error fetching orders:", error);
      toast.error("Error fetching orders. Please try again.");
    }
  };

  const fetchProducts = async () => {
    try {
      const response = await listProducts();
      setProducts(response.products);
    } catch (error) {
      console.error("Error fetching products:", error);
    }
  };

  const handleDelete = async (orderId: string) => {
    setLoading(true);

    try {
      await deleteOrder(orderId);
      toast.success("Order deleted successfully!");
      fetchOrders();
    } catch (error) {
      console.error("Error deleting order:", error);
      toast.error("Error deleting order. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  const renderProducts = (products: Product[]) => {
    return products.map((product) => (
      <div key={product.id} className="mb-4 flex items-center">
        <div className="flex-shrink-0">
          <img
            src={product.image}
            width={100}
            height={100}
            style={{ objectFit: "contain" }}
            alt={product.title}
            className="w-16 h-16 object-cover rounded-md mr-4"
          />
        </div>
        <div>
          <p className="text-lg font-semibold">{product.title}</p>
          <p className="text-gray-500">SKU: {product.sku}</p>
          <p className="text-gray-500">Price: {product.price}</p>
          <p className="text-gray-500">Quantity: {product.quantity}</p>
        </div>
      </div>
    ));
  };

  const getProducts = (orderLines?: OrderLine[]): Product[] => {
    if (!orderLines) return [];
    const productsForOrderLines: Product[] = orderLines.map((orderLine) => {
      const product = products.find((product) => product.id === orderLine.product_id);
      if (!product) return null;
  
      return {
        ...product,
        price: orderLine.price,
        quantity: orderLine.quantity,
      };
    }).filter((product): product is Product => product !== null);
  
    return productsForOrderLines;
  };

  return (
    <div className="min-h-screen flex items-center justify-center dark:bg-black">
      <div className="bg-white dark:bg-gray-800 shadow-md rounded-md p-8 w-full max-w-screen-md">
        {loading ? (
          <p className="text-gray-500">Loading orders...</p>
        ) : (
          <div>
            <div className="flex items-center justify-between mb-6">
              <h1 className="text-3xl font-semibold">Order List</h1>
              <button
                onClick={() => router.push("/dashboard/orders/create")}
                className="bg-blue-500 text-white rounded-md p-2 hover:bg-blue-200 transition duration-300"
              >
                Create Order
              </button>
            </div>
            {orders.map((order) => (
              <div
                key={order.id}
                className="border border-gray-300 rounded-md p-4 mb-6"
              >
                <div>
                  <p className="text-lg font-semibold">
                    Full name: {order?.recipient_full_name}
                  </p>
                  <p className="text-gray-500">
                    Email Address: {order?.email_address}
                  </p>
                  <p className="text-gray-500">Phone: {order?.phone}</p>
                  <p className="text-gray-500">
                    Delivery address: {order?.delivery_address}
                  </p>
                  <p className="text-gray-500">Order Status: {order?.status}</p>
                </div>
                <div className="mt-4">
                  {renderProducts(getProducts(order.order_lines))}
                </div>
                <button
                  className="text-red-500 hover:text-red-700 mt-4"
                  onClick={() => handleDelete(order.id as string)}
                  disabled={loading}
                >
                  {loading ? "Deleting..." : "Delete Order"}
                </button>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default OrderListPage;
