import ProfileSideBar from "../../components/concrete/other/ProfileSidebar"

function ProfileSidebarPage({ children }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-3">
                <ProfileSideBar />
                <div className="w-full min-w-0">{children}</div>
            </div>
        </div>
    )
}

export default ProfileSidebarPage
