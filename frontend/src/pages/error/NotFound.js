import { Link } from "react-router-dom"
import { useTranslation } from "react-i18next"
import { SVGNotFound } from "../../components/svg/SVGs"
import RoundedFrame from "../../components/container/RoundedFrame"
import Button from "../../components/basic/Button"
import NarrowPage from "../wrappers/NarrowPage"

function NotFoundFrame() {
    const { t } = useTranslation()
    return (
        <RoundedFrame title={t("not_found.title")}>
            <div className="px-10 py-8 flex flex-col relative justify-between">
                <p className="z-10">{t("not_found.message")}</p>
                <div className="flex justify-center absolute inset-0">
                    <SVGNotFound />
                </div>
                <div className="flex justify-center mt-8">
                    <Link className="z-10" to="/">
                        <Button color="indigo" minWidth="15rem">
                            {t("not_found.back_to_home")}
                        </Button>
                    </Link>
                </div>
            </div>
        </RoundedFrame>
    )
}

function NotFound() {
    return (
        <NarrowPage>
            <NotFoundFrame />
        </NarrowPage>
    )
}

export default NotFound
