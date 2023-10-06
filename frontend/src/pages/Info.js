import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { SVGTitleComponent } from "../components/container/RoundedFrame"
import ProfileSideBar from "../components/concrete/other/ProfileSidebar"
import RoundedTable from "../components/container/RoundedTable"

function InfoTable() {
    const { t } = useTranslation()
    const titleComponent = (
        <SVGTitleComponent svg={<FontAwesomeIcon icon="fa-info" className="w-4 h-4 mr-3" />} title={t("info.info")} />
    )
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-dividecol text-sm">
                <tr>
                    <td>{t("info.compiler_options")}</td>
                    <td>
                        <Link to="#" className="link">
                            options.pdf
                        </Link>
                    </td>
                </tr>
                <tr>
                    <td>{t("info.contests")}</td>
                    <td>
                        <Link to="#" className="link">
                            contests.pdf
                        </Link>
                    </td>
                </tr>
            </tbody>
        </RoundedTable>
    )
}

function Info({ data }) {
    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl lg:space-x-3 px-4">
                <ProfileSideBar />
                <div className="w-full min-w-0">
                    <InfoTable />
                </div>
            </div>
        </div>
    )
}

export default Info
