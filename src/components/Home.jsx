import React, { useState, useEffect } from "react";
import "./MainContent.css";
import axios from 'axios';

const MainContent = () => {
  const [data, setData] = useState([])
  useEffect(()=>{
    axios.get('http://localhost:8080/api/users/')
    .then(res=> setData(res.data))
    .catch(err => console.log(err))
  }, [])

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
          {data.map((d,i)=>(
            <tr key={i}>
              <td>{d.Username}</td>
              <td>{d.Birthdate}</td>
              <td>
                <center><button>
                  Edit
                </button>
                <button >Delete</button></center>
              </td>
            </tr>
      ))}
        </tbody>
      </table>
     
    </div>
  );
};

export default MainContent;
