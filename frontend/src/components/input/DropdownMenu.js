import { useEffect, useRef, useState } from "react"
import { matchPath, useLocation, useNavigate } from "react-router-dom"
import { SVGDropdownMenuArrow } from "../svg/SVGs"
import { findRouteIndex } from "../../util/findRouteIndex"

function DropdownItem({ name, onClick }) {
    return (
        <li
            className="cursor-pointer px-4 py-3 flex items-center hover:bg-framebgcol border-grey-750"
            onClick={onClick}>
            <span className="truncate">{name}</span>
        </li>
    )
}

export function DefaultDropdownButton({
    label,
    isOpen,
    onClick,
    cls,
    clsOpen = "bg-grey-775 hover:bg-grey-750",
    clsClosed = "bg-grey-850 hover:bg-grey-825",
}) {
    return (
        <button
            className={`${cls} dropdown-button ${
                isOpen ? `${clsOpen}` : `${clsClosed}`
            }`}
            onClick={onClick}>
            <span className="truncate min-w-0">{label}</span>
            <SVGDropdownMenuArrow isOpen={isOpen} />
        </button>
    )
}

function DropdownMenu({ initSelected, itemNames, button: Button, onChange }) {
    const [selected, setSelected] = useState(initSelected)
    const [isOpen, setOpen] = useState(false)
    const dropdownRef = useRef(null)
    const items = itemNames.map((itemName, index) => (
        <DropdownItem
            index={index}
            name={itemName}
            key={index}
            onClick={() => {
                if (onChange) {
                    onChange(index)
                }
                setOpen(false)
                setSelected(index)
            }}
        />
    ))
    useEffect(() => {
        setSelected(initSelected || -1)
    }, [initSelected])

    useEffect(() => {
        const handleClickOutside = (event) => {
            if (
                dropdownRef.current &&
                !dropdownRef.current.contains(event.target)
            ) {
                setOpen(false)
            }
        }
        document.addEventListener("click", handleClickOutside)
        return () => {
            document.removeEventListener("click", handleClickOutside)
        }
    }, [])

    Button = Button || DefaultDropdownButton
    return (
        <div className="relative w-full" ref={dropdownRef}>
            <Button
                label={itemNames[selected === -1 ? 0 : selected]}
                isOpen={isOpen}
                onClick={() => setOpen(!isOpen)}
            />
            <div
                className={`z-10 absolute overflow-hidden top-12 inset-x-0 ${
                    isOpen ? "max-h-60 opacity-100" : "max-h-0 opacity-0"
                } transition-height-opacity duration-[250ms]`}>
                <div
                    className={`rounded-md max-h-60 overflow-y-auto border-bordefcol border`}>
                    <ul
                        className={`divide-y divide-grey-750 bg-grey-875 rounded-md overflow-hidden text-dropdown`}>
                        {items}
                    </ul>
                </div>
            </div>
        </div>
    )
}

export function DropdownRoutes({
    routes,
    routeLabels,
    button: Button,
    onChange,
}) {
    const navigate = useNavigate()
    const location = useLocation()
    const selected = findRouteIndex(routes, location.pathname)
    const handleChange = (index) => {
        if (index !== -1 && !matchPath(routes[index], location.pathname)) {
            navigate(routes[index])
        }
        if (onChange) {
            onChange(index)
        }
    }
    return (
        <DropdownMenu
            initSelected={selected}
            button={Button}
            itemNames={routeLabels}
            onChange={handleChange}
        />
    )
}

export default DropdownMenu
