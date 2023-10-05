import ProfileSideBar from "../components/concrete/other/ProfileSidebar"
import PostFrame from "../components/container/PostFrame"
import React from "react"

function Home({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-4">
                <ProfileSideBar />
                <div className="w-full min-w-0">
                    <PostFrame posts={data.posts} />
                </div>
            </div>
        </div>
    )
}

export default Home
