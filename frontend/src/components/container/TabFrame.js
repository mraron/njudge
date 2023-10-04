import { Link, useLocation } from "react-router-dom";
import { DropdownRoutes } from "../input/DropdownMenu";
import { findRouteIndex } from "../../util/findRouteIndex";

function Tab({ isSelected, label, route }) {
    return (
        <Link
            className={`block rounded-md px-4 py-2 text-nav ${
                isSelected ? "bg-framebgcol" : "hover:bg-grey-850"
            }`}
            to={route}>
            {label}
        </Link>
    );
}

function TabFrame({ routes, routeLabels, routePatterns, children }) {
    const location = useLocation();
    const selected = findRouteIndex(routePatterns, location.pathname);
    const tabsContent = routes.map((item, index) => (
        <li className="mr-1.5" key={index}>
            <Tab
                isSelected={index === selected}
                label={routeLabels[index]}
                route={item}
                key={index}
            />
        </li>
    ));
    return (
        <div className="w-full space-y-2">
            <ul className="hidden sm:flex">{tabsContent}</ul>
            <div className="block sm:hidden">
                <DropdownRoutes
                    label="Profil"
                    routes={routes}
                    routePatterns={routePatterns}
                    routeLabels={routeLabels}
                />
            </div>
            <div>{children}</div>
        </div>
    );
}

export default TabFrame;
