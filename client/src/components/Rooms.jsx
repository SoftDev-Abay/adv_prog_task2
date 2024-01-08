import { React, useEffect, useState } from "react";
import RoomCard from "./RoomCard";
import axios from "axios";
import { useCategoryContext } from "../contex/CategoryContex";
const Rooms = () => {
  const [rooms, setRooms] = useState([]);
  const { currentCategory } = useCategoryContext();

  const getRooms = async () => {
    const responce = await axios.get("http://localhost:3000/buildings", {
      headers: {
        "Content-Type": "application/json",
      },
    });
    console.log(responce.data);
    setRooms(responce.data);
  };

  useEffect(() => {
    getRooms();
  }, []);

  return (
    <div className="w-fit  mx-auto grid grid-cols-1 min-[400px]:grid-cols-2   sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 2xl:grid-cols-7 justify-items-center justify-center gap-7 mt-10">
      {rooms
        .filter(
          (room) =>
            room.category === currentCategory || currentCategory === "All"
        )
        .map((room) => {
          return <RoomCard key={room.id} room={room} />;
        })}
    </div>
  );
};

export default Rooms;
