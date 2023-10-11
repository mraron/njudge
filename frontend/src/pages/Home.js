import PostFrame from "../components/container/PostFrame"
import React from "react"
import ProfileSidebarPage from "./wrappers/ProfileSidebarPage"

function Home({ data }) {
    return (
        <ProfileSidebarPage>
            <PostFrame posts={data.posts} />
        </ProfileSidebarPage>
    )
}

export default Home
