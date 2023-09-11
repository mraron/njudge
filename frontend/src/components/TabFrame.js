 
import { Link, useLocation } from 'react-router-dom';
import { DropdownRoutes } from './DropdownMenu';
import { trimRoute } from '../util/route';

function Tab({ isSelected, name, route }) {
    return (
        <Link className={`block rounded-md px-4 py-2 ${isSelected? "bg-grey-800": "hover:bg-grey-850"} mr-1.5`} to={route}>
            {name}
        </Link>
    )
}

function TabFrame({ routes, children }) {
	const location = useLocation()
	const selected = routes.map(pair => trimRoute(pair[1])).indexOf(trimRoute(location.pathname));
    const tabs = routes.map((pair, index) => <Tab isSelected={index === selected} name={pair[0]} route={pair[1]} />)
    return (
        <div className="w-full">
            <ul className="hidden sm:flex mb-2">
                {tabs}
            </ul>
            <div className="block sm:hidden mb-2">
                <DropdownRoutes label="Profil" routes={routes} />
            </div>
            <div>
                {children}
            </div>
        </div>
    );
}

export default TabFrame;
