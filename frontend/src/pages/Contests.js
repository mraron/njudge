import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import RoundedFrame from "../components/container/RoundedFrame"
import Button from "../components/basic/Button"
import ProfileSidebarPage from "./wrappers/ProfileSidebarPage"

function ContestFrame({ contest }) {
    const { t } = useTranslation()
    const { name, href, date, active } = contest
    return (
        <RoundedFrame>
            <div className="px-8 py-6 sm:px-10 sm:py-8">
                <div className="flex justify-between items-start space-x-4 mb-2">
                    <span className="text-base emph-strong break-words min-w-0">{name}</span>
                    <span className="date-label">{date}</span>
                </div>
                <div className="flex space-x-2">
                    <Link to={href}>
                        <Button color="gray">{t("contests.view")}</Button>
                    </Link>
                    {active && <Button color="indigo">{t("contests.register")}</Button>}
                </div>
            </div>
        </RoundedFrame>
    )
}

function Contests({ data }) {
    const contestsContent = data.contests.map((item, index) => <ContestFrame key={index} contest={item} />)
    return (
        <ProfileSidebarPage>
            <div className="space-y-3">{contestsContent}</div>
        </ProfileSidebarPage>
    )
}

export default Contests
