import React from "react";
import {
  Routes,
  Route,
  HashRouter as Router,
} from "react-router-dom";
import { HomePage,Indexes,Others } from "../pages/index";
import {Navbar} from "./index";


const Navigation = () => {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route exact path="/" element={<HomePage />} />
        <Route exact path="/indexes" element={<Indexes />} />
        <Route exact path="/others" element={<Others />} />
      </Routes>
    </Router>
  );
};


export default Navigation;
