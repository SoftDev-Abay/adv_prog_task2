import { useState } from "react";
import Navbar from "./components/Navbar";
import Categories from "./components/Categories";
import Rooms from "./components/Rooms";
import Footer from "./components/Footer";
import CategoryContex from "./contex/CategoryContex";
import AddBuildingModal from "./modals/AddBuildingModal";

function App() {
  const [isOpenAddBuildingModal, setIsOpenAddBuildingModal] = useState(false);
  return (
    <>
      {isOpenAddBuildingModal && (
        <AddBuildingModal
          setIsOpenAddBuildingModal={setIsOpenAddBuildingModal}
        />
      )}
      <Navbar setIsOpenAddBuildingModal={setIsOpenAddBuildingModal} />
      <main className="px-7 min-[500px]:px-20  mx-auto mt-5 mb-10">
        <CategoryContex>
          <Categories />
          <Rooms />
        </CategoryContex>
      </main>
      <Footer />
    </>
  );
}

export default App;
