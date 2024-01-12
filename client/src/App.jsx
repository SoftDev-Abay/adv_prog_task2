import { useState } from "react";
import Layout from "./components/Layout";
import Main from "./pages/Main";

import AuthContext from "./contex/AuthContext";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
    <>
      <AuthContext>
        <Layout>
          <Router>
            <Routes>
              <Route path="/" element={<Main />} />
            </Routes>
          </Router>
        </Layout>
      </AuthContext>
    </>
  );
}

export default App;
