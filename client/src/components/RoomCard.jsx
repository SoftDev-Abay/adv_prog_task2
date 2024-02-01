import React from "react";
import { useModalContext } from "../contex/ModalsContext";
const RoomCard = ({ room }) => {
  const { price_day, country, city, img_url, category } = room;
  const {setIsOpenEditModal} = useModalContext()
  // console.log(room)
  return (
    <div className="cursor-pointer" onClick={()=>{
      setIsOpenEditModal(room.id)
    }}>
      <img
        className="rounded-xl w-[250px] h-[250px]  min-[400px]:w-[200px] min-[400px]:h-[200px] "
        src={img_url}
        alt=""
      />
      <div className="mt-2 flex flex-col gap-2">
        <span className="flex-grow font-semibold text-lg">
          {city}, {country}
        </span>
        <span className="flex-grow opacity-30 text-base">{category}</span>
        <span className="flex-grow font-semibold">
          $ {price_day} <span className="font-normal">night</span>
        </span>
      </div>
    </div>
  );
};

export default RoomCard;
