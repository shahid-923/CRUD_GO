import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Verify from "./pages/Verify";
import Signup from "./pages/Signup";
import Cart from "./pages/Cart";
import Orders from "./pages/Orders";
import PostOrder from "./pages/PostOrder";
import FailedOrder from "./pages/FailedOrder";
import SellerProgram from "./pages/SellerProgram";
import ManageProducts from "./pages/ManageProducts";
import SellerOrderDetails from "./pages/SellerOrderDetails";
import ProductDetails from "./pages/ProductDetails";
import Profile from "./pages/Profile";
import LandingPage from "./pages/LandingPage";
import NotFound from "./pages/NotFound";

interface RouteManagerProps {}

const RouteManager: React.FC<RouteManagerProps> = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LandingPage />} />
        <Route path="/login" element={<Login />} />
        <Route path="/verify" element={<Verify />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/cart" element={<Cart />} />
        <Route path="/orders" element={<Orders />} />
        <Route path="/post-order" element={<PostOrder />} />
        <Route path="/failed-order" element={<FailedOrder />} />
        <Route path="/seller-program" element={<SellerProgram />} />
        <Route path="/manage-products" element={<ManageProducts />} />
        <Route path="/seller-order/:id" element={<SellerOrderDetails />} />
        <Route path="/product-details/:id" element={<ProductDetails />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
};

export default RouteManager;
