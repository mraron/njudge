import ProfileSideBar from '../components/ProfileSidebar'
import DropdownListFrame from '../components/DropdownListFrame'
import {useEffect, useState} from "react";

function Category({ title, children }) {
    return (
        <div className="mb-3">
            <DropdownListFrame title={title} tree={{"children": children}} />
        </div>
    )
}

function Archive() {
    const [data, setData] = useState(null)
    useEffect(() => {
        fetch("/api/v2/archive")
            .then(res => res.json())
            .then(data => setData(data))
    }, []);
    const categoriesContent = data? data["categories"].map((category, index) =>
        <Category title={category.title} children={category.children} index={index} />
    ): null
    return (
        <div className="w-full flex justify-center">
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