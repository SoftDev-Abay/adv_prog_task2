import { TbBeach, TbMountain, TbPool } from "react-icons/tb";
import {
  GiBarn,
  GiBoatFishing,
  GiCactus,
  GiCastle,
  GiCaveEntrance,
  GiForestCamp,
  GiIsland,
  GiWindmill,
} from "react-icons/gi";
import { FaSkiing } from "react-icons/fa";
import { BsSnow } from "react-icons/bs";
import { IoDiamond } from "react-icons/io5";
import { MdOutlineVilla } from "react-icons/md";
import { useCategoryContext } from "../contex/CategoryContex";

export const categories = [
  {
    label: "Beach",
    icon: TbBeach,
  },
  {
    label: "Windmills",
    icon: GiWindmill,
  },
  {
    label: "Modern",
    icon: MdOutlineVilla,
  },
  {
    label: "Countryside",
    icon: TbMountain,
  },
  {
    label: "Pools",
    icon: TbPool,
  },
  {
    label: "Islands",
    icon: GiIsland,
  },
  {
    label: "Lake",
    icon: GiBoatFishing,
  },
  {
    label: "Skiing",
    icon: FaSkiing,
  },
  {
    label: "Castles",
    icon: GiCastle,
  },
  {
    label: "Caves",
    icon: GiCaveEntrance,
  },
  {
    label: "Camping",
    icon: GiForestCamp,
  },
  {
    label: "Arctic",
    icon: BsSnow,
  },
  {
    label: "Desert",
    icon: GiCactus,
  },
  {
    label: "Barns",
    icon: GiBarn,
  },
  {
    label: "Lux",
    icon: IoDiamond,
  },
];

const Categories = () => {
  const { setCurrentCategory } = useCategoryContext();

  return (
    <div className="mt-5 flex justify-between overflow-auto gap-10 scrollbar-dark ">
      {categories.map((category) => {
        const { label, icon: Icon } = category;
        return (
          <div
            onClick={() => {
              setCurrentCategory(label);
            }}
            className="flex flex-col items-center text-gray-500 gap-1 cursor-pointer hover:text-gray-700"
          >
            <Icon className="text-4xl " />
            <span className="text-sm ">{label}</span>
          </div>
        );
      })}
    </div>
  );
};

export default Categories;
