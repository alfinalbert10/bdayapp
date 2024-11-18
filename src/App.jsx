import React, { useState } from "react";
import Navbar from "./components/Nav";
import Dashboard from "./components/dashboard";
import Welcome from "./components/welcome";
import "./App.css";
import { Outlet } from "react-router-dom";

const App = () => {
 
  return (
    <div className="app">
      <nav>   <Navbar /> </nav>
   
      <div className="container">
        <Dashboard />
       
        <Outlet/>
      </div>
    </div>
  );
};

export default App;
