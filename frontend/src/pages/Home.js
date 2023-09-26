import ProfileSideBar from "../components/concrete/other/ProfileSidebar";
import PostFrame from "../components/container/PostFrame";

function Home({ data }) {
    return (
        <div className="relative w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="ml-0 lg:ml-4">
                    <ProfileSideBar />
                </div>
                <div className="w-full px-4 lg:pl-3">
                    <PostFrame posts={data.posts} />
                </div>
            </div>
        </div>
    );
}

export default Home;
