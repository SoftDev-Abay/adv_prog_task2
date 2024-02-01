import { React, useEffect, useState } from "react";
import RoomCard from "./RoomCard";
import axios from "axios";
import { useCategoryContext } from "../contex/CategoryContex";
import Pagination from "./Pagination";
const Rooms = () => {
  const limit = 4;
  const countVisiblePages = 5;

  const [rooms, setRooms] = useState([]);
  const { currentCategory } = useCategoryContext();
  const [activePage, setActivePage] = useState(1);
  const [pagesRange, setPagesRange] = useState({
    start: 1,
    end: countVisiblePages,
  });
  const [countBuildings, setCountBuildings] = useState(0);

  const getRooms = async (page = 1) => {
    const responce = await axios.get(
      `http://localhost:3000/buildings/page?page=${page}&limit=${limit}`,
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    const { data } = responce;
    console.log(data.buildings);
    setRooms(data.buildings);
    setCountBuildings(data.count);
  };

  useEffect(() => {
    getRooms();
  }, []);

  useEffect(() => {
    const countPages = Math.ceil(countBuildings / limit);
    if (countPages < countVisiblePages && countPages != 0) {
      setPagesRange({
        start: 1,
        end: countPages,
      });
    }
  }, [countBuildings]);

  return (
    <>
      <div className="mb-16 w-fit  mx-auto grid grid-cols-1 min-[400px]:grid-cols-2   sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 2xl:grid-cols-7 justify-items-center justify-center gap-7 mt-10">
        {rooms
          .filter(
            (room) =>
              room.category === currentCategory || currentCategory === "All"
          )
          .map((room) => {
            return <RoomCard key={"RoomCard" + room.id} room={room} />;
          })}
      </div>
      <Pagination
        activePage={activePage}
        countVisiblePages={countVisiblePages}
        limit={limit}
        countBuildings={countBuildings}
        pagesRange={pagesRange}
        setActivePage={setActivePage}
        setPagesRange={setPagesRange}
        getRooms={getRooms}
      />
    </>
  );
};

export default Rooms;
