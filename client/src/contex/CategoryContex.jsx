import React, { createContext, useContext, useState } from "react";

const Context = createContext();

const CategoryContex = ({ children }) => {
  const [currentCategory, setCurrentCategory] = useState("All");

  return (
    <Context.Provider
      value={{
        currentCategory,
        setCurrentCategory,
      }}
    >
      {children}
    </Context.Provider>
  );
};

export default CategoryContex;

export const useCategoryContext = () => useContext(Context);
