import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import SVGTitleComponent from "../../svg/SVGTitleComponent"
import RoundedFrame from "../../container/RoundedFrame"

function RanklistNotPublic() {
    const { t } = useTranslation()
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon
                    icon="fa-xmark"
                    className="w-5 h-5 highlight-red mr-3"
                />
            }
            title={t("ranklist_not_public.title")}
        />
    )
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-10 py-8 flex flex-col relative justify-between">
                <p className="z-10">{t("ranklist_not_public.message")}</p>
            </div>
        </RoundedFrame>
    )
}

export default RanklistNotPublic
