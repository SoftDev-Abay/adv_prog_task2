import { React, useState } from "react";
import Footer from "./Footer";
import Navbar from "./Navbar";
import AddBuildingModal from "../modals/AddBuildingModal";
import LoginModal from "../modals/LoginModal";

const Layout = ({ children }) => {
  const [isOpenAddBuildingModal, setIsOpenAddBuildingModal] = useState(false);
  const [isOpenLoginModal, setIsOpenLoginModal] = useState(false);
  return (
    <>
      {isOpenAddBuildingModal && (
        <AddBuildingModal
          setIsOpenAddBuildingModal={setIsOpenAddBuildingModal}
        />
      )}
      {isOpenLoginModal && (
        <LoginModal setIsOpenLoginModal={setIsOpenLoginModal} />
      )}

      <Navbar
        setIsOpenAddBuildingModal={setIsOpenAddBuildingModal}
        setIsOpenLoginModal={setIsOpenLoginModal}
      />
      <main className="px-7 min-[500px]:px-20  mx-auto mt-5 mb-10">
        {children}
      </main>
      <Footer />
    </>
  );
};

export default Layout;
