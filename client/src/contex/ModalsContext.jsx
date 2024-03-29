import React, { createContext, useContext, useState } from "react";
import EditBuildingModal from "../modals/EditBuildingModal";
const Context = createContext();

const ModalsContext = ({ children }) => {
  const [isOpenEditModal, setIsOpenEditModal] = useState(false);

  return (
    <Context.Provider
      value={{
        isOpenEditModal,
        setIsOpenEditModal,
      }}
    >
      {isOpenEditModal && <EditBuildingModal id={isOpenEditModal} setIsOpenEditModal = {setIsOpenEditModal}/>}
      {children}
    </Context.Provider>
  );
};

export default ModalsContext;

export const useModalContext = () => useContext(Context);
