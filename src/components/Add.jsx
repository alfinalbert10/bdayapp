import React from "react";
import Form from 'react-bootstrap/Form';
import "./Add.css";
import { NavLink } from "react-router-dom";
const Add = () => {
    return (
        <div id="cont" style={{width:"500px",height:"50px",marginLeft:"350px"}}>
        <Form method="post" id="contact-form" >
        <p>
          <span>Name</span>
          <input
            placeholder="Enter Name"
            aria-label="Full name"
            type="text"
            name="full name"
         
          />
        </p>
        <label>
          <span>Date Of Birth</span>
          <input
            type="date"
            name="dob"
        
          />
        </label>
        <label>
          <span>Phone Number</span>
          <input
            type="text"
            name="phone"
          
          />
        </label>
       
          <button type="submit">Save</button>
          <NavLink to={"/view"}><button type="button">Cancel</button></NavLink>
       
      </Form>
      </div>
    );
  };
  
  export default Add;