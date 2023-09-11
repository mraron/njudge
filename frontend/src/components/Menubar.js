import { Link, useLocation } from 'react-router-dom';
import { useState } from 'react';
import { DropdownRoutes } from './DropdownMenu';
import { SVGClose, SVGDropdownMenuArrow, SVGHamburger } from '../svg/SVGs';
import { trimRoute } from '../util/route';

const routes = [
    ["Főoldal", "/"], 
    ["Versenyek", "/contests/"], 
    ["Archívum", "/archive/"], 
    ["Beküldések", "/problemset/status/"], 
    ["Feladatok", "/problemset/main/"], 
    ["Tudnivalók", "/info/"]];

const profileRoutes = [
    ["Profil", "/user/profile/"],
    ["Kilépés", "/"],
]

function MenuOption({ name, route, selected, horizontal, onClick }) {
    return (
        <li>
            <Link onClick={onClick} className={`flex items-center h-full px-4 ${horizontal? "border-b-3 pt-1": "border-l-3 p-3"} ${selected? "border-indigo-500 bg-grey-775": "border-transparent hover:bg-grey-800"}`} to={route}>
                {name}
            </Link>
        </li>
    )
}

function ProfileDropdownButton({ isOpen, onClick }) {
    return (
        <button className={`border-1 border-grey-675 rounded-tl-md rounded-bl-md flex items-center justify-between px-3 py-2 min-w-32 w-full h-full ${isOpen? "bg-grey-750 hover:bg-grey-700": "hover:bg-grey-800"}`} onClick={onClick}>
            <span className="flex items-center">
                <span>Profil</span>
            </span>
            <SVGDropdownMenuArrow isOpen={isOpen} />
        </button>
    );
}

function ProfileSettings() {
    return (
        <div className="flex">
            <DropdownRoutes button={ProfileDropdownButton} routes={profileRoutes} />
            <div className="px-4 flex items-center justify-center border-1 border-l-0 border-grey-675 rounded-tr-md rounded-br-md">
                <button className="px-2 bg-grey-725 rounded-md mr-1">hu</button>
                <button className="px-2 hover:bg-grey-800 rounded-md">en</button>
            </div>
        </div>
    );
}

function MenuSideBar({ selected, isOpen, onClose }) {
    const menuOptions = routes.map((item, index) => {
        return (
            <MenuOption name={item[0]} route={item[1]} selected={index === selected} horizontal={false} key={index} onClick={onClose} />
        );
    });
    return (
        <aside className={`z-20 h-full overflow-hidden lg:hidden fixed right-0 bg-grey-825 border-l-1 border-default ${isOpen? "w-72 opacity-100": "w-0 opacity-0"} ease-in-out transition-all duration-200`}>
            <div className="p-3">
                <button className="rounded-full p-3 hover:bg-grey-800 transition duration-200" onClick={onClose}>
                    <SVGClose size="w-4 h-4" />
                </button>
            </div>
            <div className="flex flex-col justify-center">
                <div className="mx-4 mb-4">
                    <ProfileSettings />
                </div>
                <ol className="divide-y divide-default border-t border-b border-grey-750">
                    {menuOptions}
                </ol>
            </div>
        </aside>
    );
}

function MenuTopBar({ selected, onOpen }) {
    const menuOptions = routes.map((item, index) => {
        return (
            <MenuOption name={item[0]} route={item[1]} selected={index === selected} horizontal={true} key={index} />
        );
    });
    return (
        <div className="z-10 flex justify-center bg-grey-825 border-b-1 border-grey-725 fixed w-full top-0">
            <div className="w-full max-w-7xl flex justify-between items-center">
                <div className="flex w-full">
                    <Link to="/" className="font-semibold text-lg mx-8 my-4">nJudge</Link>
                    <ol className="hidden lg:flex">
                        {menuOptions}
                    </ol>
                    <div className="w-full hidden lg:flex justify-end mx-4 my-2">
                        <ProfileSettings />
                    </div>
                </div>
                <div className="lg:hidden mx-4">
                    <button className="rounded-full p-2 hover:bg-grey-800 transition duration-200" onClick={onOpen}>
                        <SVGHamburger />
                    </button>
                </div>
            </div>
        </div>
    );
}

function Menubar() {
	const location = useLocation();
	const selected = routes.map(pair => trimRoute(pair[1])).indexOf(trimRoute(location.pathname));
    
    const [isOpen, setOpen] = useState(false);
    const handleClose = () => {
        setOpen(false);
    };
    const handleOpen = () => {
        setOpen(true);
    };
    return (
        <div>
            <MenuTopBar selected={selected} onOpen={handleOpen}></MenuTopBar>
            <MenuSideBar selected={selected} isOpen={isOpen} onClose={handleClose} />
        </div>
    );
}

export default Menubar;