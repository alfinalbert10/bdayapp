import React from "react";
import "./MainContent.css";

const MainContent = () => {
  return (
    <div className="main-content">
      <h2>Here you can view the birthday list.</h2>
      <table className="birthday-table">
        <thead>
          <tr>
            <th><center>Name</center></th>
            <th><center>Date of Birth</center></th>
            <th><center>Actions</center></th>
          </tr>
        </thead>
        <tbody>
            <tr>
              <td></td>
              <td></td>
              <td>
                <center><button>
                  Edit
                </button>
                <button >Delete</button></center>
              </td>
            </tr>
      
        </tbody>
      </table>
     
    </div>
  );
};

export default MainContent;
