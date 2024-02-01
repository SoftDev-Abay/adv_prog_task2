import React from "react";
import { BiSolidArrowToRight, BiSolidArrowToLeft } from "react-icons/bi";
import { IoArrowBackSharp, IoArrowForwardSharp } from "react-icons/io5";

const Pagination = ({
  countVisiblePages,
  limit,
  countBuildings,
  pagesRange,
  activePage,
  setActivePage,
  setPagesRange,
  getRooms,
}) => {
  const { start, end } = pagesRange;
  const countPages = Math.ceil(countBuildings / limit);

  const changePageHandlier = (newPage) => {
    if (newPage == start && newPage != 1) decreaseVisiblePages();
    else if (newPage == end && newPage != countPages) increaseVisiblePages();
    setActivePage(newPage);
    getRooms(newPage);
  };
  const increaseVisiblePages = () => {
    let startPage = pagesRange.start + countVisiblePages;
    let endPage = pagesRange.end + countVisiblePages;

    if (endPage > countPages) endPage = countPages;

    setPagesRange([startPage, endPage]);
  };

  const decreaseVisiblePages = () => {
    let startPage = pagesRange.start - countVisiblePages;
    let endPage = pagesRange.end - countVisiblePages;

    if (startPage < 1) startPage = 1;
    setPagesRange([startPage, endPage]);
  };

  const getPagesNumArr = () =>
    Array.from({ length: end - start + 1 }, (_, i) => start + i);

  return (
    <div className="w-full flex justify-center items-center">
      <div className="flex items-center gap-4">
        <div className="text-slate-500 cursor-pointer rounded-full flex justify-center items-center text-2xl w-12 h-12 bg-[$99C2FF]">
          <IoArrowBackSharp
            dis
            onClick={() => {
              if (activePage > 1) changePageHandlier(activePage - 1);
            }}
          />
        </div>
        <div className="flex gap-3">
          {getPagesNumArr().map((num) => (
            <PageNavItem
              key={"PageNavItem" + num}
              number={num}
              isActive={activePage == num}
              changePage={changePageHandlier}
            />
          ))}
        </div>
        <div className="text-slate-500 cursor-pointer rounded-full flex justify-center items-center text-2xl  w-12 h-12 ">
          <IoArrowForwardSharp
            onClick={() => {
              if (activePage < countPages) changePageHandlier(activePage + 1);
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default Pagination;

const PageNavItem = ({ number, changePage, isActive = false }) => {
  const fomattedNumber = number < 10 ? "0" + number : toString(number);
  return (
    <div
      className={`rounded-full  p-5 flex cursor-pointer justify-center items-center w-12 h-12 ${
        isActive ? "bg-slate-200 font-semibold" : "hover:bg-slate-50"
      }`}
      onClick={() => {
        changePage(number);
      }}
    >
      {number}
    </div>
  );
};
