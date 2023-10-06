import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import ProfileSideBar from "../components/concrete/other/ProfileSidebar"
import RoundedFrame from "../components/container/RoundedFrame"
import Button from "../components/basic/Button"

function ContestFrame({ contest }) {
    const { t } = useTranslation()
    const { name, href, date, active } = contest
    return (
        <RoundedFrame>
            <div className="px-6 py-5 sm:px-10 sm:py-8">
                <div className="flex justify-between items-start space-x-4 mb-2">
                    <span className="text-base emph-strong break-words min-w-0">
                        {name}
                    </span>
                    <span className="date-label">{date}</span>
                </div>
                <div className="flex space-x-2">
                    <Link to={href}>
                        <Button color="gray">{t("contests.view")}</Button>
                    </Link>
                    {active && (
                        <Button color="indigo">{t("contests.register")}</Button>
                    )}
                </div>
            </div>
        </RoundedFrame>
    )
}

function ContestList({ contests }) {
    const contestsContent = contests.map((item, index) => (
        <ContestFrame key={index} contest={item} />
    ))
    return <div className="space-y-3">{contestsContent}</div>
}

function Contests({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-4">
                <ProfileSideBar />
                <div className="w-full min-w-0">
                    <ContestList contests={data.contests} />
                </div>
            </div>
        </div>
    )
}

export default Contests
