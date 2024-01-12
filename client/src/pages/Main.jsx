import { React, useState } from "react";
import AddBuildingModal from "../modals/AddBuildingModal";
import CategoryContex from "../contex/CategoryContex";
import Categories from "../components/Categories";
import Rooms from "../components/Rooms";
import { useAuthContext } from "../contex/AuthContext";

function Main() {
  return (
    <>
      <CategoryContex>
        <Categories />
        <Rooms />
      </CategoryContex>
    </>
  );
}

export default Main;
