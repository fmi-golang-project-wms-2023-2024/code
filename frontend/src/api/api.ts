import axios, { AxiosInstance, AxiosResponse } from "axios";
import Cookies from "js-cookie";

export interface User {
  id: string;
  full_name: string;
  username: string;
  password: string;
}

interface CreateUserRequest {
  user: Omit<User, "id">;
}

interface CreateUserResponse {
  user: User;
}

const axiosInstance: AxiosInstance = axios.create({
  baseURL: "http://localhost:8090",
});

axiosInstance.interceptors.request.use(
  (config) => {
    const token = Cookies.get("accessToken");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export async function createUser(
  user: CreateUserRequest
): Promise<CreateUserResponse> {
  try {
    const response: AxiosResponse<CreateUserResponse> =
      await axiosInstance.post("/v1/users", user);
    console.log("Create User Response:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error creating user:", error);
    throw error;
  }
}

interface LoginUserResponse {
  user: User;
  access_token: string;
  refresh_token: string;
}

export async function loginUser(
  username: string,
  password: string
): Promise<LoginUserResponse> {
  try {
    const response: AxiosResponse<LoginUserResponse> = await axiosInstance.post(
      "/v1/users/auth",
      {
        username,
        password,
      }
    );
    console.log("Login User Response:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error creating user:", error);
    throw error;
  }
}

export async function getUsers(): Promise<User[]> {
  try {
    const response: AxiosResponse<{ user: User[] }> = await axiosInstance.get(
      "/v1/users"
    );
    console.log("Get Users Response:", response.data.user);
    return response.data.user;
  } catch (error) {
    console.error("Error getting users:", error);
    throw error;
  }
}

interface DeleteUserResponse {
  message: string;
}

export async function deleteUser(userId: string): Promise<DeleteUserResponse> {
  try {
    const response: AxiosResponse<DeleteUserResponse> =
      await axiosInstance.delete(`/v1/users/${userId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
}

export interface Product {
  id: string;
  sku: string;
  title: string;
  price: string;
  image: string;
  quantity: number;
}

interface CreateProductResponse {
  product: Product;
}

interface ListProductsResponse {
  products: Product[];
}
interface DeleteProductResponse {}

export async function createProduct(
  product: Omit<Product, "id">
): Promise<CreateProductResponse> {
  try {
    const response: AxiosResponse<CreateProductResponse> =
      await axiosInstance.post("/v1/products", {
        product,
      });
    console.log("Create Product Response:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error creating product:", error);
    throw error;
  }
}

export async function listProducts(): Promise<ListProductsResponse> {
  try {
    const response: AxiosResponse<ListProductsResponse> =
      await axiosInstance.get("/v1/products");
    console.log("List Products Response:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error listing products:", error);
    throw error;
  }
}

export async function deleteProduct(
  productId: string
): Promise<DeleteProductResponse> {
  try {
    const response: AxiosResponse<DeleteProductResponse> =
      await axiosInstance.delete(`/v1/products/${productId}`);
    return response.data;
  } catch (error) {
    throw error;
  }
}

export interface Order {
  id?: string;
  recipient_full_name?: string;
  email_address?: string;
  delivery_address?: string;
  phone?: string;
  status?: OrderStatus;
  order_lines?: OrderLine[];
}

export interface OrderLine {
  id?: string;
  order_id?: string;
  product_id?: string;
  price?: string;
  quantity?: number;
}

export enum OrderStatus {
  ORDER_STATUS_UNSPECIFIED = 0,
  ORDER_STATUS_PENDING = 1,
  ORDER_STATUS_PROCESSING = 2,
  ORDER_STATUS_COMPLETED = 3,
  ORDER_STATUS_CANCELLED = 4,
}

export interface CreateOrderRequest {
  order: Omit<Order, "id">;
}

export interface CreateOrderResponse {
  order: Order;
}

export interface CreateOrderLineRequest {
  orderLine: Omit<OrderLine, "id">;
}

export interface CreateOrderLineResponse {
  orderLine: OrderLine;
}

export interface GetOrderRequest {
  id: string;
}

export interface GetOrderResponse {
  order: Order;
}

export interface DeleteOrderRequest {
  id: string;
}

export interface DeleteOrderResponse {}

export interface ListOrdersRequest {}

export interface ListOrdersResponse {
  orders: Order[];
}

export async function createOrder(
  order: Omit<Order, "id">
): Promise<CreateOrderResponse> {
  try {
    const response: AxiosResponse<CreateOrderResponse> =
      await axiosInstance.post("/v1/orders", {
        order,
      });
    console.log("Create Order Response:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error creating order:", error);
    throw error;
  }
}

export async function deleteOrder(id: string): Promise<DeleteOrderResponse> {
  try {
    const response: AxiosResponse<DeleteOrderResponse> =
      await axiosInstance.delete(`/v1/orders/${id}`);
    return response.data;
  } catch (error) {
    throw error;
  }
}

export async function listOrders(): Promise<ListOrdersResponse> {
  try {
    const response: AxiosResponse<ListOrdersResponse> = await axiosInstance.get(
      "/v1/orders"
    );
    console.log("List Orders Response:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error listing orders:", error);
    throw error;
  }
}
