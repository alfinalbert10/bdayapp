import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import MainContent from './components/Home.jsx';
import ErrorPage from "./components/error-page.jsx";
import Add from "./components/Add.jsx";
const router = createBrowserRouter([
  {
    path: "/",
    element: <App/>,
    
    errorElement: <ErrorPage />,
    children: [
      {
        path: "view",
        element: <MainContent/>,
      },
      {
        path: "add",
        element: <Add/>,
      },
    ],
  },

]);

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
