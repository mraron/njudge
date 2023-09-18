import {Link, useLocation} from 'react-router-dom';
import {useContext, useEffect, useRef, useState} from 'react';
import {DropdownRoutes} from '../../input/DropdownMenu';
import {SVGClose, SVGDropdownMenuArrow, SVGHamburger} from '../../../svg/SVGs';
import {findRouteIndex} from '../../../util/FindRouteIndex';
import {routeMap} from "../../../config/RouteConfig";
import UserContext from "../../../contexts/user/UserContext";
import LanguageContext from "../../../contexts/language/LanguageContext";

const menuRoutes = [
    routeMap.main,
    routeMap.contests,
    routeMap.archive,
    routeMap.submissions,
    routeMap.problems,
    routeMap.info
]
const menuRouteLabels = [
    "Home",
    "Contests",
    "Archive",
    "Submissions",
    "Problems",
    "Information"
]
const translate = {
    "hu": {
        "Home": "Főoldal",
        "Contests": "Versenyek",
        "Archive": "Archívum",
        "Submissions": "Beküldések",
        "Problems": "Feladatok",
        "Information": "Tudnivalók",
    },
    "en": {
        "Home": "Home",
        "Contests": "Contests",
        "Archive": "Archive",
        "Submissions": "Submissions",
        "Problems": "Problems",
        "Information": "Information",
    }
}

const profileRouteLabels = {
    "loggedIn": [
        "Profil",
        "Kilépés"
    ],
    "loggedOut": [
        "Belépés",
        "Regisztráció"
    ]
}

function MenuOption({label, route, selected, horizontal, onClick}) {
    return (
        <li>
            <Link onClick={onClick}
                  className={`flex items-center h-full px-4 ${horizontal ? "border-b-3 pt-1" : "border-l-3 p-3"} ${selected ? "border-indigo-600 bg-grey-775" : "border-transparent hover:bg-grey-800"}`}
                  to={route}>
                {label}
            </Link>
        </li>
    )
}

function getProfileDropdownButton(isLoggedIn) {
    function ProfileDropdownButton({isOpen, onClick}) {
        return (
            <button
                className={`border-1 border-grey-675 rounded-tl-md rounded-bl-md flex items-center justify-between px-3 py-2 w-full h-full ${isOpen ? "bg-grey-750 hover:bg-grey-700" : "hover:bg-grey-800"}`}
                onClick={onClick}>
                <span className="text-left">
                    {isLoggedIn ? "Profil" : "Belépés"}
                </span>
                <SVGDropdownMenuArrow isOpen={isOpen}/>
            </button>
        );
    }

    return ProfileDropdownButton
}

function ProfileSettings() {
    const {userData, isLoggedIn} = useContext(UserContext)
    const {language, storeLanguage} = useContext(LanguageContext)
    const loginStr = isLoggedIn ? "loggedIn" : "loggedOut"

    let profileRoutes = [
        routeMap.login,
        routeMap.register
    ]
    if (isLoggedIn) {
        profileRoutes = [
            routeMap.profile.replace(":user", encodeURIComponent(userData.username)),
            routeMap.logout
        ]
    }
    return (
        <div className="flex">
            <DropdownRoutes button={getProfileDropdownButton(isLoggedIn)} routes={profileRoutes}
                            routeLabels={profileRouteLabels[loginStr]}/>
            <div
                className="px-3 flex items-center justify-center border-1 border-l-0 border-grey-675 rounded-tr-md rounded-br-md">
                <button className={`px-2 ${language === "hu"? "bg-grey-725": "hover:bg-grey-775"} rounded mr-1`} onClick={() => storeLanguage("hu")}>hu</button>
                <button className={`px-2 ${language === "en"? "bg-grey-725": "hover:bg-grey-775"} rounded`} onClick={() => storeLanguage("en")}>en</button>
            </div>
        </div>
    );
}

function MenuSideBar({selected, isOpen, onClose}) {
    const {language} = useContext(LanguageContext)
    const menuRef = useRef(null)
    const menuOptions = menuRoutes.map((item, index) => {
        return (
            <MenuOption label={translate[language][menuRouteLabels[index]]} route={item} selected={index === selected} horizontal={false}
                        key={index} onClick={onClose}/>
        );
    });
    useEffect(() => {
        const handleClickOutside = (event) => {
            if (menuRef.current && !menuRef.current.contains(event.target) && event.target.id !== "hamburgerButton" && !event.target.closest("#hamburgerButton")) {
                onClose()
            }
        };
        document.addEventListener('click', handleClickOutside);
        return () => {
            document.removeEventListener('click', handleClickOutside);
        };
    }, []);

    return (
        <aside ref={menuRef}
               className={`z-20 h-full overflow-y-auto lg:hidden fixed right-0 bg-grey-825 border-l-1 border-default ${isOpen ? "w-72 opacity-100" : "w-0 opacity-0"} ease-in-out transition-all duration-200`}>
            <div className="p-3">
                <button className="rounded-full p-3 hover:bg-grey-800 transition duration-200" onClick={onClose}>
                    <SVGClose cls="w-4 h-4"/>
                </button>
            </div>
            <div className="flex flex-col justify-center">
                <div className="mx-4 mb-4">
                    <ProfileSettings/>
                </div>
                <ol className="divide-y divide-default border-t border-b border-grey-750">
                    {menuOptions}
                </ol>
            </div>
        </aside>
    );
}

function MenuTopBar({selected, onOpen}) {
    const {language} = useContext(LanguageContext)
    const menuOptions = menuRoutes.map((item, index) => {
        return (
            <MenuOption label={translate[language][menuRouteLabels[index]]} route={item} selected={index === selected} horizontal={true}
                        key={index}/>
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
                        <div className="w-60">
                            <ProfileSettings/>
                        </div>
                    </div>
                </div>
                <div className="lg:hidden mx-4">
                    <button id="hamburgerButton" className="rounded-full p-2 hover:bg-grey-800 transition duration-200"
                            onClick={() => onOpen(this)}>
                        <SVGHamburger/>
                    </button>
                </div>
            </div>
        </div>
    );
}

function Menubar() {
    const {isLoggedIn} = useContext(UserContext)
    const location = useLocation();
    const selected = findRouteIndex(menuRoutes, location.pathname)
    const [isOpen, setOpen] = useState(false);
    const handleClose = () => {
        setOpen(false);
    };
    const handleOpen = () => {
        setOpen(true);
    };
    if (isLoggedIn !== null) {
        console.log(isLoggedIn)
    }
    return (
        isLoggedIn !== null &&
        <div>
            <MenuTopBar selected={selected} onOpen={handleOpen}></MenuTopBar>
            <MenuSideBar selected={selected} isOpen={isOpen} onClose={handleClose}/>
        </div>
    );
}

export default Menubar;