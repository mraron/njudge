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
            <div className="flex flex-col w-full">
                <div className="mb-2">
                    <SubmissionsTable submissions={data.submissions} />
                </div>
                <Pagination current={1000} last={2000} />
            </div>
    }
    return (
        <div className="relative">
            {pageContent}
        </div>
    );
}

export default ProfileSubmissions;