import ProfileSideBar from '../components/ProfileSidebar'
import DropdownListFrame from '../components/DropdownListFrame'
import {matchPath} from "react-router-dom";
import {routeMap} from "../config/RouteConfig";

function Archive({ data }) {
    if (!data || data.processed) {
        return <></>
    }
    data.processed = true
    const categoriesContent = data.categories.map((item, index) =>
        <div className="mb-3" key={index}>
            <DropdownListFrame title={item.title} tree={{"children": item.children}} />
        </div>
    )
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3">
                    {categoriesContent}
                </div>
            </div>
        </div>
    );
}

export default Archive;