import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { SVGTitleComponent } from "../components/container/RoundedFrame"
import RoundedTable from "../components/container/RoundedTable"
import ProfileSidebarPage from "./wrappers/ProfileSidebarPage"

function InfoTable() {
    const { t } = useTranslation()
    const titleComponent = (
        <SVGTitleComponent icon={<FontAwesomeIcon icon="fa-info" className="w-4 h-4 mr-3" />} title={t("info.info")} />
    )
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-dividecol">
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
        <ProfileSidebarPage>
            <InfoTable />
        </ProfileSidebarPage>
    )
}

export default Info
