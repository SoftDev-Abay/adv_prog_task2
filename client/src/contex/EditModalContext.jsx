import React, { createContext, useContext, useState } from "react";
import AddBuildingModal from "../modals/AddBuildingModal";

const Context = createContext();

const ModalsContext = ({ children }) => {
  const [isOpenAddBuildingModal, setIsOpenAddBuildingModal] = useState(false);
  return (
    <Context.Provider
      value={{
        isOpenAddBuildingModal,
        setIsOpenAddBuildingModal,
      }}
    >
      {isOpenAddBuildingModal && <AddBuildingModal />}
      {children}
    </Context.Provider>
  );
};

export default ModalsContext;

export const useModalContext = () => useContext(Context);
