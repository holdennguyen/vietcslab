"use client";

import { Category } from "@prisma/client";
import {
    FcEngineering,
    FcFilmReel,
    FcMultipleDevices,
    FcMusic,
    FcOldTimeCamera,
    FcSalesPerformance,
    FcSportsMode
} from "react-icons/fc";
import { IconType } from "react-icons";
import { CategoryItem } from "./category_items";


interface CategoriesProps {
    items: Category[];
}

const iconMap: Record<Category["name"], IconType> = {
    "Programming Fundamentals": FcMusic,
    "Software Development": FcOldTimeCamera,
    "Web Development": FcSportsMode,
    "Data Science & Machine Learning": FcSalesPerformance,
    "Cybersecurity": FcMultipleDevices,
    "Operating Systems & Networking": FcFilmReel,
    "Computer Science Theory": FcEngineering,
    "Career Guidance & Resources": FcSalesPerformance,
}

export const Categories = ({
    items,
}: CategoriesProps) => {
    return (
        <div className="flex items-center gap-x-2 overflow-x-auto pb-2">
            {items.map((item)=> (
                <CategoryItem
                    key={item.id}
                    label={item.name}
                    icon={iconMap[item.name]}
                    value={item.id}
                />
            ))}
        </div>
    )
}