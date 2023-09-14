 
import {Link, useLocation} from 'react-router-dom';
import { DropdownRoutes } from './DropdownMenu';
import {findRouteIndex} from '../util/RouteUtil';

function Tab({ isSelected, label, route }) {
    return (
        <Link className={`block rounded-md px-4 py-2 ${isSelected? "bg-grey-800": "hover:bg-grey-850"}`} to={route}>
            {label}
        </Link>
    )
}

function TabFrame({ routes, routeLabels, routePatterns, children }) {
	const location = useLocation()
    const selected = findRouteIndex(routePatterns, location.pathname)
    const tabsContent = routes.map((item, index) =>
        <div className="mr-1.5">
            <Tab isSelected={index === selected} label={routeLabels[index]} route={item} key={index} />
        </div>
    )
    return (
        <div className="w-full">
            <ul className="hidden sm:flex mb-2">
                {tabsContent}
            </ul>
            <div className="block sm:hidden mb-2">
                <DropdownRoutes label="Profil" routes={routes} routePatterns={routePatterns} routeLabels={routeLabels} />
            </div>
            <div>
                {children}
            </div>
        </div>
    );
}

export default TabFrame;
