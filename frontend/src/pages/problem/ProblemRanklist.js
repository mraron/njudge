import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { useTranslation } from "react-i18next"
import Ranklist from "../../components/concrete/other/Ranklist"
import Pagination from "../../components/util/Pagination"

function ProblemRanklist({ data }) {
    const { t } = useTranslation()
    const navigate = useNavigate()

    useEffect(() => {
        if (!data.isPublic) {
            window.flash("flash.ranklist_not_public", "failure")
            navigate(-1)
        }
    })
    return (
        data.isPublic && (
            <div className="space-y-2">
                <Ranklist
                    ranklist={data.ranklist}
                    title={t("problem_ranklist.ranklist")}
                />
                <Pagination paginationData={data.paginationData} />
            </div>
        )
    )
}

export default ProblemRanklist
