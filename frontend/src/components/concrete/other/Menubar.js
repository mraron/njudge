import { useContext, useEffect, useRef, useState } from "react"
import { Link, useLocation } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { DefaultDropdownButton, DropdownRoutes } from "../../input/DropdownMenu"
import { findRouteIndex } from "../../../util/findRouteIndex"
import { routeMap } from "../../../config/RouteConfig"
import UserContext from "../../../contexts/user/UserContext"
import ThemeContext from "../../../contexts/theme/ThemeContext"

const menuRoutes = [
    routeMap.home,
    routeMap.contests,
    routeMap.archive,
    routeMap.submissions,
    routeMap.problems.replace(":problemset", "main"),
    routeMap.info,
]
const menuRouteLabels = [
    "menubar.home",
    "menubar.contests",
    "menubar.archive",
    "menubar.submissions",
    "menubar.problems",
    "menubar.information",
]

function MenuOption({ label, route, selected, horizontal, onClick }) {
    return (
        <li>
            <Link
                className={`flex items-center h-full px-4 ${horizontal ? "border-b-3 pt-1" : "border-l-3 p-3"} ${
                    selected ? "border-highlight bg-grey-775" : "border-transparent hover:bg-framebgcol"
                }`}
                onClick={onClick}
                to={route}>
                {label}
            </Link>
        </li>
    )
}

function getProfileDropdownButton(isLoggedIn) {
    function ProfileDropdownButton({ isOpen, onClick }) {
        const { t } = useTranslation()
        return (
            <DefaultDropdownButton
                label={isLoggedIn ? t("menubar.profile") : t("menubar.login")}
                isOpen={isOpen}
                onClick={onClick}
                cls="rounded-r-none"
            />
        )
    }
    return ProfileDropdownButton
}

function ProfileSettings({ onSidebarClose }) {
    const { userData, isLoggedIn } = useContext(UserContext)
    const { t, i18n } = useTranslation()

    let profileRoutes = [routeMap.login, routeMap.register]
    let profileRouteLabels = ["menubar.login", "menubar.register"]
    if (isLoggedIn) {
        profileRoutes = [routeMap.profile.replace(":user", encodeURIComponent(userData.username)), routeMap.logout]
        profileRouteLabels = ["menubar.profile", "menubar.logout"]
        if (userData.isAdmin) {
            profileRoutes = profileRoutes.slice(0, 1).concat([routeMap.admin]).concat(profileRoutes.slice(1))
            profileRouteLabels = profileRouteLabels
                .slice(0, 1)
                .concat(["menubar.admin"])
                .concat(profileRouteLabels.slice(1))
        }
    }
    function LanguageButton({ lang }) {
        return (
            <button
                className={`px-2 py-0.5 ${
                    i18n.resolvedLanguage === lang ? "bg-grey-725" : "hover:bg-grey-775"
                } rounded`}
                onClick={() => {
                    console.log("nigga")
                    i18n.changeLanguage(lang)
                }}>
                {lang}
            </button>
        )
    }
    return (
        <div className="w-full flex items-stretch">
            <DropdownRoutes
                button={getProfileDropdownButton(isLoggedIn)}
                routes={profileRoutes}
                onChange={onSidebarClose}
                routeLabels={profileRouteLabels.map(t)}
            />
            <div className="px-3 flex items-center justify-center border border-l-0 border-bordefcol space-x-1">
                <LanguageButton lang="hu" />
                <LanguageButton lang="en" />
            </div>
            <div>
                <ThemeButton />
            </div>
        </div>
    )
}

function ThemeButton() {
    const { theme, changeTheme } = useContext(ThemeContext)
    const toggleTheme = () => {
        if (theme === "light") {
            changeTheme("dark")
        } else {
            changeTheme("light")
        }
    }
    return (
        <button
            className="h-full border border-l-0 border-bordefcol flex items-center justify-center p-2 rounded-r-md hover:bg-framebgcol"
            onClick={toggleTheme}>
            <FontAwesomeIcon icon={theme === "light" ? "fa-moon" : "fa-sun"} className="text-iconcol w-6 h-4" />
        </button>
    )
}

function MenuSideBar({ selected, isOpen, onClose }) {
    const { t } = useTranslation()
    const menuRef = useRef(null)
    const menuOptions = menuRoutes.map((item, index) => {
        return (
            <MenuOption
                route={item}
                label={t(menuRouteLabels[index])}
                selected={index === selected}
                horizontal={false}
                onClick={onClose}
                key={index}
            />
        )
    })
    useEffect(() => {
        const handleClickOutside = (event) => {
            if (
                menuRef.current &&
                !menuRef.current.contains(event.target) &&
                event.target.id !== "menuButton" &&
                !event.target.closest("#menuButton")
            ) {
                onClose()
            }
        }
        document.addEventListener("click", handleClickOutside)
        return () => {
            document.removeEventListener("click", handleClickOutside)
        }
    }, [])

    return (
        <aside
            ref={menuRef}
            className={`h-full w-80 z-20 pt-20 overflow-y-auto overflow-x-hidden xl:hidden fixed bg-grey-825 border-l border-bordefcol right-0 transform ${
                isOpen ? "translate-x-0 opacity-100" : "translate-x-80 opacity-0"
            } ease-out transition-transform-opacity duration-200`}>
            <div className="flex flex-col justify-center">
                <div className="w-full flex px-4 mb-4">
                    <ProfileSettings onSidebarClose={onClose} />
                </div>
                <ol className="divide-y divide-grey-725 border-t border-b border-grey-725">{menuOptions}</ol>
            </div>
        </aside>
    )
}

function MenuTopBar({ selected, isOpen, onToggle }) {
    const { t } = useTranslation()
    const menuOptions = menuRoutes.map((item, index) => {
        return (
            <MenuOption
                route={item}
                label={t(menuRouteLabels[index])}
                selected={index === selected}
                horizontal={true}
                key={index}
            />
        )
    })
    return (
        <div className="z-30 flex justify-center bg-grey-825 border-b border-grey-700 fixed w-full top-0">
            <div className="w-full max-w-7xl flex justify-between items-center">
                <div className="flex w-full">
                    <Link to="/" className="font-semibold text-lg mx-8 my-4">
                        nJudge
                    </Link>
                    <ol className="hidden xl:flex">{menuOptions}</ol>
                    <div className="w-full hidden xl:flex justify-end mx-4 my-2">
                        <div className="w-72">
                            <ProfileSettings />
                        </div>
                    </div>
                </div>
                <div className="xl:hidden mx-4">
                    <button
                        id="menuButton"
                        className="flex items-center justify-center p-2 rounded-full hover:bg-framebgcol"
                        onClick={() => onToggle(this)}
                        aria-label={t("aria_label.toggle_menu")}>
                        {isOpen ? (
                            <FontAwesomeIcon icon="fa-close" className="w-5 h-5" />
                        ) : (
                            <FontAwesomeIcon icon="fa-bars" className="w-5 h-5" />
                        )}
                    </button>
                </div>
            </div>
        </div>
    )
}

function Menubar() {
    const { isLoggedIn } = useContext(UserContext)
    const location = useLocation()
    const selected = findRouteIndex(menuRoutes, location.pathname)
    const [isOpen, setOpen] = useState(false)
    const handleToggle = () => {
        setOpen((prevOpen) => !prevOpen)
    }
    const handleClose = () => {
        setOpen(false)
    }
    return (
        isLoggedIn !== null && (
            <div className="text-nav">
                <MenuTopBar selected={selected} isOpen={isOpen} onToggle={handleToggle} />
                <MenuSideBar selected={selected} isOpen={isOpen} onClose={handleClose} />
            </div>
        )
    )
}

export default Menubar
