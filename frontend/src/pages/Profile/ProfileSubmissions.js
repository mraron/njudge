import Pagination from "../../components/Pagination";
import SubmissionsTable from "../../components/SubmissionsTable";
import {useEffect, useState} from "react";
import PageLoadingAnimation from "../../components/PageLoadingAnimation";
import ProfileSideBar from "../../components/ProfileSidebar";
import {useParams} from "react-router-dom";

function ProfileSubmissions() {
    const {user} = useParams()
    const [data, setData] = useState(null)

    useEffect(() => {
        fetch(`/api/v2/user/profile/${user}/submissions`)
            .then(res => res.json())
            .then(data => setData(data))
    }, []);
    let pageContent = <PageLoadingAnimation/>;
    if (data) {
        pageContent =
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar
                        src="https://st3.depositphotos.com/6672868/13701/v/450/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"
                        username="dbence"
                        score="2550"/>
                </div>
                <div className="w-full px-4 lg:pl-3 overflow-x-auto">
                    <div className="mb-2">
                        <SubmissionsTable submissions={data.submissions} />
                    </div>
                    <Pagination current={1000} last={2000} />
                </div>
            </div>
    }
    return (
        <div className="relative">
            {pageContent}
        </div>
    );
}

export default ProfileSubmissions;