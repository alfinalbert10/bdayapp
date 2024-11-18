import React from "react";
import "./Dashboard.css";
import { Outlet, Link } from "react-router-dom";

const Dashboard = () => {
  return (
    <div className="dashboard">
       <button id="btn" style={{width:"150px",height:"40px"}}> <Link to={`view`}>View</Link></button>
      <button id="btn" style={{width:"150px",height:"40px"}}><Link to={`add`}>Add</Link></button>
     
    </div>
  );
};

export default Dashboard;
