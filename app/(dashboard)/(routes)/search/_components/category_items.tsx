"use client";

import qs from "query-string";
import { cn } from "@/lib/utils";
import { 
    usePathname, 
    useRouter,
    useSearchParams,
} from "next/navigation";
import { IconType } from "react-icons";

interface CategoryItemProps {
    label: string;
    icon?: IconType;
    value?: string;
};

export const CategoryItem = ({
    label,
    value,
    icon: Icon,
}: CategoryItemProps) => {
    const pathname = usePathname();
    const router = useRouter();
    const searchParams = useSearchParams();

    const currentCategory = searchParams.get("categoryId");
    const currentTitle = searchParams.get("title");
    
    const isSelected = currentCategory === value;

    const onClick = () => {
        const url = qs.stringifyUrl({
            url: pathname,
            query: {
                title: currentTitle,
                categoryId: isSelected ? null : value,
            }
        }, { skipNull: true, skipEmptyString: true });

        router.push(url);
    };

    return (
        <button
        onClick={onClick}
            className={cn(
                "py-2 px-3 text-sm border border-slate-200 rounded-full flex items-center gap-x-1 hover:border-slate-700 transition",
                isSelected && "border-slate-700 bg-slate-200/20 text-slate-800"
            )}
            type="button"
        >
            {Icon && <Icon size={20} />}
            <div className="truncate">
                {label}
            </div>
        </button>
    )
}