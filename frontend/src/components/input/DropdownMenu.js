import {useEffect, useRef, useState} from "react";
import {matchPath, useLocation, useNavigate} from "react-router-dom";
import {SVGDropdownMenuArrow} from "../../svg/SVGs";
import {findRouteIndex} from "../../util/findRouteIndex";

function DropdownItem({name, onClick}) {
    return (
        <li className="cursor-pointer px-4 py-3 flex items-center hover:bg-grey-800 border-grey-750" onClick={onClick}>
            <span className="truncate">{name}</span>
        </li>
    );
}

function DefaultDropdownButton({label, isOpen, onClick}) {
    return (
        <button
            className={`w-full rounded-md px-3 py-2 border-1 flex items-center justify-between ${isOpen ? "bg-grey-750 hover:bg-grey-700 border-grey-650" : "bg-grey-825 hover:bg-grey-775 border-default"} transition duration-150`}
            onClick={onClick}>
            <span className="overflow-ellipsis overflow-hidden">{label}</span>
            <SVGDropdownMenuArrow isOpen={isOpen} cls="shrink-0"/>
        </button>
    )
}

function DropdownMenu({initSelected, itemNames, button: Button, onChange}) {
    const [selected, setSelected] = useState(initSelected);
    const [isOpen, setOpen] = useState(false);
    const dropdownRef = useRef(null);
    const items = itemNames.map((itemName, index) =>
        <DropdownItem index={index} name={itemName} key={index} onClick={() => {
            if (onChange) {
                onChange(index);
            }
            setOpen(false);
            setSelected(index);
        }}/>
    );
    useEffect(() => {
        setSelected(initSelected || -1)
    }, [initSelected])

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
                setOpen(false);
            }
        };
        document.addEventListener('click', handleClickOutside);
        return () => {
            document.removeEventListener('click', handleClickOutside);
        };
    }, []);

    Button = Button || DefaultDropdownButton;
    return (
        <div className="relative w-full" ref={dropdownRef}>
            <Button label={itemNames[selected === -1 ? 0 : selected]} isOpen={isOpen} onClick={() => setOpen(!isOpen)}/>
            <div
                className={`z-10 absolute overflow-hidden top-12 inset-x-0 ${isOpen ? 'max-h-60 opacity-100' : 'max-h-0 opacity-0'} transition-all duration-[250ms]`}>
                <div className={`rounded-md max-h-60 overflow-y-auto border-default border-1`}>
                    <ul className={`divide-y divide-default bg-grey-875 rounded-md`}>
                        {items}
                    </ul>
                </div>
            </div>
        </div>
    )
}

export function DropdownRoutes({routes, routeLabels, button: Button}) {
    const navigate = useNavigate();
    const location = useLocation();
    const selected = findRouteIndex(routes, location.pathname)
    const handleChange = (index) => {
        if (index !== -1 && !matchPath(routes[index], location.pathname)) {
            navigate(routes[index])
        }
    }
    return (
        <DropdownMenu initSelected={selected} button={Button} itemNames={routeLabels} onChange={handleChange}/>
    )
}

export default DropdownMenu;